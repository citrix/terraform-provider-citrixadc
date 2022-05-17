package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNspartition() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNspartitionFunc,
		Read:          readNspartitionFunc,
		Update:        updateNspartitionFunc,
		Delete:        deleteNspartitionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"partitionname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"force": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"maxbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxconn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxmemlimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"partitionmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"save": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNspartitionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNspartitionFunc")
	client := meta.(*NetScalerNitroClient).client
	nspartitionName := d.Get("partitionname").(string)
	nspartition := ns.Nspartition{
		Force:         d.Get("force").(bool),
		Maxbandwidth:  d.Get("maxbandwidth").(int),
		Maxconn:       d.Get("maxconn").(int),
		Maxmemlimit:   d.Get("maxmemlimit").(int),
		Minbandwidth:  d.Get("minbandwidth").(int),
		Partitionmac:  d.Get("partitionmac").(string),
		Partitionname: d.Get("partitionname").(string),
		Save:          d.Get("save").(bool),
	}

	_, err := client.AddResource(service.Nspartition.Type(), nspartitionName, &nspartition)
	if err != nil {
		return err
	}

	d.SetId(nspartitionName)

	err = readNspartitionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nspartition but we can't read it ?? %s", nspartitionName)
		return nil
	}
	return nil
}

func readNspartitionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNspartitionFunc")
	client := meta.(*NetScalerNitroClient).client
	nspartitionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nspartition state %s", nspartitionName)
	data, err := client.FindResource(service.Nspartition.Type(), nspartitionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nspartition state %s", nspartitionName)
		d.SetId("")
		return nil
	}
	d.Set("force", data["force"])
	d.Set("maxbandwidth", data["maxbandwidth"])
	d.Set("maxconn", data["maxconn"])
	d.Set("maxmemlimit", data["maxmemlimit"])
	// d.Set("minbandwidth", data["minbandwidth"])
	d.Set("partitionmac", data["partitionmac"])
	d.Set("partitionname", data["partitionname"])
	d.Set("save", data["save"])

	return nil

}

func updateNspartitionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNspartitionFunc")
	client := meta.(*NetScalerNitroClient).client
	nspartitionName := d.Get("partitionname").(string)

	nspartition := ns.Nspartition{
		Partitionname: d.Get("partitionname").(string),
	}
	hasChange := false
	if d.HasChange("force") {
		log.Printf("[DEBUG]  citrixadc-provider: Force has changed for nspartition %s, starting update", nspartitionName)
		nspartition.Force = d.Get("force").(bool)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxbandwidth has changed for nspartition %s, starting update", nspartitionName)
		nspartition.Maxbandwidth = d.Get("maxbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("maxconn") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxconn has changed for nspartition %s, starting update", nspartitionName)
		nspartition.Maxconn = d.Get("maxconn").(int)
		hasChange = true
	}
	if d.HasChange("maxmemlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxmemlimit has changed for nspartition %s, starting update", nspartitionName)
		nspartition.Maxmemlimit = d.Get("maxmemlimit").(int)
		hasChange = true
	}
	if d.HasChange("minbandwidth") {
		log.Printf("[DEBUG]  citrixadc-provider: Minbandwidth has changed for nspartition %s, starting update", nspartitionName)
		nspartition.Minbandwidth = d.Get("minbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("partitionmac") {
		log.Printf("[DEBUG]  citrixadc-provider: Partitionmac has changed for nspartition %s, starting update", nspartitionName)
		nspartition.Partitionmac = d.Get("partitionmac").(string)
		hasChange = true
	}
	if d.HasChange("save") {
		log.Printf("[DEBUG]  citrixadc-provider: Save has changed for nspartition %s, starting update", nspartitionName)
		nspartition.Save = d.Get("save").(bool)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nspartition.Type(), nspartitionName, &nspartition)
		if err != nil {
			return fmt.Errorf("Error updating nspartition %s", nspartitionName)
		}
	}
	return readNspartitionFunc(d, meta)
}

func deleteNspartitionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNspartitionFunc")
	client := meta.(*NetScalerNitroClient).client
	nspartitionName := d.Id()
	err := client.DeleteResource(service.Nspartition.Type(), nspartitionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
