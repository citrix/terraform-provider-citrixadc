package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationoauthidppolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationoauthidppolicyFunc,
		ReadContext:   readAuthenticationoauthidppolicyFunc,
		UpdateContext: updateAuthenticationoauthidppolicyFunc,
		DeleteContext: deleteAuthenticationoauthidppolicyFunc,
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

func createAuthenticationoauthidppolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationoauthidppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthidppolicyName := d.Get("name").(string)
	authenticationoauthidppolicy := authentication.Authenticationoauthidppolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource("authenticationoauthidppolicy", authenticationoauthidppolicyName, &authenticationoauthidppolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationoauthidppolicyName)

	return readAuthenticationoauthidppolicyFunc(ctx, d, meta)
}

func readAuthenticationoauthidppolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationoauthidppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthidppolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationoauthidppolicy state %s", authenticationoauthidppolicyName)
	data, err := client.FindResource("authenticationoauthidppolicy", authenticationoauthidppolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationoauthidppolicy state %s", authenticationoauthidppolicyName)
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

func updateAuthenticationoauthidppolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationoauthidppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthidppolicyName := d.Get("name").(string)

	authenticationoauthidppolicy := authentication.Authenticationoauthidppolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for authenticationoauthidppolicy %s, starting update", authenticationoauthidppolicyName)
		authenticationoauthidppolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for authenticationoauthidppolicy %s, starting update", authenticationoauthidppolicyName)
		authenticationoauthidppolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for authenticationoauthidppolicy %s, starting update", authenticationoauthidppolicyName)
		authenticationoauthidppolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationoauthidppolicy %s, starting update", authenticationoauthidppolicyName)
		authenticationoauthidppolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for authenticationoauthidppolicy %s, starting update", authenticationoauthidppolicyName)
		authenticationoauthidppolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationoauthidppolicy", authenticationoauthidppolicyName, &authenticationoauthidppolicy)
		if err != nil {
			return diag.Errorf("Error updating authenticationoauthidppolicy %s", authenticationoauthidppolicyName)
		}
	}
	return readAuthenticationoauthidppolicyFunc(ctx, d, meta)
}

func deleteAuthenticationoauthidppolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationoauthidppolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthidppolicyName := d.Id()
	err := client.DeleteResource("authenticationoauthidppolicy", authenticationoauthidppolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
