package slickvpn

func getUniqueHosts(udpHostToURL map[string]string) (
	hosts []string) {
	uniqueHosts := make(map[string]struct{}, len(udpHostToURL))
	for host := range udpHostToURL {
		uniqueHosts[host] = struct{}{}
	}

	hosts = make([]string, 0, len(uniqueHosts))
	for host := range uniqueHosts {
		hosts = append(hosts, host)
	}

	return hosts
}
