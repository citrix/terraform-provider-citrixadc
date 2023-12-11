variable "username" {
  type        = string
  sensitive   = true
  description = "ADC administrator username"
}
variable "password" {
  type        = string
  sensitive   = true
  description = "ADC administrator password"
}

variable "httpcsvip_config" {
  type        = map(string)
  description = "Describes the friendly name (key=csname),VIP(key=csvip), port(key=port) and service type (key=servicetype) for the CS"
}

variable "sslcsvip_config" {
  type        = map(string)
  description = "Describes the friendly name (key=csname),VIP(key=csvip), port(key=port) and service type (key=servicetype) for the CS"
}

variable "httpvip_config" {
  type        = map(string)
  description = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}
/*
variable "sslvip_config" {
  type        = map(string)
  description = "Describes the friendly name (key=lbname),VIP(key=vip), port(key=port) and service type (key=servicetype) for the LB"
}*/
variable "servers" {
  description = "Name and IP address of the server"
  type = map(object({
    ipaddress = string
    state     = string
    name      = string
  }))
  default = {
    "testservermain01" = {
      ipaddress = "192.168.50.50"
      state     = "ENABLED"
      name      = "testservermain01"
    }
    "testservermain02" = {
      ipaddress = "192.168.50.51"
      state     = "ENABLED"
      name      = "testservermain02"
    }
  }
}
variable "service_config" {
  type        = map(string)
  description = "plain text (HTTP or TCP) backend service configuration"
}
/*
variable "monitor_config" {
  description = "LB health monitor configuration"
  type        = map(string)
}*/
variable "cs_action1" {
  type        = map(string)
  description = "content switching action"
}
variable "cs_pol1" {
  type        = map(string)
  description = "content switching policy"
}
variable "cs_action2" {
  type        = map(string)
  description = "content switching action"
}
variable "cs_pol2" {
  type        = map(string)
  description = "content switching policy"
}
variable "cs_action3" {
  type        = map(string)
  description = "content switching action"
}
variable "cs_pol3" {
  type        = map(string)
  description = "content switching policy"
}
