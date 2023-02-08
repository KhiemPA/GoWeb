package routes

import (
	"fmt"
	"net/http"
	"prj/models"
	"prj/sessions"
	"prj/utils"

	
)

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := sessions.Flash(w, r)
	utils.ExcuteTemplate(w, "register.html", struct {
		Alert utils.Alert
	}{
		Alert : utils.NewAlert(message, alert),
	})
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user models.User

	user.FirstName = r.PostForm.Get("firstname")
	user.LastName = r.PostForm.Get("lastname")
	user.Email = r.PostForm.Get("email")
	user.Password = r.PostForm.Get("password")
	fmt.Println(user)
	_,err := models.NewUser(user)
	checkErrRegister(w, r,  err)
}

func checkErrRegister(w http.ResponseWriter, r *http.Request, err error) {
	
	message := "Register success"
	if err != nil {
		switch(err) {
		case models.ErrRequiredFirstName,
			models.ErrRequiredLastName,
			models.ErrRequiredEmail,
			models.ErrInvalidEmail,
			models.ErrRequiredPassword,
			models.ErrMaxLimit,
			models.ErrDuplicateKeyEmail:
			message = fmt.Sprintf("%s", err)
			break
		default:
			utils.InternalServerError(w)
			return
		}
		sessions.Message(message, "danger", w, r)
		http.Redirect(w, r, "/register", 302)
		return
	}
	sessions.Message(message, "success", w, r)
	http.Redirect(w, r, "/login", 302)
}

func usersGetHandler(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	total := int64(len(users))
	utils.ExcuteTemplate(w, "user.html", struct{
		Users []models.User
		Total int64
	}{
		Users: users,
		Total: total,
	})
}