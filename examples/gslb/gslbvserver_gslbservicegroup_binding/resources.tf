resource "citrixadc_gslbvserver_gslbservicegroup_binding" "tf_gslbvserver_gslbservicegroup_binding" {
  name             = citrixadc_gslbvserver.tf_gslbvserver.name
  servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
}

resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}

resource "citrixadc_gslbvserver" "tf_gslbvserver" {
  dnsrecordtype = "A"
  name          = "GSLB-East-Coast-Vserver"
  servicetype   = "HTTP"
  domain {
    domainname = "www.fooco.co"
    ttl        = "60"
  }
  domain {
    domainname = "www.barco.com"
    ttl        = "65"
  }
}
resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
  servicegroupname = "test_gslbvservicegroup"
  servicetype      = "HTTP"
  cip              = "DISABLED"
  healthmonitor    = "NO"
  sitename         = citrixadc_gslbsite.site_local.sitename
}



