# Since the cachepolicy resource is not yet available on Terraform,
# the tf_cachepolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add cache policy tf_cachepolicy -rule "http.req.url.query.contains(\"IssuePage\")" -action CACHE
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationvserver_cachepolicy_binding" "tf_binding" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = "tf_cachepolicy"
  bindpoint = "REQUEST" // notable to unbind for anyother values in CLI also
  priority  = 9
}