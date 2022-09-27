variable "lbvserver1_name" {
  type        = string
  description = "lb vserver1 name"
}
variable "lbvserver1_ip" {
  type        = string
  description = "lb vserver1 ip"
}
variable "lbvserver1_port" {
  type        = number
  description = "lb vserver1 Port"
}
variable "lbvserver1_servicetype" {
  description = "lb vserver1 Servicetype"
}

variable "lbvserver2_name" {
  type        = string
  description = "lb vserver2 name"
}
variable "lbvserver2_ip" {
  type        = string
  description = "lb vserver2 ip"
}
variable "lbvserver2_port" {
  type        = number
  description = "lb vserver2 Port"
}
variable "lbvserver2_servicetype" {
  description = "lb vserver2 Servicetype"
}

variable "service1_name" {
  type        = string
  description = "Service1 name"
}
variable "service1_ip" {
  type        = string
  description = "Service1 ip"
}
variable "service1_servicetype" {
  description = "Service1 Servicetype"
}
variable "service1_port" {
  type        = number
  description = "Service1 Port"
}

variable "service2_name" {
  type        = string
  description = "Service2 name"
}
variable "service2_ip" {
  type        = string
  description = "Service2 ip"
}
variable "service2_servicetype" {
  description = "Service2 Servicetype"
}
variable "service2_port" {
  type        = number
  description = "Service2 Port"
}

variable "service3_name" {
  type        = string
  description = "Service3 name"
}
variable "service3_ip" {
  type        = string
  description = "Service3 ip"
}
variable "service3_servicetype" {
  description = "Service3 Servicetype"
}
variable "service3_port" {
  type        = number
  description = "Service3 Port"
}

variable "service4_name" {
  type        = string
  description = "Service4 name"
}
variable "service4_ip" {
  type        = string
  description = "Service4 ip"
}
variable "service4_servicetype" {
  description = "Service4 Servicetype"
}
variable "service4_port" {
  type        = number
  description = "Service4 Port"
}

variable "csvserver_name" {
  type        = string
  description = "CS vserver name"
}
variable "csvserver_ipv46" {
  type        = string
  description = "CS vserver ip"
}
variable "csvserver_servicetype" {
  description = "CS vserver Servicetype"
}
variable "csvserver_port" {
  type        = number
  description = "CS vserver Port"
}

variable "cspolicy1_name" {
  type        = string
  description = "CS policy1 name"
}
variable "cspolicy1_url" {
  description = "CS Policy1 Url"
}

variable "cspolicy2_name" {
  type        = string
  description = "CS policy2 name"
}
variable "cspolicy2_url" {
  description = "CS Policy2 Url"
}

variable "cspolicy3_name" {
  type        = string
  description = "CS policy3 name"
}
variable "cspolicy3_url" {
  description = "CS Policy3 Url"
}

variable "cspolicy4_name" {
  type        = string
  description = "CS policy4 name"
}
variable "cspolicy4_url" {
  description = "CS Policy4 Url"
}

variable "sslcertkey_name" {
  type        = string
  description = "SSL CertKey Attribute"
}
variable "sslcertkey_cert" {
  description = "SSL Cert Attribute"
}
variable "sslcertkey_key" {
  description = "SSL Key Attribute"
}