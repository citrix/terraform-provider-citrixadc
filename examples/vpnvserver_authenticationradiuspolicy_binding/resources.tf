# Since the authenticationradiuspolicy resource is not yet available on Terraform,
# the tf_radiuspolicy policy must be created by hand(manually) in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add authenticationradiusaction tf_radiusaction -serverIP 2.3.2.1 -radKey WareTheLorax
# add authenticationradiuspolicy tf_radiuspolicy ns_true tf_radiusaction

resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserver"
  servicetype = "SSL"
}
resource "citrixadc_vpnvserver_authenticationradiuspolicy_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = "tf_radiuspolicy"
  priority  = 20
}
