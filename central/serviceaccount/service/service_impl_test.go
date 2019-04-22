package service

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	deploymentMocks "github.com/stackrox/rox/central/deployment/datastore/mocks"
	namespaceMocks "github.com/stackrox/rox/central/namespace/datastore/mocks"
	roleMocks "github.com/stackrox/rox/central/rbac/k8srole/datastore/mocks"
	bindingMocks "github.com/stackrox/rox/central/rbac/k8srolebinding/datastore/mocks"
	saMocks "github.com/stackrox/rox/central/serviceaccount/datastore/mocks"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stretchr/testify/suite"
)

var (
	saID = "id1"

	expectedSA = &storage.ServiceAccount{
		Id:        saID,
		Name:      "serviceaccountname",
		ClusterId: "cluster",
		Namespace: "namespace",
	}

	listDeployment = &storage.ListDeployment{
		Id:        "deploymentId",
		Name:      "deployment",
		ClusterId: "cluster",
		Namespace: "namespace",
	}

	role = &storage.K8SRole{
		Id:           "role1",
		Name:         "role1",
		ClusterId:    "cluster",
		Namespace:    "namespace",
		ClusterScope: false,
	}
	clusterRole = &storage.K8SRole{
		Id:           "role2",
		Name:         "role2",
		ClusterId:    "cluster",
		ClusterScope: true,
	}

	rolebinding = &storage.K8SRoleBinding{
		RoleId: "role1",
		Subjects: []*storage.Subject{
			{
				Name:      "serviceaccountname",
				Namespace: "namespace",
				Kind:      storage.SubjectKind_SERVICE_ACCOUNT,
			},
		},
		ClusterScope: false,
		Namespace:    "namespace",
		Id:           "binding1",
	}

	clusterRoleBinding = &storage.K8SRoleBinding{
		RoleId: "role2",
		Subjects: []*storage.Subject{
			{
				Name:      "serviceaccountname",
				Namespace: "namespace",
				Kind:      storage.SubjectKind_SERVICE_ACCOUNT,
			},
		},
		ClusterScope: true,
		Id:           "binding2",
	}

	namespaceMetadata = &storage.NamespaceMetadata{
		Name: "namespace",
	}
)

func TestServiceAccountService(t *testing.T) {
	suite.Run(t, new(ServiceAccountServiceTestSuite))
}

type ServiceAccountServiceTestSuite struct {
	suite.Suite

	mockServiceAccountStore *saMocks.MockDataStore
	mockDeploymentStore     *deploymentMocks.MockDataStore
	mockRoleStore           *roleMocks.MockDataStore
	mockBindingStore        *bindingMocks.MockDataStore
	mockNamespaceStore      *namespaceMocks.MockDataStore
	service                 Service

	mockCtrl *gomock.Controller
}

func (suite *ServiceAccountServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockServiceAccountStore = saMocks.NewMockDataStore(suite.mockCtrl)
	suite.mockDeploymentStore = deploymentMocks.NewMockDataStore(suite.mockCtrl)
	suite.mockRoleStore = roleMocks.NewMockDataStore(suite.mockCtrl)
	suite.mockBindingStore = bindingMocks.NewMockDataStore(suite.mockCtrl)
	suite.mockNamespaceStore = namespaceMocks.NewMockDataStore(suite.mockCtrl)

	suite.service = New(suite.mockServiceAccountStore, suite.mockBindingStore, suite.mockRoleStore,
		suite.mockDeploymentStore, suite.mockNamespaceStore)
}

// Test happy path for getting service accounts
func (suite *ServiceAccountServiceTestSuite) TestGetServiceAccount() {

	suite.setupMocks()

	suite.mockServiceAccountStore.EXPECT().GetServiceAccount(saID).Return(expectedSA, true, nil)

	sa, err := suite.service.GetServiceAccount((context.Context)(nil), &v1.ResourceByID{Id: saID})
	suite.NoError(err)
	suite.Equal(expectedSA, sa.SaAndRole.ServiceAccount)
	suite.Equal(1, len(sa.SaAndRole.DeploymentRelationships))
	suite.Equal(listDeployment.GetName(), sa.SaAndRole.DeploymentRelationships[0].GetName())
	suite.Equal(1, len(sa.SaAndRole.ScopedRoles))
	suite.Equal(1, len(sa.SaAndRole.ClusterRoles))
	suite.Equal("namespace", sa.SaAndRole.ScopedRoles[0].Namespace)

}

