// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vultr Reverse IPv6 resource. This can be used to create, read,
// modify, and delete reverse DNS records for IPv6 addresses. Upon success, DNS
// changes may take 6-12 hours to become active.
//
// ## Example Usage
//
// Create a new reverse DNS record for an IPv6 address:
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
// 		_, err := vultr.NewInstance(ctx, "myServer", &vultr.InstanceArgs{
// 			Plan:       pulumi.String("vc2-1c-1gb"),
// 			Region:     pulumi.String("ewr"),
// 			OsId:       pulumi.Int(167),
// 			EnableIpv6: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vultr.NewReverseIPv6(ctx, "myReverseIpv6", &vultr.ReverseIPv6Args{
// 			InstanceId: pulumi.Any(vultr_server.My_server.Id),
// 			Ip:         pulumi.Any(vultr_server.My_server.V6_networks[0].V6_main_ip),
// 			Reverse:    pulumi.String("host.example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ReverseIPv6 struct {
	pulumi.CustomResourceState

	// The ID of the server you want to set an IPv6
	// reverse DNS record for.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The IPv6 address used in the reverse DNS record.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The hostname used in the IPv6 reverse DNS record.
	Reverse pulumi.StringOutput `pulumi:"reverse"`
}

// NewReverseIPv6 registers a new resource with the given unique name, arguments, and options.
func NewReverseIPv6(ctx *pulumi.Context,
	name string, args *ReverseIPv6Args, opts ...pulumi.ResourceOption) (*ReverseIPv6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.Reverse == nil {
		return nil, errors.New("invalid value for required argument 'Reverse'")
	}
	var resource ReverseIPv6
	err := ctx.RegisterResource("vultr:index/reverseIPv6:ReverseIPv6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReverseIPv6 gets an existing ReverseIPv6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReverseIPv6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReverseIPv6State, opts ...pulumi.ResourceOption) (*ReverseIPv6, error) {
	var resource ReverseIPv6
	err := ctx.ReadResource("vultr:index/reverseIPv6:ReverseIPv6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReverseIPv6 resources.
type reverseIPv6State struct {
	// The ID of the server you want to set an IPv6
	// reverse DNS record for.
	InstanceId *string `pulumi:"instanceId"`
	// The IPv6 address used in the reverse DNS record.
	Ip *string `pulumi:"ip"`
	// The hostname used in the IPv6 reverse DNS record.
	Reverse *string `pulumi:"reverse"`
}

type ReverseIPv6State struct {
	// The ID of the server you want to set an IPv6
	// reverse DNS record for.
	InstanceId pulumi.StringPtrInput
	// The IPv6 address used in the reverse DNS record.
	Ip pulumi.StringPtrInput
	// The hostname used in the IPv6 reverse DNS record.
	Reverse pulumi.StringPtrInput
}

func (ReverseIPv6State) ElementType() reflect.Type {
	return reflect.TypeOf((*reverseIPv6State)(nil)).Elem()
}

type reverseIPv6Args struct {
	// The ID of the server you want to set an IPv6
	// reverse DNS record for.
	InstanceId string `pulumi:"instanceId"`
	// The IPv6 address used in the reverse DNS record.
	Ip string `pulumi:"ip"`
	// The hostname used in the IPv6 reverse DNS record.
	Reverse string `pulumi:"reverse"`
}

// The set of arguments for constructing a ReverseIPv6 resource.
type ReverseIPv6Args struct {
	// The ID of the server you want to set an IPv6
	// reverse DNS record for.
	InstanceId pulumi.StringInput
	// The IPv6 address used in the reverse DNS record.
	Ip pulumi.StringInput
	// The hostname used in the IPv6 reverse DNS record.
	Reverse pulumi.StringInput
}

func (ReverseIPv6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*reverseIPv6Args)(nil)).Elem()
}

type ReverseIPv6Input interface {
	pulumi.Input

	ToReverseIPv6Output() ReverseIPv6Output
	ToReverseIPv6OutputWithContext(ctx context.Context) ReverseIPv6Output
}

func (*ReverseIPv6) ElementType() reflect.Type {
	return reflect.TypeOf((**ReverseIPv6)(nil)).Elem()
}

func (i *ReverseIPv6) ToReverseIPv6Output() ReverseIPv6Output {
	return i.ToReverseIPv6OutputWithContext(context.Background())
}

func (i *ReverseIPv6) ToReverseIPv6OutputWithContext(ctx context.Context) ReverseIPv6Output {
	return pulumi.ToOutputWithContext(ctx, i).(ReverseIPv6Output)
}

// ReverseIPv6ArrayInput is an input type that accepts ReverseIPv6Array and ReverseIPv6ArrayOutput values.
// You can construct a concrete instance of `ReverseIPv6ArrayInput` via:
//
//          ReverseIPv6Array{ ReverseIPv6Args{...} }
type ReverseIPv6ArrayInput interface {
	pulumi.Input

	ToReverseIPv6ArrayOutput() ReverseIPv6ArrayOutput
	ToReverseIPv6ArrayOutputWithContext(context.Context) ReverseIPv6ArrayOutput
}

type ReverseIPv6Array []ReverseIPv6Input

func (ReverseIPv6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReverseIPv6)(nil)).Elem()
}

