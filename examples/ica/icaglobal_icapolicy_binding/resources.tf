resource "citrixadc_icaglobal_icapolicy_binding" "tf_icaglobal_icapolicy_binding" {
  policyname = "my_ica_policy"
  priority   = 100
  type       = "ICA_REQ_DEFAULT"
}
