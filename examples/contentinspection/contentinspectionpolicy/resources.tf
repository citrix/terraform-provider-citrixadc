resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
  name   = "my_ci_policy"
  rule   = "true"
  action = "RESET"
}
