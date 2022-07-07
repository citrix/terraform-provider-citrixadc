resource "citrixadc_dnsptrrec" "tf_dnsptrrec" {
  reversedomain = "0.2.0.192.in-addr.arpa"
  domain        = "example.com"
  ttl           = 3600
}
