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
resource "citrixadc_vpnglobal_authenticationnegotiatepolicy_binding" "tf_binding" {
  policyname             = citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy.name
  secondary              = "false"
  priority               = 10
  gotopriorityexpression = "END"
}