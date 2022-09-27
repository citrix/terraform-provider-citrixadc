variable "ns" {
    description = "NetScaler information to target SSH connections"
    type = "map"
    default = {
        login = "nsroot"
        port = "22"
    }
}

variable "config_done" {
    default = false
    "description" = "whether the config is done. set to false to re-apply"
}

variable "rnat_config" {
    type = "list"
    "description" = "all the rnat rules"
}

variable "vlan_config" {
    type = "list"
    "description" = "all the vlans"
}

variable "nsconfig" {
    type = "map"
    "description" = "Basic IP address information"
}
