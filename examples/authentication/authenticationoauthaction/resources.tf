resource "citrixadc_authenticationoauthaction" "tf_authenticationoauthaction" {
  name                  = "tf_authenticationoauthaction"
  authorizationendpoint = "https://example.com/"
  tokenendpoint         = "https://ssexample.com/"
  clientid              = "cliId"
  clientsecret          = "secret"
  resourceuri           = "http://www.sd.com"
}