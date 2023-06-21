package repos

import (
	"errors"

	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/components"
	"{{.moduleName}}/global"
	"{{.moduleName}}/utils/request"
	"{{.moduleName}}/utils/response"
)

type UserRoleRepo struct {
	db *components.MysqlODM
}

func newUserRoleRepo() *UserRoleRepo {
	return &UserRoleRepo{
		db: global.MysqlDb,
	}
}
func init() {
	global.DigContainer.Provide(newUserRoleRepo)
}

func (r *UserRoleRepo) PageList(c *gin.Context, req *request.PageList) (res *response.PageList, err error) {
	db := r.db.Client
	var userroles []*models.UserRole

	if req.SearchKey != `` {
		db = db.Where(`name LIKE ?`, `%`+req.SearchKey+`%`)
	}
	offset := (req.Page - 1) * req.PageSize
	db = db.Offset(int(offset)).Limit(int(req.PageSize))
	// db=db.Order(`name desc`)
	err = db.Find(&userroles).Error

	var count int64
	db.Offset(-1).Limit(-1).Count(&count)

	res = &response.PageList{
		List:  userroles,
		Pages: response.MakePages(count, req.Page, req.PageSize),
	}
	return
}

func (r *UserRoleRepo) One(c *gin.Context, id uint) (userrole *models.UserRole, err error) {
	db := r.db.Client
	err = db.Where(`id=?`, id).First(&userrole).Error
	return
}
func (r *UserRoleRepo) OneByUid(c *gin.Context, userId uint) (userrole *models.UserRole, err error) {
	db := r.db.Client
	err = db.Where(`user_id=?`, userId).First(&userrole).Error
	return
}

func (r *UserRoleRepo) AllByUserId(c *gin.Context, userId uint) (userrole []*models.UserRole, err error) {
	db := r.db.Client
	err = db.Where(`user_id=?`, userId).Find(&userrole).Error
	return
}

func (r *UserRoleRepo) Add(c *gin.Context, model *models.UserRole) (newId int, err error) {
	db := r.db.Client
	db = db.Create(&model)
	err = db.Error

	newId = int(model.ID)
	return
}

func (r *UserRoleRepo) Update(c *gin.Context, model *models.UserRole, id uint) (updated bool, err error) {
	if id == 0 {
		updated = false
		err = errors.New(`pls input id`)
		return
	}
	model.ID = uint(id)
	db := r.db.Client
	err = db.Omit(`created_at`).Save(&model).Error
	if err == nil {
		updated = true
	}
	return
}

func (r *UserRoleRepo) Delete(c *gin.Context, id uint) (deleted bool, err error) {
	if id == 0 {
		deleted = false
		err = errors.New(`pls input id`)
		return
	}
	db := r.db.Client
	model, err := r.One(c, id)
	if err != nil {
		return
	}
	if model.ID == 0 {
		err = errors.New(`obj not existe`)
		return
	}
	err = db.Delete(&model).Error
	if err == nil {
		deleted = true
	}
	return
}

func (r *UserRoleRepo) DeleteByUid(c *gin.Context, userId uint) (deleted bool, err error) {
	if userId == 0 {
		deleted = false
		err = errors.New(`pls input userId`)
		return
	}
	db := r.db.Client
	err = db.Where(`user_id=?`, userId).Delete(&models.UserRole{}).Error
	if err == nil {
		deleted = true
	}

	return
}