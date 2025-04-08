package words

type Pronunciation struct {
	IPA    string `bson:"ipa"`
	Region string `bson:"region"`
	Audio  string `bson:"audio" json:"audio"`
}
