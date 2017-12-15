package controller

import (
	"testing"
)

func TestRegister(t *testing.T) {
	if Register("summer", "12345678", "864409241@qq.com", "15521132011") {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
	if Register("123", "12345678", "864409241@qq.com", "15521132011") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	if Register("summer", "123", "864409241@qq.com", "15521132011") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	if Register("summer", "12345678", "864409241", "15521132011") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	if Register("summer", "12345678", "864409241@qq.com", "1552113") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
}

func TestLogin(t *testing.T) {
	key = ""
	if Login("summer", "12345678") {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
	key = "1b23456yf"
	if Login("summer", "12345678") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
}

func TestLogout(t *testing.T) {
	key = ""
	if Logout() {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	key = "1b23456yf"
	if Logout() {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
}

func TestListUser(t *testing.T) {
	key = ""
	if ListUser() {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	key = "1b23456yf"
	if ListUser() {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
}

func TestDeleteUser(t *testing.T) {
	key = ""
	if DeleteUser() {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	key = "1b23456yf"
	if DeleteUser() {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
}

func TestCreateMeeting(t *testing.T) {
	key = ""
	participant := []string{"p1", "p2"}
	if CreateMeeting("meeting1", participant, "2017-12-12 13:00:00", "2017-12-12 14:00:00") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	key = "1b23456yf"
	if CreateMeeting("meeting1", participant, "2017-12-12 13:00:00", "2017-12-12 14:00:00") {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
	if CreateMeeting("", participant, "2017-12-12 13:00:00", "2017-12-12 14:00:00") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	if CreateMeeting("meeting1", participant, "2017-12-12 13:00", "2017-12-12 14:00:00") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	if CreateMeeting("meeting1", participant, "2017-12-12 13:00:00", "2017-12-12 14:00") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
}

func TestModifyMeeting(t *testing.T) {
	key = ""
	add := []string{"p1", "p2"}
	delete := []string{"p3", "p4"}
	if ModifyMeeting("meeting1", add, delete) {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	key = "1b23456yf"
	if ModifyMeeting("meeting1", add, delete) {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
	if ModifyMeeting("", add, delete) {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
}

func TestQueryMeeting(t *testing.T) {
	key = ""
	if QueryMeeting("2017-12-12 13:00:00", "2017-12-12 14:00:00") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	key = "1b23456yf"
	if QueryMeeting("2017-12-12 13:00:00", "2017-12-12 14:00:00") {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
	if QueryMeeting("2017-12-12 13:00", "2017-12-12 14:00:00") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	if QueryMeeting("2017-12-12 13:00:00", "2017-12-12 14:00") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
}

func TestQuitMeeting(t *testing.T) {
	key = ""
	if QuitMeeting("meeting1") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	key = "1b23456yf"
	if QuitMeeting("meeting1") {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
	if QuitMeeting("") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
}

func TestCancelMeeting(t *testing.T) {
	key = ""
	if CancelMeeting("meeting1") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	key = "1b23456yf"
	if CancelMeeting("meeting1") {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
	if CancelMeeting("") {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
}

func TestClearMeeting(t *testing.T) {
	key = ""
	if ClearMeeting() {
		t.Error("test fail")
	} else {
		t.Log("test pass")
	}
	key = "1b23456yf"
	if ClearMeeting() {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
}
