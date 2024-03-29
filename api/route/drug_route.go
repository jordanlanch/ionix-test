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

func NewDrugRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	dr := repository.NewDrugRepository(db)
	dc := controller.NewDrugController(usecase.NewDrugUsecase(dr, timeout))

	group.POST("/drugs", dc.Create)
	group.PUT("/drugs/:id", dc.Update)
	group.GET("/drugs/:id", dc.Get)
	group.GET("/drugs", dc.FetchAll)
	group.DELETE("/drugs/:id", dc.Delete)
}
