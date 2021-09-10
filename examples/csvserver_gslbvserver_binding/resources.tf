resource "citrixadc_csvserver_gslbvserver_binding" "tf_csvserver_gslbvserver_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	vserver = citrixadc_gslbvserver.tf_gslbvserver.name
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	servicetype = "HTTP"
	targettype = "GSLB"
}

resource "citrixadc_gslbvserver" "tf_gslbvserver" {
	name = "tf_gslbvserver"
	servicetype = "HTTP"
}