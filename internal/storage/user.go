package storage

//
// TODO: add JSON and SQL specification
//
type User struct {
	StorageDB StorageDBparam

	Login      string
	PasswdHash string
	Id         uint32
}

//
// TODO: realize interface StorageDBobjects
//

func (u *User) Create() (err error)
func (u *User) Read() (err error)
func (u *User) Update() (err error)
func (u *User) Delete() (err error)
func (u *User) ExistOrNot() (err error)

// User's functions
func (u *User) GetBalance() (err error)
