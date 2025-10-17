package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVpntrafficaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpntrafficactionFunc,
		ReadContext:   readVpntrafficactionFunc,
		UpdateContext: updateVpntrafficactionFunc,
		DeleteContext: deleteVpntrafficactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"qual": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"apptimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"formssoaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fta": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hdx": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdaccount": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passwdexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlssoprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"userexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wanscaler": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVpntrafficactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpntrafficactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpntrafficactionName := d.Get("name").(string)
	vpntrafficaction := vpn.Vpntrafficaction{
		Formssoaction:    d.Get("formssoaction").(string),
		Fta:              d.Get("fta").(string),
		Hdx:              d.Get("hdx").(string),
		Kcdaccount:       d.Get("kcdaccount").(string),
		Name:             d.Get("name").(string),
		Passwdexpression: d.Get("passwdexpression").(string),
		Proxy:            d.Get("proxy").(string),
		Qual:             d.Get("qual").(string),
		Samlssoprofile:   d.Get("samlssoprofile").(string),
		Sso:              d.Get("sso").(string),
		Userexpression:   d.Get("userexpression").(string),
		Wanscaler:        d.Get("wanscaler").(string),
	}

	if raw := d.GetRawConfig().GetAttr("apptimeout"); !raw.IsNull() {
		vpntrafficaction.Apptimeout = intPtr(d.Get("apptimeout").(int))
	}

	_, err := client.AddResource(service.Vpntrafficaction.Type(), vpntrafficactionName, &vpntrafficaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vpntrafficactionName)

	return readVpntrafficactionFunc(ctx, d, meta)
}

func readVpntrafficactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpntrafficactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpntrafficactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpntrafficaction state %s", vpntrafficactionName)
	data, err := client.FindResource(service.Vpntrafficaction.Type(), vpntrafficactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpntrafficaction state %s", vpntrafficactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	setToInt("apptimeout", d, data["apptimeout"])
	d.Set("formssoaction", data["formssoaction"])
	d.Set("fta", data["fta"])
	d.Set("hdx", data["hdx"])
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("name", data["name"])
	d.Set("passwdexpression", data["passwdexpression"])
	d.Set("proxy", data["proxy"])
	d.Set("qual", data["qual"])
	d.Set("samlssoprofile", data["samlssoprofile"])
	d.Set("sso", data["sso"])
	d.Set("userexpression", data["userexpression"])
	d.Set("wanscaler", data["wanscaler"])

	return nil

}

func updateVpntrafficactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpntrafficactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpntrafficactionName := d.Get("name").(string)

	vpntrafficaction := vpn.Vpntrafficaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("apptimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Apptimeout has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Apptimeout = intPtr(d.Get("apptimeout").(int))
		hasChange = true
	}
	if d.HasChange("formssoaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Formssoaction has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Formssoaction = d.Get("formssoaction").(string)
		hasChange = true
	}
	if d.HasChange("fta") {
		log.Printf("[DEBUG]  citrixadc-provider: Fta has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Fta = d.Get("fta").(string)
		hasChange = true
	}
	if d.HasChange("hdx") {
		log.Printf("[DEBUG]  citrixadc-provider: Hdx has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Hdx = d.Get("hdx").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG]  citrixadc-provider: Kcdaccount has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("passwdexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Passwdexpression has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Passwdexpression = d.Get("passwdexpression").(string)
		hasChange = true
	}
	if d.HasChange("proxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxy has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Proxy = d.Get("proxy").(string)
		hasChange = true
	}
	if d.HasChange("qual") {
		log.Printf("[DEBUG]  citrixadc-provider: Qual has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Qual = d.Get("qual").(string)
		hasChange = true
	}
	if d.HasChange("samlssoprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlssoprofile has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Samlssoprofile = d.Get("samlssoprofile").(string)
		hasChange = true
	}
	if d.HasChange("sso") {
		log.Printf("[DEBUG]  citrixadc-provider: Sso has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Sso = d.Get("sso").(string)
		hasChange = true
	}
	if d.HasChange("userexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Userexpression has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Userexpression = d.Get("userexpression").(string)
		hasChange = true
	}
	if d.HasChange("wanscaler") {
		log.Printf("[DEBUG]  citrixadc-provider: Wanscaler has changed for vpntrafficaction %s, starting update", vpntrafficactionName)
		vpntrafficaction.Wanscaler = d.Get("wanscaler").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Vpntrafficaction.Type(), vpntrafficactionName, &vpntrafficaction)
		if err != nil {
			return diag.Errorf("Error updating vpntrafficaction %s", vpntrafficactionName)
		}
	}
	return readVpntrafficactionFunc(ctx, d, meta)
}

func deleteVpntrafficactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpntrafficactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpntrafficactionName := d.Id()
	err := client.DeleteResource(service.Vpntrafficaction.Type(), vpntrafficactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
