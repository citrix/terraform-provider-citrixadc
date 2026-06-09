package servicegroup_servicegroupmemberlist_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ServicegroupServicegroupmemberlistBindingResourceModel describes the resource data model.
type ServicegroupServicegroupmemberlistBindingResourceModel struct {
	Id               types.String                  `tfsdk:"id"`
	Servicegroupname types.String                  `tfsdk:"servicegroupname"`
	Members          []ServicegroupMemberListModel `tfsdk:"members"`
}

// ServicegroupMemberListModel describes a single servicegroup member entry.
type ServicegroupMemberListModel struct {
	Ip     types.String `tfsdk:"ip"`
	Port   types.Int64  `tfsdk:"port"`
	Weight types.Int64  `tfsdk:"weight"`
	State  types.String `tfsdk:"state"`
	Order  types.Int64  `tfsdk:"order"`
}

func (r *ServicegroupServicegroupmemberlistBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the servicegroup_servicegroupmemberlist_binding resource.",
			},
			"servicegroupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the service group.",
			},
			"members": schema.ListNestedAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "Desired servicegroupmember binding set. Any existing servicegroupmember which is not part of the input will be deleted or disabled based on graceful setting on servicegroup. Each member requires exactly one of ip or servername to be set.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							Optional:    true,
							Description: "IP Address.",
						},
						"port": schema.Int64Attribute{
							Optional:    true,
							Description: "The port number of the service to be enabled.",
						},
						"weight": schema.Int64Attribute{
							Optional:    true,
							Description: "Weight to assign to the servicegroup member.",
						},
						"state": schema.StringAttribute{
							Optional:    true,
							Description: "Initial state of the service group. Possible values = ENABLED, DISABLED",
						},
						"order": schema.Int64Attribute{
							Optional:    true,
							Description: "Order number to be assigned to the servicegroup member.",
						},
					},
				},
			},
		},
	}
}

func servicegroup_servicegroupmemberlist_bindingGetThePayloadFromthePlan(ctx context.Context, data *ServicegroupServicegroupmemberlistBindingResourceModel) basic.Servicegroupservicegroupmemberlistbinding {
	tflog.Debug(ctx, "In servicegroup_servicegroupmemberlist_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	servicegroup_servicegroupmemberlist_binding := basic.Servicegroupservicegroupmemberlistbinding{}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		servicegroup_servicegroupmemberlist_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}

	if len(data.Members) > 0 {
		members := make([]basic.Members, 0, len(data.Members))
		for _, m := range data.Members {
			member := basic.Members{}
			if !m.Ip.IsNull() && !m.Ip.IsUnknown() {
				member.Ip = m.Ip.ValueString()
			}
			if !m.Port.IsNull() && !m.Port.IsUnknown() {
				member.Port = utils.IntPtr(int(m.Port.ValueInt64()))
			}
			if !m.Weight.IsNull() && !m.Weight.IsUnknown() {
				member.Weight = utils.IntPtr(int(m.Weight.ValueInt64()))
			}
			if !m.State.IsNull() && !m.State.IsUnknown() {
				member.State = m.State.ValueString()
			}
			if !m.Order.IsNull() && !m.Order.IsUnknown() {
				member.Order = utils.IntPtr(int(m.Order.ValueInt64()))
			}
			members = append(members, member)
		}
		servicegroup_servicegroupmemberlist_binding.Members = members
	}

	return servicegroup_servicegroupmemberlist_binding
}
