package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcContentinspectionglobal_contentinspectionpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createContentinspectionglobal_contentinspectionpolicy_bindingFunc,
		ReadContext:   readContentinspectionglobal_contentinspectionpolicy_bindingFunc,
		DeleteContext: deleteContentinspectionglobal_contentinspectionpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"policyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"invoke": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labeltype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"globalbindtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createContentinspectionglobal_contentinspectionpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectionglobal_contentinspectionpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	contentinspectionglobal_contentinspectionpolicy_binding := contentinspection.Contentinspectionglobalcontentinspectionpolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
		Type:                   d.Get("type").(string),
		Globalbindtype:         d.Get("globalbindtype").(string),
	}

	err := client.UpdateUnnamedResource("contentinspectionglobal_contentinspectionpolicy_binding", &contentinspectionglobal_contentinspectionpolicy_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policyname)

	return readContentinspectionglobal_contentinspectionpolicy_bindingFunc(ctx, d, meta)
}

func readContentinspectionglobal_contentinspectionpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectionglobal_contentinspectionpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectionglobal_contentinspectionpolicy_binding state %s", policyname)
	argsmap := make(map[string]string)

	if v, ok := d.GetOk("type"); ok {
		argsmap["type"] = v.(string)
	} else {
		argsmap["type"] = "REQ_DEFAULT"
	}
	findParams := service.FindParams{
		ResourceType:             "contentinspectionglobal_contentinspectionpolicy_binding",
		ArgsMap:                  argsmap,
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
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectionglobal_contentinspectionpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["policyname"].(string) == policyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams policyname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectionglobal_contentinspectionpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("globalbindtype", data["globalbindtype"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("invoke", data["invoke"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])
	d.Set("type", data["type"])

	return nil

}

func deleteContentinspectionglobal_contentinspectionpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectionglobal_contentinspectionpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))
	args = append(args, fmt.Sprintf("priority:%v", d.Get("priority").(int)))
	if v, ok := d.GetOk("type"); ok {
		type_val := v.(string)
		args = append(args, fmt.Sprintf("type:%s", type_val))
	} else {
		args = append(args, "type:REQ_DEFAULT")
	}

	err := client.DeleteResourceWithArgs("contentinspectionglobal_contentinspectionpolicy_binding", "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
