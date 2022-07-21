package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcGslbvserver_domain_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbvserver_domain_bindingFunc,
		Read:          readGslbvserver_domain_bindingFunc,
		Delete:        deleteGslbvserver_domain_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domainname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backupip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"backupipflag": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cookiedomain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cookiedomainflag": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cookietimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sitedomainttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createGslbvserver_domain_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbvserver_domain_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	domainname := d.Get("domainname")
	
	bindingId := fmt.Sprintf("%s,%s", name, domainname)
	gslbvserver_domain_binding := gslb.Gslbvserverdomainbinding{
		Backupip:         d.Get("backupip").(string),
		Backupipflag:     d.Get("backupipflag").(bool),
		Cookiedomain:     d.Get("cookiedomain").(string),
		Cookiedomainflag: d.Get("cookiedomainflag").(bool),
		Cookietimeout:    d.Get("cookietimeout").(int),
		Domainname:       d.Get("domainname").(string),
		Name:             d.Get("name").(string),
		Sitedomainttl:    d.Get("sitedomainttl").(int),
		Ttl:              d.Get("ttl").(int),
	}

	err := client.UpdateUnnamedResource(service.Gslbvserver_domain_binding.Type(), &gslbvserver_domain_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readGslbvserver_domain_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbvserver_domain_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readGslbvserver_domain_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbvserver_domain_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	domainname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading gslbvserver_domain_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "gslbvserver_domain_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing gslbvserver_domain_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["domainname"].(string) == domainname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams domainname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbvserver_domain_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("backupip", data["backupip"])
	d.Set("backupipflag", data["backupipflag"])
	d.Set("cookiedomain", data["cookiedomain"])
	d.Set("cookiedomainflag", data["cookiedomainflag"])
	d.Set("cookietimeout", data["cookietimeout"])
	d.Set("domainname", data["domainname"])
	d.Set("name", data["name"])
	d.Set("sitedomainttl", data["sitedomainttl"])
	d.Set("ttl", data["ttl"])

	return nil

}

func deleteGslbvserver_domain_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbvserver_domain_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	domainname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("domainname:%s", domainname))

	err := client.DeleteResourceWithArgs(service.Gslbvserver_domain_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
