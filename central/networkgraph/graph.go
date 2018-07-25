package networkgraph

import (
	"sort"

	clusterDatastore "bitbucket.org/stack-rox/apollo/central/cluster/datastore"
	deploymentsDatastore "bitbucket.org/stack-rox/apollo/central/deployment/datastore"
	networkPolicyStore "bitbucket.org/stack-rox/apollo/central/networkpolicies/store"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
	"bitbucket.org/stack-rox/apollo/pkg/set"
	"github.com/deckarep/golang-set"
)

var logger = logging.LoggerForModule()

// GraphEvaluator implements the interface for the network graph generator
type GraphEvaluator interface {
	GetGraph() (*v1.GetNetworkGraphResponse, error)
}

// graphEvaluatorImpl handles all of the graph calculations
type graphEvaluatorImpl struct {
	clustersStore      clusterDatastore.DataStore
	deploymentsStore   deploymentsDatastore.DataStore
	namespaceStore     ns
	networkPolicyStore networkPolicyStore.Store
}

type ns interface {
	GetNamespaces() ([]*v1.Namespace, error)
}

// newGraphEvaluator takes in namespaces
func newGraphEvaluator(clustersStore clusterDatastore.DataStore, deploymentsStore deploymentsDatastore.DataStore,
	namespaceStore ns, networkPolicyStore networkPolicyStore.Store) *graphEvaluatorImpl {
	return &graphEvaluatorImpl{
		clustersStore:      clustersStore,
		deploymentsStore:   deploymentsStore,
		namespaceStore:     namespaceStore,
		networkPolicyStore: networkPolicyStore,
	}
}

func (g *graphEvaluatorImpl) GetGraph() (*v1.GetNetworkGraphResponse, error) {
	nodes, edges, err := g.evaluate()
	if err != nil {
		return nil, err
	}
	return &v1.GetNetworkGraphResponse{
		Nodes: nodes,
		Edges: edges,
	}, nil
}

func (g *graphEvaluatorImpl) evaluate() (nodes []*v1.NetworkNode, edges []*v1.NetworkEdge, err error) {
	clusters, err := g.clustersStore.GetClusters()
	if err != nil {
		return
	}
	networkPolicies, err := g.networkPolicyStore.GetNetworkPolicies()
	if err != nil {
		return
	}
	for _, c := range clusters {
		var deployments []*v1.Deployment
		deployments, err = g.deploymentsStore.SearchRawDeployments(&v1.ParsedSearchRequest{
			Scopes: []*v1.Scope{
				{
					Cluster: c.GetName(),
				},
			},
		})
		if err != nil {
			return
		}

		clusterNodes, clusterEdges := g.evaluateCluster(deployments, networkPolicies)
		nodes = append(nodes, clusterNodes...)
		edges = append(edges, clusterEdges...)
	}
	return
}

func networkPolicySelectorAppliesToDeployment(d *v1.Deployment, np *v1.NetworkPolicy, hasPolicy func([]v1.NetworkPolicyType) bool) bool {
	spec := np.GetSpec()
	// Check if the src matches the pod selector and deployment then the egress rules actually apply to that deployment
	if !doesPodLabelsMatchLabel(d, spec.GetPodSelector()) || d.GetNamespace() != np.GetNamespace() {
		return false
	}
	// If no egress rules are defined, then it doesn't apply
	return hasPolicy(spec.GetPolicyTypes())
}

func (g *graphEvaluatorImpl) matchPolicyPeers(d *v1.Deployment, namespace string, peers []*v1.NetworkPolicyPeer) bool {
	if len(peers) == 0 {
		return true
	}
	for _, p := range peers {
		if g.matchPolicyPeer(d, namespace, p) {
			return true
		}
	}
	return false
}

func (g *graphEvaluatorImpl) doesEgressNetworkPolicyRuleMatchDeployment(src *v1.Deployment, np *v1.NetworkPolicy) bool {
	for _, egressRule := range np.GetSpec().GetEgress() {
		if g.matchPolicyPeers(src, np.GetNamespace(), egressRule.GetTo()) {
			return true
		}
	}
	return false
}

