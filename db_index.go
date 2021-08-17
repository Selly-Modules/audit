package audit

import (
	"github.com/Selly-Modules/mongodb"
)

func (s Service) indexDB() {
	// Index key
	commonIndex := mongodb.NewIndexKey("target", "targetId")

	// Index all targets
	for _, target := range s.Targets {
		mongodb.CreateIndex(getColName(target), commonIndex)
	}
}
