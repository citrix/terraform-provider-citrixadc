data "citrixadc_nitro_info" "service_info" {
    workflow = yamldecode(file("../workflows/service.yaml"))
    primary_id = "tf_service"
}

output "nitro_object" {
  value = data.citrixadc_nitro_info.service_info.nitro_object
}

output "nitro_object_length" {
  value = length(data.citrixadc_nitro_info.service_info.nitro_object)
}

resource "citrixadc_service" "service" {
    servicetype = "HTTP"
    name = "tf_service"
    ipaddress = "10.77.33.22"
    ip = "10.77.33.22"
    port = "80"

    state = "ENABLED"
}
