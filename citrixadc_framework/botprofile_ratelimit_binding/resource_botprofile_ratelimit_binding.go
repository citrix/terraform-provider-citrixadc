package botprofile_ratelimit_binding

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

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bot_rate_limit_type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotRateLimitType.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("bot_rate_limit_url:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotRateLimitUrl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("bot_ratelimit:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotRatelimit.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("condition:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Condition.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("cookiename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Cookiename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("countrycode:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Countrycode.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
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
	// Binding with parent - delete using DeleteResourceWithArgs
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

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["bot_rate_limit_type"]; ok && val != "" {
		argsMap["bot_rate_limit_type"] = val
	}
	if val, ok := idMap["bot_rate_limit_url"]; ok && val != "" {
		argsMap["bot_rate_limit_url"] = val
	}
	if val, ok := idMap["bot_ratelimit"]; ok && val != "" {
		argsMap["bot_ratelimit"] = val
	}
	if val, ok := idMap["condition"]; ok && val != "" {
		argsMap["condition"] = val
	}
	if val, ok := idMap["cookiename"]; ok && val != "" {
		argsMap["cookiename"] = val
	}
	if val, ok := idMap["countrycode"]; ok && val != "" {
		argsMap["countrycode"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Botprofile_ratelimit_binding.Type(), name_value, argsMap)
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

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bot_rate_limit_type
		if idVal, ok := idMap["bot_rate_limit_type"]; ok {
			if val, ok := v["bot_rate_limit_type"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["bot_rate_limit_type"].(string); ok {
			match = false
			continue
		}

		// Check bot_rate_limit_url
		if idVal, ok := idMap["bot_rate_limit_url"]; ok {
			if val, ok := v["bot_rate_limit_url"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["bot_rate_limit_url"].(string); ok {
			match = false
			continue
		}

		// Check bot_ratelimit
		if idVal, ok := idMap["bot_ratelimit"]; ok {
			if val, ok := v["bot_ratelimit"].(bool); ok {
				idValBool, _ := strconv.ParseBool(idVal)
				if val != idValBool {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["bot_ratelimit"].(bool); ok {
			match = false
			continue
		}

		// Check condition
		if idVal, ok := idMap["condition"]; ok {
			if val, ok := v["condition"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["condition"].(string); ok {
			match = false
			continue
		}

		// Check cookiename
		if idVal, ok := idMap["cookiename"]; ok {
			if val, ok := v["cookiename"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["cookiename"].(string); ok {
			match = false
			continue
		}

		// Check countrycode
		if idVal, ok := idMap["countrycode"]; ok {
			if val, ok := v["countrycode"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["countrycode"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("botprofile_ratelimit_binding not found with the provided ID attributes"))
		return
	}

	botprofile_ratelimit_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
