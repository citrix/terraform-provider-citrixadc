package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/transform"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcTransformaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createTransformactionFunc,
		Read:          readTransformactionFunc,
		Update:        updateTransformactionFunc,
		Delete:        deleteTransformactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookiedomainfrom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookiedomaininto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"profilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"requrlfrom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requrlinto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resurlfrom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resurlinto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createTransformactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createTransformactionFunc")
	client := meta.(*NetScalerNitroClient).client
	transformactionName := d.Get("name").(string)

	// Create does not support all attributes
	transformactionNew := transform.Transformaction{
		Name:        d.Get("name").(string),
		Priority:    d.Get("priority").(int),
		Profilename: d.Get("profilename").(string),
		State:       d.Get("state").(string),
	}

	_, err := client.AddResource(service.Transformaction.Type(), transformactionName, &transformactionNew)
	if err != nil {
		return err
	}

	// Need to update with full set of attributes
	transformaction := transform.Transformaction{
		Comment:          d.Get("comment").(string),
		Cookiedomainfrom: d.Get("cookiedomainfrom").(string),
		Cookiedomaininto: d.Get("cookiedomaininto").(string),
		Name:             d.Get("name").(string),
		Priority:         d.Get("priority").(int),
		Requrlfrom:       d.Get("requrlfrom").(string),
		Requrlinto:       d.Get("requrlinto").(string),
		Resurlfrom:       d.Get("resurlfrom").(string),
		Resurlinto:       d.Get("resurlinto").(string),
		State:            d.Get("state").(string),
	}

	_, err = client.UpdateResource(service.Transformaction.Type(), transformactionName, &transformaction)
	if err != nil {
		return err
	}

	d.SetId(transformactionName)

	err = readTransformactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this transformaction but we can't read it ?? %s", transformactionName)
		return nil
	}
	return nil
}

func readTransformactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readTransformactionFunc")
	client := meta.(*NetScalerNitroClient).client
	transformactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading transformaction state %s", transformactionName)
	data, err := client.FindResource(service.Transformaction.Type(), transformactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing transformaction state %s", transformactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("cookiedomainfrom", data["cookiedomainfrom"])
	d.Set("cookiedomaininto", data["cookiedomaininto"])
	d.Set("name", data["name"])
	d.Set("priority", data["priority"])
	d.Set("profilename", data["profilename"])
	d.Set("requrlfrom", data["requrlfrom"])
	d.Set("requrlinto", data["requrlinto"])
	d.Set("resurlfrom", data["resurlfrom"])
	d.Set("resurlinto", data["resurlinto"])
	d.Set("state", data["state"])

	return nil

}

func updateTransformactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTransformactionFunc")
	client := meta.(*NetScalerNitroClient).client
	transformactionName := d.Get("name").(string)

	transformaction := transform.Transformaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for transformaction %s, starting update", transformactionName)
		transformaction.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("cookiedomainfrom") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookiedomainfrom has changed for transformaction %s, starting update", transformactionName)
		transformaction.Cookiedomainfrom = d.Get("cookiedomainfrom").(string)
		hasChange = true
	}
	if d.HasChange("cookiedomaininto") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookiedomaininto has changed for transformaction %s, starting update", transformactionName)
		transformaction.Cookiedomaininto = d.Get("cookiedomaininto").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for transformaction %s, starting update", transformactionName)
		transformaction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for transformaction %s, starting update", transformactionName)
		transformaction.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("profilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Profilename has changed for transformaction %s, starting update", transformactionName)
		transformaction.Profilename = d.Get("profilename").(string)
		hasChange = true
	}
	if d.HasChange("requrlfrom") {
		log.Printf("[DEBUG]  citrixadc-provider: Requrlfrom has changed for transformaction %s, starting update", transformactionName)
		transformaction.Requrlfrom = d.Get("requrlfrom").(string)
		hasChange = true
	}
	if d.HasChange("requrlinto") {
		log.Printf("[DEBUG]  citrixadc-provider: Requrlinto has changed for transformaction %s, starting update", transformactionName)
		transformaction.Requrlinto = d.Get("requrlinto").(string)
		hasChange = true
	}
	if d.HasChange("resurlfrom") {
		log.Printf("[DEBUG]  citrixadc-provider: Resurlfrom has changed for transformaction %s, starting update", transformactionName)
		transformaction.Resurlfrom = d.Get("resurlfrom").(string)
		hasChange = true
	}
	if d.HasChange("resurlinto") {
		log.Printf("[DEBUG]  citrixadc-provider: Resurlinto has changed for transformaction %s, starting update", transformactionName)
		transformaction.Resurlinto = d.Get("resurlinto").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for transformaction %s, starting update", transformactionName)
		transformaction.State = d.Get("state").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Transformaction.Type(), transformactionName, &transformaction)
		if err != nil {
			return fmt.Errorf("Error updating transformaction %s", transformactionName)
		}
	}
	return readTransformactionFunc(d, meta)
}

func deleteTransformactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTransformactionFunc")
	client := meta.(*NetScalerNitroClient).client
	transformactionName := d.Id()
	err := client.DeleteResource(service.Transformaction.Type(), transformactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
