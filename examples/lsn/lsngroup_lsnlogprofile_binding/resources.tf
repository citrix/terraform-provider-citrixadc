resource "citrixadc_lsngroup_lsnlogprofile_binding" "tf_lsngroup_lsnlogprofile_binding" {
  groupname      = "my_lsn_group"
  logprofilename = "my_lsn_logprofile"
}
