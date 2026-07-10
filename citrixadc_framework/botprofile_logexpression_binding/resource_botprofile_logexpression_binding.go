package botprofile_logexpression_binding

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
var _ resource.Resource = &BotprofileLogexpressionBindingResource{}
var _ resource.ResourceWithConfigure = (*BotprofileLogexpressionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotprofileLogexpressionBindingResource)(nil)

func NewBotprofileLogexpressionBindingResource() resource.Resource {
	return &BotprofileLogexpressionBindingResource{}
}

// BotprofileLogexpressionBindingResource defines the resource implementation.
type BotprofileLogexpressionBindingResource struct {
	client *service.NitroClient
}

func (r *BotprofileLogexpressionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotprofileLogexpressionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_logexpression_binding"
}

func (r *BotprofileLogexpressionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotprofileLogexpressionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotprofileLogexpressionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botprofile_logexpression_binding resource")
	botprofile_logexpression_binding := botprofile_logexpression_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Botprofile_logexpression_binding.Type(), &botprofile_logexpression_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botprofile_logexpression_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created botprofile_logexpression_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bot_log_expression_name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotLogExpressionName.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("logexpression:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Logexpression.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readBotprofileLogexpressionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileLogexpressionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotprofileLogexpressionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botprofile_logexpression_binding resource")

	r.readBotprofileLogexpressionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileLogexpressionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state BotprofileLogexpressionBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating botprofile_logexpression_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		botprofile_logexpression_binding := botprofile_logexpression_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Botprofile_logexpression_binding.Type(), &botprofile_logexpression_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botprofile_logexpression_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated botprofile_logexpression_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for botprofile_logexpression_binding resource, skipping update")
	}

	// Read the updated state back
	r.readBotprofileLogexpressionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileLogexpressionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotprofileLogexpressionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botprofile_logexpression_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "bot_log_expression_name"}, nil)
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
	if val, ok := idMap["bot_log_expression_name"]; ok && val != "" {
		argsMap["bot_log_expression_name"] = val
	}
	if val, ok := idMap["logexpression"]; ok && val != "" {
		argsMap["logexpression"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Botprofile_logexpression_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete botprofile_logexpression_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted botprofile_logexpression_binding binding")
}

// Helper function to read botprofile_logexpression_binding data from API
func (r *BotprofileLogexpressionBindingResource) readBotprofileLogexpressionBindingFromApi(ctx context.Context, data *BotprofileLogexpressionBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "bot_log_expression_name"}, nil)
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
		ResourceType:             service.Botprofile_logexpression_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_logexpression_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "botprofile_logexpression_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bot_log_expression_name
		if idVal, ok := idMap["bot_log_expression_name"]; ok {
			if val, ok := v["bot_log_expression_name"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["bot_log_expression_name"].(string); ok {
			match = false
			continue
		}

		// Check logexpression
		if idVal, ok := idMap["logexpression"]; ok {
			if val, ok := v["logexpression"].(bool); ok {
				idValBool, _ := strconv.ParseBool(idVal)
				if val != idValBool {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["logexpression"].(bool); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("botprofile_logexpression_binding not found with the provided ID attributes"))
		return
	}

	botprofile_logexpression_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
