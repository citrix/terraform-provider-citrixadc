# Since the authorizationpolicy resource is not yet available on Terraform,
# the tf_authorizationpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add authorization policy tf_authorizationpolicy ns_true ALLOW

resource "citrixadc_csvserver_authorizationpolicy_binding" "tf_csvserver_authorizationpolicy_binding" {
  name = citrixadc_csvserver.tf_csvserver.name
  policyname = "tf_authorizationpolicy"
  priority = 5
}

resource "citrixadc_csvserver" "tf_csvserver" {
  name = "tf_csvserver"
  ipv46 = "10.202.11.11"
  port = 8080
  servicetype = "HTTP"
}