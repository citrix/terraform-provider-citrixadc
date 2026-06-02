package metricsprofile_lbvserver_binding

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
var _ resource.Resource = &MetricsprofileLbvserverBindingResource{}
var _ resource.ResourceWithConfigure = (*MetricsprofileLbvserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*MetricsprofileLbvserverBindingResource)(nil)

func NewMetricsprofileLbvserverBindingResource() resource.Resource {
	return &MetricsprofileLbvserverBindingResource{}
}

// MetricsprofileLbvserverBindingResource defines the resource implementation.
type MetricsprofileLbvserverBindingResource struct {
	client *service.NitroClient
}

func (r *MetricsprofileLbvserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MetricsprofileLbvserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metricsprofile_lbvserver_binding"
}

func (r *MetricsprofileLbvserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *MetricsprofileLbvserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MetricsprofileLbvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating metricsprofile_lbvserver_binding resource")
	metricsprofile_lbvserver_binding := metricsprofile_lbvserver_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Metricsprofile_lbvserver_binding.Type(), &metricsprofile_lbvserver_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create metricsprofile_lbvserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created metricsprofile_lbvserver_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("entityname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Entityname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("entitytype:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Entitytype.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readMetricsprofileLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MetricsprofileLbvserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MetricsprofileLbvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading metricsprofile_lbvserver_binding resource")

	r.readMetricsprofileLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MetricsprofileLbvserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state MetricsprofileLbvserverBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for metricsprofile_lbvserver_binding; NITRO exposes no
	// update endpoint and all attributes use RequiresReplace.
	tflog.Debug(ctx, "Update is a no-op for metricsprofile_lbvserver_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readMetricsprofileLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MetricsprofileLbvserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MetricsprofileLbvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting metricsprofile_lbvserver_binding resource")
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

	err = r.client.DeleteResourceWithArgsMap(service.Metricsprofile_lbvserver_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete metricsprofile_lbvserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted metricsprofile_lbvserver_binding binding")
}

// Helper function to read metricsprofile_lbvserver_binding data from API
func (r *MetricsprofileLbvserverBindingResource) readMetricsprofileLbvserverBindingFromApi(ctx context.Context, data *MetricsprofileLbvserverBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Metricsprofile_lbvserver_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read metricsprofile_lbvserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "metricsprofile_lbvserver_binding returned empty array.")
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
		diags.AddError("Client Error", fmt.Sprintf("metricsprofile_lbvserver_binding not found with the provided ID attributes"))
		return
	}

	metricsprofile_lbvserver_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}