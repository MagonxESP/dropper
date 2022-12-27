package writer

import (
	"github.com/MagonxESP/dropper/internal/domain"
	"os"
	"path"
)

type FileSystemWriter struct {
	Destination string
}

func NewFileSystemWriter(spec *domain.FileSystemBucketSpec) *FileSystemWriter {
	return &FileSystemWriter{
		Destination: spec.DirPath,
	}
}

func (f *FileSystemWriter) Write(file *domain.DownloadedFile) error {
	return os.WriteFile(path.Join(f.Destination, file.Name), file.Content, 0755)
}
