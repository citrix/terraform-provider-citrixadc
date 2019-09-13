resource "citrixadc_csvserver" "csvserver" {
  name        = "ssltest_csvserver"
  ipv46       = "10.78.1.38"
  port        = "443"
  servicetype = "SSL"
  ciphers     = ["HIGH", "DEFAULT"]
}
