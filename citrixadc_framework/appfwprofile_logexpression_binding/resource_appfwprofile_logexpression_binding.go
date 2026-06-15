package appfwprofile_logexpression_binding

import (
	"context"
	"fmt"
	"net/url"
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
var _ resource.Resource = &AppfwprofileLogexpressionBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileLogexpressionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileLogexpressionBindingResource)(nil)

func NewAppfwprofileLogexpressionBindingResource() resource.Resource {
	return &AppfwprofileLogexpressionBindingResource{}
}

// AppfwprofileLogexpressionBindingResource defines the resource implementation.
type AppfwprofileLogexpressionBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileLogexpressionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileLogexpressionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_logexpression_binding"
}

func (r *AppfwprofileLogexpressionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileLogexpressionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileLogexpressionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_logexpression_binding resource")
	appfwprofile_logexpression_binding := appfwprofile_logexpression_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_logexpression_binding.Type(), &appfwprofile_logexpression_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_logexpression_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_logexpression_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("logexpression:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Logexpression.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileLogexpressionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileLogexpressionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileLogexpressionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_logexpression_binding resource")

	r.readAppfwprofileLogexpressionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileLogexpressionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileLogexpressionBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_logexpression_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_logexpression_binding := appfwprofile_logexpression_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_logexpression_binding.Type(), &appfwprofile_logexpression_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_logexpression_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_logexpression_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_logexpression_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileLogexpressionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileLogexpressionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileLogexpressionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_logexpression_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "logexpression"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// DeleteResourceWithArgsMap does NOT URL-encode the arg values, so values that
	// contain special characters must be encoded here, mirroring the SDK v2 resource
	// (which url.QueryEscape'd them).
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["logexpression"]; ok && val != "" {
		argsMap["logexpression"] = url.QueryEscape(val)
	}
	// ruletype is not part of the composite ID; pull it from state and include it in
	// the delete args (mirrors SDK v2 + the NITRO delete args list).
	if !data.Ruletype.IsNull() && data.Ruletype.ValueString() != "" {
		argsMap["ruletype"] = url.QueryEscape(data.Ruletype.ValueString())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_logexpression_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_logexpression_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_logexpression_binding binding")
}

// Helper function to read appfwprofile_logexpression_binding data from API
func (r *AppfwprofileLogexpressionBindingResource) readAppfwprofileLogexpressionBindingFromApi(ctx context.Context, data *AppfwprofileLogexpressionBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "logexpression"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_logexpression_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_logexpression_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwprofile_logexpression_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check logexpression
		if idVal, ok := idMap["logexpression"]; ok {
			if val, ok := v["logexpression"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["logexpression"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("appfwprofile_logexpression_binding not found with the provided ID attributes"))
		return
	}

	appfwprofile_logexpression_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
