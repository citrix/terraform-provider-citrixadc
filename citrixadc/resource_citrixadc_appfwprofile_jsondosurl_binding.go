package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
	"net/url"
)

func resourceCitrixAdcAppfwprofile_jsondosurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_jsondosurl_bindingFunc,
		Read:          readAppfwprofile_jsondosurl_bindingFunc,
		Delete:        deleteAppfwprofile_jsondosurl_bindingFunc,
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
			"jsondosurl": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"alertonly": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isautodeployed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxarraylength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxarraylengthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxcontainerdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxcontainerdepthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxdocumentlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxdocumentlengthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxobjectkeycount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxobjectkeycountcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxobjectkeylength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxobjectkeylengthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxstringlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"jsonmaxstringlengthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resourceid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ruletype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_jsondosurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_jsondosurl_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_jsondosurl_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_jsondosurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
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
	d.Set("jsonmaxarraylength", data["jsonmaxarraylength"])
	d.Set("jsonmaxarraylengthcheck", data["jsonmaxarraylengthcheck"])
	d.Set("jsonmaxcontainerdepth", data["jsonmaxcontainerdepth"])
	d.Set("jsonmaxcontainerdepthcheck", data["jsonmaxcontainerdepthcheck"])
	d.Set("jsonmaxdocumentlength", data["jsonmaxdocumentlength"])
	d.Set("jsonmaxdocumentlengthcheck", data["jsonmaxdocumentlengthcheck"])
	d.Set("jsonmaxobjectkeycount", data["jsonmaxobjectkeycount"])
	d.Set("jsonmaxobjectkeycountcheck", data["jsonmaxobjectkeycountcheck"])
	d.Set("jsonmaxobjectkeylength", data["jsonmaxobjectkeylength"])
	d.Set("jsonmaxobjectkeylengthcheck", data["jsonmaxobjectkeylengthcheck"])
	d.Set("jsonmaxstringlength", data["jsonmaxstringlength"])
	d.Set("jsonmaxstringlengthcheck", data["jsonmaxstringlengthcheck"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])

	return nil

}

func deleteAppfwprofile_jsondosurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId("")

	return nil
}
