output "sslcertkey_name" {
  description = "The installed SSL cert/key pair name."
  value       = citrixadc_sslcertkey.kp.certkey
}

output "cert_path" {
  description = "Full path of the uploaded certificate file on the appliance."
  value       = citrixadc_sslcertkey.kp.cert
}
