resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
  name                       = "tf_webauthaction"
  serverip                   = "1.2.3.4"
  serverport                 = 8080
  fullreqexpr                = "TRUE"
  scheme                     = "http"
  successrule                = "http.RES.STATUS.EQ(200)"
  defaultauthenticationgroup = "new_group"
}