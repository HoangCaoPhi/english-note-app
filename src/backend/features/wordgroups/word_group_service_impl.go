package wordgroups

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type WordGroupServiceImpl struct {
	WordGroupRepositoryRead  WordGroupRepositoryRead
	WordGroupRepositoryWrite WordGroupRepositoryWrite
}

func NewWordGroupServiceImpl(
	wordGroupRepositoryRead WordGroupRepositoryRead,
	wordGroupRepositoryWrite WordGroupRepositoryWrite,
) *WordGroupServiceImpl {
	return &WordGroupServiceImpl{
		WordGroupRepositoryRead:  wordGroupRepositoryRead,
		WordGroupRepositoryWrite: wordGroupRepositoryWrite,
	}
}
func (w *WordGroupServiceImpl) GetWordGroupsByUserId(ctx context.Context) ([]WordGroup, error) {
	userId := ctx.Value("userId").(bson.Binary)

	wordGroups, err := w.WordGroupRepositoryRead.GetWordGroupsByUserID(userId)
	if err != nil {
		return nil, err
	}
	return wordGroups, nil
}

func (w *WordGroupServiceImpl) CreateWordGroup(ctx context.Context, createRequest CreateWordGroupRequest) (bson.Binary, error) {
	userId := ctx.Value("userId").(bson.Binary)

	wordGroup := NewWordGroup(userId, createRequest.Name)
	id, err := w.WordGroupRepositoryWrite.CreateWordGroup(wordGroup)
	if err != nil {
		return bson.Binary{}, err
	}
	return id, nil
}
