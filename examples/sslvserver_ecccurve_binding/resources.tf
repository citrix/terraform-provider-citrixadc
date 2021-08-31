resource "citrixadc_sslvserver_ecccurve_binding" "tf_sslvserver_ecccurve_binding" {
	ecccurvename = "P_256"
	vservername = citrixadc_lbvserver.tf_sslvserver.name
	
}

resource "citrixadc_lbvserver" "tf_sslvserver" {
	name        = "tf_sslvserver"
	servicetype = "SSL"
}