resource "citrixadc_lsnclient_network6_binding" "tf_lsnclient_network6_binding" {
  clientname = "my_lsn_client"
  network6   = "2001:DB8:5001::/96"
}
