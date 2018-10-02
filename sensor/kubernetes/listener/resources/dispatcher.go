package resources

import (
	pkgV1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/listeners"
	"k8s.io/api/core/v1"
	networkingV1 "k8s.io/api/networking/v1"
	v1Listers "k8s.io/client-go/listers/core/v1"
)

// Dispatcher is responsible for processing resource events, and returning the sensor events that should be emitted
// in response.
type Dispatcher interface {
	ProcessEvent(obj interface{}, action pkgV1.ResourceAction, deploymentType string) []*listeners.EventWrap
}

// NewDispatcher creates and returns a new dispatcher.
func NewDispatcher(podLister v1Listers.PodLister) Dispatcher {
	serviceStore := newServiceStore()
	deploymentStore := newDeploymentStore()

	return &dispatcher{
		namespaceHandler:     newNamespaceHandler(serviceStore, deploymentStore),
		deploymentHandler:    newDeploymentHandler(serviceStore, deploymentStore, podLister),
		serviceHandler:       newServiceHandler(serviceStore, deploymentStore),
		secretHandler:        newSecretHandler(),
		networkPolicyHandler: newNetworkPolicyHandler(),
	}
}

type dispatcher struct {
	namespaceHandler     *namespaceHandler
	deploymentHandler    *deploymentHandler
	serviceHandler       *serviceHandler
	secretHandler        *secretHandler
	networkPolicyHandler *networkPolicyHandler
}

func (d *dispatcher) ProcessEvent(obj interface{}, action pkgV1.ResourceAction, deploymentType string) []*listeners.EventWrap {
	if deploymentType != "" {
		return d.deploymentHandler.Process(obj, action, deploymentType)
	}

	switch o := obj.(type) {
	case *v1.Service:
		return d.serviceHandler.Process(o, action)
	case *v1.Secret:
		return d.secretHandler.Process(o, action)
	case *networkingV1.NetworkPolicy:
		return d.networkPolicyHandler.Process(o, action)
	case *v1.Namespace:
		return d.namespaceHandler.Process(o, action)
	default:
		return nil
	}
}
