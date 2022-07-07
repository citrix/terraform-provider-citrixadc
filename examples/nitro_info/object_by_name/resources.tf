data "citrixadc_nitro_info" "server_info" {
    workflow = yamldecode(file("../workflows/server.yaml"))
    primary_id = "tf_server"
}

output "nitro_object" {
  value = data.citrixadc_nitro_info.server_info.nitro_object
}

output "nitro_object_length" {
  value = length(data.citrixadc_nitro_info.server_info.nitro_object)
}

resource "citrixadc_server" "server" {
  name      = "tf_server"
  ipaddress = "10.22.22.22"
}
