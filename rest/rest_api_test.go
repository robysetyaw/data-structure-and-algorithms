package rest

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestRestV2(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "data.json")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpFile.Close()

	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}

	data, err := json.Marshal(users)
	if err != nil {
		t.Fatal(err)
	}

	if err := ioutil.WriteFile(tmpFile.Name(), data, 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("TestLoadDataSuccess", func(t *testing.T) {
		loadedUsers, err := restV2LoadData(tmpFile.Name())
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(loadedUsers) != len(users) {
			t.Fatalf("Expected %d users, got %d", len(users), len(loadedUsers))
		}
		for i, user := range loadedUsers {
			if user != users[i] {
				t.Errorf("Expected user %v, got %v", users[i], user)
			}
		}
	})

	// Tes jika file tidak ditemukan
	t.Run("TestLoadDataFileNotFound", func(t *testing.T) {
		_, err := restV2LoadData("invalid/path/data.json")
		if err == nil {
			t.Fatal("Expected error, got none")
		}
	})

	t.Run("TestLoadDataInvalidJSON", func(t *testing.T) {
		// Membuat file sementara dengan data JSON invalid
		invalidData := []byte("invalid json data")
		tmpFileInvalid, err := ioutil.TempFile("", "invalid_data.json")
		if err != nil {
			t.Fatal(err)
		}
		defer tmpFileInvalid.Close()

		if err := ioutil.WriteFile(tmpFileInvalid.Name(), invalidData, 0644); err != nil {
			t.Fatal(err)
		}

		_, err = restV2LoadData(tmpFileInvalid.Name())
		if err == nil {
			t.Fatal("Expected error, got none")
		}
	})
}
