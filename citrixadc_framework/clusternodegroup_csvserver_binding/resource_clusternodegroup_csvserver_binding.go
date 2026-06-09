package clusternodegroup_csvserver_binding

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
var _ resource.Resource = &ClusternodegroupCsvserverBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupCsvserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupCsvserverBindingResource)(nil)

func NewClusternodegroupCsvserverBindingResource() resource.Resource {
	return &ClusternodegroupCsvserverBindingResource{}
}

// ClusternodegroupCsvserverBindingResource defines the resource implementation.
type ClusternodegroupCsvserverBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupCsvserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupCsvserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_csvserver_binding"
}

func (r *ClusternodegroupCsvserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupCsvserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupCsvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_csvserver_binding resource")
	clusternodegroup_csvserver_binding := clusternodegroup_csvserver_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Clusternodegroup_csvserver_binding.Type(), &clusternodegroup_csvserver_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_csvserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created clusternodegroup_csvserver_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vserver:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vserver.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readClusternodegroupCsvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupCsvserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupCsvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_csvserver_binding resource")

	r.readClusternodegroupCsvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupCsvserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClusternodegroupCsvserverBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for clusternodegroup_csvserver_binding: NITRO exposes only add (PUT)
	// and delete (no update/change endpoint), and all schema attributes are RequiresReplace, so Terraform
	// recreates the resource on any change rather than calling Update.
	tflog.Debug(ctx, "Update is a no-op for clusternodegroup_csvserver_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readClusternodegroupCsvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupCsvserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupCsvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_csvserver_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent (name) as the
	// resource name and the bound vserver passed as an arg. This matches the SDK v2 contract.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "vserver"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name := idMap["name"]
	args := make([]string, 0)
	if val, ok := idMap["vserver"]; ok && val != "" {
		args = append(args, fmt.Sprintf("vserver:%s", val))
	}

	err = r.client.DeleteResourceWithArgs(service.Clusternodegroup_csvserver_binding.Type(), name, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete clusternodegroup_csvserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted clusternodegroup_csvserver_binding binding")
}

// Helper function to read clusternodegroup_csvserver_binding data from API
func (r *ClusternodegroupCsvserverBindingResource) readClusternodegroupCsvserverBindingFromApi(ctx context.Context, data *ClusternodegroupCsvserverBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "vserver"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	// The NITRO binding GET requires the parent (name) to be supplied as the resource name.
	findParams := service.FindParams{
		ResourceType:             service.Clusternodegroup_csvserver_binding.Type(),
		ResourceName:             idMap["name"],
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_csvserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "clusternodegroup_csvserver_binding returned empty array")
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

		// Check vserver
		if idVal, ok := idMap["vserver"]; ok {
			if val, ok := v["vserver"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["vserver"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("clusternodegroup_csvserver_binding not found with the provided ID attributes"))
		return
	}

	clusternodegroup_csvserver_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
