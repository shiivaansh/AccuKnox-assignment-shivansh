package entities

type DeleteNote struct {
	SID string `json:"sid" bson:"sid"`
	ID  string `json:"id" bson:"_id,omitempty"`
}
