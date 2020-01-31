package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/responder"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcResponderpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createResponderpolicylabelFunc,
		Read:          readResponderpolicylabelFunc,
		Update:        updateResponderpolicylabelFunc,
		Delete:        deleteResponderpolicylabelFunc,
		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createResponderpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createResponderpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	var responderpolicylabelName string
	if v, ok := d.GetOk("name"); ok {
		responderpolicylabelName = v.(string)
	} else {
		responderpolicylabelName = resource.PrefixedUniqueId("tf-responderpolicylabel-")
		d.Set("name", responderpolicylabelName)
	}
	responderpolicylabel := responder.Responderpolicylabel{
		Comment:         d.Get("comment").(string),
		Labelname:       d.Get("labelname").(string),
		Newname:         d.Get("newname").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	_, err := client.AddResource(netscaler.Responderpolicylabel.Type(), responderpolicylabelName, &responderpolicylabel)
	if err != nil {
		return err
	}

	d.SetId(responderpolicylabelName)

	err = readResponderpolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this responderpolicylabel but we can't read it ?? %s", responderpolicylabelName)
		return nil
	}
	return nil
}

func readResponderpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readResponderpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading responderpolicylabel state %s", responderpolicylabelName)
	data, err := client.FindResource(netscaler.Responderpolicylabel.Type(), responderpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing responderpolicylabel state %s", responderpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("labelname", data["labelname"])
	d.Set("newname", data["newname"])
	d.Set("policylabeltype", data["policylabeltype"])

	return nil

}

func updateResponderpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateResponderpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicylabelName := d.Get("name").(string)

	responderpolicylabel := responder.Responderpolicylabel{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for responderpolicylabel %s, starting update", responderpolicylabelName)
		responderpolicylabel.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("labelname") {
		log.Printf("[DEBUG]  citrixadc-provider: Labelname has changed for responderpolicylabel %s, starting update", responderpolicylabelName)
		responderpolicylabel.Labelname = d.Get("labelname").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for responderpolicylabel %s, starting update", responderpolicylabelName)
		responderpolicylabel.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("policylabeltype") {
		log.Printf("[DEBUG]  citrixadc-provider: Policylabeltype has changed for responderpolicylabel %s, starting update", responderpolicylabelName)
		responderpolicylabel.Policylabeltype = d.Get("policylabeltype").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Responderpolicylabel.Type(), responderpolicylabelName, &responderpolicylabel)
		if err != nil {
			return fmt.Errorf("Error updating responderpolicylabel %s", responderpolicylabelName)
		}
	}
	return readResponderpolicylabelFunc(d, meta)
}

func deleteResponderpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteResponderpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicylabelName := d.Id()
	err := client.DeleteResource(netscaler.Responderpolicylabel.Type(), responderpolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
