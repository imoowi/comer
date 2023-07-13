package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/apps/user/repos"
	"{{.moduleName}}/global"
	"{{.moduleName}}/middlewares/token"
	"{{.moduleName}}/utils/request"
	"{{.moduleName}}/utils/response"
	"github.com/spf13/cast"
)

var User *UserService

func RegisterUserService(s *UserService) {
	User = s
}

type UserService struct {
	UserRepo *repos.UserRepo
	RoleRepo *repos.RoleRepo
}

func newUserService(r *repos.UserRepo, r2 *repos.RoleRepo) *UserService {
	return &UserService{
		UserRepo: r,
		RoleRepo: r2,
	}
}
func init() {
	global.DigContainer.Provide(newUserService)
	global.RegisterContainerProviders(RegisterUserService)
}

func (s *UserService) PageList(c *gin.Context, req *request.PageList) (res *response.PageList, err error) {
	res, err = s.UserRepo.PageList(c, req)
	return
}

func (s *UserService) One(c *gin.Context, id uint) (model *models.User, err error) {
	model, err = s.UserRepo.One(c, id)
	return
}
func (s *UserService) OneByUsername(c *gin.Context, username string) (model *models.User, err error) {
	model, err = s.UserRepo.OneByUsername(c, username)
	return
}
func (s *UserService) Add(c *gin.Context, _model *models.UserAdd) (newId uint, err error) {
	admin, _ := s.UserRepo.One(c, c.GetUint(`uid`))
	// 角色是否存在？
	role, _ := Role.One(c, _model.RoleId)
	if role == nil || role.ID <= 0 {
		newId = 0
		err = errors.New(`角色不存在`)
		return
	}
	// 用户是否存在？
	model, _ := s.UserRepo.OneByUsername(c, _model.Username)
	if model != nil && model.ID > 0 {
		newId = 0
		err = errors.New(`用户名已经存在`)
		return
	}
	newUser := &models.User{
		Username: _model.Username,
		Passwd:   _model.Passwd,
	}
	newId, err = s.UserRepo.Add(c, newUser)
	if newId > 0 {
		// 插入用户和角色关系

		UserRole.Add(c, &models.UserRole{UserID: newId, RoleId: _model.RoleId}, admin)

		go func(c *gin.Context, _model *models.UserAdd, newId uint) {
			userlog := &models.UserLog{
				UserID:     admin.ID,
				LogType:    models.USER_LOG_TYPE_USER_ADD,
				LogContent: `管理员【` + admin.Username + `】添加了用户【` + cast.ToString(newId) + `】`,
				IP:         c.ClientIP(),
			}
			UserLog.Add(c, userlog)
		}(c, _model, newId)
	}
	return
}
func (s *UserService) Update(c *gin.Context, _model *models.User, id uint) (updated bool, err error) {
	admin, _ := s.UserRepo.One(c, c.GetUint(`uid`))
	updated, err = s.UserRepo.Update(c, _model, id)
	if updated {
		userlog := &models.UserLog{
			UserID:     admin.ID,
			LogType:    models.USER_LOG_TYPE_USER_UPDATE,
			LogContent: `管理员【` + admin.Username + `】更新了用户【` + cast.ToString(id) + `】`,
			IP:         c.ClientIP(),
		}
		UserLog.Add(c, userlog)
	}
	return
}

func (s *UserService) PatchUpdate(c *gin.Context, patchData map[string]any, id uint) (updated bool, err error) {
	updated, err = s.UserRepo.PatchUpdate(c, patchData, id)
	return
}

func (s *UserService) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	admin, _ := s.UserRepo.One(c, c.GetUint(`uid`))
	deleted, err = s.UserRepo.Delete(c, id)
	go func() {
		UserRole.DeleteByUid(c, id, admin)
		userlog := &models.UserLog{
			UserID:     admin.ID,
			LogType:    models.USER_LOG_TYPE_USER_DELETE,
			LogContent: `管理员【` + admin.Username + `】删除了用户【` + cast.ToString(id) + `】`,
			IP:         c.ClientIP(),
		}
		UserLog.Add(c, userlog)
	}()
	return
}

