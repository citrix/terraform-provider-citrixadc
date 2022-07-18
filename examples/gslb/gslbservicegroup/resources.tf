resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
  servicegroupname = "test_gslbvservicegroup"
  servicetype      = "HTTP"
  cip              = "DISABLED"
  healthmonitor    = "NO"
  sitename         = citrixadc_gslbsite.site_local.sitename
}
resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}
