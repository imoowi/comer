package repos

import (
	"errors"

	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/global"
	"{{.moduleName}}/utils/request"
	"{{.moduleName}}/utils/response"
)

type UserLogRepo struct {
	db *global.MysqlODM
}

func newUserLogRepo() *UserLogRepo {
	return &UserLogRepo{
		db: global.MysqlDb,
	}
}

func init() {
	global.DigContainer.Provide(newUserLogRepo)
}
func (r *UserLogRepo) PageList(c *gin.Context, req *request.PageList) (res *response.PageList, err error) {
	db := r.db.Client
	var userlogs []*models.UserLog

	if req.SearchKey != `` {
		db = db.Where(`name LIKE ?`, `%`+req.SearchKey+`%`)
	}
	offset := (req.Page - 1) * req.PageSize
	db = db.Offset(int(offset)).Limit(int(req.PageSize))
	// db=db.Order(`name desc`)
	err = db.Find(&userlogs).Error

	var count int64
	db.Offset(-1).Limit(-1).Count(&count)

	res = &response.PageList{
		List:  userlogs,
		Pages: response.MakePages(count, req.Page, req.PageSize),
	}
	return
}

func (r *UserLogRepo) One(c *gin.Context, id uint) (userlog *models.UserLog, err error) {
	db := r.db.Client
	err = db.Where(`id=?`, id).First(&userlog).Error
	return
}

func (r *UserLogRepo) Add(c *gin.Context, model *models.UserLog) (newId uint, err error) {
	db := r.db.Client
	db = db.Create(&model)
	err = db.Error

	newId = model.ID
	return
}

func (r *UserLogRepo) Update(c *gin.Context, model *models.UserLog, id uint) (updated bool, err error) {
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

func (r *UserLogRepo) Delete(c *gin.Context, id uint) (deleted bool, err error) {
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
