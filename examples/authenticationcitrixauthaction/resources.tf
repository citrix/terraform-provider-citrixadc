resource "citrixadc_authenticationcitrixauthaction" "tf_citrixauthaction" {
  name               = "tf_citrixauthaction"
  authenticationtype = "CITRIXCONNECTOR"
  authentication     = "DISABLED"
}