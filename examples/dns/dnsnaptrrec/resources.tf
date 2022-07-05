resource "citrixadc_dnsnaptrrec" "dnsnaptrrec" {
  domain      = "example.com"
  order       = 10
  preference  = 2
  ttl         = 3600
  replacement = "example1.com"
}