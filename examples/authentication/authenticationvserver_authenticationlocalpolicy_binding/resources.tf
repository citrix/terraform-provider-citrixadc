resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationlocalpolicy" "tf_authenticationlocalpolicy" {
  name = "tf_authenticationlocalpolicy"
  rule = "ns_true"
}
resource "citrixadc_authenticationvserver_authenticationlocalpolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy.name
  priority  = 90
  bindpoint = "AAA_RESPONSE"
}