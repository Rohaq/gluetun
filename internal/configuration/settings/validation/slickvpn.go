package validation

import "github.com/qdm12/gluetun/internal/models"

func SlickVPNRegionChoices(servers []models.SlickVPNServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Region
	}
	return makeUnique(choices)
}

func SlickVPNCountryChoices(servers []models.SlickVPNServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Country
	}
	return makeUnique(choices)
}

func SlickVPNCityChoices(servers []models.SlickVPNServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].City
	}
	return makeUnique(choices)
}

func SlickVPNHostnameChoices(servers []models.SlickVPNServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Hostname
	}
	return makeUnique(choices)
}
