resource "citrixadc_nshmackey" "tf_nshmackey" {
  name     = "tf_nshmackey"
  digest   = "MD4"
  keyvalue = "AUTO"
  comment  = "Testing"
}