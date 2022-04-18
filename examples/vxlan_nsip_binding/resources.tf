resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_nsip" "tf_snip" {
  ipaddress = "10.222.74.146"
  type      = "SNIP"
  netmask   = "255.255.255.0"
  icmp      = "ENABLED"
  state     = "ENABLED"
}
resource "citrixadc_vxlan_nsip_binding" "tf_binding" {
  vxlanid   = citrixadc_vxlan.tf_vxlan.vxlanid
  ipaddress = citrixadc_nsip.tf_snip.ipaddress
  netmask   = citrixadc_nsip.tf_snip.netmask
}