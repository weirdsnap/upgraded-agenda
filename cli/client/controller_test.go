package controller

import "testing"

func TestIsTelNumberValid (t *testing.T) {
  success, _ := isTelNumberValid("15521132011")
  if success {
    t.Log("test pass")
  } else {
    t.Error("test fail")
  }
}

// func TestRegister (t *testing.T) {
//   if Register("summer", "12345678", "864409241@qq.com", "15521132011") {
//     t.Log("test pass")
//   } else {
//     t.Error("test fail")
//   }
// }
