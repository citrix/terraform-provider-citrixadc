resource "citrixadc_nscapacity" "tf_pooled" {
  bandwidth = 100
  unit      = "Mbps"
  edition   = "Platinum"

}
