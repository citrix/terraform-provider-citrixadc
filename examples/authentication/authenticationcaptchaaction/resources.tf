resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
  name                       = "tf_captchaaction"
  secretkey                  = "secret"
  sitekey                    = "key"
  serverurl                  = "http://www.example.com/"
  defaultauthenticationgroup = "new_group"
}