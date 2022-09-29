terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}
terraform {
  required_version = ">= 0.12"
}

resource "citrixadc_systemfile" "sslcert" {
  filename     = var.filename["certname"]
  filelocation = "/flash/nsconfig/ssl"
  filecontent  = file(var.ssl_certificate_path)
}
resource "citrixadc_systemfile" "sslkey" {
  filename     = var.filename["keyname"]
  filelocation = "/flash/nsconfig/ssl"
  filecontent  = file(var.ssl_key_path)
}
resource "citrixadc_sslcertkey" "tf_certkeypair" {
  certkey = var.ssl_certkey_name
  cert    = format("%s/%s", citrixadc_systemfile.sslcert.filelocation, citrixadc_systemfile.sslcert.filename)
  key     = format("%s/%s", citrixadc_systemfile.sslkey.filelocation, citrixadc_systemfile.sslkey.filename)
  inform  = "DER"
}
