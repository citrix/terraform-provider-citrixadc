package lsnappsprofile_port_binding

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
var _ resource.Resource = &LsnappsprofilePortBindingResource{}
var _ resource.ResourceWithConfigure = (*LsnappsprofilePortBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsnappsprofilePortBindingResource)(nil)

func NewLsnappsprofilePortBindingResource() resource.Resource {
	return &LsnappsprofilePortBindingResource{}
}

// LsnappsprofilePortBindingResource defines the resource implementation.
type LsnappsprofilePortBindingResource struct {
	client *service.NitroClient
}

func (r *LsnappsprofilePortBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnappsprofilePortBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnappsprofile_port_binding"
}

func (r *LsnappsprofilePortBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnappsprofilePortBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnappsprofilePortBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnappsprofile_port_binding resource")
	lsnappsprofile_port_binding := lsnappsprofile_port_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lsnappsprofile_port_binding.Type(), &lsnappsprofile_port_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnappsprofile_port_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnappsprofile_port_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("appsprofilename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Appsprofilename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("lsnport:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Lsnport.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLsnappsprofilePortBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsprofilePortBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnappsprofilePortBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnappsprofile_port_binding resource")

	r.readLsnappsprofilePortBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsprofilePortBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnappsprofilePortBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnappsprofile_port_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		lsnappsprofile_port_binding := lsnappsprofile_port_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lsnappsprofile_port_binding.Type(), &lsnappsprofile_port_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnappsprofile_port_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lsnappsprofile_port_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnappsprofile_port_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLsnappsprofilePortBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsprofilePortBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnappsprofilePortBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnappsprofile_port_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"appsprofilename", "lsnport"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	appsprofilename_value, ok := idMap["appsprofilename"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'appsprofilename' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["lsnport"]; ok && val != "" {
		argsMap["lsnport"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lsnappsprofile_port_binding.Type(), appsprofilename_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnappsprofile_port_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnappsprofile_port_binding binding")
}

// Helper function to read lsnappsprofile_port_binding data from API
func (r *LsnappsprofilePortBindingResource) readLsnappsprofilePortBindingFromApi(ctx context.Context, data *LsnappsprofilePortBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"appsprofilename", "lsnport"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	appsprofilename_Name, ok := idMap["appsprofilename"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'appsprofilename' not found in ID string")
		return
	}

	// The direct lsnappsprofile_port_binding GET endpoint does not return the bound
	// ports; they are only available nested under the parent's aggregate
	// lsnappsprofile_binding endpoint. Fetch them from there.
	dataArr, err := getLsnappsprofilePortBindings(r.client, appsprofilename_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnappsprofile_port_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "lsnappsprofile_port_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check lsnport
		if idVal, ok := idMap["lsnport"]; ok {
			if val, ok := v["lsnport"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["lsnport"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("lsnappsprofile_port_binding not found with the provided ID attributes"))
		return
	}

	lsnappsprofile_port_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
