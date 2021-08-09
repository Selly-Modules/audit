package audit

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Audit ...
type Audit struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Source    string             `bson:"source" json:"source"`
	Target    string             `bson:"target" json:"target"`
	TargetID  string             `bson:"targetId" json:"targetId"`
	Action    string             `bson:"action" json:"action"`
	Data      string             `bson:"data" json:"data"`
	Author    Author             `bson:"author" json:"author"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

// Author ...
type Author struct {
	ID   string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}
