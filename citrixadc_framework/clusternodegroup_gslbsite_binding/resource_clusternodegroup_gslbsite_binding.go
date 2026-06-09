package clusternodegroup_gslbsite_binding

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
var _ resource.Resource = &ClusternodegroupGslbsiteBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupGslbsiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupGslbsiteBindingResource)(nil)

func NewClusternodegroupGslbsiteBindingResource() resource.Resource {
	return &ClusternodegroupGslbsiteBindingResource{}
}

// ClusternodegroupGslbsiteBindingResource defines the resource implementation.
type ClusternodegroupGslbsiteBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupGslbsiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupGslbsiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_gslbsite_binding"
}

func (r *ClusternodegroupGslbsiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupGslbsiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupGslbsiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_gslbsite_binding resource")
	clusternodegroup_gslbsite_binding := clusternodegroup_gslbsite_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Clusternodegroup_gslbsite_binding.Type(), &clusternodegroup_gslbsite_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_gslbsite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created clusternodegroup_gslbsite_binding resource")

	// Set ID for the resource before reading state
	// ID order must match resource_id_mapping.json ("name,gslbsite") and the datasource setter.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("gslbsite:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Gslbsite.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readClusternodegroupGslbsiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupGslbsiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupGslbsiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_gslbsite_binding resource")

	r.readClusternodegroupGslbsiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupGslbsiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClusternodegroupGslbsiteBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for clusternodegroup_gslbsite_binding: NITRO exposes only add (PUT)
	// and delete (no update/change endpoint), and all schema attributes are RequiresReplace, so Terraform
	// recreates the resource on any change rather than calling Update.
	tflog.Debug(ctx, "Update is a no-op for clusternodegroup_gslbsite_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readClusternodegroupGslbsiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupGslbsiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupGslbsiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_gslbsite_binding resource")
	// Parent-keyed binding delete: DELETE /clusternodegroup_gslbsite_binding/<name>?args=gslbsite:<v>
	// The keyless (empty-name) form silently no-ops; live ADC requires the parent name in the URL.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "gslbsite"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name := idMap["name"]
	args := make([]string, 0)
	if val, ok := idMap["gslbsite"]; ok && val != "" {
		args = append(args, fmt.Sprintf("gslbsite:%s", utils.UrlEncode(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Clusternodegroup_gslbsite_binding.Type(), name, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete clusternodegroup_gslbsite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted clusternodegroup_gslbsite_binding binding")
}

// Helper function to read clusternodegroup_gslbsite_binding data from API
func (r *ClusternodegroupGslbsiteBindingResource) readClusternodegroupGslbsiteBindingFromApi(ctx context.Context, data *ClusternodegroupGslbsiteBindingResourceModel, diags *diag.Diagnostics) {

	// Parent-keyed binding read: GET requires the parent name in the URL (keyless yields errorcode 1095).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "gslbsite"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Clusternodegroup_gslbsite_binding.Type(),
		ResourceName:             idMap["name"],
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_gslbsite_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "clusternodegroup_gslbsite_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check gslbsite
		if idVal, ok := idMap["gslbsite"]; ok {
			if val, ok := v["gslbsite"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["gslbsite"].(string); ok {
			match = false
			continue
		}

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

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("clusternodegroup_gslbsite_binding not found with the provided ID attributes"))
		return
	}

	clusternodegroup_gslbsite_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
