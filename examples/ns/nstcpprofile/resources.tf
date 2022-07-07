resource "citrixadc_nstcpprofile" "test_profile" {
    name = "test_tf_profile"
    ws = "ENABLED"
    ackaggregation = "DISABLED"
    #mpcapablecbit = "ENABLED"
}
