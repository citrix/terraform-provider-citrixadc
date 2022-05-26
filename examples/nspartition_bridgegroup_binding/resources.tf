resource "citrixadc_nspartition" "tf_nspartition" {
  partitionname = "tf_nspartition"
  maxbandwidth  = 10240
  minbandwidth  = 512
  maxconn       = 512
  maxmemlimit   = 11
}
resource "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id     = 2
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
}
resource "citrixadc_nspartition_bridgegroup_binding" "tf_binding" {
  partitionname = citrixadc_nspartition.tf_nspartition.partitionname
  bridgegroup   = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
}