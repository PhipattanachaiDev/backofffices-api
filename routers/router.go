package router

import (
	"net/http"

	"os"

	authControllers "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/authentication"
	battery "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/battery"
	cardReader "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/card-reader"
	certificatesController "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/certificates"
	customersController "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/customers"
	dealersController "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/dealers"
	gps "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/gps"
	v2 "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v2"
	"ezview.asia/ezview-web/ezview-lite-back-office/logger"

	gpsAntenna "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/gps-antenna"
	gsmAntenna "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/gsm-antenna"
	menusController "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/menus"
	sim "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/sim"

	tracker "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/tracker"
	trackerBom "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/tracker-bom"
	userControllers "ezview.asia/ezview-web/ezview-lite-back-office/controllers/v1/users"

	"ezview.asia/ezview-web/ezview-lite-back-office/middlewares"
	response "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// "golang.org/x/time/rate"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// ตั้งค่า CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // HTTP methods ที่อนุญาต
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Headers ที่อนุญาต
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Middleware for API versioning
	r.Use(middlewares.APIVersionMiddleware)

	// Logger middleware setup
	logDir := os.Getenv("LOG_DIR")
	if logDir == "" {
		logDir = "logs" // Default log directory
	}

	logger.SetupLogger(logDir)

	r.Use(logger.LoggerMiddleware(logger.GetLogger()))

	// JWT Middleware - Uncomment if you want to use JWT authentication
	// r.Use(middlewares.JWTMiddleware())

	// Create a rate limiter with a limit of 500 request per second
	// limiter := rate.NewLimiter(500, 500)

	// r.Use(func(c *gin.Context) {
	// 	if limiter.Allow() {
	// 		c.Next()
	// 	} else {
	// 		c.AbortWithStatus(http.StatusTooManyRequests)
	// 	}
	// })

	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		basePath = "/back-office-service"
	}

	// Swagger UI setup
	r.GET(basePath+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Base path for API routes

	api := r.Group(basePath + "/api")
	{
		// Version 1 API group
		v1Group := api.Group("/v1")
		{
			v1Group.POST("/login", authControllers.Login)

			certificatesGroup := v1Group.Group("/certificates")
			{
				certificatesGroup.GET("/pdf", certificatesController.CreatePDF)
				certificatesGroup.GET("/pdf/preview", certificatesController.PreviewPDF)
			}

			v1Group.Use(middlewares.JWTMiddleware())

			authGroup := v1Group.Group("/auth")
			{
				authGroup.POST("/register/user/customer", authControllers.RegisterUserCustomer)
			}

			v1Group.GET("/refresh-token", authControllers.RefreshToken)

			v1Group.GET("/menus", menusController.GetMenus)

			userGroup := v1Group.Group("/users")
			{
				userGroup.PUT("/", userControllers.UpdateUser)
				userGroup.DELETE("/:id", userControllers.DeleteUser)
				userGroup.GET("/profile", userControllers.GetUserProfile)
			}

			customerGroup := v1Group.Group("/customers")
			{
				customerGroup.GET("/individual/:id", customersController.GetIndividualCustomerById)
				customerGroup.GET("/juristic/:id", customersController.GetJuristicCustomerById)
				customerGroup.POST("/individual", customersController.CreateIndividualCustomer)
				customerGroup.PUT("/individual/:id", customersController.UpdateIndividualCustomer)
				customerGroup.PUT("/juristic/:id", customersController.UpdateJuristicCustomer)
				customerGroup.POST("/juristic", customersController.CreateJuristicCustomer)
				customerGroup.POST("/", customersController.GetCustomersByConditions)
				customerGroup.GET("/groups", customersController.GetCustomerGroups)
				customerGroup.GET("/status", customersController.GetCustomerStatus)
				customerGroup.GET("/contacts/:id", customersController.GetCustomerContactByCustomerId)
				customerGroup.POST("/contacts", customersController.CreateCustomerContact)
				customerGroup.PUT("/contacts", customersController.UpdateCustomerContact)
				customerGroup.DELETE("/contacts/:id", customersController.DeleteCustomerContact)
				customerGroup.GET("/others/:id", customersController.GetCustomerOthersByCustomerId)
				customerGroup.DELETE("/:id", customersController.DeleteCustomerById)
				customerGroup.GET("/users/:id", customersController.GetUsersCustomerById)
			}

			dealerGroup := v1Group.Group("/dealers")
			{
				dealerGroup.POST("/", dealersController.GetDealersByConditions)
				dealerGroup.GET("/groups", dealersController.GetDealerGroups)
				dealerGroup.GET("/status", dealersController.GetDealerStatus)
			}

			v1Group.POST("/logout", authControllers.Logout)
			v1Group.GET("/provinces", gps.GetProvince)
			v1Group.GET("/districts", gps.GetDistrict)
			v1Group.GET("/subdistricts", gps.GetSubDistrict)

			// Battery
			v1Group.GET("/battery/status", battery.GetBatteryStatus)
			v1Group.GET("/battery/general", battery.GetBatteryGeneralById)
			v1Group.POST("/battery/search", battery.SearchBattery)
			v1Group.POST("/battery/insert", battery.InsertBattery)
			v1Group.PUT("/battery/update", battery.UpdateBattery)

			// Sim
			v1Group.GET("/sim/status", sim.GetSimStatus)
			v1Group.GET("/sim/operator", sim.GetSimOperator)
			v1Group.GET("/sim/general", sim.GetSimGeneralById)
			v1Group.POST("/sim/search", sim.SearchSim)
			v1Group.POST("/sim/insert", sim.InsertSim)
			v1Group.PUT("/sim/update", sim.UpdateSim)

			// GPS Antenna
			v1Group.GET("/gps-antenna/status", gpsAntenna.GetGpsAntennaStatus)
			v1Group.GET("/gps-antenna/general", gpsAntenna.GetGpsAntennaGeneralById)
			v1Group.POST("/gps-antenna/search", gpsAntenna.SearchGpsAntenna)
			v1Group.POST("/gps-antenna/insert", gpsAntenna.InsertGpsAntenna)
			v1Group.PUT("/gps-antenna/update", gpsAntenna.UpdateGpsAntenna)

			//Gsm Antenna
			v1Group.GET("/gsm-antenna/status", gsmAntenna.GetGsmAntennaStatus)
			v1Group.GET("/gsm-antenna/general", gsmAntenna.GetGsmAntennaGeneralById)
			v1Group.POST("/gsm-antenna/search", gsmAntenna.SearchGsmAntenna)
			v1Group.POST("/gsm-antenna/insert", gsmAntenna.InsertGsmAntenna)
			v1Group.PUT("/gsm-antenna/update", gsmAntenna.UpdateGsmAntenna)

			// Card Reader
			v1Group.GET("/card-reader/status", cardReader.GetCardReaderStatus)
			v1Group.GET("/card-reader/general", cardReader.GetCardReaderGeneral)
			v1Group.GET("/card-reader/model", cardReader.GetCardReaderModel)
			v1Group.GET("/card-reader/brand", cardReader.GetCardReaderBrand)
			v1Group.POST("/card-reader/insert", cardReader.InsertCardReader)
			v1Group.PUT("/card-reader/update", cardReader.UpdateCardReader)
			v1Group.POST("/card-reader/search", cardReader.SearchCardReader)

			// Tracker Bom
			v1Group.GET("/tracker-bom/master", trackerBom.GetTrackerBom)
			v1Group.POST("/tracker-bom/insert", trackerBom.InsertTrackerBom)
			v1Group.PUT("/tracker-bom/update", trackerBom.UpdateTrackerBom)
			v1Group.GET("/tracker-bom/gps", trackerBom.GetGPS)
			v1Group.GET("/tracker-bom/gps-antenna", trackerBom.GetGpsAntenna)
			v1Group.GET("/tracker-bom/gsm-antenna", trackerBom.GetGsmAntenna)
			v1Group.GET("/tracker-bom/card-reader", trackerBom.GetCardReader)
			v1Group.GET("/tracker-bom/battery", trackerBom.GetBattery)
			v1Group.GET("/tracker-bom/sim", trackerBom.GetSim)
			v1Group.GET("/tracker-bom/general", trackerBom.GetTrackerBomGeneralByID)

			// Tracker
			v1Group.POST("/tracker/search", tracker.SearchTracker)
			v1Group.GET("/tracker/brand", tracker.GetTrackerBrand)
			v1Group.GET("/tracker/model", tracker.GetTrackerModel)
			v1Group.GET("/tracker/status", tracker.GetTrackerStatus)
			v1Group.PUT("/tracker/update", tracker.UpdateTracker)
			v1Group.POST("/tracker/insert", tracker.InsertTracker)
			v1Group.GET("/tracker/general", tracker.GetTrackerGeneralByTrackerId)

			// Gps
			v1Group.GET("/gps-status", gps.GetGpsStatus)
			v1Group.GET("/gps-brands", gps.GetGpsBrands)
			v1Group.GET("/gps-models", gps.GetGpsModels)
			v1Group.POST("/search-gps", gps.SearchGPS)
			v1Group.POST("/create-gps", gps.CreateGPS)
			v1Group.PUT("/update-gps", gps.UpdateGPS)
			v1Group.GET("/gps/general", gps.GetGpsGeneralByGpsId)
		}

		// Version 2 API group
		v2Group := api.Group("/v2")
		{
			v2Group.GET("/ping", v2.Ping)
		}
	}

	// Handle 404 errors
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.NewResponse(false, http.StatusNotFound, "Not found", nil))
	})

	return r
}
