resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 40
  aliasname = "Management VLAN"
}
resource "citrixadc_vlan" "tf_vlan1" {
  vlanid    = 41
  aliasname = "Management VLAN"
}
resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
  name = "tf_vxlanvlanmp"
}
resource "citrixadc_vxlanvlanmap_vxlan_binding" "tf_binding" {
  name = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
  vxlan = citrixadc_vxlan.tf_vxlan.vxlanid
  vlan = [citrixadc_vlan.tf_vlan.vlanid,citrixadc_vlan.tf_vlan1.vlanid]
}