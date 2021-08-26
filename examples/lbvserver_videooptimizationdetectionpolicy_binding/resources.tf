# Since the videooptimizationdetectionpolicy resource is not yet available on Terraform,
# the tf_vop policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add videooptimization detectionpolicy tf_vop -rule TRUE -action DETECT_ENCRYPTED_ABR

resource "citrixadc_lbvserver_videooptimizationdetectionpolicy_binding" "tf_vopolicy_binding" {
	bindpoint = "REQUEST"
	gotopriorityexpression = "END"
	name = citrixadc_lbvserver.tf_lbvserver.name
	policyname = "tf_vop"
	priority = 1
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver"
	ipv46       = "10.10.10.33"
	port        = 80
	servicetype = "HTTP"
}