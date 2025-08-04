package db_postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	sharedConfigs "project1.v0/configs"
)

type DbDeps struct {
	DbConfig *sharedConfigs.DbConfig
}

type Db struct {
	*sql.DB
	conf *sharedConfigs.DbConfig
}

func NewDb(deps DbDeps) *Db {
	fmt.Println("conf.Dsn=", deps.DbConfig.Dsn)
	db, err := sql.Open("postgres", deps.DbConfig.Dsn)
	if err != nil {
		log.Fatalf("%v: %v", ErrDbOpenError.Error(), err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("%v: %v", ErrDbPingError.Error(), err)
	}

	db.SetMaxOpenConns(deps.DbConfig.MaxOpenConns)
	db.SetMaxIdleConns(deps.DbConfig.MaxIdleConns)
	db.SetConnMaxIdleTime(deps.DbConfig.ConnMaxIdleTime)
	return &Db{
		conf: deps.DbConfig,
		DB:   db,
	}
}

type NewDbConstructor func(deps DbDeps) *Db

/*
func (db *Db) BeginTransaction(isolation sql.IsolationLevel, readOnly bool) {
	//sql.LevelDefault
	//sql.LevelLinearizable
	//sql.LevelReadCommitted
	//sql.LevelReadUncommitted
	//sql.LevelRepeatableRead
	//sql.LevelSerializable
	//sql.LevelSnapshot
	//sql.LevelWriteCommitted
	//1. Serializable Isolation Level (Highest):
	//Ensures complete isolation; all transactions are executed serially.

	txOptions := &sql.TxOptions{
		//Isolation: sql.LevelSerializable, // Set desired isolation level
		Isolation: isolation,
		ReadOnly:  false, // Set to true for read-only transactions
	}

	// Begin a transaction with specific options
	tx, err := db.BeginTx(context.Background(), txOptions)
	if err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}
}*/
