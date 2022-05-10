resource "citrixadc_radiusnode" "tf_radiusnode" {
  nodeprefix = "10.10.10.10/32"
  radkey     = "secret"
}