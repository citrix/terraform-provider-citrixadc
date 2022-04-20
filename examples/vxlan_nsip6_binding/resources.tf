resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_nsip6" "test_nsip" {
  ipv6address = "2001:db8:100::fb/64"
  type        = "VIP"
  icmp        = "DISABLED"
}
resource "citrixadc_vxlan_nsip6_binding" "tf_binding" {
  vxlanid   = citrixadc_vxlan.tf_vxlan.vxlanid
  ipaddress = citrixadc_nsip6.test_nsip.ipv6address
  netmask   = "255.255.255.0"
}