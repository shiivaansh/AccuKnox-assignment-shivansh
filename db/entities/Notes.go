package entities

import "gopkg.in/mgo.v2/bson"

type Notes struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Notes []Note        `json:"notes"`
}
