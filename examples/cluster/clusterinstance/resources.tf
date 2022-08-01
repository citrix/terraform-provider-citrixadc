resource "citrixadc_clusterinstance" "tf_clusterinstance" {
  clid          = 1
  deadinterval  = 8
  hellointerval = 600
}
