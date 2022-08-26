resource "citrixadc_autoscaleaction" "action1" {
    name = "action1"
    type = "SCALE_UP"
    profilename = "profile1"
    vserver = "server1"
    parameters = "abc123"
}