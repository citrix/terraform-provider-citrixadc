package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationtacacsaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationtacacsactionFunc,
		ReadContext:   readAuthenticationtacacsactionFunc,
		UpdateContext: updateAuthenticationtacacsactionFunc,
		DeleteContext: deleteAuthenticationtacacsactionFunc,
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
			"accounting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auditfailedcmds": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupattrname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tacacssecret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationtacacsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationtacacsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationtacacsactionName := d.Get("name").(string)
	authenticationtacacsaction := authentication.Authenticationtacacsaction{
		Accounting:                 d.Get("accounting").(string),
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
		Attributes:                 d.Get("attributes").(string),
		Auditfailedcmds:            d.Get("auditfailedcmds").(string),
		Authorization:              d.Get("authorization").(string),
		Authtimeout:                d.Get("authtimeout").(int),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Groupattrname:              d.Get("groupattrname").(string),
		Name:                       d.Get("name").(string),
		Serverip:                   d.Get("serverip").(string),
		Serverport:                 d.Get("serverport").(int),
		Tacacssecret:               d.Get("tacacssecret").(string),
	}

	_, err := client.AddResource(service.Authenticationtacacsaction.Type(), authenticationtacacsactionName, &authenticationtacacsaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationtacacsactionName)

	return readAuthenticationtacacsactionFunc(ctx, d, meta)
}

func readAuthenticationtacacsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationtacacsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationtacacsactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationtacacsaction state %s", authenticationtacacsactionName)
	data, err := client.FindResource(service.Authenticationtacacsaction.Type(), authenticationtacacsactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationtacacsaction state %s", authenticationtacacsactionName)
		d.SetId("")
		return nil
	}
	d.Set("accounting", data["accounting"])
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
	d.Set("attributes", data["attributes"])
	d.Set("auditfailedcmds", data["auditfailedcmds"])
	d.Set("authorization", data["authorization"])
	setToInt("authtimeout", d, data["authtimeout"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("groupattrname", data["groupattrname"])
	d.Set("name", data["name"])
	d.Set("serverip", data["serverip"])
	setToInt("serverport", d, data["serverport"])
	d.Set("tacacssecret", data["tacacssecret"])

	return nil

}

func updateAuthenticationtacacsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationtacacsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationtacacsactionName := d.Get("name").(string)

	authenticationtacacsaction := authentication.Authenticationtacacsaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("accounting") {
		log.Printf("[DEBUG]  citrixadc-provider: Accounting has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Accounting = d.Get("accounting").(string)
		hasChange = true
	}
	if d.HasChange("attribute1") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute1 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute1 = d.Get("attribute1").(string)
		hasChange = true
	}
	if d.HasChange("attribute10") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute10 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute10 = d.Get("attribute10").(string)
		hasChange = true
	}
	if d.HasChange("attribute11") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute11 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute11 = d.Get("attribute11").(string)
		hasChange = true
	}
	if d.HasChange("attribute12") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute12 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute12 = d.Get("attribute12").(string)
		hasChange = true
	}
	if d.HasChange("attribute13") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute13 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute13 = d.Get("attribute13").(string)
		hasChange = true
	}
	if d.HasChange("attribute14") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute14 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute14 = d.Get("attribute14").(string)
		hasChange = true
	}
	if d.HasChange("attribute15") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute15 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute15 = d.Get("attribute15").(string)
		hasChange = true
	}
	if d.HasChange("attribute16") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute16 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute16 = d.Get("attribute16").(string)
		hasChange = true
	}
	if d.HasChange("attribute2") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute2 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute2 = d.Get("attribute2").(string)
		hasChange = true
	}
	if d.HasChange("attribute3") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute3 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute3 = d.Get("attribute3").(string)
		hasChange = true
	}
	if d.HasChange("attribute4") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute4 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute4 = d.Get("attribute4").(string)
		hasChange = true
	}
	if d.HasChange("attribute5") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute5 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute5 = d.Get("attribute5").(string)
		hasChange = true
	}
	if d.HasChange("attribute6") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute6 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute6 = d.Get("attribute6").(string)
		hasChange = true
	}
	if d.HasChange("attribute7") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute7 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute7 = d.Get("attribute7").(string)
		hasChange = true
	}
	if d.HasChange("attribute8") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute8 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute8 = d.Get("attribute8").(string)
		hasChange = true
	}
	if d.HasChange("attribute9") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute9 has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attribute9 = d.Get("attribute9").(string)
		hasChange = true
	}
	if d.HasChange("attributes") {
		log.Printf("[DEBUG]  citrixadc-provider: Attributes has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Attributes = d.Get("attributes").(string)
		hasChange = true
	}
	if d.HasChange("auditfailedcmds") {
		log.Printf("[DEBUG]  citrixadc-provider: Auditfailedcmds has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Auditfailedcmds = d.Get("auditfailedcmds").(string)
		hasChange = true
	}
	if d.HasChange("authorization") {
		log.Printf("[DEBUG]  citrixadc-provider: Authorization has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Authorization = d.Get("authorization").(string)
		hasChange = true
	}
	if d.HasChange("authtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Authtimeout has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Authtimeout = d.Get("authtimeout").(int)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("groupattrname") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupattrname has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Groupattrname = d.Get("groupattrname").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Serverport = d.Get("serverport").(int)
		hasChange = true
	}
	if d.HasChange("tacacssecret") {
		log.Printf("[DEBUG]  citrixadc-provider: Tacacssecret has changed for authenticationtacacsaction %s, starting update", authenticationtacacsactionName)
		authenticationtacacsaction.Tacacssecret = d.Get("tacacssecret").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationtacacsaction.Type(), authenticationtacacsactionName, &authenticationtacacsaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationtacacsaction %s", authenticationtacacsactionName)
		}
	}
	return readAuthenticationtacacsactionFunc(ctx, d, meta)
}

func deleteAuthenticationtacacsactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationtacacsactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationtacacsactionName := d.Id()
	err := client.DeleteResource(service.Authenticationtacacsaction.Type(), authenticationtacacsactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
