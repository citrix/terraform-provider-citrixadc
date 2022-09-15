resource "citrixadc_contentinspectionglobal_contentinspectionpolicy_binding" "tf_ci_binding" {
  policyname = "my_ci_policy"
  priority   = 100
}
