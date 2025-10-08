package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLsnappsprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnappsprofileFunc,
		ReadContext:   readLsnappsprofileFunc,
		UpdateContext: updateLsnappsprofileFunc,
		DeleteContext: deleteLsnappsprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"appsprofilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transportprotocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"filtering": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ippooling": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2info": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mapping": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			},
		},
	}
}

func createLsnappsprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnappsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsprofileName := d.Get("appsprofilename").(string)
	lsnappsprofile := lsn.Lsnappsprofile{
		Appsprofilename:   d.Get("appsprofilename").(string),
		Filtering:         d.Get("filtering").(string),
		Ippooling:         d.Get("ippooling").(string),
		L2info:            d.Get("l2info").(string),
		Mapping:           d.Get("mapping").(string),
		Tcpproxy:          d.Get("tcpproxy").(string),
		Td:                d.Get("td").(int),
		Transportprotocol: d.Get("transportprotocol").(string),
	}

	_, err := client.AddResource("lsnappsprofile", lsnappsprofileName, &lsnappsprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lsnappsprofileName)

	return readLsnappsprofileFunc(ctx, d, meta)
}

func readLsnappsprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnappsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnappsprofile state %s", lsnappsprofileName)
	data, err := client.FindResource("lsnappsprofile", lsnappsprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnappsprofile state %s", lsnappsprofileName)
		d.SetId("")
		return nil
	}
	d.Set("appsprofilename", data["appsprofilename"])
	d.Set("filtering", data["filtering"])
	d.Set("ippooling", data["ippooling"])
	d.Set("l2info", data["l2info"])
	d.Set("mapping", data["mapping"])
	d.Set("tcpproxy", data["tcpproxy"])
	setToInt("td", d, data["td"])
	d.Set("transportprotocol", data["transportprotocol"])

	return nil

}

func updateLsnappsprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsnappsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsprofileName := d.Get("appsprofilename").(string)

	lsnappsprofile := lsn.Lsnappsprofile{
		Appsprofilename: d.Get("appsprofilename").(string),
	}
	hasChange := false
	if d.HasChange("filtering") {
		log.Printf("[DEBUG]  citrixadc-provider: Filtering has changed for lsnappsprofile %s, starting update", lsnappsprofileName)
		lsnappsprofile.Filtering = d.Get("filtering").(string)
		hasChange = true
	}
	if d.HasChange("ippooling") {
		log.Printf("[DEBUG]  citrixadc-provider: Ippooling has changed for lsnappsprofile %s, starting update", lsnappsprofileName)
		lsnappsprofile.Ippooling = d.Get("ippooling").(string)
		hasChange = true
	}
	if d.HasChange("l2info") {
		log.Printf("[DEBUG]  citrixadc-provider: L2info has changed for lsnappsprofile %s, starting update", lsnappsprofileName)
		lsnappsprofile.L2info = d.Get("l2info").(string)
		hasChange = true
	}
	if d.HasChange("mapping") {
		log.Printf("[DEBUG]  citrixadc-provider: Mapping has changed for lsnappsprofile %s, starting update", lsnappsprofileName)
		lsnappsprofile.Mapping = d.Get("mapping").(string)
		hasChange = true
	}
	if d.HasChange("tcpproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpproxy has changed for lsnappsprofile %s, starting update", lsnappsprofileName)
		lsnappsprofile.Tcpproxy = d.Get("tcpproxy").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for lsnappsprofile %s, starting update", lsnappsprofileName)
		lsnappsprofile.Td = d.Get("td").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lsnappsprofile", &lsnappsprofile)
		if err != nil {
			return diag.Errorf("Error updating lsnappsprofile %s", lsnappsprofileName)
		}
	}
	return readLsnappsprofileFunc(ctx, d, meta)
}

func deleteLsnappsprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnappsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnappsprofileName := d.Id()
	err := client.DeleteResource("lsnappsprofile", lsnappsprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
