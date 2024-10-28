package controllers

import (
	"BlogSitesi/admin/helpers"
	"BlogSitesi/admin/models"
	"crypto/sha256"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type Userops struct{}

func (userops Userops) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("userops/login")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (userops Userops) Login(w http.ResponseWriter, r *http.Request, paramas httprouter.Params) {
	username := r.FormValue("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))

	user := models.User{}.Get("username=? AND password=?", username, password)
	if user.Username == username && user.Password == password {
		//Login olmuş oluyor
		helpers.SetUser(w, r, username, password)
		helpers.SetAlert(w, r, "Hoşgelmişsennn")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		//Giriş yapılamadı
		helpers.SetAlert(w, r, "Şifre Veya Kullanıcı Adını Yanlış Girdiniz.")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}

}

func (userops Userops) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helpers.RemoveUser(w, r)
	helpers.SetAlert(w, r, "Hoşçakal Olacaklar Sensiz Olsun")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}
