package api

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/evalphobia/logrus_sentry"
	"github.com/gin-contrib/cors"
	"github.com/sirupsen/logrus"
	"github.com/terraswap/terraswap-service/configs"

	"github.com/gin-gonic/gin"
	"github.com/terraswap/terraswap-service/internal/app/api/common/terraswap"
	"github.com/terraswap/terraswap-service/internal/app/api/pair"
	"github.com/terraswap/terraswap-service/internal/app/api/token"
	"github.com/terraswap/terraswap-service/internal/app/api/tx"
	"github.com/terraswap/terraswap-service/internal/pkg/cache"
	"github.com/terraswap/terraswap-service/internal/pkg/logging"
	"github.com/terraswap/terraswap-service/internal/pkg/repeater"
	terra "github.com/terraswap/terraswap-service/internal/pkg/terraswap"
	tsCache "github.com/terraswap/terraswap-service/internal/pkg/terraswap/cache"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap/databases/grpc"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap/databases/rdb"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap/router"
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

	var tsHandler terraswap.DataHandler
	isClassic := terra.IsClassic(c.Terraswap.ChainId)
	if isClassic {
		var db rdb.TerraswapRdb
		if c.Rdb.Host != "" {
			db = rdb.New(c.Rdb)
		}
		grpcClient := grpc.NewClassic(c.Terraswap.GrpcHost, c.Terraswap.ChainId, c.Terraswap.Version, c.Terraswap.InsecureConnection, c.Log)
		terraswapRepo := terraswap.NewClassicRepo(c.Terraswap.ChainId, grpcClient, db)
		tsHandler = terraswap.NewDataHandler(terraswapRepo, routerService, terraswapCache, c)
	} else {
		grpcClient := grpc.New(c.Terraswap.GrpcHost, c.Terraswap.ChainId, c.Terraswap.InsecureConnection, c.Log)
		terraswapRepo := terraswap.NewRepo(c.Terraswap.ChainId, grpcClient)
		tsHandler = terraswap.NewDataHandler(terraswapRepo, routerService, terraswapCache, c)
	}

	tsHandler.Run()
	routerService.Run()

	repeater.Enroll(tsHandler.GetLogger(), tsHandler, "TerraswapDataHandler", 2, terra.BLOCK_TIME*100)
	repeater.Enroll(tsHandler.GetLogger(), routerService, "TerraswapRouter", 2, terra.BLOCK_TIME*200)

	gin.SetMode(c.App.Mode)
	app := terraswapApi{
		db:     tsHandler,
		engine: gin.Default(),
		config: c,
		logger: logging.New(c.App.Name, c.Log),
	}

	app.setMiddlewares()
	apiVersion := c.Terraswap.Version
	app.setControllers(isClassic, apiVersion)
	if isClassic {
		app.setControllers(isClassic, "")
	}

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

	allowedOrigins := []string{`\.terraswap\.io$`, `terraswap\.netlify\.app$`, `localhost$`, `127\.0\.0\.1$`}
	conf := cors.DefaultConfig()
	conf.AllowOriginFunc = func(origin string) bool {
		for _, o := range allowedOrigins {
			matched, _ := regexp.MatchString(o, origin)
			if matched {
				return true
			}
		}
		return false
	}
	conf.AllowMethods = []string{"GET", "OPTIONS"}

	app.engine.Use(cors.New(conf))

}

func (app *terraswapApi) setControllers(isClassic bool, version string) {
	router := app.engine.Group(version)
	token.Init(app.db, router, isClassic)
	pair.Init(app.db, router)
	tx.Init(app.db, router, app.logger, isClassic)
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
