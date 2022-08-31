resource "citrixadc_tmglobal_tmtrafficpolicy_binding" "tf_tmglobal_tmtrafficpolicy_binding" {
  priority   = "100"
  policyname = "my_tmtrafficpolicy"
}
