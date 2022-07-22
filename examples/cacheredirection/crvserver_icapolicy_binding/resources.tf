# Since the icapolicy resource is not yet available on Terraform,
# the tf_icapolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add ica action tf_icaaction -accessProfileName default_ica_accessprofile
# add ica policy tf_icapolicy -rule true -action tf_icaaction
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_crvserver_icapolicy_binding" "crvserver_icapolicy_binding" {
  name       = citrixadc_crvserver.crvserver.name
  policyname = "tf_icapolicy"
  priority   = 1
}