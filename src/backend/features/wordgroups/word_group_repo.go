package wordgroups

import "go.mongodb.org/mongo-driver/v2/bson"

type (
	WordGroupRepositoryRead interface {
		GetWordGroupsByUserID(userID bson.Binary) ([]WordGroup, error)
	}

	WordGroupRepositoryWrite interface {
		CreateWordGroup(wordGroup *WordGroup) (bson.Binary, error)
		UpdateWordGroup(wordGroup WordGroup) (bson.Binary, error)
		DeleteWordGroup(id bson.Binary) (bson.Binary, error)
	}
)

var (
	wordGroupRepositoryRead  WordGroupRepositoryRead
	wordGroupRepositoryWrite WordGroupRepositoryWrite
)

func InitWordGroupRepositoryRead(read WordGroupRepositoryRead) {
	wordGroupRepositoryRead = read
}

func InitWordGroupRepositoryWrite(write WordGroupRepositoryWrite) {
	wordGroupRepositoryWrite = write
}

func NewReadRepositoryRead() WordGroupRepositoryRead {
	return wordGroupRepositoryRead
}
func NewReadRepositoryWrite() WordGroupRepositoryWrite {
	return wordGroupRepositoryWrite
}
