package vpn

type Vpnalwaysonprofile struct {
	Clientcontrol             string `json:"clientcontrol,omitempty"`
	Locationbasedvpn          string `json:"locationbasedvpn,omitempty"`
	Name                      string `json:"name,omitempty"`
	Networkaccessonvpnfailure string `json:"networkaccessonvpnfailure,omitempty"`
}
