package helpers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("123123")) // Yönergeye göre yap.

func SetAlert(w http.ResponseWriter, r *http.Request, message string) error {

	session, err := store.Get(r, "alert-go")
	if err != nil {
		fmt.Println(err)
		return err
	}
	//sessions.Values uzun süreli verileri tutmak istediğimizde value
	session.AddFlash(message)

	return session.Save(r, w)

}

func GetAlert(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	session, err := store.Get(r, "alert-go") //İsimler aynı olacak..."alert-go"
	if err != nil {
		fmt.Println(err)
		return nil
	}

	data := make(map[string]interface{})
	flashes := session.Flashes()

	if len(flashes) > 0 {
		data["is_alert"] = true
		data["message"] = flashes[0]
	} else {
		data["is_alert"] = false
		data["message"] = nil
	}

	err = session.Save(r, w)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
