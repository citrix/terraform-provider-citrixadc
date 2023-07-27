package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcBotpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBotpolicyFunc,
		Read:          readBotpolicyFunc,
		Update:        updateBotpolicyFunc,
		Delete:        deleteBotpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"profilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createBotpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	botpolicyName := d.Get("name").(string)

	botpolicy := bot.Botpolicy{
		Name:        botpolicyName,
		Rule:        d.Get("rule").(string),
		Profilename: d.Get("profilename").(string),
		Undefaction: d.Get("undefaction").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
	}

	_, err := client.AddResource("botpolicy", botpolicyName, &botpolicy)
	if err != nil {
		return err
	}

	d.SetId(botpolicyName)

	err = readBotpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this botpolicy but we can't read it ?? %s", botpolicyName)
		return nil
	}
	return nil
}

func readBotpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	botpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Readingbotpolicy state %s", botpolicyName)
	data, err := client.FindResource("botpolicy", botpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing botpolicy state %s", botpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])
	d.Set("profilename", data["profilename"])
	d.Set("undefaction", data["undefaction"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])

	return nil

}

func updateBotpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateBotpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	botpolicyName := d.Get("name").(string)

	botpolicy := bot.Botpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for botpolicy %s, starting update", botpolicyName)
		botpolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("profilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Profilename has changed for botpolicy %s, starting update", botpolicyName)
		botpolicy.Profilename = d.Get("profilename").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for botpolicy %s, starting update", botpolicyName)
		botpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for botpolicy %s, starting update", botpolicyName)
		botpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for botpolicy %s, starting update", botpolicyName)
		botpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for botpolicy %s, starting update", botpolicyName)
		botpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if hasChange {
		_, err := client.UpdateResource("botpolicy", botpolicyName, &botpolicy)
		if err != nil {
			return fmt.Errorf("Error updating botpolicy %s", botpolicyName)
		}
	}
	return readBotpolicyFunc(d, meta)
}

func deleteBotpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	botpolicyName := d.Id()
	err := client.DeleteResource("botpolicy", botpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
