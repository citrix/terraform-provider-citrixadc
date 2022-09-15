resource "citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding" "tf_ci_binding" {
  labelname  = "my_ci_label"
  policyname = "my_ci_policy"
  priority   = 100
}
