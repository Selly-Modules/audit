package audit

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Audit ...
type Audit struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Target    string             `bson:"target" json:"target"`
	TargetID  string             `bson:"targetId" json:"targetId"`
	Action    string             `bson:"action" json:"action"`
	Data      string             `bson:"data" json:"data"`
	Author    Author             `bson:"author" json:"author"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	Message   string             `bson:"message" json:"message"`
}

// Author ...
type Author struct {
	ID   string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
	Type string `bson:"type" json:"type"`
}
