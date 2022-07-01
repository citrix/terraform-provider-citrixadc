resource "citrixadc_filteraction" "tf_filteraction" {
  name  = "tf_filteraction"
  qual  = "corrupt"
  value = "X-Forwarded-For"
}