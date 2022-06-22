resource "citrixadc_lbparameter" "tf_lbparam" {
    useportforhashlb = "YES"
    lbhashalgorithm = "JARH"
    lbhashfingers = 256
}
