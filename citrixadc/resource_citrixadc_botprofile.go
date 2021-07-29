package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcBotprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBotprofileFunc,
		Read:          readBotprofileFunc,
		Update:        updateBotprofileFunc,
		Delete:        deleteBotprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"signature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"errorurl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trapurl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bot_enable_white_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bot_enable_black_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bot_enable_rate_limit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"devicefingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"devicefingerprintaction": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bot_enable_ip_reputation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trapaction": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bot_enable_tps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createBotprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createBotprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	var botprofileName string
	if v, ok := d.GetOk("name"); ok {
		botprofileName = v.(string)
	} else {
		botprofileName = resource.PrefixedUniqueId("tf-botprofile-")
		d.Set("name", botprofileName)
	}
	botprofileName = d.Get("name").(string)
	botprofile := bot.Botprofile{
		Name:                    d.Get("name").(string),
		Signature:               d.Get("signature").(string),
		Errorurl:                d.Get("errorurl").(string),
		Trapurl:                 d.Get("trapurl").(string),
		Comment:                 d.Get("comment").(string),
		Botenablewhitelist:      d.Get("bot_enable_white_list").(string),
		Botenableblacklist:      d.Get("bot_enable_black_list").(string),
		Botenableratelimit:      d.Get("bot_enable_rate_limit").(string),
		Devicefingerprint:       d.Get("devicefingerprint").(string),
		Devicefingerprintaction: toStringList(d.Get("devicefingerprintaction").([]interface{})),
		Botenableipreputation:   d.Get("bot_enable_ip_reputation").(string),
		Trap:                    d.Get("trap").(string),
		Trapaction:              toStringList(d.Get("trapaction").([]interface{})),
		Botenabletps:            d.Get("bot_enable_tps").(string),
	}

	_, err := client.AddResource("botprofile", botprofileName, &botprofile)
	if err != nil {
		return err
	}

	d.SetId(botprofileName)

	err = readBotprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this botprofile but we can't read it ?? %s", botprofileName)
		return nil
	}
	return nil
}

func readBotprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readBotprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	botprofileName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading botprofile state %s", botprofileName)
	data, err := client.FindResource("botprofile", botprofileName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing botprofile state %s", botprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("signature", data["signature"])
	d.Set("errorurl", data["errorurl"])
	d.Set("trapurl", data["trapurl"])
	d.Set("comment", data["comment"])
	d.Set("bot_enable_white_list", data["bot_enable_white_list"])
	d.Set("bot_enable_black_list", data["bot_enable_black_list"])
	d.Set("bot_enable_rate_limit", data["bot_enable_rate_limit"])
	d.Set("devicefingerprint", data["devicefingerprint"])
	d.Set("devicefingerprintaction", data["devicefingerprintaction"])
	d.Set("bot_enable_ip_reputation", data["bot_enable_ip_reputation"])
	d.Set("trap", data["trap"])
	d.Set("trapaction", data["trapaction"])
	d.Set("bot_enable_tps", data["bot_enable_tps"])

	return nil
}

func updateBotprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateBotprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	botprofileName := d.Get("name").(string)

	botprofile := bot.Botprofile{
		Name: d.Get("name").(string),
	}

	hasChange := false

	if d.HasChange("signature") {
		log.Printf("[DEBUG]  netscaler-provider: Signature has changed for botprofile %s, starting update", botprofileName)
		botprofile.Signature = d.Get("signature").(string)
		hasChange = true
	}
	if d.HasChange("errorurl") {
		log.Printf("[DEBUG]  netscaler-provider: Errorurl has changed for botprofile %s, starting update", botprofileName)
		botprofile.Errorurl = d.Get("errorurl").(string)
		hasChange = true
	}
	if d.HasChange("trapurl") {
		log.Printf("[DEBUG]  netscaler-provider: Trapurl has changed for botprofile %s, starting update", botprofileName)
		botprofile.Trapurl = d.Get("trapurl").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  netscaler-provider: Comment has changed for botprofile %s, starting update", botprofileName)
		botprofile.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("bot_enable_white_list") {
		log.Printf("[DEBUG]  netscaler-provider: Bot_enable_white_list has changed for botprofile %s, starting update", botprofileName)
		botprofile.Botenablewhitelist = d.Get("bot_enable_white_list").(string)
		hasChange = true
	}
	if d.HasChange("bot_enable_black_list") {
		log.Printf("[DEBUG]  netscaler-provider: Bot_enable_black_list has changed for botprofile %s, starting update", botprofileName)
		botprofile.Botenableblacklist = d.Get("bot_enable_black_list").(string)
		hasChange = true
	}
	if d.HasChange("bot_enable_rate_limit") {
		log.Printf("[DEBUG]  netscaler-provider: Bot_enable_rate_limit has changed for botprofile %s, starting update", botprofileName)
		botprofile.Botenableratelimit = d.Get("bot_enable_rate_limit").(string)
		hasChange = true
	}
	if d.HasChange("devicefingerprint") {
		log.Printf("[DEBUG]  netscaler-provider: Devicefingerprint has changed for botprofile %s, starting update", botprofileName)
		botprofile.Devicefingerprint = d.Get("devicefingerprint").(string)
		hasChange = true
	}
	if d.HasChange("devicefingerprintaction") {
		log.Printf("[DEBUG]  netscaler-provider: Devicefingerprintaction has changed for botprofile %s, starting update", botprofileName)
		hasChange = true
		botprofile.Devicefingerprintaction = toStringList(d.Get("devicefingerprintaction").([]interface{}))
	}
	if d.HasChange("bot_enable_ip_reputation") {
		log.Printf("[DEBUG]  netscaler-provider: Bot_enable_ip_reputation has changed for botprofile %s, starting update", botprofileName)
		botprofile.Botenableipreputation = d.Get("bot_enable_ip_reputation").(string)
		hasChange = true
	}
	if d.HasChange("trap") {
		log.Printf("[DEBUG]  netscaler-provider: Trap has changed for botprofile %s, starting update", botprofileName)
		botprofile.Trap = d.Get("trap").(string)
		hasChange = true
	}
	if d.HasChange("trapaction") {
		log.Printf("[DEBUG]  netscaler-provider: Trapaction has changed for botprofile %s, starting update", botprofileName)
		hasChange = true
		botprofile.Trapaction = toStringList(d.Get("trapaction").([]interface{}))
	}
	if d.HasChange("bot_enable_tps") {
		log.Printf("[DEBUG]  netscaler-provider: Bot_enable_tps has changed for botprofile %s, starting update", botprofileName)
		botprofile.Botenabletps = d.Get("bot_enable_tps").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("botprofile", botprofileName, &botprofile)
		if err != nil {
			return fmt.Errorf("Error updating botprofile %s", botprofileName)
		}
	}

	return readBotprofileFunc(d, meta)
}

func deleteBotprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteBotprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	botprofileName := d.Id()
	err := client.DeleteResource("botprofile", botprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