func (g *graphEvaluatorImpl) doesIngressNetworkPolicyRuleMatchDeployment(src *v1.Deployment, np *v1.NetworkPolicy) bool {
	for _, ingressRule := range np.GetSpec().GetIngress() {
		if g.matchPolicyPeers(src, np.GetNamespace(), ingressRule.GetFrom()) {
			return true
		}
	}
	return false
}

func (g *graphEvaluatorImpl) evaluateCluster(deployments []*v1.Deployment, networkPolicies []*v1.NetworkPolicy) ([]*v1.NetworkNode, []*v1.NetworkEdge) {
	selectedDeploymentsToIngressPolicies := make(map[string]mapset.Set)
	selectedDeploymentsToEgressPolicies := make(map[string]mapset.Set)

	matchedDeploymentsToIngressPolicies := make(map[string]mapset.Set)
	matchedDeploymentsToEgressPolicies := make(map[string]mapset.Set)

	var nodes []*v1.NetworkNode
	for _, d := range deployments {
		selectedDeploymentsToIngressPolicies[d.GetId()] = mapset.NewSet()
		selectedDeploymentsToEgressPolicies[d.GetId()] = mapset.NewSet()
		matchedDeploymentsToIngressPolicies[d.GetId()] = mapset.NewSet()
		matchedDeploymentsToEgressPolicies[d.GetId()] = mapset.NewSet()

		for _, n := range networkPolicies {
			if n.GetSpec() == nil {
				continue
			}
			if networkPolicySelectorAppliesToDeployment(d, n, hasIngress) {
				selectedDeploymentsToIngressPolicies[d.GetId()].Add(n.GetId())
			}
			if g.doesIngressNetworkPolicyRuleMatchDeployment(d, n) {
				matchedDeploymentsToIngressPolicies[d.GetId()].Add(n.GetId())
			}
			if networkPolicySelectorAppliesToDeployment(d, n, hasEgress) {
				selectedDeploymentsToEgressPolicies[d.GetId()].Add(n.GetId())
			}
			if g.doesEgressNetworkPolicyRuleMatchDeployment(d, n) {
				matchedDeploymentsToEgressPolicies[d.GetId()].Add(n.GetId())
			}
		}

		nodePoliciesSet := set.StringSliceFromSet(selectedDeploymentsToIngressPolicies[d.GetId()].Union(selectedDeploymentsToEgressPolicies[d.GetId()]))
		sort.Strings(nodePoliciesSet)

		nodes = append(nodes, &v1.NetworkNode{
			Id:        d.GetId(),
			Cluster:   d.GetClusterName(),
			Namespace: d.GetNamespace(),
			PolicyIds: nodePoliciesSet,
		})
	}

	var edges []*v1.NetworkEdge
	for _, src := range deployments {
		for _, dst := range deployments {
			if src.GetId() == dst.GetId() {
				continue
			}

			// This set is the set of Egress policies that are applicable to the src
			selectedEgressPoliciesSet := selectedDeploymentsToEgressPolicies[src.GetId()]
			// This set is the set if Egress policies that have rules that are applicable to the dst
			matchedEgressPoliciesSet := matchedDeploymentsToEgressPolicies[dst.GetId()]
			// If there are no values in the src set of egress then it has no Egress rules and can talk to everything
			// Otherwise, if it is not empty then ensure that the intersection of the policies that apply to the source and the rules that apply to the dst have at least one in common
			if selectedEgressPoliciesSet.Cardinality() != 0 && selectedEgressPoliciesSet.Intersect(matchedEgressPoliciesSet).Cardinality() == 0 {
				continue
			}

			// This set is the set of Ingress policies that are applicable to the dst
			selectedIngressPoliciesSet := selectedDeploymentsToIngressPolicies[dst.GetId()]
			// This set is the set if Ingress policies that have rules that are applicable to the src
			matchedIngressPoliciesSet := matchedDeploymentsToIngressPolicies[src.GetId()]
			// If there are no values in the src set of egress then it has no Egress rules and can talk to everything
			// Otherwise, if it is not empty then ensure that the intersection of the policies that apply to the source and the rules that apply to the dst have at least one in common
			if selectedIngressPoliciesSet.Cardinality() != 0 && selectedIngressPoliciesSet.Intersect(matchedIngressPoliciesSet).Cardinality() == 0 {
				continue
			}

			var policyIDs []string
			policyIDs = append(policyIDs, set.StringSliceFromSet(selectedIngressPoliciesSet.Intersect(matchedIngressPoliciesSet))...)
			policyIDs = append(policyIDs, set.StringSliceFromSet(selectedEgressPoliciesSet.Intersect(matchedEgressPoliciesSet))...)

			edges = append(edges, &v1.NetworkEdge{Source: src.GetId(), Target: dst.GetId()})
		}
	}
	return nodes, edges
}

