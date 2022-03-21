resource "citrixadc_authenticationemailaction" "tf_emailaction" {
  name      = "tf_emailaction"
  username  = "username"
  password  = "secret"
  serverurl = "www/sdsd.com"
  timeout   = 100
  type      = "SMTP"
}