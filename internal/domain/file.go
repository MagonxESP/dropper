package domain

type DownloadedFile struct {
	Name    string
	Content []byte
}

type RemoteFileDownloader interface {
	Download(url string) (*DownloadedFile, error)
}

type RemoteFileWriter interface {
	Write(file *DownloadedFile) error
}
