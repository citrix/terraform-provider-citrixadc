package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcAppfwprofile_fileuploadtype_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwprofile_fileuploadtype_bindingFunc,
		Read:          readAppfwprofile_fileuploadtype_bindingFunc,
		Delete:        deleteAppfwprofile_fileuploadtype_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"fileuploadtype": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"as_fileuploadtypes_url": {
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
			"filetype": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"isautodeployed": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"isregexfileuploadtypesurl": {
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
			"isnameregex": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwprofile_fileuploadtype_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwprofile_fileuploadtype_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	fileuploadtype := d.Get("fileuploadtype")
	as_fileuploadtypes_url := d.Get("as_fileuploadtypes_url")
	filetype := strings.Join(toStringList(d.Get("filetype").([]interface{})), " ")
	bindingId := fmt.Sprintf("%s,%s,%s,%s", name, fileuploadtype, as_fileuploadtypes_url, url.QueryEscape(filetype))
	appfwprofile_fileuploadtype_binding := appfw.Appfwprofilefileuploadtypebinding{
		Alertonly:                 d.Get("alertonly").(string),
		Asfileuploadtypesurl:      d.Get("as_fileuploadtypes_url").(string),
		Comment:                   d.Get("comment").(string),
		Filetype:                  toStringList(d.Get("filetype").([]interface{})),
		Fileuploadtype:            d.Get("fileuploadtype").(string),
		Isautodeployed:            d.Get("isautodeployed").(string),
		Isregexfileuploadtypesurl: d.Get("isregexfileuploadtypesurl").(string),
		Isnameregex:               d.Get("isnameregex").(string),
		Name:                      d.Get("name").(string),
		Resourceid:                d.Get("resourceid").(string),
		Ruletype:                  d.Get("ruletype").(string),
		State:                     d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource("appfwprofile_fileuploadtype_binding", &appfwprofile_fileuploadtype_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readAppfwprofile_fileuploadtype_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwprofile_fileuploadtype_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readAppfwprofile_fileuploadtype_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwprofile_fileuploadtype_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	name := idSlice[0]
	fileuploadtype := idSlice[1]
	as_fileuploadtypes_url := idSlice[2]
	filetype := ""
	if len(idSlice) > 3 {
		filetype = idSlice[3]
	} else {
		filetype = url.QueryEscape(strings.Join(toStringList(d.Get("filetype").([]interface{})), " "))
		bindingId = fmt.Sprintf("%s,%s,%s,%s", name, fileuploadtype, as_fileuploadtypes_url, filetype)
	}

	log.Printf("[DEBUG] citrixadc-provider: Reading appfwprofile_fileuploadtype_binding state %s", bindingId)
	d.SetId(bindingId)

	findParams := service.FindParams{
		ResourceType:             "appfwprofile_fileuploadtype_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_fileuploadtype_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["fileuploadtype"].(string) == fileuploadtype {
			if v["as_fileuploadtypes_url"].(string) == as_fileuploadtypes_url {
				// Check if filetype matches (convert slice to space-separated string for comparison)
				dataFiletype := ""
				if v["filetype"] != nil {
					if filetypeSlice, ok := v["filetype"].([]interface{}); ok {
						dataFiletype = url.QueryEscape(strings.Join(toStringList(filetypeSlice), " "))
					}
				}
				if dataFiletype == filetype {
					foundIndex = i
					break
				}
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing appfwprofile_fileuploadtype_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("alertonly", data["alertonly"])
	d.Set("as_fileuploadtypes_url", data["as_fileuploadtypes_url"])
	d.Set("comment", data["comment"])
	d.Set("filetype", data["filetype"])
	d.Set("fileuploadtype", data["fileuploadtype"])
	d.Set("isautodeployed", data["isautodeployed"])
	d.Set("isregexfileuploadtypesurl", data["isregexfileuploadtypesurl"])
	d.Set("name", data["name"])
	d.Set("resourceid", data["resourceid"])
	d.Set("ruletype", data["ruletype"])
	d.Set("state", data["state"])
	d.Set("isnameregex", data["isnameregex"])
	return nil

}

func deleteAppfwprofile_fileuploadtype_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwprofile_fileuploadtype_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	name := idSlice[0]
	fileuploadtype := idSlice[1]
	as_fileuploadtypes_url := idSlice[2]
	filetype := ""
	if len(idSlice) > 3 {
		filetype = idSlice[3]
	} else {
		filetype = url.QueryEscape(strings.Join(toStringList(d.Get("filetype").([]interface{})), " "))
	}

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("fileuploadtype:%s", fileuploadtype))
	args = append(args, fmt.Sprintf("as_fileuploadtypes_url:%s", as_fileuploadtypes_url))
	if filetype != "" {
		args = append(args, fmt.Sprintf("filetype:%s", filetype))
	}
	if val, ok := d.GetOk("ruletype"); ok {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs("appfwprofile_fileuploadtype_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
