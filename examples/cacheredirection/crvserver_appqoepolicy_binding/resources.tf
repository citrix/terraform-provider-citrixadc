# Since the appqoepolicy resource is not yet available on Terraform,
# the tf_appqoepolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add appqoe action tf_appqoeaction -priority MEDIUM
# add appqoe policy tf_appqoepolicy -rule true -action tf_appqoeaction
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_crvserver_appqoepolicy_binding" "crvserver_appqoepolicy_binding" {
  name       = citrixadc_crvserver.crvserver.name
  policyname = "tf_appqoepolicy"
  priority   = 10
}