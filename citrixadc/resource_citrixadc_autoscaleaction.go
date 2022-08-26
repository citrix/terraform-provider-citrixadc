package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/autoscale"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAutoscaleaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAutoscaleactionFunc,
		Read:          readAutoscaleactionFunc,
		Update:        updateAutoscaleactionFunc,
		Delete:        deleteAutoscaleactionFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quiettime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vmdestroygraceperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAutoscaleactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAutoscaleactionFunc")
	client := meta.(*NetScalerNitroClient).client
	var autoscaleactionName string
	if v, ok := d.GetOk("name"); ok {
		autoscaleactionName = v.(string)
	} else {
		autoscaleactionName = resource.PrefixedUniqueId("tf-autoscaleaction-")
		d.Set("name", autoscaleactionName)
	}
	autoscaleaction := autoscale.Autoscaleaction{
		Name:                 d.Get("name").(string),
		Parameters:           d.Get("parameters").(string),
		Profilename:          d.Get("profilename").(string),
		Quiettime:            d.Get("quiettime").(int),
		Type:                 d.Get("type").(string),
		Vmdestroygraceperiod: d.Get("vmdestroygraceperiod").(int),
		Vserver:              d.Get("vserver").(string),
	}

	_, err := client.AddResource(service.Autoscaleaction.Type(), autoscaleactionName, &autoscaleaction)
	if err != nil {
		return err
	}

	d.SetId(autoscaleactionName)

	err = readAutoscaleactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this autoscaleaction but we can't read it ?? %s", autoscaleactionName)
		return nil
	}
	return nil
}

func readAutoscaleactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAutoscaleactionFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading autoscaleaction state %s", autoscaleactionName)
	data, err := client.FindResource(service.Autoscaleaction.Type(), autoscaleactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing autoscaleaction state %s", autoscaleactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("name", data["name"])
	d.Set("parameters", data["parameters"])
	d.Set("profilename", data["profilename"])
	d.Set("quiettime", data["quiettime"])
	d.Set("type", data["type"])
	d.Set("vmdestroygraceperiod", data["vmdestroygraceperiod"])
	d.Set("vserver", data["vserver"])

	return nil

}

func updateAutoscaleactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAutoscaleactionFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleactionName := d.Get("name").(string)

	autoscaleaction := autoscale.Autoscaleaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG]  citrixadc-provider: Parameters has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Parameters = d.Get("parameters").(string)
		hasChange = true
	}
	if d.HasChange("profilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Profilename has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Profilename = d.Get("profilename").(string)
		hasChange = true
	}
	if d.HasChange("quiettime") {
		log.Printf("[DEBUG]  citrixadc-provider: Quiettime has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Quiettime = d.Get("quiettime").(int)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("vmdestroygraceperiod") {
		log.Printf("[DEBUG]  citrixadc-provider: Vmdestroygraceperiod has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Vmdestroygraceperiod = d.Get("vmdestroygraceperiod").(int)
		hasChange = true
	}
	if d.HasChange("vserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Vserver has changed for autoscaleaction %s, starting update", autoscaleactionName)
		autoscaleaction.Vserver = d.Get("vserver").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Autoscaleaction.Type(), autoscaleactionName, &autoscaleaction)
		if err != nil {
			return fmt.Errorf("Error updating autoscaleaction %s", autoscaleactionName)
		}
	}
	return readAutoscaleactionFunc(d, meta)
}

func deleteAutoscaleactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAutoscaleactionFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleactionName := d.Id()
	err := client.DeleteResource(service.Autoscaleaction.Type(), autoscaleactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
