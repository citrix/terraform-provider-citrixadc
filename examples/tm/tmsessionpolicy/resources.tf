resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
  name   = "my_tmsession_policy"
  rule   = "true"
  action = "tf_tmsessaction"
}
