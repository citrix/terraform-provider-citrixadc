# -----------------------------------------------------------------------------
# NetScaler HA Pair
# -----------------------------------------------------------------------------

variable "node1_nsip" {
  description = "NSIP of the first NetScaler node"
  type        = string
}

variable "node2_nsip" {
  description = "NSIP of the second NetScaler node"
  type        = string
}

variable "adc_admin_username" {
  description = "Admin username for both nodes"
  type        = string
  default     = "nsroot"
}

variable "adc_admin_password" {
  description = "Admin password for both nodes"
  type        = string
  sensitive   = true
}

# -----------------------------------------------------------------------------
# Load Balancer Configuration
# -----------------------------------------------------------------------------

variable "lb_vserver_name" {
  description = "Name of the LB virtual server"
  type        = string
  default     = "app_lb_vserver"
}

variable "lb_vip" {
  description = "Virtual IP for the load balancer"
  type        = string
}

variable "lb_port" {
  description = "Port for the LB virtual server"
  type        = number
  default     = 80
}

variable "lb_service_type" {
  description = "Service type (HTTP, HTTPS, TCP, etc.)"
  type        = string
  default     = "HTTP"
}

variable "lb_method" {
  description = "Load balancing method"
  type        = string
  default     = "ROUNDROBIN"
}

# -----------------------------------------------------------------------------
# Backend Servers
# -----------------------------------------------------------------------------

variable "backend_servers" {
  description = "List of backend servers to load balance"
  type = list(object({
    name = string
    ip   = string
    port = number
  }))
  default = [
    {
      name = "web_server1"
      ip   = "10.0.1.10"
      port = 80
    },
    {
      name = "web_server2"
      ip   = "10.0.1.11"
      port = 80
    },
  ]
}
