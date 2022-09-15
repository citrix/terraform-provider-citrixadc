resource "citrixadc_contentinspectionaction" "tf_contentinspectionaction" {
  name            = "my_ci_action"
  type            = "ICAP"
  icapprofilename = "reqmod-profile"
  servername      = "vicap"
  ifserverdown    = "DROP"
}
