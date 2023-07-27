package citrixadc

import (
	"log"

	"github.com/citrix/adc-nitro-go/resource/config/authorization"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCitrixAdcAuthorizationpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthorizationpolicylabelFunc,
		Read:          readAuthorizationpolicylabelFunc,
		Delete:        deleteAuthorizationpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthorizationpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthorizationpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicylabelName := d.Get("labelname").(string)
	authorizationpolicylabel := authorization.Authorizationpolicylabel{
		Labelname: authorizationpolicylabelName,
		Newname:   d.Get("newname").(string),
	}

	_, err := client.AddResource(service.Authorizationpolicylabel.Type(), authorizationpolicylabelName, &authorizationpolicylabel)
	if err != nil {
		return err
	}

	d.SetId(authorizationpolicylabelName)

	err = readAuthorizationpolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authorizationpolicylabel but we can't read it ?? %s", authorizationpolicylabelName)
		return nil
	}
	return nil
}

func readAuthorizationpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthorizationpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authorizationpolicylabel state %s", authorizationpolicylabelName)
	data, err := client.FindResource(service.Authorizationpolicylabel.Type(), authorizationpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authorizationpolicylabel state %s", authorizationpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("newname", data["newname"])

	return nil

}

func deleteAuthorizationpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthorizationpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	authorizationpolicylabelName := d.Id()
	err := client.DeleteResource(service.Authorizationpolicylabel.Type(), authorizationpolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
