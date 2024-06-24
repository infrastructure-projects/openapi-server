package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"openapi-server/apis"
	"openapi-server/configs"
	"openapi-server/swagger"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))
	configs.LoadConfigs()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	apis.Export(engine)
	engine.StaticFS("/swagger", http.FS(swagger.FS))
	_ = engine.Run(":8000")
}
