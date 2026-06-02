package metricsprofile_crvserver_binding

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
var _ resource.Resource = &MetricsprofileCrvserverBindingResource{}
var _ resource.ResourceWithConfigure = (*MetricsprofileCrvserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*MetricsprofileCrvserverBindingResource)(nil)

func NewMetricsprofileCrvserverBindingResource() resource.Resource {
	return &MetricsprofileCrvserverBindingResource{}
}

// MetricsprofileCrvserverBindingResource defines the resource implementation.
type MetricsprofileCrvserverBindingResource struct {
	client *service.NitroClient
}

func (r *MetricsprofileCrvserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MetricsprofileCrvserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metricsprofile_crvserver_binding"
}

func (r *MetricsprofileCrvserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *MetricsprofileCrvserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MetricsprofileCrvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating metricsprofile_crvserver_binding resource")
	metricsprofile_crvserver_binding := metricsprofile_crvserver_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Metricsprofile_crvserver_binding.Type(), &metricsprofile_crvserver_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create metricsprofile_crvserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created metricsprofile_crvserver_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("entityname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Entityname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("entitytype:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Entitytype.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readMetricsprofileCrvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MetricsprofileCrvserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MetricsprofileCrvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading metricsprofile_crvserver_binding resource")

	r.readMetricsprofileCrvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MetricsprofileCrvserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state MetricsprofileCrvserverBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for metricsprofile_crvserver_binding; NITRO exposes no
	// update endpoint for this binding and all attributes are RequiresReplace.
	tflog.Debug(ctx, "Update is a no-op for metricsprofile_crvserver_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readMetricsprofileCrvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MetricsprofileCrvserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MetricsprofileCrvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting metricsprofile_crvserver_binding resource")
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
	if val, ok := idMap["entityname"]; ok && val != "" {
		argsMap["entityname"] = val
	}
	if val, ok := idMap["entitytype"]; ok && val != "" {
		argsMap["entitytype"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Metricsprofile_crvserver_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete metricsprofile_crvserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted metricsprofile_crvserver_binding binding")
}

// Helper function to read metricsprofile_crvserver_binding data from API
func (r *MetricsprofileCrvserverBindingResource) readMetricsprofileCrvserverBindingFromApi(ctx context.Context, data *MetricsprofileCrvserverBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Metricsprofile_crvserver_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read metricsprofile_crvserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "metricsprofile_crvserver_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check entityname
		if idVal, ok := idMap["entityname"]; ok {
			if val, ok := v["entityname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["entityname"].(string); ok {
			match = false
			continue
		}

		// Check entitytype
		if idVal, ok := idMap["entitytype"]; ok {
			if val, ok := v["entitytype"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["entitytype"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("metricsprofile_crvserver_binding not found with the provided ID attributes"))
		return
	}

	metricsprofile_crvserver_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}