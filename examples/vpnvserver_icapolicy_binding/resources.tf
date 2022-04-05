# Since the icapolicy resource is not yet available on Terraform,
# the tf_icapolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add ica action tf_icaaction -accessProfileName default_ica_accessprofile
# add ica policy tf_icapolicy -rule true -action tf_icaaction
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserverexample"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_icapolicy_binding" "tf_binding" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = "tf_icapolicy"
  priority  = 30
  bindpoint = "AAA_RESPONSE"
}