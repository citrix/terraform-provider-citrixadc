package sslvserver_sslcipher_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslvserverSslcipherBindingResource{}
var _ resource.ResourceWithConfigure = (*SslvserverSslcipherBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslvserverSslcipherBindingResource)(nil)

func NewSslvserverSslcipherBindingResource() resource.Resource {
	return &SslvserverSslcipherBindingResource{}
}

// SslvserverSslcipherBindingResource defines the resource implementation.
type SslvserverSslcipherBindingResource struct {
	client *service.NitroClient
}

func (r *SslvserverSslcipherBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslvserverSslcipherBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver_sslcipher_binding"
}

func (r *SslvserverSslcipherBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslvserverSslcipherBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslvserverSslcipherBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslvserver_sslcipher_binding resource")
	sslvserver_sslcipher_binding := sslvserver_sslcipher_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslvserver_sslcipher_binding.Type(), &sslvserver_sslcipher_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslvserver_sslcipher_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslvserver_sslcipher_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslvserverSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcipherBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslvserverSslcipherBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslvserver_sslcipher_binding resource")

	r.readSslvserverSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcipherBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslvserverSslcipherBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslvserver_sslcipher_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslvserver_sslcipher_binding := sslvserver_sslcipher_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslvserver_sslcipher_binding.Type(), &sslvserver_sslcipher_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslvserver_sslcipher_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslvserver_sslcipher_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslvserver_sslcipher_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslvserverSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverSslcipherBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslvserverSslcipherBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslvserver_sslcipher_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	vservername_value, ok := idMap["vservername"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'vservername' not found in ID")
		return
	}

	// Cipher/alias names can contain '+' and other special characters, so the arg
	// value must be URL-encoded (the NITRO client does not escape arg values).
	args := []string{}
	if val, ok := idMap["ciphername"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ciphername:%s", utils.UrlEncode(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Sslvserver_sslcipher_binding.Type(), vservername_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslvserver_sslcipher_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslvserver_sslcipher_binding binding")
}

// Helper function to read sslvserver_sslcipher_binding data from API
func (r *SslvserverSslcipherBindingResource) readSslvserverSslcipherBindingFromApi(ctx context.Context, data *SslvserverSslcipherBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	vservername_Name, ok := idMap["vservername"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'vservername' not found in ID string")
		return
	}
	ciphername_Name := idMap["ciphername"]

	var dataArr []map[string]interface{}

	// This sslvserver_sslcipher_binding subresource is NOT reflected over the typed GET
	// on this firmware: the plain by-name and ?filter= GETs both return an empty
	// {"message":"Done"} body even when the binding exists, and ?args= is rejected with
	// errorcode 278 ("Invalid argument"). The bind PUT still succeeds and the binding is
	// live. It IS reflected via the umbrella sslvserver_binding endpoint under
	// sslvserver_sslciphersuite_binding[] (keyed by ciphername). So try the typed GET
	// first (in case a firmware does reflect it) and fall back to the umbrella. See the
	// sibling sslservice_sslcipher_binding / sslservicegroup_sslcipher_binding for the
	// same quirk.
	findParams := service.FindParams{
		ResourceType:             service.Sslvserver_sslcipher_binding.Type(),
		ResourceName:             vservername_Name,
		FilterMap:                map[string]string{"ciphername": ciphername_Name},
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_sslcipher_binding, got error: %s", err))
		return
	}

	if len(dataArr) != 0 {
		// Some firmwares do reflect the typed GET - honour it when present.
		foundIndex := -1
		for i, v := range dataArr {
			if val, ok := v["ciphername"].(string); ok && val == ciphername_Name {
				foundIndex = i
				break
			}
		}
		if foundIndex != -1 {
			sslvserver_sslcipher_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
			resolveUnknownComputed(data)
			return
		}
	}

	tflog.Debug(ctx, "sslvserver_sslcipher_binding typed GET returned no matching row; falling back to umbrella sslvserver_binding")
	if row := findCipherBindingViaUmbrella(ctx, r.client, vservername_Name, ciphername_Name, diags); row != nil {
		sslvserver_sslcipher_bindingSetAttrFromGet(ctx, data, row)
	} else {
		if diags.HasError() {
			return
		}
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Guard against leaving any Computed attribute unknown after apply.
	resolveUnknownComputed(data)
}

// findCipherBindingViaUmbrella reads the parent sslvserver_binding object and locates
// the bound cipher row under sslvserver_sslciphersuite_binding[] (matched on ciphername).
// It normalizes the row keys to what sslvserver_sslcipher_bindingSetAttrFromGet expects
// (vservername, ciphername, description; cipheraliasname when present). The vserver row
// does not carry cipheraliasname/cipherdefaulton. Returns nil when not found.
func findCipherBindingViaUmbrella(ctx context.Context, client *service.NitroClient, vservername, ciphername string, diags *diag.Diagnostics) map[string]interface{} {
	umbrella, err := client.FindResourceArrayWithParams(service.FindParams{
		ResourceType:             service.Sslvserver_binding.Type(),
		ResourceName:             vservername,
		ResourceMissingErrorCode: 258,
	})
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_binding (umbrella) for sslvserver_sslcipher_binding, got error: %s", err))
		return nil
	}
	for _, vs := range umbrella {
		raw, ok := vs["sslvserver_sslciphersuite_binding"]
		if !ok {
			continue
		}
		rows, ok := raw.([]interface{})
		if !ok {
			continue
		}
		for _, cb := range rows {
			m, ok := cb.(map[string]interface{})
			if !ok {
				continue
			}
			matchName := ""
			if v, ok := m["cipheraliasname"].(string); ok && v != "" {
				matchName = v
			} else if v, ok := m["ciphername"].(string); ok && v != "" {
				matchName = v
			}
			if matchName != ciphername {
				continue
			}
			// Normalize to the typed-binding key names the setter consumes.
			normalized := map[string]interface{}{
				"vservername": vservername,
				"ciphername":  ciphername,
			}
			if v, ok := m["cipheraliasname"]; ok {
				normalized["cipheraliasname"] = v
			}
			if v, ok := m["description"]; ok {
				normalized["description"] = v
			}
			return normalized
		}
	}
	return nil
}

// resolveUnknownComputed pins any Computed attribute still unknown after a Read to a
// concrete value, so the post-apply state never contains unknowns (Terraform rejects
// "unknown value after apply").
func resolveUnknownComputed(data *SslvserverSslcipherBindingResourceModel) {
	if data.Cipheraliasname.IsUnknown() {
		data.Cipheraliasname = types.StringNull()
	}
	if data.Description.IsUnknown() {
		data.Description = types.StringNull()
	}
}
