package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/pcp"

	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcPcpprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPcpprofileFunc,
		ReadContext:   readPcpprofileFunc,
		UpdateContext: updatePcpprofileFunc,
		DeleteContext: deletePcpprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"announcemulticount": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mapping": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxmaplife": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minmaplife": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"peer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"thirdparty": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createPcpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPcpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpprofileName := d.Get("name").(string)

	pcpprofile := make(map[string]interface{})
	if v, ok := d.GetOkExists("thirdparty"); ok {
		pcpprofile["thirdparty"] = v.(string)
	}
	if v, ok := d.GetOk("peer"); ok {
		pcpprofile["peer"] = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		pcpprofile["name"] = v.(string)
	}
	if v, ok := d.GetOk("minmaplife"); ok {
		pcpprofile["minmaplife"] = v.(int)
	}
	if v, ok := d.GetOkExists("maxmaplife"); ok {
		pcpprofile["maxmaplife"] = v.(int)
	}
	if v, ok := d.GetOk("mapping"); ok {
		pcpprofile["mapping"] = v.(string)
	}
	if v, ok := d.GetOkExists("announcemulticount"); ok {
		val, _ := strconv.Atoi(v.(string))
		pcpprofile["announcemulticount"] = val
	}

	_, err := client.AddResource("pcpprofile", pcpprofileName, &pcpprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(pcpprofileName)

	return readPcpprofileFunc(ctx, d, meta)
}

func readPcpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPcpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading pcpprofile state %s", pcpprofileName)
	data, err := client.FindResource("pcpprofile", pcpprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing pcpprofile state %s", pcpprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("announcemulticount", data["announcemulticount"])
	d.Set("mapping", data["mapping"])
	setToInt("maxmaplife", d, data["maxmaplife"])
	setToInt("minmaplife", d, data["minmaplife"])
	d.Set("peer", data["peer"])
	d.Set("thirdparty", data["thirdparty"])

	return nil

}

func updatePcpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePcpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpprofileName := d.Get("name").(string)

	pcpprofile := pcp.Pcpprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("announcemulticount") {
		log.Printf("[DEBUG]  citrixadc-provider: Announcemulticount has changed for pcpprofile %s, starting update", pcpprofileName)
		val, _ := strconv.Atoi(d.Get("announcemulticount").(string))
		pcpprofile.Announcemulticount = val
		hasChange = true
	}
	if d.HasChange("mapping") {
		log.Printf("[DEBUG]  citrixadc-provider: Mapping has changed for pcpprofile %s, starting update", pcpprofileName)
		pcpprofile.Mapping = d.Get("mapping").(string)
		hasChange = true
	}
	if d.HasChange("maxmaplife") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxmaplife has changed for pcpprofile %s, starting update", pcpprofileName)
		pcpprofile.Maxmaplife = d.Get("maxmaplife").(int)
		hasChange = true
	}
	if d.HasChange("minmaplife") {
		log.Printf("[DEBUG]  citrixadc-provider: Minmaplife has changed for pcpprofile %s, starting update", pcpprofileName)
		pcpprofile.Minmaplife = d.Get("minmaplife").(int)
		hasChange = true
	}
	if d.HasChange("peer") {
		log.Printf("[DEBUG]  citrixadc-provider: Peer has changed for pcpprofile %s, starting update", pcpprofileName)
		pcpprofile.Peer = d.Get("peer").(string)
		hasChange = true
	}
	if d.HasChange("thirdparty") {
		log.Printf("[DEBUG]  citrixadc-provider: Thirdparty has changed for pcpprofile %s, starting update", pcpprofileName)
		pcpprofile.Thirdparty = d.Get("thirdparty").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("pcpprofile", pcpprofileName, &pcpprofile)
		if err != nil {
			return diag.Errorf("Error updating pcpprofile %s", pcpprofileName)
		}
	}
	return readPcpprofileFunc(ctx, d, meta)
}

func deletePcpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePcpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	pcpprofileName := d.Id()
	err := client.DeleteResource("pcpprofile", pcpprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
