resource "citrixadc_aaaotpparameter" "tf_aaaotpparameter" {
  encryption = "ON"
  maxotpdevices = 5
}