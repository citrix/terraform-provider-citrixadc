package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/rewrite"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcRewritepolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRewritepolicylabelFunc,
		Read:          readRewritepolicylabelFunc,
		Update:        updateRewritepolicylabelFunc,
		Delete:        deleteRewritepolicylabelFunc,
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
			"transform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createRewritepolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRewritepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	var rewritepolicylabelName string
	if v, ok := d.GetOk("name"); ok {
		rewritepolicylabelName = v.(string)
	} else {
		rewritepolicylabelName = resource.PrefixedUniqueId("tf-rewritepolicylabel-")
		d.Set("name", rewritepolicylabelName)
	}
	rewritepolicylabel := rewrite.Rewritepolicylabel{
		Comment:   d.Get("comment").(string),
		Labelname: d.Get("labelname").(string),
		Newname:   d.Get("newname").(string),
		Transform: d.Get("transform").(string),
	}

	_, err := client.AddResource(netscaler.Rewritepolicylabel.Type(), rewritepolicylabelName, &rewritepolicylabel)
	if err != nil {
		return err
	}

	d.SetId(rewritepolicylabelName)

	err = readRewritepolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rewritepolicylabel but we can't read it ?? %s", rewritepolicylabelName)
		return nil
	}
	return nil
}

func readRewritepolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRewritepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rewritepolicylabel state %s", rewritepolicylabelName)
	data, err := client.FindResource(netscaler.Rewritepolicylabel.Type(), rewritepolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rewritepolicylabel state %s", rewritepolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("labelname", data["labelname"])
	d.Set("newname", data["newname"])
	d.Set("transform", data["transform"])

	return nil

}

func updateRewritepolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRewritepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicylabelName := d.Get("name").(string)

	rewritepolicylabel := rewrite.Rewritepolicylabel{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for rewritepolicylabel %s, starting update", rewritepolicylabelName)
		rewritepolicylabel.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("labelname") {
		log.Printf("[DEBUG]  citrixadc-provider: Labelname has changed for rewritepolicylabel %s, starting update", rewritepolicylabelName)
		rewritepolicylabel.Labelname = d.Get("labelname").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for rewritepolicylabel %s, starting update", rewritepolicylabelName)
		rewritepolicylabel.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("transform") {
		log.Printf("[DEBUG]  citrixadc-provider: Transform has changed for rewritepolicylabel %s, starting update", rewritepolicylabelName)
		rewritepolicylabel.Transform = d.Get("transform").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Rewritepolicylabel.Type(), rewritepolicylabelName, &rewritepolicylabel)
		if err != nil {
			return fmt.Errorf("Error updating rewritepolicylabel %s", rewritepolicylabelName)
		}
	}
	return readRewritepolicylabelFunc(d, meta)
}

func deleteRewritepolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewritepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicylabelName := d.Id()
	err := client.DeleteResource(netscaler.Rewritepolicylabel.Type(), rewritepolicylabelName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
