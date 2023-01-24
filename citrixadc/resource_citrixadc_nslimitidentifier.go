package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNslimitidentifier() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNslimitidentifierFunc,
		Read:          readNslimitidentifierFunc,
		Update:        updateNslimitidentifierFunc,
		Delete:        deleteNslimitidentifierFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"limitidentifier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"limittype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"selectorname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeslice": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trapsintimeslice": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNslimitidentifierFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNslimitidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	nslimitidentifierName := d.Get("limitidentifier").(string)

	nslimitidentifier := ns.Nslimitidentifier{
		Limitidentifier:  d.Get("limitidentifier").(string),
		Limittype:        d.Get("limittype").(string),
		Maxbandwidth:     d.Get("maxbandwidth").(int),
		Mode:             d.Get("mode").(string),
		Selectorname:     d.Get("selectorname").(string),
		Threshold:        d.Get("threshold").(int),
		Timeslice:        d.Get("timeslice").(int),
		Trapsintimeslice: d.Get("trapsintimeslice").(int),
	}

	_, err := client.AddResource(service.Nslimitidentifier.Type(), nslimitidentifierName, &nslimitidentifier)
	if err != nil {
		return err
	}

	d.SetId(nslimitidentifierName)

	err = readNslimitidentifierFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nslimitidentifier but we can't read it ?? %s", nslimitidentifierName)
		return nil
	}
	return nil
}

func readNslimitidentifierFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNslimitidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	nslimitidentifierName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nslimitidentifier state %s", nslimitidentifierName)
	data, err := client.FindResource(service.Nslimitidentifier.Type(), nslimitidentifierName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nslimitidentifier state %s", nslimitidentifierName)
		d.SetId("")
		return nil
	}
	d.Set("limitidentifier", data["limitidentifier"])
	d.Set("limittype", data["limittype"])
	d.Set("maxbandwidth", data["maxbandwidth"])
	d.Set("mode", data["mode"])
	d.Set("selectorname", data["selectorname"])
	d.Set("threshold", data["threshold"])
	d.Set("timeslice", data["timeslice"])
	d.Set("trapsintimeslice", data["trapsintimeslice"])

	return nil

}

func updateNslimitidentifierFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNslimitidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	nslimitidentifierName := d.Get("limitidentifier").(string)

	nslimitidentifier := ns.Nslimitidentifier{
		Limitidentifier: d.Get("limitidentifier").(string),
	}
	hasChange := false
	if d.HasChange("limittype") {
		log.Printf("[DEBUG]  citrixadc-provider: Limittype has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Limittype = d.Get("limittype").(string)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxbandwidth has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Maxbandwidth = d.Get("maxbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("mode") {
		log.Printf("[DEBUG]  citrixadc-provider: Mode has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Mode = d.Get("mode").(string)
		hasChange = true
	}
	if d.HasChange("selectorname") {
		log.Printf("[DEBUG]  citrixadc-provider: Selectorname has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Selectorname = d.Get("selectorname").(string)
		hasChange = true
	}
	if d.HasChange("threshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Threshold has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Threshold = d.Get("threshold").(int)
		hasChange = true
	}
	if d.HasChange("timeslice") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeslice has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Timeslice = d.Get("timeslice").(int)
		hasChange = true
	}
	if d.HasChange("trapsintimeslice") {
		log.Printf("[DEBUG]  citrixadc-provider: Trapsintimeslice has changed for nslimitidentifier %s, starting update", nslimitidentifierName)
		nslimitidentifier.Trapsintimeslice = d.Get("trapsintimeslice").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nslimitidentifier.Type(), nslimitidentifierName, &nslimitidentifier)
		if err != nil {
			return fmt.Errorf("Error updating nslimitidentifier %s", nslimitidentifierName)
		}
	}
	return readNslimitidentifierFunc(d, meta)
}

func deleteNslimitidentifierFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNslimitidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	nslimitidentifierName := d.Id()
	err := client.DeleteResource(service.Nslimitidentifier.Type(), nslimitidentifierName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
