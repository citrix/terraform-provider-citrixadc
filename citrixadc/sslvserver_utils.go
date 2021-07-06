package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func syncCiphersuites(d *schema.ResourceData, meta interface{}, vserverName string) error {
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] In syncCiphersuites")
	var ciphersuites interface{}
	var ok bool
	var err error
	var ciphersuiteBindings []map[string]interface{}
	if ciphersuites, ok = d.GetOk("ciphersuites"); ok {
		log.Printf("Configured ciphersuites %v", ciphersuites)
	}

	ciphersuiteBindings, err = client.FindResourceArray(service.Sslvserver_sslciphersuite_binding.Type(), vserverName)

	// When cipher suites are not set and they do not exist there is no sync
	if !ok && len(ciphersuiteBindings) == 0 {
		return nil
	}

	// Catch error in Findresources
	if err != nil && len(ciphersuiteBindings) != 0 {
		log.Printf("Error finding ciphersuite bindings for vserver %v bindings %v with len %v ", vserverName, ciphersuiteBindings, len(ciphersuiteBindings))
		return err
	}

	// Evaluating the equality of resource and target ADC set cipher
	actualCiphersuites := make([]interface{}, 0, len(ciphersuiteBindings))
	for _, ciphersuiteBinding := range ciphersuiteBindings {
		actualCiphersuites = append(actualCiphersuites, ciphersuiteBinding["ciphername"])
	}

	// When existing and set are equal there is no sync
	if slicesEqual(ciphersuites.([]interface{}), actualCiphersuites) {
		log.Printf("Found actual ciphersuites list to be equal to set cipher list")
		return nil
	}
	// Fallthrough

	// First delete all configured cipher bindings
	for _, ciphersuiteBinding := range ciphersuiteBindings {
		ciphername := ciphersuiteBinding["ciphername"].(string)
		argsMap := map[string]string{"ciphername": ciphername}
		log.Printf("Will delete ciphername %v from vserver %v", ciphername, vserverName)
		err := client.DeleteResourceWithArgsMap(service.Sslvserver_sslciphersuite_binding.Type(), vserverName, argsMap)
		if err != nil {
			log.Printf("Error deleting ciphersuite %v from vserver %v", ciphername, vserverName)
			return err
		}
	}

	// Then add all configured bindings
	ciphersuiteslice := ciphersuites.([]interface{})
	for _, ciphername := range ciphersuiteslice {
		resource := ssl.Sslvserverciphersuitebinding{
			Vservername: vserverName,
			Ciphername:  ciphername.(string),
		}
		_, err = client.AddResource(service.Sslvserver_sslciphersuite_binding.Type(), vserverName, resource)
		if err != nil {
			log.Printf("Error binding ciphersuite %v to vserver %v", ciphername, vserverName)
			return err
		}
	}

	return nil
}

func setCiphersuiteData(d *schema.ResourceData, meta interface{}, vserverName string) error {
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] In setCiphersuiteData")

	ciphersuiteBindings, err := client.FindResourceArray(service.Sslvserver_sslciphersuite_binding.Type(), vserverName)
	if err != nil && len(ciphersuiteBindings) != 0 {
		log.Printf("Error retrieving ciphersuite resource array")
		return err
	}
	log.Printf("ciphersuiteBindings %v\n", ciphersuiteBindings)
	ciphersuiteList := make([]interface{}, 0, len(ciphersuiteBindings))
	for _, ciphersuiteBinding := range ciphersuiteBindings {
		ciphersuiteList = append(ciphersuiteList, ciphersuiteBinding["ciphername"])
	}
	log.Printf("Setting ciphersuites to value %v", ciphersuiteList)
	d.Set("ciphersuites", ciphersuiteList)
	return nil
}

