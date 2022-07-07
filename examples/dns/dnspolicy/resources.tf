resource "citrixadc_dnspolicy" "dnspolicy" {
	name = "policy_A"
	rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
	drop = "YES"
  }