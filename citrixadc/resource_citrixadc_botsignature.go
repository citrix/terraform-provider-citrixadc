package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/bot"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcBotsignature() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBotsignatureFunc,
		Read:          readBotsignatureFunc,
		Delete:        deleteBotsignatureFunc,
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
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"overwrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"src": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBotsignatureFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotsignatureFunc")
	client := meta.(*NetScalerNitroClient).client

	botsignatureName := d.Get("name").(string)

	botsignature := bot.Botsignature{
		Comment:   d.Get("comment").(string),
		Name:      d.Get("name").(string),
		Overwrite: d.Get("overwrite").(bool),
		Src:       d.Get("src").(string),
	}

	err := client.ActOnResource("botsignature", &botsignature, "Import")
	if err != nil {
		return err
	}

	d.SetId(botsignatureName)

	err = readBotsignatureFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this botsignature but we can't read it ?? %s", botsignatureName)
		return nil
	}
	return nil
}

func readBotsignatureFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotsignatureFunc")
	client := meta.(*NetScalerNitroClient).client
	botsignatureName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading botsignature state %s", botsignatureName)
	data, err := client.FindResource("botsignature", botsignatureName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing botsignature state %s", botsignatureName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteBotsignatureFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotsignatureFunc")
	client := meta.(*NetScalerNitroClient).client
	botsignatureName := d.Id()
	err := client.DeleteResource("botsignature", botsignatureName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
