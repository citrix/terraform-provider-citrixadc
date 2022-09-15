resource "citrixadc_lsngroup_lsnpool_binding" "tf_lsngroup_lsnpool_binding" {
  groupname = "my_lsn_group"
  poolname  = "my_pool"
}
