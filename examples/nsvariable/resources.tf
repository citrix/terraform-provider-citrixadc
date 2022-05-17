resource "citrixadc_nsvariable" "tf_nsvariable" {
  name          = "tf_nsvariable"
  type          = "text(20)"
  scope         = "global"
  iffull        = "undef"
  ifvaluetoobig = "undef"
  ifnovalue     = "init"
  comment       = "Testing"
}