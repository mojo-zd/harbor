package source

import (
	"fmt"

	"github.com/vmware/harbor/src/common"
	"github.com/vmware/harbor/src/common/dao"
	"github.com/vmware/harbor/src/common/utils/log"
	"github.com/vmware/harbor/src/replication"
	"github.com/vmware/harbor/src/replication/models"
	"github.com/vmware/harbor/src/replication/registry"
)

// RepositoryFilter implement Filter interface to filter repository
type StatusFilter struct {
	pattern   string
	convertor Convertor
}

// StatusFilter returns an instance of StatusFilter
func NewStatusFilter(pattern string, registry registry.Adaptor) *StatusFilter {
	return &StatusFilter{
		pattern:   pattern,
		convertor: NewStatusConvertor(registry),
	}
}

// Init ...
func (s *StatusFilter) Init() error {
	return nil
}

// GetConvertor ...
func (s *StatusFilter) GetConvertor() Convertor {
	return s.convertor
}

// DoFilter filters tag of the image
func (s *StatusFilter) DoFilter(items []models.FilterItem) []models.FilterItem {
	candidates := []string{}
	for _, item := range items {
		candidates = append(candidates, item.Value)
	}
	log.Debugf("tag filter candidates: %v", candidates)
	result := []models.FilterItem{}
	if s.pattern != common.ImagePending && s.pattern != common.ImageDeveloping && s.pattern != common.ImageFinished && s.pattern != common.ImageFailed {
		log.Warning("pattern not match image status")
		return result
	}
	images, _ := dao.ListImages(map[string]interface{}{})
	for _, item := range items {
		if item.Kind != replication.FilterItemKindImageStatus {
			log.Warningf("unsupported type %s for tag filter, dropped", item.Kind)
			continue
		}

		if len(s.pattern) == 0 {
			log.Debugf("pattern is null, add %s to the tag filter result list", item.Value)
			result = append(result, item)
			continue
		}

		appendable := false
		exist := false
		for _, image := range images {
			if fmt.Sprintf("%s:%s", image.RepositoryName, image.Tag) == item.Value {
				exist = true
				if s.pattern == image.Status {
					appendable = true
				}
				break
			}
		}

		if !exist && s.pattern == common.ImagePending {
			appendable = true
		}

		if appendable {
			result = append(result, item)
		}
	}

	return result
}
