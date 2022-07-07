resource "citrixadc_nsacls" "allacls" {
  aclsname = "foo"
  acl {
    aclname    = "restrict"
    protocol   = "TCP"
    aclaction  = "DENY"
    destipval  = "192.168.1.20"
    srcportval = "49-1024"
  }

  acl {
    aclname    = "restrictudp"
    protocol   = "UDP"
    aclaction  = "DENY"
    destipval  = "192.168.1.2"
    srcportval = "45-10024"
  }

  acl {
    aclname    = "restricttcp2"
    protocol   = "TCP"
    aclaction  = "DENY"
    destipval  = "192.168.199.52"
    srcportval = "149-1524"
  }

  acl {
    aclname    = "restrictudp2"
    protocol   = "UDP"
    aclaction  = "DENY"
    destipval  = "192.168.45.55"
    srcportval = "490-1024"
    priority   = "100"
  }

  acl {
    aclname   = "restrictvlan"
    aclaction = "DENY"
    vlan      = "2000"
  }
}
