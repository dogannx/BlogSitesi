package main

import (
	admin_models "BlogSitesi/admin/models"
	"BlogSitesi/config"
	"net/http"
)

func main() {

	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()

	/*	admin_models.Post{
		Title: "C++ İle Web Programlama",
		Slug:  "c++",
	}.Add()*/

	/*post := admin_models.Post{}.Get(1) ////// post := admin_models.Post{}.Get("slug:=?","deneme")
	fmt.Println(post.Title)*/

	/*	post := admin_models.Post{}.GetAll(1)
		fmt.Println(post)*/

	/*	fmt.Println(admin_models.Post{}.GetAll("title=?", "Go İle Web Programlama")) // Tüm satırı getirir önümüze
	 */

	/*post := admin_models.Post{}.Get(3)
	post.Update("slug", "Bilal")*/ // Satırın bir bölümünü değiştirme

	/*post := admin_models.Post{}.Get(1)
	post.Updates(admin_models.Post{Title: "Yeni", Description: "Test"})*/ //Bir Satırda İstediğim tüm değişkenleri değiştirmek için

	/*	post := admin_models.Post{}.Get(1)
		post.Delete()*/

	http.ListenAndServe(":8080", config.Routes())

}
