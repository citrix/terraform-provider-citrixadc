# Since the analyticsprofile resource is not yet available on Terraform,
# the new_profile profile must be created by hand(manually) in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add analyticsprofile new_profile -type tcpinsight
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_crvserver_analyticsprofile_binding" "crvserver_analyticsprofile_binding" {
  name             = citrixadc_crvserver.crvserver.name
  analyticsprofile = "new_profile"
  count            = 10
}