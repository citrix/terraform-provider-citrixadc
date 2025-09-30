package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAppfwfieldtype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwfieldtypeFunc,
		Read:          readAppfwfieldtypeFunc,
		Update:        updateAppfwfieldtypeFunc,
		Delete:        deleteAppfwfieldtypeFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nocharmaps": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"regex": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createAppfwfieldtypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwfieldtypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwfieldtypeName := d.Get("name").(string)
	appfwfieldtype := appfw.Appfwfieldtype{
		Comment:    d.Get("comment").(string),
		Name:       appfwfieldtypeName,
		Nocharmaps: d.Get("nocharmaps").(bool),
		Priority:   d.Get("priority").(int),
		Regex:      d.Get("regex").(string),
	}

	_, err := client.AddResource(service.Appfwfieldtype.Type(), appfwfieldtypeName, &appfwfieldtype)
	if err != nil {
		return err
	}

	d.SetId(appfwfieldtypeName)

	err = readAppfwfieldtypeFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwfieldtype but we can't read it ?? %s", appfwfieldtypeName)
		return nil
	}
	return nil
}

func readAppfwfieldtypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwfieldtypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwfieldtypeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwfieldtype state %s", appfwfieldtypeName)
	data, err := client.FindResource(service.Appfwfieldtype.Type(), appfwfieldtypeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwfieldtype state %s", appfwfieldtypeName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("nocharmaps", data["nocharmaps"])
	setToInt("priority", d, data["priority"])
	d.Set("regex", data["regex"])

	return nil

}

func updateAppfwfieldtypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppfwfieldtypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwfieldtypeName := d.Get("name").(string)

	appfwfieldtype := appfw.Appfwfieldtype{
		Name:  d.Get("name").(string),
		Regex: d.Get("regex").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for appfwfieldtype %s, starting update", appfwfieldtypeName)
		appfwfieldtype.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("nocharmaps") {
		log.Printf("[DEBUG]  citrixadc-provider: Nocharmaps has changed for appfwfieldtype %s, starting update", appfwfieldtypeName)
		appfwfieldtype.Nocharmaps = d.Get("nocharmaps").(bool)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for appfwfieldtype %s, starting update", appfwfieldtypeName)
		appfwfieldtype.Priority = d.Get("priority").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Appfwfieldtype.Type(), appfwfieldtypeName, &appfwfieldtype)
		if err != nil {
			return fmt.Errorf("Error updating appfwfieldtype %s", appfwfieldtypeName)
		}
	}
	return readAppfwfieldtypeFunc(d, meta)
}

func deleteAppfwfieldtypeFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwfieldtypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwfieldtypeName := d.Id()
	err := client.DeleteResource(service.Appfwfieldtype.Type(), appfwfieldtypeName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
