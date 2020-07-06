package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/network"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcIpset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIpsetFunc,
		Read:          readIpsetFunc,
		Delete:        deleteIpsetFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nsipbinding": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true, // to avoid this error: https://github.com/hashicorp/terraform/blob/master/helper/schema/resource.go#L635
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"nsip6binding": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true, // to avoid this error: https://github.com/hashicorp/terraform/blob/master/helper/schema/resource.go#L635
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func createIpsetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createIpsetFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsetName := d.Get("name").(string)

	ipset := network.Ipset{
		Name: d.Get("name").(string),
		Td:   d.Get("td").(int),
	}

	_, err := client.AddResource(netscaler.Ipset.Type(), ipsetName, &ipset)
	if err != nil {
		return err
	}

	d.SetId(ipsetName)

	err = updateIpsetNsipBindings(d, meta)
	if err != nil {
		return err
	}

	err = updateIpsetNsip6Bindings(d, meta)
	if err != nil {
		return err
	}

	err = readIpsetFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this ipset but we can't read it ?? %s", ipsetName)
		return nil
	}
	return nil
}

func readIpsetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readIpsetFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsetName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading ipset state %s", ipsetName)
	data, err := client.FindResource(netscaler.Ipset.Type(), ipsetName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ipset state %s", ipsetName)
		d.SetId("")
		return nil
	}

	err = readIpsetNsipBindings(d, meta)
	if err != nil {
		return err
	}

	err = readIpsetNsip6Bindings(d, meta)
	if err != nil {
		return err
	}

	d.Set("name", data["name"])
	d.Set("td", data["td"])

	return nil

}

func deleteIpsetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIpsetFunc")
	client := meta.(*NetScalerNitroClient).client
	ipsetName := d.Id()
	err := client.DeleteResource(netscaler.Ipset.Type(), ipsetName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func deleteSingleIpsetNsipBinding(d *schema.ResourceData, meta interface{}, nsip string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleIpsetNsipBinding")
	client := meta.(*NetScalerNitroClient).client

	name := d.Get("name").(string)
	args := make([]string, 0, 1)

	s := fmt.Sprintf("ipaddress:%s", nsip)
	args = append(args, s)

	log.Printf("args is %v", args)

	if err := client.DeleteResourceWithArgs("ipset_nsip_binding", name, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting nsip binding %v\n", nsip)
		return err
	}

	return nil
}

func addSingleIpsetNsipBinding(d *schema.ResourceData, meta interface{}, nsip string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleIpsetNsipBinding")
	client := meta.(*NetScalerNitroClient).client

	bindingStruct := network.Ipsetnsipbinding{}
	bindingStruct.Name = d.Get("name").(string)
	bindingStruct.Ipaddress = nsip

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource("ipset_nsip_binding", bindingStruct.Name, bindingStruct); err != nil {
		return err
	}
	return nil
}

func updateIpsetNsipBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIpsetNsipBindings")
	oldSet, newSet := d.GetChange("nsipbinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, nsip := range remove.List() {
		if err := deleteSingleIpsetNsipBinding(d, meta, nsip.(string)); err != nil {
			return err
		}
	}

	for _, nsip := range add.List() {
		if err := addSingleIpsetNsipBinding(d, meta, nsip.(string)); err != nil {
			return err
		}
	}
	return nil
}

func readIpsetNsipBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readIpsetNsipBindings")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	bindings, _ := client.FindResourceArray("ipset_nsip_binding", name)
	log.Printf("bindings %v\n", bindings)

	processedBindings := make([]interface{}, len(bindings))
	for i, val := range bindings {
		processedBindings[i] = val["ipaddress"].(string)
	}

	updatedSet := processedBindings
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("nsipbinding", updatedSet); err != nil {
		return err
	}
	return nil
}

func deleteSingleIpsetNsip6Binding(d *schema.ResourceData, meta interface{}, nsip6 string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleIpsetNsip6Binding")
	client := meta.(*NetScalerNitroClient).client

	name := d.Get("name").(string)
	args := make([]string, 0, 1)

	s := fmt.Sprintf("ipaddress:%s", nsip6)
	args = append(args, s)

	log.Printf("args is %v", args)

	if err := client.DeleteResourceWithArgs("ipset_nsip6_binding", name, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting nsip6 binding %v\n", nsip6)
		return err
	}

	return nil
}

func addSingleIpsetNsip6Binding(d *schema.ResourceData, meta interface{}, nsip6 string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleIpsetNsip6Binding")
	client := meta.(*NetScalerNitroClient).client

	bindingStruct := network.Ipsetnsip6binding{}
	bindingStruct.Name = d.Get("name").(string)
	bindingStruct.Ipaddress = nsip6 //strings.SplitN(nsip6, "/", -1)[0]

	log.Printf("[DEBUG] bindingStruct: %v\n", bindingStruct)

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource("ipset_nsip6_binding", bindingStruct.Name, bindingStruct); err != nil {
		return err
	}
	return nil
}

func updateIpsetNsip6Bindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIpsetNsip6Bindings")
	oldSet, newSet := d.GetChange("nsip6binding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, nsip6 := range remove.List() {
		if err := deleteSingleIpsetNsip6Binding(d, meta, nsip6.(string)); err != nil {
			return err
		}
	}

	for _, nsip6 := range add.List() {
		if err := addSingleIpsetNsip6Binding(d, meta, nsip6.(string)); err != nil {
			return err
		}
	}
	return nil
}

func readIpsetNsip6Bindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readIpsetNsip6Bindings")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	bindings, _ := client.FindResourceArray("ipset_nsip6_binding", name)
	log.Printf("bindings %v\n", bindings)

	processedBindings := make([]interface{}, len(bindings))
	for i, val := range bindings {
		processedBindings[i] = val["ipaddress"].(string)
	}

	updatedSet := processedBindings
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("nsip6binding", updatedSet); err != nil {
		return err
	}
	return nil
}
