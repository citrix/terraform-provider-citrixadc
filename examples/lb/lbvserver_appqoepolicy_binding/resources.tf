resource "citrixadc_lbvserver_appqoepolicy_binding" "foo" {
    name = "test-server"
    policyname = "appqoe-pol-primd"
    bindpoint = "REQUEST"
    gotopriorityexpression = "END"
    priority = 56 
}
