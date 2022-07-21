
resource "citrixadc_gslbservice_lbmonitor_binding" "tf_gslbservice_lbmonitor_binding" {
  monitor_name = citrixadc_lbmonitor.tfmonitor1.monitorname
  monstate    = "DISABLED"
  servicename = citrixadc_gslbservice.tf_gslbservice.servicename
  weight      = "20"

}

resource "citrixadc_gslbservice" "tf_gslbservice" {
  ip          = "172.16.1.128"
  port        = "80"
  servicename = "test_gslb1vservice"
  servicetype = "HTTP"
  sitename    = citrixadc_gslbsite.tf_gslbsite.sitename
}

resource "citrixadc_gslbsite" "tf_gslbsite" {
  sitename      = "test_sitename"
  siteipaddress = "10.222.70.79"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "test_monitor"
  type        = "HTTP"
}
