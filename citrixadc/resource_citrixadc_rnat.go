package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcRnat() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createRnatFunc,
		ReadContext:   readRnatFunc,
		UpdateContext: updateRnatFunc,
		DeleteContext: deleteRnatFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"aclname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"connfailover": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"natip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"srcippersistency": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"useproxyport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRnatFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createRnatFunc")
	client := meta.(*NetScalerNitroClient).client
	rnatName := d.Get("name").(string)

	rnat := network.Rnat{
		Aclname:          d.Get("aclname").(string),
		Connfailover:     d.Get("connfailover").(string),
		Name:             d.Get("name").(string),
		Netmask:          d.Get("netmask").(string),
		Network:          d.Get("network").(string),
		Ownergroup:       d.Get("ownergroup").(string),
		Srcippersistency: d.Get("srcippersistency").(string),
		Useproxyport:     d.Get("useproxyport").(string),
	}

	if raw := d.GetRawConfig().GetAttr("redirectport"); !raw.IsNull() {
		rnat.Redirectport = intPtr(d.Get("redirectport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		rnat.Td = intPtr(d.Get("td").(int))
	}

	_, err := client.AddResource(service.Rnat.Type(), rnatName, &rnat)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(rnatName)

	return readRnatFunc(ctx, d, meta)
}

func readRnatFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readRnatFunc")
	client := meta.(*NetScalerNitroClient).client
	rnatName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rnat state %s", rnatName)
	data, err := client.FindResource(service.Rnat.Type(), rnatName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rnat state %s", rnatName)
		d.SetId("")
		return nil
	}
	d.Set("aclname", data["aclname"])
	d.Set("connfailover", data["connfailover"])
	d.Set("name", data["name"])
	d.Set("natip", data["natip"])
	d.Set("netmask", data["netmask"])
	d.Set("network", data["network"])
	d.Set("ownergroup", data["ownergroup"])
	setToInt("redirectport", d, data["redirectport"])
	d.Set("srcippersistency", data["srcippersistency"])
	setToInt("td", d, data["td"])
	d.Set("useproxyport", data["useproxyport"])

	return nil

}

func updateRnatFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRnatFunc")
	client := meta.(*NetScalerNitroClient).client
	rnatName := d.Get("name").(string)

	rnat := network.Rnat{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("connfailover") {
		log.Printf("[DEBUG]  citrixadc-provider: Connfailover has changed for rnat %s, starting update", rnatName)
		rnat.Connfailover = d.Get("connfailover").(string)
		hasChange = true
	}
	if d.HasChange("natip") {
		log.Printf("[DEBUG]  citrixadc-provider: Natip has changed for rnat %s, starting update", rnatName)
		rnat.Natip = d.Get("natip").(string)
		hasChange = true
	}
	if d.HasChange("ownergroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownergroup has changed for rnat %s, starting update", rnatName)
		rnat.Ownergroup = d.Get("ownergroup").(string)
		hasChange = true
	}
	if d.HasChange("redirectport") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectport has changed for rnat %s, starting update", rnatName)
		rnat.Redirectport = intPtr(d.Get("redirectport").(int))
		hasChange = true
	}
	if d.HasChange("srcippersistency") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcippersistency has changed for rnat %s, starting update", rnatName)
		rnat.Srcippersistency = d.Get("srcippersistency").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for rnat %s, starting update", rnatName)
		rnat.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("useproxyport") {
		log.Printf("[DEBUG]  citrixadc-provider: Useproxyport has changed for rnat %s, starting update", rnatName)
		rnat.Useproxyport = d.Get("useproxyport").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Rnat.Type(), rnatName, &rnat)
		if err != nil {
			return diag.Errorf("Error updating rnat %s", rnatName)
		}
	}
	return readRnatFunc(ctx, d, meta)
}

func deleteRnatFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRnatFunc")
	client := meta.(*NetScalerNitroClient).client
	rnatName := d.Id()
	err := client.DeleteResource(service.Rnat.Type(), rnatName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
