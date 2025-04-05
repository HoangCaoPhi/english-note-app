package wordgroups

import (
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
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, wordGroups)
}

func (w *WordGroupController) CreateWordGroup(ctx *gin.Context) {
	var createWordGroupRequest CreateWordGroupRequest

	if err := ctx.ShouldBindJSON(&createWordGroupRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	id, err := w.WordGroupService.CreateWordGroup(ctx, createWordGroupRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"id": id})
}
