resource "citrixadc_uservserver" "tf_uservserver" {
  name         = "my_user_vserver"
  userprotocol = "MQTT"
  ipaddress    = "10.222.74.180"
  port         = 3200
  defaultlb    = "mysv"
}