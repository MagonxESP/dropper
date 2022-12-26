package pixiv

import "testing"

func TestIllustrationOriginalUrl(t *testing.T) {
	thumbnailUrl, err := ScrapeIllustrationOriginalUrl("https://www.pixiv.net/en/artworks/99895500")
	expected := "https://i.pximg.net/img-original/img/2022/07/22/00/00/34/99895500_p0.jpg"

	if err != nil {
		t.Error(err)
	}

	if thumbnailUrl != expected {
		t.Errorf("expected %s got %s", expected, thumbnailUrl)
	}
}

func TestPixivIllustrationDownloader_Download(t *testing.T) {
	source := "https://www.pixiv.net/en/artworks/99895500"
	downloader := NewPixivIllustrationDownloader()
	file, err := downloader.Download(source)

	if err != nil {
		t.Fatal(err)
	}

	if file.Content == nil {
		t.Fatalf("The downloaded content from %s is empty", source)
	}

	if file.Name != "99895500_p0.jpg" {
		t.Fatalf("The filename retrieved is invalid")
	}
}
