
resource "netscaler_gslbsite" "site_local" {
  sitename = "Site-Local"
  siteipaddress = "172.31.96.234"
  sessionexchange = "DISABLED"
}

resource "netscaler_gslbservice" "gslb_svc1" {
  ip = "172.16.1.121"
  port = "80"
  servicename = "gslb1vservice"
  servicetype = "HTTP"
  sitename = "${netscaler_gslbsite.site_local.sitename}"
}

resource "netscaler_gslbvserver" "foo" {

  dnsrecordtype = "A"
  name = "GSLB-East-Coast-Vserver"
  servicetype = "HTTP"
  domain {
	  domainname =  "www.fooco.co"
	  ttl = "60"
  }
  domain {
	  domainname = "www.barco.com"
	  ttl = "55"
  }
}

