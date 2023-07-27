package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationepaaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationepaactionFunc,
		Read:          readAuthenticationepaactionFunc,
		Update:        updateAuthenticationepaactionFunc,
		Delete:        deleteAuthenticationepaactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"csecexpr": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"defaultepagroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deletefiles": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"killprocess": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quarantinegroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationepaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationepaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationepaactionName := d.Get("name").(string)
	authenticationepaaction := authentication.Authenticationepaaction{
		Csecexpr:        d.Get("csecexpr").(string),
		Defaultepagroup: d.Get("defaultepagroup").(string),
		Deletefiles:     d.Get("deletefiles").(string),
		Killprocess:     d.Get("killprocess").(string),
		Name:            d.Get("name").(string),
		Quarantinegroup: d.Get("quarantinegroup").(string),
	}

	_, err := client.AddResource("authenticationepaaction", authenticationepaactionName, &authenticationepaaction)
	if err != nil {
		return err
	}

	d.SetId(authenticationepaactionName)

	err = readAuthenticationepaactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationepaaction but we can't read it ?? %s", authenticationepaactionName)
		return nil
	}
	return nil
}

func readAuthenticationepaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationepaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationepaactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationepaaction state %s", authenticationepaactionName)
	data, err := client.FindResource("authenticationepaaction", authenticationepaactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationepaaction state %s", authenticationepaactionName)
		d.SetId("")
		return nil
	}
	d.Set("csecexpr", data["csecexpr"])
	d.Set("defaultepagroup", data["defaultepagroup"])
	d.Set("deletefiles", data["deletefiles"])
	d.Set("killprocess", data["killprocess"])
	d.Set("name", data["name"])
	d.Set("quarantinegroup", data["quarantinegroup"])

	return nil

}

func updateAuthenticationepaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationepaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationepaactionName := d.Get("name").(string)

	authenticationepaaction := authentication.Authenticationepaaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("csecexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Csecexpr has changed for authenticationepaaction %s, starting update", authenticationepaactionName)
		authenticationepaaction.Csecexpr = d.Get("csecexpr").(string)
		hasChange = true
	}
	if d.HasChange("defaultepagroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultepagroup has changed for authenticationepaaction %s, starting update", authenticationepaactionName)
		authenticationepaaction.Defaultepagroup = d.Get("defaultepagroup").(string)
		hasChange = true
	}
	if d.HasChange("deletefiles") {
		log.Printf("[DEBUG]  citrixadc-provider: Deletefiles has changed for authenticationepaaction %s, starting update", authenticationepaactionName)
		authenticationepaaction.Deletefiles = d.Get("deletefiles").(string)
		hasChange = true
	}
	if d.HasChange("killprocess") {
		log.Printf("[DEBUG]  citrixadc-provider: Killprocess has changed for authenticationepaaction %s, starting update", authenticationepaactionName)
		authenticationepaaction.Killprocess = d.Get("killprocess").(string)
		hasChange = true
	}
	if d.HasChange("quarantinegroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Quarantinegroup has changed for authenticationepaaction %s, starting update", authenticationepaactionName)
		authenticationepaaction.Quarantinegroup = d.Get("quarantinegroup").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationepaaction", authenticationepaactionName, &authenticationepaaction)
		if err != nil {
			return fmt.Errorf("Error updating authenticationepaaction %s", authenticationepaactionName)
		}
	}
	return readAuthenticationepaactionFunc(d, meta)
}

func deleteAuthenticationepaactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationepaactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationepaactionName := d.Id()
	err := client.DeleteResource("authenticationepaaction", authenticationepaactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
