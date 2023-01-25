package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwurlencodedformcontenttype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwurlencodedformcontenttypeFunc,
		Read:          readAppfwurlencodedformcontenttypeFunc,
		Delete:        deleteAppfwurlencodedformcontenttypeFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"urlencodedformcontenttypevalue": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"isregex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwurlencodedformcontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwurlencodedformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwurlencodedformcontenttypeName := d.Get("urlencodedformcontenttypevalue").(string)
	appfwurlencodedformcontenttype := appfw.Appfwurlencodedformcontenttype{
		Isregex:                        d.Get("isregex").(string),
		Urlencodedformcontenttypevalue: d.Get("urlencodedformcontenttypevalue").(string),
	}

	_, err := client.AddResource("appfwurlencodedformcontenttype", appfwurlencodedformcontenttypeName, &appfwurlencodedformcontenttype)
	if err != nil {
		return err
	}

	d.SetId(appfwurlencodedformcontenttypeName)

	err = readAppfwurlencodedformcontenttypeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwurlencodedformcontenttype but we can't read it ?? %s", appfwurlencodedformcontenttypeName)
		return nil
	}
	return nil
}

func readAppfwurlencodedformcontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwurlencodedformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwurlencodedformcontenttypeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwurlencodedformcontenttype state %s", appfwurlencodedformcontenttypeName)
	data, err := client.FindResource("appfwurlencodedformcontenttype", appfwurlencodedformcontenttypeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwurlencodedformcontenttype state %s", appfwurlencodedformcontenttypeName)
		d.SetId("")
		return nil
	}
	d.Set("isregex", data["isregex"])
	d.Set("urlencodedformcontenttypevalue", data["urlencodedformcontenttypevalue"])

	return nil

}

func deleteAppfwurlencodedformcontenttypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwurlencodedformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwurlencodedformcontenttypeName := d.Id()
	err := client.DeleteResource("appfwurlencodedformcontenttype", appfwurlencodedformcontenttypeName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}