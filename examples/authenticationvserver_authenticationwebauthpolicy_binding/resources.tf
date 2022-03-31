resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
  name                       = "tf_webauthaction"
  serverip                   = "1.2.3.4"
  serverport                 = 8080
  fullreqexpr                = "TRUE"
  scheme                     = "http"
  successrule                = "http.RES.STATUS.EQ(200)"
  defaultauthenticationgroup = "new_group"
}
resource "citrixadc_authenticationwebauthpolicy" "tf_webauthpolicy" {
  name   = "tf_webauthpolicy"
  rule   = "NS_TRUE"
  action = citrixadc_authenticationwebauthaction.tf_webauthaction.name
}
resource "citrixadc_authenticationvserver_authenticationwebauthpolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = citrixadc_authenticationwebauthpolicy.tf_webauthpolicy.name
  priority  = 80
  bindpoint = "AAA_RESPONSE"
}