resource "citrixadc_csaction" "tf_csaction" {
  name            = "tf_test_csaction"
  targetlbvserver = citrixadc_lbvserver.tf_image_lb2.name
  comment         = "Forwards image requests to the image_lb2"
}

resource "citrixadc_lbvserver" "tf_image_lb" {
  name        = "image_lb"
  ipv46       = "10.0.2.5"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_image_lb2" {
  name        = "image_lb2"
  ipv46       = "10.0.2.6"
  port        = "80"
  servicetype = "HTTP"
}
