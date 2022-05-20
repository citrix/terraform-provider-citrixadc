# mamlimit should be less than Maximum Memory Usage Limit
resource "citrixadc_extendedmemoryparam" "tf_extendedmemoryparam" {
  memlimit = 512
}