package db

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var Db *bolt.DB

// Can we declare byte slice constant ?
// Export or not ?
var task = []byte("Task")

func Init(filePath string) error {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	var err error
	Db, err = bolt.Open(filePath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new top-level bucket
	err = Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(task)
		if err != nil {
			log.Fatalf("Failed to create root bucket: %v", err)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func CreateTask(db *bolt.DB, content string) error {
	byteContent := []byte(content)
	return db.Update(func(tx *bolt.Tx) error {
		var taskId uint64
		b := tx.Bucket(task)
		taskId, _ = b.NextSequence()
		return b.Put(itob(taskId), byteContent)
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// btoi returns an uint64 value from b
func btoi(b []byte) uint64 { return binary.BigEndian.Uint64(b) }
