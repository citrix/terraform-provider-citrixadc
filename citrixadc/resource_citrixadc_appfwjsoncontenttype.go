package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
	"net/url"
)

func resourceCitrixAdcAppfwjsoncontenttype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwjsoncontenttypeFunc,
		Read:          readAppfwjsoncontenttypeFunc,
		Delete:        deleteAppfwjsoncontenttypeFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"isregex": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"jsoncontenttypevalue": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwjsoncontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwjsoncontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwjsoncontenttypeName := d.Get("jsoncontenttypevalue").(string)
	appfwjsoncontenttype := appfw.Appfwjsoncontenttype{
		Isregex:              d.Get("isregex").(string),
		Jsoncontenttypevalue: appfwjsoncontenttypeName,
	}

	_, err := client.AddResource(service.Appfwjsoncontenttype.Type(), appfwjsoncontenttypeName, &appfwjsoncontenttype)
	if err != nil {
		return err
	}

	d.SetId(appfwjsoncontenttypeName)

	err = readAppfwjsoncontenttypeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwjsoncontenttype but we can't read it ?? %s", appfwjsoncontenttypeName)
		return nil
	}
	return nil
}

func readAppfwjsoncontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwjsoncontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwjsoncontenttypeName := d.Id()
	appfwjsoncontenttypeNameEscaped := url.PathEscape(url.QueryEscape(appfwjsoncontenttypeName))
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwjsoncontenttype state %s", appfwjsoncontenttypeName)
	data, err := client.FindResource(service.Appfwjsoncontenttype.Type(), appfwjsoncontenttypeNameEscaped)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwjsoncontenttype state %s", appfwjsoncontenttypeName)
		d.SetId("")
		return nil
	}
	d.Set("isregex", data["isregex"])
	d.Set("jsoncontenttypevalue", data["jsoncontenttypevalue"])

	return nil

}

func deleteAppfwjsoncontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwjsoncontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwjsoncontenttypeName := d.Id()
	appfwjsoncontenttypeNameEscaped := url.PathEscape(url.QueryEscape(appfwjsoncontenttypeName))
	err := client.DeleteResource(service.Appfwjsoncontenttype.Type(), appfwjsoncontenttypeNameEscaped)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
