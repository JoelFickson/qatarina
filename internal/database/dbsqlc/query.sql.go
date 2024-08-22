// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package dbsqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const countTestCasesNotLinkedToProject = `-- name: CountTestCasesNotLinkedToProject :one
SELECT COUNT(*) FROM test_cases
RIGHT OUTER JOIN test_plans p ON p.test_case_id = test_cases.id
WHERE p.project_id IS NULL
`

func (q *Queries) CountTestCasesNotLinkedToProject(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTestCasesNotLinkedToProject)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createNewTestRun = `-- name: CreateNewTestRun :one
INSERT INTO test_runs (
id, project_id, test_plan_id, test_case_id, owner_id, tested_by_id, assigned_to_id, code, created_at, updated_at,
result_state, is_closed, assignee_can_change_code, notes,reactions, tested_on, expected_result
)
VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
'pending', false, false, 'None', '{}'::jsonb, now(), 'Test to Pass'
)
RETURNING id
`

type CreateNewTestRunParams struct {
	ID           uuid.UUID
	ProjectID    int32
	TestPlanID   int32
	TestCaseID   uuid.UUID
	OwnerID      int32
	TestedByID   int32
	AssignedToID int32
	Code         string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}

func (q *Queries) CreateNewTestRun(ctx context.Context, arg CreateNewTestRunParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createNewTestRun,
		arg.ID,
		arg.ProjectID,
		arg.TestPlanID,
		arg.TestCaseID,
		arg.OwnerID,
		arg.TestedByID,
		arg.AssignedToID,
		arg.Code,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createProject = `-- name: CreateProject :one
INSERT INTO projects (
    title, description, version, is_active, is_public, website_url,
    github_url, trello_url, jira_url, monday_url,
    owner_user_id, created_at, updated_at, deleted_at
)
VALUES(
    $1, $2, $3, $4, $5, $6,
    $7, $8, $9, $10,
    $11, $12, $13, $14
) RETURNING id
`

type CreateProjectParams struct {
	Title       string
	Description string
	Version     sql.NullString
	IsActive    sql.NullBool
	IsPublic    sql.NullBool
	WebsiteUrl  sql.NullString
	GithubUrl   sql.NullString
	TrelloUrl   sql.NullString
	JiraUrl     sql.NullString
	MondayUrl   sql.NullString
	OwnerUserID int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createProject,
		arg.Title,
		arg.Description,
		arg.Version,
		arg.IsActive,
		arg.IsPublic,
		arg.WebsiteUrl,
		arg.GithubUrl,
		arg.TrelloUrl,
		arg.JiraUrl,
		arg.MondayUrl,
		arg.OwnerUserID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createTestCase = `-- name: CreateTestCase :one
INSERT INTO test_cases (
    id, kind, code, feature_or_module, title, description, parent_test_case_id,
    is_draft, tags, created_by_id, created_at, updated_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7,
    $8, $9, $10, $11, $12
)
RETURNING id
`

type CreateTestCaseParams struct {
	ID               uuid.UUID
	Kind             TestKind
	Code             string
	FeatureOrModule  sql.NullString
	Title            string
	Description      string
	ParentTestCaseID sql.NullInt32
	IsDraft          sql.NullBool
	Tags             []string
	CreatedByID      int32
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
}

func (q *Queries) CreateTestCase(ctx context.Context, arg CreateTestCaseParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createTestCase,
		arg.ID,
		arg.Kind,
		arg.Code,
		arg.FeatureOrModule,
		arg.Title,
		arg.Description,
		arg.ParentTestCaseID,
		arg.IsDraft,
		pq.Array(arg.Tags),
		arg.CreatedByID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createTestPlan = `-- name: CreateTestPlan :one
INSERT INTO test_plans (
    project_id, assigned_to_id, created_by_id, updated_by_id,
    kind, description, start_at, closed_at, scheduled_end_at,
    num_test_cases, num_failures, is_complete, is_locked, has_report,
    created_at, updated_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $7, $8,
    $9, $10, $11, $12, $13, $14, $15
)
RETURNING id
`

type CreateTestPlanParams struct {
	ProjectID      int32
	AssignedToID   int32
	CreatedByID    int32
	UpdatedByID    int32
	Kind           TestKind
	Description    sql.NullString
	StartAt        sql.NullTime
	ScheduledEndAt sql.NullTime
	NumTestCases   int32
	NumFailures    int32
	IsComplete     sql.NullBool
	IsLocked       sql.NullBool
	HasReport      sql.NullBool
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}

func (q *Queries) CreateTestPlan(ctx context.Context, arg CreateTestPlanParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createTestPlan,
		arg.ProjectID,
		arg.AssignedToID,
		arg.CreatedByID,
		arg.UpdatedByID,
		arg.Kind,
		arg.Description,
		arg.StartAt,
		arg.ScheduledEndAt,
		arg.NumTestCases,
		arg.NumFailures,
		arg.IsComplete,
		arg.IsLocked,
		arg.HasReport,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    first_name, last_name, display_name, email, password, phone,
    org_id, country_iso, city, address,
    is_activated, is_reviewed, is_super_admin, is_verified,
    last_login_at, email_confirmed_at, created_at, updated_at, deleted_at
)
VALUES(
    $1, $2, $3, $4, $5, $6,
    $7, $8, $9, $10,
    $11, $12, $13, $14,
    $15, $16, $17, $18, $19
)
RETURNING id
`

