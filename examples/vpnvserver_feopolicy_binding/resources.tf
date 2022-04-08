# Since the feopolicy resource is not yet available on Terraform,
# the tf_feopolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add feo policy tf_feopolicy TRUE BASIC
# bindpoint = RESPONSE does not work ("message": "Invalid bind point.")
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vservercom"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_feopolicy_binding" "tf_bind" {
  name       = citrixadc_vpnvserver.tf_vpnvserver.name
  policy     = "tf_feopolicy"
  priority   = 90
  bindpoint  = "REQUEST"  // doesnot unbind for other values
}