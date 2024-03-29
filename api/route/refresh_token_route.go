package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanlanch/ionix-test/api/controller"
	"github.com/jordanlanch/ionix-test/bootstrap"
	"github.com/jordanlanch/ionix-test/repository"
	"github.com/jordanlanch/ionix-test/usecase"
	"gorm.io/gorm"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
