resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile" {
  appsprofilename   = "my_lsn_appsprofile"
  transportprotocol = "TCP"
  mapping           = "ENDPOINT-INDEPENDENT"
}
