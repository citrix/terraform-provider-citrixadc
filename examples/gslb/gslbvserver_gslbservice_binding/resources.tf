resource "citrixadc_gslbvserver_gslbservice_binding" "tf_gslbvserver_gslbservice_binding" {
  name        = citrixadc_gslbvserver.tf_gslbvserver.name
  servicename = citrixadc_gslbservice.gslb_svc1.servicename
}

resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}

resource "citrixadc_gslbservice" "gslb_svc1" {
  ip          = "172.16.1.121"
  port        = "80"
  servicename = "gslb1vservice"
  servicetype = "HTTP"
  sitename    = citrixadc_gslbsite.site_local.sitename
}

resource "citrixadc_gslbvserver" "tf_gslbvserver" {
  dnsrecordtype = "A"
  name          = "gslb_vserver"
  servicetype   = "HTTP"
  domain {
    domainname = "www.fooco.co"
    ttl        = "60"
  }
}

