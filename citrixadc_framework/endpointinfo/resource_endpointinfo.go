package endpointinfo

import (
	"context"
	"fmt"
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
var _ resource.Resource = &EndpointinfoResource{}
var _ resource.ResourceWithConfigure = (*EndpointinfoResource)(nil)
var _ resource.ResourceWithImportState = (*EndpointinfoResource)(nil)

func NewEndpointinfoResource() resource.Resource {
	return &EndpointinfoResource{}
}

// EndpointinfoResource defines the resource implementation.
type EndpointinfoResource struct {
	client *service.NitroClient
}

func (r *EndpointinfoResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *EndpointinfoResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpointinfo"
}

func (r *EndpointinfoResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *EndpointinfoResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data EndpointinfoResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating endpointinfo resource")
	endpointinfo := endpointinfoGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// NITRO add is POST (HTTP Method: POST) - use AddResource
	_, err := r.client.AddResource(service.Endpointinfo.Type(), "", endpointinfo)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create endpointinfo, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created endpointinfo resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("endpointkind:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Endpointkind.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("endpointname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Endpointname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readEndpointinfoFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EndpointinfoResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data EndpointinfoResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading endpointinfo resource")

	r.readEndpointinfoFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Object is gone out-of-band: remove from state so a subsequent apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EndpointinfoResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state EndpointinfoResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating endpointinfo resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Endpointlabelsjson.Equal(state.Endpointlabelsjson) {
		tflog.Debug(ctx, fmt.Sprintf("endpointlabelsjson has changed for endpointinfo"))
		hasChange = true
	}
	if !data.Endpointmetadata.Equal(state.Endpointmetadata) {
		tflog.Debug(ctx, fmt.Sprintf("endpointmetadata has changed for endpointinfo"))
		hasChange = true
	}
	if !data.Endpointname.Equal(state.Endpointname) {
		tflog.Debug(ctx, fmt.Sprintf("endpointname has changed for endpointinfo"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		endpointinfo := endpointinfoGetThePayloadFromthePlan(ctx, &data)
		// NITRO update is PUT /config/endpointinfo (no name in URL - identity
		// travels in the body). UpdateUnnamedResource emits exactly that PUT.
		err := r.client.UpdateUnnamedResource(service.Endpointinfo.Type(), endpointinfo)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update endpointinfo, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated endpointinfo resource")
	} else {
		tflog.Debug(ctx, "No changes detected for endpointinfo resource, skipping update")
	}

	// Read the updated state back
	r.readEndpointinfoFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EndpointinfoResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data EndpointinfoResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting endpointinfo resource")
	// NITRO delete URL: /config/endpointinfo/<endpointkind> with
	// args=endpointname:<value>. endpointkind is the URL key, endpointname is
	// the query arg.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}
	endpointkind_value := idMap["endpointkind"]
	endpointname_value := idMap["endpointname"]
	args := []string{
		fmt.Sprintf("endpointname:%s", endpointname_value),
	}

	err = r.client.DeleteResourceWithArgs(service.Endpointinfo.Type(), endpointkind_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete endpointinfo, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted endpointinfo resource")
}

// Helper function to read endpointinfo data from API
func (r *EndpointinfoResource) readEndpointinfoFromApi(ctx context.Context, data *EndpointinfoResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Endpointinfo.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read endpointinfo, got error: %s", err))
		return
	}

	// Resource is gone: signal removal via null Id.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check endpointkind
		if idVal, ok := idMap["endpointkind"]; ok {
			if val, ok := v["endpointkind"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["endpointkind"].(string); ok {
			match = false
			continue
		}

		// Check endpointname
		if idVal, ok := idMap["endpointname"]; ok {
			if val, ok := v["endpointname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["endpointname"].(string); ok {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Binding is gone: signal removal via null Id.
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	endpointinfoSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
