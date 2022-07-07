resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction" {
  name                       = "tf_negotiateaction"
  domain                     = "DomainName"
  domainuser                 = "usersame"
  domainuserpasswd           = "password"
  ntlmpath                   = "http://www.example.com/"
  defaultauthenticationgroup = "new_grpname"
}
resource "citrixadc_authenticationnegotiatepolicy" "tf_negotiatepolicy" {
  name      = "tf_negotiatepolicy"
  rule      = "ns_true"
  reqaction = citrixadc_authenticationnegotiateaction.tf_negotiateaction.name
}
resource "citrixadc_authenticationvserver_authenticationnegotiatepolicy_binding" "tf_binding" {
  name            = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy          = citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy.name
  priority        = 9
  groupextraction = "false"
  bindpoint       = "REQUEST"
}