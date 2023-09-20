package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsassignment() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsassignmentFunc,
		Read:          readNsassignmentFunc,
		Update:        updateNsassignmentFunc,
		Delete:        deleteNsassignmentFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"variable": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"add": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"append": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clear": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"set": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sub": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsassignmentFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsassignmentFunc")
	client := meta.(*NetScalerNitroClient).client
	nsassignmentName := d.Get("name").(string)
	nsassignment := ns.Nsassignment{
		Add:      d.Get("add").(string),
		Append:   d.Get("append").(string),
		Clear:    d.Get("clear").(bool),
		Comment:  d.Get("comment").(string),
		Name:     d.Get("name").(string),
		Newname:  d.Get("newname").(string),
		Set:      d.Get("set").(string),
		Sub:      d.Get("sub").(string),
		Variable: d.Get("variable").(string),
	}

	_, err := client.AddResource(service.Nsassignment.Type(), nsassignmentName, &nsassignment)
	if err != nil {
		return err
	}

	d.SetId(nsassignmentName)

	err = readNsassignmentFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsassignment but we can't read it ?? %s", nsassignmentName)
		return nil
	}
	return nil
}

func readNsassignmentFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsassignmentFunc")
	client := meta.(*NetScalerNitroClient).client
	nsassignmentName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsassignment state %s", nsassignmentName)
	data, err := client.FindResource(service.Nsassignment.Type(), nsassignmentName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsassignment state %s", nsassignmentName)
		d.SetId("")
		return nil
	}
	d.Set("add", data["add"])
	d.Set("append", data["append"])
	d.Set("clear", data["clear"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("set", data["set"])
	d.Set("sub", data["sub"])
	d.Set("variable", data["variable"])

	return nil

}

func updateNsassignmentFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsassignmentFunc")
	client := meta.(*NetScalerNitroClient).client
	nsassignmentName := d.Get("name").(string)

	nsassignment := ns.Nsassignment{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("add") {
		log.Printf("[DEBUG]  citrixadc-provider: Add has changed for nsassignment %s, starting update", nsassignmentName)
		nsassignment.Add = d.Get("add").(string)
		hasChange = true
	}
	if d.HasChange("append") {
		log.Printf("[DEBUG]  citrixadc-provider: Append has changed for nsassignment %s, starting update", nsassignmentName)
		nsassignment.Append = d.Get("append").(string)
		hasChange = true
	}
	if d.HasChange("clear") {
		log.Printf("[DEBUG]  citrixadc-provider: Clear has changed for nsassignment %s, starting update", nsassignmentName)
		nsassignment.Clear = d.Get("clear").(bool)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for nsassignment %s, starting update", nsassignmentName)
		nsassignment.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for nsassignment %s, starting update", nsassignmentName)
		nsassignment.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("set") {
		log.Printf("[DEBUG]  citrixadc-provider: Set has changed for nsassignment %s, starting update", nsassignmentName)
		nsassignment.Set = d.Get("set").(string)
		hasChange = true
	}
	if d.HasChange("sub") {
		log.Printf("[DEBUG]  citrixadc-provider: Sub has changed for nsassignment %s, starting update", nsassignmentName)
		nsassignment.Sub = d.Get("sub").(string)
		hasChange = true
	}
	if d.HasChange("variable") {
		log.Printf("[DEBUG]  citrixadc-provider: Variable has changed for nsassignment %s, starting update", nsassignmentName)
		nsassignment.Variable = d.Get("variable").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nsassignment.Type(), nsassignmentName, &nsassignment)
		if err != nil {
			return fmt.Errorf("Error updating nsassignment %s", nsassignmentName)
		}
	}
	return readNsassignmentFunc(d, meta)
}

func deleteNsassignmentFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsassignmentFunc")
	client := meta.(*NetScalerNitroClient).client
	nsassignmentName := d.Id()
	err := client.DeleteResource(service.Nsassignment.Type(), nsassignmentName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
