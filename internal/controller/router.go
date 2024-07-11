package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mlchain/dify-sandbox/internal/middleware"
	"github.com/mlchain/dify-sandbox/internal/static"
)

func Setup(eng *gin.Engine) {
	eng.Use(middleware.Auth())

	eng.POST(
		"/v1/sandbox/run",
		middleware.MaxRequest(static.GetDifySandboxGlobalConfigurations().MaxRequests),
		middleware.MaxWorker(static.GetDifySandboxGlobalConfigurations().MaxWorkers),
		RunSandboxController,
	)
	eng.GET(
		"/v1/sandbox/dependencies",
		GetDependencies,
	)

	eng.POST(
		"/v1/sandbox/dependencies/update",
		UpdateDependencies,
	)
}
