# Since the authenticationlocalpolicy resource is not yet available on Terraform,
# the tf_localpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add authenticationlocalpolicy tf_localpolicy ns_true
resource "citrixadc_vpnglobal_authenticationlocalpolicy_binding" "tf_bind" {
  policyname = "tf_localpolicy"
  priority   = 20
}