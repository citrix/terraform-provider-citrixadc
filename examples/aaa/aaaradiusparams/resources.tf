resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
  radkey             = "sslvpn"
  radnasip           = "ENABLED"
  serverip           = "10.222.74.158"
  authtimeout        = 8
}
