package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcInat() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createInatFunc,
		ReadContext:   readInatFunc,
		UpdateContext: updateInatFunc,
		DeleteContext: deleteInatFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ftp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"privateip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxyip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publicip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tcpproxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tftp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"useproxyport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usnip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createInatFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In createInatFunc")
	client := meta.(*NetScalerNitroClient).client
	var inatName string
	if v, ok := d.GetOk("name"); ok {
		inatName = v.(string)
	} else {
		inatName = resource.PrefixedUniqueId("tf-inat-")
		d.Set("name", inatName)
	}
	inat := network.Inat{
		Ftp:          d.Get("ftp").(string),
		Mode:         d.Get("mode").(string),
		Name:         d.Get("name").(string),
		Privateip:    d.Get("privateip").(string),
		Proxyip:      d.Get("proxyip").(string),
		Publicip:     d.Get("publicip").(string),
		Tcpproxy:     d.Get("tcpproxy").(string),
		Tftp:         d.Get("tftp").(string),
		Useproxyport: d.Get("useproxyport").(string),
		Usip:         d.Get("usip").(string),
		Usnip:        d.Get("usnip").(string),
	}

	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		inat.Td = intPtr(d.Get("td").(int))
	}

	_, err := client.AddResource(service.Inat.Type(), inatName, &inat)
	if err != nil {
		fmt.Printf("[DEBUG] netscaler-provider add inat failed, name=%s", inatName)
		return diag.FromErr(err)
	}

	d.SetId(inatName)

	return readInatFunc(ctx, d, meta)
}

func readInatFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In readInatFunc")
	client := meta.(*NetScalerNitroClient).client
	inatName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading inat state %s", inatName)
	data, err := client.FindResource(service.Inat.Type(), inatName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing inat state %s", inatName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("ftp", data["ftp"])
	d.Set("mode", data["mode"])
	d.Set("name", data["name"])
	d.Set("privateip", data["privateip"])
	d.Set("proxyip", data["proxyip"])
	d.Set("publicip", data["publicip"])
	d.Set("tcpproxy", data["tcpproxy"])
	setToInt("td", d, data["td"])
	d.Set("tftp", data["tftp"])
	d.Set("useproxyport", data["useproxyport"])
	d.Set("usip", data["usip"])
	d.Set("usnip", data["usnip"])

	return nil

}

func updateInatFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In updateInatFunc")
	client := meta.(*NetScalerNitroClient).client
	inatName := d.Get("name").(string)

	inat := network.Inat{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("ftp") {
		log.Printf("[DEBUG]  netscaler-provider: Ftp has changed for inat %s, starting update", inatName)
		inat.Ftp = d.Get("ftp").(string)
		hasChange = true
	}
	if d.HasChange("mode") {
		log.Printf("[DEBUG]  netscaler-provider: Mode has changed for inat %s, starting update", inatName)
		inat.Mode = d.Get("mode").(string)
		hasChange = true
	}
	if d.HasChange("privateip") {
		log.Printf("[DEBUG]  netscaler-provider: Privateip has changed for inat %s, starting update", inatName)
		inat.Privateip = d.Get("privateip").(string)
		hasChange = true
	}
	if d.HasChange("proxyip") {
		log.Printf("[DEBUG]  netscaler-provider: Proxyip has changed for inat %s, starting update", inatName)
		inat.Proxyip = d.Get("proxyip").(string)
		hasChange = true
	}
	if d.HasChange("publicip") {
		log.Printf("[DEBUG]  netscaler-provider: Publicip has changed for inat %s, starting update", inatName)
		inat.Publicip = d.Get("publicip").(string)
		hasChange = true
	}
	if d.HasChange("tcpproxy") {
		log.Printf("[DEBUG]  netscaler-provider: Tcpproxy has changed for inat %s, starting update", inatName)
		inat.Tcpproxy = d.Get("tcpproxy").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  netscaler-provider: Td has changed for inat %s, starting update", inatName)
		inat.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("tftp") {
		log.Printf("[DEBUG]  netscaler-provider: Tftp has changed for inat %s, starting update", inatName)
		inat.Tftp = d.Get("tftp").(string)
		hasChange = true
	}
	if d.HasChange("useproxyport") {
		log.Printf("[DEBUG]  netscaler-provider: Useproxyport has changed for inat %s, starting update", inatName)
		inat.Useproxyport = d.Get("useproxyport").(string)
		hasChange = true
	}
	if d.HasChange("usip") {
		log.Printf("[DEBUG]  netscaler-provider: Usip has changed for inat %s, starting update", inatName)
		inat.Usip = d.Get("usip").(string)
		hasChange = true
	}
	if d.HasChange("usnip") {
		log.Printf("[DEBUG]  netscaler-provider: Usnip has changed for inat %s, starting update", inatName)
		inat.Usnip = d.Get("usnip").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Inat.Type(), inatName, &inat)
		if err != nil {
			return diag.Errorf("Error updating inat %s", inatName)
		}
	}
	return readInatFunc(ctx, d, meta)
}

func deleteInatFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In deleteInatFunc")
	client := meta.(*NetScalerNitroClient).client
	inatName := d.Id()
	err := client.DeleteResource(service.Inat.Type(), inatName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
