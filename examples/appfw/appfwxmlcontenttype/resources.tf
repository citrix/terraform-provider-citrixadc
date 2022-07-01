resource "citrixadc_appfwxmlcontenttype" "demo_appfwxmlcontenttype" {
  xmlcontenttypevalue = "demo.*test" 
  isregex = "REGEX"
}
