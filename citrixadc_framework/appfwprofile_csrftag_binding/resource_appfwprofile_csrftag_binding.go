package appfwprofile_csrftag_binding

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
var _ resource.Resource = &AppfwprofileCsrftagBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileCsrftagBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileCsrftagBindingResource)(nil)

func NewAppfwprofileCsrftagBindingResource() resource.Resource {
	return &AppfwprofileCsrftagBindingResource{}
}

// AppfwprofileCsrftagBindingResource defines the resource implementation.
type AppfwprofileCsrftagBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileCsrftagBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileCsrftagBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_csrftag_binding"
}

func (r *AppfwprofileCsrftagBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileCsrftagBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileCsrftagBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_csrftag_binding resource")
	appfwprofile_csrftag_binding := appfwprofile_csrftag_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_csrftag_binding.Type(), &appfwprofile_csrftag_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_csrftag_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_csrftag_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("csrfformactionurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Csrfformactionurl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("csrftag:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Csrftag.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileCsrftagBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_csrftag_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCsrftagBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileCsrftagBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_csrftag_binding resource")

	r.readAppfwprofileCsrftagBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwprofileCsrftagBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileCsrftagBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_csrftag_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_csrftag_binding := appfwprofile_csrftag_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_csrftag_binding.Type(), &appfwprofile_csrftag_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_csrftag_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_csrftag_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_csrftag_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileCsrftagBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_csrftag_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCsrftagBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileCsrftagBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_csrftag_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "csrftag", "csrfformactionurl"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// The delete-arg values must be URL-encoded because the NITRO client joins them
	// into the delete URL ?args= list verbatim (no escaping). csrfformactionurl/csrftag
	// can contain regex/URL special characters. Mirrors SDK v2 (url.QueryEscape).
	args := make([]string, 0)
	if val, ok := idMap["csrftag"]; ok && val != "" {
		args = append(args, fmt.Sprintf("csrftag:%s", url.QueryEscape(val)))
	}
	if val, ok := idMap["csrfformactionurl"]; ok && val != "" {
		args = append(args, fmt.Sprintf("csrfformactionurl:%s", url.QueryEscape(val)))
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() && data.Ruletype.ValueString() != "" {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(data.Ruletype.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Appfwprofile_csrftag_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_csrftag_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_csrftag_binding binding")
}

// Helper function to read appfwprofile_csrftag_binding data from API
func (r *AppfwprofileCsrftagBindingResource) readAppfwprofileCsrftagBindingFromApi(ctx context.Context, data *AppfwprofileCsrftagBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "csrftag", "csrfformactionurl"}, nil)
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
		ResourceType:             service.Appfwprofile_csrftag_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_csrftag_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check csrfformactionurl
		if idVal, ok := idMap["csrfformactionurl"]; ok {
			if val, ok := v["csrfformactionurl"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["csrfformactionurl"].(string); ok {
			match = false
			continue
		}

		// Check csrftag
		if idVal, ok := idMap["csrftag"]; ok {
			if val, ok := v["csrftag"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["csrftag"].(string); ok {
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
		data.Id = types.StringNull()
		return
	}

	appfwprofile_csrftag_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
