package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/unknowntpo/task/db"
)

func TestNewAddCmd(t *testing.T) {
	err := initDB(t)
	if err != nil {
		t.Errorf("Fail to init DB: %v", err)
	}

	// change to ./task -h
	buf, err := initTestCmd(t, []string{"add", "test123"})
	if err != nil {
		t.Errorf("Fail to init cmds: %v", err)
	}

	got, err := ioutil.ReadAll(buf)
	want := `Add "test123" to your task list.
`
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("expected \"%s\" got \"%s\"", want, string(got))
	}

}

func initTestCmd(t *testing.T, args []string) (*bytes.Buffer, error) {
	t.Helper()
	rootCmd, err := Init()
	if err != nil {
		return nil, fmt.Errorf("Fail to init test cmd: %v", err)
	}

	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs(args)
	rootCmd.Execute()
	return b, nil
}
func initDB(t *testing.T) error {
	t.Helper()
	tempFile, removeFunc := createTempFile(t)
	defer removeFunc()

	err := db.Init(tempFile.Name())
	if err != nil {
		return fmt.Errorf("Fail to init database")
	}
	return nil
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
