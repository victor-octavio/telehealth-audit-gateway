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
	router.GET("/history/:id", dh.GetHistory)
	router.GET("/record/:id", dh.GetById)
}

func (dh *DiagnosisHandler) Add(ctx *gin.Context) {
	var req models.DiagnosisRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "An error occurred during the diagnosis insertion request"})
		return
	}
	if err := dh.uc.Diagnosis.Add(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred during the diagnosis insertion", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": req.ID, "message": "diagnosis record created in ledger"})
}

func (dh *DiagnosisHandler) GetHistory(ctx *gin.Context) {
	record, err := dh.uc.Diagnosis.GetHistory(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "An error occurred during the history fetch", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, record)
}

func (dh *DiagnosisHandler) GetById(ctx *gin.Context) {
	record, err := dh.uc.Diagnosis.GetById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "An error occurred during the history fetch", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, record)
}
