package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSsldtlsprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSsldtlsprofileFunc,
		ReadContext:   readSsldtlsprofileFunc,
		UpdateContext: updateSsldtlsprofileFunc,
		DeleteContext: deleteSsldtlsprofileFunc,
		Schema: map[string]*schema.Schema{
			"initialretrytimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"helloverifyrequest": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxbadmacignorecount": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxholdqlen": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxpacketsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxrecordsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxretrytime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pmtudiscovery": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"terminatesession": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSsldtlsprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSsldtlsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	var ssldtlsprofileName string
	if v, ok := d.GetOk("name"); ok {
		ssldtlsprofileName = v.(string)
	} else {
		ssldtlsprofileName = resource.PrefixedUniqueId("tf-ssldtlsprofile-")
		d.Set("name", ssldtlsprofileName)
	}
	ssldtlsprofile := ssl.Ssldtlsprofile{
		Helloverifyrequest: d.Get("helloverifyrequest").(string),
		Name:               ssldtlsprofileName,
		Pmtudiscovery:      d.Get("pmtudiscovery").(string),
		Terminatesession:   d.Get("terminatesession").(string),
	}
	if raw := d.GetRawConfig().GetAttr("initialretrytimeout"); !raw.IsNull() {
		ssldtlsprofile.Initialretrytimeout = intPtr(d.Get("initialretrytimeout").(int))
	}

	if raw := d.GetRawConfig().GetAttr("maxbadmacignorecount"); !raw.IsNull() {
		ssldtlsprofile.Maxbadmacignorecount = intPtr(d.Get("maxbadmacignorecount").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxholdqlen"); !raw.IsNull() {
		ssldtlsprofile.Maxholdqlen = intPtr(d.Get("maxholdqlen").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxpacketsize"); !raw.IsNull() {
		ssldtlsprofile.Maxpacketsize = intPtr(d.Get("maxpacketsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxrecordsize"); !raw.IsNull() {
		ssldtlsprofile.Maxrecordsize = intPtr(d.Get("maxrecordsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxretrytime"); !raw.IsNull() {
		ssldtlsprofile.Maxretrytime = intPtr(d.Get("maxretrytime").(int))
	}

	_, err := client.AddResource(service.Ssldtlsprofile.Type(), ssldtlsprofileName, &ssldtlsprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ssldtlsprofileName)

	return readSsldtlsprofileFunc(ctx, d, meta)
}

func readSsldtlsprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSsldtlsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ssldtlsprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading ssldtlsprofile state %s", ssldtlsprofileName)
	data, err := client.FindResource(service.Ssldtlsprofile.Type(), ssldtlsprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ssldtlsprofile state %s", ssldtlsprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	setToInt("initialretrytimeout", d, data["initialretrytimeout"])
	d.Set("helloverifyrequest", data["helloverifyrequest"])
	setToInt("maxbadmacignorecount", d, data["maxbadmacignorecount"])
	setToInt("maxholdqlen", d, data["maxholdqlen"])
	setToInt("maxpacketsize", d, data["maxpacketsize"])
	setToInt("maxrecordsize", d, data["maxrecordsize"])
	setToInt("maxretrytime", d, data["maxretrytime"])
	d.Set("name", data["name"])
	d.Set("pmtudiscovery", data["pmtudiscovery"])
	d.Set("terminatesession", data["terminatesession"])

	return nil

}

func updateSsldtlsprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSsldtlsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ssldtlsprofileName := d.Get("name").(string)

	ssldtlsprofile := ssl.Ssldtlsprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("initialretrytimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Initialretrytimeout has changed for ssldtlsprofile, starting update")
		ssldtlsprofile.Initialretrytimeout = intPtr(d.Get("initialretrytimeout").(int))
		hasChange = true
	}
	if d.HasChange("helloverifyrequest") {
		log.Printf("[DEBUG]  citrixadc-provider: Helloverifyrequest has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Helloverifyrequest = d.Get("helloverifyrequest").(string)
		hasChange = true
	}
	if d.HasChange("maxbadmacignorecount") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxbadmacignorecount has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Maxbadmacignorecount = intPtr(d.Get("maxbadmacignorecount").(int))
		hasChange = true
	}
	if d.HasChange("maxholdqlen") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxholdqlen has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Maxholdqlen = intPtr(d.Get("maxholdqlen").(int))
		hasChange = true
	}
	if d.HasChange("maxpacketsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxpacketsize has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Maxpacketsize = intPtr(d.Get("maxpacketsize").(int))
		hasChange = true
	}
	if d.HasChange("maxrecordsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxrecordsize has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Maxrecordsize = intPtr(d.Get("maxrecordsize").(int))
		hasChange = true
	}
	if d.HasChange("maxretrytime") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxretrytime has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Maxretrytime = intPtr(d.Get("maxretrytime").(int))
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("pmtudiscovery") {
		log.Printf("[DEBUG]  citrixadc-provider: Pmtudiscovery has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Pmtudiscovery = d.Get("pmtudiscovery").(string)
		hasChange = true
	}
	if d.HasChange("terminatesession") {
		log.Printf("[DEBUG]  citrixadc-provider: Terminatesession has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Terminatesession = d.Get("terminatesession").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Ssldtlsprofile.Type(), ssldtlsprofileName, &ssldtlsprofile)
		if err != nil {
			return diag.Errorf("Error updating ssldtlsprofile %s", ssldtlsprofileName)
		}
	}
	return readSsldtlsprofileFunc(ctx, d, meta)
}

func deleteSsldtlsprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSsldtlsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ssldtlsprofileName := d.Id()
	err := client.DeleteResource(service.Ssldtlsprofile.Type(), ssldtlsprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
