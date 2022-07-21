resource "citrixadc_authenticationloginschema" "tf_loginschema" {
  name                    = "tf_loginschema"
  authenticationschema    = "LoginSchema/SingleAuth.xml"
  ssocredentials          = "YES"
  authenticationstrength  = "30"
  passwordcredentialindex = "10"
}