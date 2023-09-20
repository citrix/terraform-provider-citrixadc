package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLsnsipalgprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsnsipalgprofileFunc,
		Read:          readLsnsipalgprofileFunc,
		Update:        updateLsnsipalgprofileFunc,
		Delete:        deleteLsnsipalgprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"sipalgprofilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"datasessionidletimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"opencontactpinhole": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"openrecordroutepinhole": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"openregisterpinhole": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"openroutepinhole": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"openviapinhole": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"registrationtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sipdstportrange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sipsessiontimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sipsrcportrange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"siptransportprotocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsnsipalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnsipalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnsipalgprofileName := d.Get("sipalgprofilename").(string)
	lsnsipalgprofile := lsn.Lsnsipalgprofile{
		Datasessionidletimeout: d.Get("datasessionidletimeout").(int),
		Opencontactpinhole:     d.Get("opencontactpinhole").(string),
		Openrecordroutepinhole: d.Get("openrecordroutepinhole").(string),
		Openregisterpinhole:    d.Get("openregisterpinhole").(string),
		Openroutepinhole:       d.Get("openroutepinhole").(string),
		Openviapinhole:         d.Get("openviapinhole").(string),
		Registrationtimeout:    d.Get("registrationtimeout").(int),
		Rport:                  d.Get("rport").(string),
		Sipalgprofilename:      d.Get("sipalgprofilename").(string),
		Sipdstportrange:        d.Get("sipdstportrange").(string),
		Sipsessiontimeout:      d.Get("sipsessiontimeout").(int),
		Sipsrcportrange:        d.Get("sipsrcportrange").(string),
		Siptransportprotocol:   d.Get("siptransportprotocol").(string),
	}

	_, err := client.AddResource("lsnsipalgprofile", lsnsipalgprofileName, &lsnsipalgprofile)
	if err != nil {
		return err
	}

	d.SetId(lsnsipalgprofileName)

	err = readLsnsipalgprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsnsipalgprofile but we can't read it ?? %s", lsnsipalgprofileName)
		return nil
	}
	return nil
}

func readLsnsipalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnsipalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnsipalgprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnsipalgprofile state %s", lsnsipalgprofileName)
	data, err := client.FindResource("lsnsipalgprofile", lsnsipalgprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnsipalgprofile state %s", lsnsipalgprofileName)
		d.SetId("")
		return nil
	}
	d.Set("datasessionidletimeout", data["datasessionidletimeout"])
	d.Set("opencontactpinhole", data["opencontactpinhole"])
	d.Set("openrecordroutepinhole", data["openrecordroutepinhole"])
	d.Set("openregisterpinhole", data["openregisterpinhole"])
	d.Set("openroutepinhole", data["openroutepinhole"])
	d.Set("openviapinhole", data["openviapinhole"])
	d.Set("registrationtimeout", data["registrationtimeout"])
	d.Set("rport", data["rport"])
	d.Set("sipalgprofilename", data["sipalgprofilename"])
	d.Set("sipdstportrange", data["sipdstportrange"])
	d.Set("sipsessiontimeout", data["sipsessiontimeout"])
	d.Set("sipsrcportrange", data["sipsrcportrange"])
	d.Set("siptransportprotocol", data["siptransportprotocol"])

	return nil

}

func updateLsnsipalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsnsipalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnsipalgprofileName := d.Get("sipalgprofilename").(string)

	lsnsipalgprofile := lsn.Lsnsipalgprofile{
		Sipalgprofilename: d.Get("sipalgprofilename").(string),
	}
	hasChange := false
	if d.HasChange("datasessionidletimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Datasessionidletimeout has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Datasessionidletimeout = d.Get("datasessionidletimeout").(int)
		hasChange = true
	}
	if d.HasChange("opencontactpinhole") {
		log.Printf("[DEBUG]  citrixadc-provider: Opencontactpinhole has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Opencontactpinhole = d.Get("opencontactpinhole").(string)
		hasChange = true
	}
	if d.HasChange("openrecordroutepinhole") {
		log.Printf("[DEBUG]  citrixadc-provider: Openrecordroutepinhole has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Openrecordroutepinhole = d.Get("openrecordroutepinhole").(string)
		hasChange = true
	}
	if d.HasChange("openregisterpinhole") {
		log.Printf("[DEBUG]  citrixadc-provider: Openregisterpinhole has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Openregisterpinhole = d.Get("openregisterpinhole").(string)
		hasChange = true
	}
	if d.HasChange("openroutepinhole") {
		log.Printf("[DEBUG]  citrixadc-provider: Openroutepinhole has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Openroutepinhole = d.Get("openroutepinhole").(string)
		hasChange = true
	}
	if d.HasChange("openviapinhole") {
		log.Printf("[DEBUG]  citrixadc-provider: Openviapinhole has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Openviapinhole = d.Get("openviapinhole").(string)
		hasChange = true
	}
	if d.HasChange("registrationtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Registrationtimeout has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Registrationtimeout = d.Get("registrationtimeout").(int)
		hasChange = true
	}
	if d.HasChange("rport") {
		log.Printf("[DEBUG]  citrixadc-provider: Rport has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Rport = d.Get("rport").(string)
		hasChange = true
	}
	if d.HasChange("sipdstportrange") {
		log.Printf("[DEBUG]  citrixadc-provider: Sipdstportrange has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Sipdstportrange = d.Get("sipdstportrange").(string)
		hasChange = true
	}
	if d.HasChange("sipsessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sipsessiontimeout has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Sipsessiontimeout = d.Get("sipsessiontimeout").(int)
		hasChange = true
	}
	if d.HasChange("sipsrcportrange") {
		log.Printf("[DEBUG]  citrixadc-provider: Sipsrcportrange has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Sipsrcportrange = d.Get("sipsrcportrange").(string)
		hasChange = true
	}
	if d.HasChange("siptransportprotocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Siptransportprotocol has changed for lsnsipalgprofile %s, starting update", lsnsipalgprofileName)
		lsnsipalgprofile.Siptransportprotocol = d.Get("siptransportprotocol").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lsnsipalgprofile", &lsnsipalgprofile)
		if err != nil {
			return fmt.Errorf("Error updating lsnsipalgprofile %s", lsnsipalgprofileName)
		}
	}
	return readLsnsipalgprofileFunc(d, meta)
}

func deleteLsnsipalgprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnsipalgprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnsipalgprofileName := d.Id()
	err := client.DeleteResource("lsnsipalgprofile", lsnsipalgprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
