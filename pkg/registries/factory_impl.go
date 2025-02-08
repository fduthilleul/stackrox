package registries

import (
	"fmt"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/registries/types"
)

var disableRepoListForAll = env.DisableRegistryRepoList.BooleanSetting()

type factoryImpl struct {
	creators                map[string]types.Creator
	creatorsWithoutRepoList map[string]types.Creator
}

type registryWithDataSource struct {
	types.Registry
	datasource *storage.DataSource
	source     *storage.ImageIntegration
}

func (r *registryWithDataSource) DataSource() *storage.DataSource {
	return r.datasource
}

func (r *registryWithDataSource) Source() *storage.ImageIntegration {
	return r.source
}

func (e *factoryImpl) CreateRegistry(source *storage.ImageIntegration, options ...types.CreatorOption) (types.ImageRegistry, error) {
	var creator types.Creator
	var exists bool

	// The source != nil condition here ensures that sourced integrations (ie: with namespace and secret name)
	// do not use repo lists for matching, these integrations will be matched elsewhere based on source.
	if !source.Autogenerated || disableRepoListForAll || source.GetSource() != nil {
		creator, exists = e.creatorsWithoutRepoList[source.GetType()]
	}

	// If a creator was previously not found nor applicable, attempt to
	// find one among the default creators for the factory.
	if !exists {
		creator, exists = e.creators[source.GetType()]
	}

	if !exists {
		return nil, fmt.Errorf("registry with type '%s' does not exist", source.GetType())
	}
	integration, err := creator(source, options...)
	if err != nil {
		return nil, err
	}

	return &registryWithDataSource{
		Registry: integration,
		datasource: &storage.DataSource{
			Id:   source.GetId(),
			Name: source.GetName(),
		},
		source: source,
	}, nil
}
