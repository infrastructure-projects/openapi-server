package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Export(engine *gin.Engine) {
	engine.GET("/v3/api-docs/swagger-config", func(context *gin.Context) {
		var ruote routes
		_ = viper.UnmarshalKey("application", &ruote)
		request := context.Request
		config := &openapiDesc{
			ConfigUrl:         "/v3/api-docs/swagger-config",
			Oauth2RedirectUrl: fmt.Sprintf("http://%s/swagger-ui/oauth2-redirect.html", request.Host),
			ValidatorUrl:      "",
			Urls:              ruote.getUrlInfo(),
		}
		context.JSON(http.StatusOK, config)
	})
}
