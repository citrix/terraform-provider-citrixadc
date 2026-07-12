package botprofile_whitelist_binding

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
var _ resource.Resource = &BotprofileWhitelistBindingResource{}
var _ resource.ResourceWithConfigure = (*BotprofileWhitelistBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotprofileWhitelistBindingResource)(nil)

func NewBotprofileWhitelistBindingResource() resource.Resource {
	return &BotprofileWhitelistBindingResource{}
}

// BotprofileWhitelistBindingResource defines the resource implementation.
type BotprofileWhitelistBindingResource struct {
	client *service.NitroClient
}

func (r *BotprofileWhitelistBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotprofileWhitelistBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_whitelist_binding"
}

func (r *BotprofileWhitelistBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotprofileWhitelistBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotprofileWhitelistBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botprofile_whitelist_binding resource")
	botprofile_whitelist_binding := botprofile_whitelist_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Botprofile_whitelist_binding.Type(), &botprofile_whitelist_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botprofile_whitelist_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created botprofile_whitelist_binding resource")

	// Set ID for the resource before reading state.
	// Use the SDK v2 legacy attribute order "name,bot_whitelist_value"
	// (see resource_id_mapping.json) so the new key:value ID and any imported
	// legacy positional ID decode identically via utils.ParseIdString.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("bot_whitelist_value:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotWhitelistValue.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readBotprofileWhitelistBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "botprofile_whitelist_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileWhitelistBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotprofileWhitelistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botprofile_whitelist_binding resource")

	r.readBotprofileWhitelistBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *BotprofileWhitelistBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state BotprofileWhitelistBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating botprofile_whitelist_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		botprofile_whitelist_binding := botprofile_whitelist_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Botprofile_whitelist_binding.Type(), &botprofile_whitelist_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botprofile_whitelist_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated botprofile_whitelist_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for botprofile_whitelist_binding resource, skipping update")
	}

	// Read the updated state back
	r.readBotprofileWhitelistBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "botprofile_whitelist_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileWhitelistBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotprofileWhitelistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botprofile_whitelist_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Parse name + bot_whitelist_value from the ID (handles new key:value and
	// legacy positional "name,bot_whitelist_value" formats).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "bot_whitelist_value"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	botWhitelistValue, ok := idMap["bot_whitelist_value"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Attribute 'bot_whitelist_value' not found in ID")
		return
	}

	// NITRO requires BOTH bot_whitelist and bot_whitelist_value as delete args
	// (errorcode 1093 "Argument pre-requisite missing [value, whiteList]"
	// otherwise). bot_whitelist comes from prior state (it is not in the ID),
	// mirroring the SDK v2 behavior. URL-encode values for slashy/special
	// whitelist entries.
	args := make([]string, 0, 2)
	if !data.BotWhitelist.IsNull() && !data.BotWhitelist.IsUnknown() {
		args = append(args, fmt.Sprintf("bot_whitelist:%t", data.BotWhitelist.ValueBool()))
	}
	args = append(args, fmt.Sprintf("bot_whitelist_value:%s", url.QueryEscape(botWhitelistValue)))

	err = r.client.DeleteResourceWithArgs(service.Botprofile_whitelist_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete botprofile_whitelist_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted botprofile_whitelist_binding binding")
}

// Helper function to read botprofile_whitelist_binding data from API
func (r *BotprofileWhitelistBindingResource) readBotprofileWhitelistBindingFromApi(ctx context.Context, data *BotprofileWhitelistBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "bot_whitelist_value"}, nil)
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
		ResourceType:             service.Botprofile_whitelist_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_whitelist_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// The binding under a profile is uniquely identified by bot_whitelist_value
	// (matching the SDK v2 read, which keyed on bot_whitelist_value only).
	botWhitelistValue, ok := idMap["bot_whitelist_value"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'bot_whitelist_value' not found in ID string")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["bot_whitelist_value"].(string); ok && val == botWhitelistValue {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		// Binding not present in the returned set: signal removal via a null Id (see above).
		data.Id = types.StringNull()
		return
	}

	botprofile_whitelist_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
