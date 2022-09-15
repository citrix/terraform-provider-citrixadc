resource "citrixadc_lsnlogprofile" "tf_lsnlogprofile" {
  logprofilename = "my_lsn_logprofile"
  logsubscrinfo   = "ENABLED"
  logcompact      = "ENABLED"
  logipfix        = "ENABLED"
}
