# Since the authorizationpolicy resource is not yet available on Terraform,
# the tf_authorizationpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add authorization policy tf_authorizationpolicy ns_true ALLOW

resource "citrixadc_lbvserver_authorizationpolicy_binding" "tf_lbvserver_authorizationpolicy_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  policyname = "tf_authorizationpolicy"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.33"
  port        = 80
  servicetype = "HTTP"
}