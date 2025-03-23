package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func cicdResource() resource.Resource {
	return &resourceCICD{}
}

type resourceCICD struct{}

func (n *resourceCICD) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

func (n *resourceCICD) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"build": schema.StringAttribute{
				MarkdownDescription: "Build command to be executed.",
				Optional:            true,
			},
			"test": schema.StringAttribute{
				MarkdownDescription: "Test command to be executed.",
				Optional:            true,
			},
			"working_directory": schema.StringAttribute{
				MarkdownDescription: "Working directory for the commands.",
				Optional:            true,
			},
			"build_and_test": schema.StringAttribute{
				MarkdownDescription: "Build and test command to be executed.",
				Optional:            true,
			},
			"docker_build": schema.StringAttribute{
				MarkdownDescription: "Docker build command to be executed.",
				Optional:            true,
			},
			"docker_push": schema.StringAttribute{
				MarkdownDescription: "Docker push command to be executed.",
				Optional:            true,
			},
			"container_registry_url": schema.StringAttribute{
				MarkdownDescription: "Container registry URL.",
				Optional:            true,
			},
			"dockerfile_directory": schema.StringAttribute{
				MarkdownDescription: "Directory containing the Dockerfile.",
				Optional:            true,
			},
			"timestamp": schema.StringAttribute{
				MarkdownDescription: "Timestamp of the last update.",
				Computed:            true,
			},
		},
	}
}

func (n *resourceCICD) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan cicdModel
}

func (r *resourceCICD) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *resourceCICD) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.Create(ctx, resource.CreateRequest(req), &resource.CreateResponse(resp))
}

func (r *resourceCICD) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type cicdModel struct {
	ID                   types.String `tfsdk:"id"`
	Build                types.String `tfsdk:"build"`
	Test                 types.String `tfsdk:"test"`
	WorkingDirectory     types.String `tfsdk:"working_directory"`
	BuildAndTest         types.String `tfsdk:"build_and_test"`
	DockerBuild          types.String `tfsdk:"docker_build"`
	DockerPush           types.String `tfsdk:"docker_push"`
	ContainerRegistryURL types.String `tfsdk:"container_registry_url"`
	DockerfileDirectory  types.String `tfsdk:"dockerfile_directory"`
	Timestamp            types.String `tfsdk:"timestamp"`
}