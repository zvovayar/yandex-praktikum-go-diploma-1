package storage

//
// TODO: add JSON and SQL specification
//
type User struct {
	StorageDB StorageDBparam

	Login      string `json:"login"`
	PasswdHash string `json:"password"`
	ID         uint32
}

//
// TODO: realize interface StorageDBobjects
//

func (u *User) Create() (err error)     { return nil }
func (u *User) Read() (err error)       { return nil }
func (u *User) Update() (err error)     { return nil }
func (u *User) Delete() (err error)     { return nil }
func (u *User) ExistOrNot() (err error) { return nil }

// User's functions
func (u *User) GetBalance() (err error) { return nil }
