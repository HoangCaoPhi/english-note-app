package words

import (
	"hoangcaophi/english-note-app/src/backend/pkg/response"
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
		response.BadRequest(ctx, "Invalid request payload")
		return
	}

	id, err := w.WordService.CreateWord(ctx, createWordRequest)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}

	response.Created(ctx, gin.H{"id": id})
}

func (h *WordController) GetWordsByGroupID(c *gin.Context) {
	groupIDStr := c.Query("groupId")
	if groupIDStr == "" {
		response.BadRequest(c, "groupId is required")
		return
	}

	groupID, err := bson.ObjectIDFromHex(groupIDStr)
	if err != nil {
		response.BadRequest(c, "invalid groupId")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	words, total, err := h.WordService.GetWordsByGroupID(c.Request.Context(), groupID, page, limit)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	wordResponses := make([]WordResponse, len(words))
	for i, word := range words {
		wordResponses[i] = word.ToResponse()
	}

	response.Success(c, WordListResponse{
		Data: wordResponses,
		Pagination: PaginationResponse{
			Page:  page,
			Limit: limit,
			Total: int(total),
		},
	})
}
