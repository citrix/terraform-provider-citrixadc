resource "citrixadc_authenticationpolicylabel" "tf_authenticationpolicylabel" {
  labelname = "tf_authenticationpolicylabel"
  type      = "AAATM_REQ"
  comment   = "Testing"
}