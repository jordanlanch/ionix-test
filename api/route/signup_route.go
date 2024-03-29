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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
