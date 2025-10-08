package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLsntransportprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsntransportprofileFunc,
		ReadContext:   readLsntransportprofileFunc,
		UpdateContext: updateLsntransportprofileFunc,
		DeleteContext: deleteLsntransportprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"transportprofilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transportprotocol": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"finrsttimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"groupsessionlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"portpreserveparity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portpreserverange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portquota": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sessionquota": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sessiontimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"stuntimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"syncheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"synidletimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsntransportprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsntransportprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsntransportprofileName := d.Get("transportprofilename").(string)
	lsntransportprofile := lsn.Lsntransportprofile{
		Finrsttimeout:        d.Get("finrsttimeout").(int),
		Groupsessionlimit:    d.Get("groupsessionlimit").(int),
		Portpreserveparity:   d.Get("portpreserveparity").(string),
		Portpreserverange:    d.Get("portpreserverange").(string),
		Portquota:            d.Get("portquota").(int),
		Sessionquota:         d.Get("sessionquota").(int),
		Sessiontimeout:       d.Get("sessiontimeout").(int),
		Stuntimeout:          d.Get("stuntimeout").(int),
		Syncheck:             d.Get("syncheck").(string),
		Synidletimeout:       d.Get("synidletimeout").(int),
		Transportprofilename: d.Get("transportprofilename").(string),
		Transportprotocol:    d.Get("transportprotocol").(string),
	}

	_, err := client.AddResource("lsntransportprofile", lsntransportprofileName, &lsntransportprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lsntransportprofileName)

	return readLsntransportprofileFunc(ctx, d, meta)
}

func readLsntransportprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsntransportprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsntransportprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsntransportprofile state %s", lsntransportprofileName)
	data, err := client.FindResource("lsntransportprofile", lsntransportprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsntransportprofile state %s", lsntransportprofileName)
		d.SetId("")
		return nil
	}
	d.Set("transportprofilename", data["transportprofilename"])
	setToInt("finrsttimeout", d, data["finrsttimeout"])
	setToInt("groupsessionlimit", d, data["groupsessionlimit"])
	d.Set("portpreserveparity", data["portpreserveparity"])
	d.Set("portpreserverange", data["portpreserverange"])
	setToInt("portquota", d, data["portquota"])
	setToInt("sessionquota", d, data["sessionquota"])
	setToInt("sessiontimeout", d, data["sessiontimeout"])
	setToInt("stuntimeout", d, data["stuntimeout"])
	d.Set("syncheck", data["syncheck"])
	setToInt("synidletimeout", d, data["synidletimeout"])
	d.Set("transportprofilename", data["transportprofilename"])
	d.Set("transportprotocol", data["transportprotocol"])

	return nil

}

func updateLsntransportprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsntransportprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsntransportprofileName := d.Get("transportprofilename").(string)

	lsntransportprofile := lsn.Lsntransportprofile{
		Transportprofilename: d.Get("transportprofilename").(string),
	}
	hasChange := false
	if d.HasChange("finrsttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Finrsttimeout has changed for lsntransportprofile %s, starting update", lsntransportprofileName)
		lsntransportprofile.Finrsttimeout = d.Get("finrsttimeout").(int)
		hasChange = true
	}
	if d.HasChange("groupsessionlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupsessionlimit has changed for lsntransportprofile %s, starting update", lsntransportprofileName)
		lsntransportprofile.Groupsessionlimit = d.Get("groupsessionlimit").(int)
		hasChange = true
	}
	if d.HasChange("portpreserveparity") {
		log.Printf("[DEBUG]  citrixadc-provider: Portpreserveparity has changed for lsntransportprofile %s, starting update", lsntransportprofileName)
		lsntransportprofile.Portpreserveparity = d.Get("portpreserveparity").(string)
		hasChange = true
	}
	if d.HasChange("portpreserverange") {
		log.Printf("[DEBUG]  citrixadc-provider: Portpreserverange has changed for lsntransportprofile %s, starting update", lsntransportprofileName)
		lsntransportprofile.Portpreserverange = d.Get("portpreserverange").(string)
		hasChange = true
	}
	if d.HasChange("portquota") {
		log.Printf("[DEBUG]  citrixadc-provider: Portquota has changed for lsntransportprofile %s, starting update", lsntransportprofileName)
		lsntransportprofile.Portquota = d.Get("portquota").(int)
		hasChange = true
	}
	if d.HasChange("sessionquota") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionquota has changed for lsntransportprofile %s, starting update", lsntransportprofileName)
		lsntransportprofile.Sessionquota = d.Get("sessionquota").(int)
		hasChange = true
	}
	if d.HasChange("sessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessiontimeout has changed for lsntransportprofile %s, starting update", lsntransportprofileName)
		lsntransportprofile.Sessiontimeout = d.Get("sessiontimeout").(int)
		hasChange = true
	}
	if d.HasChange("stuntimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Stuntimeout has changed for lsntransportprofile %s, starting update", lsntransportprofileName)
		lsntransportprofile.Stuntimeout = d.Get("stuntimeout").(int)
		hasChange = true
	}
	if d.HasChange("syncheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Syncheck has changed for lsntransportprofile %s, starting update", lsntransportprofileName)
		lsntransportprofile.Syncheck = d.Get("syncheck").(string)
		hasChange = true
	}
	if d.HasChange("synidletimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Synidletimeout has changed for lsntransportprofile %s, starting update", lsntransportprofileName)
		lsntransportprofile.Synidletimeout = d.Get("synidletimeout").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lsntransportprofile", &lsntransportprofile)
		if err != nil {
			return diag.Errorf("Error updating lsntransportprofile %s", lsntransportprofileName)
		}
	}
	return readLsntransportprofileFunc(ctx, d, meta)
}

func deleteLsntransportprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsntransportprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsntransportprofileName := d.Id()
	err := client.DeleteResource("lsntransportprofile", lsntransportprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
