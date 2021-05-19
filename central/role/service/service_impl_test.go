package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	clusterMocks "github.com/stackrox/rox/central/cluster/datastore/mocks"
	namespaceMocks "github.com/stackrox/rox/central/namespace/datastore/mocks"
	roleMocks "github.com/stackrox/rox/central/role/datastore/mocks"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/labels"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

////////////////////////////////////////////////////////////////////////////////
// Cluster and namespace configuration                                        //
//                                                                            //
// Queen       { genre: rock }                                                //
//   Queen        { released: 1973 }                                          //
//   Jazz         { released: 1978 }                                          //
//   Innuendo     { released: 1991 }                                          //
//                                                                            //
// Pink Floyd  { genre: psychedelic_rock }                                    //
//   The Wall     { released: 1979 }                                          //
//                                                                            //
// Deep Purple { genre: hard_rock }                                           //
//   Machine Head { released: 1972 }                                          //
//                                                                            //

var clusters = []*storage.Cluster{
	{
		Id:   "band.queen",
		Name: "Queen",
		Labels: map[string]string{
			"genre": "rock",
		},
	},
	{
		Id:   "band.pinkfloyd",
		Name: "Pink Floyd",
		Labels: map[string]string{
			"genre": "psychedelic_rock",
		},
	},
	{
		Id:   "band.deeppurple",
		Name: "Deep Purple",
		Labels: map[string]string{
			"genre": "hard_rock",
		},
	},
}

var namespaces = []*storage.NamespaceMetadata{
	// Queen
	{
		Id:          "album.queen",
		Name:        "Queen",
		ClusterId:   "band.queen",
		ClusterName: "Queen",
		Labels: map[string]string{
			"released": "1973",
		},
	},
	{
		Id:          "album.jazz",
		Name:        "Jazz",
		ClusterId:   "band.queen",
		ClusterName: "Queen",
		Labels: map[string]string{
			"released": "1978",
		},
	},
	{
		Id:          "album.innuendo",
		Name:        "Innuendo",
		ClusterId:   "band.queen",
		ClusterName: "Queen",
		Labels: map[string]string{
			"released": "1991",
		},
	},
	// Pink Floyd
	{
		Id:          "album.thewall",
		Name:        "The Wall",
		ClusterId:   "band.pinkfloyd",
		ClusterName: "Pink Floyd",
		Labels: map[string]string{
			"released": "1979",
		},
	},
	// Deep Purple
	{
		Id:          "album.machinehead",
		Name:        "Machine Head",
		ClusterId:   "band.deeppurple",
		ClusterName: "Deep Purple",
		Labels: map[string]string{
			"released": "1972",
		},
	},
}

////////////////////////////////////////////////////////////////////////////////
// Access scope rules and expected effective access scopes                    //
//                                                                            //
// Valid rules:                                                               //
//   `namespace: "Queen::Jazz" OR cluster.labels: genre in (psychedelic_rock)`//
//     => { "Queen::Jazz", "Pink Floyd::*" }                                  //
//                                                                            //
// Invalid rules:                                                             //
//   `namespace: "::Jazz"` => { }                                             //
//                                                                            //

var validRules = &storage.SimpleAccessScope_Rules{
	IncludedNamespaces: []*storage.SimpleAccessScope_Rules_Namespace{
		{
			ClusterName:   "Queen",
			NamespaceName: "Jazz",
		},
	},
	ClusterLabelSelectors: labels.LabelSelectors("genre", storage.SetBasedLabelSelector_IN, []string{"psychedelic_rock"}),
}

var validExpectedHigh = &storage.EffectiveAccessScope{
	Clusters: []*storage.EffectiveAccessScope_Cluster{
		{
			Id:    "band.deeppurple",
			Name:  "Deep Purple",
			State: storage.EffectiveAccessScope_EXCLUDED,
			Labels: map[string]string{
				"genre": "hard_rock",
			},
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.machinehead",
					Name:  "Machine Head",
					State: storage.EffectiveAccessScope_EXCLUDED,
					Labels: map[string]string{
						"released": "1972",
					},
				},
			},
		},
		{
			Id:    "band.pinkfloyd",
			Name:  "Pink Floyd",
			State: storage.EffectiveAccessScope_INCLUDED,
			Labels: map[string]string{
				"genre": "psychedelic_rock",
			},
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.thewall",
					Name:  "The Wall",
					State: storage.EffectiveAccessScope_INCLUDED,
					Labels: map[string]string{
						"released": "1979",
					},
				},
			},
		},
		{
			Id:    "band.queen",
			Name:  "Queen",
			State: storage.EffectiveAccessScope_PARTIAL,
			Labels: map[string]string{
				"genre": "rock",
			},
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.innuendo",
					Name:  "Innuendo",
					State: storage.EffectiveAccessScope_EXCLUDED,
					Labels: map[string]string{
						"released": "1991",
					},
				},
				{
					Id:    "album.jazz",
					Name:  "Jazz",
					State: storage.EffectiveAccessScope_INCLUDED,
					Labels: map[string]string{
						"released": "1978",
					},
				},
				{
					Id:    "album.queen",
					Name:  "Queen",
					State: storage.EffectiveAccessScope_EXCLUDED,
					Labels: map[string]string{
						"released": "1973",
					},
				},
			},
		},
	},
}

