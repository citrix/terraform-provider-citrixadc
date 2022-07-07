resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
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
resource "citrixadc_bridgetable" "tf_bridgetable" {
  mac       = "00:00:00:00:00:01"
  vxlan     = citrixadc_vxlan.tf_vxlan.vxlanid
  vtep      = "2.34.5.6"
  bridgeage = "300"
}