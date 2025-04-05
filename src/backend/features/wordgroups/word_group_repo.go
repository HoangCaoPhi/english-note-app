package wordgroups

import "go.mongodb.org/mongo-driver/v2/bson"

type (
	WordGroupRepositoryRead interface {
		GetWordGroupsByUserID(userID bson.ObjectID) ([]WordGroup, error)
	}

	WordGroupRepositoryWrite interface {
		CreateWordGroup(wordGroup *WordGroup) (bson.ObjectID, error)
		UpdateWordGroup(wordGroup WordGroup) (bson.ObjectID, error)
		DeleteWordGroup(id bson.ObjectID) (bson.ObjectID, error)
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