type CreateUserParams struct {
	FirstName        string
	LastName         string
	DisplayName      sql.NullString
	Email            string
	Password         string
	Phone            string
	OrgID            sql.NullInt32
	CountryIso       string
	City             sql.NullString
	Address          string
	IsActivated      sql.NullBool
	IsReviewed       sql.NullBool
	IsSuperAdmin     sql.NullBool
	IsVerified       sql.NullBool
	LastLoginAt      sql.NullTime
	EmailConfirmedAt sql.NullTime
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
	DeletedAt        sql.NullTime
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.DisplayName,
		arg.Email,
		arg.Password,
		arg.Phone,
		arg.OrgID,
		arg.CountryIso,
		arg.City,
		arg.Address,
		arg.IsActivated,
		arg.IsReviewed,
		arg.IsSuperAdmin,
		arg.IsVerified,
		arg.LastLoginAt,
		arg.EmailConfirmedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteAllTestPlansInProject = `-- name: DeleteAllTestPlansInProject :execrows
DELETE FROM test_plans WHERE project_id = $1
`

func (q *Queries) DeleteAllTestPlansInProject(ctx context.Context, projectID int32) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteAllTestPlansInProject, projectID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const deleteProject = `-- name: DeleteProject :execrows
DELETE FROM projects WHERE id = $1
`

func (q *Queries) DeleteProject(ctx context.Context, id int32) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteProject, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const deleteTestPlan = `-- name: DeleteTestPlan :execrows
DELETE FROM test_plans WHERE id = $1
`

func (q *Queries) DeleteTestPlan(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteTestPlan, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const findUserLoginByEmail = `-- name: FindUserLoginByEmail :one
SELECT id, display_name, email, password, last_login_at FROM users WHERE email = $1 AND is_activated AND deleted_at IS NULL
`

type FindUserLoginByEmailRow struct {
	ID          int32
	DisplayName sql.NullString
	Email       string
	Password    string
	LastLoginAt sql.NullTime
}

func (q *Queries) FindUserLoginByEmail(ctx context.Context, email string) (FindUserLoginByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, findUserLoginByEmail, email)
	var i FindUserLoginByEmailRow
	err := row.Scan(
		&i.ID,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.LastLoginAt,
	)
	return i, err
}

const getProject = `-- name: GetProject :one
SELECT id, title, description, version, is_active, is_public, website_url, github_url, trello_url, jira_url, monday_url, owner_user_id, created_at, updated_at, deleted_at FROM projects WHERE id = $1
`

func (q *Queries) GetProject(ctx context.Context, id int32) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Version,
		&i.IsActive,
		&i.IsPublic,
		&i.WebsiteUrl,
		&i.GithubUrl,
		&i.TrelloUrl,
		&i.JiraUrl,
		&i.MondayUrl,
		&i.OwnerUserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getTestCase = `-- name: GetTestCase :one
SELECT id, kind, code, feature_or_module, title, description, parent_test_case_id, is_draft, tags, created_by_id, created_at, updated_at FROM test_cases WHERE id = $1
`

func (q *Queries) GetTestCase(ctx context.Context, id uuid.UUID) (TestCase, error) {
	row := q.db.QueryRowContext(ctx, getTestCase, id)
	var i TestCase
	err := row.Scan(
		&i.ID,
		&i.Kind,
		&i.Code,
		&i.FeatureOrModule,
		&i.Title,
		&i.Description,
		&i.ParentTestCaseID,
		&i.IsDraft,
		pq.Array(&i.Tags),
		&i.CreatedByID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTestPlan = `-- name: GetTestPlan :one
SELECT id, project_id, assigned_to_id, created_by_id, updated_by_id, kind, description, start_at, closed_at, scheduled_end_at, num_test_cases, num_failures, is_complete, is_locked, has_report, created_at, updated_at FROM test_plans WHERE id = $1
`

func (q *Queries) GetTestPlan(ctx context.Context, id int64) (TestPlan, error) {
	row := q.db.QueryRowContext(ctx, getTestPlan, id)
	var i TestPlan
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.AssignedToID,
		&i.CreatedByID,
		&i.UpdatedByID,
		&i.Kind,
		&i.Description,
		&i.StartAt,
		&i.ClosedAt,
		&i.ScheduledEndAt,
		&i.NumTestCases,
		&i.NumFailures,
		&i.IsComplete,
		&i.IsLocked,
		&i.HasReport,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, display_name, email, password, phone, org_id, country_iso, city, address, is_activated, is_reviewed, is_super_admin, is_verified, last_login_at, email_confirmed_at, created_at, updated_at, deleted_at FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.Phone,
		&i.OrgID,
		&i.CountryIso,
		&i.City,
		&i.Address,
		&i.IsActivated,
		&i.IsReviewed,
		&i.IsSuperAdmin,
		&i.IsVerified,
		&i.LastLoginAt,
		&i.EmailConfirmedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const isTestCaseLinkedToProject = `-- name: IsTestCaseLinkedToProject :one
SELECT EXISTS(
    SELECT test_cases.id, test_cases.kind, code, feature_or_module, title, test_cases.description, parent_test_case_id, is_draft, tags, test_cases.created_by_id, test_cases.created_at, test_cases.updated_at, p.id, project_id, assigned_to_id, p.created_by_id, updated_by_id, p.kind, p.description, start_at, closed_at, scheduled_end_at, num_test_cases, num_failures, is_complete, is_locked, has_report, p.created_at, p.updated_at FROM test_cases
    INNER JOIN test_plans p ON p.test_case_id = test_cases.id
    WHERE p.project_id = $1
)
`

func (q *Queries) IsTestCaseLinkedToProject(ctx context.Context, projectID int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, isTestCaseLinkedToProject, projectID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const listProjects = `-- name: ListProjects :many
SELECT id, title, description, version, is_active, is_public, website_url, github_url, trello_url, jira_url, monday_url, owner_user_id, created_at, updated_at, deleted_at FROM projects ORDER BY created_at DESC
`

func (q *Queries) ListProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, listProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Version,
			&i.IsActive,
			&i.IsPublic,
			&i.WebsiteUrl,
			&i.GithubUrl,
			&i.TrelloUrl,
			&i.JiraUrl,
			&i.MondayUrl,
			&i.OwnerUserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTestCases = `-- name: ListTestCases :many
SELECT id, kind, code, feature_or_module, title, description, parent_test_case_id, is_draft, tags, created_by_id, created_at, updated_at FROM test_cases ORDER BY created_at DESC
`

func (q *Queries) ListTestCases(ctx context.Context) ([]TestCase, error) {
	rows, err := q.db.QueryContext(ctx, listTestCases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestCase
	for rows.Next() {
		var i TestCase
		if err := rows.Scan(
			&i.ID,
			&i.Kind,
			&i.Code,
			&i.FeatureOrModule,
			&i.Title,
			&i.Description,
			&i.ParentTestCaseID,
			&i.IsDraft,
			pq.Array(&i.Tags),
			&i.CreatedByID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTestCasesByCreator = `-- name: ListTestCasesByCreator :many
SELECT id, kind, code, feature_or_module, title, description, parent_test_case_id, is_draft, tags, created_by_id, created_at, updated_at FROM test_cases WHERE created_by_id = $1
`

func (q *Queries) ListTestCasesByCreator(ctx context.Context, createdByID int32) ([]TestCase, error) {
	rows, err := q.db.QueryContext(ctx, listTestCasesByCreator, createdByID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestCase
	for rows.Next() {
		var i TestCase
		if err := rows.Scan(
			&i.ID,
			&i.Kind,
			&i.Code,
			&i.FeatureOrModule,
			&i.Title,
			&i.Description,
			&i.ParentTestCaseID,
			&i.IsDraft,
			pq.Array(&i.Tags),
			&i.CreatedByID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTestCasesByProject = `-- name: ListTestCasesByProject :many
SELECT test_cases.id, test_cases.kind, code, feature_or_module, title, test_cases.description, parent_test_case_id, is_draft, tags, test_cases.created_by_id, test_cases.created_at, test_cases.updated_at, p.id, project_id, assigned_to_id, p.created_by_id, updated_by_id, p.kind, p.description, start_at, closed_at, scheduled_end_at, num_test_cases, num_failures, is_complete, is_locked, has_report, p.created_at, p.updated_at FROM test_cases
INNER JOIN test_plans p ON p.test_case_id = test_cases.id
WHERE p.project_id = $1
`

type ListTestCasesByProjectRow struct {
	ID               uuid.UUID
	Kind             TestKind
	Code             string
	FeatureOrModule  sql.NullString
	Title            string
	Description      string
	ParentTestCaseID sql.NullInt32
	IsDraft          sql.NullBool
	Tags             []string
	CreatedByID      int32
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
	ID_2             int64
	ProjectID        int32
	AssignedToID     int32
	CreatedByID_2    int32
	UpdatedByID      int32
	Kind_2           TestKind
	Description_2    sql.NullString
	StartAt          sql.NullTime
	ClosedAt         sql.NullTime
	ScheduledEndAt   sql.NullTime
	NumTestCases     int32
	NumFailures      int32
	IsComplete       sql.NullBool
	IsLocked         sql.NullBool
	HasReport        sql.NullBool
	CreatedAt_2      sql.NullTime
	UpdatedAt_2      sql.NullTime
}

func (q *Queries) ListTestCasesByProject(ctx context.Context, projectID int32) ([]ListTestCasesByProjectRow, error) {
	rows, err := q.db.QueryContext(ctx, listTestCasesByProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListTestCasesByProjectRow
	for rows.Next() {
		var i ListTestCasesByProjectRow
		if err := rows.Scan(
			&i.ID,
			&i.Kind,
			&i.Code,
			&i.FeatureOrModule,
			&i.Title,
			&i.Description,
			&i.ParentTestCaseID,
			&i.IsDraft,
			pq.Array(&i.Tags),
			&i.CreatedByID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.ProjectID,
			&i.AssignedToID,
			&i.CreatedByID_2,
			&i.UpdatedByID,
			&i.Kind_2,
			&i.Description_2,
			&i.StartAt,
			&i.ClosedAt,
			&i.ScheduledEndAt,
			&i.NumTestCases,
			&i.NumFailures,
			&i.IsComplete,
			&i.IsLocked,
			&i.HasReport,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTestPlans = `-- name: ListTestPlans :many
SELECT id, project_id, assigned_to_id, created_by_id, updated_by_id, kind, description, start_at, closed_at, scheduled_end_at, num_test_cases, num_failures, is_complete, is_locked, has_report, created_at, updated_at FROM test_plans ORDER BY created_at DESC
`

func (q *Queries) ListTestPlans(ctx context.Context) ([]TestPlan, error) {
	rows, err := q.db.QueryContext(ctx, listTestPlans)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestPlan
	for rows.Next() {
		var i TestPlan
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.AssignedToID,
			&i.CreatedByID,
			&i.UpdatedByID,
			&i.Kind,
			&i.Description,
			&i.StartAt,
			&i.ClosedAt,
			&i.ScheduledEndAt,
			&i.NumTestCases,
			&i.NumFailures,
			&i.IsComplete,
			&i.IsLocked,
			&i.HasReport,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTestPlansByProject = `-- name: ListTestPlansByProject :many
SELECT id, project_id, assigned_to_id, created_by_id, updated_by_id, kind, description, start_at, closed_at, scheduled_end_at, num_test_cases, num_failures, is_complete, is_locked, has_report, created_at, updated_at FROM test_plans WHERE project_id = $1
`

func (q *Queries) ListTestPlansByProject(ctx context.Context, projectID int32) ([]TestPlan, error) {
	rows, err := q.db.QueryContext(ctx, listTestPlansByProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestPlan
	for rows.Next() {
		var i TestPlan
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.AssignedToID,
			&i.CreatedByID,
			&i.UpdatedByID,
			&i.Kind,
			&i.Description,
			&i.StartAt,
			&i.ClosedAt,
			&i.ScheduledEndAt,
			&i.NumTestCases,
			&i.NumFailures,
			&i.IsComplete,
			&i.IsLocked,
			&i.HasReport,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, first_name, last_name, display_name, email, password, phone, org_id, country_iso, city, address, is_activated, is_reviewed, is_super_admin, is_verified, last_login_at, email_confirmed_at, created_at, updated_at, deleted_at FROM users
ORDER BY created_at DESC
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.DisplayName,
			&i.Email,
			&i.Password,
			&i.Phone,
			&i.OrgID,
			&i.CountryIso,
			&i.City,
			&i.Address,
			&i.IsActivated,
			&i.IsReviewed,
			&i.IsSuperAdmin,
			&i.IsVerified,
			&i.LastLoginAt,
			&i.EmailConfirmedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserLastLogin = `-- name: UpdateUserLastLogin :execrows
UPDATE users SET last_login_at = $1 WHERE id = $2 AND is_activated AND deleted_at IS NULL
`

type UpdateUserLastLoginParams struct {
	LastLoginAt sql.NullTime
	ID          int32
}

func (q *Queries) UpdateUserLastLogin(ctx context.Context, arg UpdateUserLastLoginParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, updateUserLastLogin, arg.LastLoginAt, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const userExists = `-- name: UserExists :one
SELECT EXISTS(SELECT id FROM users WHERE id = $1)
`

func (q *Queries) UserExists(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, userExists, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
