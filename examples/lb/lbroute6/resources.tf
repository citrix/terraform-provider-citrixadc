resource "citrixadc_nsip6" "tf1_nsip6" {
    ipv6address = "22::1/64"
	vserver = "DISABLED"
}

resource "citrixadc_nsip6" "tf2_nsip6" {
    ipv6address = "33::1/64"
	vserver = "DISABLED"
}

resource "citrixadc_nsip6" "tf3_nsip6" {
    ipv6address = "44::1/64"
	vserver = "DISABLED"
}

resource "citrixadc_lbvserver" "llb6" {
    name = "llb6"
    servicetype = "ANY"
    persistencetype = "NONE"
    lbmethod = "ROUNDROBIN"
}

resource "citrixadc_service" "r4" {
    name = "r4"
    ip = "22::2"
    servicetype  = "ANY"
    port = 65535

    depends_on = [citrixadc_nsip6.tf1_nsip6]
}

resource "citrixadc_service" "r5" {
    name = "r5"
    ip = "33::2"
    servicetype  = "ANY"
    port = 65535

    depends_on = [citrixadc_nsip6.tf2_nsip6]

}

resource "citrixadc_service" "r6" {
    name = "r6"
    ip = "44::2"
    servicetype  = "ANY"
    port = 65535

    depends_on = [citrixadc_nsip6.tf3_nsip6]

}

resource "citrixadc_lbvserver_service_binding" "tf_binding4" {
  name = citrixadc_lbvserver.llb6.name
  servicename = citrixadc_service.r4.name
  weight = 10
}

resource "citrixadc_lbvserver_service_binding" "tf_binding5" {
  name = citrixadc_lbvserver.llb6.name
  servicename = citrixadc_service.r5.name
  weight = 10
}

resource "citrixadc_lbvserver" "llb7" {
    name = "llb7"
    servicetype = "ANY"
    persistencetype = "NONE"
    lbmethod = "ROUNDROBIN"
}

resource "citrixadc_lbvserver_service_binding" "tf_binding76" {
  name = citrixadc_lbvserver.llb7.name
  servicename = citrixadc_service.r6.name
  weight = 10
}

resource "citrixadc_lbroute6" "demo_route6" {
    network = "66::/64"
    gatewayname = citrixadc_lbvserver.llb6.name
}

resource "citrixadc_lbroute6" "demo_route7" {
    network = "68::/64"
    gatewayname = citrixadc_lbvserver.llb7.name
}
