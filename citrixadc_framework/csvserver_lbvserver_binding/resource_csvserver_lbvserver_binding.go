package csvserver_lbvserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
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
	r.readCsvserverLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

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

	r.readCsvserverLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

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

	bindingId := data.Id.ValueString()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name_Name := idSlice[0]
	lbvserver_Name := idSlice[1]

	// Build args for delete
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("lbvserver:%s", lbvserver_Name))

	err := r.client.DeleteResourceWithArgs(service.Csvserver_lbvserver_binding.Type(), name_Name, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete csvserver_lbvserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted csvserver_lbvserver_binding resource")
}

// Helper function to read csvserver_lbvserver_binding data from API
func (r *CsvserverLbvserverBindingResource) readCsvserverLbvserverBindingFromApi(ctx context.Context, data *CsvserverLbvserverBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	bindingId := data.Id.ValueString()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name_Name := idSlice[0]
	lbvserver_Name := idSlice[1]

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Csvserver_lbvserver_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_lbvserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "csvserver_lbvserver_binding returned empty array.")
		return
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
		diags.AddError("Client Error", fmt.Sprintf("csvserver_lbvserver_binding not found with the provided ID attributes"))
		return
	}

	csvserver_lbvserver_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
