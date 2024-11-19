package database

import (
	"TesisAclifim/pkg/util"
	"database/sql"
	_ "github.com/lib/pq"

	"log"
	"os"
	"testing"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(conn)
	os.Exit(m.Run())
}
