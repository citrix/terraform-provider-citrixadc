
resource "netscaler_nsacl" "acl1" {
  aclname = "restrict"
  protocol = "TCP"
  aclaction = "DENY"
  destipval = "192.168.1.20"
  srcportval = "49-1024"
}

resource "netscaler_nsacl" "acl2" {
  aclname = "restrictudp"
  protocol = "UDP"
  aclaction = "DENY"
  destipval = "192.168.1.2"
  srcportval = "45-10024"
}

resource "netscaler_nsacl" "acl3" {
  aclname = "restricttcp2"
  protocol = "TCP"
  aclaction = "ALLOW"
  destipval = "192.168.1.40"
  srcportval = "149-1024"
}

resource "netscaler_nsacl" "acl4" {
  aclname = "restrictudp2"
  protocol = "UDP"
  aclaction = "ALLOW"
  destipval = "192.168.10.2"
  srcportval = "490-1024"
}

resource "netscaler_nsacl" "acl5" {
  aclname = "restrictvlan"
  aclaction = "DENY"
  vlan = "2000"
}

resource "netscaler_nsacls" "allacls" {
}
