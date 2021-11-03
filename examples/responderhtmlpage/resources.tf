resource "citrixadc_systemfile" "tf_html_page" {
    filename = "tf_html_page.html"
    filelocation = "/var/tmp"
    filecontent = "<h1>Hello Responder</h1>"
}

resource "citrixadc_responderhtmlpage" "tf_responder_page" {
    name = "tf_responder_page"
    src = "local://tf_html_page.html"
    depends_on = [citrixadc_systemfile.tf_html_page]
}
