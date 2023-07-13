// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/jackc/pgx/v4"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/central/role/resources"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
	"github.com/stackrox/rox/pkg/sync"
	"gorm.io/gorm"
)

const (
	baseTable = "test_grandparents"
	storeName = "TestGrandparent"

	batchAfter = 100

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.TestGrandparentsSchema
	targetResource = resources.Namespace
)

type storeType = storage.TestGrandparent

// Store is the interface to interact with the storage for storage.TestGrandparent
type Store interface {
	Upsert(ctx context.Context, obj *storeType) error
	UpsertMany(ctx context.Context, objs []*storeType) error
	Delete(ctx context.Context, id string) error
	DeleteByQuery(ctx context.Context, q *v1.Query) error
	DeleteMany(ctx context.Context, identifiers []string) error

	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)

	Get(ctx context.Context, id string) (*storeType, bool, error)
	GetByQuery(ctx context.Context, query *v1.Query) ([]*storeType, error)
	GetMany(ctx context.Context, identifiers []string) ([]*storeType, []int, error)
	GetIDs(ctx context.Context) ([]string, error)

	Walk(ctx context.Context, fn func(obj *storeType) error) error
}

type storeImpl struct {
	*pgSearch.GenericStore[storeType, *storeType]
	mutex sync.RWMutex
}

// New returns a new Store instance using the provided sql instance.
func New(db postgres.DB) Store {
	return &storeImpl{
		GenericStore: pgSearch.NewGenericStore[storeType, *storeType](
			db,
			schema,
			pkGetter,
			metricsSetAcquireDBConnDuration,
			metricsSetPostgresOperationDurationTime,
			targetResource,
		),
	}
}

// region Helper functions

func pkGetter(obj *storeType) string {
	return obj.GetId()
}

func metricsSetPostgresOperationDurationTime(start time.Time, op ops.Op) {
	metrics.SetPostgresOperationDurationTime(start, op, storeName)
}

func metricsSetAcquireDBConnDuration(start time.Time, op ops.Op) {
	metrics.SetAcquireDBConnDuration(start, op, storeName)
}

func insertIntoTestGrandparents(ctx context.Context, batch *pgx.Batch, obj *storage.TestGrandparent) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetId(),
		obj.GetVal(),
		obj.GetPriority(),
		obj.GetRiskScore(),
		serialized,
	}

	finalStr := "INSERT INTO test_grandparents (Id, Val, Priority, RiskScore, serialized) VALUES($1, $2, $3, $4, $5) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Val = EXCLUDED.Val, Priority = EXCLUDED.Priority, RiskScore = EXCLUDED.RiskScore, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

	var query string

	for childIndex, child := range obj.GetEmbedded() {
		if err := insertIntoTestGrandparentsEmbeddeds(ctx, batch, child, obj.GetId(), childIndex); err != nil {
			return err
		}
	}

	query = "delete from test_grandparents_embeddeds where test_grandparents_Id = $1 AND idx >= $2"
	batch.Queue(query, obj.GetId(), len(obj.GetEmbedded()))
	return nil
}

func insertIntoTestGrandparentsEmbeddeds(ctx context.Context, batch *pgx.Batch, obj *storage.TestGrandparent_Embedded, testGrandparentID string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		testGrandparentID,
		idx,
		obj.GetVal(),
	}

	finalStr := "INSERT INTO test_grandparents_embeddeds (test_grandparents_Id, idx, Val) VALUES($1, $2, $3) ON CONFLICT(test_grandparents_Id, idx) DO UPDATE SET test_grandparents_Id = EXCLUDED.test_grandparents_Id, idx = EXCLUDED.idx, Val = EXCLUDED.Val"
	batch.Queue(finalStr, values...)

	var query string

	for childIndex, child := range obj.GetEmbedded2() {
		if err := insertIntoTestGrandparentsEmbeddedsEmbedded2(ctx, batch, child, testGrandparentID, idx, childIndex); err != nil {
			return err
		}
	}

	query = "delete from test_grandparents_embeddeds_embedded2 where test_grandparents_Id = $1 AND test_grandparents_embeddeds_idx = $2 AND idx >= $3"
	batch.Queue(query, testGrandparentID, idx, len(obj.GetEmbedded2()))
	return nil
}

func insertIntoTestGrandparentsEmbeddedsEmbedded2(_ context.Context, batch *pgx.Batch, obj *storage.TestGrandparent_Embedded_Embedded2, testGrandparentID string, testGrandparentEmbeddedIdx int, idx int) error {

	values := []interface{}{
		// parent primary keys start
		testGrandparentID,
		testGrandparentEmbeddedIdx,
		idx,
		obj.GetVal(),
	}

	finalStr := "INSERT INTO test_grandparents_embeddeds_embedded2 (test_grandparents_Id, test_grandparents_embeddeds_idx, idx, Val) VALUES($1, $2, $3, $4) ON CONFLICT(test_grandparents_Id, test_grandparents_embeddeds_idx, idx) DO UPDATE SET test_grandparents_Id = EXCLUDED.test_grandparents_Id, test_grandparents_embeddeds_idx = EXCLUDED.test_grandparents_embeddeds_idx, idx = EXCLUDED.idx, Val = EXCLUDED.Val"
	batch.Queue(finalStr, values...)

	return nil
}

