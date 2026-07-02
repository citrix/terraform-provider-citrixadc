package tmglobal_tmsessionpolicy_binding

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
var _ resource.Resource = &TmglobalTmsessionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*TmglobalTmsessionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*TmglobalTmsessionpolicyBindingResource)(nil)

func NewTmglobalTmsessionpolicyBindingResource() resource.Resource {
	return &TmglobalTmsessionpolicyBindingResource{}
}

// TmglobalTmsessionpolicyBindingResource defines the resource implementation.
type TmglobalTmsessionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *TmglobalTmsessionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmglobalTmsessionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmglobal_tmsessionpolicy_binding"
}

func (r *TmglobalTmsessionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmglobalTmsessionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmglobalTmsessionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmglobal_tmsessionpolicy_binding resource")
	tmglobal_tmsessionpolicy_binding := tmglobal_tmsessionpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Tmglobal_tmsessionpolicy_binding.Type(), &tmglobal_tmsessionpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmglobal_tmsessionpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created tmglobal_tmsessionpolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	// Read the updated state back
	r.readTmglobalTmsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmglobalTmsessionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmglobalTmsessionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmglobal_tmsessionpolicy_binding resource")

	r.readTmglobalTmsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmglobalTmsessionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state TmglobalTmsessionpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// NITRO exposes no set/update verb for tmglobal_tmsessionpolicy_binding.
	// All configurable attributes are RequiresReplace, so Update is a documented no-op.
	tflog.Debug(ctx, "Update is a no-op for tmglobal_tmsessionpolicy_binding; NITRO has no set endpoint and all attributes are RequiresReplace")

	// Read the current state back
	r.readTmglobalTmsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmglobalTmsessionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmglobalTmsessionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmglobal_tmsessionpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	policyname_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("policyname:%s", policyname_value),
	}

	err := r.client.DeleteResourceWithArgs(service.Tmglobal_tmsessionpolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete tmglobal_tmsessionpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted tmglobal_tmsessionpolicy_binding binding")
}

// Helper function to read tmglobal_tmsessionpolicy_binding data from API
func (r *TmglobalTmsessionpolicyBindingResource) readTmglobalTmsessionpolicyBindingFromApi(ctx context.Context, data *TmglobalTmsessionpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID.
	// The ID is a single plain value (policyname), so supply the legacy attr order
	// so ParseIdString can map it (otherwise it errors "no attribute order provided").
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	// NOTE (global-binding read quirk): on this firmware, neither the typed
	// GET tmglobal_tmsessionpolicy_binding nor the umbrella GET tmglobal_binding
	// return the bound policies (both come back empty {"message":"Done"} / [{}]).
	// The bound policies are only exposed by the base "tmglobal" endpoint, which
	// returns an array of {policyname, priority, gotopriorityexpression, feature, ...}.
	// We therefore read from "tmglobal" and match on policyname.
	findParams := service.FindParams{
		ResourceType:             "tmglobal",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmglobal_tmsessionpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "tmglobal_tmsessionpolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if idVal, ok := idMap["policyname"]; ok {
			if val, ok := v["policyname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["policyname"].(string); ok {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("tmglobal_tmsessionpolicy_binding not found with the provided ID attributes"))
		return
	}

	tmglobal_tmsessionpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
