config_done = false

ns = {
    ip = "localhost"
    port = "32770"
    password = "nsroot"
}

rnat_config = [
  {
    network = "192.168.33.0"
    netmask = "255.255.255.0"
  },
  {
    network = "192.168.35.0",
    netmask = "255.255.255.0"
    nat = "true"
    natip   = "172.17.0.2"
  },
  {
    acl = "someacl"
    nat = "true"
    natip   = "172.17.0.2"
  }
]

nsconfig = {
    ipaddress = "192.168.33.3"
    netmask = "255.255.255.0"
}

vlan_config = [
  {
    vlanid = "400"
  },
  {
    vlanid = "500"
    ipaddress = "172.17.20.2"
    netmask = "255.255.255.0"
  }
]
