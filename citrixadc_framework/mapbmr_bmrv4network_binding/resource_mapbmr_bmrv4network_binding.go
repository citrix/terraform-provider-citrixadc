package mapbmr_bmrv4network_binding

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
var _ resource.Resource = &MapbmrBmrv4networkBindingResource{}
var _ resource.ResourceWithConfigure = (*MapbmrBmrv4networkBindingResource)(nil)
var _ resource.ResourceWithImportState = (*MapbmrBmrv4networkBindingResource)(nil)

func NewMapbmrBmrv4networkBindingResource() resource.Resource {
	return &MapbmrBmrv4networkBindingResource{}
}

// MapbmrBmrv4networkBindingResource defines the resource implementation.
type MapbmrBmrv4networkBindingResource struct {
	client *service.NitroClient
}

func (r *MapbmrBmrv4networkBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MapbmrBmrv4networkBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mapbmr_bmrv4network_binding"
}

func (r *MapbmrBmrv4networkBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *MapbmrBmrv4networkBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MapbmrBmrv4networkBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating mapbmr_bmrv4network_binding resource")
	mapbmr_bmrv4network_binding := mapbmr_bmrv4network_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Mapbmr_bmrv4network_binding.Type(), &mapbmr_bmrv4network_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create mapbmr_bmrv4network_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created mapbmr_bmrv4network_binding resource")

	// Set ID for the resource before reading state
	// Composite ID matches legacy SDK v2 order (resource_id_mapping.json: "name,network").
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("network:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Network.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readMapbmrBmrv4networkBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapbmrBmrv4networkBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MapbmrBmrv4networkBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading mapbmr_bmrv4network_binding resource")

	r.readMapbmrBmrv4networkBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapbmrBmrv4networkBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state MapbmrBmrv4networkBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating mapbmr_bmrv4network_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		mapbmr_bmrv4network_binding := mapbmr_bmrv4network_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Mapbmr_bmrv4network_binding.Type(), &mapbmr_bmrv4network_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update mapbmr_bmrv4network_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated mapbmr_bmrv4network_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for mapbmr_bmrv4network_binding resource, skipping update")
	}

	// Read the updated state back
	r.readMapbmrBmrv4networkBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapbmrBmrv4networkBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MapbmrBmrv4networkBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting mapbmr_bmrv4network_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "network"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// network is the mandatory delete arg (red/bold in NITRO doc); netmask is optional.
	// netmask is not part of the ID, so read it from state. URL-encode arg values since
	// the NITRO client joins them raw into the delete URL.
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["network"]; ok && val != "" {
		argsMap["network"] = url.QueryEscape(val)
	}
	if !data.Netmask.IsNull() && data.Netmask.ValueString() != "" {
		argsMap["netmask"] = url.QueryEscape(data.Netmask.ValueString())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Mapbmr_bmrv4network_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete mapbmr_bmrv4network_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted mapbmr_bmrv4network_binding binding")
}

// Helper function to read mapbmr_bmrv4network_binding data from API
func (r *MapbmrBmrv4networkBindingResource) readMapbmrBmrv4networkBindingFromApi(ctx context.Context, data *MapbmrBmrv4networkBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "network"}, nil)
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
		ResourceType:             service.Mapbmr_bmrv4network_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read mapbmr_bmrv4network_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "mapbmr_bmrv4network_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id.
	// Identity is (name, network); netmask is not part of the ID.
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check network
		if idVal, ok := idMap["network"]; ok {
			if val, ok := v["network"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["network"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("mapbmr_bmrv4network_binding not found with the provided ID attributes"))
		return
	}

	mapbmr_bmrv4network_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
