# Since the videooptimizationpacingpolicy resource is not yet available on Terraform,
# the tf_videooptimizationpacingpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add videooptimization pacingaction myPacingaction -rate 2000
# add videooptimization pacingpolicy tf_pacingpolicy -rule TRUE -action myPacingaction

resource "citrixadc_lbvserver_videooptimizationpacingpolicy_binding" "tf_lbvserver_videooptimizationpacingpolicy_binding" {
  bindpoint = "REQUEST"
      gotopriorityexpression = "END"
      name = citrixadc_lbvserver.tf_lbvserver.name
      policyname = "tf_pacingpolicy"
      priority = 1
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.33"
  port        = 80
  servicetype = "HTTP"
}