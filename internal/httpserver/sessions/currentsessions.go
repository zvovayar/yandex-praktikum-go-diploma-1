package sessions

import "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/storage"

type CurrentSessions struct {
	DB storage.StorageDBparam
}

func (cs *CurrentSessions) KillSessions()             {}
func (cs *CurrentSessions) BeginSession() (id string) { return "" }
func (cs *CurrentSessions) EndSession(id string)      {}
