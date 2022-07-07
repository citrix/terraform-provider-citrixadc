# Since the cachepolicy resource is not yet available on Terraform,
# the tf_cachepolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add cache policy tf_cachepolicy -rule "http.req.url.query.contains(\"IssuePage\")" -action CACHE
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_examplecom"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_cachepolicy_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = "tf_cachepolicy"
  priority  = 5
  bindpoint = "REQUEST" 
}
