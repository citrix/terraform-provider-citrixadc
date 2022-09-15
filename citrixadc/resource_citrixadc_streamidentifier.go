package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/stream"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcStreamidentifier() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createStreamidentifierFunc,
		Read:          readStreamidentifierFunc,
		Update:        updateStreamidentifierFunc,
		Delete:        deleteStreamidentifierFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"acceptancethreshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"breachthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxtransactionthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mintransactionthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"samplecount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"selectorname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmptrap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sort": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trackackonlypackets": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tracktransactions": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createStreamidentifierFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createStreamidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	streamidentifierName := d.Get("name").(string)
	streamidentifier := stream.Streamidentifier{
		Acceptancethreshold:     d.Get("acceptancethreshold").(string),
		Appflowlog:              d.Get("appflowlog").(string),
		Breachthreshold:         d.Get("breachthreshold").(int),
		Interval:                d.Get("interval").(int),
		Maxtransactionthreshold: d.Get("maxtransactionthreshold").(int),
		Mintransactionthreshold: d.Get("mintransactionthreshold").(int),
		Name:                    d.Get("name").(string),
		Samplecount:             d.Get("samplecount").(int),
		Selectorname:            d.Get("selectorname").(string),
		Snmptrap:                d.Get("snmptrap").(string),
		Sort:                    d.Get("sort").(string),
		Trackackonlypackets:     d.Get("trackackonlypackets").(string),
		Tracktransactions:       d.Get("tracktransactions").(string),
	}

	_, err := client.AddResource(service.Streamidentifier.Type(), streamidentifierName, &streamidentifier)
	if err != nil {
		return err
	}

	d.SetId(streamidentifierName)

	err = readStreamidentifierFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this streamidentifier but we can't read it ?? %s", streamidentifierName)
		return nil
	}
	return nil
}

func readStreamidentifierFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readStreamidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	streamidentifierName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading streamidentifier state %s", streamidentifierName)
	data, err := client.FindResource(service.Streamidentifier.Type(), streamidentifierName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing streamidentifier state %s", streamidentifierName)
		d.SetId("")
		return nil
	}
	d.Set("acceptancethreshold", data["acceptancethreshold"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("breachthreshold", data["breachthreshold"])
	d.Set("interval", data["interval"])
	d.Set("maxtransactionthreshold", data["maxtransactionthreshold"])
	d.Set("mintransactionthreshold", data["mintransactionthreshold"])
	d.Set("name", data["name"])
	d.Set("samplecount", data["samplecount"])
	d.Set("selectorname", data["selectorname"])
	d.Set("snmptrap", data["snmptrap"])
	d.Set("sort", data["sort"])
	d.Set("trackackonlypackets", data["trackackonlypackets"])
	d.Set("tracktransactions", data["tracktransactions"])

	return nil

}

func updateStreamidentifierFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateStreamidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	streamidentifierName := d.Get("name").(string)

	streamidentifier := stream.Streamidentifier{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("acceptancethreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Acceptancethreshold has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Acceptancethreshold = d.Get("acceptancethreshold").(string)
		hasChange = true
	}
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowlog has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("breachthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Breachthreshold has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Breachthreshold = d.Get("breachthreshold").(int)
		hasChange = true
	}
	if d.HasChange("interval") {
		log.Printf("[DEBUG]  citrixadc-provider: Interval has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Interval = d.Get("interval").(int)
		hasChange = true
	}
	if d.HasChange("maxtransactionthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxtransactionthreshold has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Maxtransactionthreshold = d.Get("maxtransactionthreshold").(int)
		hasChange = true
	}
	if d.HasChange("mintransactionthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Mintransactionthreshold has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Mintransactionthreshold = d.Get("mintransactionthreshold").(int)
		hasChange = true
	}
	if d.HasChange("samplecount") {
		log.Printf("[DEBUG]  citrixadc-provider: Samplecount has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Samplecount = d.Get("samplecount").(int)
		hasChange = true
	}
	if d.HasChange("selectorname") {
		log.Printf("[DEBUG]  citrixadc-provider: Selectorname has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Selectorname = d.Get("selectorname").(string)
		hasChange = true
	}
	if d.HasChange("snmptrap") {
		log.Printf("[DEBUG]  citrixadc-provider: Snmptrap has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Snmptrap = d.Get("snmptrap").(string)
		hasChange = true
	}
	if d.HasChange("sort") {
		log.Printf("[DEBUG]  citrixadc-provider: Sort has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Sort = d.Get("sort").(string)
		hasChange = true
	}
	if d.HasChange("trackackonlypackets") {
		log.Printf("[DEBUG]  citrixadc-provider: Trackackonlypackets has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Trackackonlypackets = d.Get("trackackonlypackets").(string)
		hasChange = true
	}
	if d.HasChange("tracktransactions") {
		log.Printf("[DEBUG]  citrixadc-provider: Tracktransactions has changed for streamidentifier %s, starting update", streamidentifierName)
		streamidentifier.Tracktransactions = d.Get("tracktransactions").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Streamidentifier.Type(), &streamidentifier)
		if err != nil {
			return fmt.Errorf("Error updating streamidentifier %s", streamidentifierName)
		}
	}
	return readStreamidentifierFunc(d, meta)
}

func deleteStreamidentifierFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteStreamidentifierFunc")
	client := meta.(*NetScalerNitroClient).client
	streamidentifierName := d.Id()
	err := client.DeleteResource(service.Streamidentifier.Type(), streamidentifierName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
