package diagnosis

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-octavio/telehealth-audit-api/internal/usecases"
)

type DiagnosisHandler struct {
	uc *usecases.Usecases
}

func Init(uc *usecases.Usecases, router *gin.RouterGroup) {
	dh := &DiagnosisHandler{
		uc: uc,
	}

	router.GET("/history", dh.GetHistory)
}

func (dh *DiagnosisHandler) Add(ctx *gin.Context) {

}

func (dh *DiagnosisHandler) GetHistory(ctx *gin.Context) {

}

func (dh *DiagnosisHandler) GetById(ctx *gin.Context) {

}
