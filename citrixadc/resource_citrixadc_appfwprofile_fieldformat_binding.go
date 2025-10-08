package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAppfwprofile_fieldformat_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwprofile_fieldformat_bindingFunc,
		ReadContext:   readAppfwprofile_fieldformat_bindingFunc,
		DeleteContext: deleteAppfwprofile_fieldformat_bindingFunc,
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
			"fieldformat": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"formactionurl_ff": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"alertonly": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"fieldformatmaxlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"fieldformatminlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"fieldtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isautodeployed": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isregexff": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resourceid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ruletype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_fieldformat_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_fieldformat_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	fieldformat := d.Get("fieldformat")
	formactionurl_ff := d.Get("formactionurl_ff")
	bindingId := fmt.Sprintf("%s,%s,%s", name, fieldformat, formactionurl_ff)
	appfwprofile_fieldformat_binding := appfw.Appfwprofilefieldformatbinding{
		Alertonly:            d.Get("alertonly").(string),
		Comment:              d.Get("comment").(string),
		Fieldformat:          d.Get("fieldformat").(string),
		Fieldformatmaxlength: d.Get("fieldformatmaxlength").(int),
		Fieldformatminlength: d.Get("fieldformatminlength").(int),
		Fieldtype:            d.Get("fieldtype").(string),
		Formactionurlff:      d.Get("formactionurl_ff").(string),
		Isautodeployed:       d.Get("isautodeployed").(string),
		Isregexff:            d.Get("isregexff").(string),
		Name:                 d.Get("name").(string),
		Resourceid:           d.Get("resourceid").(string),
		Ruletype:             d.Get("ruletype").(string),
		State:                d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_fieldformat_binding.Type(), &appfwprofile_fieldformat_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAppfwprofile_fieldformat_bindingFunc(ctx, d, meta)
}

func readAppfwprofile_fieldformat_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_fieldformat_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	fieldformat := idSlice[1]
	formactionurl_ff := idSlice[2]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_fieldformat_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_fieldformat_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_fieldformat_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true
		if v["fieldformat"].(string) != fieldformat {
			match = false
		}
		if v["formactionurl_ff"].(string) != formactionurl_ff {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_fieldformat_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	d.Set("fieldformat", data["fieldformat"])
	setToInt("fieldformatmaxlength", d, data["fieldformatmaxlength"])
	setToInt("fieldformatminlength", d, data["fieldformatminlength"])
	d.Set("fieldtype", data["fieldtype"])
	d.Set("formactionurl_ff", data["formactionurl_ff"])
	d.Set("isautodeployed", data["isautodeployed"])
	// d.Set("isregexff", data["isregexff"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_fieldformat_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_fieldformat_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	fieldformat := idSlice[1]
	formactionurl_ff := idSlice[2]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("fieldformat:%s", fieldformat))
	args = append(args, fmt.Sprintf("formactionurl_ff:%s", url.QueryEscape(formactionurl_ff)))

	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Appfwprofile_fieldformat_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
