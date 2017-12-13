package entities

type userInfoDao DaoSource

var userInfoInsertStmt = "INSERT INTO userinfo(username,password,tel,email) values(?,?,?,?)"

// Save .
func (dao *userInfoDao) Save(u *UserInfo) error {
	stmt, err := dao.Prepare(userInfoInsertStmt)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(u.Username, u.Password, u.Tel,u.Email)
	checkErr(err)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.UID = int(id)
	return nil
}

var userInfoQueryAll = "SELECT * FROM userinfo"
var userInfoQueryByID = "SELECT * FROM userinfo where uid = ?"
var userInfoQueryByUsername = "SELECT * FROM userinfo where username = ?"

// FindAll .
func (dao *userInfoDao) FindAll() []UserInfo {
	rows, err := dao.Query(userInfoQueryAll)
	checkErr(err)
	defer rows.Close()

	ulist := make([]UserInfo, 0, 0)
	for rows.Next() {
		u := UserInfo{}
		err := rows.Scan(&u.UID, &u.Username, &u.Password, &u.Tel, u.Email)
		checkErr(err)
		ulist = append(ulist, u)
	}
	return ulist
}

// FindByID .
func (dao *userInfoDao) FindByID(id int) *UserInfo {
	stmt, err := dao.Prepare(userInfoQueryByID)
	checkErr(err)
	defer stmt.Close()

	row := stmt.QueryRow(id)
	u := UserInfo{}
	err = row.Scan(&u.UID, &u.Username, &u.Password, &u.Tel, u.Email)
	checkErr(err)

	return &u
}

// FindByUsername .
func (dao *userInfoDao) FindByUsername(Username string) *UserInfo {
	stmt, err := dao.Prepare(userInfoQueryByUsername)
	checkErr(err)
	defer stmt.Close()

	row := stmt.QueryRow(Username)
	u := UserInfo{}
	err = row.Scan(&u.UID, &u.Username, &u.Password, &u.Tel, &u.Email)
	checkErr(err)

	return &u
}