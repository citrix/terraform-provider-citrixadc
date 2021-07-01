package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strconv"
	"strings"
)

func resourceCitrixAdcServicegroup_servicegroupmember_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createServicegroup_servicegroupmember_bindingFunc,
		Read:          readServicegroup_servicegroupmember_bindingFunc,
		Delete:        deleteServicegroup_servicegroupmember_bindingFunc,
		Schema: map[string]*schema.Schema{
			"customserverid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"dbsttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"hashid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ip": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"servername"},
			},
			"nameserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"serverid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servername": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"ip"},
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createServicegroup_servicegroupmember_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createServicegroup_servicegroupmember_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingIdSlice := make([]string, 0, 3)

	bindingIdSlice = append(bindingIdSlice, d.Get("servicegroupname").(string))

	// Second id component will either be servername or ip
	// ConflictsWith restriction ensures that only one of them is defined
	if v, ok := d.GetOk("servername"); ok {
		bindingIdSlice = append(bindingIdSlice, v.(string))
	}
	if v, ok := d.GetOk("ip"); ok {
		bindingIdSlice = append(bindingIdSlice, v.(string))
	}

	// Third component will be the port if defined
	if v, ok := d.GetOk("port"); ok {
		val := strconv.Itoa(v.(int))
		bindingIdSlice = append(bindingIdSlice, val)
	}

	bindingId := strings.Join(bindingIdSlice, ",")

	servicegroup_servicegroupmember_binding := basic.Servicegroupservicegroupmemberbinding{
		Customserverid:   d.Get("customserverid").(string),
		Dbsttl:           uint64(d.Get("dbsttl").(int)),
		Hashid:           uint32(d.Get("hashid").(int)),
		Ip:               d.Get("ip").(string),
		Nameserver:       d.Get("nameserver").(string),
		Port:             int32(d.Get("port").(int)),
		Serverid:         uint32(d.Get("serverid").(int)),
		Servername:       d.Get("servername").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
		State:            d.Get("state").(string),
		Weight:           uint32(d.Get("weight").(int)),
	}

	err := client.UpdateUnnamedResource(service.Servicegroup_servicegroupmember_binding.Type(), &servicegroup_servicegroupmember_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readServicegroup_servicegroupmember_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this servicegroup_servicegroupmember_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readServicegroup_servicegroupmember_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readServicegroup_servicegroupmember_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)
	servicegroupname := idSlice[0]

	// When ip is defined ADC will create a server with name equal to the ip address
	// So no matter what the user actually defined servername will always be valid search criterion
	servername := idSlice[1]

	// Port is optional
	// Need to check if it actually exists
	port := 0
	var err error
	if len(idSlice) == 3 {
		if port, err = strconv.Atoi(idSlice[2]); err != nil {
			return err
		}
	}

	log.Printf("[DEBUG] citrixadc-provider: Reading servicegroup_servicegroupmember_binding state %v", bindingId)

	findParams := service.FindParams{
		ResourceType:             "servicegroup_servicegroupmember_binding",
		ResourceName:             servicegroupname,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	log.Printf("[DEBUG] citrixadc-provider: dataArr: %v", dataArr)
	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing servicegroup_servicegroupmember_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if port != 0 {
			portEqual := int(v["port"].(float64)) == port
			servernameEqual := v["servername"] == servername
			if servernameEqual && portEqual {
				foundIndex = i
				break
			}
		} else {
			log.Printf("[WARN] citrixadc-provider: port is zero")
			if v["servername"].(string) == servername {
				foundIndex = i
				break
			}
		}
	}
	log.Printf("[WARN] citrixadc-provider: foundIndex: %v", foundIndex)

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing servicegroup_servicegroupmember_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("customserverid", data["customserverid"])
	d.Set("dbsttl", data["dbsttl"])
	d.Set("hashid", data["hashid"])
	d.Set("ip", data["ip"])
	d.Set("nameserver", data["nameserver"])
	d.Set("port", data["port"])
	d.Set("serverid", data["serverid"])
	d.Set("servername", data["servername"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("state", data["state"])
	d.Set("weight", data["weight"])

	return nil

}

func deleteServicegroup_servicegroupmember_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteServicegroup_servicegroupmember_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)
	servicegroupname := idSlice[0]
	servername := idSlice[1]
	port := 0
	var err error
	if len(idSlice) == 3 {
		if port, err = strconv.Atoi(idSlice[2]); err != nil {
			return err
		}
	}
	args := make([]string, 0, 3)

	// When ip is defined ADC will create a server with name equal to the ip address
	// So no matter what the user actually defined servername will always be valid search criterion
	args = append(args, fmt.Sprintf("servername:%s", servername))

	// Port is optional
	if port != 0 {
		args = append(args, fmt.Sprintf("port:%v", port))
	}

	err = client.DeleteResourceWithArgs("servicegroup_servicegroupmember_binding", servicegroupname, args)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
