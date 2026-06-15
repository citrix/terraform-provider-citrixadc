package botprofile_blacklist_binding

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
var _ resource.Resource = &BotprofileBlacklistBindingResource{}
var _ resource.ResourceWithConfigure = (*BotprofileBlacklistBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotprofileBlacklistBindingResource)(nil)

func NewBotprofileBlacklistBindingResource() resource.Resource {
	return &BotprofileBlacklistBindingResource{}
}

// BotprofileBlacklistBindingResource defines the resource implementation.
type BotprofileBlacklistBindingResource struct {
	client *service.NitroClient
}

func (r *BotprofileBlacklistBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotprofileBlacklistBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_blacklist_binding"
}

func (r *BotprofileBlacklistBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotprofileBlacklistBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotprofileBlacklistBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botprofile_blacklist_binding resource")
	botprofile_blacklist_binding := botprofile_blacklist_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Botprofile_blacklist_binding.Type(), &botprofile_blacklist_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botprofile_blacklist_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created botprofile_blacklist_binding resource")

	// Set ID for the resource before reading state.
	// ID order matches resource_id_mapping.json ("name,bot_blacklist_value") so
	// legacy SDK v2 state imports parse correctly via ParseIdString.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("bot_blacklist_value:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotBlacklistValue.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readBotprofileBlacklistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileBlacklistBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotprofileBlacklistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botprofile_blacklist_binding resource")

	r.readBotprofileBlacklistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileBlacklistBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state BotprofileBlacklistBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating botprofile_blacklist_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		botprofile_blacklist_binding := botprofile_blacklist_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Botprofile_blacklist_binding.Type(), &botprofile_blacklist_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botprofile_blacklist_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated botprofile_blacklist_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for botprofile_blacklist_binding resource, skipping update")
	}

	// Read the updated state back
	r.readBotprofileBlacklistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileBlacklistBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotprofileBlacklistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botprofile_blacklist_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// ID carries the parent name + bot_blacklist_value (legacy order); bot_blacklist
	// is read from state because the NITRO delete endpoint requires it as a delete arg
	// (args=bot_blacklist:<bool>,bot_blacklist_value:<string>), mirroring SDK v2.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "bot_blacklist_value"}, nil)
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
	if val, ok := idMap["bot_blacklist_value"]; ok && val != "" {
		args = append(args, fmt.Sprintf("bot_blacklist_value:%s", url.QueryEscape(val)))
	}
	if !data.BotBlacklist.IsNull() && !data.BotBlacklist.IsUnknown() {
		args = append(args, fmt.Sprintf("bot_blacklist:%t", data.BotBlacklist.ValueBool()))
	}

	err = r.client.DeleteResourceWithArgs(service.Botprofile_blacklist_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete botprofile_blacklist_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted botprofile_blacklist_binding binding")
}

// Helper function to read botprofile_blacklist_binding data from API
func (r *BotprofileBlacklistBindingResource) readBotprofileBlacklistBindingFromApi(ctx context.Context, data *BotprofileBlacklistBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "bot_blacklist_value"}, nil)
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
		ResourceType:             service.Botprofile_blacklist_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_blacklist_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "botprofile_blacklist_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id.
	// The binding is uniquely identified within the parent profile by
	// bot_blacklist_value (the legacy SDK v2 read matched on this same key).
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bot_blacklist_value
		if idVal, ok := idMap["bot_blacklist_value"]; ok {
			if val, ok := v["bot_blacklist_value"].(string); ok {
				if val != idVal {
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

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("botprofile_blacklist_binding not found with the provided ID attributes"))
		return
	}

	botprofile_blacklist_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
