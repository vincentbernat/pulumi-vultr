// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about operating systems that can be launched when creating a Vultr VPS.
//
// ## Example Usage
//
// Get the information for an operating system by `name`:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vultr/sdk/go/vultr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vultr.GetOS(ctx, &GetOSArgs{
// 			Filters: []GetOSFilter{
// 				GetOSFilter{
// 					Name: "name",
// 					Values: []string{
// 						"CentOS 7 x64",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOS(ctx *pulumi.Context, args *GetOSArgs, opts ...pulumi.InvokeOption) (*GetOSResult, error) {
	var rv GetOSResult
	err := ctx.Invoke("vultr:index/getOS:getOS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOS.
type GetOSArgs struct {
	// Query parameters for finding operating systems.
	Filters []GetOSFilter `pulumi:"filters"`
}

// A collection of values returned by getOS.
type GetOSResult struct {
	// The architecture of the operating system.
	Arch string `pulumi:"arch"`
	// The family of the operating system.
	Family  string        `pulumi:"family"`
	Filters []GetOSFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the operating system.
	Name string `pulumi:"name"`
}

func GetOSOutput(ctx *pulumi.Context, args GetOSOutputArgs, opts ...pulumi.InvokeOption) GetOSResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOSResult, error) {
			args := v.(GetOSArgs)
			r, err := GetOS(ctx, &args, opts...)
			return *r, err
		}).(GetOSResultOutput)
}

// A collection of arguments for invoking getOS.
type GetOSOutputArgs struct {
	// Query parameters for finding operating systems.
	Filters GetOSFilterArrayInput `pulumi:"filters"`
}

func (GetOSOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOSArgs)(nil)).Elem()
}

// A collection of values returned by getOS.
type GetOSResultOutput struct{ *pulumi.OutputState }

func (GetOSResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOSResult)(nil)).Elem()
}

func (o GetOSResultOutput) ToGetOSResultOutput() GetOSResultOutput {
	return o
}

func (o GetOSResultOutput) ToGetOSResultOutputWithContext(ctx context.Context) GetOSResultOutput {
	return o
}

// The architecture of the operating system.
func (o GetOSResultOutput) Arch() pulumi.StringOutput {
	return o.ApplyT(func(v GetOSResult) string { return v.Arch }).(pulumi.StringOutput)
}

// The family of the operating system.
func (o GetOSResultOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v GetOSResult) string { return v.Family }).(pulumi.StringOutput)
}

func (o GetOSResultOutput) Filters() GetOSFilterArrayOutput {
	return o.ApplyT(func(v GetOSResult) []GetOSFilter { return v.Filters }).(GetOSFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOSResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOSResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the operating system.
func (o GetOSResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOSResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOSResultOutput{})
}