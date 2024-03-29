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

func NewVaccinationRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	dr := repository.NewVaccinationRepository(db)
	dc := controller.NewVaccinationController(usecase.NewVaccinationUsecase(dr, timeout))

	group.POST("/vaccination", dc.Create)
	group.PUT("/vaccination/:id", dc.Update)
	group.GET("/vaccination/:id", dc.Get)
	group.GET("/vaccination", dc.FetchAll)
	group.DELETE("/vaccination/:id", dc.Delete)
}
