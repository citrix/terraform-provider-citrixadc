package clusternodegroup_streamidentifier_binding

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
var _ resource.Resource = &ClusternodegroupStreamidentifierBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupStreamidentifierBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupStreamidentifierBindingResource)(nil)

func NewClusternodegroupStreamidentifierBindingResource() resource.Resource {
	return &ClusternodegroupStreamidentifierBindingResource{}
}

// ClusternodegroupStreamidentifierBindingResource defines the resource implementation.
type ClusternodegroupStreamidentifierBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupStreamidentifierBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_streamidentifier_binding"
}

func (r *ClusternodegroupStreamidentifierBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_streamidentifier_binding resource")
	clusternodegroup_streamidentifier_binding := clusternodegroup_streamidentifier_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Clusternodegroup_streamidentifier_binding.Type(), &clusternodegroup_streamidentifier_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_streamidentifier_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created clusternodegroup_streamidentifier_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("identifiername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Identifiername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readClusternodegroupStreamidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_streamidentifier_binding resource")

	r.readClusternodegroupStreamidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for clusternodegroup_streamidentifier_binding: NITRO exposes only add (PUT)
	// and delete (no update/change endpoint), and all schema attributes are RequiresReplace, so Terraform
	// recreates the resource on any change rather than calling Update.
	tflog.Debug(ctx, "Update is a no-op for clusternodegroup_streamidentifier_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readClusternodegroupStreamidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_streamidentifier_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgsMap.
	// Per NITRO doc: DELETE .../clusternodegroup_streamidentifier_binding/<name_value>?args=identifiername:<value>
	// name is the URL PATH key (parent); identifiername is the ONLY query arg.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "identifiername"}, nil)
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
	if val, ok := idMap["identifiername"]; ok && val != "" {
		argsMap["identifiername"] = utils.UrlEncode(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Clusternodegroup_streamidentifier_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete clusternodegroup_streamidentifier_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted clusternodegroup_streamidentifier_binding binding")
}

// Helper function to read clusternodegroup_streamidentifier_binding data from API
func (r *ClusternodegroupStreamidentifierBindingResource) readClusternodegroupStreamidentifierBindingFromApi(ctx context.Context, data *ClusternodegroupStreamidentifierBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	// name is the URL PATH key (parent); identifiername is the filter within the parent's bindings.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "identifiername"}, nil)
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
		ResourceType:             service.Clusternodegroup_streamidentifier_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_streamidentifier_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "clusternodegroup_streamidentifier_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check name
		if idVal, ok := idMap["name"]; ok {
			if val, ok := v["name"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["name"].(string); ok {
			match = false
			continue
		}

		// Check identifiername
		if idVal, ok := idMap["identifiername"]; ok {
			if val, ok := v["identifiername"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["identifiername"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("clusternodegroup_streamidentifier_binding not found with the provided ID attributes"))
		return
	}

	clusternodegroup_streamidentifier_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