// Test that when we fail to find a service account, an error is returned.
func (suite *ServiceAccountServiceTestSuite) TestGetSAWithStoreSANotExists() {
	saID := "id1"

	suite.mockServiceAccountStore.EXPECT().GetServiceAccount(saID).Return((*storage.ServiceAccount)(nil), false, nil)

	_, err := suite.service.GetServiceAccount((context.Context)(nil), &v1.ResourceByID{Id: saID})
	suite.Error(err)
}

// Test that when we fail to read the db for a secret, an error is returned.
func (suite *ServiceAccountServiceTestSuite) TestGetSAWithStoreSAFailure() {
	saID := "id1"

	expectedErr := fmt.Errorf("failure")
	suite.mockServiceAccountStore.EXPECT().GetServiceAccount(saID).Return((*storage.ServiceAccount)(nil), true, expectedErr)

	_, actualErr := suite.service.GetServiceAccount((context.Context)(nil), &v1.ResourceByID{Id: saID})
	suite.Error(actualErr)
}

// Test happy path for searching secrets and relationships
func (suite *ServiceAccountServiceTestSuite) TestSearchServiceAccount() {
	suite.setupMocks()

	suite.mockServiceAccountStore.EXPECT().ListServiceAccounts().Return([]*storage.ServiceAccount{expectedSA}, nil)

	q := search.NewQueryBuilder().AddExactMatches(search.ClusterID, expectedSA.ClusterId).
		AddExactMatches(search.Namespace, expectedSA.GetNamespace()).
		AddExactMatches(search.ServiceAccountName, expectedSA.GetName()).ProtoQuery()

	suite.mockDeploymentStore.EXPECT().SearchListDeployments(q).AnyTimes().Return([]*storage.ListDeployment{listDeployment}, nil)

	_, err := suite.service.ListServiceAccounts((context.Context)(nil), &v1.RawQuery{})
	suite.NoError(err)
}

// Test that when searching fails, that error is returned.
func (suite *ServiceAccountServiceTestSuite) TestSearchServiceAccountFailure() {
	expectedError := fmt.Errorf("failure")

	suite.mockServiceAccountStore.EXPECT().ListServiceAccounts().Return(nil, expectedError)

	_, actualErr := suite.service.ListServiceAccounts((context.Context)(nil), &v1.RawQuery{})
	suite.True(strings.Contains(actualErr.Error(), expectedError.Error()))
}

func (suite *ServiceAccountServiceTestSuite) setupMocks() {

	q := search.NewQueryBuilder().AddExactMatches(search.ClusterID, expectedSA.ClusterId).
		AddExactMatches(search.Namespace, expectedSA.GetNamespace()).
		AddExactMatches(search.ServiceAccountName, expectedSA.GetName()).ProtoQuery()

	suite.mockDeploymentStore.EXPECT().SearchListDeployments(q).Return([]*storage.ListDeployment{listDeployment}, nil)

	suite.mockRoleStore.EXPECT().GetRole("role1").AnyTimes().Return(role, true, nil)
	suite.mockRoleStore.EXPECT().GetRole("role2").AnyTimes().Return(clusterRole, true, nil)

	namespaceQ := search.NewQueryBuilder().AddExactMatches(search.ClusterID, "cluster").ProtoQuery()
	suite.mockNamespaceStore.EXPECT().SearchNamespaces(namespaceQ).AnyTimes().
		Return([]*storage.NamespaceMetadata{namespaceMetadata}, nil)

	clusterScopeQuery := search.NewQueryBuilder().
		AddExactMatches(search.ClusterID, "cluster").
		AddBools(search.ClusterScope, true).ProtoQuery()
	suite.mockBindingStore.EXPECT().SearchRawRoleBindings(clusterScopeQuery).AnyTimes().
		Return([]*storage.K8SRoleBinding{clusterRoleBinding}, nil)

	namespaceScopeQuery := search.NewQueryBuilder().
		AddExactMatches(search.ClusterID, "cluster").
		AddExactMatches(search.Namespace, "namespace").
		AddBools(search.ClusterScope, false).ProtoQuery()
	suite.mockBindingStore.EXPECT().SearchRawRoleBindings(namespaceScopeQuery).AnyTimes().
		Return([]*storage.K8SRoleBinding{rolebinding}, nil)

}
