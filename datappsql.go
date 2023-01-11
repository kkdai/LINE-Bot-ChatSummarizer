package main

import (
	"log"

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
		(*MemStorage)(nil),
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

// DBStorage: for orm db storage.
type DBStorage struct {
	Id      int64     `bson:"_id"`
	RoomID  string    `json:"roomid" bson:"roomid"`
	Dataset GroupData `json:"dataset" bson:"dataset"`
}

func (u *DBStorage) Add(conn *PSqlDB) {
	_, err := conn.Db.Model(u).Insert()
	if err != nil {
		log.Println(err)
	}
}

func (u *DBStorage) Get(conn *PSqlDB) (result *DBStorage, err error) {
	log.Println("***Get dataset roomID=", u.RoomID)
	data := DBStorage{}
	err = conn.Db.Model(&data).
		Where("Room ID = ?", u.RoomID).
		Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("DB result= ", data)
	return &data, nil
}

func (u *DBStorage) Update(conn *PSqlDB) (err error) {
	log.Println("***Update DB group data=", u)

	_, err = conn.Db.Model(u).
		Set("dataset = ?", u.Dataset).
		Where("roomid = ?", u.RoomID).
		Update()
	if err != nil {
		log.Println(err)
	}
	return nil
}
