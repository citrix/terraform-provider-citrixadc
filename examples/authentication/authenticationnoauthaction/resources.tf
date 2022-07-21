resource "citrixadc_authenticationnoauthaction" "tf_noauthaction" {
  name                       = "tf_noauthaction"
  defaultauthenticationgroup = "group"
}