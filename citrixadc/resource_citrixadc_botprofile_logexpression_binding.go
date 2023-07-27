package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/bot"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcBotprofile_logexpression_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBotprofile_logexpression_bindingFunc,
		Read:          readBotprofile_logexpression_bindingFunc,
		Delete:        deleteBotprofile_logexpression_bindingFunc,
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
			"bot_log_expression_name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"bot_bind_comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_log_expression_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_log_expression_value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"logexpression": {
				Type:     schema.TypeBool,
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
		},
	}
}

func createBotprofile_logexpression_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotprofile_logexpression_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	bot_log_expression_name := d.Get("bot_log_expression_name")
	bindingId := fmt.Sprintf("%s,%s", name, bot_log_expression_name)
	botprofile_logexpression_binding := bot.Botprofilelogexpressionbinding{
		Botbindcomment:          d.Get("bot_bind_comment").(string),
		Botlogexpressionenabled: d.Get("bot_log_expression_enabled").(string),
		Botlogexpressionname:    d.Get("bot_log_expression_name").(string),
		Botlogexpressionvalue:   d.Get("bot_log_expression_value").(string),
		Logexpression:           d.Get("logexpression").(bool),
		Logmessage:              d.Get("logmessage").(string),
		Name:                    d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource("botprofile_logexpression_binding", &botprofile_logexpression_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readBotprofile_logexpression_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this botprofile_logexpression_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readBotprofile_logexpression_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotprofile_logexpression_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_log_expression_name := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading botprofile_logexpression_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "botprofile_logexpression_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_logexpression_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["bot_log_expression_name"].(string) == bot_log_expression_name {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_logexpression_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bot_bind_comment", data["bot_bind_comment"])
	d.Set("bot_log_expression_enabled", data["bot_log_expression_enabled"])
	d.Set("bot_log_expression_name", data["bot_log_expression_name"])
	d.Set("bot_log_expression_value", data["bot_log_expression_value"])
	d.Set("logexpression", data["logexpression"])
	d.Set("logmessage", data["logmessage"])
	d.Set("name", data["name"])

	return nil

}

func deleteBotprofile_logexpression_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotprofile_logexpression_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_log_expression_name := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("bot_log_expression_name:%s", bot_log_expression_name))
	if val, ok := d.GetOk("logexpression"); ok {
		args = append(args, fmt.Sprintf("logexpression:%t", val.(bool)))
	}

	err := client.DeleteResourceWithArgs("botprofile_logexpression_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