func (g *graphEvaluatorImpl) getNamespace(deployment *v1.Deployment) *v1.Namespace {
	namespaces, err := g.namespaceStore.GetNamespaces()
	if err != nil {
		return &v1.Namespace{
			Name: deployment.GetNamespace(),
		}
	}
	for _, n := range namespaces {
		if n.GetName() == deployment.GetNamespace() && n.GetClusterId() == deployment.GetClusterId() {
			return n
		}
	}
	return &v1.Namespace{
		Name: deployment.GetNamespace(),
	}
}

func (g *graphEvaluatorImpl) matchPolicyPeer(deployment *v1.Deployment, policyNamespace string, peer *v1.NetworkPolicyPeer) bool {
	if peer.IpBlock != nil {
		logger.Infof("IP Block network policy is currently not handled")
		return false
	}

	// If namespace selector is specified, then make sure the namespace matches
	// Other you fall back to the fact that the deployment must be in the policy's namespace
	if peer.GetNamespaceSelector() != nil {
		namespace := g.getNamespace(deployment)
		if !doesNamespaceMatchLabel(namespace, peer.GetNamespaceSelector()) {
			return false
		}
	} else if deployment.GetNamespace() != policyNamespace {
		return false
	}

	if peer.GetPodSelector() != nil {
		return doesPodLabelsMatchLabel(deployment, peer.GetPodSelector())
	}
	return true
}

func doesNamespaceMatchLabel(namespace *v1.Namespace, selector *v1.LabelSelector) bool {
	if len(selector.MatchLabels) == 0 {
		return true
	}
	for k, v := range namespace.GetLabels() {
		if selector.MatchLabels[k] == v {
			return true
		}
	}
	return false
}

func doesPodLabelsMatchLabel(deployment *v1.Deployment, podSelector *v1.LabelSelector) bool {
	// No values equals match all
	if len(podSelector.GetMatchLabels()) == 0 {
		return true
	}
	deploymentLabelMap := make(map[string]string)
	for _, keyValue := range deployment.GetLabels() {
		deploymentLabelMap[keyValue.GetKey()] = keyValue.GetValue()
	}
	for k, v := range podSelector.GetMatchLabels() {
		if deploymentLabelMap[k] != v {
			return false
		}
	}
	return true
}

func hasPolicyType(types []v1.NetworkPolicyType, t v1.NetworkPolicyType) bool {
	for _, pType := range types {
		if pType == t {
			return true
		}
	}
	return false
}

func hasEgress(types []v1.NetworkPolicyType) bool {
	return hasPolicyType(types, v1.NetworkPolicyType_EGRESS_NETWORK_POLICY_TYPE)
}

func hasIngress(types []v1.NetworkPolicyType) bool {
	if len(types) == 0 {
		return true
	}
	return hasPolicyType(types, v1.NetworkPolicyType_INGRESS_NETWORK_POLICY_TYPE)
}
