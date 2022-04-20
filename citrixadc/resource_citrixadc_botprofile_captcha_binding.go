package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/bot"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
	"net/url"
)

func resourceCitrixAdcBotprofile_captcha_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBotprofile_captcha_bindingFunc,
		Read:          readBotprofile_captcha_bindingFunc,
		Delete:        deleteBotprofile_captcha_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"bot_captcha_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"captcharesource": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_bind_comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_captcha_action": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_captcha_enabled": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"graceperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"logmessage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"muteperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"requestsizelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"retryattempts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"waittime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBotprofile_captcha_bindingFunc(d *schema.ResourceData, meta interface{}) error {
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
		Graceperiod:       d.Get("graceperiod").(int),
		Logmessage:        d.Get("logmessage").(string),
		Muteperiod:        d.Get("muteperiod").(int),
		Name:              d.Get("name").(string),
		Requestsizelimit:  d.Get("requestsizelimit").(int),
		Retryattempts:     d.Get("retryattempts").(int),
		Waittime:          d.Get("waittime").(int),
	}

	err := client.UpdateUnnamedResource("botprofile_captcha_binding", &botprofile_captcha_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readBotprofile_captcha_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this botprofile_captcha_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readBotprofile_captcha_bindingFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
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
	d.Set("graceperiod", data["graceperiod"])
	d.Set("logmessage", data["logmessage"])
	d.Set("muteperiod", data["muteperiod"])
	d.Set("name", data["name"])
	d.Set("requestsizelimit", data["requestsizelimit"])
	d.Set("retryattempts", data["retryattempts"])
	d.Set("waittime", data["waittime"])

	return nil

}

func deleteBotprofile_captcha_bindingFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId("")

	return nil
}
