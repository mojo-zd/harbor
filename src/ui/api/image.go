package api

import (
	"strings"

	"github.com/vmware/harbor/src/common"
	"github.com/vmware/harbor/src/common/dao"
	"github.com/vmware/harbor/src/common/dao/project"
	"github.com/vmware/harbor/src/common/models"
)

// ImageAPI handlers request for images management
type ImageAPI struct {
	BaseController
}

// fetch current user
func (*ImageAPI) Prepare() {

}

func (ia *ImageAPI) validate(image *models.Image) {
	if image.Status != common.ImagePending &&
		image.Status != common.ImageDeveloping &&
		image.Status != common.ImageFinished &&
		image.Status != common.ImageFailed {
		ia.HandleBadRequest("invalid image status")
		return
	}
	if err := image.Valid(); err != nil {
		ia.HandleBadRequest(err.Error())
		return
	}
	username, password, _ := ia.Ctx.Request.BasicAuth()
	user, err := dao.LoginByDb(models.AuthModel{Principal: username, Password: password})
	if err != nil || user == nil {
		ia.HandleUnauthorized()
		return
	}

	info := strings.Split(image.RepositoryName, "/")
	if len(info) != 2 {
		ia.HandleInternalServerError("repository name invalid")
		return
	}

	projectEntity, err := dao.GetProjectByName(info[0])
	if err != nil {
		ia.HandleNotFound("project not found")
		return
	}

	members, _ := project.GetProjectMember(models.Member{ProjectID: projectEntity.ProjectID})
	var currentMember *models.Member
	for _, member := range members {
		if member.EntityID == user.UserID {
			currentMember = member
			break
		}
	}

	if currentMember == nil || currentMember.Role != models.PROJECTADMIN {
		ia.HandleForbidden(username)
		return
	}

	return
}

// change the image status
func (ia *ImageAPI) Put() {
	image := &models.Image{}
	ia.DecodeJSONReqAndValidate(image)
	ia.validate(image)
	if err := dao.UpdateImage(image); err != nil {
		ia.HandleInternalServerError("modify image failed")
	}
}
