resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 40
  aliasname = "Management VLAN"
}
resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  vlan               = citrixadc_vlan.tf_vlan.vlanid
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}