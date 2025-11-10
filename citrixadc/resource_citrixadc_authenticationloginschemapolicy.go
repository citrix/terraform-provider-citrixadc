package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationloginschemapolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationloginschemapolicyFunc,
		ReadContext:   readAuthenticationloginschemapolicyFunc,
		UpdateContext: updateAuthenticationloginschemapolicyFunc,
		DeleteContext: deleteAuthenticationloginschemapolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationloginschemapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationloginschemapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationloginschemapolicyName := d.Get("name").(string)
	authenticationloginschemapolicy := authentication.Authenticationloginschemapolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource(service.Authenticationloginschemapolicy.Type(), authenticationloginschemapolicyName, &authenticationloginschemapolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationloginschemapolicyName)

	return readAuthenticationloginschemapolicyFunc(ctx, d, meta)
}

func readAuthenticationloginschemapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateAuthenticationloginschemapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("Error updating authenticationloginschemapolicy %s", authenticationloginschemapolicyName)
		}
	}
	return readAuthenticationloginschemapolicyFunc(ctx, d, meta)
}

func deleteAuthenticationloginschemapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationloginschemapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationloginschemapolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationloginschemapolicy.Type(), authenticationloginschemapolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
