package entities

type Note struct {
	// ID  bson.ObjectId `json:"id" bson:"_id,omitempty"`
	SID string `json:"sid" bson:"sid"`
	Nte string `json:"note" bson:"note"`
}
