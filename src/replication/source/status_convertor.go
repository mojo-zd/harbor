package source

import (
	"github.com/vmware/harbor/src/replication"
	"github.com/vmware/harbor/src/replication/models"
	"github.com/vmware/harbor/src/replication/registry"
)

// StatusConvertor implement Convertor interface, convert tag with status
type StatusConvertor struct {
	registry registry.Adaptor
}

// NewRepositoryConvertor returns an instance of StatusConvertor
func NewStatusConvertor(registry registry.Adaptor) *StatusConvertor {
	return &StatusConvertor{
		registry: registry,
	}
}

// Convert projects to repositories
func (s *StatusConvertor) Convert(items []models.FilterItem) []models.FilterItem {
	result := []models.FilterItem{}
	for _, item := range items {
		if item.Kind == replication.FilterItemKindTag {
			// just put it to the result list if the item is not a repository
			item.Kind = replication.FilterItemKindImageStatus
			result = append(result, item)
			continue
		}
	}
	return result
}
