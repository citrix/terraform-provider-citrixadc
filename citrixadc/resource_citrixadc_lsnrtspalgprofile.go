package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLsnrtspalgprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnrtspalgprofileFunc,
		Read:          readLsnrtspalgprofileFunc,
		Update:        updateLsnrtspalgprofileFunc,
		Delete:        deleteLsnrtspalgprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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

func createLsnrtspalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId(lsnrtspalgprofileName)

	err = readLsnrtspalgprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnrtspalgprofile but we can't read it ?? %s", lsnrtspalgprofileName)
		return nil
	}
	return nil
}

func readLsnrtspalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("rtspidletimeout", data["rtspidletimeout"])
	d.Set("rtspportrange", data["rtspportrange"])
	d.Set("rtsptransportprotocol", data["rtsptransportprotocol"])

	return nil

}

func updateLsnrtspalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
			return fmt.Errorf("Error updating lsnrtspalgprofile %s", lsnrtspalgprofileName)
		}
	}
	return readLsnrtspalgprofileFunc(d, meta)
}

func deleteLsnrtspalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnrtspalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnrtspalgprofileName := d.Id()
	err := client.DeleteResource("lsnrtspalgprofile", lsnrtspalgprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
