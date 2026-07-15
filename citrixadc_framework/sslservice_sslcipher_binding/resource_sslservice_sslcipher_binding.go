package sslservice_sslcipher_binding

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
var _ resource.Resource = &SslserviceSslcipherBindingResource{}
var _ resource.ResourceWithConfigure = (*SslserviceSslcipherBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslserviceSslcipherBindingResource)(nil)

func NewSslserviceSslcipherBindingResource() resource.Resource {
	return &SslserviceSslcipherBindingResource{}
}

// SslserviceSslcipherBindingResource defines the resource implementation.
type SslserviceSslcipherBindingResource struct {
	client *service.NitroClient
}

func (r *SslserviceSslcipherBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslserviceSslcipherBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservice_sslcipher_binding"
}

func (r *SslserviceSslcipherBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslserviceSslcipherBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslserviceSslcipherBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservice_sslcipher_binding resource")
	sslservice_sslcipher_binding := sslservice_sslcipher_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslservice_sslcipher_binding.Type(), &sslservice_sslcipher_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservice_sslcipher_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslservice_sslcipher_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslserviceSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslcipherBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslserviceSslcipherBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservice_sslcipher_binding resource")

	r.readSslserviceSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// If the binding is gone out-of-band, remove it from state so a subsequent apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslcipherBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslserviceSslcipherBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslservice_sslcipher_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslservice_sslcipher_binding := sslservice_sslcipher_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslservice_sslcipher_binding.Type(), &sslservice_sslcipher_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservice_sslcipher_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslservice_sslcipher_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslservice_sslcipher_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslserviceSslcipherBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslcipherBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslserviceSslcipherBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservice_sslcipher_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicename_value, ok := idMap["servicename"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicename' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["ciphername"]; ok && val != "" {
		argsMap["ciphername"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslservice_sslcipher_binding.Type(), servicename_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslservice_sslcipher_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslservice_sslcipher_binding binding")
}

// Helper function to read sslservice_sslcipher_binding data from API
func (r *SslserviceSslcipherBindingResource) readSslserviceSslcipherBindingFromApi(ctx context.Context, data *SslserviceSslcipherBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	servicename_Name, ok := idMap["servicename"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'servicename' not found in ID string")
		return
	}
	ciphername_Name := idMap["ciphername"]

	var dataArr []map[string]interface{}

	// This sslservice_sslcipher_binding subresource is NOT reflected over the typed
	// GET on this firmware: the plain by-name and ?filter= GETs both return an empty
	// {"message":"Done"} body even when the binding exists, and ?args= is rejected
	// with errorcode 278 ("Invalid argument"). The bind PUT still succeeds and the
	// binding is live. It IS reflected via the umbrella sslservice_binding endpoint
	// under sslservice_sslciphersuite_binding[] (keyed by ciphername). So try the
	// typed GET first (in case a firmware does reflect it) and fall back to the
	// umbrella. See the sibling sslprofile_sslciphersuite_binding for the same quirk.
	findParams := service.FindParams{
		ResourceType:             service.Sslservice_sslcipher_binding.Type(),
		ResourceName:             servicename_Name,
		FilterMap:                map[string]string{"ciphername": ciphername_Name},
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservice_sslcipher_binding, got error: %s", err))
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
			sslservice_sslcipher_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
			resolveUnknownComputed(data)
			return
		}
	}

	tflog.Debug(ctx, "sslservice_sslcipher_binding typed GET returned no matching row; falling back to umbrella sslservice_binding")
	if row := findCipherBindingViaUmbrella(ctx, r.client, servicename_Name, ciphername_Name, diags); row != nil {
		sslservice_sslcipher_bindingSetAttrFromGet(ctx, data, row)
	} else {
		if diags.HasError() {
			return
		}
		// Neither the typed GET nor the umbrella endpoint has the binding: it was
		// deleted out-of-band. Signal "gone" so Read removes it from state and a
		// subsequent apply re-creates it.
		data.Id = types.StringNull()
		return
	}

	// Guard against leaving any Computed attribute unknown after apply.
	resolveUnknownComputed(data)
}

// findCipherBindingViaUmbrella reads the parent sslservice_binding object and locates
// the bound cipher row under sslservice_sslciphersuite_binding[] (matched on ciphername).
// It normalizes the row keys to what sslservice_sslcipher_bindingSetAttrFromGet expects
// (servicename, ciphername, description, cipherdefaulton). Returns nil when not found.
func findCipherBindingViaUmbrella(ctx context.Context, client *service.NitroClient, servicename, ciphername string, diags *diag.Diagnostics) map[string]interface{} {
	umbrella, err := client.FindResourceArrayWithParams(service.FindParams{
		ResourceType:             service.Sslservice_binding.Type(),
		ResourceName:             servicename,
		ResourceMissingErrorCode: 258,
	})
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservice_binding (umbrella) for sslservice_sslcipher_binding, got error: %s", err))
		return nil
	}
	for _, svc := range umbrella {
		raw, ok := svc["sslservice_sslciphersuite_binding"]
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
				"servicename": servicename,
				"ciphername":  ciphername,
			}
			if v, ok := m["cipheraliasname"]; ok {
				normalized["cipheraliasname"] = v
			}
			if v, ok := m["cipherdefaulton"]; ok {
				normalized["cipherdefaulton"] = v
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
func resolveUnknownComputed(data *SslserviceSslcipherBindingResourceModel) {
	if data.Cipheraliasname.IsUnknown() {
		data.Cipheraliasname = types.StringNull()
	}
	if data.Cipherdefaulton.IsUnknown() {
		data.Cipherdefaulton = types.Int64Null()
	}
	if data.Description.IsUnknown() {
		data.Description = types.StringNull()
	}
}
