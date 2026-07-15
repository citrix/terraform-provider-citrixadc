package aaagroup_intranetip6_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AaagroupIntranetip6BindingResource{}
var _ resource.ResourceWithConfigure = (*AaagroupIntranetip6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaagroupIntranetip6BindingResource)(nil)

func NewAaagroupIntranetip6BindingResource() resource.Resource {
	return &AaagroupIntranetip6BindingResource{}
}

// AaagroupIntranetip6BindingResource defines the resource implementation.
type AaagroupIntranetip6BindingResource struct {
	client *service.NitroClient
}

func (r *AaagroupIntranetip6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaagroupIntranetip6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup_intranetip6_binding"
}

func (r *AaagroupIntranetip6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaagroupIntranetip6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaagroupIntranetip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaagroup_intranetip6_binding resource")
	aaagroup_intranetip6_binding := aaagroup_intranetip6_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaagroup_intranetip6_binding.Type(), &aaagroup_intranetip6_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaagroup_intranetip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaagroup_intranetip6_binding resource")

	// Set ID for the resource before reading state.
	// Composite ID = groupname,intranetip6,numaddr (all three persisted so Delete
	// has every arg it needs). IPv6 colons are percent-encoded inside the helper.
	data.Id = types.StringValue(aaagroup_intranetip6_bindingComposeId(
		data.Groupname.ValueString(),
		data.Intranetip6.ValueString(),
		data.Numaddr.ValueInt64(),
	))

	// Read the updated state back
	r.readAaagroupIntranetip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupIntranetip6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaagroupIntranetip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaagroup_intranetip6_binding resource")

	r.readAaagroupIntranetip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource was deleted out-of-band - remove from state for self-healing
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupIntranetip6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AaagroupIntranetip6BindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Pattern 5: aaagroup_intranetip6_binding has no NITRO update endpoint and
	// every schema attribute uses RequiresReplace, so Terraform will never call
	// Update for an attribute change (it forces recreation instead). This body is
	// a documented no-op that simply re-reads current state.
	tflog.Debug(ctx, "Update is a no-op for aaagroup_intranetip6_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readAaagroupIntranetip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupIntranetip6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaagroupIntranetip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaagroup_intranetip6_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	groupname_value, ok := idMap["groupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'groupname' not found in ID")
		return
	}

	// Live ADC (10.101.132.121, errorcode evidence) determines the delete arg set,
	// NOT the NITRO doc:
	//   - args=numaddr only       -> 1093 "Argument pre-requisite missing [numaddr, intranetIP6]"
	//   - args=intranetip6 only   -> accepted (2984 only when the binding is absent)
	//   - args=intranetip6,numaddr-> accepted, but a Go map is UNORDERED so numaddr
	//                                could be emitted first and trip the 1093 prereq.
	// Therefore Delete passes ONLY intranetip6 (the doc's claim that numaddr is a
	// required delete arg is wrong - numaddr is a create-only payload field).
	//
	// IPv6-colon handling: ParseIdString returns the intranetip6 value already
	// URL-DECODED (raw ':' restored). The NITRO client joins delete args as
	// "key:value" WITHOUT re-encoding, so a raw IPv6 value like "fd00::1" would
	// corrupt the args string. Re-encode it (utils.UrlEncode) so every ':' becomes
	// '%3A'.
	var deleteArgs []string
	if val, ok := idMap["intranetip6"]; ok && val != "" {
		deleteArgs = append(deleteArgs, "intranetip6:"+utils.UrlEncode(val))
	}

	err = r.client.DeleteResourceWithArgs(service.Aaagroup_intranetip6_binding.Type(), groupname_value, deleteArgs)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaagroup_intranetip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaagroup_intranetip6_binding binding")
}

// Helper function to read aaagroup_intranetip6_binding data from API
func (r *AaagroupIntranetip6BindingResource) readAaagroupIntranetip6BindingFromApi(ctx context.Context, data *AaagroupIntranetip6BindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	groupname_Name, ok := idMap["groupname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'groupname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Aaagroup_intranetip6_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup_intranetip6_binding, got error: %s", err))
		return
	}

	// Resource is missing - signal deletion for self-healing
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id.
	// Filter on both intranetip6 AND numaddr (both are part of the composite ID
	// and the delete key). idMap values are URL-decoded by ParseIdString, so the
	// intranetip6 value here has its raw ':' restored and matches the API value
	// directly. numaddr is compared numerically.
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check intranetip6
		if idVal, ok := idMap["intranetip6"]; ok {
			if val, ok := v["intranetip6"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["intranetip6"].(string); ok {
			match = false
			continue
		}

		// Check numaddr
		if idVal, ok := idMap["numaddr"]; ok {
			if val, ok := v["numaddr"]; ok && val != nil {
				apiNum, apiErr := utils.ConvertToInt64(val)
				idNum, idErr := utils.ConvertToInt64(idVal)
				if apiErr != nil || idErr != nil || apiNum != idNum {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing - signal deletion for self-healing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	aaagroup_intranetip6_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
