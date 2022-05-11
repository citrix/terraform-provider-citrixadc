package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcVrid6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVrid6Func,
		Read:          readVrid6Func,
		Update:        updateVrid6Func,
		Delete:        deleteVrid6Func,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vrid6_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"all": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ownernode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"preemption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preemptiondelaytimer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sharing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trackifnumpriority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tracking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVrid6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVrid6Func")
	client := meta.(*NetScalerNitroClient).client
	vrid6Id := d.Get("vrid6_id").(int)
	vrid6 := network.Vrid6{
		All:                  d.Get("all").(bool),
		Id:                   d.Get("vrid6_id").(int),
		Ownernode:            d.Get("ownernode").(int),
		Preemption:           d.Get("preemption").(string),
		Preemptiondelaytimer: d.Get("preemptiondelaytimer").(int),
		Priority:             d.Get("priority").(int),
		Sharing:              d.Get("sharing").(string),
		Trackifnumpriority:   d.Get("trackifnumpriority").(int),
		Tracking:             d.Get("tracking").(string),
	}
	vrid6IdStr := strconv.Itoa(vrid6Id)
	_, err := client.AddResource(service.Vrid6.Type(), vrid6IdStr, &vrid6)
	if err != nil {
		return err
	}

	d.SetId(vrid6IdStr)

	err = readVrid6Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vrid6 but we can't read it ?? %s", vrid6IdStr)
		return nil
	}
	return nil
}

func readVrid6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVrid6Func")
	client := meta.(*NetScalerNitroClient).client
	vrid6IdStr := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vrid6 state %s", vrid6IdStr)
	data, err := client.FindResource(service.Vrid6.Type(), vrid6IdStr)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vrid6 state %s", vrid6IdStr)
		d.SetId("")
		return nil
	}
	d.Set("all", data["all"])
	d.Set("vrid6_id", data["id"])
	d.Set("ownernode", data["ownernode"])
	d.Set("preemption", data["preemption"])
	d.Set("preemptiondelaytimer", data["preemptiondelaytimer"])
	d.Set("priority", data["priority"])
	d.Set("sharing", data["sharing"])
	d.Set("trackifnumpriority", data["trackifnumpriority"])
	d.Set("tracking", data["tracking"])

	return nil

}

func updateVrid6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVrid6Func")
	client := meta.(*NetScalerNitroClient).client
	vrid6Id := d.Get("vrid6_id").(int)

	vrid6 := network.Vrid6{
		Id: d.Get("vrid6_id").(int),
	}
	hasChange := false
	if d.HasChange("all") {
		log.Printf("[DEBUG]  citrixadc-provider: All has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.All = d.Get("all").(bool)
		hasChange = true
	}
	if d.HasChange("id") {
		log.Printf("[DEBUG]  citrixadc-provider: Id has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Id = d.Get("id").(int)
		hasChange = true
	}
	if d.HasChange("ownernode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ownernode has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Ownernode = d.Get("ownernode").(int)
		hasChange = true
	}
	if d.HasChange("preemption") {
		log.Printf("[DEBUG]  citrixadc-provider: Preemption has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Preemption = d.Get("preemption").(string)
		hasChange = true
	}
	if d.HasChange("preemptiondelaytimer") {
		log.Printf("[DEBUG]  citrixadc-provider: Preemptiondelaytimer has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Preemptiondelaytimer = d.Get("preemptiondelaytimer").(int)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("sharing") {
		log.Printf("[DEBUG]  citrixadc-provider: Sharing has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Sharing = d.Get("sharing").(string)
		hasChange = true
	}
	if d.HasChange("trackifnumpriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Trackifnumpriority has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Trackifnumpriority = d.Get("trackifnumpriority").(int)
		hasChange = true
	}
	if d.HasChange("tracking") {
		log.Printf("[DEBUG]  citrixadc-provider: Tracking has changed for vrid6 %d, starting update", vrid6Id)
		vrid6.Tracking = d.Get("tracking").(string)
		hasChange = true
	}
	vrid6IdStr := strconv.Itoa(vrid6Id)
	if hasChange {
		_, err := client.UpdateResource(service.Vrid6.Type(), vrid6IdStr, &vrid6)
		if err != nil {
			return fmt.Errorf("Error updating vrid6 %s", vrid6IdStr)
		}
	}
	return readVrid6Func(d, meta)
}

func deleteVrid6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVrid6Func")
	client := meta.(*NetScalerNitroClient).client
	vrid6Id := d.Id()
	err := client.DeleteResource(service.Vrid6.Type(), vrid6Id)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
