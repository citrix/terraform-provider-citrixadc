package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLsntransportprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsntransportprofileFunc,
		Read:          readLsntransportprofileFunc,
		Update:        updateLsntransportprofileFunc,
		Delete:        deleteLsntransportprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"transportprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transportprotocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"finrsttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"groupsessionlimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"portpreserveparity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portpreserverange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portquota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sessionquota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sessiontimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"stuntimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"syncheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"synidletimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsntransportprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId(lsntransportprofileName)

	err = readLsntransportprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsntransportprofile but we can't read it ?? %s", lsntransportprofileName)
		return nil
	}
	return nil
}

func readLsntransportprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("finrsttimeout", data["finrsttimeout"])
	d.Set("groupsessionlimit", data["groupsessionlimit"])
	d.Set("portpreserveparity", data["portpreserveparity"])
	d.Set("portpreserverange", data["portpreserverange"])
	d.Set("portquota", data["portquota"])
	d.Set("sessionquota", data["sessionquota"])
	d.Set("sessiontimeout", data["sessiontimeout"])
	d.Set("stuntimeout", data["stuntimeout"])
	d.Set("syncheck", data["syncheck"])
	d.Set("synidletimeout", data["synidletimeout"])
	d.Set("transportprofilename", data["transportprofilename"])
	d.Set("transportprotocol", data["transportprotocol"])

	return nil

}

func updateLsntransportprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
			return fmt.Errorf("Error updating lsntransportprofile %s", lsntransportprofileName)
		}
	}
	return readLsntransportprofileFunc(d, meta)
}

func deleteLsntransportprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsntransportprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsntransportprofileName := d.Id()
	err := client.DeleteResource("lsntransportprofile", lsntransportprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
