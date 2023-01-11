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

type DBStorage struct {
	ID      string
	Dataset GroupData
}

func (u *DBStorage) Add(conn *PSqlDB) {
	_, err := conn.Db.Model(u).Insert()
	if err != nil {
		log.Println(err)
	}
}

func (u *DBStorage) Get(conn *PSqlDB) (result *DBStorage, err error) {
	log.Println("***Get dataset uUID=", u.ID)
	data := DBStorage{}
	err = conn.Db.Model(&data).
		Where("ID = ?", u.ID).
		Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("DB result= ", data)
	return &data, nil
}

// func (u *UserFavorite) Update(meta *models.Model) (err error) {
// 	log.Println("***Update Fav User=", u)

// 	_, err = meta.Db.Model(u).
// 		Set("favorites = ?", u.Favorites).
// 		Where("user_id = ?", u.UserId).
// 		Update()
// 	if err != nil {
// 		meta.Log.Println(err)
// 	}
// 	return nil
// }
