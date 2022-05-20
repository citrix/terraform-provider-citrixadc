resource "citrixadc_nsspparams" "tf_nsspparams" {
  basethreshold = 200
  throttle      = "Aggressive"
}