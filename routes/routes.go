package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fp/controllers"
	"fp/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // set db to gin context
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    // r.GET("/movies", controllers.GetAllMovie)
    // r.GET("/:id", controllers.GetMovieById)

    // moviesMiddlewareRoute := r.Group("/movies")
    // moviesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    // moviesMiddlewareRoute.POST("/", controllers.CreateMovie)
    // moviesMiddlewareRoute.PATCH("/:id", controllers.UpdateMovie)
    // moviesMiddlewareRoute.DELETE("/:id", controllers.DeleteMovie)

    r.GET("/product-categories", controllers.GetAllProductCategory)
    r.GET("/product-categories/:id", controllers.GetCategoryById)
    // r.GET("/age-rating-categories/:id/movies", controllers.GetMoviesByRatingId)

    categoryMiddlewareRoute := r.Group("/product-categories")
    categoryMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    categoryMiddlewareRoute.POST("/", controllers.CreateCategory)
    categoryMiddlewareRoute.PATCH("/:id", controllers.UpdateCategory)
    categoryMiddlewareRoute.DELETE("/:id", controllers.DeleteProductCategory)

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}