var validExpectedStandard = &storage.EffectiveAccessScope{
	Clusters: []*storage.EffectiveAccessScope_Cluster{
		{
			Id:    "band.deeppurple",
			Name:  "Deep Purple",
			State: storage.EffectiveAccessScope_EXCLUDED,
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.machinehead",
					Name:  "Machine Head",
					State: storage.EffectiveAccessScope_EXCLUDED,
				},
			},
		},
		{
			Id:    "band.pinkfloyd",
			Name:  "Pink Floyd",
			State: storage.EffectiveAccessScope_INCLUDED,
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.thewall",
					Name:  "The Wall",
					State: storage.EffectiveAccessScope_INCLUDED,
				},
			},
		},
		{
			Id:    "band.queen",
			Name:  "Queen",
			State: storage.EffectiveAccessScope_PARTIAL,
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.innuendo",
					Name:  "Innuendo",
					State: storage.EffectiveAccessScope_EXCLUDED,
				},
				{
					Id:    "album.jazz",
					Name:  "Jazz",
					State: storage.EffectiveAccessScope_INCLUDED,
				},
				{
					Id:    "album.queen",
					Name:  "Queen",
					State: storage.EffectiveAccessScope_EXCLUDED,
				},
			},
		},
	},
}

var validExpectedMinimal = &storage.EffectiveAccessScope{
	Clusters: []*storage.EffectiveAccessScope_Cluster{
		{
			Id:    "band.pinkfloyd",
			State: storage.EffectiveAccessScope_INCLUDED,
		},
		{
			Id:    "band.queen",
			State: storage.EffectiveAccessScope_PARTIAL,
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.jazz",
					State: storage.EffectiveAccessScope_INCLUDED,
				},
			},
		},
	},
}

var invalidRules = &storage.SimpleAccessScope_Rules{
	IncludedNamespaces: []*storage.SimpleAccessScope_Rules_Namespace{
		{
			NamespaceName: "Jazz",
		},
	},
}

var invalidExpectedHigh = &storage.EffectiveAccessScope{
	Clusters: []*storage.EffectiveAccessScope_Cluster{
		{
			Id:    "band.deeppurple",
			Name:  "Deep Purple",
			State: storage.EffectiveAccessScope_EXCLUDED,
			Labels: map[string]string{
				"genre": "hard_rock",
			},
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.machinehead",
					Name:  "Machine Head",
					State: storage.EffectiveAccessScope_EXCLUDED,
					Labels: map[string]string{
						"released": "1972",
					},
				},
			},
		},
		{
			Id:    "band.pinkfloyd",
			Name:  "Pink Floyd",
			State: storage.EffectiveAccessScope_EXCLUDED,
			Labels: map[string]string{
				"genre": "psychedelic_rock",
			},
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.thewall",
					Name:  "The Wall",
					State: storage.EffectiveAccessScope_EXCLUDED,
					Labels: map[string]string{
						"released": "1979",
					},
				},
			},
		},
		{
			Id:    "band.queen",
			Name:  "Queen",
			State: storage.EffectiveAccessScope_EXCLUDED,
			Labels: map[string]string{
				"genre": "rock",
			},
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.innuendo",
					Name:  "Innuendo",
					State: storage.EffectiveAccessScope_EXCLUDED,
					Labels: map[string]string{
						"released": "1991",
					},
				},
				{
					Id:    "album.jazz",
					Name:  "Jazz",
					State: storage.EffectiveAccessScope_EXCLUDED,
					Labels: map[string]string{
						"released": "1978",
					},
				},
				{
					Id:    "album.queen",
					Name:  "Queen",
					State: storage.EffectiveAccessScope_EXCLUDED,
					Labels: map[string]string{
						"released": "1973",
					},
				},
			},
		},
	},
}

var invalidExpectedStandard = &storage.EffectiveAccessScope{
	Clusters: []*storage.EffectiveAccessScope_Cluster{
		{
			Id:    "band.deeppurple",
			Name:  "Deep Purple",
			State: storage.EffectiveAccessScope_EXCLUDED,
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.machinehead",
					Name:  "Machine Head",
					State: storage.EffectiveAccessScope_EXCLUDED,
				},
			},
		},
		{
			Id:    "band.pinkfloyd",
			Name:  "Pink Floyd",
			State: storage.EffectiveAccessScope_EXCLUDED,
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.thewall",
					Name:  "The Wall",
					State: storage.EffectiveAccessScope_EXCLUDED,
				},
			},
		},
		{
			Id:    "band.queen",
			Name:  "Queen",
			State: storage.EffectiveAccessScope_EXCLUDED,
			Namespaces: []*storage.EffectiveAccessScope_Namespace{
				{
					Id:    "album.innuendo",
					Name:  "Innuendo",
					State: storage.EffectiveAccessScope_EXCLUDED,
				},
				{
					Id:    "album.jazz",
					Name:  "Jazz",
					State: storage.EffectiveAccessScope_EXCLUDED,
				},
				{
					Id:    "album.queen",
					Name:  "Queen",
					State: storage.EffectiveAccessScope_EXCLUDED,
				},
			},
		},
	},
}

