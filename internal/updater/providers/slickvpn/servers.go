/*
Package slickvpn contains code to obtain the server information for the SlickVPN provider.
Note: As far as I'm aware, SlickVPN only provide UDP config files.
If we want to implement TCP, we would need to perhaps take the UDP config file and adjust
its config to match SlickVPN's TCP settings (which I believe that they do support)
*/
package slickvpn

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/updater/resolver"
)

var ErrNotEnoughServers = errors.New("not enough servers found")

func GetServers(ctx context.Context, client *http.Client,
	presolver resolver.Parallel, minServers int) (
	servers []models.SlickVPNServer, warnings []string, err error) {
	udpHostToURL, err := getAllHostToURL(ctx, client)
	if err != nil {
		return nil, nil, err
	}

	hosts := getUniqueHosts(udpHostToURL)

	if len(hosts) < minServers {
		return nil, nil, fmt.Errorf("%w: %d and expected at least %d",
			ErrNotEnoughServers, len(hosts), minServers)
	}

	hostToIPs, warnings, err := resolveHosts(ctx, presolver, hosts, minServers)
	if err != nil {
		return nil, warnings, err
	}

	servers = make([]models.SlickVPNServer, 0, len(hostToIPs))
	for host, IPs := range hostToIPs {
		udpURL, udp := udpHostToURL[host]

		// These two are only used to extract the country, region and city.
		var url, protocol string
		if udp {
			url = udpURL
			protocol = "UDP"
		}
		country, region, city := parseOpenvpnURL(url, protocol)

		server := models.SlickVPNServer{
			Country:  country,
			Region:   region,
			City:     city,
			Hostname: host,
			IPs:      IPs,
			UDP:      udp,
		}
		servers = append(servers, server)
	}

	sortServers(servers)

	return servers, warnings, nil
}
