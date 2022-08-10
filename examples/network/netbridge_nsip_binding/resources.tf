resource "citrixadc_netbridge_nsip_binding" "tf_netbridge_nsip_binding" {
  name      = "my_netbridge"
  netmask   = "255.255.255.192"
  ipaddress = "10.222.74.128"
}
