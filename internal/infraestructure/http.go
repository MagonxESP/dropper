package infraestructure

import (
	"io"
	"net/http"
)

func ReadFileUrlContent(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}
