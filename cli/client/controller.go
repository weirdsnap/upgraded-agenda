package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type User struct {
	Username  string
	Password  string
	Email     string
	Telephone string
}
type Meeting struct {
	Host         string
	Title        string
	Participants []string
	Start        string
	End          string
}

//测试时用
// var key string = "1b23456yf"
var key string = ""

func Register(username, password, email, telphone string) bool {
	//合法性检查
	a, err := isUserNameValid(username)
	if false == a {
		fmt.Println("username fail", err)
		return false
	}
	b, err := isPasswordValid(password)
	if false == b {
		fmt.Println("password fail", err)
		return false
	}
	c, err := isEmailValid(email)
	if false == c {
		fmt.Println("email fail", err)
		return false
	}
	d, err := isTelNumberValid(telphone)
	if false == d {
		fmt.Println("telphone fail", err)
		return false
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
		return false
	} else {
		fmt.Println("register success!")
		return true
	}
}

func isLogined() bool {
	if key == "" {
		return false
	} else {
		return true
	}
}

func Login(username, password string) bool {
	//检测是否已经有登陆用户
	if isLogined() {
		fmt.Println("Login failed! Error : already Logined. Please logout first")
		return false
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
	resKey := struct {
		Key string
	}{"1b23456yf"}

	err1 := json.Unmarshal(res_body, &resKey)
	if err != nil {
		panic(err1)
	}
	if resKey.Key == "" {
		fmt.Println("Login failed! Error: username and password unmatch!")
		return false
	}
	//登陆成功后在本地记录返回的key
	key = resKey.Key
	fmt.Println("Login success!")
	return true

}

//检查服务器端该key对应的用户是否处在登陆状态
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

func Logout() bool {
	//检查当前是否有用户登陆
	if !isLogined() {
		fmt.Println("Logout failed! Error: no user login now.")
		return false
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
			return true
		}
		fmt.Println("Logout failed!")
		return false
	} else {
		fmt.Println("Logout failed! Error: no user login now.")
		return false
	}
}

func ListUser() bool {
	if !isLogined() {
		fmt.Println("Please Log in first!")
		return false
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
		return true

	} else {
		fmt.Println("Please Log in first!")
		return false
	}

}

func DeleteUser() bool {
	if !isLogined() {
		fmt.Println("Delete failed! Error: no user login now.")
		return false
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
			return true
		}
		fmt.Println("Delete failed!")
		return false
	} else {
		fmt.Println("Delete failed! Error: no user login now.")
		return false
	}
}

func CreateMeeting(title string, participators []string, starttime string, endtime string) bool {
	//检测是否已经有登陆用户
	if !isLogined() {
		fmt.Println("Create meeting failed! Error: no user login now.")
		return false
	}
	//检查会议名称是否为空
	if title == "" {
		fmt.Println("Create meeting failed! Error: meeting must have a title!")
		return false
	}
	//检查时间格式的合法性
	s, _ := isTimeValid(starttime)
	e, _ := isTimeValid(endtime)
	if s == false || e == false {
		fmt.Println("Create meeting failed! Error: time format invalid!")
		return false
	}
	if checkKey() {
		meeting := struct {
			Key          string
			Title        string
			Participants []string
			Start        string
			End          string
		}{key, title, participators, starttime, endtime}
		Jmeeting, err := json.Marshal(meeting)
		if err != nil {
			panic(err)
		}
		post_body := bytes.NewBuffer([]byte(Jmeeting))
		res, err := http.Post("https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/meeting/create", "application/json;charset=utf-8", post_body)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		//转换回user对象
		res_body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		err1 := json.Unmarshal(res_body, &meeting)
		if err != nil {
			panic(err1)
		}
		//只检测key可以吗？
		if meeting.Key == "null" {
			fmt.Println("Create meeting failed!")
			return false
		} else {
			fmt.Println("Create meeting success!")
			fmt.Println(meeting.Key)
			fmt.Println(meeting.Title)
			fmt.Println(meeting.Participants)
			fmt.Println(meeting.Start)
			fmt.Println(meeting.End)
			return true
		}
	} else {
		fmt.Println("Create meeting failed! Error : no user log in now.")
		return false
	}
}

