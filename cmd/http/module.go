package http

import (
	"mnc-rest-api/internal/config"

	cRegister "mnc-rest-api/internal/controller/register"

	cLogin "mnc-rest-api/internal/controller/login"
	cPayment "mnc-rest-api/internal/controller/payment"
	cTopup "mnc-rest-api/internal/controller/topup"
	cTransaction "mnc-rest-api/internal/controller/transaction"
	cTransfer "mnc-rest-api/internal/controller/transfer"
	cUser "mnc-rest-api/internal/controller/user"
	"mnc-rest-api/internal/domain"

	uLogin "mnc-rest-api/internal/usecase/login"
	uPayment "mnc-rest-api/internal/usecase/payment"
	uRegister "mnc-rest-api/internal/usecase/register"
	uTopup "mnc-rest-api/internal/usecase/topup"
	uTransaction "mnc-rest-api/internal/usecase/transaction"
	uTransfer "mnc-rest-api/internal/usecase/transfer"
	uUser "mnc-rest-api/internal/usecase/user"

	rRegister "mnc-rest-api/internal/repository/register"
	rTopup "mnc-rest-api/internal/repository/topup"
	rTransaction "mnc-rest-api/internal/repository/transaction"
	rUser "mnc-rest-api/internal/repository/user"

	middleware "mnc-rest-api/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	repoTransaction domain.TransactionRepository
	repoRegister    domain.RegisterRepository
	repoUser        domain.UserRepository
	repoTopup       domain.TopupRepository

	usecaseTransaction domain.TransactionUsecase
	usecaseLogin       domain.LoginUsecase
	usecaseRegister    domain.RegisterUsecase
	usecaseTopup       domain.TopupUsecase
	usecasePayment     domain.PaymentUsecase
	usecaseTransfer    domain.TransferUsecase
	usecaseUser        domain.UserUsecase
)

func InitializeRepositories(db *gorm.DB) {

	repoTransaction = rTransaction.New(db)
	repoRegister = rRegister.New(db)
	repoUser = rUser.New(db)
	repoTopup = rTopup.New(db)

}

func InitializeUsecases(config config.Config) {

	usecaseTransaction = uTransaction.New(repoTransaction)
	usecaseRegister = uRegister.New(repoRegister, repoUser)
	usecaseLogin = uLogin.New("xxxsecret", repoUser)
	usecaseTopup = uTopup.New(repoTopup, repoUser, repoTransaction)
	usecasePayment = uPayment.New(repoUser, repoTransaction)
	usecaseTransfer = uTransfer.New(repoUser, repoTransaction)
	usecaseUser = uUser.New(repoUser)
}

func InitializeControllers(router *gin.Engine, config config.Config) {

	publicApiV1_0 := router.Group("public/api/v1.0")
	{
		registerPublicApiV1_0(publicApiV1_0, config)
	}

}

func registerPublicApiV1_0(v1_0 *gin.RouterGroup, config config.Config) {
	controllerRegister := cRegister.New(usecaseRegister)
	controllerLogin := cLogin.New(usecaseLogin)
	controllerTopup := cTopup.New(usecaseTopup)
	controllerPayment := cPayment.New(usecasePayment)
	controllerTransfer := cTransfer.New(usecaseTransfer)
	controllerTransaction := cTransaction.New(usecaseTransaction)
	controllerUser := cUser.New(usecaseUser)

	authMiddleware := middleware.AuthMiddleware()

	BaseRoutes := v1_0.Group("")
	{
		BaseRoutes.POST("/register", controllerRegister.Register)
		BaseRoutes.POST("/login", controllerLogin.Login)
		BaseRoutes.POST("/topup", authMiddleware, controllerTopup.TopUp)
		BaseRoutes.POST("/pay", authMiddleware, controllerPayment.Pay)
		BaseRoutes.POST("/transfer", authMiddleware, controllerTransfer.Transfer)
		BaseRoutes.GET("/transactions", authMiddleware, controllerTransaction.GetTransactions)
		BaseRoutes.PUT("/profile", authMiddleware, controllerUser.UpdateProfile)
	}
}