func (s *UserService) Login(c *gin.Context, login *models.UserLogin) (user *models.User, err error) {
	user, err = s.UserRepo.Login(c, login)
	if user != nil && user.ID > 0 {
		go func(c *gin.Context,user *models.User) {
			userlog := &models.UserLog{
				UserID:     user.ID,
				LogType:    models.USER_LOG_TYPE_USER_LOGIN,
				LogContent: `用户【` + user.Username + `】登录了系统`,
				IP:         c.ClientIP(),
			}
			UserLog.Add(c, userlog)
		}(c,user)
	}

	return
}
func (s *UserService) Logout(c *gin.Context, token string) bool {
	// err := global.Redis.SAdd(c, JwtTokenBlackListSetKey, token).Err()
	var err error
	admin, _ := s.UserRepo.One(c, c.GetUint(`uid`))
	if admin.ID > 0 {
		go func(c *gin.Context, admin *models.User) {
			userlog := &models.UserLog{
				UserID:     admin.ID,
				LogType:    models.USER_LOG_TYPE_USER_LOGOUT,
				LogContent: `用户【` + admin.Username + `】退出了系统`,
				IP:         c.ClientIP(),
			}
			UserLog.Add(c, userlog)
		}(c, admin)
	}
	return err == nil
}

func (s *UserService) ChgPwd(c *gin.Context, userChgPwd *models.UserChgPwd) (newToken string, err error) {
	admin, _ := s.UserRepo.One(c, c.GetUint(`uid`))
	ok, err := s.UserRepo.ChgPwd(c, userChgPwd)
	if err == nil && ok {
		if admin.ID > 0 {
			newToken, _ = token.GenToken(admin.Username, admin.ID)
			go func(c *gin.Context, admin *models.User) {
				userlog := &models.UserLog{
					UserID:     admin.ID,
					LogType:    models.USER_LOG_TYPE_USER_CHGPWD,
					LogContent: `用户【` + admin.Username + `】修改了密码`,
					IP:         c.ClientIP(),
				}
				UserLog.Add(c, userlog)
			}(c, admin)
		}
	}
	return
}

const JwtTokenBlackListSetKey = `lynkros-monitor:jwt:token:blacklist`

func (s *UserService) IsLogouted(c *gin.Context, token string) bool {
	// ok, _ := global.Redis.SIsMember(c, JwtTokenBlackListSetKey, token).Result()
	var ok bool
	return ok
}
func (s *UserService) IsSuper(c *gin.Context, userId uint) bool {
	userRoles, err := UserRole.AllByUserId(c, userId)
	if err != nil {
		return false
	}
	if len(userRoles) <= 0 {
		return false
	}
	for _, v := range userRoles {
		role, err := Role.One(c, v.RoleId)
		if err != nil {
			continue
		}
		if role.Level == models.ROLE_LEVEL_SUPER {
			return true
		}
	}
	return false
}
func (s *UserService) UserRoleIds(c *gin.Context, userId uint) []uint {
	userRoleIds, err := UserRole.AllIdsByUserId(c, userId)
	if err != nil {
		return make([]uint, 0)
	}
	if len(userRoleIds) > 0 {
		return userRoleIds
	}
	return make([]uint, 0)
}
func (s *UserService) UserRolesByUid(c *gin.Context, userId uint) (roles []*models.Role) {
	userRoles, err := UserRole.AllByUserId(c, userId)
	if err != nil {
		return nil
	}
	if len(userRoles) <= 0 {
		return nil
	}
	for _, v := range userRoles {
		role, err := Role.One(c, v.RoleId)
		if err != nil {
			continue
		}
		roles = append(roles, role)
	}
	return
}
