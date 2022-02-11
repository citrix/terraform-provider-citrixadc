resource "citrixadc_vpntrafficaction" "tf_action" {
  name       = "Testing"
  qual       = "tcp"
  apptimeout = 20
  fta        = "OFF"
  hdx        = "OFF"
  sso        = "ON"
}