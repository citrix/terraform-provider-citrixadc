package sslprofile_sslciphersuite_binding

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
var _ resource.Resource = &SslprofileSslciphersuiteBindingResource{}
var _ resource.ResourceWithConfigure = (*SslprofileSslciphersuiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslprofileSslciphersuiteBindingResource)(nil)

func NewSslprofileSslciphersuiteBindingResource() resource.Resource {
	return &SslprofileSslciphersuiteBindingResource{}
}

// SslprofileSslciphersuiteBindingResource defines the resource implementation.
type SslprofileSslciphersuiteBindingResource struct {
	client *service.NitroClient
}

func (r *SslprofileSslciphersuiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslprofileSslciphersuiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslprofile_sslciphersuite_binding"
}

func (r *SslprofileSslciphersuiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslprofileSslciphersuiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslprofileSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslprofile_sslciphersuite_binding resource")
	sslprofile_sslciphersuite_binding := sslprofile_sslciphersuite_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslprofile_sslciphersuite_binding.Type(), &sslprofile_sslciphersuite_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslprofile_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslprofile_sslciphersuite_binding resource")

	// Set ID for the resource before reading state
	// Composite ID = name,ciphername
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslprofileSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslciphersuiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslprofileSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslprofile_sslciphersuite_binding resource")

	r.readSslprofileSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslprofileSslciphersuiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslprofileSslciphersuiteBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Pattern 5: no NITRO update endpoint for this binding; all attributes are
	// RequiresReplace, so Update is a no-op that just reconciles state.
	tflog.Debug(ctx, "Update is a no-op for sslprofile_sslciphersuite_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSslprofileSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileSslciphersuiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslprofileSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslprofile_sslciphersuite_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs (parent name + member arg)
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "ciphername"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	args := make([]string, 0, 1)
	if val, ok := idMap["ciphername"]; ok && val != "" {
		args = append(args, "ciphername:"+utils.UrlEncode(val))
	}

	err = r.client.DeleteResourceWithArgs(service.Sslprofile_sslciphersuite_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslprofile_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslprofile_sslciphersuite_binding binding")
}

// Helper function to read sslprofile_sslciphersuite_binding data from API
func (r *SslprofileSslciphersuiteBindingResource) readSslprofileSslciphersuiteBindingFromApi(ctx context.Context, data *SslprofileSslciphersuiteBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "ciphername"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}
	ciphername_Name := idMap["ciphername"]

	var dataArr []map[string]interface{}

	// This sslprofile_* binding subresource rejects the plain by-name GET (returns an
	// empty {"message":"Done"} body) and also rejects ?args=. It only narrows results
	// via the documented ?filter=<key>:<value> query parameter, which FindParams.FilterMap
	// emits. See the sibling nstimer_autoscalepolicy_binding for the same firmware quirk.
	findParams := service.FindParams{
		ResourceType:             service.Sslprofile_sslciphersuite_binding.Type(),
		ResourceName:             name_Name,
		FilterMap:                map[string]string{"ciphername": ciphername_Name},
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile_sslciphersuite_binding, got error: %s", err))
		return
	}

	// The typed binding GET is not reflected over REST on this firmware (always an empty
	// "Done" body even when the binding exists - the create PUT succeeds and the binding
	// is live). It IS reflected via the umbrella sslprofile_binding endpoint under
	// sslprofile_sslcipher_binding[] (keyed by cipheraliasname for aliases / ciphername
	// for individual ciphers). So when the typed GET returns nothing, fall back to the
	// umbrella to populate the model from live config; only if the umbrella also has no
	// matching row do we preserve existing state.
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
			sslprofile_sslciphersuite_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
			resolveUnknownComputed(data)
			return
		}
	}

	tflog.Debug(ctx, "sslprofile_sslciphersuite_binding typed GET returned no matching row; falling back to umbrella sslprofile_binding")
	if row := findCipherBindingViaUmbrella(ctx, r.client, name_Name, ciphername_Name, diags); row != nil {
		sslprofile_sslciphersuite_bindingSetAttrFromGet(ctx, data, row)
	} else if diags.HasError() {
		return
	} else {
		// Neither the typed GET nor the umbrella endpoint has the binding: it was
		// deleted out-of-band. Signal "gone" so Read removes it from state and a
		// subsequent apply re-creates it.
		data.Id = types.StringNull()
		return
	}

	// description (and cipherpriority when the user omitted it) may be absent; guard
	// against leaving any Computed attribute unknown after apply.
	resolveUnknownComputed(data)
}

// findCipherBindingViaUmbrella reads the parent sslprofile_binding object and locates the
// bound cipher row (matched on cipheraliasname for aliases or ciphername for individual
// ciphers). It normalizes the row keys to what sslprofile_sslciphersuite_bindingSetAttrFromGet
// expects (ciphername, cipherpriority, description). Returns nil when not found.
func findCipherBindingViaUmbrella(ctx context.Context, client *service.NitroClient, name, ciphername string, diags *diag.Diagnostics) map[string]interface{} {
	umbrella, err := client.FindResourceArrayWithParams(service.FindParams{
		ResourceType:             service.Sslprofile_binding.Type(),
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	})
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile_binding (umbrella) for sslprofile_sslciphersuite_binding, got error: %s", err))
		return nil
	}
	for _, prof := range umbrella {
		raw, ok := prof["sslprofile_sslcipher_binding"]
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
				"name":       name,
				"ciphername": ciphername,
			}
			if v, ok := m["cipherpriority"]; ok {
				normalized["cipherpriority"] = v
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
// "unknown value after apply"). This endpoint never echoes description; cipherpriority,
// when supplied, is already carried in the plan/state.
func resolveUnknownComputed(data *SslprofileSslciphersuiteBindingResourceModel) {
	if data.Description.IsUnknown() {
		data.Description = types.StringNull()
	}
	if data.Cipherpriority.IsUnknown() {
		data.Cipherpriority = types.Int64Null()
	}
}
