
resource "netscaler_gslbsite" "site_local" {
  sitename = "Site-Local"
  siteipaddress = "172.31.96.234"
  sessionexchange = "DISABLED"
  
}

