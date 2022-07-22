resource "citrixadc_gslbservicegroup_lbmonitor_binding" "tf_gslbservicegroup_lbmonitor_binding" {
  weight           = 20
  servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
  monitor_name      = citrixadc_lbmonitor.tfmonitor1.monitorname

}

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

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "test_monitor"
  type        = "HTTP"
}
