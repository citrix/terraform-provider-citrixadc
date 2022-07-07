resource "citrixadc_authenticationloginschema" "tf_loginschema" {
  name                    = "tf_loginschema"
  authenticationschema    = "LoginSchema/SingleAuth.xml"
  ssocredentials          = "YES"
  authenticationstrength  = "30"
  passwordcredentialindex = "10"
}
resource "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy" {
  name      = "tf_loginschemapolicy"
  rule      = "true"
  action    = citrixadc_authenticationloginschema.tf_loginschema.name
  comment   = "samplenew_testing"
}