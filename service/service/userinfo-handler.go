package service

import (
	"github.com/unrolled/render"
	"github.com/weirdsnap/upgraded-agenda/service/entities"
	// "github.com/weirdsnap/agendaweb/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var index string = "86253e9"

type User struct {
	Username string
	Password string
	Email    string
	Telphone string
}

func getUserLoginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		username := req.Form["username"]
		password := req.Form["password"]
		fmt.Println("username:", username, " - ", password)
		u := entities.UserInfoService.FindByUsername(string(username[0]))
		if strings.EqualFold(u.Password, password[0]) {
			formatter.JSON(w, http.StatusOK, struct {
				KEY string `json:"key"`
			}{KEY: index + strconv.Itoa(u.UID)})
		} else {
			formatter.JSON(w, http.StatusOK, struct {
				KEY string `json:"key"`
			}{KEY: "null"})
		}

	}
}

func psotUserRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		result, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()

		fmt.Printf("%s\n", result)

		var user User
		json.Unmarshal([]byte(result), &user)

		u := entities.NewUserInfo(entities.UserInfo{Username: user.Username})
		u.Password = user.Password
		u.Email = user.Email
		u.Tel = user.Telphone
		// fmt.Println("u",u)
		entities.UserInfoService.Save(u)

		formatter.JSON(w, http.StatusOK, user)
	}
}

func getUserVerifyHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		//username := req.Form["username"]
		//	password := req.Form["password"]

	}
}

func getUserLogoutHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		//	username := req.Form["username"]
		//	password := req.Form["password"]

	}
}

// func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {

// 	return func(w http.ResponseWriter, req *http.Request) {
// 		req.ParseForm()
// 		if len(req.Form["username"][0]) == 0 {
// 			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
// 			return
// 		}
// 		u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
// 		u.DepartName = req.Form["departname"][0]
// 		entities.UserInfoService.Save(u)
// 		formatter.JSON(w, http.StatusOK, u)
// 	}
// }

// func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {

// 	return func(w http.ResponseWriter, req *http.Request) {
// 		req.ParseForm()
// 		if len(req.Form["userid"][0]) != 0 {
// 			i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)

// 			u := entities.UserInfoService.FindByID(int(i))
// 			formatter.JSON(w, http.StatusBadRequest, u)
// 			return
// 		}
// 		ulist := entities.UserInfoService.FindAll()
// 		formatter.JSON(w, http.StatusOK, ulist)
// 	}
// }
