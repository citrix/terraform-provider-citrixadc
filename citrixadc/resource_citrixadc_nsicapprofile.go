package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsicapprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsicapprofileFunc,
		Read:          readNsicapprofileFunc,
		Update:        updateNsicapprofileFunc,
		Delete:        deleteNsicapprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"uri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"allow204": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connectionkeepalive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inserthttprequest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inserticapheaders": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preview": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"previewlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"queryparams": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reqtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reqtimeoutaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"useragent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsicapprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsicapprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nsicapprofileName := d.Get("name").(string)
	nsicapprofile := ns.Nsicapprofile{
		Allow204:            d.Get("allow204").(string),
		Connectionkeepalive: d.Get("connectionkeepalive").(string),
		Hostheader:          d.Get("hostheader").(string),
		Inserthttprequest:   d.Get("inserthttprequest").(string),
		Inserticapheaders:   d.Get("inserticapheaders").(string),
		Logaction:           d.Get("logaction").(string),
		Mode:                d.Get("mode").(string),
		Name:                d.Get("name").(string),
		Preview:             d.Get("preview").(string),
		Previewlength:       d.Get("previewlength").(int),
		Queryparams:         d.Get("queryparams").(string),
		Reqtimeout:          d.Get("reqtimeout").(int),
		Reqtimeoutaction:    d.Get("reqtimeoutaction").(string),
		Uri:                 d.Get("uri").(string),
		Useragent:           d.Get("useragent").(string),
	}

	_, err := client.AddResource("nsicapprofile", nsicapprofileName, &nsicapprofile)
	if err != nil {
		return err
	}

	d.SetId(nsicapprofileName)

	err = readNsicapprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsicapprofile but we can't read it ?? %s", nsicapprofileName)
		return nil
	}
	return nil
}

func readNsicapprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsicapprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nsicapprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsicapprofile state %s", nsicapprofileName)
	data, err := client.FindResource(("nsicapprofile"), nsicapprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsicapprofile state %s", nsicapprofileName)
		d.SetId("")
		return nil
	}
	d.Set("allow204", data["allow204"])
	d.Set("connectionkeepalive", data["connectionkeepalive"])
	d.Set("hostheader", data["hostheader"])
	d.Set("inserthttprequest", data["inserthttprequest"])
	d.Set("inserticapheaders", data["inserticapheaders"])
	d.Set("logaction", data["logaction"])
	d.Set("mode", data["mode"])
	d.Set("name", data["name"])
	d.Set("preview", data["preview"])
	d.Set("previewlength", data["previewlength"])
	d.Set("queryparams", data["queryparams"])
	d.Set("reqtimeout", data["reqtimeout"])
	d.Set("reqtimeoutaction", data["reqtimeoutaction"])
	d.Set("uri", data["uri"])
	d.Set("useragent", data["useragent"])

	return nil

}

func updateNsicapprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsicapprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nsicapprofileName := d.Get("name").(string)

	nsicapprofile := ns.Nsicapprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("allow204") {
		log.Printf("[DEBUG]  citrixadc-provider: Allow204 has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Allow204 = d.Get("allow204").(string)
		hasChange = true
	}
	if d.HasChange("connectionkeepalive") {
		log.Printf("[DEBUG]  citrixadc-provider: Connectionkeepalive has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Connectionkeepalive = d.Get("connectionkeepalive").(string)
		hasChange = true
	}
	if d.HasChange("hostheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Hostheader has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Hostheader = d.Get("hostheader").(string)
		hasChange = true
	}
	if d.HasChange("inserthttprequest") {
		log.Printf("[DEBUG]  citrixadc-provider: Inserthttprequest has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Inserthttprequest = d.Get("inserthttprequest").(string)
		hasChange = true
	}
	if d.HasChange("inserticapheaders") {
		log.Printf("[DEBUG]  citrixadc-provider: Inserticapheaders has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Inserticapheaders = d.Get("inserticapheaders").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("mode") {
		log.Printf("[DEBUG]  citrixadc-provider: Mode has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Mode = d.Get("mode").(string)
		hasChange = true
	}
	if d.HasChange("preview") {
		log.Printf("[DEBUG]  citrixadc-provider: Preview has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Preview = d.Get("preview").(string)
		hasChange = true
	}
	if d.HasChange("previewlength") {
		log.Printf("[DEBUG]  citrixadc-provider: Previewlength has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Previewlength = d.Get("previewlength").(int)
		hasChange = true
	}
	if d.HasChange("queryparams") {
		log.Printf("[DEBUG]  citrixadc-provider: Queryparams has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Queryparams = d.Get("queryparams").(string)
		hasChange = true
	}
	if d.HasChange("reqtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqtimeout has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Reqtimeout = d.Get("reqtimeout").(int)
		hasChange = true
	}
	if d.HasChange("reqtimeoutaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqtimeoutaction has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Reqtimeoutaction = d.Get("reqtimeoutaction").(string)
		hasChange = true
	}
	if d.HasChange("uri") {
		log.Printf("[DEBUG]  citrixadc-provider: Uri has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Uri = d.Get("uri").(string)
		hasChange = true
	}
	if d.HasChange("useragent") {
		log.Printf("[DEBUG]  citrixadc-provider: Useragent has changed for nsicapprofile %s, starting update", nsicapprofileName)
		nsicapprofile.Useragent = d.Get("useragent").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("nsicapprofile", nsicapprofileName, &nsicapprofile)
		if err != nil {
			return fmt.Errorf("Error updating nsicapprofile %s", nsicapprofileName)
		}
	}
	return readNsicapprofileFunc(d, meta)
}

func deleteNsicapprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsicapprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nsicapprofileName := d.Id()
	err := client.DeleteResource("nsicapprofile", nsicapprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
