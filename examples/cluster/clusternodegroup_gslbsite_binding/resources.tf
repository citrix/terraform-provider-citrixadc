resource "citrixadc_clusternodegroup_gslbsite_binding" "tf_clusternodegroup_gslbsite_binding" {
  gslbsite = citrixadc_gslbsite.site_remote.sitename
  name     = "my_group"
}

resource "citrixadc_gslbsite" "site_remote" {
  sitename        = "my_local_site"
  siteipaddress   = "10.222.74.169"
  sessionexchange = "DISABLED"
  sitetype        = "LOCAL"
}
