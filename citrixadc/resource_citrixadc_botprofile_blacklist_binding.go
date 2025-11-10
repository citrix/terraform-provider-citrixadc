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

func resourceCitrixAdcBotprofile_blacklist_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createBotprofile_blacklist_bindingFunc,
		ReadContext:   readBotprofile_blacklist_bindingFunc,
		DeleteContext: deleteBotprofile_blacklist_bindingFunc,
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
			"bot_blacklist_value": {
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
			"bot_blacklist": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_blacklist_action": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_blacklist_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_blacklist_type": {
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

func createBotprofile_blacklist_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotprofile_blacklist_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	bot_blacklist_value := d.Get("bot_blacklist_value")
	bindingId := fmt.Sprintf("%s,%s", name, bot_blacklist_value)
	botprofile_blacklist_binding := bot.Botprofileblacklistbinding{
		Botbindcomment:      d.Get("bot_bind_comment").(string),
		Botblacklist:        d.Get("bot_blacklist").(bool),
		Botblacklistaction:  toStringList(d.Get("bot_blacklist_action").([]interface{})),
		Botblacklistenabled: d.Get("bot_blacklist_enabled").(string),
		Botblacklisttype:    d.Get("bot_blacklist_type").(string),
		Botblacklistvalue:   d.Get("bot_blacklist_value").(string),
		Logmessage:          d.Get("logmessage").(string),
		Name:                d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource("botprofile_blacklist_binding", &botprofile_blacklist_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readBotprofile_blacklist_bindingFunc(ctx, d, meta)
}

func readBotprofile_blacklist_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotprofile_blacklist_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_blacklist_value := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading botprofile_blacklist_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "botprofile_blacklist_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_blacklist_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["bot_blacklist_value"].(string) == bot_blacklist_value {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_blacklist_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bot_bind_comment", data["bot_bind_comment"])
	d.Set("bot_blacklist", data["bot_blacklist"])
	d.Set("bot_blacklist_action", data["bot_blacklist_action"])
	d.Set("bot_blacklist_enabled", data["bot_blacklist_enabled"])
	d.Set("bot_blacklist_type", data["bot_blacklist_type"])
	d.Set("bot_blacklist_value", data["bot_blacklist_value"])
	d.Set("logmessage", data["logmessage"])
	d.Set("name", data["name"])

	return nil

}

func deleteBotprofile_blacklist_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotprofile_blacklist_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bot_blacklist_value := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("bot_blacklist_value:%s", bot_blacklist_value))
	if val, ok := d.GetOk("bot_blacklist"); ok {
		args = append(args, fmt.Sprintf("bot_blacklist:%t", val.(bool)))
	}

	err := client.DeleteResourceWithArgs("botprofile_blacklist_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
