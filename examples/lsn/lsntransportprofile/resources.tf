resource "citrixadc_lsntransportprofile" "tf_lsntransportprofile" {
  transportprofilename = "my_lsn_transportprofile"
  transportprotocol    = "TCP"
  portquota            = 10
  sessionquota         = 10
  groupsessionlimit    = 1000
}
