package wordgroups

import (
	"hoangcaophi/english-note-app/src/backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type WordGroupController struct {
	WordGroupService WordGroupService
}

func NewWordGroupController(wordGroupService WordGroupService) *WordGroupController {
	return &WordGroupController{
		WordGroupService: wordGroupService,
	}
}

func (w *WordGroupController) GetWordGroupsByUserId(ctx *gin.Context) {
	wordGroups, err := w.WordGroupService.GetWordGroupsByUserId(ctx)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}
	response.Success(ctx, wordGroups)
}

func (w *WordGroupController) CreateWordGroup(ctx *gin.Context) {
	var createWordGroupRequest CreateWordGroupRequest

	if err := ctx.ShouldBindJSON(&createWordGroupRequest); err != nil {
		response.BadRequest(ctx, "Invalid request payload")
		return
	}

	id, err := w.WordGroupService.CreateWordGroup(ctx, createWordGroupRequest)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}

	response.Created(ctx, gin.H{"id": id})
}
