package nget

import "fmt"

// HTTPDownloader ...
type HTTPDownloader struct {
	url string
}

// NewHTTPDownloader ...
func NewHTTPDownloader(url string) *HTTPDownloader {
	return &HTTPDownloader{
		url,
	}
}

func (h *HTTPDownloader) Do() {
	fmt.Println("Do")
}
