resource "citrixadc_nspartition" "tf_nspartition" {
  partitionname = "tf_nspartition"
  maxbandwidth  = 10240
  minbandwidth  = 512
  maxconn       = 512
  maxmemlimit   = 11
}
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_nspartition_vlan_binding" "tf_binding" {
  partitionname = citrixadc_nspartition.tf_nspartition.partitionname
  vlan          = citrixadc_vlan.tf_vlan.vlanid
}