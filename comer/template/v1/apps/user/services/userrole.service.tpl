package services

import (
	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/apps/user/repos"
	"{{.moduleName}}/global"
	"{{.moduleName}}/utils/request"
	"{{.moduleName}}/utils/response"
	"github.com/spf13/cast"
)

var UserRole *UserRoleService

func RegisterUserRoleService(s *UserRoleService) {
	UserRole = s
}

type UserRoleService struct {
	UserRoleRepo *repos.UserRoleRepo
}

func newUserRoleService(r *repos.UserRoleRepo) *UserRoleService {
	return &UserRoleService{
		UserRoleRepo: r,
	}
}

func init() {
	global.DigContainer.Provide(newUserRoleService)
	global.RegisterContainerProviders(RegisterUserRoleService)
}

func (s *UserRoleService) PageList(c *gin.Context, req *request.PageList) (res *response.PageList, err error) {
	res, err = s.UserRoleRepo.PageList(c, req)
	return
}

func (s *UserRoleService) One(c *gin.Context, id uint) (model *models.UserRole, err error) {
	model, err = s.UserRoleRepo.One(c, id)
	return
}

func (s *UserRoleService) AllByUserId(c *gin.Context, userId uint) (model []*models.UserRole, err error) {
	model, err = s.UserRoleRepo.AllByUserId(c, userId)
	return
}

func (s *UserRoleService) AllIdsByUserId(c *gin.Context, userId uint) (ids []uint, err error) {
	model, err := s.UserRoleRepo.AllByUserId(c, userId)
	if len(model) > 0 {
		for _, v := range model {
			ids = append(ids, v.ID)
		}

	}
	return
}

func (s *UserRoleService) Add(c *gin.Context, _model *models.UserRole, admin *models.User) (newId int, err error) {

	newId, err = s.UserRoleRepo.Add(c, _model)
	if newId > 0 {
		go func(c *gin.Context, _model *models.UserRole) {
			userlog := &models.UserLog{
				UserID:     admin.ID,
				LogType:    models.USER_LOG_TYPE_ROLE_ADD,
				LogContent: `管理员【` + admin.Username + `】添加了用户角色【` + cast.ToString(_model.RoleId) + `】`,
				IP:         c.ClientIP(),
			}
			UserLog.Add(c, userlog)
		}(c, _model)
	}
	return
}

func (s *UserRoleService) Update(c *gin.Context, _model *models.UserRole, id uint, admin *models.User) (updated bool, err error) {
	updated, err = s.UserRoleRepo.Update(c, _model, id)
	if updated {
		go func() {
			userlog := &models.UserLog{
				UserID:     admin.ID,
				LogType:    models.USER_LOG_TYPE_ROLE_ADD,
				LogContent: `管理员【` + admin.Username + `】修改了用户角色关联关系【` + cast.ToString(id) + `】`,
				IP:         c.ClientIP(),
			}
			UserLog.Add(c, userlog)
		}()
	}
	return
}
func (s *UserRoleService) Delete(c *gin.Context, id uint, admin *models.User) (deleted bool, err error) {
	deleted, err = s.UserRoleRepo.Delete(c, id)
	if deleted {
		go func() {
			userlog := &models.UserLog{
				UserID:     admin.ID,
				LogType:    models.USER_LOG_TYPE_ROLE_ADD,
				LogContent: `管理员【` + admin.Username + `】删除了用户角色关联关系【` + cast.ToString(id) + `】`,
				IP:         c.ClientIP(),
			}
			UserLog.Add(c, userlog)
		}()
	}
	return
}

func (s *UserRoleService) DeleteByUid(c *gin.Context, userId uint, admin *models.User) (deleted bool, err error) {
	deleted, err = s.UserRoleRepo.DeleteByUid(c, userId)
	if deleted {
		go func() {
			userlog := &models.UserLog{
				UserID:     admin.ID,
				LogType:    models.USER_LOG_TYPE_ROLE_ADD,
				LogContent: `管理员【` + admin.Username + `】删除了用户【` + cast.ToString(userId) + `】的所有角色关联关系`,
				IP:         c.ClientIP(),
			}
			UserLog.Add(c, userlog)
		}()
	}
	return
}
