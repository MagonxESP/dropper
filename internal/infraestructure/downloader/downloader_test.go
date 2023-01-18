package downloader

import (
	"testing"
)

func TestGetDownloaderByUrl_Pixiv(t *testing.T) {
	url := "https://www.pixiv.net/en/artworks/99895500"
	downloader, err := GetDownloaderByUrl(url)

	if err != nil {
		t.Error(err)
	}

	if downloader == nil {
		t.Errorf("missing downloader for %s", url)
	}

	if _, ok := downloader.(*PixivIllustrationDownloader); !ok {
		t.Errorf("mismatch downloader for %s", url)
	}
}

func TestGetDownloaderByUrl_Example(t *testing.T) {
	url := "http://example.com"
	downloader, err := GetDownloaderByUrl(url)

	if err != nil {
		t.Error(err)
	}

	if downloader == nil {
		t.Errorf("missing downloader for %s", url)
	}

	if _, ok := downloader.(*SimpleDownloader); !ok {
		t.Errorf("mismatch downloader for %s", url)
	}
}
