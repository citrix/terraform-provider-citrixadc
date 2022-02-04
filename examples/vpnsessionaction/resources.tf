resource "citrixadc_vpnsessionaction" "tf_sessionaction" {
  name                       = "newsession"
  sesstimeout                = "10"
  defaultauthorizationaction = "ALLOW"
  transparentinterception    = "ON"
  clientidletimeout          = "10"
  sso                        = "ON"
  icaproxy                   = "ON"
  wihome                     = "https://citrix.lab.com"
  clientlessvpnmode          = "DISABLED"
  httpport                   = [8080, 8000, 808]
}