package businesslogic

import (
	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
	httpcs "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/httpserver/sessions"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/storage"
)

type BusinessSession struct {
	HTTPsession httpcs.CurrentSession
}

func (bs *BusinessSession) RegisterNewUser(u storage.User) (err error) {

	config.LoggerCLS.Debug("register new user " + u.Login)
	return nil
}
func (bs *BusinessSession) UserLogin(u storage.User) (err error) {

	config.LoggerCLS.Debug("login user " + u.Login)
	return nil
}
func (bs *BusinessSession) Buy() (err error)                { return nil }
func (bs *BusinessSession) LoadOrder(oc string) (err error) { return nil }
func (bs *BusinessSession) Withdraw() (err error)           { return nil }
