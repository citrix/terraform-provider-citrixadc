package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/rewrite"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcRewritepolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRewritepolicylabelFunc,
		Read:          readRewritepolicylabelFunc,
		Delete:        deleteRewritepolicylabelFunc,
		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createRewritepolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRewritepolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicylabelName := d.Get("labelname").(string)
	d.Set("name", rewritepolicylabelName)

	rewritepolicylabel := rewrite.Rewritepolicylabel{
		Comment:   d.Get("comment").(string),
		Labelname: d.Get("labelname").(string),
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
	d.Set("transform", data["transform"])

	return nil

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
