resource "citrixadc_nsvariable" "tf_nsvariable" {
  name          = "tf_nsvariable"
  type          = "text(20)"
  scope         = "global"
  iffull        = "undef"
  ifvaluetoobig = "undef"
  ifnovalue     = "init"
  comment       = "Testing"
}

# Refer the Documentation for "variable" argument value
resource "citrixadc_nsassignment" "tf_nsassignment" {
  name     = "tf_nsassignment"
  variable = join("", ["$", citrixadc_nsvariable.tf_nsvariable.name])
  set      = 1
  comment  = "Testing"
}