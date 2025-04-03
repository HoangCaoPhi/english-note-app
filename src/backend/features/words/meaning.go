package words

type Meaning struct {
	PartOfSpeech string   `bson:"part_of_speech"`
	Definition   string   `bson:"definition"`
	Examples     []string `bson:"examples"`
}
