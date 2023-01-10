package main

type PSqlDB struct {
	db GroupStorage
}

func (mdb *PSqlDB) ReadGroupInfo(roomID string) GroupData {
	return mdb.db[roomID]
}

func (mdb *PSqlDB) AppendGroupInfo(roomID string, m MsgDetail) {
	mdb.db[roomID] = append(mdb.db[roomID], m)
}

func NewPQSql() *PSqlDB {
	return &PSqlDB{
		db: make(GroupStorage),
	}
}
