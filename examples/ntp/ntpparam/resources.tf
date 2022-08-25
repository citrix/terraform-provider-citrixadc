resource "citrixadc_ntpparam" "tf_ntpparam" {
  authentication = "YES"
  trustedkey     = [123, 456]
  autokeylogsec  = 15
  revokelogsec   = 20
}
