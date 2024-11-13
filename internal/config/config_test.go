package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"testing"
)

func createTempFile(t testing.TB) *os.File {
	t.Helper()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("Could not find home dir")
	}
	t.TempDir()
	tempFile, err := os.CreateTemp(homeDir, ".gatorconfig.test.json")
	if err != nil {
		t.Fatal("Could not open temporary file")
	}
	configFilename = path.Base(tempFile.Name())

	t.Cleanup(func() {
		tempFile.Close()
		if err := os.Remove(tempFile.Name()); err != nil {
			fmt.Printf("Error removing file, please remove it manually at: %s", tempFile.Name())
		}
	})
	return tempFile
}

func TestReadConfigFull(t *testing.T) {
	tempFile := createTempFile(t)
	fakeContents := []byte(`{"db_url":"postgres://example","current_user_name":"some_username"}`)
	tempFile.Write(fakeContents)

	got, err := Read()
	if err != nil {
		t.Fatal(err)
	}
	if got.CurrentUserName != "some_username" {
		t.Errorf("cfg.CurrentUserName = %v, want 'some_username'", got.CurrentUserName)
	}
	if got.DBURL != "postgres://example" {
		t.Errorf("cfg.DBURL = %v, want 'postgres://example'", got.DBURL)
	}
}

func TestReadConfigOnlyUrl(t *testing.T) {
	tempFile := createTempFile(t)
	fakeContents := []byte(`{"db_url":"postgres://example"}`)
	tempFile.Write(fakeContents)

	got, err := Read()
	if err != nil {
		t.Fatal(err)
	}
	if got.CurrentUserName != "" {
		t.Errorf("cfg.CurrentUserName = %v, want empty string", got.CurrentUserName)
	}
	if got.DBURL != "postgres://example" {
		t.Errorf("cfg.DBURL = %v, want 'postgres://example'", got.DBURL)
	}
}

func TestReadConfigUsername(t *testing.T) {
	tempFile := createTempFile(t)
	fakeContents := []byte(`{"current_user_name":"username"}`)
	tempFile.Write(fakeContents)

	got, err := Read()
	if err != nil {
		t.Fatal(err)
	}
	if got.CurrentUserName != "username" {
		t.Errorf("cfg.CurrentUserName = %v, want 'username'", got.CurrentUserName)
	}
	if got.DBURL != "" {
		t.Errorf("cfg.DBURL = %v, want empty string", got.DBURL)
	}
}

func TestSetUser(t *testing.T) {
	tempFile := createTempFile(t)
	fakeContents := []byte(`{"db_url": "sqlite://file.db"}`)
	tempFile.Write(fakeContents)
	cfg, err := Read()
	if err != nil {
		t.Fatal(err)
	}
	cfg.SetUser("cool_user123")
	expectedContent := []byte(`{"db_url":"sqlite://file.db","current_user_name":"cool_user123"}`)
	tempFile.Seek(0, 0) // Go to the start of the file before reading all
	contents, err := io.ReadAll(tempFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(contents, expectedContent) {
		t.Errorf("file contents: %v, expected %v", string(contents), string(expectedContent))
	}
	if cfg.CurrentUserName != "cool_user123" {
		t.Errorf("cfg.CurrentUserName: %v, expected 'cool_user123'", cfg.CurrentUserName)
	}
}
