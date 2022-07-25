resource "citrixadc_snmpcommunity" "tf_snmpcommunity" {
  communityname = "test_community"
  permissions   = "GET_BULK"
}