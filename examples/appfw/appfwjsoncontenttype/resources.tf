resource "citrixadc_appfwjsoncontenttype" "demo_appfwjsoncontenttype" {
  jsoncontenttypevalue = "demo.*test" 
  isregex = "REGEX"
}
