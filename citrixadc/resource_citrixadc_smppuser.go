package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/smpp"
	
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)


func resourceCitrixAdcSmppuser() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSmppuserFunc,
		Read:          readSmppuserFunc,
		Update:        updateSmppuserFunc,
		Delete:        deleteSmppuserFunc,
		Schema: map[string]*schema.Schema{
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Sensitive: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			
			
		},
	}
}

func createSmppuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSmppuserFunc")
	client := meta.(*NetScalerNitroClient).client
	var smppuserName string
	if v, ok := d.GetOk("username"); ok {
		smppuserName = v.(string)
	} else {
		smppuserName= resource.PrefixedUniqueId("tf-smppuser-")
		d.Set("username", smppuserName)
	}
	smppuser := smpp.Smppuser{
		Password:           d.Get("password").(string),
		Username:           d.Get("username").(string),
		
	}

	_, err := client.AddResource("smppuser", smppuserName, &smppuser)
	if err != nil {
		return err
	}

	d.SetId(smppuserName)
	
	err = readSmppuserFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this smppuser but we can't read it ?? %s", smppuserName)
		return nil
	}
	return nil
}

func readSmppuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSmppuserFunc")
	client := meta.(*NetScalerNitroClient).client
	smppuserName:= d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading smppuser state %s", smppuserName)
	data, err := client.FindResource("smppuser", smppuserName)
	if err != nil {
	log.Printf("[WARN] citrixadc-provider: Clearing smppuser state %s", smppuserName)
		d.SetId("")
		return nil
	}
	d.Set("username", data["username"])
	

	return nil

}

func updateSmppuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSmppuserFunc")
	client := meta.(*NetScalerNitroClient).client
	smppuserName := d.Get("username").(string)

	smppuser := smpp.Smppuser{
		Username : d.Get("username").(string),
	}
	hasChange := false
	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for smppuser %s, starting update", smppuserName)
		smppuser.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("username") {
		log.Printf("[DEBUG]  citrixadc-provider: Username has changed for smppuser %s, starting update", smppuserName)
		smppuser.Username = d.Get("username").(string)
		hasChange = true
	}
	

	if hasChange {
		_, err := client.UpdateResource("smppuser", smppuserName, &smppuser)
		if err != nil {
			return fmt.Errorf("Error updating smppuser %s", smppuserName)
		}
	}
	return readSmppuserFunc(d, meta)
}

func deleteSmppuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSmppuserFunc")
	client := meta.(*NetScalerNitroClient).client
	smppuserName := d.Id()
	err := client.DeleteResource("smppuser", smppuserName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
