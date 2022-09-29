
variable "servers" {
  description = "Name and IP address of the server"
  type = map(object({
    ipaddress = string
    state     = string
    name      = string
  }))
}
variable "service_config" {
  description = "service configuration"
  type        = map(string)
}
