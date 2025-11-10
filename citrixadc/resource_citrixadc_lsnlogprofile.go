package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLsnlogprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnlogprofileFunc,
		ReadContext:   readLsnlogprofileFunc,
		UpdateContext: updateLsnlogprofileFunc,
		DeleteContext: deleteLsnlogprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"logprofilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"analyticsprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logcompact": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logipfix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logsessdeletion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logsubscrinfo": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsnlogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnlogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnlogprofileName := d.Get("logprofilename").(string)
	lsnlogprofile := lsn.Lsnlogprofile{
		Analyticsprofile: d.Get("analyticsprofile").(string),
		Logcompact:       d.Get("logcompact").(string),
		Logipfix:         d.Get("logipfix").(string),
		Logprofilename:   d.Get("logprofilename").(string),
		Logsessdeletion:  d.Get("logsessdeletion").(string),
		Logsubscrinfo:    d.Get("logsubscrinfo").(string),
	}

	_, err := client.AddResource("lsnlogprofile", lsnlogprofileName, &lsnlogprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lsnlogprofileName)

	return readLsnlogprofileFunc(ctx, d, meta)
}

func readLsnlogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnlogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnlogprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnlogprofile state %s", lsnlogprofileName)
	data, err := client.FindResource("lsnlogprofile", lsnlogprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnlogprofile state %s", lsnlogprofileName)
		d.SetId("")
		return nil
	}
	d.Set("logprofilename", data["logprofilename"])
	d.Set("analyticsprofile", data["analyticsprofile"])
	d.Set("logcompact", data["logcompact"])
	d.Set("logipfix", data["logipfix"])
	d.Set("logsessdeletion", data["logsessdeletion"])
	d.Set("logsubscrinfo", data["logsubscrinfo"])

	return nil

}

func updateLsnlogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsnlogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnlogprofileName := d.Get("logprofilename").(string)

	lsnlogprofile := lsn.Lsnlogprofile{
		Logprofilename: d.Get("logprofilename").(string),
	}
	hasChange := false
	if d.HasChange("analyticsprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Analyticsprofile has changed for lsnlogprofile %s, starting update", lsnlogprofileName)
		lsnlogprofile.Analyticsprofile = d.Get("analyticsprofile").(string)
		hasChange = true
	}
	if d.HasChange("logcompact") {
		log.Printf("[DEBUG]  citrixadc-provider: Logcompact has changed for lsnlogprofile %s, starting update", lsnlogprofileName)
		lsnlogprofile.Logcompact = d.Get("logcompact").(string)
		hasChange = true
	}
	if d.HasChange("logipfix") {
		log.Printf("[DEBUG]  citrixadc-provider: Logipfix has changed for lsnlogprofile %s, starting update", lsnlogprofileName)
		lsnlogprofile.Logipfix = d.Get("logipfix").(string)
		hasChange = true
	}
	if d.HasChange("logsessdeletion") {
		log.Printf("[DEBUG]  citrixadc-provider: Logsessdeletion has changed for lsnlogprofile %s, starting update", lsnlogprofileName)
		lsnlogprofile.Logsessdeletion = d.Get("logsessdeletion").(string)
		hasChange = true
	}
	if d.HasChange("logsubscrinfo") {
		log.Printf("[DEBUG]  citrixadc-provider: Logsubscrinfo has changed for lsnlogprofile %s, starting update", lsnlogprofileName)
		lsnlogprofile.Logsubscrinfo = d.Get("logsubscrinfo").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("lsnlogprofile", lsnlogprofileName, &lsnlogprofile)
		if err != nil {
			return diag.Errorf("Error updating lsnlogprofile %s", lsnlogprofileName)
		}
	}
	return readLsnlogprofileFunc(ctx, d, meta)
}

func deleteLsnlogprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnlogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnlogprofileName := d.Id()
	err := client.DeleteResource("lsnlogprofile", lsnlogprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
