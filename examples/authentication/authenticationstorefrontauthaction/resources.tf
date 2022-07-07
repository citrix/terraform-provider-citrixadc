resource "citrixadc_authenticationstorefrontauthaction" "tf_storefront" {
  name                       = "tf_storefront"
  serverurl                  = "http://www.example.com/"
  domain                     = "domainname"
  defaultauthenticationgroup = "group_name"
}