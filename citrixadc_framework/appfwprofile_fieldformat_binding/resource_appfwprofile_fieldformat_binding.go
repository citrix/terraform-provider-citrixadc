package appfwprofile_fieldformat_binding

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
var _ resource.Resource = &AppfwprofileFieldformatBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileFieldformatBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileFieldformatBindingResource)(nil)

func NewAppfwprofileFieldformatBindingResource() resource.Resource {
	return &AppfwprofileFieldformatBindingResource{}
}

// AppfwprofileFieldformatBindingResource defines the resource implementation.
type AppfwprofileFieldformatBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileFieldformatBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileFieldformatBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_fieldformat_binding"
}

func (r *AppfwprofileFieldformatBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileFieldformatBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileFieldformatBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_fieldformat_binding resource")
	appfwprofile_fieldformat_binding := appfwprofile_fieldformat_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_fieldformat_binding.Type(), &appfwprofile_fieldformat_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_fieldformat_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_fieldformat_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("fieldformat:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Fieldformat.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_ff:%s", utils.UrlEncode(fmt.Sprintf("%v", data.FormactionurlFf.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileFieldformatBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_fieldformat_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFieldformatBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileFieldformatBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_fieldformat_binding resource")

	r.readAppfwprofileFieldformatBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwprofileFieldformatBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileFieldformatBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_fieldformat_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_fieldformat_binding := appfwprofile_fieldformat_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_fieldformat_binding.Type(), &appfwprofile_fieldformat_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_fieldformat_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_fieldformat_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_fieldformat_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileFieldformatBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_fieldformat_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFieldformatBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileFieldformatBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_fieldformat_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "fieldformat", "formactionurl_ff"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// Mirror SDK v2: URL-encode special-char values (formactionurl_ff is a
	// regex/URL, ruletype may contain special chars) in the delete args.
	args := make([]string, 0)
	if val, ok := idMap["fieldformat"]; ok && val != "" {
		args = append(args, fmt.Sprintf("fieldformat:%s", val))
	}
	if val, ok := idMap["formactionurl_ff"]; ok && val != "" {
		args = append(args, fmt.Sprintf("formactionurl_ff:%s", url.QueryEscape(val)))
	}
	if !data.Ruletype.IsNull() && data.Ruletype.ValueString() != "" {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(data.Ruletype.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Appfwprofile_fieldformat_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_fieldformat_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_fieldformat_binding binding")
}

// Helper function to read appfwprofile_fieldformat_binding data from API
func (r *AppfwprofileFieldformatBindingResource) readAppfwprofileFieldformatBindingFromApi(ctx context.Context, data *AppfwprofileFieldformatBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "fieldformat", "formactionurl_ff"}, nil)
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
		ResourceType:             service.Appfwprofile_fieldformat_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_fieldformat_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check fieldformat
		if idVal, ok := idMap["fieldformat"]; ok {
			if val, ok := v["fieldformat"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["fieldformat"].(string); ok {
			match = false
			continue
		}

		// Check formactionurl_ff
		if idVal, ok := idMap["formactionurl_ff"]; ok {
			if val, ok := v["formactionurl_ff"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["formactionurl_ff"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	appfwprofile_fieldformat_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
