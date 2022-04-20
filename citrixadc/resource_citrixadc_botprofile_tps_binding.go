package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/bot"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcBotprofile_tps_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBotprofile_tps_bindingFunc,
		Read:          readBotprofile_tps_bindingFunc,
		Delete:        deleteBotprofile_tps_bindingFunc,
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
			"bot_tps_type": &schema.Schema{
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
			"bot_tps": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_tps_action": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"percentage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBotprofile_tps_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotprofile_tps_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	bot_tps_type := d.Get("bot_tps_type")
	bindingId := fmt.Sprintf("%s,%s", name, bot_tps_type)
	botprofile_tps_binding := bot.Botprofiletpsbinding{
		Botbindcomment: d.Get("bot_bind_comment").(string),
		Bottps:         d.Get("bot_tps").(bool),
		Bottpsaction:   toStringList(d.Get("bot_tps_action").([]interface{})),
		Bottpstype:     d.Get("bot_tps_type").(string),
		Logmessage:     d.Get("logmessage").(string),
		Name:           d.Get("name").(string),
		Percentage:     d.Get("percentage").(int),
		Threshold:      d.Get("threshold").(int),
	}

	err := client.UpdateUnnamedResource("botprofile_tps_binding", &botprofile_tps_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readBotprofile_tps_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this botprofile_tps_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readBotprofile_tps_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotprofile_tps_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_tps_type := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading botprofile_tps_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "botprofile_tps_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_tps_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["bot_tps_type"].(string) == bot_tps_type {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_tps_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bot_bind_comment", data["bot_bind_comment"])
	d.Set("bot_tps", data["bot_tps"])
	d.Set("bot_tps_action", data["bot_tps_action"])
	d.Set("bot_tps_type", data["bot_tps_type"])
	d.Set("logmessage", data["logmessage"])
	d.Set("name", data["name"])
	d.Set("percentage", data["percentage"])
	d.Set("threshold", data["threshold"])

	return nil

}

func deleteBotprofile_tps_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotprofile_tps_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_tps_type := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("bot_tps_type:%s", bot_tps_type))
	if val, ok := d.GetOk("bot_tps"); ok {
		args = append(args, fmt.Sprintf("bot_tps:%t", (val.(bool))))
	}

	err := client.DeleteResourceWithArgs("botprofile_tps_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
