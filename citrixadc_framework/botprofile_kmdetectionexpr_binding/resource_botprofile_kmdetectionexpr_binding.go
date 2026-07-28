package botprofile_kmdetectionexpr_binding

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
var _ resource.Resource = &BotprofileKmdetectionexprBindingResource{}
var _ resource.ResourceWithConfigure = (*BotprofileKmdetectionexprBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BotprofileKmdetectionexprBindingResource)(nil)

func NewBotprofileKmdetectionexprBindingResource() resource.Resource {
	return &BotprofileKmdetectionexprBindingResource{}
}

// BotprofileKmdetectionexprBindingResource defines the resource implementation.
type BotprofileKmdetectionexprBindingResource struct {
	client *service.NitroClient
}

func (r *BotprofileKmdetectionexprBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotprofileKmdetectionexprBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile_kmdetectionexpr_binding"
}

func (r *BotprofileKmdetectionexprBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotprofileKmdetectionexprBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotprofileKmdetectionexprBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botprofile_kmdetectionexpr_binding resource")
	botprofile_kmdetectionexpr_binding := botprofile_kmdetectionexpr_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Botprofile_kmdetectionexpr_binding.Type(), &botprofile_kmdetectionexpr_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botprofile_kmdetectionexpr_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created botprofile_kmdetectionexpr_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bot_km_expression_name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BotKmExpressionName.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("kmdetectionexpr:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Kmdetectionexpr.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readBotprofileKmdetectionexprBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileKmdetectionexprBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotprofileKmdetectionexprBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botprofile_kmdetectionexpr_binding resource")

	r.readBotprofileKmdetectionexprBindingFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Object gone out-of-band - remove from state so a subsequent apply re-creates it
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileKmdetectionexprBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state BotprofileKmdetectionexprBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for botprofile_kmdetectionexpr_binding; the NITRO resource has no
	// set/update endpoint and every attribute is RequiresReplace (immutable binding).
	tflog.Debug(ctx, "Update is a no-op for botprofile_kmdetectionexpr_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readBotprofileKmdetectionexprBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileKmdetectionexprBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotprofileKmdetectionexprBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botprofile_kmdetectionexpr_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
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
	if val, ok := idMap["bot_km_expression_name"]; ok && val != "" {
		argsMap["bot_km_expression_name"] = val
	}
	if val, ok := idMap["kmdetectionexpr"]; ok && val != "" {
		argsMap["kmdetectionexpr"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Botprofile_kmdetectionexpr_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete botprofile_kmdetectionexpr_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted botprofile_kmdetectionexpr_binding binding")
}

// Helper function to read botprofile_kmdetectionexpr_binding data from API
func (r *BotprofileKmdetectionexprBindingResource) readBotprofileKmdetectionexprBindingFromApi(ctx context.Context, data *BotprofileKmdetectionexprBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
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
		ResourceType:             service.Botprofile_kmdetectionexpr_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botprofile_kmdetectionexpr_binding, got error: %s", err))
		return
	}

	// Resource is missing (parent gone) - signal removal via null Id
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bot_km_expression_name
		if idVal, ok := idMap["bot_km_expression_name"]; ok {
			if val, ok := v["bot_km_expression_name"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["bot_km_expression_name"].(string); ok {
			match = false
			continue
		}

		// Check kmdetectionexpr
		if idVal, ok := idMap["kmdetectionexpr"]; ok {
			if val, ok := v["kmdetectionexpr"].(bool); ok {
				idValBool, _ := strconv.ParseBool(idVal)
				if val != idValBool {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["kmdetectionexpr"].(bool); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing - signal removal via null Id
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	botprofile_kmdetectionexpr_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
