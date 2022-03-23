resource "citrixadc_authenticationpushservice" "tf_pushservice" {
  name            = "tf_pushservice"
  clientid        = "cliId"
  clientsecret    = "secret"
  customerid      = "cusID"
  refreshinterval = 50
}