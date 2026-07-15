package appfwprofile_jsonblockkeyword_binding

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
var _ resource.Resource = &AppfwprofileJsonblockkeywordBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileJsonblockkeywordBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileJsonblockkeywordBindingResource)(nil)

func NewAppfwprofileJsonblockkeywordBindingResource() resource.Resource {
	return &AppfwprofileJsonblockkeywordBindingResource{}
}

// AppfwprofileJsonblockkeywordBindingResource defines the resource implementation.
type AppfwprofileJsonblockkeywordBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileJsonblockkeywordBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileJsonblockkeywordBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_jsonblockkeyword_binding"
}

func (r *AppfwprofileJsonblockkeywordBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileJsonblockkeywordBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileJsonblockkeywordBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_jsonblockkeyword_binding resource")
	appfwprofile_jsonblockkeyword_binding := appfwprofile_jsonblockkeyword_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsonblockkeyword_binding.Type(), &appfwprofile_jsonblockkeyword_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_jsonblockkeyword_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_jsonblockkeyword_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("jsonblockkeyword:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Jsonblockkeyword.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("jsonblockkeywordurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Jsonblockkeywordurl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("keyname_json_blockkeyword:%s", utils.UrlEncode(fmt.Sprintf("%v", data.KeynameJsonBlockkeyword.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileJsonblockkeywordBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsonblockkeywordBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileJsonblockkeywordBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_jsonblockkeyword_binding resource")

	r.readAppfwprofileJsonblockkeywordBindingFromApi(ctx, &data, &resp.Diagnostics)

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

func (r *AppfwprofileJsonblockkeywordBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileJsonblockkeywordBindingResourceModel

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
	// endpoint for appfwprofile_jsonblockkeyword_binding (only add/delete/get), and
	// every schema attribute is RequiresReplace, so Terraform never invokes Update
	// with an actual changed value. Just re-read and persist state.
	tflog.Debug(ctx, "Update is a no-op for appfwprofile_jsonblockkeyword_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readAppfwprofileJsonblockkeywordBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsonblockkeywordBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileJsonblockkeywordBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_jsonblockkeyword_binding resource")
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
	if val, ok := idMap["jsonblockkeyword"]; ok && val != "" {
		argsMap["jsonblockkeyword"] = val
	}
	if val, ok := idMap["jsonblockkeywordurl"]; ok && val != "" {
		// jsonblockkeywordurl is a URL containing reserved characters (':' and '/').
		// nitro-go does not encode ?args= values, so encode it to avoid a 400 from NITRO.
		argsMap["jsonblockkeywordurl"] = utils.UrlEncode(val)
	}
	if val, ok := idMap["keyname_json_blockkeyword"]; ok && val != "" {
		argsMap["keyname_json_blockkeyword"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_jsonblockkeyword_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_jsonblockkeyword_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_jsonblockkeyword_binding binding")
}

// Helper function to read appfwprofile_jsonblockkeyword_binding data from API
func (r *AppfwprofileJsonblockkeywordBindingResource) readAppfwprofileJsonblockkeywordBindingFromApi(ctx context.Context, data *AppfwprofileJsonblockkeywordBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Appfwprofile_jsonblockkeyword_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_jsonblockkeyword_binding, got error: %s", err))
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

		// Check jsonblockkeyword
		if idVal, ok := idMap["jsonblockkeyword"]; ok {
			if val, ok := v["jsonblockkeyword"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["jsonblockkeyword"].(string); ok {
			match = false
			continue
		}

		// Check jsonblockkeywordurl
		if idVal, ok := idMap["jsonblockkeywordurl"]; ok {
			if val, ok := v["jsonblockkeywordurl"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["jsonblockkeywordurl"].(string); ok {
			match = false
			continue
		}

		// Check keyname_json_blockkeyword
		if idVal, ok := idMap["keyname_json_blockkeyword"]; ok {
			if val, ok := v["keyname_json_blockkeyword"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["keyname_json_blockkeyword"].(string); ok {
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

	// Backfill the identity / ID-component attributes from the parsed ID so that
	// `terraform import` (which has no prior plan/state) fully round-trips these
	// RequiresReplace identity attributes. On a normal Read the values are
	// identical to the prior state, so this is a no-op there. These are the exact
	// values the ID is composed from, so data.Id remains unchanged. Placed after
	// the found/len self-heal checks so null-Id-on-not-found is preserved.
	if val, ok := idMap["name"]; ok {
		data.Name = types.StringValue(val)
	}
	if val, ok := idMap["jsonblockkeyword"]; ok {
		data.Jsonblockkeyword = types.StringValue(val)
	}
	if val, ok := idMap["keyname_json_blockkeyword"]; ok {
		data.KeynameJsonBlockkeyword = types.StringValue(val)
	}
	if val, ok := idMap["jsonblockkeywordurl"]; ok {
		data.Jsonblockkeywordurl = types.StringValue(val)
	}

	appfwprofile_jsonblockkeyword_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
