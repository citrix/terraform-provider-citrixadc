# Since the spilloverpolicy resource is not yet available on Terraform,
# the tf_spilloverpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add spillover policy tf_spilloverpolicy -rule TRUE -action SPILLOVER

resource "citrixadc_lbvserver_spilloverpolicy_binding" "tf_lbvserver_spilloverpolicy_binding" {
	name = citrixadc_lbvserver.tf_lbvserver.name
	policyname = "tf_spilloverpolicy"
	bindpoint = "REQUEST"
	gotopriorityexpression = "END"
	invoke = false
	priority = 1
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver"
	ipv46       = "10.10.10.33"
	port        = 80
	servicetype = "HTTP"
}