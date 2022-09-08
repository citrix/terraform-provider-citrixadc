package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/db"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcDbuser() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDbuserFunc,
		Read:          readDbuserFunc,
		Update:        updateDbuserFunc,
		Delete:        deleteDbuserFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"loggedin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDbuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDbuserFunc")
	client := meta.(*NetScalerNitroClient).client
	dbuserName := d.Get("username").(string)
	dbuser := db.Dbuser{
		Password: d.Get("password").(string),
		Username: d.Get("username").(string),
	}

	_, err := client.AddResource(service.Dbuser.Type(), dbuserName, &dbuser)
	if err != nil {
		return err
	}

	d.SetId(dbuserName)

	err = readDbuserFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dbuser but we can't read it ?? %s", dbuserName)
		return nil
	}
	return nil
}

func readDbuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDbuserFunc")
	client := meta.(*NetScalerNitroClient).client
	dbuserName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dbuser state %s", dbuserName)
	data, err := client.FindResource(service.Dbuser.Type(), dbuserName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dbuser state %s", dbuserName)
		d.SetId("")
		return nil
	}
	d.Set("username", data["username"])
	d.Set("loggedin", data["loggedin"])
	//d.Set("password", data["password"])

	return nil

}

func updateDbuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDbuserFunc")
	client := meta.(*NetScalerNitroClient).client
	dbuserName := d.Get("username").(string)

	dbuser := db.Dbuser{
		Username: d.Get("username").(string),
	}
	hasChange := false
	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for dbuser %s, starting update", dbuserName)
		dbuser.Password = d.Get("password").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Dbuser.Type(), &dbuser)
		if err != nil {
			return fmt.Errorf("Error updating dbuser %s", dbuserName)
		}
	}
	return readDbuserFunc(d, meta)
}

func deleteDbuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDbuserFunc")
	client := meta.(*NetScalerNitroClient).client
	dbuserName := d.Id()
	err := client.DeleteResource(service.Dbuser.Type(), dbuserName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
