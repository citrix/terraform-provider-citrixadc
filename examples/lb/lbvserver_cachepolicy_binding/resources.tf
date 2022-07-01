# Since the cachepolicy resource is not yet available on Terraform,
# the tf_cachepolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add cache policy tf_cachepolicy -rule "http.req.url.query.contains(\"IssuePage\")" -action CACHE

resource "citrixadc_lbvserver_cachepolicy_binding" "tf_citrixadc_lbvserver_cachepolicy_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  policyname = "tf_cachepolicy"
  priority = 1
  bindpoint = "REQUEST"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.33"
  port        = 80
  servicetype = "HTTP"
}