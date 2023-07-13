package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/apps/user/repos"
	"{{.moduleName}}/global"
	"{{.moduleName}}/utils/request"
	"{{.moduleName}}/utils/response"
	"github.com/spf13/cast"
)

var Role *RoleService

func RegisterRoleService(s *RoleService) {
	Role = s
}

type RoleService struct {
	RoleRepo *repos.RoleRepo
}

func newRoleService(r *repos.RoleRepo) *RoleService {
	return &RoleService{
		RoleRepo: r,
	}
}

func init() {
	global.DigContainer.Provide(newRoleService)
	global.RegisterContainerProviders(RegisterRoleService)
}

func (s *RoleService) PageList(c *gin.Context, req *request.PageList) (res *response.PageList, err error) {
	res, err = s.RoleRepo.PageList(c, req)
	return
}

func (s *RoleService) One(c *gin.Context, id uint) (model *models.Role, err error) {
	model, err = s.RoleRepo.One(c, id)
	return
}
func (s *RoleService) OneByName(c *gin.Context, name string) (model *models.Role, err error) {
	model, err = s.RoleRepo.OneByName(c, name)
	return
}
func (s *RoleService) Add(c *gin.Context, _role *models.Role) (newId uint, err error) {
	model, _ := s.RoleRepo.OneByName(c, _role.Name)
	if model != nil && model.ID > 0 {
		newId = 0
		err = errors.New(`name existed`)
		return
	}
	newId, err = s.RoleRepo.Add(c, _role)
	admin, _ := User.One(c, c.GetUint(`uid`))
	if newId > 0 {
		go func(c *gin.Context, newId uint) {
			userlog := &models.UserLog{
				UserID:     admin.ID,
				LogType:    models.USER_LOG_TYPE_ROLE_ADD,
				LogContent: `管理员【` + admin.Username + `】添加了角色【` + cast.ToString(newId) + `】`,
				IP:         c.ClientIP(),
			}
			UserLog.Add(c, userlog)
		}(c, newId)
	}
	return
}
func (s *RoleService) Update(c *gin.Context, _role *models.Role, id uint) (updated bool, err error) {
	updated, err = s.RoleRepo.Update(c, _role, id)
	if updated {
		admin, _ := User.One(c, c.GetUint(`uid`))
		go func() {
			userlog := &models.UserLog{
				UserID:     admin.ID,
				LogType:    models.USER_LOG_TYPE_ROLE_ADD,
				LogContent: `管理员【` + admin.Username + `】修改了角色【` + cast.ToString(id) + `】`,
				IP:         c.ClientIP(),
			}
			UserLog.Add(c, userlog)
		}()
	}
	return
}

func (s *RoleService) PatchUpdate(c *gin.Context, patchData map[string]any, id uint) (updated bool, err error) {
	updated, err = s.RoleRepo.PatchUpdate(c, patchData, id)
	return
}

func (s *RoleService) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	deleted, err = s.RoleRepo.Delete(c, id)
	if deleted {
		admin, _ := User.One(c, c.GetUint(`uid`))
		go func() {
			userlog := &models.UserLog{
				UserID:     admin.ID,
				LogType:    models.USER_LOG_TYPE_ROLE_ADD,
				LogContent: `管理员【` + admin.Username + `】删除了角色【` + cast.ToString(id) + `】`,
				IP:         c.ClientIP(),
			}
			UserLog.Add(c, userlog)
		}()
	}
	return
}
