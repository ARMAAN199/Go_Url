package modal

import "go.mongodb.org/mongo-driver/bson/primitive"

type Url struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Url       string             `json:"url,omitempty"`
	Shortened string             `json:"shortened,omitempty"`
}
