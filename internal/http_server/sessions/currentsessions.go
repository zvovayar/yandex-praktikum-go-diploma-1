package sessions

import "github.com/zvovayar/yandex-praktikum-go-diploma-1/storage"

type CurrentSessions struct {
	DB storage.StorageDBparam
}

func (cs *CurrentSessions) KillSessions()
func (cs *CurrentSessions) BeginSession() (id string)
func (cs *CurrentSessions) EndSession(id string)