func syncCiphers(d *schema.ResourceData, meta interface{}, vserverName string) error {
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] In syncCiphers")
	var ciphers interface{}
	var ok bool
	var err error
	var cipherBindings []map[string]interface{}
	if ciphers, ok = d.GetOk("ciphers"); ok {
		log.Printf("Configured ciphers %v", ciphers)
	}

	findParams := service.FindParams{
		ResourceType:             "sslvserver_sslcipher_binding",
		ResourceName:             vserverName,
		ResourceMissingErrorCode: 258,
	}
	cipherBindings, err = client.FindResourceArrayWithParams(findParams)

	if err != nil {
		return err
	}

	// When ciphers are not set and they do not exist there is no sync
	if !ok && len(cipherBindings) == 0 {
		return nil
	}

	// Catch error in Findresources
	if err != nil && len(cipherBindings) != 0 {
		log.Printf("Error finding ciphersuite bindings for vserver %v bindings %v with len %v ", vserverName, cipherBindings, len(cipherBindings))
		return err
	}

	// Evaluating the equality of resource and target ADC set cipher
	actualCiphers := make([]interface{}, 0, len(cipherBindings))
	for _, cipherBinding := range cipherBindings {
		actualCiphers = append(actualCiphers, cipherBinding["cipheraliasname"])
	}

	// When existing and set are equal there is no sync
	if slicesEqual(ciphers.([]interface{}), actualCiphers) {
		log.Printf("Found actual cipher list to be equal to set cipher list")
		return nil
	}
	// Fallthrough

	// First delete all configured cipher bindings
	for _, cipherBinding := range cipherBindings {
		cipheraliasname := cipherBinding["cipheraliasname"].(string)
		argsMap := map[string]string{"ciphername": cipheraliasname}
		log.Printf("Will delete cipheraliasname %v from vserver %v", cipheraliasname, vserverName)
		err := client.DeleteResourceWithArgsMap(service.Sslvserver_sslcipher_binding.Type(), vserverName, argsMap)
		if err != nil {
			log.Printf("Error deleting cipheraliasname %v from vserver %v", cipheraliasname, vserverName)
			return err
		}
	}

	// Then add all configured bindings
	cipherslice := ciphers.([]interface{})
	for _, cipheraliasname := range cipherslice {
		resource := ssl.Sslvserverciphersuitebinding{
			Vservername: vserverName,
			Ciphername:  cipheraliasname.(string),
		}
		_, err = client.AddResource(service.Sslvserver_sslcipher_binding.Type(), vserverName, resource)
		if err != nil {
			log.Printf("Error binding cipher %v to vserver %v", cipheraliasname, vserverName)
			return err
		}
	}

	return nil
}

func setCipherData(d *schema.ResourceData, meta interface{}, vserverName string) error {
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] In setCipherData")

	cipherBindings, err := client.FindResourceArray(service.Sslvserver_sslcipher_binding.Type(), vserverName)
	if err != nil && len(cipherBindings) != 0 {
		log.Printf("Error retrieving cipher resource array")
		return err
	}
	log.Printf("cipherBindings %v\n", cipherBindings)
	cipherList := make([]interface{}, 0, len(cipherBindings))
	for _, cipherBinding := range cipherBindings {
		cipherList = append(cipherList, cipherBinding["cipheraliasname"])
	}
	log.Printf("Setting ciphers to value %v", cipherList)
	d.Set("ciphers", cipherList)
	return nil
}

func slicesEqual(a, b []interface{}) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func syncSnisslcert(d *schema.ResourceData, meta interface{}, sslvserverName string) error {
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] In syncSnisslcert")
	old, new := d.GetChange("snisslcertkeys")

	var oldlist, newlist []interface{}
	if old == nil {
		oldlist = make([]interface{}, 0)
	} else {
		oldlist = old.(*schema.Set).List()
	}

	newlist = new.(*schema.Set).List()

	var toadd, todelete []string
	toadd = make([]string, 0, len(newlist))

	if old == nil {
		todelete = make([]string, 0)
	} else {
		todelete = make([]string, 0, len(oldlist))
	}

	// certificate key bindings that exist only in the old data will be deleted
	for _, oldcertkey := range oldlist {
		exists := false
		for _, newcertkey := range newlist {
			if oldcertkey.(string) == newcertkey.(string) {
				exists = true
				break
			}
		}
		if !exists {
			todelete = append(todelete, oldcertkey.(string))
		}
	}
	log.Printf("[DEBUG] The following sni certificate key bindings are marked for deletion %v", todelete)

	// certificate key bindings that exist only in the new data will be created
	for _, newcertkey := range newlist {
		exists := false
		for _, oldcertkey := range oldlist {
			if oldcertkey.(string) == newcertkey.(string) {
				exists = true
				break
			}
		}
		if !exists {
			toadd = append(toadd, newcertkey.(string))
		}
	}
	log.Printf("[DEBUG] The following sni certificate key bindings are marked for addition %v", toadd)

	// Do the unbindings first
	for _, snisslcertkey := range todelete {

		args := map[string]string{"certkeyname": snisslcertkey, "snicert": "true"}
		err := client.DeleteResourceWithArgsMap(service.Sslvserver_sslcertkey_binding.Type(), sslvserverName, args)
		if err != nil {
			return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding sni sslcertkey from sslvserver %s", snisslcertkey)
		}
		log.Printf("[DEBUG] netscaler-provider: sni sslcertkey has been unbound from sslvserver for sslcertkey %s ", snisslcertkey)
	}

	// Do the bindings
	for _, snisslcertkey := range toadd {
		binding := ssl.Sslvservercertkeybinding{
			Vservername: sslvserverName,
			Certkeyname: snisslcertkey,
			Snicert:     true,
		}
		log.Printf("[INFO] netscaler-provider:  Binding sni ssl cert %s to sslvserver %s", snisslcertkey, sslvserverName)
		err := client.BindResource(service.Sslvserver.Type(), sslvserverName, service.Sslcertkey.Type(), snisslcertkey, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind sni ssl cert %s to sslvserver %s", snisslcertkey, sslvserverName)
			err2 := client.DeleteResource(service.Lbvserver.Type(), sslvserverName)
			if err2 != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to delete sslvserver %s after bind to sni ssl cert failed", sslvserverName)
				return fmt.Errorf("[ERROR] netscaler-provider:  Failed to delete sslvserver %s after bind to sni ssl cert failed", sslvserverName)
			}
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind sni ssl cert %s to sslvserver %s", snisslcertkey, sslvserverName)
		}
	}

	return nil
}

