package citrixadc

import (
	"fmt"
	"log"
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCitrixAdcLinkset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLinksetFunc,
		Read:          readLinksetFunc,
		Delete:        deleteLinksetFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"linkset_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"interfacebinding": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true, // to avoid this error: https://github.com/hashicorp/terraform-plugin-sdk/blob/master/helper/schema/resource.go#L635
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func createLinksetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLinksetFunc")
	client := meta.(*NetScalerNitroClient).client
	linksetName := d.Get("linkset_id").(string)
	linkset := network.Linkset{
		Id: linksetName,
	}

	_, err := client.AddResource(service.Linkset.Type(), "", &linkset)
	if err != nil {
		return err
	}

	d.SetId(linksetName)

	err = updateLinksetInterfaceBindings(d, meta)
	if err != nil {
		return err
	}

	err = readLinksetFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this linkset but we can't read it ?? %s", linksetName)
		return nil
	}
	return nil
}

func readLinksetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLinksetFunc")
	client := meta.(*NetScalerNitroClient).client
	linksetName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading linkset state %s", linksetName)
	// double encode value part as it contains `/`
	//linksetNameEscaped := url.QueryEscape(url.QueryEscape(linksetName))
	linksetNameEscaped := url.PathEscape(url.QueryEscape(linksetName))
	data, err := client.FindResource(service.Linkset.Type(), linksetNameEscaped)

	err = readLinksetInterfaceBindings(d, meta)
	if err != nil {
		return err
	}

	d.Set("linkset_id", data["id"])

	return nil

}

func deleteLinksetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLinksetFunc")
	client := meta.(*NetScalerNitroClient).client
	linksetName := d.Id()
	// double encode value part as it contains `/`
	linksetNameEscaped := url.QueryEscape(url.QueryEscape(linksetName))
	err := client.DeleteResource(service.Linkset.Type(), linksetNameEscaped)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func deleteSingleLinksetInterfaceBinding(d *schema.ResourceData, meta interface{}, ifnum string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleLinksetInterfaceBinding")
	client := meta.(*NetScalerNitroClient).client

	linksetName := d.Get("linkset_id").(string)
	args := make([]string, 0, 1)
	// double encode value part as it contains `/`
	ifnumEscaped := url.QueryEscape(url.QueryEscape(ifnum))

	s := fmt.Sprintf("ifnum:%s", ifnumEscaped)
	args = append(args, s)

	log.Printf("args is %v", args)
	linksetNameEscaped := url.QueryEscape(url.QueryEscape(linksetName))

	if err := client.DeleteResourceWithArgs(service.Linkset_interface_binding.Type(), linksetNameEscaped, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting interface binding %v\n", ifnum)
		return err
	}

	return nil
}

func addSingleLinksetInterfaceBinding(d *schema.ResourceData, meta interface{}, ifnum string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleLinksetInterfaceBinding")
	client := meta.(*NetScalerNitroClient).client

	bindingStruct := network.Linksetinterfacebinding{}
	bindingStruct.Id = d.Get("linkset_id").(string)
	bindingStruct.Ifnum = ifnum

	// We need to do a HTTP PUT hence the UpdateResource
	if err := client.UpdateUnnamedResource(service.Linkset_interface_binding.Type(), bindingStruct); err != nil {
		return err
	}
	return nil
}

func updateLinksetInterfaceBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLinksetInterfaceBindings")
	oldSet, newSet := d.GetChange("interfacebinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, ifnum := range remove.List() {
		if err := deleteSingleLinksetInterfaceBinding(d, meta, ifnum.(string)); err != nil {
			return err
		}
	}

	for _, ifnum := range add.List() {
		if err := addSingleLinksetInterfaceBinding(d, meta, ifnum.(string)); err != nil {
			return err
		}
	}
	return nil
}

func readLinksetInterfaceBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readLinksetInterfaceBindings")
	client := meta.(*NetScalerNitroClient).client
	linksetName := d.Get("linkset_id").(string)
	// double encode value part as it contains `/`
	linksetNameEscaped := url.QueryEscape(url.QueryEscape(linksetName))

	bindings, _ := client.FindResourceArray(service.Linkset_interface_binding.Type(), linksetNameEscaped)
	log.Printf("bindings %v\n", bindings)

	processedBindings := make([]interface{}, len(bindings))
	for i, val := range bindings {
		processedBindings[i] = val["ifnum"].(string)
	}

	updatedSet := processedBindings
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("interfacebinding", updatedSet); err != nil {
		return err
	}
	return nil
}
