// Make sure the dhfile name does not exist in the target ADC
// citrixadc_sslparam does not support UPDATE operation, so to change any attributes here, first delete the dhfile, if present
resource "citrixadc_ssldhparam" "tf_dhparam" {
  dhfile = "/nsconfig/ssl/tf_dhfile"
  bits   = "512"
  gen    = "2"
}
