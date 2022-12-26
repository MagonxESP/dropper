package downloader

import (
	"errors"
	"fmt"
	"github.com/MagonxESP/dropper/internal/domain"
	"regexp"
)

type HttpFileDownloader struct{}

var downloaders = map[string]domain.FileDownloader{
	".*\\.?pixiv\\.net": NewPixivIllustrationDownloader(),
	".+":                NewSimpleDownloader(),
}

func NewHttpFileDownloader() *HttpFileDownloader {
	return &HttpFileDownloader{}
}

func GetDownloaderByUrl(url string) (domain.FileDownloader, error) {
	for pattern, downloader := range downloaders {
		regex, err := regexp.Compile(pattern)

		if err != nil {
			return nil, err
		}

		if regex.Match([]byte(url)) {
			return downloader, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("missing downloader for url %s", url))
}

func (h *HttpFileDownloader) Download(url string) (*domain.DownloadedFile, error) {
	downloader, err := GetDownloaderByUrl(url)

	if err != nil {
		return nil, err
	}

	return downloader.Download(url)
}
