// Author: pat@patfairbank.com (Patrick Fairbank)
// Modified for Isotope Robotics by: Ethen Brandenburg
//
// Functions for manipulating the per-event Bolt datastore.

package model

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.etcd.io/bbolt"
)

const backupsDir = "db/backups"

var BaseDir = "."

type Database struct {
	Path               string
	bolt               *bbolt.DB
	userSessionTable   *table[UserSession]
	eventSettingsTable *table[EventSettings]
	teamTable          *table[Team]
}

// Opens the Bolt database at the given path, creating it if it doesn't exist.
func OpenDatabase(filename string) (*Database, error) {
	database := Database{Path: filename}
	var err error
	database.bolt, err = bbolt.Open(database.Path, 0644, &bbolt.Options{NoSync: true, Timeout: time.Second})
	if err != nil {
		return nil, err
	}

	// Register tables.
	//TODO

	return &database, nil
}

func (database *Database) Close() error {
	return database.bolt.Close()
}

// Creates a copy of the current database and saves it to the backups directory.
func (database *Database) Backup(eventName, reason string) error {
	backupsPath := filepath.Join(BaseDir, backupsDir)
	err := os.MkdirAll(backupsPath, 0755)
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%s/%s_%s_%s.db", backupsPath, strings.Replace(eventName, " ", "_", -1),
		time.Now().Format("20060102150405"), reason)

	dest, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dest.Close()

	if err = database.WriteBackup(dest); err != nil {
		return err
	}
	return nil
}

// Takes a snapshot of Bolt database and writes it to the given writer.
func (database *Database) WriteBackup(writer io.Writer) error {
	return database.bolt.View(func(tx *bbolt.Tx) error {
		_, err := tx.WriteTo(writer)
		return err
	})
}
