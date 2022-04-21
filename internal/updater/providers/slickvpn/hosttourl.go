package slickvpn

import (
	"context"
	"net/http"

	"github.com/qdm12/gluetun/internal/updater/openvpn"
)

func getAllHostToURL(ctx context.Context, client *http.Client) (
	udpHostToURL map[string]string, err error) {
	udpHostToURL, err = getHostToURL(ctx, client)
	if err != nil {
		return nil, err
	}

	return udpHostToURL, nil
}

func getHostToURL(ctx context.Context, client *http.Client) (
	hostToURL map[string]string, err error) {
	const baseURL = "https://www.slickvpn.com/locations/"

	urls, err := fetchIndex(ctx, client, baseURL)
	if err != nil {
		return nil, err
	}

	return openvpn.FetchMultiFiles(ctx, client, urls)
}
