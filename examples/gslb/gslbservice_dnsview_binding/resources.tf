resource "citrixadc_gslbservice_dnsview_binding" "tf_gslbservice_dnsview_binding" {
  servicename = citrixadc_gslbservice.gslb_svc1.servicename
  viewname    = citrixadc_dnsview.tf_dnsview.viewname
  viewip      = "192.168.2.1"
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
  sitename    = citrixadc_gslbsite.site_remote.sitename
}

resource "citrixadc_dnsview" "tf_dnsview" {
  viewname = "view1"
}

