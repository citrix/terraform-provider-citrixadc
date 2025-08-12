package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

/**
* Binding class showing the servicegroup that can be bound to lbvserver.
 */
type Lbvserverservicegroupbinding struct {
	/**
	* The service group name bound to the selected load balancing virtual server.
	 */
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* Order number to be assigned to the service when it is bound to the lb vserver.
	 */
	Order int `json:"order,omitempty"`
	/**
	* Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
	*/
	Name string `json:"name,omitempty"`
}

func resourceCitrixAdcLbvserver_servicegroup_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbvserver_servicegroup_bindingFunc,
		Read:          readLbvserver_servicegroup_bindingFunc,
		Delete:        deleteLbvserver_servicegroup_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"servicegroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createLbvserver_servicegroup_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbvserver_servicegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	name := d.Get("name").(string)
	servicegroupname := d.Get("servicegroupname").(string)

	bindingId := fmt.Sprintf("%s,%s", name, servicegroupname)

	lbvserver_servicegroup_binding := Lbvserverservicegroupbinding{
		Name:             d.Get("name").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
		Order:            d.Get("order").(int),
	}

	_, err := client.AddResource(service.Lbvserver_servicegroup_binding.Type(), name, &lbvserver_servicegroup_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLbvserver_servicegroup_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbvserver_servicegroup_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLbvserver_servicegroup_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbvserver_servicegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()

	idSlice := strings.SplitN(bindingId, ",", 2)
	name := idSlice[0]
	servicegroupname := idSlice[1]

	findParams := service.FindParams{
		ResourceType:             "lbvserver_servicegroup_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_servicegroup_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right monitor name
	foundIndex := -1
	for i, v := range dataArr {
		if v["servicegroupname"].(string) == servicegroupname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_servicegroup_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	log.Printf("[DEBUG] citrixadc-provider: Reading lbvserver_servicegroup_binding state %s", bindingId)
	d.Set("name", data["name"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("order", data["order"])

	return nil

}

func deleteLbvserver_servicegroup_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbvserver_servicegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()

	idSlice := strings.SplitN(bindingId, ",", 2)
	name := idSlice[0]
	servicegroupname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("servicegroupname:%s", servicegroupname))

	err := client.DeleteResourceWithArgs("lbvserver_servicegroup_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