var invalidExpectedMinimal = &storage.EffectiveAccessScope{}

////////////////////////////////////////////////////////////////////////////////
// Tests                                                                      //
//                                                                            //

func TestEffectiveAccessScopeForSimpleAccessScope(t *testing.T) {
	type testCase struct {
		desc             string
		rules            *storage.SimpleAccessScope_Rules
		expectedHigh     *storage.EffectiveAccessScope
		expectedStandard *storage.EffectiveAccessScope
		expectedMinimal  *storage.EffectiveAccessScope
	}

	testCases := []testCase{
		{
			"valid access scope rules",
			validRules,
			validExpectedHigh,
			validExpectedStandard,
			validExpectedMinimal,
		},
		{
			"invalid access scope rules",
			invalidRules,
			invalidExpectedHigh,
			invalidExpectedStandard,
			invalidExpectedMinimal,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			resHigh, err := effectiveAccessScopeForSimpleAccessScope(tc.rules, clusters, namespaces, v1.ComputeEffectiveAccessScopeRequest_HIGH)
			assert.NoError(t, err)
			assert.Exactly(t, tc.expectedHigh, resHigh, "detail: HIGH")

			resStandard, err := effectiveAccessScopeForSimpleAccessScope(tc.rules, clusters, namespaces, v1.ComputeEffectiveAccessScopeRequest_STANDARD)
			assert.NoError(t, err)
			assert.Exactly(t, tc.expectedStandard, resStandard, "detail: STANDARD")

			resMinimal, err := effectiveAccessScopeForSimpleAccessScope(tc.rules, clusters, namespaces, v1.ComputeEffectiveAccessScopeRequest_MINIMAL)
			assert.NoError(t, err)
			assert.Exactly(t, tc.expectedMinimal, resMinimal, "detail: MINIMAL")

			resUnknown, err := effectiveAccessScopeForSimpleAccessScope(tc.rules, clusters, namespaces, 42)
			assert.NoError(t, err)
			assert.Exactly(t, tc.expectedStandard, resUnknown, "unknown detail maps to STANDARD")
		})
	}
}

func TestServiceImpl(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ServiceTestSuite))
}

type ServiceTestSuite struct {
	suite.Suite

	envIsolator    *envisolator.EnvIsolator
	requestContext context.Context

	mockCtrl *gomock.Controller

	mockRoles      *roleMocks.MockDataStore
	mockClusters   *clusterMocks.MockDataStore
	mockNamespaces *namespaceMocks.MockDataStore

	svc Service
}

func (s *ServiceTestSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.requestContext = context.Background()

	s.mockCtrl = gomock.NewController(s.T())

	s.mockRoles = roleMocks.NewMockDataStore(s.mockCtrl)
	s.mockClusters = clusterMocks.NewMockDataStore(s.mockCtrl)
	s.mockNamespaces = namespaceMocks.NewMockDataStore(s.mockCtrl)

	s.svc = New(s.mockRoles, s.mockClusters, s.mockNamespaces)
}

func (s *ServiceTestSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
	s.mockCtrl.Finish()
}

func (s *ServiceTestSuite) TestAccessScopeAPIDisabledByDefault() {
	s.envIsolator.Setenv(features.ScopedAccessControl.EnvVar(), "false")
	s.False(features.ScopedAccessControl.Enabled())

	ctx := sac.WithGlobalAccessScopeChecker(context.Background(), sac.AllowAllAccessScopeChecker())

	_, err := s.svc.GetSimpleAccessScope(ctx, &v1.ResourceByID{Id: "someid"})
	s.ErrorIs(err, status.Error(codes.Unimplemented, "feature not enabled"))

	_, err = s.svc.ListSimpleAccessScopes(ctx, &v1.Empty{})
	s.ErrorIs(err, status.Error(codes.Unimplemented, "feature not enabled"))

	_, err = s.svc.PostSimpleAccessScope(ctx, &storage.SimpleAccessScope{})
	s.ErrorIs(err, status.Error(codes.Unimplemented, "feature not enabled"))

	_, err = s.svc.PutSimpleAccessScope(ctx, &storage.SimpleAccessScope{})
	s.ErrorIs(err, status.Error(codes.Unimplemented, "feature not enabled"))

	_, err = s.svc.DeleteSimpleAccessScope(ctx, &v1.ResourceByID{Id: "someid"})
	s.ErrorIs(err, status.Error(codes.Unimplemented, "feature not enabled"))

	_, err = s.svc.ComputeEffectiveAccessScope(ctx, &v1.ComputeEffectiveAccessScopeRequest{})
	s.ErrorIs(err, status.Error(codes.Unimplemented, "feature not enabled"))
}
