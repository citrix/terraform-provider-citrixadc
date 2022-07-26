# Since the feopolicy resource is not yet available on Terraform,
# the tf_feopolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add feo policy tf_feopolicy TRUE BASIC
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_crvserver_feopolicy_binding" "crvserver_feopolicy_binding" {
  name       = citrixadc_crvserver.crvserver.name
  policyname = "tf_feopolicy"
  priority   = 10
}