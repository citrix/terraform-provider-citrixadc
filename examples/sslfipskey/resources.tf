resource "citrixadc_sslfipskey" "demo_sslfipskey" {
    fipskeyname = "f1"
    keytype = "ECDSA"
    curve = "P_256"
}
