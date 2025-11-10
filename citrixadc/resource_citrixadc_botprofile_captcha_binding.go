package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/bot"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcBotprofile_captcha_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createBotprofile_captcha_bindingFunc,
		ReadContext:   readBotprofile_captcha_bindingFunc,
		DeleteContext: deleteBotprofile_captcha_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"bot_captcha_url": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"captcharesource": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_bind_comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_captcha_action": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_captcha_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"graceperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"logmessage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"muteperiod": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"requestsizelimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"retryattempts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"waittime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBotprofile_captcha_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotprofile_captcha_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	bot_captcha_url := d.Get("bot_captcha_url")
	bindingId := fmt.Sprintf("%s,%s", name, bot_captcha_url)
	botprofile_captcha_binding := bot.Botprofilecaptchabinding{
		Botbindcomment:    d.Get("bot_bind_comment").(string),
		Botcaptchaaction:  toStringList(d.Get("bot_captcha_action").([]interface{})),
		Botcaptchaenabled: d.Get("bot_captcha_enabled").(string),
		Botcaptchaurl:     d.Get("bot_captcha_url").(string),
		Captcharesource:   d.Get("captcharesource").(bool),
		Logmessage:        d.Get("logmessage").(string),
		Name:              d.Get("name").(string),
	}

	if raw := d.GetRawConfig().GetAttr("graceperiod"); !raw.IsNull() {
		botprofile_captcha_binding.Graceperiod = intPtr(d.Get("graceperiod").(int))
	}
	if raw := d.GetRawConfig().GetAttr("muteperiod"); !raw.IsNull() {
		botprofile_captcha_binding.Muteperiod = intPtr(d.Get("muteperiod").(int))
	}
	if raw := d.GetRawConfig().GetAttr("requestsizelimit"); !raw.IsNull() {
		botprofile_captcha_binding.Requestsizelimit = intPtr(d.Get("requestsizelimit").(int))
	}
	if raw := d.GetRawConfig().GetAttr("retryattempts"); !raw.IsNull() {
		botprofile_captcha_binding.Retryattempts = intPtr(d.Get("retryattempts").(int))
	}
	if raw := d.GetRawConfig().GetAttr("waittime"); !raw.IsNull() {
		botprofile_captcha_binding.Waittime = intPtr(d.Get("waittime").(int))
	}

	err := client.UpdateUnnamedResource("botprofile_captcha_binding", &botprofile_captcha_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readBotprofile_captcha_bindingFunc(ctx, d, meta)
}

func readBotprofile_captcha_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotprofile_captcha_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_captcha_url := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading botprofile_captcha_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "botprofile_captcha_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	log.Print("Helloooooo")
	log.Println(dataArr)
	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return diag.FromErr(err)
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_captcha_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["bot_captcha_url"].(string) == bot_captcha_url {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_captcha_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bot_bind_comment", data["bot_bind_comment"])
	d.Set("bot_captcha_action", data["bot_captcha_action"])
	d.Set("bot_captcha_enabled", data["bot_captcha_enabled"])
	d.Set("bot_captcha_url", data["bot_captcha_url"])
	d.Set("captcharesource", data["captcharesource"])
	setToInt("graceperiod", d, data["graceperiod"])
	d.Set("logmessage", data["logmessage"])
	setToInt("muteperiod", d, data["muteperiod"])
	d.Set("name", data["name"])
	setToInt("requestsizelimit", d, data["requestsizelimit"])
	setToInt("retryattempts", d, data["retryattempts"])
	setToInt("waittime", d, data["waittime"])

	return nil

}

func deleteBotprofile_captcha_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotprofile_captcha_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_captcha_url := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("bot_captcha_url:%s", url.QueryEscape(bot_captcha_url)))
	if val, ok := d.GetOk("captcharesource"); ok {
		args = append(args, fmt.Sprintf("captcharesource:%t", (val.(bool))))
	}

	err := client.DeleteResourceWithArgs("botprofile_captcha_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
