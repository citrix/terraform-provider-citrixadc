variable "cluster_id" {
  description = "Cluster ID"
  type        = number
  default     = 1
}

variable "clip" {
  description = "Cluster IP address"
  type        = string
  default     = "10.102.201.228"
}

variable "adc_admin_username" {
  description = "Admin username for all nodes"
  type        = string
  default     = "nsroot"
}

variable "adc_admin_password" {
  description = "Admin password for all nodes"
  type        = string
  sensitive   = true
  default     = ""
}

variable "netscaler_attributes" {
  description = "Map of node name to attributes including node_id and backplane"
  type = map(object({
    node_id   = number
    backplane = string
  }))
  default = {
    "node1" = {
      node_id   = 0
      backplane = "0/1/1"
    }
    "node2" = {
      node_id   = 1
      backplane = "1/1/1"
    }
    "node2" = {
      node_id   = 2
      backplane = "2/1/1"
    }
   
  }
}

variable "nsips" {
  description = "Map of node name to NSIP address"
  type        = map(string)
  default = {
    "node1" = "10.102.201.213"
    "node2" = "10.102.201.42"
    "node3" = "10.102.201.22"
  }
}
 