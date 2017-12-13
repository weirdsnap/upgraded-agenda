package entities
// import ("fmt")
//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {

	tx, err := mydb.Begin()
	checkErr(err)
	// fmt.Println("开始保存")
	dao := userInfoDao{tx}
	err = dao.Save(u)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	dao := userInfoDao{mydb}
	return dao.FindAll()
}

// FindByUsername .
func (*UserInfoAtomicService) FindByUsername(username string) *UserInfo {
	dao := userInfoDao{mydb}
	return dao.FindByUsername(username)
}
