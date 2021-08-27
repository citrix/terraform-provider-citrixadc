resource "citrixadc_sslservicegroup_ecccurve_binding" "tf_sslservicegroup_ecccurve_binding" {
	ecccurvename = "P_256"
	servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
	servicegroupname = "tf_servicegroup"
	servicetype = "SSL"
}