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

type RoleRepo struct {
	db *components.MysqlODM
}

func newRoleRepo() *RoleRepo {
	return &RoleRepo{
		db: global.MysqlDb,
	}
}
func init() {
	global.DigContainer.Provide(newRoleRepo)
}

func (r *RoleRepo) PageList(c *gin.Context, req *request.PageList) (res *response.PageList, err error) {
	db := r.db.Client
	var roles []*models.Role

	if req.SearchKey != `` {
		db = db.Where(`name LIKE ?`, `%`+req.SearchKey+`%`)
	}
	offset := (req.Page - 1) * req.PageSize
	db = db.Offset(int(offset)).Limit(int(req.PageSize))
	// db=db.Order(`name desc`)
	err = db.Find(&roles).Error

	var count int64
	db.Offset(-1).Limit(-1).Count(&count)

	res = &response.PageList{
		List:  roles,
		Pages: response.MakePages(count, req.Page, req.PageSize),
	}
	return
}

func (r *RoleRepo) One(c *gin.Context, id uint) (role *models.Role, err error) {
	db := r.db.Client
	err = db.Where(`id=?`, id).First(&role).Error
	return
}

func (r *RoleRepo) OneByName(c *gin.Context, name string) (user *models.Role, err error) {
	db := r.db.Client
	err = db.Where(`name=?`, name).First(&user).Error
	return
}
func (r *RoleRepo) Add(c *gin.Context, model *models.Role) (newId uint, err error) {
	db := r.db.Client
	db = db.Create(&model)
	err = db.Error

	newId = model.ID
	return
}

func (r *RoleRepo) Update(c *gin.Context, model *models.Role, id uint) (updated bool, err error) {
	if id == 0 {
		updated = false
		err = errors.New(`pls input id`)
		return
	}
	model.ID = id
	db := r.db.Client
	err = db.Omit(`created_at`).Save(&model).Error
	if err == nil {
		updated = true
	}
	return
}


func (r *RoleRepo) PatchUpdate(c *gin.Context, patchData map[string]any, id uint) (updated bool, err error) {
	if id == 0 {
		updated = false
		err = errors.New(`pls input id`)
		return
	}
	model, err := r.One(c, id)
	if err != nil {
		return
	}
	if model == nil {
		err = errors.New(`no data existed`)
		return
	}

	patchDataBytes, err := json.Marshal(patchData)
	if err != nil {
		return
	}
	err = json.Unmarshal(patchDataBytes, &model)
	if err != nil {
		return
	}

	db := r.Db.Client
	err = db.Omit(`created_at`).Save(&model).Error
	if err == nil {
		updated = true
	}
	return
}


func (r *RoleRepo) Delete(c *gin.Context, id uint) (deleted bool, err error) {
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
