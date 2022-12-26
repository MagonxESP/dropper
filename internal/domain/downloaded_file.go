package domain

import "time"

type DownloadedFile struct {
	Name         string
	Content      []byte
	DownloadedAt time.Time
}

type FileDownloader interface {
	Download(url string) (*DownloadedFile, error)
}

type DownloadedFileWriter interface {
	Write(file *DownloadedFile) error
}
