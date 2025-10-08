package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcAppfwprofile_xmlattachmenturl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwprofile_xmlattachmenturl_bindingFunc,
		ReadContext:   readAppfwprofile_xmlattachmenturl_bindingFunc,
		DeleteContext: deleteAppfwprofile_xmlattachmenturl_bindingFunc,
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
			"xmlattachmenturl": {
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
			"isautodeployed": {
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
			"xmlattachmentcontenttype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlattachmentcontenttypecheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxattachmentsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxattachmentsizecheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_xmlattachmenturl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_xmlattachmenturl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	xmlattachmenturl := d.Get("xmlattachmenturl")
	bindingId := fmt.Sprintf("%s,%s", name, xmlattachmenturl)
	appfwprofile_xmlattachmenturl_binding := appfw.Appfwprofilexmlattachmenturlbinding{
		Alertonly:                     d.Get("alertonly").(string),
		Comment:                       d.Get("comment").(string),
		Isautodeployed:                d.Get("isautodeployed").(string),
		Name:                          d.Get("name").(string),
		Resourceid:                    d.Get("resourceid").(string),
		Ruletype:                      d.Get("ruletype").(string),
		State:                         d.Get("state").(string),
		Xmlattachmentcontenttype:      d.Get("xmlattachmentcontenttype").(string),
		Xmlattachmentcontenttypecheck: d.Get("xmlattachmentcontenttypecheck").(string),
		Xmlattachmenturl:              d.Get("xmlattachmenturl").(string),
		Xmlmaxattachmentsize:          d.Get("xmlmaxattachmentsize").(int),
		Xmlmaxattachmentsizecheck:     d.Get("xmlmaxattachmentsizecheck").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_xmlattachmenturl_binding.Type(), &appfwprofile_xmlattachmenturl_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAppfwprofile_xmlattachmenturl_bindingFunc(ctx, d, meta)
}

func readAppfwprofile_xmlattachmenturl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_xmlattachmenturl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	xmlattachmenturl := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_xmlattachmenturl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_xmlattachmenturl_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_xmlattachmenturl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["xmlattachmenturl"].(string) == xmlattachmenturl {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_xmlattachmenturl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	// d.Set("isautodeployed", data["isautodeployed"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])
	d.Set("xmlattachmentcontenttype", data["xmlattachmentcontenttype"])
	d.Set("xmlattachmentcontenttypecheck", data["xmlattachmentcontenttypecheck"])
	d.Set("xmlattachmenturl", data["xmlattachmenturl"])
	setToInt("xmlmaxattachmentsize", d, data["xmlmaxattachmentsize"])
	d.Set("xmlmaxattachmentsizecheck", data["xmlmaxattachmentsizecheck"])

	return nil

}

func deleteAppfwprofile_xmlattachmenturl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_xmlattachmenturl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	xmlattachmenturl := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("xmlattachmenturl:%s", url.QueryEscape(xmlattachmenturl)))
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Appfwprofile_xmlattachmenturl_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
