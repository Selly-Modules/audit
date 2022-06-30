package audit

import (
	"context"
	"sync"

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
func (s Service) All(query AllQuery) (result []Audit, total int64) {
	var (
		ctx     = context.Background()
		colName = getColName(query.Target)
		skip    = query.Page * query.Limit
		wg      sync.WaitGroup
	)
	cond := bson.D{
		{"target", query.Target},
		{"targetId", query.TargetID},
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		opts := options.Find().SetLimit(query.Limit).SetSkip(skip).SetSort(bson.M{"_id": -1})
		if query.Sort != nil {
			opts.SetSort(query.Sort)
		}

		// Find db
		cursor, err := s.DB.Collection(colName).Find(ctx, cond, opts)
		if err != nil {
			return
		}
		defer cursor.Close(ctx)
		cursor.All(ctx, &result)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		total, _ = s.DB.Collection(colName).CountDocuments(ctx, cond)
	}()
	wg.Wait()
	return result, total
}
