package botprofile_ratelimit_binding

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
var _ resource.Resource = &BotprofileRatelimitBindingResource{}
var _ resource.ResourceWithConfigure = (*BotprofileRatelimitBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotprofileRatelimitBindingResource)(nil)

func NewBotprofileRatelimitBindingResource() resource.Resource {
	return &BotprofileRatelimitBindingResource{}
}

// BotprofileRatelimitBindingResource defines the resource implementation.
type BotprofileRatelimitBindingResource struct {
	client *service.NitroClient
}

func (r *BotprofileRatelimitBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotprofileRatelimitBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_ratelimit_binding"
}

func (r *BotprofileRatelimitBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotprofileRatelimitBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotprofileRatelimitBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botprofile_ratelimit_binding resource")
	botprofile_ratelimit_binding := botprofile_ratelimit_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Botprofile_ratelimit_binding.Type(), &botprofile_ratelimit_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botprofile_ratelimit_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created botprofile_ratelimit_binding resource")

	// Set ID for the resource before reading state.
	// Composite key uses the SDK v2 legacy order: name,bot_rate_limit_type
	// (matches resource_id_mapping.json so imported legacy state still parses).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("bot_rate_limit_type:%s", utils.UrlEncode(data.BotRateLimitType.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readBotprofileRatelimitBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileRatelimitBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotprofileRatelimitBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botprofile_ratelimit_binding resource")

	r.readBotprofileRatelimitBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileRatelimitBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state BotprofileRatelimitBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating botprofile_ratelimit_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		botprofile_ratelimit_binding := botprofile_ratelimit_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Botprofile_ratelimit_binding.Type(), &botprofile_ratelimit_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botprofile_ratelimit_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated botprofile_ratelimit_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for botprofile_ratelimit_binding resource, skipping update")
	}

	// Read the updated state back
	r.readBotprofileRatelimitBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileRatelimitBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotprofileRatelimitBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botprofile_ratelimit_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Parent name + bot_rate_limit_type come from the composite ID (legacy order:
	// name,bot_rate_limit_type). ParseIdString handles both the new key:value form and
	// the legacy "name,bot_rate_limit_type" form for imported SDK v2 state.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "bot_rate_limit_type"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// Build the delete args matching the SDK v2 resource: the bot_rate_limit_type
	// discriminator plus any of the optional disambiguating args that are set in state.
	// Slashy/special values are URL-encoded.
	args := make([]string, 0)
	if val, ok := idMap["bot_rate_limit_type"]; ok && val != "" {
		args = append(args, fmt.Sprintf("bot_rate_limit_type:%s", url.QueryEscape(val)))
	}
	if !data.BotRatelimit.IsNull() && !data.BotRatelimit.IsUnknown() {
		args = append(args, fmt.Sprintf("bot_ratelimit:%t", data.BotRatelimit.ValueBool()))
	}
	if !data.BotRateLimitUrl.IsNull() && !data.BotRateLimitUrl.IsUnknown() && data.BotRateLimitUrl.ValueString() != "" {
		args = append(args, fmt.Sprintf("bot_rate_limit_url:%s", url.QueryEscape(data.BotRateLimitUrl.ValueString())))
	}
	if !data.Cookiename.IsNull() && !data.Cookiename.IsUnknown() && data.Cookiename.ValueString() != "" {
		args = append(args, fmt.Sprintf("cookiename:%s", url.QueryEscape(data.Cookiename.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Botprofile_ratelimit_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete botprofile_ratelimit_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted botprofile_ratelimit_binding binding")
}

// Helper function to read botprofile_ratelimit_binding data from API
func (r *BotprofileRatelimitBindingResource) readBotprofileRatelimitBindingFromApi(ctx context.Context, data *BotprofileRatelimitBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "bot_rate_limit_type"}, nil)
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
		ResourceType:             service.Botprofile_ratelimit_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_ratelimit_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "botprofile_ratelimit_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the matching bot_rate_limit_type.
	// The binding's identity (per SDK v2) is the parent name + bot_rate_limit_type,
	// so only the second key needs filtering within the parent's binding array.
	bot_rate_limit_type, ok := idMap["bot_rate_limit_type"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'bot_rate_limit_type' not found in ID string")
		return
	}

	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["bot_rate_limit_type"].(string); ok && val == bot_rate_limit_type {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("botprofile_ratelimit_binding not found with the provided ID attributes"))
		return
	}

	botprofile_ratelimit_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
