package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/bot"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcBotprofile_trapinsertionurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBotprofile_trapinsertionurl_bindingFunc,
		Read:          readBotprofile_trapinsertionurl_bindingFunc,
		Delete:        deleteBotprofile_trapinsertionurl_bindingFunc,
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
			"bot_trap_url": &schema.Schema{
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
			"bot_trap_url_insertion_enabled": &schema.Schema{
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
			"trapinsertionurl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBotprofile_trapinsertionurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotprofile_trapinsertionurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	bot_trap_url := d.Get("bot_trap_url")
	bindingId := fmt.Sprintf("%s,%s", name, bot_trap_url)
	botprofile_trapinsertionurl_binding := bot.Botprofiletrapinsertionurlbinding{
		Botbindcomment:             d.Get("bot_bind_comment").(string),
		Bottrapurl:                 d.Get("bot_trap_url").(string),
		Bottrapurlinsertionenabled: d.Get("bot_trap_url_insertion_enabled").(string),
		Logmessage:                 d.Get("logmessage").(string),
		Name:                       d.Get("name").(string),
		Trapinsertionurl:           d.Get("trapinsertionurl").(bool),
	}

	err := client.UpdateUnnamedResource("botprofile_trapinsertionurl_binding", &botprofile_trapinsertionurl_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readBotprofile_trapinsertionurl_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this botprofile_trapinsertionurl_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readBotprofile_trapinsertionurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotprofile_trapinsertionurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_trap_url := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading botprofile_trapinsertionurl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "botprofile_trapinsertionurl_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_trapinsertionurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["bot_trap_url"].(string) == bot_trap_url {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_trapinsertionurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bot_bind_comment", data["bot_bind_comment"])
	d.Set("bot_trap_url", data["bot_trap_url"])
	d.Set("bot_trap_url_insertion_enabled", data["bot_trap_url_insertion_enabled"])
	d.Set("logmessage", data["logmessage"])
	d.Set("name", data["name"])
	d.Set("trapinsertionurl", data["trapinsertionurl"])

	return nil

}

func deleteBotprofile_trapinsertionurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotprofile_trapinsertionurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_trap_url := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("bot_trap_url:%s", bot_trap_url))
	if val, ok := d.GetOk("trapinsertionurl"); ok {
		args = append(args, fmt.Sprintf("trapinsertionurl:%t", val.(bool)))
	}

	err := client.DeleteResourceWithArgs("botprofile_trapinsertionurl_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
