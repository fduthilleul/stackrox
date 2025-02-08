package imageintegrations

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	clusterDatastore "github.com/stackrox/rox/central/cluster/datastore"
	"github.com/stackrox/rox/central/enrichment"
	"github.com/stackrox/rox/central/imageintegration/datastore"
	countMetrics "github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/central/reprocessor"
	"github.com/stackrox/rox/central/sensor/service/common"
	"github.com/stackrox/rox/central/sensor/service/pipeline"
	"github.com/stackrox/rox/central/sensor/service/pipeline/reconciliation"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/centralsensor"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/errox"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/openshift"
	"github.com/stackrox/rox/pkg/set"
	"github.com/stackrox/rox/pkg/tlscheck"
	"github.com/stackrox/rox/pkg/urlfmt"
)

var (
	log = logging.LoggerForModule()

	autogeneratedRegistriesDisabled = env.AutogeneratedRegistriesDisabled.BooleanSetting() || env.ManagedCentral.BooleanSetting()

	_ pipeline.Fragment = (*pipelineImpl)(nil)
)

// GetPipeline returns an instantiation of this particular pipeline
func GetPipeline() pipeline.Fragment {
	return NewPipeline(enrichment.ManagerSingleton(),
		datastore.Singleton(),
		clusterDatastore.Singleton(),
		reprocessor.Singleton(),
	)
}

// NewPipeline returns a new instance of Pipeline.
func NewPipeline(integrationManager enrichment.Manager,
	datastore datastore.DataStore,
	clusterDatastore clusterDatastore.DataStore,
	enrichAndDetectLoop reprocessor.Loop) pipeline.Fragment {
	return &pipelineImpl{
		integrationManager:  integrationManager,
		datastore:           datastore,
		clusterDatastore:    clusterDatastore,
		enrichAndDetectLoop: enrichAndDetectLoop,
	}
}

type pipelineImpl struct {
	integrationManager enrichment.Manager

	datastore           datastore.DataStore
	clusterDatastore    clusterDatastore.DataStore
	enrichAndDetectLoop reprocessor.Loop
}

func (s *pipelineImpl) Capabilities() []centralsensor.CentralCapability {
	return nil
}

func (s *pipelineImpl) Reconcile(ctx context.Context, clusterID string, storeMap *reconciliation.StoreMap) error {
	sourcedAutogenEnabled := features.SourcedAutogeneratedIntegrations.Enabled()
	globalPullAutogenEnabled := env.AutogenerateGlobalPullSecRegistries.BooleanSetting()

	if !sourcedAutogenEnabled && !globalPullAutogenEnabled {
		return nil
	}

	existingIDs := set.NewStringSet()

	integrations, err := s.datastore.GetImageIntegrations(ctx, &v1.GetImageIntegrationsRequest{})
	if err != nil {
		return errors.Wrap(err, "getting image integrations for reconciliation")
	}
	for _, integration := range integrations {
		// Skipping ECR image integrations since they are special and use AWS IAM.
		if integration.GetEcr() != nil {
			continue
		}
		if integration.GetClusterId() != clusterID {
			continue
		}

		if !sourcedAutogenEnabled && globalPullAutogenEnabled && !openshift.GlobalPullSecretIntegration(integration) {
			// Only integrations created from the OCP global pull secret are eligible
			// for reconciliation, ignore others.
			continue
		}

		existingIDs.Add(integration.GetId())
	}
	store := storeMap.Get((*central.SensorEvent_ImageIntegration)(nil))
	return reconciliation.Perform(store, existingIDs, "imageintegrations", func(id string) error {
		log.Debugf("Reconciling (removing) image integration %q", id)
		return s.datastore.RemoveImageIntegration(ctx, id)
	})
}

func (s *pipelineImpl) Match(msg *central.MsgFromSensor) bool {
	return msg.GetEvent().GetImageIntegration() != nil
}

// matchesDockerAuth matches Docker integrations based on their authentication
// attributes.
func matchesDockerAuth(new, old *storage.ImageIntegration) bool {
	return new.GetDocker().GetUsername() == old.GetDocker().GetUsername() &&
		new.GetDocker().GetPassword() == old.GetDocker().GetPassword()
}

// matchesECRAuth matches ECR integrations based on their Authentication Data.
func matchesECRAuth(new, old *storage.ImageIntegration) bool {
	newAuth := new.GetEcr().GetAuthorizationData()
	oldAuth := old.GetEcr().GetAuthorizationData()
	return newAuth == nil && oldAuth == nil ||
		newAuth.GetUsername() == oldAuth.GetUsername() &&
			newAuth.GetPassword() == oldAuth.GetPassword() &&
			newAuth.GetExpiresAt().AsTime().Equal(oldAuth.GetExpiresAt().AsTime())
}

// matchingFunc returns true if the new integration matches the old integration.
type matchingFunc func(new, old *storage.ImageIntegration) bool

