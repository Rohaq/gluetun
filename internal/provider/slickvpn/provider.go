package slickvpn

import (
	"math/rand"

	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/utils"
)

type SlickVPN struct {
	servers    []models.SlickVPNServer
	randSource rand.Source
	utils.NoPortForwarder
}

func New(servers []models.SlickVPNServer, randSource rand.Source) *SlickVPN {
	return &SlickVPN{
		servers:         servers,
		randSource:      randSource,
		NoPortForwarder: utils.NewNoPortForwarding(constants.SlickVPN),
	}
}
