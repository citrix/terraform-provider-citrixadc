node1_nsip         = "10.102.201.42"
node2_nsip         = "10.102.201.213"
adc_admin_username = "nsroot"
adc_admin_password = "###"

lb_vip  = "10.102.201.100"
lb_port = 80

backend_servers = [
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
