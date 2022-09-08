resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
  name             = "my_traffic_action"
  apptimeout       = 5
  sso              = "OFF"
  persistentcookie = "ON"
}
