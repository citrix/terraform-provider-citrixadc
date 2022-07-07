resource "citrixadc_dnsprofile" "dnsprofile" {
  dnsprofilename         = "tf_profile1"
  dnsquerylogging        = "DISABLED"
  dnsanswerseclogging    = "DISABLED"
  dnsextendedlogging     = "DISABLED"
  dnserrorlogging        = "DISABLED"
  cacherecords           = "ENABLED"
  cachenegativeresponses = "ENABLED"
  dropmultiqueryrequest  = "DISABLED"
  cacheecsresponses      = "DISABLED"
}
