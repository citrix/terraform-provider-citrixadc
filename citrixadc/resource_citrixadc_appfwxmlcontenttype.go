package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwxmlcontenttype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwxmlcontenttypeFunc,
		Read:          readAppfwxmlcontenttypeFunc,
		Delete:        deleteAppfwxmlcontenttypeFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"isregex": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"xmlcontenttypevalue": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwxmlcontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwxmlcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlcontenttypeName := d.Get("xmlcontenttypevalue").(string)
	appfwxmlcontenttype := appfw.Appfwxmlcontenttype{
		Isregex:             d.Get("isregex").(string),
		Xmlcontenttypevalue: appfwxmlcontenttypeName,
	}

	_, err := client.AddResource(service.Appfwxmlcontenttype.Type(), appfwxmlcontenttypeName, &appfwxmlcontenttype)
	if err != nil {
		return err
	}

	d.SetId(appfwxmlcontenttypeName)

	err = readAppfwxmlcontenttypeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwxmlcontenttype but we can't read it ?? %s", appfwxmlcontenttypeName)
		return nil
	}
	return nil
}

func readAppfwxmlcontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwxmlcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlcontenttypeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwxmlcontenttype state %s", appfwxmlcontenttypeName)
	data, err := client.FindResource(service.Appfwxmlcontenttype.Type(), appfwxmlcontenttypeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwxmlcontenttype state %s", appfwxmlcontenttypeName)
		d.SetId("")
		return nil
	}
	d.Set("isregex", data["isregex"])
	d.Set("xmlcontenttypevalue", data["xmlcontenttypevalue"])

	return nil

}

func deleteAppfwxmlcontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwxmlcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlcontenttypeName := d.Id()
	err := client.DeleteResource(service.Appfwxmlcontenttype.Type(), appfwxmlcontenttypeName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
