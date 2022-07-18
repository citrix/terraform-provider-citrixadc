resource "citrixadc_lbroute" "tf_lbroute" {
	network = "55.0.0.0"
	netmask = "255.0.0.0"
	gatewayname = citrixadc_lbvserver.tf_lbvserver.name

	depends_on = [citrixadc_lbvserver_service_binding.tf_lbvserver_service_binding, citrixadc_nsip.nsip]
}

resource "citrixadc_nsip" "nsip" {
	ipaddress = "22.2.2.1"
	netmask   = "255.255.255.0"
}

resource "citrixadc_lbvserver_service_binding" "tf_lbvserver_service_binding" {
	name = citrixadc_lbvserver.tf_lbvserver.name
	servicename = citrixadc_service.tf_service.name
}

resource "citrixadc_service" "tf_service" {
	name = "tf_service"
	port = 65535
	ip = "22.2.2.2"
	servicetype = "ANY"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name = "tf_lbvserver"
	ipv46 = "0.0.0.0"
	servicetype = "ANY"
	lbmethod = "ROUNDROBIN"
	persistencetype = "NONE"
	clttimeout = 120
	port = 0
}