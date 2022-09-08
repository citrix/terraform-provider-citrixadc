resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
  name             = "my_analyticsprofile"
  type             = "webinsight"
  httppagetracking = "DISABLED"
  httpurl          = "ENABLED"
}
