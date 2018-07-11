package dao

import (
	"strings"

	"github.com/jinzhu/copier"
	"github.com/vmware/harbor/src/common/models"
)

// UpdateImage insert or update image
func UpdateImage(image *models.Image) (err error) {
	target := &models.Image{}
	copier.Copy(target, image)
	GetImageByRepoAndTag(image)
	if image.ID > 0 {
		target.ID = image.ID
		_, err = GetOrmer().Update(target)
		return
	}
	_, err = GetOrmer().Insert(target)
	return
}

// GetImageByRepoAndTag get image with repository name and tag
func GetImageByRepoAndTag(image *models.Image) (err error) {
	err = GetOrmer().Read(image, "repository_name", "tag")
	return
}

// ListImages list images by condition eg: map[string]interface{}{"repository_name":"library/registry-ui"}
func ListImages(condition map[string]interface{}) (images []*models.Image, err error) {
	query := GetOrmer().QueryTable(&models.Image{})
	for k, v := range condition {
		// not equal handler
		if strings.Contains(k, "__not") {
			query = query.Exclude(strings.Replace(k, "__not", "", -1), v)
			continue
		}
		query = query.Filter(k, v)
	}
	_, err = query.All(&images)
	return
}
