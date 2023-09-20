package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcBotpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBotpolicylabelFunc,
		Read:          readBotpolicylabelFunc,
		Delete:        deleteBotpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBotpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createBotpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	var Botpolicylabelname string
	Botpolicylabelname = d.Get("labelname").(string)
	Botpolicylabel := bot.Botpolicylabel{
		Labelname: d.Get("labelname").(string),
		Comment:   d.Get("comment").(string),
	}

	_, err := client.AddResource("botpolicylabel", Botpolicylabelname, &Botpolicylabel)
	if err != nil {
		return err
	}

	d.SetId(Botpolicylabelname)

	err = readBotpolicylabelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this Botpolicylabel but we can't read it ?? %s", Botpolicylabelname)
		return nil
	}
	return nil
}

func readBotpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readBotpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	Botpolicylabelname := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading Botpolicylabel state %s", Botpolicylabelname)
	data, err := client.FindResource("botpolicylabel", Botpolicylabelname)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing Botpolicylabel state %s", Botpolicylabelname)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("comment", data["comment"])

	return nil
}

func deleteBotpolicylabelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteBotpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	Botpolicylabelname := d.Id()
	err := client.DeleteResource("botpolicylabel", Botpolicylabelname)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
