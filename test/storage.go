package test

import (
	"os"

	"github.com/0xDezzy/fyne"
	"github.com/0xDezzy/fyne/internal"
	"github.com/0xDezzy/fyne/storage"
)

type testStorage struct {
	*internal.Docs
}

func (s *testStorage) RootURI() fyne.URI {
	return storage.NewFileURI(os.TempDir())
}

func (s *testStorage) docRootURI() (fyne.URI, error) {
	return storage.Child(s.RootURI(), "Documents")
}
