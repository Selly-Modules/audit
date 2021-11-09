package audit

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AllQuery ...
type AllQuery struct {
	Target   string
	TargetID string
	Page     int64
	Limit    int64
	Sort     interface{}
}

// All ...
func (s Service) All(query AllQuery) []Audit {
	var (
		ctx     = context.Background()
		colName = getColName(query.Target)
		skip    = query.Page * query.Limit
		result  = make([]Audit, 0)
	)
	opts := options.Find().SetLimit(query.Limit).SetSkip(skip).SetSort(bson.M{"_id": -1})
	if query.Sort != nil {
		opts.SetSort(query.Sort)
	}

	// Find db
	cursor, err := s.DB.Collection(colName).Find(ctx, bson.D{
		{"target", query.Target},
		{"targetId", query.TargetID},
	}, opts)
	if err != nil {
		return result
	}
	defer cursor.Close(ctx)
	cursor.All(ctx, &result)
	return result
}
