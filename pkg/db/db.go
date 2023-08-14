package db

import (
	"log"
	"os"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

func StartDB() (*pg.DB, error) {
	var opts *pg.Options
	var err error

	//set up options. if PROD use env variabels otherwise use manual variables
	if os.Getenv("ENV") == "PROD" {
		opts, err = pg.ParseURL(os.Getenv("DATABASE_URL"))
		if err != nil {
			return nil, err
		}
	} else {
		opts = &pg.Options{
			Addr:     "db:5432",
			User:     "postgres",
			Password: "admin",
		}
	}

	//connect to the DB and crab the migrations from the directory
	db := pg.Connect(opts)
	collection := migrations.NewCollection()
	err = collection.DiscoverSQLMigrations("migrations")
	if err != nil {
		return nil, err
	}

	//run the migrations to intialize the DB if needed
	_, _, err = collection.Run(db, "init")
	if err != nil {
		return nil, err
	}

	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
		return nil, err
	}

	if newVersion != oldVersion {
		log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		log.Printf("version is %d\n", oldVersion)
	}

	return db, err
}
