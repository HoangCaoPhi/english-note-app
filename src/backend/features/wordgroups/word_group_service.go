package wordgroups

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type WordGroupService interface {
	GetWordGroupsByUserId(ctx context.Context) ([]WordGroup, error)
	CreateWordGroup(ctx context.Context, createRequest CreateWordGroupRequest) (bson.Binary, error)
}

var wordGroupService WordGroupService

func InitWordGroupService(service WordGroupService) {
	wordGroupService = service
}

func NewWordGroupService() WordGroupService {
	return wordGroupService
}
