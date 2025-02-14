package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang-malawi/qatarina/internal/database/dbsqlc"
	"github.com/golang-malawi/qatarina/internal/logging"
	"github.com/golang-malawi/qatarina/internal/schema"
	"github.com/google/uuid"
)

// TestCaseService contains functionality to manage services
type TestCaseService interface {
	// FindAll retrieves all test cases in the database
	FindAll(context.Context) ([]dbsqlc.TestCase, error)

	// FindAllByProjectID retrieves all test cases in the database by Project ID
	FindAllByProjectID(context.Context, int64) ([]dbsqlc.TestCase, error)

	// FindAllCreatedBy retrieves all test cases in the database created by a specific user
	FindAllCreatedBy(context.Context, int64) ([]dbsqlc.TestCase, error)

	// Create creates a new test case
	Create(context.Context, *schema.CreateTestCaseRequest) (*dbsqlc.TestCase, error)

	// BulkCreate creates test cases in bulk. Only valid test cases are created.
	BulkCreate(context.Context, *schema.BulkCreateTestCases) ([]dbsqlc.TestCase, error)

	// Update updates a test case in the system
	Update(context.Context, *schema.UpdateTestCaseRequest) (*dbsqlc.TestCase, error)

	// DeleteByID deletes a single test case by ID
	DeleteByID(context.Context, string) error

	// DeleteByID delets all test cases linked to the given ProjectID
	DeleteByProjectID(context.Context, string) error

	// DeleteByTestRunID delets all test cases linked to the given TestRun
	DeleteByTestRunID(context.Context, string) error

	// BulkDelete deletes multiple test-cases by ID
	BulkDelete(context.Context, []string) error
}

var _ TestCaseService = &testCaseServiceImpl{}

type testCaseServiceImpl struct {
	queries *dbsqlc.Queries
	logger  logging.Logger
}

func NewTestCaseService(conn *dbsqlc.Queries, logger logging.Logger) TestCaseService {
	return &testCaseServiceImpl{
		queries: conn,
		logger:  logger,
	}
}

// BulkCreate implements TestCaseService.
func (t *testCaseServiceImpl) BulkCreate(ctx context.Context, bulkRequest *schema.BulkCreateTestCases) ([]dbsqlc.TestCase, error) {
	createList := make([]uuid.UUID, 0)
	for _, request := range bulkRequest.TestCases {
		uuidVal, _ := uuid.NewV7()
		params := dbsqlc.CreateTestCaseParams{
			ID:               uuidVal,
			Kind:             dbsqlc.TestKind(request.Kind),
			Code:             request.Code,
			FeatureOrModule:  sql.NullString{String: request.FeatureOrModule, Valid: true},
			Title:            request.Title,
			Description:      request.Description,
			ParentTestCaseID: sql.NullInt32{},
			IsDraft:          sql.NullBool{Bool: request.IsDraft, Valid: true},
			Tags:             request.Tags,
			CreatedByID:      1,
			CreatedAt:        sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedAt:        sql.NullTime{Time: time.Now(), Valid: true},
		}

		createdID, err := t.queries.CreateTestCase(ctx, params)
		if err != nil {
			return nil, err
		}
		createList = append(createList, createdID)
	}

	return []dbsqlc.TestCase{}, nil
}

// BulkDelete implements TestCaseService.
func (t *testCaseServiceImpl) BulkDelete(context.Context, []string) error {
	panic("unimplemented")
}

// Create implements TestCaseService.
func (t *testCaseServiceImpl) Create(context.Context, *schema.CreateTestCaseRequest) (*dbsqlc.TestCase, error) {
	panic("unimplemented")
}

// DeleteByID implements TestCaseService.
func (t *testCaseServiceImpl) DeleteByID(context.Context, string) error {
	panic("unimplemented")
}

// DeleteByProjectID implements TestCaseService.
func (t *testCaseServiceImpl) DeleteByProjectID(context.Context, string) error {
	panic("unimplemented")
}

// DeleteByTestRunID implements TestCaseService.
func (t *testCaseServiceImpl) DeleteByTestRunID(context.Context, string) error {
	panic("unimplemented")
}

// FindAll implements TestCaseService.
func (t *testCaseServiceImpl) FindAll(ctx context.Context) ([]dbsqlc.TestCase, error) {
	return t.queries.ListTestCases(ctx)
}

// FindAllByProjectID implements TestCaseService.
func (t *testCaseServiceImpl) FindAllByProjectID(context.Context, int64) ([]dbsqlc.TestCase, error) {
	panic("unimplemented")
}

// FindAllCreatedBy implements TestCaseService.
func (t *testCaseServiceImpl) FindAllCreatedBy(context.Context, int64) ([]dbsqlc.TestCase, error) {
	panic("unimplemented")
}

// Update implements TestCaseService.
func (t *testCaseServiceImpl) Update(context.Context, *schema.UpdateTestCaseRequest) (*dbsqlc.TestCase, error) {
	panic("unimplemented")
}
