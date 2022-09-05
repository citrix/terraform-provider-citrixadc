resource "citrixadc_feopolicy" "tf_feopolicy" {
  name   = "my_feopolicy"
  action = "my_feoaction"
  rule   = "true"
}
