package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/bot"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcBotprofile_whitelist_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createBotprofile_whitelist_bindingFunc,
		ReadContext:   readBotprofile_whitelist_bindingFunc,
		DeleteContext: deleteBotprofile_whitelist_bindingFunc,
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
			"bot_whitelist_value": {
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
			"bot_whitelist": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_whitelist_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_whitelist_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"log": {
				Type:     schema.TypeString,
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

func createBotprofile_whitelist_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotprofile_whitelist_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	bot_whitelist_value := d.Get("bot_whitelist_value")
	bindingId := fmt.Sprintf("%s,%s", name, bot_whitelist_value)
	botprofile_whitelist_binding := bot.Botprofilewhitelistbinding{
		Botbindcomment:      d.Get("bot_bind_comment").(string),
		Botwhitelist:        d.Get("bot_whitelist").(bool),
		Botwhitelistenabled: d.Get("bot_whitelist_enabled").(string),
		Botwhitelisttype:    d.Get("bot_whitelist_type").(string),
		Botwhitelistvalue:   d.Get("bot_whitelist_value").(string),
		Log:                 d.Get("log").(string),
		Logmessage:          d.Get("logmessage").(string),
		Name:                d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource("botprofile_whitelist_binding", &botprofile_whitelist_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readBotprofile_whitelist_bindingFunc(ctx, d, meta)
}

func readBotprofile_whitelist_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotprofile_whitelist_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_whitelist_value := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading botprofile_whitelist_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "botprofile_whitelist_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return diag.FromErr(err)
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_whitelist_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["bot_whitelist_value"].(string) == bot_whitelist_value {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_whitelist_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bot_bind_comment", data["bot_bind_comment"])
	d.Set("bot_whitelist", data["bot_whitelist"])
	d.Set("bot_whitelist_enabled", data["bot_whitelist_enabled"])
	d.Set("bot_whitelist_type", data["bot_whitelist_type"])
	d.Set("bot_whitelist_value", data["bot_whitelist_value"])
	d.Set("log", data["log"])
	d.Set("logmessage", data["logmessage"])
	d.Set("name", data["name"])

	return nil

}

func deleteBotprofile_whitelist_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotprofile_whitelist_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_whitelist_value := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("bot_whitelist_value:%s", bot_whitelist_value))
	if val, ok := d.GetOk("bot_whitelist"); ok {
		args = append(args, fmt.Sprintf("bot_whitelist:%t", val.(bool)))
	}

	err := client.DeleteResourceWithArgs("botprofile_whitelist_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
