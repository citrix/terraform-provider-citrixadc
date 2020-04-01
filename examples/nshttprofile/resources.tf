resource "citrixadc_nshttpprofile" "test_profile" {
  name  = "tf_httpprofile"
  http2 = "ENABLED"
}
