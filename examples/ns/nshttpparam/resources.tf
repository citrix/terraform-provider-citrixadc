resource "citrixadc_nshttpparam" "tf_nshttpparam" {
  dropinvalreqs             = "OFF"
  markconnreqinval          = "OFF"
  maxreusepool              = 1
  markhttp09inval           = "OFF"
  insnssrvrhdr              = "OFF"
  logerrresp                = "OFF"
  conmultiplex              = "ENABLED"
  http2serverside           = "OFF"
  ignoreconnectcodingscheme = "DISABLED"
}