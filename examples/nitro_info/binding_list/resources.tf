data "citrixadc_nitro_info" "sample" {
    workflow = yamldecode(file("../workflows/sslcertkey_sslvserver_binding.yaml"))
    primary_id = citrixadc_sslcertkey.tf_sslcertkey.certkey
}

output "list_output" {
    value = data.citrixadc_nitro_info.sample.nitro_list
}

output "object_output" {
    value = [ for item in data.citrixadc_nitro_info.sample.nitro_list: item.object["servername"] ]
}


resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/var/tmp/certificate1.crt"
  key = "/var/tmp/key1.pem"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.44"
  name        = "tf_lbvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}


resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  sslprofile = "ns_default_ssl_profile_frontend"
  port        = 443
  servicetype = "SSL"
}

resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
    vservername = citrixadc_lbvserver.tf_lbvserver.name
    certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
}

resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding2" {
    vservername = citrixadc_csvserver.tf_csvserver.name
    certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
}
