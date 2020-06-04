resource "citrixadc_routerdynamicrouting" "tf_dynamicrouting" {
    commandlines = [
        "router bgp 101",
        "neighbor 192.168.5.1 remote-as 100",
        "redistribute kernel",
    ]
}
