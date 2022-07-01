package business_logic

import (
	httpcs "github.com/zvovayar/yandex-praktikum-go-diploma-1/http_server/sessions"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/storage"
)

type BusinessSession struct {
	HTTPsession httpcs.CurrentSession
}

func (bs *BusinessSession) RegisterNewUser(u storage.Order) (err error)
func (bs *BusinessSession) Buy() (err error)
func (bs *BusinessSession) LoadOrder() (err error)
func (bs *BusinessSession) Withdraw() (err error)
