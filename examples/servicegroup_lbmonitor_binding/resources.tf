resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.13.46"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor2" {
  monitorname = "tfmonitor2"
  type        = "PING"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  lbvservers       = [citrixadc_lbvserver.tf_lbvserver.name]
  servicetype      = "HTTP"
  servicegroupmembers = [
    "192.168.33.33:80:1",
  ]

  # Don't use lbmonitor inside servicegroup when using the explicit bindings
}

# Import command example
# terraform import citrixadc_servicegroup_lbmonitor_binding.bind1 tf_servicegroup,tfmonitor1

resource "citrixadc_servicegroup_lbmonitor_binding" "bind1" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
    weight = 80
}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind2" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor2.monitorname
    weight = 20
}
