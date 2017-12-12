package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type User struct {
	Username  string
	Password  string
	Email     string
	Telephone string
}

var key string = "1b23456yf"

// var currentUser *User = nil

func Register(username, password, email, telphone string) {
	// var a, b, c, d bool
	// var err error
	//合法性检查
	a, err := isUserNameValid(username)
	if false == a {
		fmt.Println("username fail", err)
		return
	}
	b, err := isPasswordValid(password)
	if false == b {
		fmt.Println("password fail", err)
		return
	}
	c, err := isEmailValid(email)
	if false == c {
		fmt.Println("email fail", err)
		return
	}
	d, err := isTelNumberValid(telphone)
	if false == d {
		fmt.Println("telphone fail", err)
		return
	}
	//创建user对象
	user := struct {
		Username  string
		Password  string
		Email     string
		Telephone string
	}{username, password, email, telphone}

	bUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json error:", err)
	}

	//发送请求
	post_body := bytes.NewBuffer([]byte(bUser))
	res, err := http.Post("https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/user/register", "application/json;charset=utf-8", post_body)
	if err != nil {
		fmt.Println("Post failed, error:", err)
	}
	defer res.Body.Close()
	//转换回user对象
	res_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err1 := json.Unmarshal(res_body, &user)
	if err != nil {
		panic(err1)
	}
	//检查返回的结果
	if user.Username == "null" && user.Password == "null" && user.Email == "null" && user.Telephone == "null" {
		fmt.Println("register failed!")
		return
	} else {
		fmt.Println("register success!")
		return
	}
}

func isLogined() bool {
	if key == "" {
		return false
	} else {
		return true
	}
}

func Login(username, password string) {
	//检测是否已经有登陆用户
	if isLogined() {
		fmt.Println("Login failed! Error : already Logined. Please logout first")
		return
	}
	//发送请求
	prefix := "https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/user/login/"
	parameters := "?username=" + username + "&password=" + password
	res, err := http.Get(prefix + parameters)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	//解析返回的响应
	res_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(res_body))
	resKey := struct {
		Key string
	}{"1b23456yf"}

	err1 := json.Unmarshal(res_body, &resKey)
	if err != nil {
		panic(err1)
	}
	if resKey.Key == "" {
		fmt.Println("Login failed! Error: username and password unmatch!")
		return
	}
	key = resKey.Key
	fmt.Println("Login success!")
	return

}

func checkKey() bool {
	prefix := "https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/user/verify/"
	parameters := "?key=" + key
	res, err := http.Get(prefix + parameters)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	success := struct {
		Success bool
	}{true}
	res_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err1 := json.Unmarshal(res_body, &success)
	if err != nil {
		panic(err1)
	}
	if success.Success {
		return true
	}
	return false
}

func Logout() {
	if !isLogined() {
		fmt.Println("Logout failed! Error: no user login now.")
		return
	}
	//先检查用户是否在服务器上处于登陆状态
	if checkKey() {
		prefix := "https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/user/logout/"
		parameters := "?key=" + key
		res, err := http.Get(prefix + parameters)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		success := struct {
			Success bool
		}{true}
		res_body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		err1 := json.Unmarshal(res_body, &success)
		if err != nil {
			panic(err1)
		}
		//退出登陆的同时把key的信息清除掉
		if success.Success {
			fmt.Println("Logout success!")
			key = ""
			return
		}
		fmt.Println("Logout failed!")
		return
	} else {
		fmt.Println("Logout failed! Error: no user login now.")
		return
	}
}

func ListUser() {
	if !isLogined() {
		fmt.Println("Please Log in first!")
		return
	}
	if checkKey() {
		prefix := "https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/users/"
		parameters := "?key=" + key
		res, err := http.Get(prefix + parameters)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		//创建数组用来装用户
		users := make([]User, 0)

		res_body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		err1 := json.Unmarshal(res_body, &users)
		if err != nil {
			panic(err1)
		}
		//打印所有的用户
		for _, user := range users {
			fmt.Println(user.Username)
			fmt.Println(user.Email)
			fmt.Println(user.Telephone)
			fmt.Println(" ")
		}

	} else {
		fmt.Println("Please Log in first!")
		return
	}

}

func DeleteUser() {
	if !isLogined() {
		fmt.Println("Delete failed! Error: no user login now.")
		return
	}
	//先检查用户是否在服务器上处于登陆状态
	if checkKey() {
		prefix := "https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/user/delete/"
		parameters := "?key=" + key
		res, err := http.Get(prefix + parameters)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		success := struct {
			Success bool
		}{true}
		res_body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		err1 := json.Unmarshal(res_body, &success)
		if err != nil {
			panic(err1)
		}
		//删除的同时把key的信息清除掉
		if success.Success {
			fmt.Println("Delete success!")
			key = ""
			return
		}
		fmt.Println("Delete failed!")
		return
	} else {
		fmt.Println("Delete failed! Error: no user login now.")
		return
	}
}

