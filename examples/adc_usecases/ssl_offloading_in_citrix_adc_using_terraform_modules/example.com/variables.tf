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
variable "sslvip_config" {
  type        = map(string)
  description = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}
variable "sslvsparam" {
  type        = map(string)
  description = "SSL Virtual server parameters"
}
variable "httpvip_config" {
  type        = map(string)
  description = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}
variable "servers" {
  description = "Name and IP address of the server"
  type = map(object({
    ipaddress = string
    state     = string
    name      = string
  }))
  default = {
    "testserver11" = {
      ipaddress = "10.10.10.11"
      state     = "ENABLED"
      name      = "testserver11"
    }
    "testserver12" = {
      ipaddress = "10.10.10.12"
      state     = "ENABLED"
      name      = "testserver12"
    }
  }
}
variable "service_config" {
  description = "service configuration"
  type        = map(string)
}
variable "ssl_certkey_name" {
  type = string
}
variable "ssl_certificate_path" {
  type = string
}
variable "ssl_key_path" {
  type = string
}
variable "filename" {
  type = map(string)
}

