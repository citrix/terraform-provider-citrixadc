resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
  td        = 2
  aliasname = "tf_trafficdomain"
  vmac      = "DISABLED"
}
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_nstrafficdomain_vlan_binding" "tf_binding" {
  td   = citrixadc_nstrafficdomain.tf_trafficdomain.td
  vlan = citrixadc_vlan.tf_vlan.vlanid
}