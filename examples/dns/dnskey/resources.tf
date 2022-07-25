resource "citrixadc_dnskey" "dnskey" {
  keyname            = "adckey_1"
  publickey          = "/nsconfig/dns/demo.key"
  privatekey         = "/nsconfig/dns/demo.private"
  expires            = 120
  units1             = "DAYS"
  notificationperiod = 7
  units2             = "DAYS"
  ttl                = 3600
}
