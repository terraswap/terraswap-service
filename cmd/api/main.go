package main

import (
	"os"
	"runtime/debug"

	"github.com/davecgh/go-spew/spew"
	"github.com/terraswap/terraswap-service/configs"
	"github.com/terraswap/terraswap-service/internal/app/api"
	"github.com/terraswap/terraswap-service/internal/pkg/logging"
)

func main() {
	config := configs.New()
	logger := logging.New("main", config.Log)
	if config.Sentry.DSN != "" {
		logging.ConfigureReporter(logger, config.Sentry.DSN)
	}
	defer catch(logger)

	api.RunServer(config)
}

func catch(logger logging.Logger) {
	recovered := recover()

	if recovered != nil {
		defer os.Exit(1)

		err, ok := recovered.(error)
		if !ok {
			logger.Errorf("could not convert recovered error into error: %s\n", spew.Sdump(recovered))
			return
		}

		stack := string(debug.Stack())
		logger.WithField("err", logging.NewErrorField(err)).WithField("stack", stack).Errorf("panic caught")
	}
}
