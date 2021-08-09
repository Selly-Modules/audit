package audit

import (
	"github.com/Selly-Modules/mongodb"
)

func (s Service) indexDB() {
	// Get list targets
	var targets = make([]string, 0)
	if s.Source == SourceSelly {
		targets = SellyTargets
	}

	// Index key
	commonIndex := mongodb.NewIndexKey("source", "target", "targetId")

	// Index all allowed sources
	for _, target := range targets {
		mongodb.CreateIndex(getColName(target), commonIndex)
	}
}
