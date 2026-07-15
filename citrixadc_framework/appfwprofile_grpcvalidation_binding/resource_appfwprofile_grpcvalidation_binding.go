package appfwprofile_grpcvalidation_binding

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
var _ resource.Resource = &AppfwprofileGrpcvalidationBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileGrpcvalidationBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileGrpcvalidationBindingResource)(nil)

func NewAppfwprofileGrpcvalidationBindingResource() resource.Resource {
	return &AppfwprofileGrpcvalidationBindingResource{}
}

// AppfwprofileGrpcvalidationBindingResource defines the resource implementation.
type AppfwprofileGrpcvalidationBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileGrpcvalidationBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileGrpcvalidationBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_grpcvalidation_binding"
}

func (r *AppfwprofileGrpcvalidationBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileGrpcvalidationBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileGrpcvalidationBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_grpcvalidation_binding resource")
	appfwprofile_grpcvalidation_binding := appfwprofile_grpcvalidation_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_grpcvalidation_binding.Type(), &appfwprofile_grpcvalidation_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_grpcvalidation_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_grpcvalidation_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("grpc_relax_validation_action:%s", utils.UrlEncode(fmt.Sprintf("%v", data.GrpcRelaxValidationAction.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("grpcvalidation:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Grpcvalidation.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileGrpcvalidationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileGrpcvalidationBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileGrpcvalidationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_grpcvalidation_binding resource")

	r.readAppfwprofileGrpcvalidationBindingFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Self-heal: object was deleted out-of-band, remove from state so a
	// subsequent apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileGrpcvalidationBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileGrpcvalidationBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Pattern 5: Update is a no-op for this binding. NITRO exposes no update
	// endpoint for appfwprofile_grpcvalidation_binding (only add/delete/get), and
	// every schema attribute is RequiresReplace, so Terraform never invokes Update
	// with an actual changed value. Just re-read and persist state.
	tflog.Debug(ctx, "Update is a no-op for appfwprofile_grpcvalidation_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readAppfwprofileGrpcvalidationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileGrpcvalidationBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileGrpcvalidationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_grpcvalidation_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
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
	if val, ok := idMap["grpc_relax_validation_action"]; ok && val != "" {
		argsMap["grpc_relax_validation_action"] = val
	}
	if val, ok := idMap["grpcvalidation"]; ok && val != "" {
		argsMap["grpcvalidation"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_grpcvalidation_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_grpcvalidation_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_grpcvalidation_binding binding")
}

// Helper function to read appfwprofile_grpcvalidation_binding data from API
func (r *AppfwprofileGrpcvalidationBindingResource) readAppfwprofileGrpcvalidationBindingFromApi(ctx context.Context, data *AppfwprofileGrpcvalidationBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
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
		ResourceType:             service.Appfwprofile_grpcvalidation_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_grpcvalidation_binding, got error: %s", err))
		return
	}

	// Resource is missing - signal removal for self-heal.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check grpc_relax_validation_action
		if idVal, ok := idMap["grpc_relax_validation_action"]; ok {
			if val, ok := v["grpc_relax_validation_action"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["grpc_relax_validation_action"].(string); ok {
			match = false
			continue
		}

		// Check grpcvalidation
		if idVal, ok := idMap["grpcvalidation"]; ok {
			if val, ok := v["grpcvalidation"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["grpcvalidation"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing - signal removal for self-heal.
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	appfwprofile_grpcvalidation_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])

	// Backfill the identity / composite-ID attributes from the parsed ID so that
	// `terraform import` (which has no prior plan/state) fully round-trips. These
	// are category (a) attributes: name, grpcvalidation and
	// grpc_relax_validation_action are all components of data.Id, so they can
	// always be recovered from idMap. This is done after the found/len self-heal
	// checks above (which return early on not-found, preserving null-Id self-heal)
	// and does not alter the ID composition.
	data.Name = types.StringValue(name_Name)
	if val, ok := idMap["grpcvalidation"]; ok {
		data.Grpcvalidation = types.StringValue(val)
	}
	if val, ok := idMap["grpc_relax_validation_action"]; ok {
		data.GrpcRelaxValidationAction = types.StringValue(val)
	}
}
