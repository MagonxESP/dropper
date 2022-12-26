package domain

type DownloadedFile struct {
	Name    string
	Content []byte
}

type FileDownloader interface {
	Download(url string) (*DownloadedFile, error)
}

type DownloadedFileWriter interface {
	Write(file *DownloadedFile) error
}
