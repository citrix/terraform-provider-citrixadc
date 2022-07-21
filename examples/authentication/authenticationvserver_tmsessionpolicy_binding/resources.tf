# Since the tmsessionpolicy resource is not yet available on Terraform,
# the tf_tmsesspolicy policy must be created by hand(manually) in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add tmsessionaction tf_tmsessaction  -sessTimeout 30 -defaultAuthorization ALLOW
# add tmsessionpolicy tf_tmsesspolicy true tf_tmsessaction
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationvserver_tmsessionpolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = "tf_tmsesspolicy"
  priority  = 90
  bindpoint = "REQUEST"
}