// getMatchingImageIntegration returns the image integration that exists and should be updated
// the second return value
func (s *pipelineImpl) getMatchingImageIntegration(matches matchingFunc, auto *storage.ImageIntegration, existingIntegrations []*storage.ImageIntegration) (*storage.ImageIntegration, bool) {
	for _, existing := range existingIntegrations {
		if auto.GetName() != existing.GetName() {
			continue
		}

		// If there exists a registry with an auto-generated name, and the user has somehow managed
		// (most likely through the API) to mark it as non-autogenerated, we do not want to overwrite
		// it.
		if !existing.GetAutogenerated() {
			return nil, false
		}

		// At this point, we just want to see if we already have an exact match
		// if so then we don't want to reprocess everything for no change.
		// The cluster ID can only be different if a cluster with this name was deleted
		// and then another with the same name was created. In this case, we do want to update
		// the integration, if only for the sake of updating the cluster ID.
		if matches(auto, existing) && auto.GetType() == existing.GetType() && auto.GetClusterId() == existing.GetClusterId() {
			return nil, false
		}
		return existing, true
	}

	return nil, true
}

func parseEndpointForURL(endpoint string) string {
	url := urlfmt.FormatURL(endpoint, urlfmt.HTTPS, urlfmt.NoTrailingSlash)

	server := urlfmt.GetServerFromURL(url)
	if strings.HasSuffix(server, "docker.io") || strings.HasSuffix(server, "docker.io:443") {
		return "https://registry-1.docker.io"
	}

	scheme := urlfmt.GetSchemeFromURL(url)
	defaultScheme := urlfmt.HTTPS
	if scheme == "http" {
		defaultScheme = urlfmt.InsecureHTTP
	}
	return urlfmt.FormatURL(server, defaultScheme, urlfmt.NoTrailingSlash)
}

// setUpIntegrationParams sets up the integration parameters based on its type
// and returns them.
func setUpIntegrationParams(ctx context.Context, imageIntegration *storage.ImageIntegration, includePathInDescription bool) (description string, matches matchingFunc, err error) {
	switch config := imageIntegration.GetIntegrationConfig().(type) {
	case *storage.ImageIntegration_Docker:
		dockerCfg := config.Docker
		validTLS, tlsErr := tlscheck.CheckTLS(ctx, dockerCfg.GetEndpoint())
		if tlsErr != nil {
			log.Debugf("reaching out for TLS check to %s: %v", dockerCfg.GetEndpoint(), tlsErr)
			// Not enough evidence that we can skip TLS, so conservatively require it.
			dockerCfg.Insecure = false
		} else {
			dockerCfg.Insecure = !validTLS
		}

		origEndpoint := dockerCfg.GetEndpoint()
		dockerCfg.Endpoint = parseEndpointForURL(origEndpoint)
		description = dockerCfg.GetEndpoint()

		// Ensure the endpoint path is maintained in the description for sourced integrations so
		// that endpoints with paths (ie: quay.io/stackrox-io) do not overwrite each other.
		if includePathInDescription {
			description = urlfmt.FormatURL(origEndpoint, urlfmt.HTTPS, urlfmt.NoTrailingSlash)
		}

		matches = matchesDockerAuth
	case *storage.ImageIntegration_Ecr:
		description = config.Ecr.GetEndpoint()
		if description == "" {
			description = fmt.Sprintf("AWS account %s and region %s", config.Ecr.GetRegistryId(), config.Ecr.GetRegion())
		}
		matches = matchesECRAuth
	default:
		err = errors.Errorf("unsupported autogenerated integration type: %T", config)
	}
	return
}

func (s *pipelineImpl) legacyRun(ctx context.Context, imageIntegration *storage.ImageIntegration, matches matchingFunc) error {
	// Fetch existing integration and determine if we should update by matching its configuration type.
	existingIntegrations, err := s.datastore.GetImageIntegrations(ctx, &v1.GetImageIntegrationsRequest{})
	if err != nil {
		return errors.Wrap(err, "fetching all image integrations")
	}
	integrationToUpdate, shouldInsert := s.getMatchingImageIntegration(matches, imageIntegration, existingIntegrations)
	if !shouldInsert {
		return nil
	}

	// Update or create.
	if integrationToUpdate == nil {
		if _, err := s.datastore.AddImageIntegration(ctx, imageIntegration); err != nil {
			return errors.Wrap(err, "adding integration")
		}
		if err := s.integrationManager.Upsert(imageIntegration); err != nil {
			return errors.Wrap(err, "notifying of image integration update")
		}
		// Only when adding the integration the first time do we need to run processing
		// Central receives many updates from OpenShift about the image integrations due to service accounts
		// So we can assume the other credentials were valid up to this point.
		// Also, they will eventually be picked up within an hour.
		s.enrichAndDetectLoop.ShortCircuit()
		return nil
	}

	imageIntegration.Id = integrationToUpdate.GetId()
	imageIntegration.Name = integrationToUpdate.GetName()
	if err := s.integrationManager.Upsert(imageIntegration); err != nil {
		return errors.Wrap(err, "notifying of image integration update")
	}
	if err := s.datastore.UpdateImageIntegration(ctx, imageIntegration); err != nil {
		return errors.Wrap(err, "updating integration")
	}
	return nil
}

