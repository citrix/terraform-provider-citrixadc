package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationnegotiateaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationnegotiateactionFunc,
		ReadContext:   readAuthenticationnegotiateactionFunc,
		UpdateContext: updateAuthenticationnegotiateactionFunc,
		DeleteContext: deleteAuthenticationnegotiateactionFunc,
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
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domainuser": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domainuserpasswd": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keytab": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntlmpath": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ou": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationnegotiateactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationnegotiateactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnegotiateactionName := d.Get("name").(string)
	authenticationnegotiateaction := authentication.Authenticationnegotiateaction{
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Domain:                     d.Get("domain").(string),
		Domainuser:                 d.Get("domainuser").(string),
		Domainuserpasswd:           d.Get("domainuserpasswd").(string),
		Keytab:                     d.Get("keytab").(string),
		Name:                       d.Get("name").(string),
		Ntlmpath:                   d.Get("ntlmpath").(string),
		Ou:                         d.Get("ou").(string),
	}

	_, err := client.AddResource(service.Authenticationnegotiateaction.Type(), authenticationnegotiateactionName, &authenticationnegotiateaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationnegotiateactionName)

	return readAuthenticationnegotiateactionFunc(ctx, d, meta)
}

func readAuthenticationnegotiateactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationnegotiateactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnegotiateactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationnegotiateaction state %s", authenticationnegotiateactionName)
	data, err := client.FindResource(service.Authenticationnegotiateaction.Type(), authenticationnegotiateactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationnegotiateaction state %s", authenticationnegotiateactionName)
		d.SetId("")
		return nil
	}
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("domain", data["domain"])
	d.Set("domainuser", data["domainuser"])
	//d.Set("domainuserpasswd", data["domainuserpasswd"]) different value is recieved from ADC
	d.Set("keytab", data["keytab"])
	d.Set("name", data["name"])
	d.Set("ntlmpath", data["ntlmpath"])
	d.Set("ou", data["ou"])

	return nil

}

func updateAuthenticationnegotiateactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationnegotiateactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnegotiateactionName := d.Get("name").(string)

	authenticationnegotiateaction := authentication.Authenticationnegotiateaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationnegotiateaction %s, starting update", authenticationnegotiateactionName)
		authenticationnegotiateaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG]  citrixadc-provider: Domain has changed for authenticationnegotiateaction %s, starting update", authenticationnegotiateactionName)
		authenticationnegotiateaction.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("domainuser") {
		log.Printf("[DEBUG]  citrixadc-provider: Domainuser has changed for authenticationnegotiateaction %s, starting update", authenticationnegotiateactionName)
		authenticationnegotiateaction.Domainuser = d.Get("domainuser").(string)
		hasChange = true
	}
	if d.HasChange("domainuserpasswd") {
		log.Printf("[DEBUG]  citrixadc-provider: Domainuserpasswd has changed for authenticationnegotiateaction %s, starting update", authenticationnegotiateactionName)
		authenticationnegotiateaction.Domainuserpasswd = d.Get("domainuserpasswd").(string)
		hasChange = true
	}
	if d.HasChange("keytab") {
		log.Printf("[DEBUG]  citrixadc-provider: Keytab has changed for authenticationnegotiateaction %s, starting update", authenticationnegotiateactionName)
		authenticationnegotiateaction.Keytab = d.Get("keytab").(string)
		hasChange = true
	}
	if d.HasChange("ntlmpath") {
		log.Printf("[DEBUG]  citrixadc-provider: Ntlmpath has changed for authenticationnegotiateaction %s, starting update", authenticationnegotiateactionName)
		authenticationnegotiateaction.Ntlmpath = d.Get("ntlmpath").(string)
		hasChange = true
	}
	if d.HasChange("ou") {
		log.Printf("[DEBUG]  citrixadc-provider: Ou has changed for authenticationnegotiateaction %s, starting update", authenticationnegotiateactionName)
		authenticationnegotiateaction.Ou = d.Get("ou").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationnegotiateaction.Type(), authenticationnegotiateactionName, &authenticationnegotiateaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationnegotiateaction %s", authenticationnegotiateactionName)
		}
	}
	return readAuthenticationnegotiateactionFunc(ctx, d, meta)
}

func deleteAuthenticationnegotiateactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationnegotiateactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnegotiateactionName := d.Id()
	err := client.DeleteResource(service.Authenticationnegotiateaction.Type(), authenticationnegotiateactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
