resource "citrixadc_systemuser" "user" {
    username = "george"
    password = "12345"
    timeout = 900

    cmdpolicybinding {
        policyname = "superuser"
        priority = 100
    }
}
