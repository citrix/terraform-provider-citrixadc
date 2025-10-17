package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcBotprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createBotprofileFunc,
		ReadContext:   readBotprofileFunc,
		UpdateContext: updateBotprofileFunc,
		DeleteContext: deleteBotprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bot_enable_black_list": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bot_enable_ip_reputation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bot_enable_rate_limit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bot_enable_tps": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bot_enable_white_list": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientipexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"devicefingerprint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"devicefingerprintaction": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"devicefingerprintmobile": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"errorurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kmdetection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kmeventspostbodylimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"kmjavascriptname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signature": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signaturemultipleuseragentheaderaction": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"signaturenouseragentheaderaction": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trap": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trapaction": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trapurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createBotprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotprofileFunc")
	client := meta.(*NetScalerNitroClient).client

	botprofileName := d.Get("name").(string)
	botprofile := bot.Botprofile{
		Botenableblacklist:                     d.Get("bot_enable_black_list").(string),
		Botenableipreputation:                  d.Get("bot_enable_ip_reputation").(string),
		Botenableratelimit:                     d.Get("bot_enable_rate_limit").(string),
		Botenabletps:                           d.Get("bot_enable_tps").(string),
		Botenablewhitelist:                     d.Get("bot_enable_white_list").(string),
		Clientipexpression:                     d.Get("clientipexpression").(string),
		Comment:                                d.Get("comment").(string),
		Devicefingerprint:                      d.Get("devicefingerprint").(string),
		Devicefingerprintaction:                toStringList(d.Get("devicefingerprintaction").([]interface{})),
		Devicefingerprintmobile:                toStringList(d.Get("devicefingerprintmobile").([]interface{})),
		Errorurl:                               d.Get("errorurl").(string),
		Kmdetection:                            d.Get("kmdetection").(string),
		Kmjavascriptname:                       d.Get("kmjavascriptname").(string),
		Name:                                   d.Get("name").(string),
		Signature:                              d.Get("signature").(string),
		Signaturemultipleuseragentheaderaction: toStringList(d.Get("signaturemultipleuseragentheaderaction").([]interface{})),
		Signaturenouseragentheaderaction:       toStringList(d.Get("signaturenouseragentheaderaction").([]interface{})),
		Trap:                                   d.Get("trap").(string),
		Trapaction:                             toStringList(d.Get("trapaction").([]interface{})),
		Trapurl:                                d.Get("trapurl").(string),
	}

	if raw := d.GetRawConfig().GetAttr("kmeventspostbodylimit"); !raw.IsNull() {
		botprofile.Kmeventspostbodylimit = intPtr(d.Get("kmeventspostbodylimit").(int))
	}

	_, err := client.AddResource("botprofile", botprofileName, &botprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(botprofileName)

	return readBotprofileFunc(ctx, d, meta)
}

func readBotprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	botprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading botprofile state %s", botprofileName)
	data, err := client.FindResource("botprofile", botprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile state %s", botprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("bot_enable_black_list", data["bot_enable_black_list"])
	d.Set("bot_enable_ip_reputation", data["bot_enable_ip_reputation"])
	d.Set("bot_enable_rate_limit", data["bot_enable_rate_limit"])
	d.Set("bot_enable_tps", data["bot_enable_tps"])
	d.Set("bot_enable_white_list", data["bot_enable_white_list"])
	d.Set("clientipexpression", data["clientipexpression"])
	d.Set("comment", data["comment"])
	d.Set("devicefingerprint", data["devicefingerprint"])
	d.Set("devicefingerprintaction", data["devicefingerprintaction"])
	d.Set("devicefingerprintmobile", data["devicefingerprintmobile"])
	d.Set("errorurl", data["errorurl"])
	d.Set("kmdetection", data["kmdetection"])
	setToInt("kmeventspostbodylimit", d, data["kmeventspostbodylimit"])
	d.Set("kmjavascriptname", data["kmjavascriptname"])
	d.Set("signature", data["signature"])
	d.Set("signaturemultipleuseragentheaderaction", data["signaturemultipleuseragentheaderaction"])
	d.Set("signaturenouseragentheaderaction", data["signaturenouseragentheaderaction"])
	d.Set("trap", data["trap"])
	d.Set("trapaction", data["trapaction"])
	d.Set("trapurl", data["trapurl"])

	return nil

}

func updateBotprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateBotprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	botprofileName := d.Get("name").(string)

	botprofile := bot.Botprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("bot_enable_black_list") {
		log.Printf("[DEBUG]  citrixadc-provider: Botenableblacklist has changed for botprofile %s, starting update", botprofileName)
		botprofile.Botenableblacklist = d.Get("bot_enable_black_list").(string)
		hasChange = true
	}
	if d.HasChange("bot_enable_ip_reputation") {
		log.Printf("[DEBUG]  citrixadc-provider: Botenableipreputation has changed for botprofile %s, starting update", botprofileName)
		botprofile.Botenableipreputation = d.Get("bot_enable_ip_reputation").(string)
		hasChange = true
	}
	if d.HasChange("bot_enable_rate_limit") {
		log.Printf("[DEBUG]  citrixadc-provider: Botenableratelimit has changed for botprofile %s, starting update", botprofileName)
		botprofile.Botenableratelimit = d.Get("bot_enable_rate_limit").(string)
		hasChange = true
	}
	if d.HasChange("bot_enable_tps") {
		log.Printf("[DEBUG]  citrixadc-provider: Botenabletps has changed for botprofile %s, starting update", botprofileName)
		botprofile.Botenabletps = d.Get("bot_enable_tps").(string)
		hasChange = true
	}
	if d.HasChange("bot_enable_white_list") {
		log.Printf("[DEBUG]  citrixadc-provider: Botenablewhitelist has changed for botprofile %s, starting update", botprofileName)
		botprofile.Botenablewhitelist = d.Get("bot_enable_white_list").(string)
		hasChange = true
	}
	if d.HasChange("clientipexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientipexpression has changed for botprofile %s, starting update", botprofileName)
		botprofile.Clientipexpression = d.Get("clientipexpression").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for botprofile %s, starting update", botprofileName)
		botprofile.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("devicefingerprint") {
		log.Printf("[DEBUG]  citrixadc-provider: Devicefingerprint has changed for botprofile %s, starting update", botprofileName)
		botprofile.Devicefingerprint = d.Get("devicefingerprint").(string)
		hasChange = true
	}
	if d.HasChange("devicefingerprintaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Devicefingerprintaction has changed for botprofile %s, starting update", botprofileName)
		botprofile.Devicefingerprintaction = toStringList(d.Get("devicefingerprintaction").([]interface{}))
		hasChange = true
	}
	if d.HasChange("devicefingerprintmobile") {
		log.Printf("[DEBUG]  citrixadc-provider: Devicefingerprintmobile has changed for botprofile %s, starting update", botprofileName)
		botprofile.Devicefingerprintmobile = toStringList(d.Get("devicefingerprintmobile").([]interface{}))
		hasChange = true
	}
	if d.HasChange("errorurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Errorurl has changed for botprofile %s, starting update", botprofileName)
		botprofile.Errorurl = d.Get("errorurl").(string)
		hasChange = true
	}
	if d.HasChange("kmdetection") {
		log.Printf("[DEBUG]  citrixadc-provider: Kmdetection has changed for botprofile %s, starting update", botprofileName)
		botprofile.Kmdetection = d.Get("kmdetection").(string)
		hasChange = true
	}
	if d.HasChange("kmeventspostbodylimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Kmeventspostbodylimit has changed for botprofile %s, starting update", botprofileName)
		botprofile.Kmeventspostbodylimit = intPtr(d.Get("kmeventspostbodylimit").(int))
		hasChange = true
	}
	if d.HasChange("kmjavascriptname") {
		log.Printf("[DEBUG]  citrixadc-provider: Kmjavascriptname has changed for botprofile %s, starting update", botprofileName)
		botprofile.Kmjavascriptname = d.Get("kmjavascriptname").(string)
		hasChange = true
	}
	if d.HasChange("signature") {
		log.Printf("[DEBUG]  citrixadc-provider: Signature has changed for botprofile %s, starting update", botprofileName)
		botprofile.Signature = d.Get("signature").(string)
		hasChange = true
	}
	if d.HasChange("signaturemultipleuseragentheaderaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Signaturemultipleuseragentheaderaction has changed for botprofile %s, starting update", botprofileName)
		botprofile.Signaturemultipleuseragentheaderaction = toStringList(d.Get("signaturemultipleuseragentheaderaction").([]interface{}))
		hasChange = true
	}
	if d.HasChange("signaturenouseragentheaderaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Signaturenouseragentheaderaction has changed for botprofile %s, starting update", botprofileName)
		botprofile.Signaturenouseragentheaderaction = toStringList(d.Get("signaturenouseragentheaderaction").([]interface{}))
		hasChange = true
	}
	if d.HasChange("trap") {
		log.Printf("[DEBUG]  citrixadc-provider: Trap has changed for botprofile %s, starting update", botprofileName)
		botprofile.Trap = d.Get("trap").(string)
		hasChange = true
	}
	if d.HasChange("trapaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Trapaction has changed for botprofile %s, starting update", botprofileName)
		botprofile.Trapaction = toStringList(d.Get("trapaction").([]interface{}))
		hasChange = true
	}
	if d.HasChange("trapurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Trapurl has changed for botprofile %s, starting update", botprofileName)
		botprofile.Trapurl = d.Get("trapurl").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("botprofile", botprofileName, &botprofile)
		if err != nil {
			return diag.Errorf("Error updating botprofile %s", botprofileName)
		}
	}
	return readBotprofileFunc(ctx, d, meta)
}

func deleteBotprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	botprofileName := d.Id()
	err := client.DeleteResource("botprofile", botprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
