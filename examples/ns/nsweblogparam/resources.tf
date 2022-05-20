resource "citrixadc_nsweblogparam" "tf_nsweblofparam" {
  buffersizemb  = 32
  customreqhdrs = ["req1", "req2"]
  customrsphdrs = ["res1", "res2"]
}