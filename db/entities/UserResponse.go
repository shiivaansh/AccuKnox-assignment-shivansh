package entities

type UserRespnse struct {
	// ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Id      string `json:"id" bson:"id"`
	FName   string `json:"fname" bson:"fname"`
	City    string `json:"city" bson:"city"`
	Phone   int64  `json:"phone" bson:"phone"`
	Height  int    `json:"height"`
	Married bool   `json:"Married"`
}
