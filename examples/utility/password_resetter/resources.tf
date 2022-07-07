resource "citrixadc_password_resetter" "tf_resetter" {
    username = "nsroot"
    password = "nsroot"
    new_password = "newnsroot"
}
