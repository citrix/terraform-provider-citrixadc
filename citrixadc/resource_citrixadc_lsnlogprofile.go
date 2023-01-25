package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLsnlogprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnlogprofileFunc,
		Read:          readLsnlogprofileFunc,
		Update:        updateLsnlogprofileFunc,
		Delete:        deleteLsnlogprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"logprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"analyticsprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logcompact": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logipfix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logsessdeletion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logsubscrinfo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsnlogprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId(lsnlogprofileName)

	err = readLsnlogprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnlogprofile but we can't read it ?? %s", lsnlogprofileName)
		return nil
	}
	return nil
}

func readLsnlogprofileFunc(d *schema.ResourceData, meta interface{}) error {
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

func updateLsnlogprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
			return fmt.Errorf("Error updating lsnlogprofile %s", lsnlogprofileName)
		}
	}
	return readLsnlogprofileFunc(d, meta)
}

func deleteLsnlogprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnlogprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnlogprofileName := d.Id()
	err := client.DeleteResource("lsnlogprofile", lsnlogprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
