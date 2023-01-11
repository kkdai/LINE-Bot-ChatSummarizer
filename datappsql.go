package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type PSqlDB struct {
	Db *pg.DB
}

func (mdb *PSqlDB) ReadGroupInfo(roomID string) GroupData {
	return GroupData{}
}

func (mdb *PSqlDB) AppendGroupInfo(roomID string, m MsgDetail) {
	// mdb.db[roomID] = append(mdb.db[roomID], m)
}

func NewPQSql(url string) *PSqlDB {
	options, _ := pg.ParseURL(url)
	db := pg.Connect(options)

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	return &PSqlDB{
		Db: db,
	}
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*GroupStorage)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true})
		if err != nil {
			return err
		}
	}
	return nil
}
