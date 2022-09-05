resource "citrixadc_lldpparam" "tf_lldpparam" {
  holdtimetxmult = 3
  mode           = "TRANSMITTER"
  timer          = 40
}
