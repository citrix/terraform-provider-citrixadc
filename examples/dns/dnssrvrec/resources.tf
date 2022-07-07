resource "citrixadc_dnssrvrec" "dnssrvrec" {
  domain   = "example.com"
  target   = "_sip._udp.example.com"
  priority = 1
  weight   = 1
  port     = 22
  ttl      = 3600
}