func ModifyMeeting(title string, addedparticipators []string, deletedparticipators []string) bool {
	//检测是否已经有登陆用户
	if !isLogined() {
		fmt.Println("Modify meeting failed! Error: no user login now.")
		return false
	}
	if title == "" {
		fmt.Println("Modify meeting failed! Error: meeting must have a title!")
		return false
	}
	if checkKey() {
		meeting := struct {
			Key    string
			Title  string
			Add    []string
			Delete []string
		}{key, title, addedparticipators, deletedparticipators}
		Jmeeting, err := json.Marshal(meeting)
		if err != nil {
			panic(err)
		}
		post_body := bytes.NewBuffer([]byte(Jmeeting))
		res, err := http.Post("https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/meeting/create", "application/json;charset=utf-8", post_body)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		//转换回user对象
		res_body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		err1 := json.Unmarshal(res_body, &meeting)
		if err != nil {
			panic(err1)
		}
		//只检测key可以吗？
		if meeting.Key == "null" {
			fmt.Println("Modify meeting failed!")
			return false
		} else {
			fmt.Println("Modify meeting success!")
			return true
		}
	} else {
		fmt.Println("Modify meeting failed! Error : no user log in now.")
		return false
	}
}

func QueryMeeting(starttime string, endtime string) bool {
	//检测是否已经有登陆用户
	if !isLogined() {
		fmt.Println("Query meeting failed! Error: no user login now.")
		return false
	}
	s, _ := isTimeValid(starttime)
	e, _ := isTimeValid(endtime)
	if s == false || e == false {
		fmt.Println("Query meeting failed! Error: time format invalid!")
		return false
	}
	//将空格替换成%，这样发出去的请求格式才是正确的
	starttime = strings.Replace(starttime, " ", "%", -1)
	endtime = strings.Replace(endtime, " ", "%", -1)
	if checkKey() {
		prefix := "https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/meetings/query/"
		parameters := "?key=" + key + "&start=" + starttime + "&end=" + endtime
		res, err := http.Get(prefix + parameters)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		//创建数组用来装会议
		meetings := make([]Meeting, 0)

		res_body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		err1 := json.Unmarshal(res_body, &meetings)
		if err != nil {
			panic(err1)
		}
		//打印所有的用户
		for _, meeting := range meetings {
			fmt.Println(meeting.Host)
			fmt.Println(meeting.Title)
			fmt.Println(meeting.Participants)
			fmt.Println(meeting.Start)
			fmt.Println(meeting.End)
			fmt.Println(" ")
		}
		return true

	} else {
		fmt.Println("Please Log in first!")
		return false
	}

}

func QuitMeeting(title string) bool {
	//检测是否已经有登陆用户
	if !isLogined() {
		fmt.Println("Quit meeting failed! Error: no user login now.")
		return false
	}
	if title == "" {
		fmt.Println("Modify meeting failed! Error: meeting must have a title!")
		return false
	}
	if checkKey() {
		prefix := "https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/meeting/quit/"
		parameters := "?key=" + key + "&title=" + title
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
			fmt.Println("Quit success!")
			return true
		} else {
			fmt.Println("Quit failed!")
			return false
		}

	} else {
		fmt.Println("Please Log in first!")
		return false
	}
}

func CancelMeeting(title string) bool {
	//检测是否已经有登陆用户
	if !isLogined() {
		fmt.Println("cancel meeting failed! Error: no user login now.")
		return false
	}
	if title == "" {
		fmt.Println("Modify meeting failed! Error: meeting must have a title!")
		return false
	}
	if checkKey() {
		prefix := "https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/meeting/cancel/"
		parameters := "?key=" + key + "&title=" + title
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
			fmt.Println("Cancel success!")
			return true
		} else {
			fmt.Println("Cancel failed!")
			return false
		}

	} else {
		fmt.Println("Please Log in first!")
		return false
	}
}

func ClearMeeting() bool {
	//检测是否已经有登陆用户
	if !isLogined() {
		fmt.Println("delete meetings failed! Error: no user login now.")
		return false
	}
	if checkKey() {
		prefix := "https://private-6e5eb4a-agendav2.apiary-mock.com/agenda/v2/meetings/delete/"
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
			fmt.Println("Delete success!")
			return true
		} else {
			fmt.Println("Delete failed!")
			return false
		}

	} else {
		fmt.Println("Please Log in first!")
		return false
	}
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
