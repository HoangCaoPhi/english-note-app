package words

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type WordController struct {
	WordService WordService
}

func NewWordController(wordService WordService) *WordController {
	return &WordController{
		WordService: wordService,
	}
}

func (w *WordController) CreateWord(ctx *gin.Context) {
	var createWordRequest CreateWordRequest

	if err := ctx.ShouldBindJSON(&createWordRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	id, err := w.WordService.CreateWord(ctx, createWordRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"id": id})
}

func (h *WordController) GetWordsByGroupID(c *gin.Context) {
	groupIDStr := c.Query("groupId")
	if groupIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "groupId is required"})
		return
	}

	groupID, err := bson.ObjectIDFromHex(groupIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid groupId"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	words, total, err := h.WordService.GetWordsByGroupID(c.Request.Context(), groupID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": words,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}
