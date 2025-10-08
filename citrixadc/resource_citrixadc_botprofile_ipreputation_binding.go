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

func resourceCitrixAdcBotprofile_ipreputation_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createBotprofile_ipreputation_bindingFunc,
		ReadContext:   readBotprofile_ipreputation_bindingFunc,
		DeleteContext: deleteBotprofile_ipreputation_bindingFunc,
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
			"category": {
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
			"bot_iprep_action": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_iprep_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bot_ipreputation": {
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

func createBotprofile_ipreputation_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createBotprofile_ipreputation_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	category := d.Get("category")
	bindingId := fmt.Sprintf("%s,%s", name, category)
	botprofile_ipreputation_binding := bot.Botprofileipreputationbinding{
		Botbindcomment:  d.Get("bot_bind_comment").(string),
		Botiprepaction:  toStringList(d.Get("bot_iprep_action").([]interface{})),
		Botiprepenabled: d.Get("bot_iprep_enabled").(string),
		Botipreputation: d.Get("bot_ipreputation").(bool),
		Category:        d.Get("category").(string),
		Logmessage:      d.Get("logmessage").(string),
		Name:            d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource("botprofile_ipreputation_binding", &botprofile_ipreputation_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readBotprofile_ipreputation_bindingFunc(ctx, d, meta)
}

func readBotprofile_ipreputation_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readBotprofile_ipreputation_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	category := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading botprofile_ipreputation_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "botprofile_ipreputation_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_ipreputation_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["category"].(string) == category {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing botprofile_ipreputation_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bot_bind_comment", data["bot_bind_comment"])
	d.Set("bot_iprep_action", data["bot_iprep_action"])
	d.Set("bot_iprep_enabled", data["bot_iprep_enabled"])
	d.Set("bot_ipreputation", data["bot_ipreputation"])
	d.Set("category", data["category"])
	d.Set("logmessage", data["logmessage"])
	d.Set("name", data["name"])

	return nil

}

func deleteBotprofile_ipreputation_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteBotprofile_ipreputation_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	category := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("category:%s", category))
	if val, ok := d.GetOk("bot_ipreputation"); ok {
		args = append(args, fmt.Sprintf("bot_ipreputation:%t", val.(bool)))
	}

	err := client.DeleteResourceWithArgs("botprofile_ipreputation_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
