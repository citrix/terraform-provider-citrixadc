package citrixadc

import (
	"strconv"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcGslbservicegroup_gslbservicegroupmember_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbservicegroup_gslbservicegroupmember_bindingFunc,
		Read:          readGslbservicegroup_gslbservicegroupmember_bindingFunc,
		Delete:        deleteGslbservicegroup_gslbservicegroupmember_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"hashid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"publicip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publicport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicegroupname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"siteprefix": {
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
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createGslbservicegroup_gslbservicegroupmember_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbservicegroup_gslbservicegroupmember_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingIdSlice := make([]string, 0, 3)

	bindingIdSlice = append(bindingIdSlice, d.Get("servicegroupname").(string))

	// Second id component will either be servername or ip
	// ConflictsWith restriction ensures that only one of them is defined
	if v, ok := d.GetOk("servername"); ok {
		bindingIdSlice = append(bindingIdSlice, v.(string))
	} else if v, ok := d.GetOk("ip"); ok {
		bindingIdSlice = append(bindingIdSlice, v.(string))
	}

	// Third component will be the port if defined
	if v, ok := d.GetOk("port"); ok {
		val := strconv.Itoa(v.(int))
		bindingIdSlice = append(bindingIdSlice, val)
	}

	bindingId := strings.Join(bindingIdSlice, ",")

	gslbservicegroup_gslbservicegroupmember_binding := gslb.Gslbservicegroupgslbservicegroupmemberbinding{
		Hashid:           d.Get("hashid").(int),
		Ip:               d.Get("ip").(string),
		Port:             d.Get("port").(int),
		Publicip:         d.Get("publicip").(string),
		Publicport:       d.Get("publicport").(int),
		Servername:       d.Get("servername").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
		Siteprefix:       d.Get("siteprefix").(string),
		State:            d.Get("state").(string),
		Weight:           d.Get("weight").(int),
		Order:            d.Get("order").(int),
	}

	_, err := client.AddResource("gslbservicegroup_gslbservicegroupmember_binding", bindingId, &gslbservicegroup_gslbservicegroupmember_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readGslbservicegroup_gslbservicegroupmember_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbservicegroup_gslbservicegroupmember_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readGslbservicegroup_gslbservicegroupmember_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbservicegroup_gslbservicegroupmember_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	servicegroupname := idSlice[0]

	// When ip is defined ADC will create a server with name equal to the ip address
	// So no matter what the user actually defined servername will always be valid search criterion
	servername := idSlice[1]

	port := 0
	var err error
	if port, err = strconv.Atoi(idSlice[2]); err != nil {
		return err
	}

	log.Printf("[DEBUG] citrixadc-provider: Reading gslbservicegroup_gslbservicegroupmember_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "gslbservicegroup_gslbservicegroupmember_binding",
		ResourceName:             servicegroupname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservicegroup_gslbservicegroupmember_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		portEqual := int(v["port"].(float64)) == port
		servernameEqual := v["servername"] == servername
		if servernameEqual && portEqual {
			foundIndex = i
			break
		}
	}
	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservicegroup_gslbservicegroupmember_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("hashid", data["hashid"])
	d.Set("ip", data["ip"])
	d.Set("port", data["port"])
	d.Set("publicip", data["publicip"])
	d.Set("publicport", data["publicport"])
	d.Set("servername", data["servername"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("siteprefix", data["siteprefix"])
	d.Set("state", data["state"])
	setToInt("weight", d, data["weight"])
	setToInt("order", d, data["order"])

	return nil

}

func deleteGslbservicegroup_gslbservicegroupmember_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbservicegroup_gslbservicegroupmember_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	name := idSlice[0]
	servername := idSlice[1]
	port := idSlice[2]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("servername:%s", servername))
	args = append(args, fmt.Sprintf("port:%s", port))

	err := client.DeleteResourceWithArgs("gslbservicegroup_gslbservicegroupmember_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
