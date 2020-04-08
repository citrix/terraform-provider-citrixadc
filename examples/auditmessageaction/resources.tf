resource "citrixadc_auditmessageaction" "tf_msgaction" {
    name = "tf_msgaction"
    loglevel = "NOTICE"
    stringbuilderexpr = "\"hello bye\""
    logtonewnslog = "YES"
}
