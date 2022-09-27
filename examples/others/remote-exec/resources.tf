/* this template assumes a NetScaler CPX. */
/* For other ADCs this will not work. Using a combination of NITRO API calls and local-exec is recommended */
data "template_file" "rnat_cmd" {
  count = "${length(var.rnat_config)}"
  template = "/var/netscaler/bins/cli_script.sh 'set rnat $${network} $${netmask} $${acl} $${nat} $${natip}'"

  vars {
    network = "${lookup(var.rnat_config[count.index], "network", "")}"
    netmask = "${lookup(var.rnat_config[count.index], "netmask", "")}"
    acl = "${lookup(var.rnat_config[count.index], "acl", "")}"
    natip = "${lookup(var.rnat_config[count.index], "natip", "")}"
    nat = "${lookup(map("true","-natip"), lookup(var.rnat_config[count.index], "nat", ""), "")}"
    
  }
}

data "template_file" "vlan_cmd" {
  count = "${length(var.vlan_config)}"
  template = <<EOF
  /var/netscaler/bins/cli_script.sh 'add vlan $${vlanid} '
  $${do_bind_ip}/var/netscaler/bins/cli_script.sh 'bind vlan $${vlanid} -IPAddress $${ipaddress}  $${netmask}'
  EOF

  vars {
    vlanid = "${lookup(var.vlan_config[count.index], "vlanid", "")}"
    netmask = "${lookup(var.vlan_config[count.index], "netmask", "")}"
    ipaddress = "${lookup(var.vlan_config[count.index], "ipaddress", "")}"
    do_bind_ip = "${lookup(map("false","#"), lookup(var.vlan_config[count.index], "ipaddress", "false"), "")}"
  }
}

resource "null_resource" "rnat" {
  count = "${length(var.rnat_config)}"

  connection {
    type = "ssh"
    user = "${var.ns["login"]}"
    password = "${var.ns["password"]}"
    host = "${var.ns["ip"]}"
    port = "${var.ns["port"]}"
  }

  provisioner "remote-exec" {
      inline = [
         "${data.template_file.rnat_cmd.*.rendered[count.index]}"
      ]
  }
}

resource "null_resource" "nsconfig" {
  connection {
    type = "ssh"
    user = "${var.ns["login"]}"
    password = "${var.ns["password"]}"
    host = "${var.ns["ip"]}"
    port = "${var.ns["port"]}"
  }

  provisioner "remote-exec" {
      inline = [
         "/var/netscaler/bins/cli_script.sh 'set ns config -IPAddress ${var.nsconfig["ipaddress"]} -netmask ${var.nsconfig["netmask"]}'"
      ]
  }
}

resource "null_resource" "vlans" {
  count = "${length(var.vlan_config)}"
  connection {
    type = "ssh"
    user = "${var.ns["login"]}"
    password = "${var.ns["password"]}"
    host = "${var.ns["ip"]}"
    port = "${var.ns["port"]}"
  }

  provisioner "remote-exec" {
      inline = [
         "${data.template_file.vlan_cmd.*.rendered[count.index]}"
      ]
  }
}

output "rnatcmds" {
   value = "${data.template_file.rnat_cmd.*.rendered}"
}

output "vlancmds" {
   value = "${data.template_file.vlan_cmd.*.rendered}"
}
