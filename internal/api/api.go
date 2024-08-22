package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-malawi/qatarina/internal/config"
	"github.com/golang-malawi/qatarina/internal/database/dbsqlc"
	"github.com/golang-malawi/qatarina/internal/logging"
	"github.com/golang-malawi/qatarina/internal/repository"
	"github.com/golang-malawi/qatarina/internal/services"
	"github.com/jackc/pgx/v5"
	"github.com/riverqueue/river"
)

type API struct {
	logger           logging.Logger
	app              *fiber.App
	Config           *config.Config
	RiverClient      *river.Client[pgx.Tx]
	AuthService      services.AuthService
	OrgRepo          repository.OrgRepository
	UserService      services.UserService
	ProjectsService  services.ProjectService
	TestersRepo      repository.TesterRepository
	TestCasesService services.TestCaseService
	TestPlansService services.TestPlanService
	TestRunsService  services.TestRunService
	TesterService    services.TesterService
}

func NewAPI(config *config.Config) *API {

	dbConn := dbsqlc.New(config.OpenDB())
	logger := logging.NewFromConfig(&config.Logging)

	return &API{
		logger:           logger,
		app:              fiber.New(),
		Config:           config,
		AuthService:      services.NewAuthService(&config.Auth, dbConn, logger),
		ProjectsService:  services.NewProjectService(dbConn, logger),
		TestCasesService: services.NewTestCaseService(dbConn, logger),
		TestPlansService: services.NewTestPlanService(dbConn, logger),
		TestRunsService:  services.NewTestRunService(dbConn, logger),
		OrgRepo:          nil,
		UserService:      nil,
		TesterService:    nil,
		TestersRepo:      nil,
	}
}

func (api *API) registerRoutes() {
	api.middleware()
	api.routes()
}

func (api *API) Start(address string) error {
	api.registerRoutes()
	api.logger.Debug("Starting API on ", "address", address)
	return api.app.Listen(address)
}
