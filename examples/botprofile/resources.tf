resource "citrixadc_botprofile" "tf_botprofile_name" {
  name     		 = "botprofile_name"
  comment		 = "Botprofile comment 1"
  bot_enable_white_list  = "ON"
  devicefingerprint	 = "ON"
}
