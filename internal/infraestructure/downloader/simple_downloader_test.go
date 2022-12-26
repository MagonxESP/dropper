package downloader

import (
	"fmt"
	"testing"
)

func TestSimpleDownloader_Download(t *testing.T) {
	url := "https://raw.githubusercontent.com/cat-milk/Anime-Girls-Holding-Programming-Books/master/Go/Komi_holding_Go_Programming_Language.jpg"
	expectedFilename := "Komi_holding_Go_Programming_Language.jpg"

	downloader := NewSimpleDownloader()
	file, err := downloader.Download(url)

	if err != nil {
		t.Error(err)
	}

	AssertDownloadedFile(t, file, expectedFilename)
}

func TestSimpleDownloader_Download_html(t *testing.T) {
	url := "http://example.com/"

	downloader := NewSimpleDownloader()
	file, err := downloader.Download(url)

	if err != nil {
		t.Error(err)
	}

	AssertDownloadedFile(t, file, fmt.Sprintf("untitled_%d.html", file.DownloadedAt.UnixMilli()))
}
