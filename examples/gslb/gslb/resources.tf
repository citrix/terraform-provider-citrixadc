resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}

resource "citrixadc_gslbsite" "site_remote" {
  sitename        = "Site-Remote"
  siteipaddress   = "172.31.48.18"
  sessionexchange = "ENABLED"
}

resource "citrixadc_gslbservice" "gslb_svc1" {
  ip          = "172.16.1.121"
  port        = "80"
  servicename = "gslb1vservice"
  servicetype = "HTTP"
  sitename    = citrixadc_gslbsite.site_local.sitename
}

resource "citrixadc_gslbservice" "gslb_svc2" {
  ip          = "172.16.17.121"
  port        = "80"
  servicename = "gslb2vservice"
  servicetype = "HTTP"
  sitename    = citrixadc_gslbsite.site_remote.sitename
}

resource "citrixadc_gslbvserver" "foo" {
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
  service {
    servicename = citrixadc_gslbservice.gslb_svc1.servicename
    weight      = "75"
  }
  service {
    servicename = citrixadc_gslbservice.gslb_svc2.servicename
    weight      = "100"
  }
}

