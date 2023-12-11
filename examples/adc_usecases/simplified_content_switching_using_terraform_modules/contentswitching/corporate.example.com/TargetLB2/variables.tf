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
variable "targetlb2_config" {
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
    "testserver20" = {
      ipaddress = "192.168.50.20"
      state     = "ENABLED"
      name      = "testserver20"
      comment   = "webapp2"
    }
    "testserver21" = {
      ipaddress = "192.168.50.21"
      state     = "ENABLED"
      name      = "testserver21"
      comment   = "webapp2"
    }
  }
}
variable "service_config" {
  description = "service configuration"
  type        = map(string)
}
