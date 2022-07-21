# Since the cachepolicy resource is not yet available on Terraform,
# the tf_cachepolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add cache policy tf_cachepolicy -rule "http.req.url.query.contains(\"IssuePage\")" -action CACHE
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_crvserver_cachepolicy_binding" "crvserver_cachepolicy_binding" {
  name       = citrixadc_crvserver.crvserver.name
  policyname = "tf_cachepolicy"
  priority   = 10
  bindpoint =  "REQUEST"

}