package clusternode_routemonitor_binding

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
var _ resource.Resource = &ClusternodeRoutemonitorBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodeRoutemonitorBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodeRoutemonitorBindingResource)(nil)

func NewClusternodeRoutemonitorBindingResource() resource.Resource {
	return &ClusternodeRoutemonitorBindingResource{}
}

// ClusternodeRoutemonitorBindingResource defines the resource implementation.
type ClusternodeRoutemonitorBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodeRoutemonitorBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodeRoutemonitorBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternode_routemonitor_binding"
}

func (r *ClusternodeRoutemonitorBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodeRoutemonitorBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodeRoutemonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternode_routemonitor_binding resource")
	clusternode_routemonitor_binding := clusternode_routemonitor_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Clusternode_routemonitor_binding.Type(), &clusternode_routemonitor_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternode_routemonitor_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created clusternode_routemonitor_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("netmask:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Netmask.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("nodeid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Nodeid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("routemonitor:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Routemonitor.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readClusternodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodeRoutemonitorBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodeRoutemonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternode_routemonitor_binding resource")

	r.readClusternodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodeRoutemonitorBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClusternodeRoutemonitorBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating clusternode_routemonitor_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		clusternode_routemonitor_binding := clusternode_routemonitor_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Clusternode_routemonitor_binding.Type(), &clusternode_routemonitor_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusternode_routemonitor_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated clusternode_routemonitor_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for clusternode_routemonitor_binding resource, skipping update")
	}

	// Read the updated state back
	r.readClusternodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodeRoutemonitorBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodeRoutemonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternode_routemonitor_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"nodeid", "routemonitor"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	nodeid_value, ok := idMap["nodeid"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'nodeid' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["netmask"]; ok && val != "" {
		argsMap["netmask"] = val
	}
	if val, ok := idMap["routemonitor"]; ok && val != "" {
		argsMap["routemonitor"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Clusternode_routemonitor_binding.Type(), nodeid_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete clusternode_routemonitor_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted clusternode_routemonitor_binding binding")
}

// Helper function to read clusternode_routemonitor_binding data from API
func (r *ClusternodeRoutemonitorBindingResource) readClusternodeRoutemonitorBindingFromApi(ctx context.Context, data *ClusternodeRoutemonitorBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"nodeid", "routemonitor"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	nodeid_Name, ok := idMap["nodeid"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'nodeid' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Clusternode_routemonitor_binding.Type(),
		ResourceName:             nodeid_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternode_routemonitor_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "clusternode_routemonitor_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check netmask
		if idVal, ok := idMap["netmask"]; ok {
			if val, ok := v["netmask"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["netmask"].(string); ok {
			match = false
			continue
		}

		// Check routemonitor
		if idVal, ok := idMap["routemonitor"]; ok {
			if val, ok := v["routemonitor"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["routemonitor"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("clusternode_routemonitor_binding not found with the provided ID attributes"))
		return
	}

	clusternode_routemonitor_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
