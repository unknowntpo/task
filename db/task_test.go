package db

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/boltdb/bolt"
)

func TestInit(t *testing.T) {
	tempFile, removeFunc := createTempFile(t)
	defer removeFunc()

	// call Init(tempFilePath)
	// How to get tempFile's path?
	err := Init(tempFile.Name())
	if err != nil {
		t.Errorf("Fail to init database")
	}

	err = Db.View(func(tx *bolt.Tx) error {
		var root *bolt.Bucket
		if root = tx.Bucket(task); root == nil {
			t.Errorf("Fail to get root bucket")
		}

		return nil
	})
	if err != nil {
		t.Errorf("Something wrong with View transaction: %v", err)
	}
}

func createTempFile(t *testing.T) (tempFile *os.File, removeFunc func()) {
	t.Helper()
	tempFile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Errorf("could not create temp file: %v", err)
	}

	removeFunc = func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	return
}

func TestCreateTask(t *testing.T) {
	tempFile, removeFunc := createTempFile(t)
	defer removeFunc()

	err := Init(tempFile.Name())
	if err != nil {
		t.Errorf("Fail to init database")
	}

	wantContent := "test123"
	err = CreateTask(Db, wantContent)
	if err != nil {
		t.Errorf("Fail to Create test: %v", err)
	}

	err = Db.View(func(tx *bolt.Tx) error {
		root := tx.Bucket(task)
		if root == nil {
			t.Errorf("Fail to get root bucket")
		}
		c := root.Cursor()

		// actually there's only one k-v in our root bucket now
		for k, v := c.First(); k != nil; k, v = c.Next() {
			gotContent := string(v)
			if gotContent != wantContent {
				t.Errorf("Content of task not correct, got %v, want %v", gotContent, wantContent)
			}
		}
		return nil
	})
	if err != nil {
		t.Errorf("Fail to open a View transaction: %v", err)
	}
}
