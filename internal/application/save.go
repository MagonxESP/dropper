package application

import "github.com/MagonxESP/dropper/internal/domain"

type RemoteFileSaver struct {
	Downloader domain.FileDownloader
	Writer     domain.DownloadedFileWriter
}

func NewRemoteFileSaver(downloader domain.FileDownloader, writer domain.DownloadedFileWriter) *RemoteFileSaver {
	return &RemoteFileSaver{
		Downloader: downloader,
		Writer:     writer,
	}
}

func (r *RemoteFileSaver) Save(source string) error {
	file, err := r.Downloader.Download(source)

	if err != nil {
		return err
	}

	return r.Writer.Write(file)
}
