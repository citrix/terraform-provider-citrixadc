resource "citrixadc_arp" "tf_arp" {
  ipaddress = "10.222.74.175"
  mac       = "3B:FD:37:27:A1:F8"
  vxlan     =  4
}