func (i ReverseIPv6Array) ToReverseIPv6ArrayOutput() ReverseIPv6ArrayOutput {
	return i.ToReverseIPv6ArrayOutputWithContext(context.Background())
}

func (i ReverseIPv6Array) ToReverseIPv6ArrayOutputWithContext(ctx context.Context) ReverseIPv6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReverseIPv6ArrayOutput)
}

// ReverseIPv6MapInput is an input type that accepts ReverseIPv6Map and ReverseIPv6MapOutput values.
// You can construct a concrete instance of `ReverseIPv6MapInput` via:
//
//          ReverseIPv6Map{ "key": ReverseIPv6Args{...} }
type ReverseIPv6MapInput interface {
	pulumi.Input

	ToReverseIPv6MapOutput() ReverseIPv6MapOutput
	ToReverseIPv6MapOutputWithContext(context.Context) ReverseIPv6MapOutput
}

type ReverseIPv6Map map[string]ReverseIPv6Input

func (ReverseIPv6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReverseIPv6)(nil)).Elem()
}

func (i ReverseIPv6Map) ToReverseIPv6MapOutput() ReverseIPv6MapOutput {
	return i.ToReverseIPv6MapOutputWithContext(context.Background())
}

func (i ReverseIPv6Map) ToReverseIPv6MapOutputWithContext(ctx context.Context) ReverseIPv6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReverseIPv6MapOutput)
}

type ReverseIPv6Output struct{ *pulumi.OutputState }

func (ReverseIPv6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**ReverseIPv6)(nil)).Elem()
}

func (o ReverseIPv6Output) ToReverseIPv6Output() ReverseIPv6Output {
	return o
}

func (o ReverseIPv6Output) ToReverseIPv6OutputWithContext(ctx context.Context) ReverseIPv6Output {
	return o
}

type ReverseIPv6ArrayOutput struct{ *pulumi.OutputState }

func (ReverseIPv6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReverseIPv6)(nil)).Elem()
}

func (o ReverseIPv6ArrayOutput) ToReverseIPv6ArrayOutput() ReverseIPv6ArrayOutput {
	return o
}

func (o ReverseIPv6ArrayOutput) ToReverseIPv6ArrayOutputWithContext(ctx context.Context) ReverseIPv6ArrayOutput {
	return o
}

func (o ReverseIPv6ArrayOutput) Index(i pulumi.IntInput) ReverseIPv6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReverseIPv6 {
		return vs[0].([]*ReverseIPv6)[vs[1].(int)]
	}).(ReverseIPv6Output)
}

type ReverseIPv6MapOutput struct{ *pulumi.OutputState }

func (ReverseIPv6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReverseIPv6)(nil)).Elem()
}

func (o ReverseIPv6MapOutput) ToReverseIPv6MapOutput() ReverseIPv6MapOutput {
	return o
}

func (o ReverseIPv6MapOutput) ToReverseIPv6MapOutputWithContext(ctx context.Context) ReverseIPv6MapOutput {
	return o
}

func (o ReverseIPv6MapOutput) MapIndex(k pulumi.StringInput) ReverseIPv6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReverseIPv6 {
		return vs[0].(map[string]*ReverseIPv6)[vs[1].(string)]
	}).(ReverseIPv6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReverseIPv6Input)(nil)).Elem(), &ReverseIPv6{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReverseIPv6ArrayInput)(nil)).Elem(), ReverseIPv6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReverseIPv6MapInput)(nil)).Elem(), ReverseIPv6Map{})
	pulumi.RegisterOutputType(ReverseIPv6Output{})
	pulumi.RegisterOutputType(ReverseIPv6ArrayOutput{})
	pulumi.RegisterOutputType(ReverseIPv6MapOutput{})
}
