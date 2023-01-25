package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/bot"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
	"net/url"
)

func resourceCitrixAdcBotprofile_ratelimit_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBotprofile_ratelimit_bindingFunc,
		Read:          readBotprofile_ratelimit_bindingFunc,
		Delete:        deleteBotprofile_ratelimit_bindingFunc,
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
			"bot_rate_limit_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"bot_bind_comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_ratelimit": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_rate_limit_action": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_rate_limit_enabled": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_rate_limit_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cookiename": &schema.Schema{
				Type:     schema.TypeString,
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
			"rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"timeslice": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBotprofile_ratelimit_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotprofile_ratelimit_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	bot_rate_limit_type := d.Get("bot_rate_limit_type")
	bindingId := fmt.Sprintf("%s,%s", name, bot_rate_limit_type)
	botprofile_ratelimit_binding := bot.Botprofileratelimitbinding{
		Botbindcomment:      d.Get("bot_bind_comment").(string),
		Botratelimit:        d.Get("bot_ratelimit").(bool),
		Botratelimitaction:  toStringList(d.Get("bot_rate_limit_action").([]interface{})),
		Botratelimitenabled: d.Get("bot_rate_limit_enabled").(string),
		Botratelimittype:    d.Get("bot_rate_limit_type").(string),
		Botratelimiturl:     d.Get("bot_rate_limit_url").(string),
		Cookiename:          d.Get("cookiename").(string),
		Logmessage:          d.Get("logmessage").(string),
		Name:                d.Get("name").(string),
		Rate:                d.Get("rate").(int),
		Timeslice:           d.Get("timeslice").(int),
	}

	err := client.UpdateUnnamedResource("botprofile_ratelimit_binding", &botprofile_ratelimit_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readBotprofile_ratelimit_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this botprofile_ratelimit_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readBotprofile_ratelimit_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotprofile_ratelimit_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_rate_limit_type := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading botprofile_ratelimit_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "botprofile_ratelimit_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_ratelimit_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["bot_rate_limit_type"].(string) == bot_rate_limit_type {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_ratelimit_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bot_bind_comment", data["bot_bind_comment"])
	d.Set("bot_ratelimit", data["bot_ratelimit"])
	d.Set("bot_rate_limit_action", data["bot_rate_limit_action"])
	d.Set("bot_rate_limit_enabled", data["bot_rate_limit_enabled"])
	d.Set("bot_rate_limit_type", data["bot_rate_limit_type"])
	d.Set("bot_rate_limit_url", data["bot_rate_limit_url"])
	d.Set("cookiename", data["cookiename"])
	d.Set("logmessage", data["logmessage"])
	d.Set("name", data["name"])
	d.Set("rate", data["rate"])
	d.Set("timeslice", data["timeslice"])

	return nil

}

func deleteBotprofile_ratelimit_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotprofile_ratelimit_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_rate_limit_type := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("bot_rate_limit_type:%s", bot_rate_limit_type))
	if val, ok := d.GetOk("bot_ratelimit"); ok {
		args = append(args, fmt.Sprintf("bot_ratelimit:%t", val.(bool)))
	}
	if val, ok := d.GetOk("bot_rate_limit_url"); ok {
		args = append(args, fmt.Sprintf("bot_rate_limit_url:%s", url.QueryEscape(val.(string))))
	}
	if val, ok := d.GetOk("cookiename"); ok {
		args = append(args, fmt.Sprintf("cookiename:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs("botprofile_ratelimit_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
