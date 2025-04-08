package words

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type WordServiceImpl struct {
	wordRepositoryRead  WordRepositoryRead
	wordRepostioryWrite WordRepostioryWrite
}

func NewWordServiceImpl(r WordRepositoryRead, w WordRepostioryWrite) *WordServiceImpl {
	return &WordServiceImpl{
		wordRepositoryRead:  r,
		wordRepostioryWrite: w,
	}
}

func (s *WordServiceImpl) CreateWord(ctx context.Context, createWordRequest CreateWordRequest) (bson.ObjectID, error) {
	userId := ctx.Value("userId").(bson.ObjectID)

	exists, err := s.wordRepositoryRead.IsWordExistsInGroup(ctx, createWordRequest.GroupID, userId, createWordRequest.Word)
	if err != nil {
		return bson.ObjectID{}, err
	}

	if exists {
		return bson.ObjectID{}, errors.New("word already exists in this group")
	}

	word := CreateWord(
		createWordRequest.GroupID,
		userId,
		createWordRequest.Word,
		createWordRequest.Language,
		createWordRequest.Pronunciations,
		createWordRequest.Meanings,
		createWordRequest.Synonyms,
		createWordRequest.Antonyms,
	)

	return s.wordRepostioryWrite.CreateWord(ctx, *word)
}

func (s *WordServiceImpl) GetWordsByGroupID(ctx context.Context, groupID bson.ObjectID, page, limit int) ([]Word, int64, error) {
	return s.wordRepositoryRead.GetWordsByGroupID(ctx, groupID, page, limit)
}
