resource "citrixadc_csvserver_analyticsprofile_binding" "tf_csvserver_analyticsprofile_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	analyticsprofile = "ns_analytics_global_profile"
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "1.1.1.2"
	port = 80
	servicetype = "HTTP"
}