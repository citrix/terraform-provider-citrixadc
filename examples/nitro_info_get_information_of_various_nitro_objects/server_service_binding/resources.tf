data "citrixadc_nitro_info" "service_bindings" {
    workflow = yamldecode(file("../workflows/server_service_binding.yaml"))
    primary_id = "tf_server"
}

output "list_output" {
    value = data.citrixadc_nitro_info.service_bindings.nitro_list
}

output "object_output" {
    value = [ for item in data.citrixadc_nitro_info.service_bindings.nitro_list: item.object["servicename"] ]
}

resource "citrixadc_server" "tf_server" {
  name      = "tf_server"
  ipaddress = "192.168.2.2"
}

resource "citrixadc_service" "tf_service" {
    servicetype = "HTTP"
    name = "tf_service"
    servername = citrixadc_server.tf_server.name
    port = "80"

    state = "ENABLED"
}

resource "citrixadc_service" "tf_service2" {
    servicetype = "HTTP"
    name = "tf_service2"
    servername = citrixadc_server.tf_server.name
    port = "90"

    state = "ENABLED"
}
