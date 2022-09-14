resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
  name   = "my_appqoepolicy"
  rule   = "true"
  action = "my_act"
}
