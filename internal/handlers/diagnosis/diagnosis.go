package diagnosis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/victor-octavio/telehealth-audit-api/internal/models/diagnosis"
	"github.com/victor-octavio/telehealth-audit-api/internal/usecases"
)

type DiagnosisHandler struct {
	uc *usecases.Usecases
}

func Init(uc *usecases.Usecases, router *gin.RouterGroup) {
	dh := &DiagnosisHandler{
		uc: uc,
	}

	router.POST("/insert", dh.Add)
}

func (dh *DiagnosisHandler) Add(ctx *gin.Context) {
	var req models.DiagnosisRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "An error occurred during the diagnosis insertion request"})
		return
	}

	if err := dh.uc.Diagnosis.Add(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred during the diagnosis insertion"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": req.ID, "message": "diagnosis record created in ledger"})
}

func (dh *DiagnosisHandler) GetHistory(ctx *gin.Context) {

}

func (dh *DiagnosisHandler) GetById(ctx *gin.Context) {

}
