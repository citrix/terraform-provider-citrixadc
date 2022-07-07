resource "citrixadc_linkset" "tflinkset" {
  linkset_id = "LS/1"

  interfacebinding = [
    "1/1/1",
    "2/1/1",
  ]
}
