resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tfmonitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor2" {
  monitorname = "tfmonitor2"
  type        = "PING"
}

resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "192.168.22.19"
  sessionexchange = "DISABLED"
}

resource "citrixadc_gslbservice" "tfgslbservice" {
  ip          = "192.168.18.81"
  port        = "80"
  servicename = "tfgslbservice"
  servicetype = "HTTP"
  sitename = citrixadc_gslbsite.site_local.sitename

  lbmonitorbinding {
      monitor_name = citrixadc_lbmonitor.tfmonitor1.monitorname
      weight = 80
  }

  lbmonitorbinding {
      monitor_name = citrixadc_lbmonitor.tfmonitor2.monitorname
      weight = 20
  }
}
