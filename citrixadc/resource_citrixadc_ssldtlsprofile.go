package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSsldtlsprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSsldtlsprofileFunc,
		Read:          readSsldtlsprofileFunc,
		Update:        updateSsldtlsprofileFunc,
		Delete:        deleteSsldtlsprofileFunc,
		Schema: map[string]*schema.Schema{
			"helloverifyrequest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxbadmacignorecount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxholdqlen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxpacketsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxrecordsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxretrytime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pmtudiscovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"terminatesession": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSsldtlsprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
		Helloverifyrequest:   d.Get("helloverifyrequest").(string),
		Maxbadmacignorecount: d.Get("maxbadmacignorecount").(int),
		Maxholdqlen:          d.Get("maxholdqlen").(int),
		Maxpacketsize:        d.Get("maxpacketsize").(int),
		Maxrecordsize:        d.Get("maxrecordsize").(int),
		Maxretrytime:         d.Get("maxretrytime").(int),
		Name:                 ssldtlsprofileName,
		Pmtudiscovery:        d.Get("pmtudiscovery").(string),
		Terminatesession:     d.Get("terminatesession").(string),
	}

	_, err := client.AddResource(service.Ssldtlsprofile.Type(), ssldtlsprofileName, &ssldtlsprofile)
	if err != nil {
		return err
	}

	d.SetId(ssldtlsprofileName)

	err = readSsldtlsprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ssldtlsprofile but we can't read it ?? %s", ssldtlsprofileName)
		return nil
	}
	return nil
}

func readSsldtlsprofileFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("helloverifyrequest", data["helloverifyrequest"])
	d.Set("maxbadmacignorecount", data["maxbadmacignorecount"])
	d.Set("maxholdqlen", data["maxholdqlen"])
	d.Set("maxpacketsize", data["maxpacketsize"])
	d.Set("maxrecordsize", data["maxrecordsize"])
	d.Set("maxretrytime", data["maxretrytime"])
	d.Set("name", data["name"])
	d.Set("pmtudiscovery", data["pmtudiscovery"])
	d.Set("terminatesession", data["terminatesession"])

	return nil

}

func updateSsldtlsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSsldtlsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ssldtlsprofileName := d.Get("name").(string)

	ssldtlsprofile := ssl.Ssldtlsprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("helloverifyrequest") {
		log.Printf("[DEBUG]  citrixadc-provider: Helloverifyrequest has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Helloverifyrequest = d.Get("helloverifyrequest").(string)
		hasChange = true
	}
	if d.HasChange("maxbadmacignorecount") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxbadmacignorecount has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Maxbadmacignorecount = d.Get("maxbadmacignorecount").(int)
		hasChange = true
	}
	if d.HasChange("maxholdqlen") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxholdqlen has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Maxholdqlen = d.Get("maxholdqlen").(int)
		hasChange = true
	}
	if d.HasChange("maxpacketsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxpacketsize has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Maxpacketsize = d.Get("maxpacketsize").(int)
		hasChange = true
	}
	if d.HasChange("maxrecordsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxrecordsize has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Maxrecordsize = d.Get("maxrecordsize").(int)
		hasChange = true
	}
	if d.HasChange("maxretrytime") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxretrytime has changed for ssldtlsprofile %s, starting update", ssldtlsprofileName)
		ssldtlsprofile.Maxretrytime = d.Get("maxretrytime").(int)
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
			return fmt.Errorf("Error updating ssldtlsprofile %s", ssldtlsprofileName)
		}
	}
	return readSsldtlsprofileFunc(d, meta)
}

func deleteSsldtlsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSsldtlsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	ssldtlsprofileName := d.Id()
	err := client.DeleteResource(service.Ssldtlsprofile.Type(), ssldtlsprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
