resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
  name         = "tf_idpprofile"
  clientid     = "cliId"
  clientsecret = "secret"
  redirecturl  = "http://www.example.com/1/"
}
resource "citrixadc_authenticationoauthidppolicy" "tf_idppolicy" {
  name    = "tf_idppolicy"
  rule    = "true"
  action  = citrixadc_authenticationoauthidpprofile.tf_idpprofile.name
  comment = "aboutpolicy"
}
resource "citrixadc_authenticationvserver_authenticationoauthidppolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = citrixadc_authenticationoauthidppolicy.tf_idppolicy.name
  priority  = 90
  bindpoint = "REQUEST" //unbind doesnot work for other values
}