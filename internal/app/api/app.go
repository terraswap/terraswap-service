package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/delight-labs/terraswap-service/configs"
	"github.com/evalphobia/logrus_sentry"
	"github.com/gin-contrib/cors"
	"github.com/sirupsen/logrus"

	"github.com/delight-labs/terraswap-service/internal/app/api/common/terraswap"
	"github.com/delight-labs/terraswap-service/internal/app/api/pair"
	"github.com/delight-labs/terraswap-service/internal/app/api/token"
	"github.com/delight-labs/terraswap-service/internal/app/api/tx"
	"github.com/delight-labs/terraswap-service/internal/pkg/cache"
	"github.com/delight-labs/terraswap-service/internal/pkg/logging"
	"github.com/delight-labs/terraswap-service/internal/pkg/repeater"
	terra "github.com/delight-labs/terraswap-service/internal/pkg/terraswap"
	tsCache "github.com/delight-labs/terraswap-service/internal/pkg/terraswap/cache"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap/databases/grpc"
	"github.com/delight-labs/terraswap-service/internal/pkg/terraswap/router"
	"github.com/gin-gonic/gin"
)

type terraswapApi struct {
	db     terraswap.DataHandler
	engine *gin.Engine
	config configs.Config
	logger logging.Logger
}

func RunServer(c configs.Config) *terraswapApi {
	cacheStore := cache.New(c.Cache)
	terraswapCache := tsCache.New(cacheStore)

	routerRepo := router.NewRepo(terraswapCache)
	routerService := router.New(routerRepo, c)

	grpcClient := grpc.New(c.Terraswap.GrpcHost, c.Terraswap.ChainId, c.Log)
	terraswapRepo := terraswap.NewRepo(c.Terraswap.ChainId, grpcClient)

	tsHandler := terraswap.NewDataHandler(terraswapRepo, routerService, terraswapCache, c)

	tsHandler.Run()
	routerService.Run()

	repeater.Enroll(tsHandler.GetLogger(), tsHandler, "TerraswapDataHandler", 2, terra.BLOCK_TIME)
	repeater.Enroll(tsHandler.GetLogger(), routerService, "TerraswapRouter", 2, terra.BLOCK_TIME*10)

	gin.SetMode(c.App.Mode)
	app := terraswapApi{
		db:     tsHandler,
		engine: gin.Default(),
		config: c,
		logger: logging.New(c.App.Name, c.Log),
	}

	app.setMiddlewares()
	app.setControllers()
	if c.Sentry.DSN != "" {
		app.configureReporter(c.Sentry.DSN)
	}
	app.run()

	return &app
}

func (app *terraswapApi) run() {

	type NotFound struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	app.engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, NotFound{Code: http.StatusNotFound, Message: "Not Found"})
	})
	app.engine.Run(fmt.Sprintf(":%s", strconv.Itoa(app.config.App.Port)))
}

func (app *terraswapApi) setMiddlewares() {
	app.engine.Use(gin.CustomRecovery(codedErrorHandle))

	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{"https://app-classic.terraswap.io", "https://app.terraswap.io", "https://app-dev.terraswap.io", "http://127.0.0.1", "http://localhost"}
	conf.AllowMethods = []string{"GET", "OPTIONS"}

	app.engine.Use(cors.New(conf))

}

func (app *terraswapApi) setControllers() {
	token.Init(app.db, app.engine)
	pair.Init(app.db, app.engine)
	tx.Init(app.db, app.engine, app.logger)
}

func (app *terraswapApi) configureReporter(dsn string) error {
	hook, err := logrus_sentry.NewSentryHook(dsn, []logrus.Level{
		logrus.WarnLevel,
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})
	if err != nil {
		return err
	}
	hook.StacktraceConfiguration.Enable = true
	logging.AddHookToLogger(app.logger, hook)
	return nil
}
