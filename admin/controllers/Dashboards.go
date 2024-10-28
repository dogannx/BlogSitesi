package controllers

import (
	"BlogSitesi/admin/helpers"
	"BlogSitesi/admin/models"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Dashboards struct {
}

func (dasboards Dashboards) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory": func(Category_id int) string {
			return models.Category{}.Get(Category_id).Title
		},
	}).ParseFiles(helpers.Include("dashboard/list")...)

	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (dasboards Dashboards) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	err = view.ExecuteTemplate(w, "index", data)
	if err != nil {
		return
	}
}

func (dasboards Dashboards) Yeni(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	title := r.FormValue("blog-title") //r. fonksiyonlar ile veriyi alıyoruz.
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")

	//Upload İşlemleri
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("blog-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(f, file)
	// Upload End
	if err != nil {
		fmt.Println(err)
		return
	}

	models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		Content:     content,
		Category_id: categoryID,
		Picture_url: "uploads/" + header.Filename,
	}.Add()

	helpers.SetAlert(w, r, "Kaydedildi...Sonunda")

	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}

func (dashboards Dashboards) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	post := models.Post{}.Get(params.ByName("id"))
	post.Delete()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboards) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)

	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Post"] = models.Post{}.Get(params.ByName("id"))
	data["Categories"] = models.Category{}.GetAll()
	err = view.ExecuteTemplate(w, "index", data)
	if err != nil {
		return
	}
}

func (dashboard Dashboards) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if !helpers.CheckUser(w, r) {
		return
	}

	post := models.Post{}.Get(params.ByName("id")) // Hangi id'li elemanı değiştrieceksek onu çekmiş olduk.
	title := r.FormValue("blog-title")             //r. fonksiyonlar ile veriyi alıyoruz.
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")
	is_selected := r.FormValue("is_selected")
	var picture_url string

	if is_selected == "1" {
		//is_selecetd değişir.
		//Upload İşlemleri
		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("blog-picture")
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = io.Copy(f, file)
		// Upload End
		picture_url = "uploads/" + header.Filename
		os.Remove(post.Picture_url)
	} else {
		//is_selected aynı kalır.
		picture_url = post.Picture_url
	}
	post.Updates(models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		Category_id: categoryID,
		Content:     content,
		Picture_url: picture_url,
	})
	http.Redirect(w, r, "/admin/edit/"+params.ByName("id"), http.StatusSeeOther)
}
