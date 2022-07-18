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