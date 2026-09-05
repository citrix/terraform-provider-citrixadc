package appfwprofile_trustedlearningclients_binding

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
var _ resource.Resource = &AppfwprofileTrustedlearningclientsBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileTrustedlearningclientsBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileTrustedlearningclientsBindingResource)(nil)

func NewAppfwprofileTrustedlearningclientsBindingResource() resource.Resource {
	return &AppfwprofileTrustedlearningclientsBindingResource{}
}

// AppfwprofileTrustedlearningclientsBindingResource defines the resource implementation.
type AppfwprofileTrustedlearningclientsBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_trustedlearningclients_binding"
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileTrustedlearningclientsBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_trustedlearningclients_binding resource")
	appfwprofile_trustedlearningclients_binding := appfwprofile_trustedlearningclients_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_trustedlearningclients_binding.Type(), &appfwprofile_trustedlearningclients_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_trustedlearningclients_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_trustedlearningclients_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("trustedlearningclients:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Trustedlearningclients.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileTrustedlearningclientsBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_trustedlearningclients_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileTrustedlearningclientsBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_trustedlearningclients_binding resource")

	r.readAppfwprofileTrustedlearningclientsBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwprofileTrustedlearningclientsBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileTrustedlearningclientsBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_trustedlearningclients_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_trustedlearningclients_binding := appfwprofile_trustedlearningclients_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_trustedlearningclients_binding.Type(), &appfwprofile_trustedlearningclients_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_trustedlearningclients_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_trustedlearningclients_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_trustedlearningclients_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileTrustedlearningclientsBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_trustedlearningclients_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileTrustedlearningclientsBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_trustedlearningclients_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "trustedlearningclients"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// Mirror SDK v2: build the delete args slice with url.QueryEscape-encoded
	// values (family pattern b). The NITRO client's *WithArgsMap helpers do NOT
	// URL-encode arg values, so an unencoded '/' in trustedlearningclients (e.g.
	// "1.2.31.1/32") breaks the existence pre-check and the delete is silently
	// skipped. DeleteResourceWithArgs takes the args verbatim, so we encode here.
	args := make([]string, 0)
	if val, ok := idMap["trustedlearningclients"]; ok && val != "" {
		args = append(args, fmt.Sprintf("trustedlearningclients:%s", url.QueryEscape(val)))
	}
	if !data.Ruletype.IsNull() && data.Ruletype.ValueString() != "" {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(data.Ruletype.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Appfwprofile_trustedlearningclients_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_trustedlearningclients_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_trustedlearningclients_binding binding")
}

// Helper function to read appfwprofile_trustedlearningclients_binding data from API
func (r *AppfwprofileTrustedlearningclientsBindingResource) readAppfwprofileTrustedlearningclientsBindingFromApi(ctx context.Context, data *AppfwprofileTrustedlearningclientsBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "trustedlearningclients"}, nil)
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
		ResourceType:             service.Appfwprofile_trustedlearningclients_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_trustedlearningclients_binding, got error: %s", err))
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

		// Check trustedlearningclients
		if idVal, ok := idMap["trustedlearningclients"]; ok {
			if val, ok := v["trustedlearningclients"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["trustedlearningclients"].(string); ok {
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

	appfwprofile_trustedlearningclients_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
