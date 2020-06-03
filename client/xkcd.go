package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/vald-phoenix/xkcd/model"
)

// ComicNumber is the number of the Comic
type ComicNumber int

const (
	// BaseURL of xkcd
	BaseURL string = "https://xkcd.com"

	// DefaultClientTimeout is time to wait before cancelling the request
	DefaultClientTimeout time.Duration = 30 * time.Second

	// LatestComic is the latest comic number according to the xkcd API
	LatestComic ComicNumber = 0
)

// XKCDClient is the client for XKCD
type XKCDClient struct {
	client  *http.Client
	baseURL string
}

// NewXKCDClient creates a new XKCDClient
func NewXKCDClient() *XKCDClient {
	return &XKCDClient{
		client:  &http.Client{Timeout: DefaultClientTimeout},
		baseURL: BaseURL,
	}
}

// SetTimeout overrides the default ClientTimeout
func (hc *XKCDClient) SeetTimeout(d time.Duration) {
	hc.client.Timeout = d
}

// Fetch retrieves the comic as per provided comic number
func (hc *XKCDClient) Fetch(n ComicNumber, save bool) (model.Comic, error) {
	resp, err := hc.client.Get(hc.buildURL(n))
	if err != nil {
		return model.Comic{}, err
	}
	defer resp.Body.Close()

	var xkcdResponse model.XkcdResponse
	if err := json.NewDecoder(resp.Body).Decode(&xkcdResponse); err != nil {
		return model.Comic{}, err
	}

	if save {
		if err := hc.SaveToDisk(xkcdResponse.Img, "."); err != nil {
			fmt.Println("Failed to save image!")
		}
	}
	return xkcdResponse.Comic(), nil
}

// SaveToDisk downloads and saves the comic locally
func (hc *XKCDClient) SaveToDisk(url, savePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	absSavePath, _ := filepath.Abs(savePath)
	filePath := fmt.Sprintf("%s/%s", absSavePath, path.Base(url))

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// buildURL builds an URL to fetch a comic
func (hc *XKCDClient) buildURL(n ComicNumber) string {
	var finalURL string
	if n == LatestComic {
		finalURL = fmt.Sprintf("%s/info.0.json", hc.baseURL)
	} else {
		finalURL = fmt.Sprintf("%s/%d/info.0.json", hc.baseURL, n)
	}
	return finalURL
}
