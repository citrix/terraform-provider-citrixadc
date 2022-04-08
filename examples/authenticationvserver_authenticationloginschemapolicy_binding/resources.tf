resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationloginschema" "tf_loginschema" {
  name                    = "tf_loginschema"
  authenticationschema    = "LoginSchema/SingleAuth.xml"
  ssocredentials          = "YES"
  authenticationstrength  = "30"
  passwordcredentialindex = "10"
}
resource "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy" {
  name    = "tf_loginschemapolicy"
  rule    = "true"
  action  = citrixadc_authenticationloginschema.tf_loginschema.name
  comment = "samplenew_testing"
}
resource "citrixadc_authenticationvserver_authenticationloginschemapolicy_binding" "tf_binding" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy.name
  priority  = 77
  bindpoint = "REQUEST" // unbind doesnot work for other values
}