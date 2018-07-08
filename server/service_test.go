package server

import (
	"testing"

	"github.com/moov-io/ach"
)

// test mocks are in mock_test.go

// CreateFile tests
func TestCreateFile(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateFile(mockFileHeader())
	if id != "12345" {
		t.Errorf("expected %s received %s w/ error %s", "12345", id, err)
	}
}
func TestCreateFileIDExists(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateFile(ach.FileHeader{ID: "98765"})
	if err != ErrAlreadyExists {
		t.Errorf("expected %s received %s w/ error %s", "ErrAlreadyExists", id, err)
	}
}

func TestCreateFileNoID(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateFile(ach.NewFileHeader())
	if len(id) < 3 {
		t.Errorf("expected %s received %s w/ error %s", "NextID", id, err)
	}
}

// Service.GetFile tests

func TestGetFile(t *testing.T) {
	s := mockServiceInMemory()
	f, err := s.GetFile("98765")
	if err != nil {
		t.Errorf("expected %s received %s w/ error %s", "98765", f.ID, err)
	}
}

func TestGetFileNotFound(t *testing.T) {
	s := mockServiceInMemory()
	f, err := s.GetFile("12345")
	if err != ErrNotFound {
		t.Errorf("expected %s received %s w/ error %s", "ErrNotFound", f.ID, err)
	}
}

// Service.GetFiles tests

func TestGetFiles(t *testing.T) {
	s := mockServiceInMemory()
	files := s.GetFiles()
	if len(files) != 1 {
		t.Errorf("expected %s received %v", "1", len(files))
	}
}

// Service.DeleteFile tests

func TestDeleteFile(t *testing.T) {
	s := mockServiceInMemory()
	err := s.DeleteFile("98765")
	if err != nil {
		t.Errorf("expected %s received %s", "nil", err)
	}
	_, err = s.GetFile("98765")
	if err != ErrNotFound {
		t.Errorf("expected %s received %s", "ErrNotFound", err)
	}
}

// Service.CreateBatch tests

// TestCreateBatch tests creating a new batch when file.ID exists and batch.id does not exist
func TestCreateBatch(t *testing.T) {
	s := mockServiceInMemory()
	bh := mockBatchHeaderWeb()
	bh.ID = "11111"
	id, err := s.CreateBatch("98765", *bh)
	if id != "11111" {
		t.Errorf("expected %s received %s w/ error %s", "11111", id, err)
	}
}

// TestCreateBatchIDExists Create a new batch with batch.id already present. Should fail.
func TestCreateBatchIDExists(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateBatch("98765", *mockBatchHeaderWeb())
	if err != ErrAlreadyExists {
		t.Errorf("expected %s received %s w/ error %s", "ErrAlreadyExists", id, err)
	}
}

// TestCreateBatchFileIDExits create a batch when the file.id does not exist. Should fail.
func TestCreateBatchFileIDExits(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateBatch("55555", *mockBatchHeaderWeb())
	if err != ErrNotFound {
		t.Errorf("expected %s received %s w/ error %s", "ErrNotFound", id, err)
	}

}

// TestCreateBatchIDBank create a new batch when the batch.id is nil but file.id is valid. Should generate batch.id and save.
func TestCreateBatchIDBlank(t *testing.T) {
	s := mockServiceInMemory()
	bh := mockBatchHeaderWeb()
	bh.ID = ""
	id, err := s.CreateBatch("98765", *bh)
	if len(id) < 3 {
		t.Errorf("expected %s received %s w/ error %s", "NextID", id, err)
	}
}

// Service.GetBatch

// TestGetBatch return a batch for the existing file.id and batch.id
func TestGetBatch(t *testing.T) {
	s := mockServiceInMemory()
	b, err := s.GetBatch("98765", "54321")
	if b.ID() != "54321" {
		t.Errorf("expected %s received %s w/ error %s", "54321", b.ID(), err)
	}
}

// TestGetBatchNotFound return a failure if the batch.id is not found
func TestGetBatchNotFound(t *testing.T) {
	s := mockServiceInMemory()
	b, err := s.GetBatch("98765", "55555")
	if err != ErrNotFound {
		t.Errorf("expected %s received %s w/ error %s", "ErrNotFound", b.ID(), err)
	}
}

// Service.GetBatches

// TestGetBatches return a list of batches for the supplied file.id
func TestGetBatches(t *testing.T) {

}

// Service.DeleteBatch

// TestDeleteBatch removes a batch with existing file and batch id.
func TestDeleteBatch(t *testing.T) {

}
