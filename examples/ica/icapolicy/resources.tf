resource "citrixadc_icapolicy" "tf_icapolicy" {
  name   = "my_ica_policy"
  rule   = true
  action = "my_ica_action"
}
