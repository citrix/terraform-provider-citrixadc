# Since the dospolicy resource is not yet available on Terraform,
# the tf_dospolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add dospolicy tf_dospolicy -qDepth 25
resource "citrixadc_service" "tf_service" {
  servicetype         = "HTTP"
  name                = "tf_service"
  ipaddress           = "10.77.33.22"
  ip                  = "10.77.33.22"
  port                = "80"
  state               = "ENABLED"
  wait_until_disabled = true
}
resource "citrixadc_service_dospolicy_binding" "tf_binding" {
  name       = citrixadc_service.tf_service.name
  policyname = "tf_dospolicy"
}