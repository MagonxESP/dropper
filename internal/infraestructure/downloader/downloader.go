package downloader

import (
	"github.com/MagonxESP/dropper/internal/domain"
	"log"
)

type Resolver func(url string) (domain.FileDownloader, error)
type HttpFileDownloader struct{}

var resolvers = []Resolver{
	GetPixivIllustrationDownloaderForUrl,
}

func NewHttpFileDownloader() *HttpFileDownloader {
	return &HttpFileDownloader{}
}

func GetDownloaderByUrl(url string) (domain.FileDownloader, error) {
	for _, resolver := range resolvers {
		downloader, err := resolver(url)
		if err != nil {
			log.Println(err)
		}

		if downloader != nil {
			return downloader, nil
		}
	}

	return NewSimpleDownloader(), nil
}

func (h *HttpFileDownloader) Download(url string) (*domain.DownloadedFile, error) {
	downloader, err := GetDownloaderByUrl(url)

	if err != nil {
		return nil, err
	}

	return downloader.Download(url)
}
