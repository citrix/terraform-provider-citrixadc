resource "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id     = 2
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
}
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_bridgegroup_vlan_binding" "tf_binding" {
  bridgegroup_id = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
  vlan           = citrixadc_vlan.tf_vlan.vlanid
}