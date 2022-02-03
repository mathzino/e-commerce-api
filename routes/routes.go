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

//category
    r.GET("/product-categories", controllers.GetAllProductCategory)
    r.GET("/product-categories/:id", controllers.GetCategoryById)
    // r.GET("/age-rating-categories/:id/movies", controllers.GetMoviesByRatingId)

    categoryMiddlewareRoute := r.Group("/product-categories")
    categoryMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    categoryMiddlewareRoute.POST("/", controllers.CreateCategory)
    categoryMiddlewareRoute.PATCH("/:id", controllers.UpdateCategory)
    categoryMiddlewareRoute.DELETE("/:id", controllers.DeleteProductCategory)
//product
    r.GET("/product", controllers.GetAllProduct)
    r.GET("/product/:id", controllers.GetProductById)
    //with auth
    productMiddlewareRoute := r.Group("/product")
    productMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    productMiddlewareRoute.POST("/", controllers.CreateProduct)
    productMiddlewareRoute.PATCH("/:id", controllers.UpdateProduct)
    productMiddlewareRoute.DELETE("/:id", controllers.DeleteProduct)
    
//seller
    r.GET("/seller", controllers.GetAllSeleer)
    r.GET("/seller/:id", controllers.GetSellerById)
    //with auth
    SellerMiddlewareRoute := r.Group("/seller")
    SellerMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    SellerMiddlewareRoute.POST("/", controllers.CreateSeller)
    SellerMiddlewareRoute.PATCH("/:id", controllers.UpdateSeller)
    SellerMiddlewareRoute.DELETE("/:id", controllers.DeleteSeller)

   
//cart item
    r.GET("/cart-item", controllers.GetAllCartItem)
    r.GET("/cart-item/:id", controllers.GetCartItemByIdUser)
    //with auth
   CartItemMiddlewareRoute := r.Group("/cart-item")
   CartItemMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
   CartItemMiddlewareRoute.POST("/", controllers.CreateCartItem)
   CartItemMiddlewareRoute.PATCH("/:id", controllers.UpdateCartItem)
   CartItemMiddlewareRoute.DELETE("/:id", controllers.DeleteCartItem)
//Discussion
    r.GET("/discussion", controllers.GetAllDiscussion)
    r.GET("/discussion/:id", controllers.GetDiscussionById)
    //with auth
   DiscussionMiddlewareRoute := r.Group("/discussion")
   DiscussionMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
   DiscussionMiddlewareRoute.POST("/", controllers.CreateDiscussion)
   DiscussionMiddlewareRoute.PATCH("/:id", controllers.UpdateDiscussion)
   DiscussionMiddlewareRoute.DELETE("/:id", controllers.DeleteDiscussion)

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}