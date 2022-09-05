resource "citrixadc_feoaction" "tf_feoaction" {
  name              = "my_feoaction"
  cachemaxage       = 50
  imgshrinktoattrib = "true"
  imggiftopng       = "true"
}
