resource "citrixadc_lsnip6profile" "tf_lsnaip6profile" {
  name     = "my_lsn_ip6profile"
  type     = "DS-Lite"
  network6 = "2003::/64"
}
