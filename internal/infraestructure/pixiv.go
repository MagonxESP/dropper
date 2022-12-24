package infraestructure

import (
	"encoding/binary"
	"errors"
	"github.com/MagonxESP/dropper/internal/domain"
	"github.com/everpcpc/pixiv"
	"os"
	"regexp"
)

type PixivApiClient struct {
	Account *pixiv.Account
	Client  *pixiv.AppPixivAPI
}

type PixivIllustrationDownloader struct{}

// Hay 2 opciones:
// 1: Intentando iniciar sesion (tienen recaptcha)
// 2: Sin iniciar sesion resolviendo la url del fichero original. ¿Como?
//	Pues muy sencillo, extraemos el id de la url, buscamos la url de la miniatura que se ve y de esta url sustituimos
// 	al final el nombre del fichero por {id del artwork}_p0.{extension que tiene la url de la miniatura}.
// 	Y para descargar el contenido hay que enviar la cabecera Referer: https://www.pixiv.net/

func NewPixivApiClient() (*PixivApiClient, error) {
	account, err := pixiv.Login(os.Getenv("PIXIV_USERNAME"), os.Getenv("PIXIV_PASSWORD"))

	if err != nil {
		return nil, err
	}

	return &PixivApiClient{
		Account: account,
		Client:  pixiv.NewApp(),
	}, nil
}

func getIllustrationIdByUrl(url string) (uint64, error) {
	regex, err := regexp.Compile("^.+:\\/\\/.*pixiv\\.net\\/[a-z\\/]+([0-9]+)$")

	if err != nil {
		return 0, err
	}

	if matches := regex.FindSubmatch([]byte(url)); len(matches) > 1 {
		return binary.LittleEndian.Uint64(matches[1]), nil
	}

	return 0, errors.New("invalid url")
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

func (d *PixivIllustrationDownloader) Download(url string) (*domain.DownloadedFile, error) {
	var client *PixivApiClient
	var err error
	var fileName string
	var fileContent []byte
	var id uint64
	var illustration *pixiv.Illust

	if client, err = NewPixivApiClient(); err != nil {
		return nil, err
	}

	if id, err = getIllustrationIdByUrl(url); err != nil {
		return nil, err
	}

	if illustration, err = client.Client.IllustDetail(id); err != nil {
		return nil, err
	}

	if fileName, err = getIllustrationFileNameByUrl(illustration.Images.Original); err != nil {
		return nil, err
	}

	if fileContent, err = ReadFileUrlContent(illustration.Images.Original); err != nil {
		return nil, err
	}

	return &domain.DownloadedFile{
		Name:    fileName,
		Content: fileContent,
	}, nil
}
