cluster_id = 1
clip       = "10.102.201.228"

adc_admin_username = "nsroot"
adc_admin_password = "ConfigADC#123"

netscaler_attributes = {
  "node1" = {
    node_id   = 0
    backplane = "0/1/1"
  }
  "node2" = {
    node_id   = 1
    backplane = "1/1/1"
  }
  "node3" = {
    node_id   = 2
    backplane = "2/1/1"
  }
}

nsips = {
  "node1" = "10.102.201.213"
  "node2" = "10.102.201.42"
  "node3" = "10.102.201.222"
}
