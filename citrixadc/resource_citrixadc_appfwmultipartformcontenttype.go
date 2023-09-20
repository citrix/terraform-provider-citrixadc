package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
	"net/url"
)

func resourceCitrixAdcAppfwmultipartformcontenttype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwmultipartformcontenttypeFunc,
		Read:          readAppfwmultipartformcontenttypeFunc,
		Delete:        deleteAppfwmultipartformcontenttypeFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"multipartformcontenttypevalue": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"isregex": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwmultipartformcontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwmultipartformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwmultipartformcontenttypeName := d.Get("multipartformcontenttypevalue").(string)
	appfwmultipartformcontenttype := appfw.Appfwmultipartformcontenttype{
		Isregex:                       d.Get("isregex").(string),
		Multipartformcontenttypevalue: d.Get("multipartformcontenttypevalue").(string),
	}

	_, err := client.AddResource("appfwmultipartformcontenttype", appfwmultipartformcontenttypeName, &appfwmultipartformcontenttype)
	if err != nil {
		return err
	}

	d.SetId(appfwmultipartformcontenttypeName)

	err = readAppfwmultipartformcontenttypeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwmultipartformcontenttype but we can't read it ?? %s", appfwmultipartformcontenttypeName)
		return nil
	}
	return nil
}

func readAppfwmultipartformcontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwmultipartformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwmultipartformcontenttypeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwmultipartformcontenttype state %s", appfwmultipartformcontenttypeName)
	appfwmultipartformcontenttypeNameEscaped := url.PathEscape(url.QueryEscape(appfwmultipartformcontenttypeName))
	data, err := client.FindResource("appfwmultipartformcontenttype", appfwmultipartformcontenttypeNameEscaped)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwmultipartformcontenttype state %s", appfwmultipartformcontenttypeNameEscaped)
		d.SetId("")
		return nil
	}
	d.Set("isregex", data["isregex"])
	d.Set("multipartformcontenttypevalue", data["multipartformcontenttypevalue"])

	return nil

}

func deleteAppfwmultipartformcontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwmultipartformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwmultipartformcontenttypeName := d.Id()
	appfwmultipartformcontenttypeNameEscaped := url.PathEscape(url.QueryEscape(appfwmultipartformcontenttypeName))
	err := client.DeleteResource("appfwmultipartformcontenttype", appfwmultipartformcontenttypeNameEscaped)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
