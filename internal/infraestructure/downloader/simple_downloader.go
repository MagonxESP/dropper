package downloader

import (
	"errors"
	"fmt"
	"github.com/MagonxESP/dropper/internal/domain"
	"io"
	"mime"
	"net/http"
	"strings"
	"time"
)

type SimpleDownloader struct{}

func NewSimpleDownloader() *SimpleDownloader {
	return &SimpleDownloader{}
}

func extractFileName(url string, time time.Time) string {
	parts := strings.Split(url, "/")
	partsLen := len(parts)

	if partsLen > 0 && parts[partsLen-1] != "" {
		return parts[partsLen-1]
	}

	return fmt.Sprintf("untitled_%d", time.UnixMilli())
}

func hasExtension(fileName string, extensions []string) bool {
	hasExtension := false

	for _, extension := range extensions {
		if strings.HasSuffix(fileName, extension) {
			hasExtension = true
			break
		}
	}

	return hasExtension
}

func resolveFileName(url string, extensions []string, time time.Time) string {
	name := extractFileName(url, time)
	extension := ""

	if !hasExtension(name, extensions) {
		extension = extensions[len(extensions)-1]
	}

	return fmt.Sprintf("%s%s", name, extension)
}

func (s *SimpleDownloader) Download(url string) (*domain.DownloadedFile, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	contentType := response.Header.Get("Content-Type")
	extensions, err := mime.ExtensionsByType(contentType)

	if extensions == nil {
		return nil, errors.New(fmt.Sprintf("unknown file extension for the content type %s", contentType))
	}

	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &domain.DownloadedFile{
		Name:         resolveFileName(url, extensions, now),
		Content:      content,
		DownloadedAt: now,
	}, nil
}
