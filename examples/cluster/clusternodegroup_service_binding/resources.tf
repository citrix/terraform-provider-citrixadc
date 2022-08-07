resource "citrixadc_clusternodegroup_service_binding" "tf_clusternodegroup_service_binding" {
  name    = "my_gslb_group"
  service = citrixadc_service.tf_service.name
}

resource "citrixadc_service" "tf_service" {
    name = "tf_service"
    servicetype = "ADNS"
    ip = "10.77.33.22"
    port = "53"
}
