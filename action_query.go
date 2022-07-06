package audit

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AllQuery ...
type AllQuery struct {
	Target   string
	TargetID string

	// Additional filter
	Author         string
	CreateTimeFrom time.Time
	CreateTimeTo   time.Time

	// Pagination
	Page  int64
	Limit int64
	Sort  interface{}
}

// All ...
func (s Service) All(query AllQuery) (result []Audit, total int64) {
	var (
		ctx     = context.Background()
		colName = getColName(query.Target)
		skip    = query.Page * query.Limit
		wg      sync.WaitGroup
	)
	cond := s.getQueryCondition(query)
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

func (s Service) getQueryCondition(query AllQuery) bson.D {
	cond := bson.D{
		{"target", query.Target},
		{"targetId", query.TargetID},
	}
	if query.Author != "" {
		cond = append(cond, bson.E{
			Key:   "author.id",
			Value: query.Author,
		})
	}
	if !query.CreateTimeFrom.IsZero() || !query.CreateTimeTo.IsZero() {
		v := bson.M{}
		if !query.CreateTimeFrom.IsZero() {
			v["$gte"] = query.CreateTimeFrom
		}
		if !query.CreateTimeTo.IsZero() {
			v["$lt"] = query.CreateTimeTo
		}
		cond = append(cond, bson.E{
			Key:   "createdAt",
			Value: v,
		})
	}
	return cond
}
