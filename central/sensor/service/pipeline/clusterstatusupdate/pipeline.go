package clusterstatusupdate

import (
	"context"

	"github.com/pkg/errors"
	cloudSourcesManager "github.com/stackrox/rox/central/cloudsources/manager"
	clusterDataStore "github.com/stackrox/rox/central/cluster/datastore"
	cveFetcher "github.com/stackrox/rox/central/cve/fetcher"
	"github.com/stackrox/rox/central/deploymentenvs"
	"github.com/stackrox/rox/central/sensor/service/common"
	"github.com/stackrox/rox/central/sensor/service/pipeline"
	"github.com/stackrox/rox/central/sensor/service/pipeline/reconciliation"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/pkg/centralsensor"
	"github.com/stackrox/rox/pkg/features"
)

var (
	_ pipeline.Fragment = (*pipelineImpl)(nil)
)

// Template design pattern. We define control flow here and defer logic to subclasses.
//////////////////////////////////////////////////////////////////////////////////////

// GetPipeline returns an instantiation of this particular pipeline
func GetPipeline() pipeline.Fragment {
	return NewPipeline(clusterDataStore.Singleton(), deploymentenvs.ManagerSingleton(),
		cveFetcher.SingletonManager(), cloudSourcesManager.Singleton())
}

// NewPipeline returns a new instance of Pipeline.
func NewPipeline(clusters clusterDataStore.DataStore, deploymentEnvsMgr deploymentenvs.Manager,
	cveFetcher cveFetcher.OrchestratorIstioCVEManager, manager cloudSourcesManager.Manager) pipeline.Fragment {
	return &pipelineImpl{
		clusters:            clusters,
		deploymentEnvsMgr:   deploymentEnvsMgr,
		cveFetcher:          cveFetcher,
		cloudSourcesManager: manager,
	}
}

type pipelineImpl struct {
	clusters            clusterDataStore.DataStore
	deploymentEnvsMgr   deploymentenvs.Manager
	cveFetcher          cveFetcher.OrchestratorIstioCVEManager
	cloudSourcesManager cloudSourcesManager.Manager
}

func (s *pipelineImpl) Capabilities() []centralsensor.CentralCapability {
	return nil
}

func (s *pipelineImpl) Reconcile(_ context.Context, _ string, _ *reconciliation.StoreMap) error {
	// Nothing to reconcile
	return nil
}

func (s *pipelineImpl) Match(msg *central.MsgFromSensor) bool {
	return msg.GetClusterStatusUpdate() != nil
}

// Run runs the pipeline template on the input and returns the output.
func (s *pipelineImpl) Run(ctx context.Context, clusterID string, msg *central.MsgFromSensor, _ common.MessageInjector) error {
	switch m := msg.GetClusterStatusUpdate().Msg.(type) {
	case *central.ClusterStatusUpdate_DeploymentEnvUpdate:
		s.deploymentEnvsMgr.UpdateDeploymentEnvironments(clusterID, m.DeploymentEnvUpdate.Environments)
		return nil
	case *central.ClusterStatusUpdate_Status:
		if err := s.clusters.UpdateClusterStatus(ctx, clusterID, m.Status); err != nil {
			return err
		}
		go s.cveFetcher.HandleClusterConnection()
		if features.CloudSources.Enabled() {
			s.cloudSourcesManager.MarkClusterSecured(clusterID)
		}
		return nil
	default:
		return errors.Errorf("unknown cluster status update message type %T", m)
	}
}

func (s *pipelineImpl) OnFinish(clusterID string) {
	s.deploymentEnvsMgr.MarkClusterInactive(clusterID)
	if features.CloudSources.Enabled() {
		s.cloudSourcesManager.MarkClusterUnsecured(clusterID)
	}
}