func (s *pipelineImpl) runRemove(ctx context.Context, id string) error {
	if err := s.integrationManager.Remove(id); err != nil {
		return errors.Wrap(err, "removing integration from manager")
	}
	return s.datastore.RemoveImageIntegration(ctx, id)
}

func sourcedImageIntegration(imageIntegration *storage.ImageIntegration) bool {
	if imageIntegration.GetSource() == nil {
		// Can't be a sourced integration if it has no source data.
		return false
	}

	if features.SourcedAutogeneratedIntegrations.Enabled() {
		// All integrations are sourced (that have a source field) when feature enabled.
		return true
	}

	isGlobalPullSecret := openshift.GlobalPullSecretIntegration(imageIntegration)
	if isGlobalPullSecret && env.AutogenerateGlobalPullSecRegistries.BooleanSetting() {
		// Only the OCP Global pull secret is treated as sourced if the feature is enabled.
		return true
	}

	return false
}

// Run runs the pipeline template on the input and returns the output.
// Action is currently always update.
func (s *pipelineImpl) Run(ctx context.Context, clusterID string, msg *central.MsgFromSensor, _ common.MessageInjector) error {
	// Ignore autogenerated registries if they are disabled
	if autogeneratedRegistriesDisabled {
		return nil
	}
	defer countMetrics.IncrementResourceProcessedCounter(pipeline.ActionToOperation(msg.GetEvent().GetAction()), metrics.ImageIntegration)

	// Set up the integration and update its parameters.
	imageIntegration := msg.GetEvent().GetImageIntegration()
	sourcedIntegration := sourcedImageIntegration(imageIntegration)

	if sourcedIntegration && msg.GetEvent().GetAction() == central.ResourceAction_REMOVE_RESOURCE {
		// Remove is only supported if the feature flag is enabled.
		return s.runRemove(ctx, imageIntegration.GetId())
	}

	// Extract the cluster name.
	clusterName, exists, err := s.clusterDatastore.GetClusterName(ctx, clusterID)
	if err != nil {
		return errors.Wrapf(err, "error getting cluster name for cluster ID: %s", clusterID)
	}
	if !exists {
		return errox.NotFound.Newf("cluster with id %q does not exist", clusterID)
	}
	imageIntegration.ClusterId = clusterID

	description, matches, err := setUpIntegrationParams(ctx, imageIntegration, sourcedIntegration)
	if err != nil {
		return errors.Wrapf(err, "setting up integration params for %q", imageIntegration.GetId())
	}

	if !sourcedIntegration {
		imageIntegration.Name = fmt.Sprintf("Autogenerated %s for cluster %s", description, clusterName)
		return s.legacyRun(ctx, imageIntegration, matches)
	}

	source := imageIntegration.GetSource()
	imageIntegration.Name = fmt.Sprintf("Autogenerated %s for cluster %s from %s/%s",
		description, clusterName, source.GetNamespace(), source.GetImagePullSecretName())

	integrationExisted, err := s.upsertImageIntegration(ctx, imageIntegration)
	if err != nil {
		return err
	}

	// Currently, we will only trigger a reprocess of deployments when an image integration is newly added instead of
	// updated. Reasoning behind this is that we currently do not have a way of scoping a reprocessing to specific
	// deployments based on properties (i.e. deployments from the same cluster, deployments using a specific image
	// integration). This combined with the frequent updates of image integrations on OpenShift and the fact that
	// deployments will be reprocessed in ~1hr, we will skip reprocessing on updates for now.
	if !integrationExisted {
		s.enrichAndDetectLoop.ShortCircuit()
	}

	return nil
}

func (s *pipelineImpl) OnFinish(_ string) {}

// upsertImageIntegration will add the given image integration within the datastore.
// It will return any errors during the creation and whether the integration existed previously.
func (s *pipelineImpl) upsertImageIntegration(ctx context.Context, imageIntegration *storage.ImageIntegration) (existed bool, err error) {
	_, existed, err = s.datastore.GetImageIntegration(ctx, imageIntegration.GetId())
	if err != nil {
		return existed, errors.Wrapf(err, "retrieving image integration %q", imageIntegration.GetId())
	}

	if _, err := s.datastore.AddImageIntegration(ctx, imageIntegration); err != nil {
		return existed, errors.Wrapf(err, "adding image integration %q", imageIntegration.GetId())
	}
	if err := s.integrationManager.Upsert(imageIntegration); err != nil {
		return existed, errors.Wrapf(err, "notifying of update for image integration %q", imageIntegration.GetId())
	}
	return existed, nil
}
