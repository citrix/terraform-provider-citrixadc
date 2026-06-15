package rnat6_nsip6_binding

import (
	"context"
	"fmt"
	neturl "net/url"
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
var _ resource.Resource = &Rnat6Nsip6BindingResource{}
var _ resource.ResourceWithConfigure = (*Rnat6Nsip6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*Rnat6Nsip6BindingResource)(nil)

func NewRnat6Nsip6BindingResource() resource.Resource {
	return &Rnat6Nsip6BindingResource{}
}

// Rnat6Nsip6BindingResource defines the resource implementation.
type Rnat6Nsip6BindingResource struct {
	client *service.NitroClient
}

func (r *Rnat6Nsip6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Rnat6Nsip6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnat6_nsip6_binding"
}

func (r *Rnat6Nsip6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Rnat6Nsip6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Rnat6Nsip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rnat6_nsip6_binding resource")
	rnat6_nsip6_binding := rnat6_nsip6_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Rnat6_nsip6_binding.Type(), &rnat6_nsip6_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rnat6_nsip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created rnat6_nsip6_binding resource")

	// Set ID for the resource before reading state.
	// Identity is "name,natip6" (matches the SDK v2 ID and resource_id_mapping.json).
	// ownergroup is a delete arg / read-back attribute, not part of the identity.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("natip6:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Natip6.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readRnat6Nsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Rnat6Nsip6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Rnat6Nsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rnat6_nsip6_binding resource")

	r.readRnat6Nsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Rnat6Nsip6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Rnat6Nsip6BindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating rnat6_nsip6_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		rnat6_nsip6_binding := rnat6_nsip6_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Rnat6_nsip6_binding.Type(), &rnat6_nsip6_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rnat6_nsip6_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated rnat6_nsip6_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for rnat6_nsip6_binding resource, skipping update")
	}

	// Read the updated state back
	r.readRnat6Nsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Rnat6Nsip6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Rnat6Nsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rnat6_nsip6_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// ParseIdString handles both the new "name:..,natip6:.." format and the legacy
	// positional "name,natip6" format from imported SDK v2 state.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "natip6"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// Build delete args. Values may contain ':' and '/' (e.g. IPv6 addresses), which are
	// NITRO arg delimiters, so URL-encode each value (mirrors the SDK v2 delete args).
	args := make([]string, 0)
	if val, ok := idMap["natip6"]; ok && val != "" {
		args = append(args, fmt.Sprintf("natip6:%s", neturl.QueryEscape(val)))
	}
	// ownergroup is read from state (not the ID) to match the SDK v2 delete behaviour.
	if !data.Ownergroup.IsNull() && data.Ownergroup.ValueString() != "" {
		args = append(args, fmt.Sprintf("ownergroup:%s", neturl.QueryEscape(data.Ownergroup.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Rnat6_nsip6_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete rnat6_nsip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted rnat6_nsip6_binding binding")
}

// Helper function to read rnat6_nsip6_binding data from API
func (r *Rnat6Nsip6BindingResource) readRnat6Nsip6BindingFromApi(ctx context.Context, data *Rnat6Nsip6BindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "natip6"}, nil)
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
		ResourceType:             service.Rnat6_nsip6_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rnat6_nsip6_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "rnat6_nsip6_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check natip6 (the identity discriminator within the parent's array).
		if idVal, ok := idMap["natip6"]; ok {
			if val, ok := v["natip6"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["natip6"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("rnat6_nsip6_binding not found with the provided ID attributes"))
		return
	}

	rnat6_nsip6_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
