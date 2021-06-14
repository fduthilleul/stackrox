package resources

import (
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/process/filter"
	"github.com/stackrox/rox/sensor/common/clusterentities"
	"github.com/stackrox/rox/sensor/common/config"
	"github.com/stackrox/rox/sensor/common/detector"
	complianceOperatorDispatchers "github.com/stackrox/rox/sensor/kubernetes/listener/resources/complianceoperator/dispatchers"
	"github.com/stackrox/rox/sensor/kubernetes/orchestratornamespaces"
	v1Listers "k8s.io/client-go/listers/core/v1"
)

// Dispatcher is responsible for processing resource events, and returning the sensor events that should be emitted
// in response.
//go:generate mockgen-wrapper Dispatcher
type Dispatcher interface {
	ProcessEvent(obj, oldObj interface{}, action central.ResourceAction) []*central.SensorEvent
}

// DispatcherRegistry provides dispatchers to use.
type DispatcherRegistry interface {
	ForDeployments(deploymentType string) Dispatcher
	ForJobs() Dispatcher

	ForNamespaces() Dispatcher
	ForNetworkPolicies() Dispatcher
	ForNodes() Dispatcher
	ForSecrets() Dispatcher
	ForServices() Dispatcher
	ForServiceAccounts() Dispatcher
	ForRBAC() Dispatcher
	ForClusterOperators() Dispatcher

	ForComplianceOperatorResults() Dispatcher
	ForComplianceOperatorProfiles() Dispatcher
	ForComplianceOperatorRules() Dispatcher
	ForComplianceOperatorScanSettingBindings() Dispatcher
}

// NewDispatcherRegistry creates and returns a new DispatcherRegistry.
func NewDispatcherRegistry(syncedRBAC *concurrency.Flag, clusterID string, podLister v1Listers.PodLister,
	entityStore *clusterentities.Store, processFilter filter.Filter,
	configHandler config.Handler, detector detector.Detector, namespaces *orchestratornamespaces.OrchestratorNamespaces) DispatcherRegistry {
	serviceStore := newServiceStore()
	deploymentStore := DeploymentStoreSingleton()
	podStore := PodStoreSingleton()
	nodeStore := newNodeStore()
	nsStore := newNamespaceStore()
	endpointManager := newEndpointManager(serviceStore, deploymentStore, podStore, nodeStore, entityStore)
	rbacUpdater := newRBACUpdater(syncedRBAC)

	return &registryImpl{
		deploymentHandler: newDeploymentHandler(clusterID, serviceStore, deploymentStore, podStore, endpointManager, nsStore,
			rbacUpdater, podLister, processFilter, configHandler, detector, namespaces),

		rbacDispatcher:            newRBACDispatcher(rbacUpdater),
		namespaceDispatcher:       newNamespaceDispatcher(nsStore, serviceStore, deploymentStore, podStore),
		serviceDispatcher:         newServiceDispatcher(serviceStore, deploymentStore, endpointManager),
		secretDispatcher:          newSecretDispatcher(),
		networkPolicyDispatcher:   newNetworkPolicyDispatcher(),
		nodeDispatcher:            newNodeDispatcher(serviceStore, deploymentStore, nodeStore, endpointManager),
		serviceAccountDispatcher:  newServiceAccountDispatcher(),
		clusterOperatorDispatcher: newClusterOperatorDispatcher(namespaces),

		complianceOperatorResultDispatcher:              complianceOperatorDispatchers.NewResultDispatcher(),
		complianceOperatorRulesDispatcher:               complianceOperatorDispatchers.NewRulesDispatcher(),
		complianceOperatorProfileDispatcher:             complianceOperatorDispatchers.NewProfileDispatcher(),
		complianceOperatorScanSettingBindingsDispatcher: complianceOperatorDispatchers.NewScanSettingBindingsDispatcher(),
	}
}

type registryImpl struct {
	deploymentHandler *deploymentHandler

	rbacDispatcher            *rbacDispatcher
	namespaceDispatcher       *namespaceDispatcher
	serviceDispatcher         *serviceDispatcher
	secretDispatcher          *secretDispatcher
	networkPolicyDispatcher   *networkPolicyDispatcher
	nodeDispatcher            *nodeDispatcher
	serviceAccountDispatcher  *serviceAccountDispatcher
	clusterOperatorDispatcher *clusterOperatorDispatcher

	complianceOperatorResultDispatcher              *complianceOperatorDispatchers.ResultDispatcher
	complianceOperatorProfileDispatcher             *complianceOperatorDispatchers.ProfileDispatcher
	complianceOperatorScanSettingBindingsDispatcher *complianceOperatorDispatchers.ScanSettingBindings
	complianceOperatorRulesDispatcher               *complianceOperatorDispatchers.RulesDispatcher
}

func (d *registryImpl) ForDeployments(deploymentType string) Dispatcher {
	return newDeploymentDispatcher(deploymentType, d.deploymentHandler)
}

func (d *registryImpl) ForJobs() Dispatcher {
	return newJobDispatcherImpl(d.deploymentHandler)
}

func (d *registryImpl) ForNamespaces() Dispatcher {
	return d.namespaceDispatcher
}

func (d *registryImpl) ForNetworkPolicies() Dispatcher {
	return d.networkPolicyDispatcher
}

func (d *registryImpl) ForNodes() Dispatcher {
	return d.nodeDispatcher
}

func (d *registryImpl) ForSecrets() Dispatcher {
	return d.secretDispatcher
}

func (d *registryImpl) ForServices() Dispatcher {
	return d.serviceDispatcher
}

func (d *registryImpl) ForServiceAccounts() Dispatcher {
	return d.serviceAccountDispatcher
}

func (d *registryImpl) ForRBAC() Dispatcher {
	return d.rbacDispatcher
}

func (d *registryImpl) ForClusterOperators() Dispatcher {
	return d.clusterOperatorDispatcher
}

func (d *registryImpl) ForComplianceOperatorResults() Dispatcher {
	return d.complianceOperatorResultDispatcher
}

func (d *registryImpl) ForComplianceOperatorProfiles() Dispatcher {
	return d.complianceOperatorProfileDispatcher
}

func (d *registryImpl) ForComplianceOperatorRules() Dispatcher {
	return d.complianceOperatorRulesDispatcher
}

func (d *registryImpl) ForComplianceOperatorScanSettingBindings() Dispatcher {
	return d.complianceOperatorScanSettingBindingsDispatcher
}
