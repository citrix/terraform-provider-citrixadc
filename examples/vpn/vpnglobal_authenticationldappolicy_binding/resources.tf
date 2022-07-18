# Since the authenticationldappolicy resource is not yet available on Terraform,
# the tf_ldappolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add authenticationldapaction tf_ldapaction -serverIP 5.5.5.5
# add authenticationldappolicy tf_ldappolicy ns_true tf_ldapaction

resource "citrixadc_vpnglobal_authenticationldappolicy_binding" "tf_bind" {
  policyname = "tf_ldappolicy"
  priority   = 20
}