resource "citrixadc_nsvpxparam" "tf_vpxparam0" {
    cpuyield = "DEFAULT"
    masterclockcpu1 = "NO"
    ownernode = 0
}

resource "citrixadc_nsvpxparam" "tf_vpxparam1" {
    cpuyield = "DEFAULT"
    masterclockcpu1 = "NO"
    ownernode = 1
}
