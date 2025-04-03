package words

type Pronunciation struct {
	IPA    string `bson:"ipa"`
	Region string `bson:"region"`
}
