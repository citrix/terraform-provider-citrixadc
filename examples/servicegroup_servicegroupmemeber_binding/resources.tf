resource "citrixadc_server" "tf_server" {
    name = "tf_server"
    domain = "example.com"
    querytype = "SRV"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
  autoscale = "DNS"
}

resource "citrixadc_servicegroup_servicegroupmember_binding" "tf_binding" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    servername = citrixadc_server.tf_server.name
}

resource "citrixadc_servicegroup_servicegroupmember_binding" "tf_binding2" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    ip = "10.78.22.33"
    port = 80
}
