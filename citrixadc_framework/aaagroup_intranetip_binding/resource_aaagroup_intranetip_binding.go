package aaagroup_intranetip_binding

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
var _ resource.Resource = &AaagroupIntranetipBindingResource{}
var _ resource.ResourceWithConfigure = (*AaagroupIntranetipBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaagroupIntranetipBindingResource)(nil)

func NewAaagroupIntranetipBindingResource() resource.Resource {
	return &AaagroupIntranetipBindingResource{}
}

// AaagroupIntranetipBindingResource defines the resource implementation.
type AaagroupIntranetipBindingResource struct {
	client *service.NitroClient
}

func (r *AaagroupIntranetipBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaagroupIntranetipBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup_intranetip_binding"
}

func (r *AaagroupIntranetipBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaagroupIntranetipBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaagroupIntranetipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaagroup_intranetip_binding resource")
	aaagroup_intranetip_binding := aaagroup_intranetip_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaagroup_intranetip_binding.Type(), &aaagroup_intranetip_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaagroup_intranetip_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaagroup_intranetip_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("intranetip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Intranetip.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAaagroupIntranetipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupIntranetipBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaagroupIntranetipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaagroup_intranetip_binding resource")

	r.readAaagroupIntranetipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupIntranetipBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AaagroupIntranetipBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating aaagroup_intranetip_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		aaagroup_intranetip_binding := aaagroup_intranetip_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Aaagroup_intranetip_binding.Type(), &aaagroup_intranetip_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaagroup_intranetip_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaagroup_intranetip_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaagroup_intranetip_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAaagroupIntranetipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupIntranetipBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaagroupIntranetipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaagroup_intranetip_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "intranetip"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	groupname_value, ok := idMap["groupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'groupname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["intranetip"]; ok && val != "" {
		argsMap["intranetip"] = val
	}
	// netmask is a required delete argument for this binding (NITRO errorcode 1095
	// otherwise). It is not part of the ID, so read it from prior state.
	if !data.Netmask.IsNull() && data.Netmask.ValueString() != "" {
		argsMap["netmask"] = data.Netmask.ValueString()
	}

	err = r.client.DeleteResourceWithArgsMap(service.Aaagroup_intranetip_binding.Type(), groupname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaagroup_intranetip_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaagroup_intranetip_binding binding")
}

// Helper function to read aaagroup_intranetip_binding data from API
func (r *AaagroupIntranetipBindingResource) readAaagroupIntranetipBindingFromApi(ctx context.Context, data *AaagroupIntranetipBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "intranetip"}, nil)
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
		ResourceType:             service.Aaagroup_intranetip_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup_intranetip_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "aaagroup_intranetip_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check intranetip
		if idVal, ok := idMap["intranetip"]; ok {
			if val, ok := v["intranetip"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["intranetip"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("aaagroup_intranetip_binding not found with the provided ID attributes"))
		return
	}

	aaagroup_intranetip_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
