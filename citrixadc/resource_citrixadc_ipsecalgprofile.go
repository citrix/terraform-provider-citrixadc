package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ipsecalg"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcIpsecalgprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIpsecalgprofileFunc,
		ReadContext:   readIpsecalgprofileFunc,
		UpdateContext: updateIpsecalgprofileFunc,
		DeleteContext: deleteIpsecalgprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"connfailover": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"espgatetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"espsessiontimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ikesessiontimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createIpsecalgprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIpsecalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecalgprofileName := d.Get("name").(string)
	ipsecalgprofile := ipsecalg.Ipsecalgprofile{
		Connfailover:      d.Get("connfailover").(string),
		Espgatetimeout:    d.Get("espgatetimeout").(int),
		Espsessiontimeout: d.Get("espsessiontimeout").(int),
		Ikesessiontimeout: d.Get("ikesessiontimeout").(int),
		Name:              d.Get("name").(string),
	}

	_, err := client.AddResource("ipsecalgprofile", ipsecalgprofileName, &ipsecalgprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ipsecalgprofileName)

	return readIpsecalgprofileFunc(ctx, d, meta)
}

func readIpsecalgprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIpsecalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecalgprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading ipsecalgprofile state %s", ipsecalgprofileName)
	data, err := client.FindResource("ipsecalgprofile", ipsecalgprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ipsecalgprofile state %s", ipsecalgprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("connfailover", data["connfailover"])
	setToInt("espgatetimeout", d, data["espgatetimeout"])
	setToInt("espsessiontimeout", d, data["espsessiontimeout"])
	setToInt("ikesessiontimeout", d, data["ikesessiontimeout"])

	return nil

}

func updateIpsecalgprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIpsecalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecalgprofileName := d.Get("name").(string)

	ipsecalgprofile := ipsecalg.Ipsecalgprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("connfailover") {
		log.Printf("[DEBUG]  citrixadc-provider: Connfailover has changed for ipsecalgprofile %s, starting update", ipsecalgprofileName)
		ipsecalgprofile.Connfailover = d.Get("connfailover").(string)
		hasChange = true
	}
	if d.HasChange("espgatetimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Espgatetimeout has changed for ipsecalgprofile %s, starting update", ipsecalgprofileName)
		ipsecalgprofile.Espgatetimeout = d.Get("espgatetimeout").(int)
		hasChange = true
	}
	if d.HasChange("espsessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Espsessiontimeout has changed for ipsecalgprofile %s, starting update", ipsecalgprofileName)
		ipsecalgprofile.Espsessiontimeout = d.Get("espsessiontimeout").(int)
		hasChange = true
	}
	if d.HasChange("ikesessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Ikesessiontimeout has changed for ipsecalgprofile %s, starting update", ipsecalgprofileName)
		ipsecalgprofile.Ikesessiontimeout = d.Get("ikesessiontimeout").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("ipsecalgprofile", &ipsecalgprofile)
		if err != nil {
			return diag.Errorf("Error updating ipsecalgprofile %s", ipsecalgprofileName)
		}
	}
	return readIpsecalgprofileFunc(ctx, d, meta)
}

func deleteIpsecalgprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIpsecalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsecalgprofileName := d.Id()
	err := client.DeleteResource("ipsecalgprofile", ipsecalgprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
