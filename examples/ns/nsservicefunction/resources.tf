resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_nsservicefunction" "tf_servicefunc" {
  servicefunctionname = "tf_servicefunc"
  ingressvlan         = citrixadc_vlan.tf_vlan.vlanid
}