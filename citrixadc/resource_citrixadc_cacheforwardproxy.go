package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"strconv"
	"log"
)

func resourceCitrixAdcCacheforwardproxy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCacheforwardproxyFunc,
		Read:          readCacheforwardproxyFunc,
		Delete:        deleteCacheforwardproxyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createCacheforwardproxyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCacheforwardproxyFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheforwardproxyName := d.Get("ipaddress").(string)
	cacheforwardproxyid := fmt.Sprintf("%s,%s", cacheforwardproxyName, strconv.Itoa(d.Get("port").(int)))
	cacheforwardproxy := cache.Cacheforwardproxy{
		Ipaddress: d.Get("ipaddress").(string),
		Port:      d.Get("port").(int),
	}

	_, err := client.AddResource(service.Cacheforwardproxy.Type(), cacheforwardproxyName, &cacheforwardproxy)
	if err != nil {
		return err
	}

	d.SetId(cacheforwardproxyid)

	err = readCacheforwardproxyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cacheforwardproxy but we can't read it ?? %s", cacheforwardproxyid)
		return nil
	}
	return nil
}

func readCacheforwardproxyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCacheforwardproxyFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheforwardproxyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cacheforwardproxy state %s", cacheforwardproxyName)
	dataArr, err := client.FindAllResources(service.Cacheforwardproxy.Type())
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cacheforwardproxy state %s", cacheforwardproxyName)
		d.SetId("")
		return nil
	}
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing cacheforwardproxy state %s", cacheforwardproxyName)
		d.SetId("")
		return nil
	}
	foundIndex := -1
	for i, v := range dataArr {
		if v["ipaddress"].(string) == d.Get("ipaddress").(string) &&
			int(v["port"].(float64)) == d.Get("port").(int) {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindAllResources ipaddress or port not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing cacheforwardproxy state %s", cacheforwardproxyName)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]
	d.Set("ipaddress", data["ipaddress"])
	d.Set("port", data["port"])

	return nil

}

func deleteCacheforwardproxyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCacheforwardproxyFunc")
	client := meta.(*NetScalerNitroClient).client
	cacheforwardproxyName := d.Get("ipaddress").(string)
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("port:%v", d.Get("port").(int)))
	err := client.DeleteResourceWithArgs(service.Cacheforwardproxy.Type(), cacheforwardproxyName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
