resource "citrixadc_lsnappsprofile_port_binding" "tf_lsnappsprofile_port_binding" {
  appsprofilename = "my_lsn_profile"
  lsnport         = "80"
}
