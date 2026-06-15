package botprofile_whitelist_binding

import (
	"context"
	"fmt"
	"strconv"
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

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bot_whitelist:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotWhitelist.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("bot_whitelist_value:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotWhitelistValue.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readBotprofileWhitelistBindingFromApi(ctx, &data, &resp.Diagnostics)

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
	// Binding with parent - delete using DeleteResourceWithArgs
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

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["bot_whitelist"]; ok && val != "" {
		argsMap["bot_whitelist"] = val
	}
	if val, ok := idMap["bot_whitelist_value"]; ok && val != "" {
		argsMap["bot_whitelist_value"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Botprofile_whitelist_binding.Type(), name_value, argsMap)
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
		diags.AddError("Client Error", "botprofile_whitelist_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bot_whitelist
		if idVal, ok := idMap["bot_whitelist"]; ok {
			if val, ok := v["bot_whitelist"].(bool); ok {
				idValBool, _ := strconv.ParseBool(idVal)
				if val != idValBool {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["bot_whitelist"].(bool); ok {
			match = false
			continue
		}

		// Check bot_whitelist_value
		if idVal, ok := idMap["bot_whitelist_value"]; ok {
			if val, ok := v["bot_whitelist_value"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["bot_whitelist_value"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("botprofile_whitelist_binding not found with the provided ID attributes"))
		return
	}

	botprofile_whitelist_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
