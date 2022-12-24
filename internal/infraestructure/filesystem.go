package infraestructure

import (
	"github.com/MagonxESP/dropper/internal/domain"
	"os"
	"path"
)

type FileSystemWriter struct {
	Destination string
}

func NewFileSystemWriter(destination string) *FileSystemWriter {
	return &FileSystemWriter{
		Destination: destination,
	}
}

func (f *FileSystemWriter) Write(file *domain.DownloadedFile) error {
	return os.WriteFile(path.Join(f.Destination, file.Name), file.Content, 0755)
}
