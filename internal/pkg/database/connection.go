package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Roach holds the connection pool to the database - created by a configuration
// object (`Config`).
type Roach struct {
	// Db holds a sql.DB pointer that represents a pool of zero or more
	// underlying connections - safe for concurrent use by multiple
	// goroutines -, with freeing/creation of new connections all managed
	// by `sql/database` package.
	Db  *sql.DB
	cfg Config
}

// Config holds the configuration used for instantiating a new Roach.
type Config struct {
	// Address that locates our postgres instance
	Host string
	// Port to connect to
	Port string
	// User that has access to the database
	User string
	// Password so that the user can login
	Password string
	// Database to connect to (must have been created priorly)
	Database string
}

var (
	once sync.Once

	instance Roach
)

// New returns a Roach with the sql.DB set with the postgres
// DB connection string in the configuration
func NewConnection(cfg Config) (roach Roach, err error) {
	once.Do(func() {
		if cfg.Host == "" || cfg.Port == "" || cfg.User == "" ||
			cfg.Password == "" || cfg.Database == "" {
			err = errors.Errorf(
				"All fields must be set (%s)",
				spew.Sdump(cfg))
			return
		}
		roach.cfg = cfg
		connection := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
		db, err := sql.Open("postgres", connection)
		if err != nil {
			err = errors.Wrapf(err,
				"Couldn't open connection to postgre database (%s)",
				spew.Sdump(cfg))
			return
		}

		if err = db.Ping(); err != nil {
			err = errors.Wrapf(err,
				"Couldn't ping postgre database (%s)",
				spew.Sdump(cfg))
			return
		}
		log.Print("Connect DB Success")
		// Config connection pool
		db.SetMaxOpenConns(70)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(5 * time.Minute)
		roach.Db = db
		instance = roach
		//go CheckOpenConnection(db)
	})

	return instance, nil
}

func CheckOpenConnection(db *sql.DB) {
	println(db.Stats().OpenConnections)
	time.Sleep(time.Second * 2)
	CheckOpenConnection(db)
}

// Todo: chưa check dc instance connection nil
func GetConnection() Roach {
	return instance
}

// Close performs the release of any resources that
// `sql/database` DB pool created. This is usually meant
// to be used in the exitting of a program or `panic`ing.
func (r *Roach) Close() (err error) {
	if r.Db == nil {
		return
	}

	if err = r.Db.Close(); err != nil {
		err = errors.Wrapf(err,
			"Errored closing database connection",
			spew.Sdump(r.cfg))
	}
	return
}

