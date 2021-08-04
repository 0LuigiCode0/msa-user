package database

import (
	"fmt"

	corehelper "x-msa-core/helper"
	"x-msa-user/helper"

	mongoStore "x-msa-user/store/mongo"

	"github.com/0LuigiCode0/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	_mongo = "mongo"
)

type DB interface {
	Mongo() *mongo.Database
	MongoStore() mongoStore.Store
	Close()
}

type DBForHandler interface {
	MongoStore() mongoStore.Store
}

type db struct {
	_mongo *d
}

type d struct {
	store  interface{}
	conn   interface{}
	dbName string
}

func InitDB(conf *helper.Config) (DB DB, err error) {
	db := &db{}
	DB = db

	var conn interface{}
	if v, ok := conf.DBS[_mongo]; ok {
		conn, err = connMongo(v)
		if err != nil {
			return nil, fmt.Errorf("db not initializing: %v", err)
		}
		db._mongo = &d{
			store:  mongoStore.InitStore((*mongo.Database)(conn.(*mongo.Client).Database(v.Name))),
			dbName: v.Name,
			conn:   conn,
		}
		logger.Log.Servicef("db %q initializing", _mongo)
	}

	logger.Log.Service("db initializing")
	return
}

func (d *db) Close() {
	d._mongo.conn.(*mongo.Client).Disconnect(corehelper.Ctx)
	logger.Log.Servicef("db %q stoped", _mongo)
}

func (d *db) Mongo() *mongo.Database       { return d._mongo.conn.(*mongo.Client).Database(d._mongo.dbName) }
func (d *db) MongoStore() mongoStore.Store { return d._mongo.store.(mongoStore.Store) }

func connMongo(v *helper.DbConfig) (conn *mongo.Client, err error) {
	opt := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", v.Host, v.Port)).SetAuth(options.Credential{AuthMechanism: "SCRAM-SHA-256", Username: v.User, Password: v.Password})
	conn, err = mongo.Connect(corehelper.Ctx, opt)
	if err != nil {
		return conn, fmt.Errorf("db not connected: %v", err)
	}
	if err = conn.Ping(corehelper.Ctx, nil); err != nil {
		return conn, fmt.Errorf("db not pinged: %v", err)
	}
	return
}
