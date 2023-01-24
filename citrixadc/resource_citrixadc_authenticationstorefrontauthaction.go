package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationstorefrontauthaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationstorefrontauthactionFunc,
		Read:          readAuthenticationstorefrontauthactionFunc,
		Update:        updateAuthenticationstorefrontauthactionFunc,
		Delete:        deleteAuthenticationstorefrontauthactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"serverurl": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"defaultauthenticationgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationstorefrontauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationstorefrontauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationstorefrontauthactionName := d.Get("name").(string)
	authenticationstorefrontauthaction := authentication.Authenticationstorefrontauthaction{
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Domain:                     d.Get("domain").(string),
		Name:                       d.Get("name").(string),
		Serverurl:                  d.Get("serverurl").(string),
	}

	_, err := client.AddResource("authenticationstorefrontauthaction", authenticationstorefrontauthactionName, &authenticationstorefrontauthaction)
	if err != nil {
		return err
	}

	d.SetId(authenticationstorefrontauthactionName)

	err = readAuthenticationstorefrontauthactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationstorefrontauthaction but we can't read it ?? %s", authenticationstorefrontauthactionName)
		return nil
	}
	return nil
}

func readAuthenticationstorefrontauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationstorefrontauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationstorefrontauthactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationstorefrontauthaction state %s", authenticationstorefrontauthactionName)
	data, err := client.FindResource("authenticationstorefrontauthaction", authenticationstorefrontauthactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationstorefrontauthaction state %s", authenticationstorefrontauthactionName)
		d.SetId("")
		return nil
	}
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("domain", data["domain"])
	d.Set("name", data["name"])
	d.Set("serverurl", data["serverurl"])

	return nil

}

func updateAuthenticationstorefrontauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationstorefrontauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationstorefrontauthactionName := d.Get("name").(string)

	authenticationstorefrontauthaction := authentication.Authenticationstorefrontauthaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationstorefrontauthaction %s, starting update", authenticationstorefrontauthactionName)
		authenticationstorefrontauthaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG]  citrixadc-provider: Domain has changed for authenticationstorefrontauthaction %s, starting update", authenticationstorefrontauthactionName)
		authenticationstorefrontauthaction.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("serverurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverurl has changed for authenticationstorefrontauthaction %s, starting update", authenticationstorefrontauthactionName)
		authenticationstorefrontauthaction.Serverurl = d.Get("serverurl").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationstorefrontauthaction", authenticationstorefrontauthactionName, &authenticationstorefrontauthaction)
		if err != nil {
			return fmt.Errorf("Error updating authenticationstorefrontauthaction %s", authenticationstorefrontauthactionName)
		}
	}
	return readAuthenticationstorefrontauthactionFunc(d, meta)
}

func deleteAuthenticationstorefrontauthactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationstorefrontauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationstorefrontauthactionName := d.Id()
	err := client.DeleteResource("authenticationstorefrontauthaction", authenticationstorefrontauthactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
