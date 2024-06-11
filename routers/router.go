package routers

import (
	"hmdp-go/repository/repository_impl"
	"hmdp-go/service/impl"
	"log"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"

	"hmdp-go/common/datasource"
	"hmdp-go/common/logger"
	"hmdp-go/common/middleware/cors"
	"hmdp-go/common/middleware/jwt"
	"hmdp-go/common/setting"
	"hmdp-go/controller"
)

// InitRouter 初始化Router
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())
	gin.SetMode(setting.Config.APP.RunMode)
	Configure(r)
	return r
}

// Configure 配置router
func Configure(r *gin.Engine) {
	//controller declare
	var (
		blogController controller.Blog
		//blogCommentsController controller.BlogComments
		//followController       controller.Follow
		uploadController   controller.Upload
		shopController     controller.Shop
		shopTypeController controller.ShopType
		//signController         controller.Sign
		userController         controller.User
		voucherController      controller.Voucher
		voucherOrderController controller.VoucherOrder
		myJwt                  jwt.JWT
	)
	db := datasource.Db{}
	zap := logger.Logger{}
	//zap log init
	zap.Init()
	//database connect
	if err := db.Connect(); err != nil {
		log.Fatal("db fatal:", err)
	}
	//Injection
	var injector inject.Graph
	if err := injector.Provide(
		&inject.Object{Value: &db},
		&inject.Object{Value: &zap},
		&inject.Object{Value: &myJwt},

		// repository
		&inject.Object{Value: &repository_impl.BaseRepository{}},
		&inject.Object{Value: &repository_impl.BlogRepository{}},
		&inject.Object{Value: &repository_impl.BlogCommentsRepository{}},
		&inject.Object{Value: &repository_impl.FollowRepository{}},
		&inject.Object{Value: &repository_impl.SeckillVoucherRepository{}},
		&inject.Object{Value: &repository_impl.ShopRepository{}},
		&inject.Object{Value: &repository_impl.ShopTypeRepository{}},
		&inject.Object{Value: &repository_impl.SignRepository{}},
		&inject.Object{Value: &repository_impl.UserRepository{}},
		&inject.Object{Value: &repository_impl.UserInfoRepository{}},
		&inject.Object{Value: &repository_impl.VoucherRepository{}},
		&inject.Object{Value: &repository_impl.VoucherOrderRepository{}},

		// service
		&inject.Object{Value: &impl.BlogService{}},
		&inject.Object{Value: &impl.BlogCommentsService{}},
		&inject.Object{Value: &impl.FollowService{}},
		&inject.Object{Value: &impl.SeckillVoucherService{}},
		&inject.Object{Value: &impl.ShopService{}},
		&inject.Object{Value: &impl.ShopTypeService{}},
		&inject.Object{Value: &impl.SignService{}},
		&inject.Object{Value: &impl.UserService{}},
		&inject.Object{Value: &impl.UserInfoService{}},
		&inject.Object{Value: &impl.VoucherService{}},
		&inject.Object{Value: &impl.VoucherOrderService{}},

		// controller
		//&inject.Object{Value: &blogController},
		//&inject.Object{Value: &blogCommentsController},
		//&inject.Object{Value: &followController},
		//&inject.Object{Value: &uploadController},
		//&inject.Object{Value: &shopController},
		//&inject.Object{Value: &shopTypeController},
		//&inject.Object{Value: &signController},
		//&inject.Object{Value: &userController},
		//&inject.Object{Value: &voucherController},
		//&inject.Object{Value: &voucherOrderController},
	); err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("injector fatal: ", err)
	}
	var authMiddleware = myJwt.GinJWTMiddlewareInit(&jwt.AllUserAuthorizator{})
	r.NoRoute(authMiddleware.MiddlewareFunc(), jwt.NoRouteHandler)
	r.POST("/login", authMiddleware.LoginHandler)
	userAPI := r.Group("/user")
	{
		userAPI.GET("/refresh_token", authMiddleware.RefreshHandler)
		userAPI.POST("/code", userController.SendCode)
	}
	userAPI.Use(authMiddleware.MiddlewareFunc())
	{
		userAPI.GET("/logout", userController.Logout)
		userAPI.GET("/me", userController.Me)
		userAPI.GET("/info/:id", userController.Info)
	}

	blogCommentsApi := r.Group("/blog-comments")
	blogCommentsApi.Use(authMiddleware.MiddlewareFunc())
	{
	}

	blogApi := r.Group("/blog")
	blogApi.Use(authMiddleware.MiddlewareFunc())
	{
		blogApi.POST("", blogController.SaveBlog)
		blogApi.PUT("/like/:id", blogController.LikeBlog)
		blogApi.GET("/of/me", blogController.QueryMyBlog)
		blogApi.GET("/hot", blogController.QueryHotBlog)
	}

	followApi := r.Group("/follow")
	followApi.Use(authMiddleware.MiddlewareFunc())
	{
	}

	shopApi := r.Group("/shop")
	shopApi.Use(authMiddleware.MiddlewareFunc())
	{
		shopApi.GET("/fix_conflict/:id", shopController.QueryShopById)
		shopApi.GET("/of/type", shopController.QueryShopByType)
		shopApi.GET("/of/name", shopController.QueryShopByName)
		shopApi.POST("", shopController.SaveShop)
		shopApi.PUT("", shopController.UpdateShop)
	}

	shopTypeApi := r.Group("/shop-type")
	{
		shopTypeApi.GET("/list", shopTypeController.QueryTypeList)
	}

	uploadApi := r.Group("/upload")
	uploadApi.Use(authMiddleware.MiddlewareFunc())
	{
		uploadApi.POST("/blog", uploadController.UploadImage)
		uploadApi.GET("/blog/delete", uploadController.DeleteBlogImg)
	}

	voucherApi := r.Group("/voucher")
	voucherApi.Use(authMiddleware.MiddlewareFunc())
	{
		voucherApi.POST("/", voucherController.AddVoucher)
		voucherApi.POST("/seckill", voucherController.AddSeckillVoucher)
		voucherApi.GET("/list/:shopId", voucherController.QueryVoucherOfShop)
	}

	voucherOrderApi := r.Group("/voucher-order")
	voucherOrderApi.Use(authMiddleware.MiddlewareFunc())
	{
		voucherOrderApi.POST("/", voucherOrderController.SeckillVoucher)
	}
}
