package audit

import (
	"context"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
)

// CreatePayload ...
type CreatePayload struct {
	Target   string
	TargetID string
	Action   string
	Data     string
	Author   CreatePayloadAuthor
}

// CreatePayloadAuthor ...
type CreatePayloadAuthor struct {
	ID   string
	Name string
}

// Create ...
func (s Service) Create(payload CreatePayload) {
	ctx := context.Background()

	// Get document
	doc := Audit{
		ID:       mongodb.NewObjectID(),
		Target:   payload.Target,
		TargetID: payload.TargetID,
		Action:   payload.Action,
		Data:     payload.Data,
		Author: Author{
			ID:   payload.Author.ID,
			Name: payload.Author.Name,
		},
		CreatedAt: now(),
	}

	// Insert to db
	colName := getColName(payload.Target)
	if _, err := s.DB.Collection(colName).InsertOne(ctx, doc); err != nil {
		logger.Error("audit - Create", logger.LogData{
			"payload": payload,
			"error":   err.Error(),
		})
	}
}
