package downloader

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MagonxESP/dropper/internal/domain"
	"github.com/MagonxESP/dropper/internal/infraestructure"
	colly2 "github.com/gocolly/colly/v2"
	"io"
	"net/http"
	"regexp"
	"time"
)

type PixivIllustrationDownloader struct{}

type PixivPreloadMetadata struct {
	Illustration map[int]interface{} `json:"illust"`
}

func getIllustrationFileNameByUrl(url string) (string, error) {
	regex, err := regexp.Compile("^.+:\\/\\/i\\.pximg\\.net\\/.+\\/([0-9]+_p[0-9]+\\.[a-z]+)$")

	if err != nil {
		return "", err
	}

	if matches := regex.FindSubmatch([]byte(url)); len(matches) > 1 {
		return string(matches[1]), nil
	}

	return "", errors.New("invalid url")
}

func NewPixivIllustrationDownloader() *PixivIllustrationDownloader {
	return &PixivIllustrationDownloader{}
}

func ScrapeIllustrationOriginalUrl(url string) (string, error) {
	var originalUrl string
	var unmarshallErr error
	collector := infraestructure.NewFirefoxCollector()

	collector.OnHTML("meta[name=\"preload-data\"]", func(element *colly2.HTMLElement) {
		var metadata PixivPreloadMetadata

		if unmarshallErr := json.Unmarshal([]byte(element.Attr("content")), &metadata); unmarshallErr != nil {
			return
		}

		for _, illustration := range metadata.Illustration {
			illustration := illustration.(map[string]interface{})
			urls := illustration["urls"].(map[string]interface{})
			originalUrl = urls["original"].(string)
			break
		}
	})

	if err := collector.Visit(url); err != nil {
		return "", err
	}

	collector.Wait()

	if unmarshallErr != nil {
		return "", unmarshallErr
	}

	if originalUrl == "" {
		return "", errors.New("missing illustration metadata")
	}

	return originalUrl, nil
}

func readUrlContent(url string) ([]byte, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Referer", "https://www.pixiv.net/")
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}

func (d *PixivIllustrationDownloader) Download(url string) (*domain.DownloadedFile, error) {
	var fileName string
	var fileContent []byte
	originalUrl, err := ScrapeIllustrationOriginalUrl(url)

	if err != nil {
		return nil, err
	}

	if fileName, err = getIllustrationFileNameByUrl(originalUrl); err != nil {
		return nil, err
	}

	if fileContent, err = readUrlContent(originalUrl); err != nil {
		return nil, err
	}

	return &domain.DownloadedFile{
		Name:         fileName,
		Content:      fileContent,
		DownloadedAt: time.Now(),
	}, nil
}

func GetPixivIllustrationDownloaderForUrl(url string) (domain.FileDownloader, error) {
	regex, err := regexp.Compile(".*\\.?pixiv\\.net")
	if err != nil {
		return nil, err
	}

	if regex.MatchString(url) {
		return NewPixivIllustrationDownloader(), nil
	}

	return nil, fmt.Errorf("the url %s is not valid for the pixiv downloader", url)
}
