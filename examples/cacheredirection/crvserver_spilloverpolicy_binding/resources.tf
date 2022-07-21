resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}

resource "citrixadc_crvserver_spilloverpolicy_binding" "crvserver_spilloverpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = "tf_spilloverpolicy"
    bindpoint = "REQUEST"
    gotopriorityexpression = "END"
    invoke = false
    priority = 1
}
