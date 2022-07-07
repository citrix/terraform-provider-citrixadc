resource "citrixadc_nsservicepath" "tf_servicepath" {
  servicepathname = "tf_servicepath"
}
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_nsservicefunction" "tf_servicefunc" {
  servicefunctionname = "tf_servicefunc"
  ingressvlan         = citrixadc_vlan.tf_vlan.vlanid
}
resource "citrixadc_nsservicepath_nsservicefunction_binding" "tf_binding" {
  servicepathname = citrixadc_nsservicepath.tf_servicepath.servicepathname
  servicefunction = citrixadc_nsservicefunction.tf_servicefunc.servicefunctionname
  index           = 2
}