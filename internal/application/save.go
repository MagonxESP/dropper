package application

import "github.com/MagonxESP/dropper/internal/domain"

type RemoteFileSaver struct {
	Downloader domain.RemoteFileDownloader
	Writer     domain.RemoteFileWriter
}

func NewRemoteFileSaver(downloader domain.RemoteFileDownloader, writer domain.RemoteFileWriter) *RemoteFileSaver {
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
