package main

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type PGSqlDB struct {
	Db *pg.DB
}

func (mdb *PGSqlDB) ReadGroupInfo(roomID string) GroupData {
	pgsql := &DBStorage{
		RoomID: roomID,
	}
	if ret, err := pgsql.Get(mdb); err == nil {
		return ret.Dataset
	} else {
		log.Println("DB read err:", err)
	}

	return GroupData{}
}

func (mdb *PGSqlDB) AppendGroupInfo(roomID string, m MsgDetail) {
	u := mdb.ReadGroupInfo(roomID)
	u = append(u, m)
	pgsql := &DBStorage{
		RoomID: roomID,
	}
	if err := pgsql.Update(mdb); err != nil {
		log.Println("DB update err:", err)
	}
}

func NewPGSql(url string) *PGSqlDB {
	options, _ := pg.ParseURL(url)
	db := pg.Connect(options)

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	return &PGSqlDB{
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

// DBStorage for orm db storage.
type DBStorage struct {
	Id      int64     `bson:"_id"`
	RoomID  string    `json:"roomid" bson:"roomid"`
	Dataset GroupData `json:"dataset" bson:"dataset"`
}

func (u *DBStorage) Add(conn *PGSqlDB) {
	_, err := conn.Db.Model(u).Insert()
	if err != nil {
		log.Println(err)
	}
}

func (u *DBStorage) Get(conn *PGSqlDB) (result *DBStorage, err error) {
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

func (u *DBStorage) Update(conn *PGSqlDB) (err error) {
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
