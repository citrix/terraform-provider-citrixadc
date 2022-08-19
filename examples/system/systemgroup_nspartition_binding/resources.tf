resource "citrixadc_systemgroup_nspartition_binding" "tf_systemgroup_nspartition_binding" {
  groupname     = citrixadc_systemgroup.tf_systemgroup.groupname
  partitionname = citrixadc_nspartition.tf_nspartition.partitionname
}

resource "citrixadc_systemgroup" "tf_systemgroup" {
  groupname = "tf_systemgroup"
  timeout   = 999
}

resource "citrixadc_nspartition" "tf_nspartition" {
  partitionname = "tf_nspartition"
  maxbandwidth  = 10240
  minbandwidth  = 512
  maxconn       = 512
  maxmemlimit   = 11
}
