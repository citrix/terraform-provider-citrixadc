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

func resourceCitrixAdcBotprofile_tps_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createBotprofile_tps_bindingFunc,
		ReadContext:   readBotprofile_tps_bindingFunc,
		DeleteContext: deleteBotprofile_tps_bindingFunc,
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
			"bot_tps_type": {
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
			"bot_tps": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_tps_action": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"percentage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createBotprofile_tps_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readBotprofile_tps_bindingFunc(ctx, d, meta)
}

func readBotprofile_tps_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
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
	setToInt("percentage", d, data["percentage"])
	setToInt("threshold", d, data["threshold"])

	return nil

}

func deleteBotprofile_tps_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
