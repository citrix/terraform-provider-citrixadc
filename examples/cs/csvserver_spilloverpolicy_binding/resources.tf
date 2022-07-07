# Since the spilloverpolicy resource is not yet available on Terraform,
# the tf_spilloverpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add spillover policy tf_spilloverpolicy -rule TRUE -action SPILLOVER

resource "citrixadc_csvserver_spilloverpolicy_binding" "tf_csvserver_spilloverpolicy_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	policyname = "tf_spilloverpolicy"
	bindpoint = "REQUEST"
	gotopriorityexpression = "END"
	invoke = false
	priority = 1
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "HTTP"
}