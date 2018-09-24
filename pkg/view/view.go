package view

import (
	"net/http"

	"github.com/CRUD/pkg/model"
)

type IndexData struct {
	List []*model.News
}

func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}
func News(w http.ResponseWriter, data *model.News) {
	render(tpNews, w, data)
}
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

type AdminListData struct {
	List []*model.News
}

func AdminList(w http.ResponseWriter, data *AdminListData) {
	render(tpAdminList, w, data)
}
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreate, w, data)
}
func AdminEdit(w http.ResponseWriter, data *model.News) {
	render(tpAdminEdit, w, data)
}
