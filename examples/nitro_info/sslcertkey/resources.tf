data "citrixadc_nitro_info" "sslcertificate_info" {
    workflow = yamldecode(file("../workflows/sslcertkey.yaml"))
    primary_id = "tf_sslcertkey"
}

output "nitro_object" {
  value = data.citrixadc_nitro_info.sslcertificate_info.nitro_object
}

output "nitro_object_length" {
  value = length(data.citrixadc_nitro_info.sslcertificate_info.nitro_object)
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/var/tmp/certificate1.crt"
  key = "/var/tmp/key1.pem"
}