func (s *storeImpl) copyFromTestGrandparents(ctx context.Context, tx *postgres.Tx, objs ...*storage.TestGrandparent) error {

	inputRows := [][]interface{}{}

	var err error

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	var deletes []string

	copyCols := []string{

		"id",

		"val",

		"priority",

		"riskscore",

		"serialized",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		serialized, marshalErr := obj.Marshal()
		if marshalErr != nil {
			return marshalErr
		}

		inputRows = append(inputRows, []interface{}{

			obj.GetId(),

			obj.GetVal(),

			obj.GetPriority(),

			obj.GetRiskScore(),

			serialized,
		})

		// Add the ID to be deleted.
		deletes = append(deletes, obj.GetId())

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if err := s.DeleteMany(ctx, deletes); err != nil {
				return err
			}
			// clear the inserts and vals for the next batch
			deletes = nil

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"test_grandparents"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {
		_ = idx // idx may or may not be used depending on how nested we are, so avoid compile-time errors.

		if err = s.copyFromTestGrandparentsEmbeddeds(ctx, tx, obj.GetId(), obj.GetEmbedded()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromTestGrandparentsEmbeddeds(ctx context.Context, tx *postgres.Tx, testGrandparentID string, objs ...*storage.TestGrandparent_Embedded) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"test_grandparents_id",

		"idx",

		"val",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			testGrandparentID,

			idx,

			obj.GetVal(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"test_grandparents_embeddeds"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {
		_ = idx // idx may or may not be used depending on how nested we are, so avoid compile-time errors.

		if err = s.copyFromTestGrandparentsEmbeddedsEmbedded2(ctx, tx, testGrandparentID, idx, obj.GetEmbedded2()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromTestGrandparentsEmbeddedsEmbedded2(ctx context.Context, tx *postgres.Tx, testGrandparentID string, testGrandparentEmbeddedIdx int, objs ...*storage.TestGrandparent_Embedded_Embedded2) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"test_grandparents_id",

		"test_grandparents_embeddeds_idx",

		"idx",

		"val",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			testGrandparentID,

			testGrandparentEmbeddedIdx,

			idx,

			obj.GetVal(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"test_grandparents_embeddeds_embedded2"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

func (s *storeImpl) copyFrom(ctx context.Context, objs ...*storeType) error {
	conn, err := s.AcquireConn(ctx, ops.Get)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if err := s.copyFromTestGrandparents(ctx, tx, objs...); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (s *storeImpl) upsert(ctx context.Context, objs ...*storeType) error {
	conn, err := s.AcquireConn(ctx, ops.Get)
	if err != nil {
		return err
	}
	defer conn.Release()

	for _, obj := range objs {
		batch := &pgx.Batch{}
		if err := insertIntoTestGrandparents(ctx, batch, obj); err != nil {
			return err
		}
		batchResults := conn.SendBatch(ctx, batch)
		var result *multierror.Error
		for i := 0; i < batch.Len(); i++ {
			_, err := batchResults.Exec()
			result = multierror.Append(result, err)
		}
		if err := batchResults.Close(); err != nil {
			return err
		}
		if err := result.ErrorOrNil(); err != nil {
			return err
		}
	}
	return nil
}

// endregion Helper functions
// region Interface functions

// Upsert saves the current state of an object in storage.
func (s *storeImpl) Upsert(ctx context.Context, obj *storeType) error {
	defer metricsSetPostgresOperationDurationTime(time.Now(), ops.Upsert)

	scopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_WRITE_ACCESS).Resource(targetResource)
	if !scopeChecker.IsAllowed() {
		return sac.ErrResourceAccessDenied
	}

	return pgutils.Retry(func() error {
		return s.upsert(ctx, obj)
	})
}

// UpsertMany saves the state of multiple objects in the storage.
func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storeType) error {
	defer metricsSetPostgresOperationDurationTime(time.Now(), ops.UpdateMany)

	scopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_WRITE_ACCESS).Resource(targetResource)
	if !scopeChecker.IsAllowed() {
		return sac.ErrResourceAccessDenied
	}

	return pgutils.Retry(func() error {
		// Lock since copyFrom requires a delete first before being executed.  If multiple processes are updating
		// same subset of rows, both deletes could occur before the copyFrom resulting in unique constraint
		// violations
		if len(objs) < batchAfter {
			s.mutex.RLock()
			defer s.mutex.RUnlock()

			return s.upsert(ctx, objs...)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()

		return s.copyFrom(ctx, objs...)
	})
}

// endregion Interface functions

// region Used for testing

// CreateTableAndNewStore returns a new Store instance for testing.
func CreateTableAndNewStore(ctx context.Context, db postgres.DB, gormDB *gorm.DB) Store {
	pkgSchema.ApplySchemaForTable(ctx, gormDB, baseTable)
	return New(db)
}

// Destroy drops the tables associated with the target object type.
func Destroy(ctx context.Context, db postgres.DB) {
	dropTableTestGrandparents(ctx, db)
}

func dropTableTestGrandparents(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS test_grandparents CASCADE")
	dropTableTestGrandparentsEmbeddeds(ctx, db)

}

func dropTableTestGrandparentsEmbeddeds(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS test_grandparents_embeddeds CASCADE")
	dropTableTestGrandparentsEmbeddedsEmbedded2(ctx, db)

}

func dropTableTestGrandparentsEmbeddedsEmbedded2(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS test_grandparents_embeddeds_embedded2 CASCADE")

}

// endregion Used for testing
