package db

type StorageDBparam interface {
	GetDriverDB() (err error)
	GetURIDB() (err error)
	Ping() (err error)
	Connect() (err error)
	Close() (err error)
	// may be replaced by access to db struct
	ExecuteSQL(sql string) (err error)
	BeginTransaction() (err error)
	CommitTransaction() (err error)
	RollbackTransaction() (err error)
	InTransactionNow() (t bool, err error)
}
