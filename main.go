package main

import (
	"fmt"
	"log"

	_ "tagd/docs" // Import auto-generated docs package

	"tagd/common"
	"tagd/handlers"
	"tagd/models"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title Tag Management
// @version 1.0
// @description Tag Management System
// @termsOfService http://zgsm.ai
// @contact.name Bochun Zheng
// @contact.url http://zgsm.ai
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /tagd/api/v1
// @query.collection.format multi
func main() {
	var c = &common.Config{}
	// Initialize configuration file
	if err := c.Init("./env.yaml"); err != nil {
		panic(fmt.Errorf("配置文件初始化失败:%v", err))
	}
	dbName := c.Db.DatabaseName
	if dbName == "" {
		dbName = "tagd"
	}
	dbName = fmt.Sprintf("%s.db", dbName)
	// Initialize database
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Auto migrate models
	db.AutoMigrate(&models.TagPosition{})
	db.AutoMigrate(&models.Tag{})

	// Initialize gin router
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Register routes
	registerRoutes(r, db)

	r.Run(c.Server.ListenAddr)
}

func registerRoutes(r *gin.Engine, db *gorm.DB) {
	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tagHandler := handlers.NewTagHandler(db)
	// Tag management routes
	tags := r.Group("/tagd/api/v1/tags")
	{
		tags.GET("", tagHandler.GetTags)
		tags.GET("/:tagid", tagHandler.GetTag)
		tags.POST("", tagHandler.AddTag)
		tags.PUT("/:tagid", tagHandler.UpdateTag)
		tags.PUT("/:tagid/:key", tagHandler.UpdateTagPair)
		tags.DELETE("/:tagid", tagHandler.DeleteTag)
	}
}
