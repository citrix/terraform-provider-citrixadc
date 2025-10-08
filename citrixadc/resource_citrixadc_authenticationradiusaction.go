package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationradiusaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationradiusactionFunc,
		ReadContext:   readAuthenticationradiusactionFunc,
		UpdateContext: updateAuthenticationradiusactionFunc,
		DeleteContext: deleteAuthenticationradiusactionFunc,
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
			"radkey": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"accounting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authservretry": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"callingstationid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipattributetype": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipvendorid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"passencoding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pwdattributetype": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pwdvendorid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radattributetype": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radgroupseparator": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radgroupsprefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radnasid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radnasip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radvendorid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"serverip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tunnelendpointclientip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationradiusactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationradiusactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiusactionName := d.Get("name").(string)
	authenticationradiusaction := authentication.Authenticationradiusaction{
		Accounting:                 d.Get("accounting").(string),
		Authentication:             d.Get("authentication").(string),
		Authservretry:              d.Get("authservretry").(int),
		Authtimeout:                d.Get("authtimeout").(int),
		Callingstationid:           d.Get("callingstationid").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Ipattributetype:            d.Get("ipattributetype").(int),
		Ipvendorid:                 d.Get("ipvendorid").(int),
		Name:                       d.Get("name").(string),
		Passencoding:               d.Get("passencoding").(string),
		Pwdattributetype:           d.Get("pwdattributetype").(int),
		Pwdvendorid:                d.Get("pwdvendorid").(int),
		Radattributetype:           d.Get("radattributetype").(int),
		Radgroupseparator:          d.Get("radgroupseparator").(string),
		Radgroupsprefix:            d.Get("radgroupsprefix").(string),
		Radkey:                     d.Get("radkey").(string),
		Radnasid:                   d.Get("radnasid").(string),
		Radnasip:                   d.Get("radnasip").(string),
		Radvendorid:                d.Get("radvendorid").(int),
		Serverip:                   d.Get("serverip").(string),
		Servername:                 d.Get("servername").(string),
		Serverport:                 d.Get("serverport").(int),
		Tunnelendpointclientip:     d.Get("tunnelendpointclientip").(string),
	}

	_, err := client.AddResource(service.Authenticationradiusaction.Type(), authenticationradiusactionName, &authenticationradiusaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationradiusactionName)

	return readAuthenticationradiusactionFunc(ctx, d, meta)
}

func readAuthenticationradiusactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationradiusactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiusactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationradiusaction state %s", authenticationradiusactionName)
	data, err := client.FindResource(service.Authenticationradiusaction.Type(), authenticationradiusactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationradiusaction state %s", authenticationradiusactionName)
		d.SetId("")
		return nil
	}
	d.Set("accounting", data["accounting"])
	d.Set("authentication", data["authentication"])
	setToInt("authservretry", d, data["authservretry"])
	setToInt("authtimeout", d, data["authtimeout"])
	d.Set("callingstationid", data["callingstationid"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	setToInt("ipattributetype", d, data["ipattributetype"])
	setToInt("ipvendorid", d, data["ipvendorid"])
	d.Set("name", data["name"])
	d.Set("passencoding", data["passencoding"])
	setToInt("pwdattributetype", d, data["pwdattributetype"])
	setToInt("pwdvendorid", d, data["pwdvendorid"])
	setToInt("radattributetype", d, data["radattributetype"])
	d.Set("radgroupseparator", data["radgroupseparator"])
	d.Set("radgroupsprefix", data["radgroupsprefix"])
	// d.Set("radkey", data["radkey"]) Everytime it gives different encrypted key
	d.Set("radnasid", data["radnasid"])
	d.Set("radnasip", data["radnasip"])
	setToInt("radvendorid", d, data["radvendorid"])
	d.Set("serverip", data["serverip"])
	d.Set("servername", data["servername"])
	setToInt("serverport", d, data["serverport"])
	d.Set("tunnelendpointclientip", data["tunnelendpointclientip"])

	return nil

}

func updateAuthenticationradiusactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationradiusactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiusactionName := d.Get("name").(string)

	authenticationradiusaction := authentication.Authenticationradiusaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("accounting") {
		log.Printf("[DEBUG]  citrixadc-provider: Accounting has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Accounting = d.Get("accounting").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authservretry") {
		log.Printf("[DEBUG]  citrixadc-provider: Authservretry has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Authservretry = d.Get("authservretry").(int)
		hasChange = true
	}
	if d.HasChange("authtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Authtimeout has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Authtimeout = d.Get("authtimeout").(int)
		hasChange = true
	}
	if d.HasChange("callingstationid") {
		log.Printf("[DEBUG]  citrixadc-provider: Callingstationid has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Callingstationid = d.Get("callingstationid").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("ipattributetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipattributetype has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Ipattributetype = d.Get("ipattributetype").(int)
		hasChange = true
	}
	if d.HasChange("ipvendorid") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipvendorid has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Ipvendorid = d.Get("ipvendorid").(int)
		hasChange = true
	}
	if d.HasChange("passencoding") {
		log.Printf("[DEBUG]  citrixadc-provider: Passencoding has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Passencoding = d.Get("passencoding").(string)
		hasChange = true
	}
	if d.HasChange("pwdattributetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Pwdattributetype has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Pwdattributetype = d.Get("pwdattributetype").(int)
		hasChange = true
	}
	if d.HasChange("pwdvendorid") {
		log.Printf("[DEBUG]  citrixadc-provider: Pwdvendorid has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Pwdvendorid = d.Get("pwdvendorid").(int)
		hasChange = true
	}
	if d.HasChange("radattributetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Radattributetype has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Radattributetype = d.Get("radattributetype").(int)
		hasChange = true
	}
	if d.HasChange("radgroupseparator") {
		log.Printf("[DEBUG]  citrixadc-provider: Radgroupseparator has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Radgroupseparator = d.Get("radgroupseparator").(string)
		hasChange = true
	}
	if d.HasChange("radgroupsprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Radgroupsprefix has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Radgroupsprefix = d.Get("radgroupsprefix").(string)
		hasChange = true
	}
	if d.HasChange("radkey") {
		log.Printf("[DEBUG]  citrixadc-provider: Radkey has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Radkey = d.Get("radkey").(string)
		hasChange = true
	}
	if d.HasChange("radnasid") {
		log.Printf("[DEBUG]  citrixadc-provider: Radnasid has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Radnasid = d.Get("radnasid").(string)
		hasChange = true
	}
	if d.HasChange("radnasip") {
		log.Printf("[DEBUG]  citrixadc-provider: Radnasip has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Radnasip = d.Get("radnasip").(string)
		hasChange = true
	}
	if d.HasChange("radvendorid") {
		log.Printf("[DEBUG]  citrixadc-provider: Radvendorid has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Radvendorid = d.Get("radvendorid").(int)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  citrixadc-provider: Servername has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Serverport = d.Get("serverport").(int)
		hasChange = true
	}
	if d.HasChange("tunnelendpointclientip") {
		log.Printf("[DEBUG]  citrixadc-provider: Tunnelendpointclientip has changed for authenticationradiusaction %s, starting update", authenticationradiusactionName)
		authenticationradiusaction.Tunnelendpointclientip = d.Get("tunnelendpointclientip").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationradiusaction.Type(), authenticationradiusactionName, &authenticationradiusaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationradiusaction %s", authenticationradiusactionName)
		}
	}
	return readAuthenticationradiusactionFunc(ctx, d, meta)
}

func deleteAuthenticationradiusactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationradiusactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiusactionName := d.Id()
	err := client.DeleteResource(service.Authenticationradiusaction.Type(), authenticationradiusactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
