resource "citrixadc_nsip" "tf_srcip" {
  ipaddress = "11.22.33.44"
  type      = "SNIP"
  netmask   = "255.255.255.0"
}
resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_vxlan_srcip_binding" "tf_binding" {
  vxlanid = citrixadc_vxlan.tf_vxlan.vxlanid
  srcip   = citrixadc_nsip.tf_srcip.ipaddress
}