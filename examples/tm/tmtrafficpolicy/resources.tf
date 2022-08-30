resource "citrixadc_tmtrafficpolicy" "tf_tmtrafficpolicy" {
  name   = "my_tmtraffic_policy"
  rule   = "true"
  action = "my_tmtraffic_action"
}
