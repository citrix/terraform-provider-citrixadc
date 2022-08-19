# Since the tmsessionpolicy resource is not yet available on Terraform,
# the tf_tmsesspolicy policy must be created by hand(manually) in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add tmsessionaction tf_tmsessaction  -sessTimeout 30 -defaultAuthorization ALLOW
# add tmsessionpolicy tf_tmsesspolicy true tf_tmsessaction
resource "citrixadc_aaauser_tmsessionpolicy_binding" "tf_aaauser_tmsessionpolicy_binding" {
  username = "user1"
  policy    = "tf_tmsesspolicy"
  priority  = 100
}