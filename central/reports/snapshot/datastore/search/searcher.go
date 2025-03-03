package search

import (
	"context"

	pgStore "github.com/stackrox/rox/central/reports/snapshot/datastore/store/postgres"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/paginated"
	pkgPostgres "github.com/stackrox/rox/pkg/search/scoped/postgres"
)

var (
	defaultSortOption = &v1.QuerySortOption{
		Field:    search.ReportCompletionTime.String(),
		Reversed: true,
	}
)

// Searcher provides search functionality on existing ReportSnapshots.
//
//go:generate mockgen-wrapper
type Searcher interface {
	Search(ctx context.Context, query *v1.Query) ([]search.Result, error)
	Count(ctx context.Context, query *v1.Query) (int, error)
	SearchResults(ctx context.Context, query *v1.Query) ([]*v1.SearchResult, error)
	SearchReportSnapshots(context.Context, *v1.Query) ([]*storage.ReportSnapshot, error)
}

// New returns a new instance of Searcher for the given storage.
func New(storage pgStore.Store) Searcher {
	return &searcherImpl{
		storage:  storage,
		searcher: formatSearcher(storage),
	}
}

func formatSearcher(searcher search.Searcher) search.Searcher {
	scopedSafeSearcher := pkgPostgres.WithScoping(searcher)
	return paginated.WithDefaultSortOption(scopedSafeSearcher, defaultSortOption)
}
