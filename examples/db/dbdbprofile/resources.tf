resource "citrixadc_dbdbprofile" "tf_dbdbprofile" {
  name           = "my_dbprofile"
  stickiness     = "YES"
  conmultiplex   = "ENABLED"
  interpretquery = "YES"
}
