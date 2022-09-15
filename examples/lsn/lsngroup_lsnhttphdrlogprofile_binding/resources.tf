resource "citrixadc_lsngroup_lsnhttphdrlogprofile_binding" "tf_lsngroup_lsnhttphdrlogprofile_binding" {
  groupname             = "my_lsn_group"
  httphdrlogprofilename = "my_httplogprofile"
}
