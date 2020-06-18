resource "citrixadc_policyexpression" "tf_advanced_policyexpression" {
    name = "tf_advanced_policyexrpession"
    value = "HTTP.REQ.URL.SUFFIX.EQ(\"cgi\")"
    comment = "comment"
}

resource "citrixadc_policyexpression" "tf_classic_policyexpression" {
    name = "tf_classic_policyexrpession"
    value = "HEADER Cookie EXISTS"
    clientsecuritymessage = "security message"
}
