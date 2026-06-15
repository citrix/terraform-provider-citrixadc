package appfwprofile_fieldconsistency_binding

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
var _ resource.Resource = &AppfwprofileFieldconsistencyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileFieldconsistencyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileFieldconsistencyBindingResource)(nil)

func NewAppfwprofileFieldconsistencyBindingResource() resource.Resource {
	return &AppfwprofileFieldconsistencyBindingResource{}
}

// AppfwprofileFieldconsistencyBindingResource defines the resource implementation.
type AppfwprofileFieldconsistencyBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileFieldconsistencyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileFieldconsistencyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_fieldconsistency_binding"
}

func (r *AppfwprofileFieldconsistencyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileFieldconsistencyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileFieldconsistencyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_fieldconsistency_binding resource")
	appfwprofile_fieldconsistency_binding := appfwprofile_fieldconsistency_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_fieldconsistency_binding.Type(), &appfwprofile_fieldconsistency_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_fieldconsistency_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_fieldconsistency_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("fieldconsistency:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Fieldconsistency.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_ffc:%s", utils.UrlEncode(fmt.Sprintf("%v", data.FormactionurlFfc.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileFieldconsistencyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFieldconsistencyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileFieldconsistencyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_fieldconsistency_binding resource")

	r.readAppfwprofileFieldconsistencyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFieldconsistencyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileFieldconsistencyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_fieldconsistency_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_fieldconsistency_binding := appfwprofile_fieldconsistency_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_fieldconsistency_binding.Type(), &appfwprofile_fieldconsistency_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_fieldconsistency_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_fieldconsistency_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_fieldconsistency_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileFieldconsistencyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFieldconsistencyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileFieldconsistencyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_fieldconsistency_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Arg values must be URL-encoded because the NITRO client joins them into the
	// delete URL query string verbatim (it does not encode them). This mirrors the
	// SDK v2 resource which called url.QueryEscape on formactionurl_ffc and ruletype.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "fieldconsistency", "formactionurl_ffc"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["fieldconsistency"]; ok && val != "" {
		args = append(args, fmt.Sprintf("fieldconsistency:%s", url.QueryEscape(val)))
	}
	if val, ok := idMap["formactionurl_ffc"]; ok && val != "" {
		args = append(args, fmt.Sprintf("formactionurl_ffc:%s", url.QueryEscape(val)))
	}
	// ruletype is not part of the ID; read it from prior state to disambiguate the
	// binding if it was set (mirrors SDK v2 delete behavior).
	if !data.Ruletype.IsNull() && data.Ruletype.ValueString() != "" {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(data.Ruletype.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Appfwprofile_fieldconsistency_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_fieldconsistency_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_fieldconsistency_binding binding")
}

// Helper function to read appfwprofile_fieldconsistency_binding data from API
func (r *AppfwprofileFieldconsistencyBindingResource) readAppfwprofileFieldconsistencyBindingFromApi(ctx context.Context, data *AppfwprofileFieldconsistencyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "fieldconsistency", "formactionurl_ffc"}, nil)
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
		ResourceType:             service.Appfwprofile_fieldconsistency_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_fieldconsistency_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwprofile_fieldconsistency_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check fieldconsistency
		if idVal, ok := idMap["fieldconsistency"]; ok {
			if val, ok := v["fieldconsistency"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["fieldconsistency"].(string); ok {
			match = false
			continue
		}

		// Check formactionurl_ffc
		if idVal, ok := idMap["formactionurl_ffc"]; ok {
			if val, ok := v["formactionurl_ffc"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["formactionurl_ffc"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("appfwprofile_fieldconsistency_binding not found with the provided ID attributes"))
		return
	}

	appfwprofile_fieldconsistency_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
