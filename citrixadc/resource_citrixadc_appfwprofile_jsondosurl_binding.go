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

func resourceCitrixAdcAppfwprofile_jsondosurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwprofile_jsondosurl_bindingFunc,
		ReadContext:   readAppfwprofile_jsondosurl_bindingFunc,
		DeleteContext: deleteAppfwprofile_jsondosurl_bindingFunc,
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
			"jsondosurl": {
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
			"jsonmaxarraylength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxarraylengthcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxcontainerdepth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxcontainerdepthcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxdocumentlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxdocumentlengthcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxobjectkeycount": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxobjectkeycountcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxobjectkeylength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxobjectkeylengthcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxstringlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxstringlengthcheck": {
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

func createAppfwprofile_jsondosurl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_jsondosurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	jsondosurl := d.Get("jsondosurl")
	bindingId := fmt.Sprintf("%s,%s", name, jsondosurl)
	appfwprofile_jsondosurl_binding := appfw.Appfwprofilejsondosurlbinding{
		Alertonly:                   d.Get("alertonly").(string),
		Comment:                     d.Get("comment").(string),
		Isautodeployed:              d.Get("isautodeployed").(string),
		Jsondosurl:                  d.Get("jsondosurl").(string),
		Jsonmaxarraylength:          d.Get("jsonmaxarraylength").(int),
		Jsonmaxarraylengthcheck:     d.Get("jsonmaxarraylengthcheck").(string),
		Jsonmaxcontainerdepth:       d.Get("jsonmaxcontainerdepth").(int),
		Jsonmaxcontainerdepthcheck:  d.Get("jsonmaxcontainerdepthcheck").(string),
		Jsonmaxdocumentlength:       d.Get("jsonmaxdocumentlength").(int),
		Jsonmaxdocumentlengthcheck:  d.Get("jsonmaxdocumentlengthcheck").(string),
		Jsonmaxobjectkeycount:       d.Get("jsonmaxobjectkeycount").(int),
		Jsonmaxobjectkeycountcheck:  d.Get("jsonmaxobjectkeycountcheck").(string),
		Jsonmaxobjectkeylength:      d.Get("jsonmaxobjectkeylength").(int),
		Jsonmaxobjectkeylengthcheck: d.Get("jsonmaxobjectkeylengthcheck").(string),
		Jsonmaxstringlength:         d.Get("jsonmaxstringlength").(int),
		Jsonmaxstringlengthcheck:    d.Get("jsonmaxstringlengthcheck").(string),
		Name:                        d.Get("name").(string),
		Resourceid:                  d.Get("resourceid").(string),
		Ruletype:                    d.Get("ruletype").(string),
		State:                       d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource("appfwprofile_jsondosurl_binding", &appfwprofile_jsondosurl_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAppfwprofile_jsondosurl_bindingFunc(ctx, d, meta)
}

func readAppfwprofile_jsondosurl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_jsondosurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	jsondosurl := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_jsondosurl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_jsondosurl_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_jsondosurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["jsondosurl"].(string) == jsondosurl {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_jsondosurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	// d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	// d.Set("isautodeployed", data["isautodeployed"])
	d.Set("jsondosurl", data["jsondosurl"])
	setToInt("jsonmaxarraylength", d, data["jsonmaxarraylength"])
	d.Set("jsonmaxarraylengthcheck", data["jsonmaxarraylengthcheck"])
	setToInt("jsonmaxcontainerdepth", d, data["jsonmaxcontainerdepth"])
	d.Set("jsonmaxcontainerdepthcheck", data["jsonmaxcontainerdepthcheck"])
	setToInt("jsonmaxdocumentlength", d, data["jsonmaxdocumentlength"])
	d.Set("jsonmaxdocumentlengthcheck", data["jsonmaxdocumentlengthcheck"])
	setToInt("jsonmaxobjectkeycount", d, data["jsonmaxobjectkeycount"])
	d.Set("jsonmaxobjectkeycountcheck", data["jsonmaxobjectkeycountcheck"])
	setToInt("jsonmaxobjectkeylength", d, data["jsonmaxobjectkeylength"])
	d.Set("jsonmaxobjectkeylengthcheck", data["jsonmaxobjectkeylengthcheck"])
	setToInt("jsonmaxstringlength", d, data["jsonmaxstringlength"])
	d.Set("jsonmaxstringlengthcheck", data["jsonmaxstringlengthcheck"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_jsondosurl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_jsondosurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	jsondosurl := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("jsondosurl:%s", url.QueryEscape(jsondosurl)))
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs("appfwprofile_jsondosurl_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
