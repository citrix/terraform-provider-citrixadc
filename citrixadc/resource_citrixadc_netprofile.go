package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNetprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNetprofileFunc,
		ReadContext:   readNetprofileFunc,
		UpdateContext: updateNetprofileFunc,
		DeleteContext: deleteNetprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"mbf": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"overridelsn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxyprotocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxyprotocoltxversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcip": {
				Type:     schema.TypeString,
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
			"proxyprotocolaftertlshandshake": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNetprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNetprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	var netprofileName string
	if v, ok := d.GetOk("name"); ok {
		netprofileName = v.(string)
	} else {
		netprofileName = resource.PrefixedUniqueId("tf-netprofile-")
		d.Set("name", netprofileName)
	}
	netprofile := network.Netprofile{
		Mbf:                            d.Get("mbf").(string),
		Name:                           d.Get("name").(string),
		Overridelsn:                    d.Get("overridelsn").(string),
		Proxyprotocol:                  d.Get("proxyprotocol").(string),
		Proxyprotocoltxversion:         d.Get("proxyprotocoltxversion").(string),
		Srcip:                          d.Get("srcip").(string),
		Srcippersistency:               d.Get("srcippersistency").(string),
		Proxyprotocolaftertlshandshake: d.Get("proxyprotocolaftertlshandshake").(string),
	}

	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		netprofile.Td = intPtr(d.Get("td").(int))
	}

	_, err := client.AddResource(service.Netprofile.Type(), netprofileName, &netprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(netprofileName)

	return readNetprofileFunc(ctx, d, meta)
}

func readNetprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNetprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	netprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading netprofile state %s", netprofileName)
	data, err := client.FindResource(service.Netprofile.Type(), netprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing netprofile state %s", netprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("mbf", data["mbf"])
	d.Set("name", data["name"])
	d.Set("overridelsn", data["overridelsn"])
	d.Set("proxyprotocol", data["proxyprotocol"])
	d.Set("proxyprotocoltxversion", data["proxyprotocoltxversion"])
	d.Set("srcip", data["srcip"])
	d.Set("srcippersistency", data["srcippersistency"])
	setToInt("td", d, data["td"])
	d.Set("proxyprotocolaftertlshandshake", data["proxyprotocolaftertlshandshake"])

	return nil

}

func updateNetprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNetprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	netprofileName := d.Get("name").(string)

	netprofile := network.Netprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("mbf") {
		log.Printf("[DEBUG]  citrixadc-provider: Mbf has changed for netprofile %s, starting update", netprofileName)
		netprofile.Mbf = d.Get("mbf").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for netprofile %s, starting update", netprofileName)
		netprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("overridelsn") {
		log.Printf("[DEBUG]  citrixadc-provider: Overridelsn has changed for netprofile %s, starting update", netprofileName)
		netprofile.Overridelsn = d.Get("overridelsn").(string)
		hasChange = true
	}
	if d.HasChange("proxyprotocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyprotocol has changed for netprofile %s, starting update", netprofileName)
		netprofile.Proxyprotocol = d.Get("proxyprotocol").(string)
		hasChange = true
	}
	if d.HasChange("proxyprotocoltxversion") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyprotocoltxversion has changed for netprofile %s, starting update", netprofileName)
		netprofile.Proxyprotocoltxversion = d.Get("proxyprotocoltxversion").(string)
		hasChange = true
	}
	if d.HasChange("srcip") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcip has changed for netprofile %s, starting update", netprofileName)
		netprofile.Srcip = d.Get("srcip").(string)
		hasChange = true
	}
	if d.HasChange("srcippersistency") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcippersistency has changed for netprofile %s, starting update", netprofileName)
		netprofile.Srcippersistency = d.Get("srcippersistency").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for netprofile %s, starting update", netprofileName)
		netprofile.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("proxyprotocolaftertlshandshake") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyprotocolaftertlshandshake has changed for netprofile %s, starting update", netprofileName)
		netprofile.Proxyprotocolaftertlshandshake = d.Get("proxyprotocolaftertlshandshake").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Netprofile.Type(), netprofileName, &netprofile)
		if err != nil {
			return diag.Errorf("Error updating netprofile %s", netprofileName)
		}
	}
	return readNetprofileFunc(ctx, d, meta)
}

func deleteNetprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNetprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	netprofileName := d.Id()
	err := client.DeleteResource(service.Netprofile.Type(), netprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
