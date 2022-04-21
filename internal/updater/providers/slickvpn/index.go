package slickvpn

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func fetchIndex(ctx context.Context, client *http.Client, indexURL string) (openvpnURLs []string, err error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, indexURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	doc.Find("a[href$='.ovpn']").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "OpenVPN" {
			link, _ := s.Attr("href")

			if targetURLExists(ctx, client, link) {
				openvpnURLs = append(openvpnURLs, link)
			}
		}
	})

	return openvpnURLs, nil
}

// This is used to determine whether the URLs provided actually return a file.
// SlickVPN really need to update their links...
func targetURLExists(ctx context.Context, client *http.Client, link string) bool {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return false
	}

	response, err := client.Do(request)
	response.Body.Close()
	if err != nil || response.StatusCode == http.StatusNotFound {
		return false
	}

	return response.StatusCode >= 200 && response.StatusCode < 400
}
