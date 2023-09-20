package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsservicefunction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsservicefunctionFunc,
		Read:          readNsservicefunctionFunc,
		Update:        updateNsservicefunctionFunc,
		Delete:        deleteNsservicefunctionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"servicefunctionname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"ingressvlan": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createNsservicefunctionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsservicefunctionFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicefunctionName := d.Get("servicefunctionname").(string)

	nsservicefunction := ns.Nsservicefunction{
		Ingressvlan:         d.Get("ingressvlan").(int),
		Servicefunctionname: d.Get("servicefunctionname").(string),
	}

	_, err := client.AddResource(service.Nsservicefunction.Type(), nsservicefunctionName, &nsservicefunction)
	if err != nil {
		return err
	}

	d.SetId(nsservicefunctionName)

	err = readNsservicefunctionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsservicefunction but we can't read it ?? %s", nsservicefunctionName)
		return nil
	}
	return nil
}

func readNsservicefunctionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsservicefunctionFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicefunctionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsservicefunction state %s", nsservicefunctionName)
	data, err := client.FindResource(service.Nsservicefunction.Type(), nsservicefunctionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsservicefunction state %s", nsservicefunctionName)
		d.SetId("")
		return nil
	}
	d.Set("ingressvlan", data["ingressvlan"])
	d.Set("servicefunctionname", data["servicefunctionname"])

	return nil

}

func updateNsservicefunctionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsservicefunctionFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicefunctionName := d.Get("servicefunctionname").(string)

	nsservicefunction := ns.Nsservicefunction{
		Servicefunctionname: d.Get("servicefunctionname").(string),
	}
	hasChange := false
	if d.HasChange("ingressvlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Ingressvlan has changed for nsservicefunction %s, starting update", nsservicefunctionName)
		nsservicefunction.Ingressvlan = d.Get("ingressvlan").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nsservicefunction.Type(), nsservicefunctionName, &nsservicefunction)
		if err != nil {
			return fmt.Errorf("Error updating nsservicefunction %s", nsservicefunctionName)
		}
	}
	return readNsservicefunctionFunc(d, meta)
}

func deleteNsservicefunctionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsservicefunctionFunc")
	client := meta.(*NetScalerNitroClient).client
	nsservicefunctionName := d.Id()
	err := client.DeleteResource(service.Nsservicefunction.Type(), nsservicefunctionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
