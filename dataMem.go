package main

type GroupStorage map[string]GroupData

type MemDB struct {
	db GroupStorage
}

func (mdb *MemDB) ReadGroupInfo(roomID string) GroupData {
	return mdb.db[roomID]
}

func (mdb *MemDB) AppendGroupInfo(roomID string, m MsgDetail) {
	mdb.db[roomID] = append(mdb.db[roomID], m)
}

func NewMemDB() *MemDB {
	return &MemDB{
		db: make(GroupStorage),
	}
}
