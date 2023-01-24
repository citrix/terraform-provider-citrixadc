package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
	"net/url"
)

func resourceCitrixAdcAppfwprofile_xmldosurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_xmldosurl_bindingFunc,
		Read:          readAppfwprofile_xmldosurl_bindingFunc,
		Delete:        deleteAppfwprofile_xmldosurl_bindingFunc,
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
			"xmldosurl": &schema.Schema{
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
			"xmlblockdtd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlblockexternalentities": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlblockpi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxattributenamelength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxattributenamelengthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxattributes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxattributescheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxattributevaluelength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxattributevaluelengthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxchardatalength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxchardatalengthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxelementchildren": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxelementchildrencheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxelementdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxelementdepthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxelementnamelength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxelementnamelengthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxelements": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxelementscheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxentityexpansiondepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxentityexpansiondepthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxentityexpansions": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxentityexpansionscheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxfilesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxfilesizecheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxnamespaces": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxnamespacescheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxnamespaceurilength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxnamespaceurilengthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxnodes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxnodescheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxsoaparrayrank": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlmaxsoaparraysize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlminfilesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlminfilesizecheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xmlsoaparraycheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_xmldosurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_xmldosurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	xmldosurl := d.Get("xmldosurl")
	bindingId := fmt.Sprintf("%s,%s", name, xmldosurl)
	appfwprofile_xmldosurl_binding := appfw.Appfwprofilexmldosurlbinding{
		Alertonly:                       d.Get("alertonly").(string),
		Comment:                         d.Get("comment").(string),
		Isautodeployed:                  d.Get("isautodeployed").(string),
		Name:                            d.Get("name").(string),
		Resourceid:                      d.Get("resourceid").(string),
		Ruletype:                        d.Get("ruletype").(string),
		State:                           d.Get("state").(string),
		Xmlblockdtd:                     d.Get("xmlblockdtd").(string),
		Xmlblockexternalentities:        d.Get("xmlblockexternalentities").(string),
		Xmlblockpi:                      d.Get("xmlblockpi").(string),
		Xmldosurl:                       d.Get("xmldosurl").(string),
		Xmlmaxattributenamelength:       d.Get("xmlmaxattributenamelength").(int),
		Xmlmaxattributenamelengthcheck:  d.Get("xmlmaxattributenamelengthcheck").(string),
		Xmlmaxattributes:                d.Get("xmlmaxattributes").(int),
		Xmlmaxattributescheck:           d.Get("xmlmaxattributescheck").(string),
		Xmlmaxattributevaluelength:      d.Get("xmlmaxattributevaluelength").(int),
		Xmlmaxattributevaluelengthcheck: d.Get("xmlmaxattributevaluelengthcheck").(string),
		Xmlmaxchardatalength:            d.Get("xmlmaxchardatalength").(int),
		Xmlmaxchardatalengthcheck:       d.Get("xmlmaxchardatalengthcheck").(string),
		Xmlmaxelementchildren:           d.Get("xmlmaxelementchildren").(int),
		Xmlmaxelementchildrencheck:      d.Get("xmlmaxelementchildrencheck").(string),
		Xmlmaxelementdepth:              d.Get("xmlmaxelementdepth").(int),
		Xmlmaxelementdepthcheck:         d.Get("xmlmaxelementdepthcheck").(string),
		Xmlmaxelementnamelength:         d.Get("xmlmaxelementnamelength").(int),
		Xmlmaxelementnamelengthcheck:    d.Get("xmlmaxelementnamelengthcheck").(string),
		Xmlmaxelements:                  d.Get("xmlmaxelements").(int),
		Xmlmaxelementscheck:             d.Get("xmlmaxelementscheck").(string),
		Xmlmaxentityexpansiondepth:      d.Get("xmlmaxentityexpansiondepth").(int),
		Xmlmaxentityexpansiondepthcheck: d.Get("xmlmaxentityexpansiondepthcheck").(string),
		Xmlmaxentityexpansions:          d.Get("xmlmaxentityexpansions").(int),
		Xmlmaxentityexpansionscheck:     d.Get("xmlmaxentityexpansionscheck").(string),
		Xmlmaxfilesize:                  d.Get("xmlmaxfilesize").(int),
		Xmlmaxfilesizecheck:             d.Get("xmlmaxfilesizecheck").(string),
		Xmlmaxnamespaces:                d.Get("xmlmaxnamespaces").(int),
		Xmlmaxnamespacescheck:           d.Get("xmlmaxnamespacescheck").(string),
		Xmlmaxnamespaceurilength:        d.Get("xmlmaxnamespaceurilength").(int),
		Xmlmaxnamespaceurilengthcheck:   d.Get("xmlmaxnamespaceurilengthcheck").(string),
		Xmlmaxnodes:                     d.Get("xmlmaxnodes").(int),
		Xmlmaxnodescheck:                d.Get("xmlmaxnodescheck").(string),
		Xmlmaxsoaparrayrank:             d.Get("xmlmaxsoaparrayrank").(int),
		Xmlmaxsoaparraysize:             d.Get("xmlmaxsoaparraysize").(int),
		Xmlminfilesize:                  d.Get("xmlminfilesize").(int),
		Xmlminfilesizecheck:             d.Get("xmlminfilesizecheck").(string),
		Xmlsoaparraycheck:               d.Get("xmlsoaparraycheck").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwprofile_xmldosurl_binding.Type(), &appfwprofile_xmldosurl_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_xmldosurl_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_xmldosurl_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_xmldosurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_xmldosurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	xmldosurl := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_xmldosurl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_xmldosurl_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_xmldosurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["xmldosurl"].(string) == xmldosurl {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_xmldosurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("alertonly", data["alertonly"])
	d.Set("comment", data["comment"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])
	d.Set("xmlblockdtd", data["xmlblockdtd"])
	d.Set("xmlblockexternalentities", data["xmlblockexternalentities"])
	d.Set("xmlblockpi", data["xmlblockpi"])
	d.Set("xmldosurl", data["xmldosurl"])
	d.Set("xmlmaxattributenamelength", data["xmlmaxattributenamelength"])
	d.Set("xmlmaxattributenamelengthcheck", data["xmlmaxattributenamelengthcheck"])
	d.Set("xmlmaxattributes", data["xmlmaxattributes"])
	d.Set("xmlmaxattributescheck", data["xmlmaxattributescheck"])
	d.Set("xmlmaxattributevaluelength", data["xmlmaxattributevaluelength"])
	d.Set("xmlmaxattributevaluelengthcheck", data["xmlmaxattributevaluelengthcheck"])
	d.Set("xmlmaxchardatalength", data["xmlmaxchardatalength"])
	d.Set("xmlmaxchardatalengthcheck", data["xmlmaxchardatalengthcheck"])
	d.Set("xmlmaxelementchildren", data["xmlmaxelementchildren"])
	d.Set("xmlmaxelementchildrencheck", data["xmlmaxelementchildrencheck"])
	d.Set("xmlmaxelementdepth", data["xmlmaxelementdepth"])
	d.Set("xmlmaxelementdepthcheck", data["xmlmaxelementdepthcheck"])
	d.Set("xmlmaxelementnamelength", data["xmlmaxelementnamelength"])
	d.Set("xmlmaxelementnamelengthcheck", data["xmlmaxelementnamelengthcheck"])
	d.Set("xmlmaxelements", data["xmlmaxelements"])
	d.Set("xmlmaxelementscheck", data["xmlmaxelementscheck"])
	d.Set("xmlmaxentityexpansiondepth", data["xmlmaxentityexpansiondepth"])
	d.Set("xmlmaxentityexpansiondepthcheck", data["xmlmaxentityexpansiondepthcheck"])
	d.Set("xmlmaxentityexpansions", data["xmlmaxentityexpansions"])
	d.Set("xmlmaxentityexpansionscheck", data["xmlmaxentityexpansionscheck"])
	d.Set("xmlmaxfilesize", data["xmlmaxfilesize"])
	d.Set("xmlmaxfilesizecheck", data["xmlmaxfilesizecheck"])
	d.Set("xmlmaxnamespaces", data["xmlmaxnamespaces"])
	d.Set("xmlmaxnamespacescheck", data["xmlmaxnamespacescheck"])
	d.Set("xmlmaxnamespaceurilength", data["xmlmaxnamespaceurilength"])
	d.Set("xmlmaxnamespaceurilengthcheck", data["xmlmaxnamespaceurilengthcheck"])
	d.Set("xmlmaxnodes", data["xmlmaxnodes"])
	d.Set("xmlmaxnodescheck", data["xmlmaxnodescheck"])
	d.Set("xmlmaxsoaparrayrank", data["xmlmaxsoaparrayrank"])
	d.Set("xmlmaxsoaparraysize", data["xmlmaxsoaparraysize"])
	d.Set("xmlminfilesize", data["xmlminfilesize"])
	d.Set("xmlminfilesizecheck", data["xmlminfilesizecheck"])
	d.Set("xmlsoaparraycheck", data["xmlsoaparraycheck"])

	return nil

}

func deleteAppfwprofile_xmldosurl_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_xmldosurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	xmldosurl := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("xmldosurl:%s", url.QueryEscape(xmldosurl)))
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Appfwprofile_xmldosurl_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
