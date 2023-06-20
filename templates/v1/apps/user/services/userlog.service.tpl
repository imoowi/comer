package services

import (
	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/apps/user/repos"
	"{{.moduleName}}/global"
	"{{.moduleName}}/utils/request"
	"{{.moduleName}}/utils/response"
)

var UserLog *UserLogService

func RegisterUserLogService(s *UserLogService) {
	UserLog = s
}

type UserLogService struct {
	UserLogRepo *repos.UserLogRepo
}

func newUserLogService(r *repos.UserLogRepo) *UserLogService {
	return &UserLogService{
		UserLogRepo: r,
	}
}

func init() {
	global.DigContainer.Provide(newUserLogService)
	global.RegisterContainerProviders(RegisterUserLogService)
}

func (s *UserLogService) PageList(c *gin.Context, req *request.PageList) (res *response.PageList, err error) {
	res, err = s.UserLogRepo.PageList(c, req)
	return
}

func (s *UserLogService) One(c *gin.Context, id uint) (model *models.UserLog, err error) {
	model, err = s.UserLogRepo.One(c, id)
	return
}
func (s *UserLogService) Add(c *gin.Context, dpt *models.UserLog) (newId uint, err error) {
	newId, err = s.UserLogRepo.Add(c, dpt)
	return
}
func (s *UserLogService) Update(c *gin.Context, dpt *models.UserLog, id uint) (updated bool, err error) {
	updated, err = s.UserLogRepo.Update(c, dpt, id)
	return
}
func (s *UserLogService) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	deleted, err = s.UserLogRepo.Delete(c, id)
	return
}
