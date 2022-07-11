resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_lbvserver" "foo_lbvserver" {
  name        = "test_lbvserver"
  servicetype = "HTTP"
  ipv46       = "192.0.0.0"
  port        = 8000
  comment     = "hello"
}
resource "citrixadc_service" "tf_service" {
  lbvserver   = citrixadc_lbvserver.foo_lbvserver.name
  name        = "tf_service"
  port        = 8081
  ip          = "10.33.4.5"
  servicetype = "HTTP"
  cachetype   = "TRANSPARENT"
}
resource "citrixadc_crvserver_lbvserver_binding" "crvserver_lbvserver_binding" {
  name      = citrixadc_crvserver.crvserver.name
  lbvserver = citrixadc_lbvserver.foo_lbvserver.name
  depends_on = [
    citrixadc_service.tf_service
  ]
}