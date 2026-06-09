package clusternodegroup_clusternode_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ClusternodegroupClusternodeBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupClusternodeBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupClusternodeBindingResource)(nil)

func NewClusternodegroupClusternodeBindingResource() resource.Resource {
	return &ClusternodegroupClusternodeBindingResource{}
}

// ClusternodegroupClusternodeBindingResource defines the resource implementation.
type ClusternodegroupClusternodeBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupClusternodeBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupClusternodeBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_clusternode_binding"
}

func (r *ClusternodegroupClusternodeBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupClusternodeBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupClusternodeBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_clusternode_binding resource")
	clusternodegroup_clusternode_binding := clusternodegroup_clusternode_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Clusternodegroup_clusternode_binding.Type(), &clusternodegroup_clusternode_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_clusternode_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created clusternodegroup_clusternode_binding resource")

	// Set ID for the resource before reading state.
	// Composite ID = name,node; name is URL-encoded, node is a plain integer.
	data.Id = types.StringValue(clusternodegroup_clusternode_bindingComposeId(data.Name.ValueString(), data.Node.ValueInt64()))

	// Read the updated state back
	r.readClusternodegroupClusternodeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupClusternodeBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupClusternodeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_clusternode_binding resource")

	r.readClusternodegroupClusternodeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupClusternodeBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClusternodegroupClusternodeBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// NITRO has no update endpoint for this binding (only add/delete), and all
	// schema attributes are RequiresReplace, so Update is a no-op.
	tflog.Debug(ctx, "Update is a no-op for clusternodegroup_clusternode_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readClusternodegroupClusternodeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupClusternodeBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupClusternodeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_clusternode_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent
	// (name) as the resource name and the bound node passed as an arg. This matches
	// the SDK v2 contract. Parse the ID to recover both (handles legacy + new format).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "node"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name := idMap["name"]
	args := make([]string, 0)
	if val, ok := idMap["node"]; ok && val != "" {
		args = append(args, fmt.Sprintf("node:%s", val))
	}

	err = r.client.DeleteResourceWithArgs(service.Clusternodegroup_clusternode_binding.Type(), name, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete clusternodegroup_clusternode_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted clusternodegroup_clusternode_binding binding")
}

// Helper function to read clusternodegroup_clusternode_binding data from API
func (r *ClusternodegroupClusternodeBindingResource) readClusternodegroupClusternodeBindingFromApi(ctx context.Context, data *ClusternodegroupClusternodeBindingResourceModel, diags *diag.Diagnostics) {

	// Parse the composite ID (handles both legacy comma "name,node" and new
	// "name:<enc>,node:<int>" formats) to recover the parent name and node filter.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "node"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	nameFilter := idMap["name"]
	nodeFilter, _ := utils.ConvertToInt64(idMap["node"])

	var dataArr []map[string]interface{}

	// The NITRO binding GET requires the parent (name) to be supplied as the
	// resource name. This matches the SDK v2 contract.
	findParams := service.FindParams{
		ResourceType:             service.Clusternodegroup_clusternode_binding.Type(),
		ResourceName:             nameFilter,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_clusternode_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "clusternodegroup_clusternode_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check name
		if val, ok := v["name"].(string); ok {
			if val != nameFilter {
				match = false
				continue
			}
		} else {
			match = false
			continue
		}

		// Check node (NITRO may return it as a number or string)
		if val, ok := v["node"]; ok {
			nodeVal, _ := utils.ConvertToInt64(val)
			if nodeVal != nodeFilter {
				match = false
				continue
			}
		} else {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("clusternodegroup_clusternode_binding not found with the provided ID attributes"))
		return
	}

	clusternodegroup_clusternode_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
