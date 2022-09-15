resource "citrixadc_lsngroup_lsnappsprofile_binding" "tf_lsngroup_lsnappsprofile_binding" {
  groupname       = "my_lsn_group"
  appsprofilename = "my_lsn_profile"
}
