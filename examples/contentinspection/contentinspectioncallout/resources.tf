resource "citrixadc_contentinspectioncallout" "tf_contentinspectioncalloout" {
  name        = "my_ci_callout"
  type        = "ICAP"
  profilename = "reqmod-profile"
  servername  = "icapsv1"
  returntype  = "TEXT"
  resultexpr  = "icap.res.header(\"ISTag\")"
}
