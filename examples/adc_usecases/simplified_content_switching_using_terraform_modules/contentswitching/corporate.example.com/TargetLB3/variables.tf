variable "username" {
  description = "ADC administrator username"
  type        = string
  sensitive   = true
}
variable "password" {
  description = "ADC administrator password"
  type        = string
  sensitive   = true
}
variable "targetlb3_config" {
  type        = map(string)
  description = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}
variable "servers" {
  description = "Name and IP address of the server"
  type = map(object({
    ipaddress = string
    state     = string
    name      = string
    comment   = string
  }))
  default = {
    "testserver30" = {
      ipaddress = "192.168.50.30"
      state     = "ENABLED"
      name      = "testserver30"
      comment   = "webapp3"
    }
    "testserver31" = {
      ipaddress = "192.168.50.31"
      state     = "ENABLED"
      name      = "testserver31"
      comment   = "webapp3"
    }
  }
}
variable "service_config" {
  description = "service configuration"
  type        = map(string)
}