func snisslcertkeysExist(snisslcertkeys, meta interface{}) error {
	log.Printf("[DEBUG] In snisslcertkeysExist")
	client := meta.(*NetScalerNitroClient).client
	allkeys := snisslcertkeys.(*schema.Set).List()
	missingKeys := make([]string, 0, len(allkeys))
	for _, certkey := range allkeys {
		log.Printf("[DEBUG] checking existence of sslcertkey %v", certkey)
		exists := client.ResourceExists(service.Sslcertkey.Type(), certkey.(string))
		if !exists {
			missingKeys = append(missingKeys, certkey.(string))
		}
	}
	if len(missingKeys) > 0 {
		return fmt.Errorf("The following ssl certificate keys do not exist on target ADC %v", missingKeys)
	} else {
		return nil
	}
}

func readSslcerts(d *schema.ResourceData, meta interface{}, sslvserverName string) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcerts")
	client := meta.(*NetScalerNitroClient).client
	bindings, err := client.FindAllBoundResources(service.Sslvserver.Type(), sslvserverName, service.Sslcertkey.Type())
	if err != nil {
		log.Printf("[WARN] netscaler-provider: sslvserver binding to ssl error %s", sslvserverName)
		return nil
	}
	var boundCert string
	snicerts := make([]string, 0, len(bindings))
	for _, binding := range bindings {
		cert, ok := binding["certkeyname"]
		snicert, ok2 := binding["snicert"]
		log.Printf("Reading ssl binding certkeyname %v, %v", cert, ok)
		log.Printf("Reading ssl binding snicert %v, %v", snicert, ok2)
		if ok && ok2 && snicert == false {
			boundCert = cert.(string)
		} else if ok && ok2 && snicert == true {
			snicerts = append(snicerts, cert.(string))
		}
	}
	d.Set("sslcertkey", boundCert)
	d.Set("snisslcertkeys", snicerts)
	return nil
}

func sslpolicybindingMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In sslpolicybindingMappingHash")
	var buf bytes.Buffer

	// All keys added in alphabetical order.
	m := v.(map[string]interface{})
	if d, ok := m["gotopriorityexpression"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["invoke"]; ok {
		buf.WriteString(fmt.Sprintf("%t-", d.(bool)))
	}

	if d, ok := m["labelname"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["labeltype"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["policyname"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["priotity"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}

	if d, ok := m["type"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	return hashcode.String(buf.String())
}

func readSslpolicyBindings(d *schema.ResourceData, meta interface{}, sslvserverName string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readSslpolicyBindings")
	client := meta.(*NetScalerNitroClient).client

	// Ignore sslpolicy bindings unless explicitely in configuration
	// If not it is confused with binding defined by the explicit sslvserver_sslpolicy_binding resource
	if _, ok := d.GetOk("sslpolicybinding"); !ok {
		return nil
	}

	findParams := service.FindParams{
		ResourceType:             "sslvserver_sslpolicy_binding",
		ResourceName:             sslvserverName,
		ResourceMissingErrorCode: 258,
	}

	bindings, err := client.FindResourceArrayWithParams(findParams)

	if err != nil {
		// 1544 is returned when we try the binding on a non SSL vserver
		// We abort the rest of the execution effectively setting the sslpolicybindings to the empty set
		if strings.Contains(err.Error(), "\"errorcode\": 1544") {
			return nil
		} else {
			return err
		}
	}

	processedBindings := make([]interface{}, 0, len(bindings))
	for _, singleBinding := range bindings {
		var err error
		processedBindingEntry := make(map[string]interface{})
		processedBindingEntry["gotopriorityexpression"] = singleBinding["gotopriorityexpression"].(string)
		processedBindingEntry["invoke"] = singleBinding["invoke"].(bool)
		if _, ok := singleBinding["labelname"]; ok {
			processedBindingEntry["labelname"] = singleBinding["labelname"].(string)
		}
		if _, ok := singleBinding["labeltype"]; ok {
			processedBindingEntry["labeltype"] = singleBinding["labeltype"].(string)
		}
		processedBindingEntry["policyname"] = singleBinding["policyname"].(string)
		if processedBindingEntry["priority"], err = strconv.Atoi(singleBinding["priority"].(string)); err != nil {
			return err
		}
		if _, ok := singleBinding["type"]; ok {
			processedBindingEntry["type"] = singleBinding["type"].(string)
		}
		processedBindings = append(processedBindings, processedBindingEntry)
	}

	updatedSet := schema.NewSet(sslpolicybindingMappingHash, processedBindings)
	log.Printf("[DEBUG] citrixadc-provider: Updated sslpolicybinding set %v", updatedSet)
	if err := d.Set("sslpolicybinding", updatedSet); err != nil {
		return err
	}
	return nil
}

func updateSslpolicyBindings(d *schema.ResourceData, meta interface{}, sslvserverName string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslpolicyBindings")

	oldSet, newSet := d.GetChange("sslpolicybinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, binding := range remove.List() {
		if err := deleteSingleSslpolicyBinding(d, meta, binding.(map[string]interface{}), sslvserverName); err != nil {
			return err
		}
	}

	for _, binding := range add.List() {
		if err := addSingleSslpolicyBinding(d, meta, binding.(map[string]interface{}), sslvserverName); err != nil {
			return err
		}
	}
	return nil
}

func deleteSslpolicyBindings(d *schema.ResourceData, meta interface{}, sslvserverName string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslpolicyBindings")

	if bindings, ok := d.GetOk("sslpolicybinding"); ok {
		for _, binding := range bindings.(*schema.Set).List() {
			if err := deleteSingleSslpolicyBinding(d, meta, binding.(map[string]interface{}), sslvserverName); err != nil {
				return err
			}
		}
	}
	return nil
}

func addSingleSslpolicyBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}, sslvserverName string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleSslpolicyBinding")

	client := meta.(*NetScalerNitroClient).client

	bindingStruct := ssl.Sslvserverpolicybinding{}
	bindingStruct.Vservername = sslvserverName

	if val, ok := binding["gotopriorityexpression"]; ok {
		bindingStruct.Gotopriorityexpression = val.(string)
	}

	if val, ok := binding["invoke"]; ok {
		bindingStruct.Invoke = val.(bool)
	}

	if val, ok := binding["labelname"]; ok {
		bindingStruct.Labelname = val.(string)
	}

	if val, ok := binding["labeltype"]; ok {
		bindingStruct.Labeltype = val.(string)
	}

	if val, ok := binding["policyname"]; ok {
		bindingStruct.Policyname = val.(string)
	}

	if val, ok := binding["priority"]; ok {
		bindingStruct.Priority = uint32(val.(int))
	}

	if val, ok := binding["type"]; ok {
		bindingStruct.Type = val.(string)
	}

	if _, err := client.UpdateResource("sslvserver_sslpolicy_binding", sslvserverName, bindingStruct); err != nil {
		return err
	}
	return nil
}

func deleteSingleSslpolicyBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}, sslvserverName string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleSslpolicyBinding")
	client := meta.(*NetScalerNitroClient).client

	// Construct args from binding data
	args := make([]string, 0, 3)

	if d, ok := binding["policyname"]; ok {
		s := fmt.Sprintf("policyname:%s", d.(string))
		args = append(args, s)
	}

	if d, ok := binding["priority"]; ok {
		s := fmt.Sprintf("priority:%d", d.(int))
		args = append(args, s)
	}

	if d, ok := binding["type"]; ok {
		if d != "" {
			s := fmt.Sprintf("type:%s", d.(string))
			args = append(args, s)
		}
	}

	log.Printf("args %v", args)
	if err := client.DeleteResourceWithArgs("sslvserver_sslpolicy_binding", sslvserverName, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting sslpolicy binding %v\n", binding)
		return err
	}
	return nil
}
