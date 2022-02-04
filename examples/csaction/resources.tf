resource "citrixadc_csaction" "tf_csaction" {
  name            = "tf_test_csaction1"
  targetlbvserver = citrixadc_lbvserver.tf_image_lb.name
  comment         = "Forwards image requests to the image_lb"
}

resource "citrixadc_lbvserver" "tf_image_lb" {
  name        = "image_lb"
  ipv46       = "10.0.2.5"
  port        = "80"
  servicetype = "HTTP"
}

# For update, change the `targetlbvserser` accordingly
resource "citrixadc_lbvserver" "tf_image_lb2" {
  name        = "image_lb2"
  ipv46       = "10.0.2.6"
  port        = "80"
  servicetype = "HTTP"
}
