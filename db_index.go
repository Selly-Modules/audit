package audit

import (
	"github.com/Selly-Modules/mongodb"
)

func (s Service) indexDB() {
	// Index key
	commonIndex := mongodb.NewIndexKey("source", "target", "targetId")

	// Index all allowed sources
	for _, target := range s.Targets {
		mongodb.CreateIndex(getColName(target), commonIndex)
	}
}
