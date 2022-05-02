package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/snkonoplev/file-manager/docs"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Swagger struct {
	engine *gin.Engine
	config *viper.Viper
}

func NewSwagger(engine *gin.Engine, config *viper.Viper) *Swagger {
	return &Swagger{
		engine: engine,
		config: config,
	}
}

func (s *Swagger) EnableSwagger() {
	docs.SwaggerInfo.Title = "File Manager API"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = s.config.GetString("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = s.config.GetString("SWAGGER_BASE_PATH")
	docs.SwaggerInfo.Schemes = []string{s.config.GetString("SWAGGER_SCHEMA")}

	if s.config.GetBool("SWAGGER_ENABLED") {
		urlString := s.config.GetString("SWAGGER_SCHEMA") + "://" + s.config.GetString("SWAGGER_HOST")
		basePath := s.config.GetString("SWAGGER_BASE_PATH")

		if basePath != "" {
			urlString = urlString + "/" + basePath
		}

		urlString = urlString + "/swagger/doc.json"

		url := ginSwagger.URL(urlString)
		s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
