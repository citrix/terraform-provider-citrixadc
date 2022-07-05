resource "citrixadc_dnsmxrec" "dnsmxrec" {
  domain = "example.com"
  mx     = "mail.example.com"
  pref   = 5
  ttl    = 3600
}