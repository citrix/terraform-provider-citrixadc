package csvserver_lbvserver_binding

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
var _ resource.Resource = &CsvserverLbvserverBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverLbvserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverLbvserverBindingResource)(nil)

func NewCsvserverLbvserverBindingResource() resource.Resource {
	return &CsvserverLbvserverBindingResource{}
}

// CsvserverLbvserverBindingResource defines the resource implementation.
type CsvserverLbvserverBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverLbvserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverLbvserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_lbvserver_binding"
}

func (r *CsvserverLbvserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverLbvserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverLbvserverBindingResourceModel
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_lbvserver_binding resource")
	csvserver_lbvserver_binding := csvserver_lbvserver_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	err := r.client.UpdateUnnamedResource(service.Csvserver_lbvserver_binding.Type(), &csvserver_lbvserver_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_lbvserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created csvserver_lbvserver_binding resource")

	// Set ID for the resource before reading state
	bindingId := fmt.Sprintf("%s,%s", data.Name.ValueString(), data.Lbvserver.ValueString())
	data.Id = types.StringValue(bindingId)

	// Read the updated state back
	if !r.readCsvserverLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "csvserver_lbvserver_binding not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverLbvserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverLbvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_lbvserver_binding resource")

	found := r.readCsvserverLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverLbvserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CsvserverLbvserverBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating csvserver_lbvserver_binding resource")

	// For binding resources, updates typically require delete and recreate
	// This should not be called as all fields are ForceNew
	resp.Diagnostics.AddError("Update Not Supported", "csvserver_lbvserver_binding does not support updates. All fields are ForceNew.")
}

func (r *CsvserverLbvserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverLbvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_lbvserver_binding resource")

	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "lbvserver"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}
	name_Name, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}
	lbvserver_Name, ok := idMap["lbvserver"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "ID attribute 'lbvserver' not found in ID string")
		return
	}

	// Build args for delete
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("lbvserver:%s", lbvserver_Name))

	err = r.client.DeleteResourceWithArgs(service.Csvserver_lbvserver_binding.Type(), name_Name, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete csvserver_lbvserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted csvserver_lbvserver_binding resource")
}

// Helper function to read csvserver_lbvserver_binding data from API
func (r *CsvserverLbvserverBindingResource) readCsvserverLbvserverBindingFromApi(ctx context.Context, data *CsvserverLbvserverBindingResourceModel, diags *diag.Diagnostics) bool {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "lbvserver"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return false
	}
	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return false
	}
	lbvserver_Name, ok := idMap["lbvserver"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'lbvserver' not found in ID string")
		return false
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Csvserver_lbvserver_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_lbvserver_binding, got error: %s", err))
		return false
	}

	// Resource is missing
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check lbvserver
		if val, ok := v["lbvserver"].(string); ok {
			if val != lbvserver_Name {
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

	//  Resource is missing
	if foundIndex == -1 {
		return false
	}

	csvserver_lbvserver_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
