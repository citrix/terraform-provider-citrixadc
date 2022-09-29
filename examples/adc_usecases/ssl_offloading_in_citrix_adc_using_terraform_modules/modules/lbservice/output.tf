output "lbservices" {
    value = [ for service in citrixadc_service.tf_service: service.name ]
}
