package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVpnurl() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnurlFunc,
		ReadContext:   readVpnurlFunc,
		UpdateContext: updateVpnurlFunc,
		DeleteContext: deleteVpnurlFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"urlname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"actualurl": {
				Type:     schema.TypeString,
				Required: true,
			},
			"linkname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"appjson": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"applicationtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientlessaccess": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iconurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlssoprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssotype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vservername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVpnurlFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnurlFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlName := d.Get("urlname").(string)
	vpnurl := vpn.Vpnurl{
		Actualurl:        d.Get("actualurl").(string),
		Appjson:          d.Get("appjson").(string),
		Applicationtype:  d.Get("applicationtype").(string),
		Clientlessaccess: d.Get("clientlessaccess").(string),
		Comment:          d.Get("comment").(string),
		Iconurl:          d.Get("iconurl").(string),
		Linkname:         d.Get("linkname").(string),
		Samlssoprofile:   d.Get("samlssoprofile").(string),
		Ssotype:          d.Get("ssotype").(string),
		Urlname:          d.Get("urlname").(string),
		Vservername:      d.Get("vservername").(string),
	}

	_, err := client.AddResource(service.Vpnurl.Type(), vpnurlName, &vpnurl)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vpnurlName)

	return readVpnurlFunc(ctx, d, meta)
}

func readVpnurlFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnurlFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnurl state %s", vpnurlName)
	data, err := client.FindResource(service.Vpnurl.Type(), vpnurlName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnurl state %s", vpnurlName)
		d.SetId("")
		return nil
	}
	d.Set("urlname", data["urlname"])
	d.Set("actualurl", data["actualurl"])
	d.Set("appjson", data["appjson"])
	d.Set("applicationtype", data["applicationtype"])
	d.Set("clientlessaccess", data["clientlessaccess"])
	d.Set("comment", data["comment"])
	d.Set("iconurl", data["iconurl"])
	d.Set("linkname", data["linkname"])
	d.Set("samlssoprofile", data["samlssoprofile"])
	d.Set("ssotype", data["ssotype"])
	d.Set("urlname", data["urlname"])
	d.Set("vservername", data["vservername"])

	return nil

}

func updateVpnurlFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnurlFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlName := d.Get("urlname").(string)

	vpnurl := vpn.Vpnurl{
		Urlname: d.Get("urlname").(string),
	}
	hasChange := false
	if d.HasChange("actualurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Actualurl has changed for vpnurl %s, starting update", vpnurlName)
		vpnurl.Actualurl = d.Get("actualurl").(string)
		hasChange = true
	}
	if d.HasChange("appjson") {
		log.Printf("[DEBUG]  citrixadc-provider: Appjson has changed for vpnurl %s, starting update", vpnurlName)
		vpnurl.Appjson = d.Get("appjson").(string)
		hasChange = true
	}
	if d.HasChange("applicationtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Applicationtype has changed for vpnurl %s, starting update", vpnurlName)
		vpnurl.Applicationtype = d.Get("applicationtype").(string)
		hasChange = true
	}
	if d.HasChange("clientlessaccess") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientlessaccess has changed for vpnurl %s, starting update", vpnurlName)
		vpnurl.Clientlessaccess = d.Get("clientlessaccess").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for vpnurl %s, starting update", vpnurlName)
		vpnurl.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("iconurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Iconurl has changed for vpnurl %s, starting update", vpnurlName)
		vpnurl.Iconurl = d.Get("iconurl").(string)
		hasChange = true
	}
	if d.HasChange("linkname") {
		log.Printf("[DEBUG]  citrixadc-provider: Linkname has changed for vpnurl %s, starting update", vpnurlName)
		vpnurl.Linkname = d.Get("linkname").(string)
		hasChange = true
	}
	if d.HasChange("samlssoprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlssoprofile has changed for vpnurl %s, starting update", vpnurlName)
		vpnurl.Samlssoprofile = d.Get("samlssoprofile").(string)
		hasChange = true
	}
	if d.HasChange("ssotype") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssotype has changed for vpnurl %s, starting update", vpnurlName)
		vpnurl.Ssotype = d.Get("ssotype").(string)
		hasChange = true
	}
	if d.HasChange("vservername") {
		log.Printf("[DEBUG]  citrixadc-provider: Vservername has changed for vpnurl %s, starting update", vpnurlName)
		vpnurl.Vservername = d.Get("vservername").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Vpnurl.Type(), vpnurlName, &vpnurl)
		if err != nil {
			return diag.Errorf("Error updating vpnurl %s", vpnurlName)
		}
	}
	return readVpnurlFunc(ctx, d, meta)
}

func deleteVpnurlFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnurlFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlName := d.Id()
	err := client.DeleteResource(service.Vpnurl.Type(), vpnurlName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
