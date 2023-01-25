package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationloginschemapolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationloginschemapolicyFunc,
		Read:          readAuthenticationloginschemapolicyFunc,
		Update:        updateAuthenticationloginschemapolicyFunc,
		Delete:        deleteAuthenticationloginschemapolicyFunc,
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
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationloginschemapolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationloginschemapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationloginschemapolicyName := d.Get("name").(string)
	authenticationloginschemapolicy := authentication.Authenticationloginschemapolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Newname:     d.Get("newname").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource(service.Authenticationloginschemapolicy.Type(), authenticationloginschemapolicyName, &authenticationloginschemapolicy)
	if err != nil {
		return err
	}

	d.SetId(authenticationloginschemapolicyName)

	err = readAuthenticationloginschemapolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationloginschemapolicy but we can't read it ?? %s", authenticationloginschemapolicyName)
		return nil
	}
	return nil
}

func readAuthenticationloginschemapolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationloginschemapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationloginschemapolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationloginschemapolicy state %s", authenticationloginschemapolicyName)
	data, err := client.FindResource(service.Authenticationloginschemapolicy.Type(), authenticationloginschemapolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationloginschemapolicy state %s", authenticationloginschemapolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateAuthenticationloginschemapolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationloginschemapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationloginschemapolicyName := d.Get("name").(string)

	authenticationloginschemapolicy := authentication.Authenticationloginschemapolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for authenticationloginschemapolicy %s, starting update", authenticationloginschemapolicyName)
		authenticationloginschemapolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for authenticationloginschemapolicy %s, starting update", authenticationloginschemapolicyName)
		authenticationloginschemapolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for authenticationloginschemapolicy %s, starting update", authenticationloginschemapolicyName)
		authenticationloginschemapolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for authenticationloginschemapolicy %s, starting update", authenticationloginschemapolicyName)
		authenticationloginschemapolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationloginschemapolicy %s, starting update", authenticationloginschemapolicyName)
		authenticationloginschemapolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for authenticationloginschemapolicy %s, starting update", authenticationloginschemapolicyName)
		authenticationloginschemapolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationloginschemapolicy.Type(), authenticationloginschemapolicyName, &authenticationloginschemapolicy)
		if err != nil {
			return fmt.Errorf("Error updating authenticationloginschemapolicy %s", authenticationloginschemapolicyName)
		}
	}
	return readAuthenticationloginschemapolicyFunc(d, meta)
}

func deleteAuthenticationloginschemapolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationloginschemapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationloginschemapolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationloginschemapolicy.Type(), authenticationloginschemapolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
