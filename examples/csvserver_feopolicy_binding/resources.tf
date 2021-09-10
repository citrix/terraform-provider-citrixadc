# Since the feopolicy resource is not yet available on Terraform,
# the tf_feopolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add feo policy tf_feopolicy TRUE BASIC

resource "citrixadc_csvserver_feopolicy_binding" "tf_csvserver_feopolicy_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	policyname = "tf_feopolicy"
	bindpoint = "REQUEST"
	gotopriorityexpression = "END"
	priority = 1  
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "HTTP"
}