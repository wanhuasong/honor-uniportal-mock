package db

import (
	"github.com/juju/errors"
	bolt "github.com/k8scat/bbolt"
	"log"
	"os"
	"path/filepath"
)

const DefaultDBFile = "uniportal.db"

var (
	DefaultBucket = []byte("default")
	DefaultStore  *Store
)

type Store struct {
	db *bolt.DB
}

func InitDefaultStore() error {
	dbFile := filepath.Join(os.TempDir(), DefaultDBFile)
	log.Printf("db file: %s", dbFile)
	db, err := bolt.Open(dbFile, 0666, nil)
	if err != nil {
		err = errors.Trace(err)
	}
	DefaultStore = &Store{db: db}
	return nil
}

func (s *Store) Get(key string) string {
	var v []byte
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(DefaultBucket)
		v = b.Get([]byte(key))
		return nil
	})
	if err != nil {
		log.Printf("%+v", errors.Trace(err))
	}
	return string(v)
}

func (s *Store) Set(key, value string) {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(DefaultBucket)
		if err != nil {
			return errors.Trace(err)
		}
		err = b.Put([]byte(key), []byte(value))
		return errors.Trace(err)
	})
	if err != nil {
		log.Printf("%+v", errors.Trace(err))
	}
}

func (s *Store) Close() error {
	return s.db.Close()
}