func CreateMeeting(title string, participators []string, starttime string, endtime string) {
	// initialization()
	// if isLogined() {
	// 	for _, s := range participators {
	// 		if users.QueryUser(s) == nil {
	// 			fmt.Println("Create Meeting failed! invalid user")
	// 			return
	// 		}
	// 	}
	// 	t, _ := isTimeValid(starttime)
	// 	r, _ := isTimeValid(endtime)
	// 	if t == false || r == false {
	// 		fmt.Println("wrong time")
	// 		return
	// 	}
	// 	if meetings.AddMeeting(NewMeeting(title, starttime, endtime, currentUser.Username, participators)) == false {
	// 		fmt.Println("filed!")
	// 		return
	// 	}
	// 	fmt.Println("Create Meeting successed!")
	// } else {
	// 	fmt.Println("Please login first!")
	// }
	// update()
	// return
}

func ModifyMeeting(title string, addedparticipators []string, deletedparticipators []string) {
	// initialization()
	// if isLogined() {
	// 	//fmt.Println("add user", addedparticipators[0], len(addedparticipators))
	// 	if addedparticipators != nil && addedparticipators[0] != "" {
	// 		for _, s := range addedparticipators {
	// 			if users.QueryUser(s) == nil {
	// 				fmt.Println("add participators failed! invalid user")
	// 				return
	// 			}
	// 		}
	// 		if meetings.AddParticipants(title, addedparticipators) == false {
	// 			fmt.Println("Modify Meeting failed! invalid title or add user")
	// 			return
	// 		}
	// 	}
	// 	if deletedparticipators != nil && deletedparticipators[0] != "" {
	// 		for _, s := range deletedparticipators {
	// 			if users.QueryUser(s) == nil {
	// 				fmt.Println("delete participators failed! invalid user")
	// 				return
	// 			}
	// 		}
	// 		if meetings.DeleteParticipants(title, deletedparticipators) == false {
	// 			fmt.Println("Modify Meeting failed! invalid title or delete user")
	// 			return
	// 		}
	// 	}
	// 	fmt.Println("Modify Meeting successed!")
	// }
	// update()
	// return
}

func QueryMeeting(starttime string, endtime string) {
	// initialization()
	//
	// if isLogined() {
	// 	t, _ := isTimeValid(starttime)
	// 	r, _ := isTimeValid(endtime)
	// 	if t == false || r == false {
	// 		fmt.Println("time wrong!")
	// 	}
	// 	meeting := meetings.QueryMeeting(starttime, endtime, currentUser.Username)
	// 	for _, value := range meeting {
	// 		fmt.Println(value)
	// 	}
	// }
	// update()
	// return
}

func QuitMeeting(title string) {
	// initialization()
	// if isLogined() {
	// 	if meetings.QuitMeeting(title, currentUser.Username) {
	// 		fmt.Println("quit successed!")
	// 	} else {
	// 		fmt.Println("title wrong or you aren't hostor!")
	// 	}
	// }
	// update()
	// return
}

func CancelMeeting(title string) {

	// initialization()
	// if isLogined() {
	// 	if meetings.CancelMeeting(title, currentUser.Username) {
	// 		fmt.Println("meeting cancle successed!")
	// 	} else {
	// 		fmt.Println("meeting title wrong or you aren't hostor!")
	// 	}
	// }
	// update()
	// return
}

func ClearMeeting() {

	// initialization()
	// if isLogined() {
	// 	if meetings.ClearMeeting(currentUser.Username) {
	// 		fmt.Println("clear meeting successed!")
	// 	}
	// }
	// update()
	// return
}

func isUserNameValid(username string) (bool, error) {
	m, err := regexp.MatchString("^[a-zA-Z]{4,30}$", username)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func isPasswordValid(password string) (bool, error) {
	m, err := regexp.MatchString("^[0-9a-zA-Z@.]{6,30}$", password)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func isEmailValid(email string) (bool, error) {
	m, err := regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", email)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func isTelNumberValid(telNum string) (bool, error) {
	m, err := regexp.MatchString("^[0-9]{11}$", telNum)
	if m {
		return true, err
	} else {
		return false, err
	}
}

func isTimeValid(time string) (bool, error) {
	m, err := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}\\s[0-9]{2}:[0-9]{2}:[0-9]{2}$", time)
	if m {
		return true, err
	} else {
		return false, err
	}
}
