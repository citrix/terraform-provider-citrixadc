package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLsnrtspalgprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnrtspalgprofileFunc,
		ReadContext:   readLsnrtspalgprofileFunc,
		UpdateContext: updateLsnrtspalgprofileFunc,
		DeleteContext: deleteLsnrtspalgprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"rtspalgprofilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rtspportrange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtspidletimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rtsptransportprotocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsnrtspalgprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnrtspalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnrtspalgprofileName := d.Get("rtspalgprofilename").(string)
	lsnrtspalgprofile := lsn.Lsnrtspalgprofile{
		Rtspalgprofilename:    d.Get("rtspalgprofilename").(string),
		Rtspidletimeout:       d.Get("rtspidletimeout").(int),
		Rtspportrange:         d.Get("rtspportrange").(string),
		Rtsptransportprotocol: d.Get("rtsptransportprotocol").(string),
	}

	_, err := client.AddResource("lsnrtspalgprofile", lsnrtspalgprofileName, &lsnrtspalgprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lsnrtspalgprofileName)

	return readLsnrtspalgprofileFunc(ctx, d, meta)
}

func readLsnrtspalgprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnrtspalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnrtspalgprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnrtspalgprofile state %s", lsnrtspalgprofileName)
	data, err := client.FindResource("lsnrtspalgprofile", lsnrtspalgprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnrtspalgprofile state %s", lsnrtspalgprofileName)
		d.SetId("")
		return nil
	}
	d.Set("rtspalgprofilename", data["rtspalgprofilename"])
	setToInt("rtspidletimeout", d, data["rtspidletimeout"])
	d.Set("rtspportrange", data["rtspportrange"])
	d.Set("rtsptransportprotocol", data["rtsptransportprotocol"])

	return nil

}

func updateLsnrtspalgprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsnrtspalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnrtspalgprofileName := d.Get("rtspalgprofilename").(string)

	lsnrtspalgprofile := lsn.Lsnrtspalgprofile{
		Rtspalgprofilename: d.Get("rtspalgprofilename").(string),
	}
	hasChange := false
	if d.HasChange("rtspidletimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Rtspidletimeout has changed for lsnrtspalgprofile %s, starting update", lsnrtspalgprofileName)
		lsnrtspalgprofile.Rtspidletimeout = d.Get("rtspidletimeout").(int)
		hasChange = true
	}
	if d.HasChange("rtspportrange") {
		log.Printf("[DEBUG]  citrixadc-provider: Rtspportrange has changed for lsnrtspalgprofile %s, starting update", lsnrtspalgprofileName)
		lsnrtspalgprofile.Rtspportrange = d.Get("rtspportrange").(string)
		hasChange = true
	}
	if d.HasChange("rtsptransportprotocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Rtsptransportprotocol has changed for lsnrtspalgprofile %s, starting update", lsnrtspalgprofileName)
		lsnrtspalgprofile.Rtsptransportprotocol = d.Get("rtsptransportprotocol").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("lsnrtspalgprofile", lsnrtspalgprofileName, &lsnrtspalgprofile)
		if err != nil {
			return diag.Errorf("Error updating lsnrtspalgprofile %s", lsnrtspalgprofileName)
		}
	}
	return readLsnrtspalgprofileFunc(ctx, d, meta)
}

func deleteLsnrtspalgprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnrtspalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnrtspalgprofileName := d.Id()
	err := client.DeleteResource("lsnrtspalgprofile", lsnrtspalgprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
