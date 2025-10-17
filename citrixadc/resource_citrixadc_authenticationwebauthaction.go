package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationwebauthaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationwebauthactionFunc,
		ReadContext:   readAuthenticationwebauthactionFunc,
		UpdateContext: updateAuthenticationwebauthactionFunc,
		DeleteContext: deleteAuthenticationwebauthactionFunc,
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
			"scheme": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"serverip": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"serverport": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
			"successrule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"attribute1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute10": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute11": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute12": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute13": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute14": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute15": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute16": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute3": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute4": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute5": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute7": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute8": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute9": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fullreqexpr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationwebauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationwebauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthactionName := d.Get("name").(string)
	authenticationwebauthaction := authentication.Authenticationwebauthaction{
		Attribute1:                 d.Get("attribute1").(string),
		Attribute10:                d.Get("attribute10").(string),
		Attribute11:                d.Get("attribute11").(string),
		Attribute12:                d.Get("attribute12").(string),
		Attribute13:                d.Get("attribute13").(string),
		Attribute14:                d.Get("attribute14").(string),
		Attribute15:                d.Get("attribute15").(string),
		Attribute16:                d.Get("attribute16").(string),
		Attribute2:                 d.Get("attribute2").(string),
		Attribute3:                 d.Get("attribute3").(string),
		Attribute4:                 d.Get("attribute4").(string),
		Attribute5:                 d.Get("attribute5").(string),
		Attribute6:                 d.Get("attribute6").(string),
		Attribute7:                 d.Get("attribute7").(string),
		Attribute8:                 d.Get("attribute8").(string),
		Attribute9:                 d.Get("attribute9").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Fullreqexpr:                d.Get("fullreqexpr").(string),
		Name:                       d.Get("name").(string),
		Scheme:                     d.Get("scheme").(string),
		Serverip:                   d.Get("serverip").(string),
		Successrule:                d.Get("successrule").(string),
	}

	if raw := d.GetRawConfig().GetAttr("serverport"); !raw.IsNull() {
		authenticationwebauthaction.Serverport = intPtr(d.Get("serverport").(int))
	}

	_, err := client.AddResource(service.Authenticationwebauthaction.Type(), authenticationwebauthactionName, &authenticationwebauthaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationwebauthactionName)

	return readAuthenticationwebauthactionFunc(ctx, d, meta)
}

func readAuthenticationwebauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationwebauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationwebauthaction state %s", authenticationwebauthactionName)
	data, err := client.FindResource(service.Authenticationwebauthaction.Type(), authenticationwebauthactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationwebauthaction state %s", authenticationwebauthactionName)
		d.SetId("")
		return nil
	}
	d.Set("attribute1", data["attribute1"])
	d.Set("attribute10", data["attribute10"])
	d.Set("attribute11", data["attribute11"])
	d.Set("attribute12", data["attribute12"])
	d.Set("attribute13", data["attribute13"])
	d.Set("attribute14", data["attribute14"])
	d.Set("attribute15", data["attribute15"])
	d.Set("attribute16", data["attribute16"])
	d.Set("attribute2", data["attribute2"])
	d.Set("attribute3", data["attribute3"])
	d.Set("attribute4", data["attribute4"])
	d.Set("attribute5", data["attribute5"])
	d.Set("attribute6", data["attribute6"])
	d.Set("attribute7", data["attribute7"])
	d.Set("attribute8", data["attribute8"])
	d.Set("attribute9", data["attribute9"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("fullreqexpr", data["fullreqexpr"])
	d.Set("name", data["name"])
	d.Set("scheme", data["scheme"])
	d.Set("serverip", data["serverip"])
	setToInt("serverport", d, data["serverport"])
	d.Set("successrule", data["successrule"])

	return nil

}

func updateAuthenticationwebauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationwebauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthactionName := d.Get("name").(string)
	authenticationwebauthaction := authentication.Authenticationwebauthaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("attribute1") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute1 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute1 = d.Get("attribute1").(string)
		hasChange = true
	}
	if d.HasChange("attribute10") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute10 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute10 = d.Get("attribute10").(string)
		hasChange = true
	}
	if d.HasChange("attribute11") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute11 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute11 = d.Get("attribute11").(string)
		hasChange = true
	}
	if d.HasChange("attribute12") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute12 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute12 = d.Get("attribute12").(string)
		hasChange = true
	}
	if d.HasChange("attribute13") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute13 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute13 = d.Get("attribute13").(string)
		hasChange = true
	}
	if d.HasChange("attribute14") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute14 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute14 = d.Get("attribute14").(string)
		hasChange = true
	}
	if d.HasChange("attribute15") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute15 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute15 = d.Get("attribute15").(string)
		hasChange = true
	}
	if d.HasChange("attribute16") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute16 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute16 = d.Get("attribute16").(string)
		hasChange = true
	}
	if d.HasChange("attribute2") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute2 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute2 = d.Get("attribute2").(string)
		hasChange = true
	}
	if d.HasChange("attribute3") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute3 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute3 = d.Get("attribute3").(string)
		hasChange = true
	}
	if d.HasChange("attribute4") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute4 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute4 = d.Get("attribute4").(string)
		hasChange = true
	}
	if d.HasChange("attribute5") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute5 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute5 = d.Get("attribute5").(string)
		hasChange = true
	}
	if d.HasChange("attribute6") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute6 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute6 = d.Get("attribute6").(string)
		hasChange = true
	}
	if d.HasChange("attribute7") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute7 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute7 = d.Get("attribute7").(string)
		hasChange = true
	}
	if d.HasChange("attribute8") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute8 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute8 = d.Get("attribute8").(string)
		hasChange = true
	}
	if d.HasChange("attribute9") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute9 has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Attribute9 = d.Get("attribute9").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("fullreqexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Fullreqexpr has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Fullreqexpr = d.Get("fullreqexpr").(string)
		hasChange = true
	}
	if d.HasChange("scheme") {
		log.Printf("[DEBUG]  citrixadc-provider: Scheme has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Scheme = d.Get("scheme").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Serverport = intPtr(d.Get("serverport").(int))
		hasChange = true
	}
	if d.HasChange("successrule") {
		log.Printf("[DEBUG]  citrixadc-provider: Successrule has changed for authenticationwebauthaction %s, starting update", authenticationwebauthactionName)
		authenticationwebauthaction.Successrule = d.Get("successrule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationwebauthaction.Type(), authenticationwebauthactionName, &authenticationwebauthaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationwebauthaction %s", authenticationwebauthactionName)
		}
	}
	return readAuthenticationwebauthactionFunc(ctx, d, meta)
}

func deleteAuthenticationwebauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationwebauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthactionName := d.Id()
	err := client.DeleteResource(service.Authenticationwebauthaction.Type(), authenticationwebauthactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
