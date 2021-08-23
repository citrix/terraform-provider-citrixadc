# Since the feopolicy resource is not yet available on Terraform,
# the tf_feopolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add feo policy tf_feopolicy TRUE BASIC

resource "citrixadc_lbvserver_feopolicy_binding" "tf_lbvserver_feopolicy_binding" {
	bindpoint = "REQUEST"
	gotopriorityexpression = "END"
	name = citrixadc_lbvserver.tf_lbvserver.name
	policyname = "tf_feopolicy"
	priority = 1  
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver"
	ipv46       = "10.10.10.33"
	port        = 80
	servicetype = "HTTP"
}