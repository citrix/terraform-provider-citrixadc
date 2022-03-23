resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
  name         = "tf_idpprofile"
  clientid     = "cliId"
  clientsecret = "secret"
  redirecturl  = "http://www.example.com/1/"
}