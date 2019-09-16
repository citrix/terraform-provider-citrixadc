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
    aclname   = "RNAT_ACL_1"
    aclaction = "ALLOW"
    priority  = "100"
    srcipval  = "192.168.10.22"
    destipval = "172.17.0.20"
  }
}

resource "citrixadc_inat" "foo" {
  name      = "ip4ip4"
  privateip = "192.168.2.5"
  publicip  = "172.17.1.2"
}

resource "citrixadc_rnat" "allrnat" {
  depends_on = [citrixadc_nsacls.allacls]

  rnatsname = "rnatsall"
  rnat {
    network = "192.168.20.0"
    netmask = "255.255.255.0"
  }

  rnat {
    network = "192.168.88.0"
    netmask = "255.255.255.0"
    natip   = "172.17.0.2"
  }

  rnat {
    aclname = "RNAT_ACL_1"
    natip   = "172.17.0.2"
  }
}
