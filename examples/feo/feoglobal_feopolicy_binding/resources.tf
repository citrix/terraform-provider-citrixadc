resource "citrixadc_feoglobal_feopolicy_binding" "tf_feoglobal_feopolicy_binding" {
  policyname = "my_feopolicy"
  type       = "REQ_DEFAULT"
  priority   = 100
}
