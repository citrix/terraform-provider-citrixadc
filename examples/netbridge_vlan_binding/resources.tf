resource "citrixadc_netbridge" "tf_netbridge" {
  name         = "tf_netbridge"
}
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_netbridge_vlan_binding" "tf_binding" {
  name = citrixadc_netbridge.tf_netbridge.name
  vlan = citrixadc_vlan.tf_vlan.vlanid
}