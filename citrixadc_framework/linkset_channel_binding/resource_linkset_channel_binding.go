package linkset_channel_binding

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
var _ resource.Resource = &LinksetChannelBindingResource{}
var _ resource.ResourceWithConfigure = (*LinksetChannelBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LinksetChannelBindingResource)(nil)

func NewLinksetChannelBindingResource() resource.Resource {
	return &LinksetChannelBindingResource{}
}

// LinksetChannelBindingResource defines the resource implementation.
type LinksetChannelBindingResource struct {
	client *service.NitroClient
}

func (r *LinksetChannelBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LinksetChannelBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_linkset_channel_binding"
}

func (r *LinksetChannelBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LinksetChannelBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LinksetChannelBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating linkset_channel_binding resource")
	linkset_channel_binding := linkset_channel_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Linkset_channel_binding.Type(), &linkset_channel_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create linkset_channel_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created linkset_channel_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("linkset_id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.LinksetId.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLinksetChannelBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LinksetChannelBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LinksetChannelBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading linkset_channel_binding resource")

	r.readLinksetChannelBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LinksetChannelBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LinksetChannelBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating linkset_channel_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		linkset_channel_binding := linkset_channel_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Linkset_channel_binding.Type(), &linkset_channel_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update linkset_channel_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated linkset_channel_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for linkset_channel_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLinksetChannelBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LinksetChannelBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LinksetChannelBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting linkset_channel_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgsMap.
	// The user-facing "linkset_id" attribute maps to the NITRO resource name.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"linkset_id", "ifnum"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	linkset_id_value, ok := idMap["linkset_id"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'linkset_id' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["ifnum"]; ok && val != "" {
		argsMap["ifnum"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Linkset_channel_binding.Type(), linkset_id_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete linkset_channel_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted linkset_channel_binding binding")
}

// Helper function to read linkset_channel_binding data from API
func (r *LinksetChannelBindingResource) readLinksetChannelBindingFromApi(ctx context.Context, data *LinksetChannelBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"linkset_id", "ifnum"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	linkset_id_Name, ok := idMap["linkset_id"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'linkset_id' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	// The linkset id (e.g. "LS/3") contains a '/', so it must be double
	// URL-encoded for the NITRO GET path (matches SDK v2 behavior).
	findParams := service.FindParams{
		ResourceType:             service.Linkset_channel_binding.Type(),
		ResourceName:             url.QueryEscape(url.QueryEscape(linkset_id_Name)),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read linkset_channel_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "linkset_channel_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ifnum
		if idVal, ok := idMap["ifnum"]; ok {
			if val, ok := v["ifnum"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ifnum"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("linkset_channel_binding not found with the provided ID attributes"))
		return
	}

	linkset_channel_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
