package validators

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"github.com/imoowi/comer/utils/response"
)

func TestValidator(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	InitValidators()
	postData := PostData{
		Name:   `中文名`,
		NameEn: `english_name`,
		Mobile: `13800138000`,
		Email:  `imoowi@qq.com`,
		IDCard: `11010120240402781X`,
	}
	dataByte, _ := json.Marshal(postData)
	payload := strings.NewReader(string(dataByte))
	req, _ := http.NewRequest(`POST`, `/dopost`, payload)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	body, _ := io.ReadAll(w.Body)
	_body := string(body)
	assert.Equal(t, true, _body == `true`)

	log.Println(_body)
}

type PostData struct {
	Name   string `json:"name" form:"name" binding:"required,chinese"`
	NameEn string `json:"name_en" form:"name_en" binding:"required,english"`
	Mobile string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Email  string `json:"email" form:"email" binding:"required,email"`
	IDCard string `json:"id_card" form:"id_card" binding:"required,idcard"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST(`/dopost`, func(ctx *gin.Context) {
		postData := PostData{}
		err := ctx.ShouldBindJSON(&postData)
		if err != nil {
			response.Error(err.Error(), http.StatusBadRequest, ctx)
			return
		}
		response.OK(true, ctx)
	})
	return r
}
