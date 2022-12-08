package routes

import (
	"backend/config"
	"backend/controller"
	"backend/middleware"
	"backend/repository"
	"backend/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	// adminRepository        repository.AdminRepository        = repository.NewAdminRepository(db)
	addressRepository      repository.AddressRepository      = repository.NewAddressRepository(db)
	shippingRepository     repository.ShippingRepository     = repository.NewShippingRepository(db)
	paymentRepository      repository.PaymentRepository      = repository.NewPaymentRepository(db)
	transactionRepository  repository.TransactionRepository  = repository.NewTransactionRepository(db)
	sourceOfFundRepository repository.SourceOfFundRepository = repository.NewSourceOfFundRepository(db)
	promoRepository        repository.PromoRepository        = repository.NewPromoRepository(db)
	userPromoRepository    repository.UserPromoRepository    = repository.NewUserPromoRepository(db)
	jwtService             service.JWTService                = service.NewJWTService()
	userService            service.UserService               = service.NewUserService(userRepository)
	// adminService           service.AdminService              = service.NewAdminService(adminRepository)
	addressService     service.AddressService     = service.NewAddressService(addressRepository)
	shippingService    service.ShippingService    = service.NewShippingService(shippingRepository)
	paymentService     service.PaymentService     = service.NewPaymentService(paymentRepository)
	transactionService service.TransactionService = service.NewTransactionService(transactionRepository, userRepository, sourceOfFundRepository, addressRepository, paymentRepository, shippingRepository, promoRepository, userPromoRepository)
	promoService       service.PromoService       = service.NewPromoService(promoRepository)
	userPromoService   service.UserPromoService   = service.NewUserPromoService(userPromoRepository)
	authService        service.AuthService        = service.NewAuthService(userRepository)
	authController     controller.AuthController  = controller.NewAuthController(authService, jwtService)
	userController     controller.UserController  = controller.NewUserController(userService, jwtService)
	// adminController    controller.AdminController    = controller.NewAdminController(adminService, jwtService)
	addressController     controller.AddressController     = controller.NewAddressController(addressService, jwtService)
	shippingController    controller.ShippingController    = controller.NewShippingController(shippingService, jwtService)
	promoController       controller.PromoController       = controller.NewPromoController(promoService, jwtService)
	userPromoController   controller.UserPromoController   = controller.NewUserPromoController(userPromoService, jwtService)
	paymentController     controller.PaymentController     = controller.NewPaymentController(paymentService, jwtService)
	transactionController controller.TransactionController = controller.NewTransactionController(transactionService, jwtService)
)

func CourierRoutes() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("/api/user", middleware.AuthorizeJWT(jwtService, userService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile/:id", userController.Update)
		userRoutes.POST("/top-up", transactionController.TopUp)
		userRoutes.GET("/transactions", transactionController.GetTransactions)
		userRoutes.GET("/shipping", shippingController.All)
		userRoutes.POST("/shipping", shippingController.Insert)
		userRoutes.GET("/shipping/:id", shippingController.FindByID)
		userRoutes.DELETE("/shipping/:id", shippingController.Delete)
		userRoutes.PUT("/shipping/:id", shippingController.Update)
		userRoutes.POST("/payment", transactionController.Payment)
		userRoutes.GET("/payment", paymentController.All)
		userRoutes.GET("/address", addressController.All)
		userRoutes.POST("/address", addressController.Insert)
		userRoutes.GET("/address/:id", addressController.FindByID)
		userRoutes.DELETE("/address/:id", addressController.Delete)
		userRoutes.PUT("/address/:id", addressController.Update)
		userRoutes.GET("/user-promo", userPromoController.All)
		userRoutes.POST("/user-promo", userPromoController.Insert)
		userRoutes.PUT("/user-promo/:id", userPromoController.Update)
	}

	adminRoutes := r.Group("/api/admin", middleware.AuthorizeJWT(jwtService, userService))
	{
		adminRoutes.GET("/address", addressController.AllAddresses)
		adminRoutes.GET("/shipping", shippingController.AllShippings)
		adminRoutes.PUT("/shipping/:id", shippingController.Update)
		adminRoutes.GET("/user", userController.AllUsers)
		adminRoutes.GET("/user/referrals", userController.AllUserReferrals)
		adminRoutes.GET("/promo", promoController.AllPromos)
		adminRoutes.GET("/promo/:id", promoController.FindByID)
		adminRoutes.PUT("/promo/:id", promoController.Update)
	}

	r.Static("./dist", "dist/")
	r.Run(":8080")
}
