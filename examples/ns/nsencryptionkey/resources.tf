resource "citrixadc_nsencryptionkey" "tf_encryptionkey" {
  name     = "tf_encryptionkey"
  method   = "AES256"
  keyvalue = "26ea5537b7e0746089476e5658f9327c0b10c3b4778c673a5b38cee182874711"
  padding  = "DEFAULT"
  iv       = "c2bf0b2e15c15004d6b14bcdc7e5e365"
  comment  = "Testing"
}