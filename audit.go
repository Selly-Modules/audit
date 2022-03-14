package audit

import (
	"errors"
	"fmt"

	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// Config ...
type Config struct {
	// Targets: staff, article, ...
	Targets []string
	// MongoDB config, for save documents
	MongoDB mongodb.Config
}

// Service ...
type Service struct {
	Config
	DB *mongo.Database
}

var s *Service

// NewInstance ...
func NewInstance(config Config) error {
	if len(config.Targets) == 0 || config.MongoDB.Host == "" {
		return errors.New("please provide all necessary information: targets, mongodb")
	}

	// Connect MongoDB
	db, err := mongodb.Connect(config.MongoDB)
	if err != nil {
		fmt.Println("Cannot init module AUDIT", err)
		return err
	}

	s = &Service{
		Config: config,
		DB:     db,
	}

	// index mongo
	s.indexDB()

	return nil
}

// GetInstance ...
func GetInstance() *Service {
	return s
}

// getColName ...
func getColName(target string) string {
	return target
}
