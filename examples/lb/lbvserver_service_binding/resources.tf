resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
    name = "tf_service"
    ip = "192.168.43.33"
    servicetype  = "HTTP"
    port = 80
}

resource "citrixadc_lbvserver_service_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicename = citrixadc_service.tf_service.name
  weight = 1
}
