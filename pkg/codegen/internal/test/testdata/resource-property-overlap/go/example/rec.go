// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Rec struct {
	pulumi.CustomResourceState

	Rec RecPtrOutput `pulumi:"rec"`
}

// NewRec registers a new resource with the given unique name, arguments, and options.
func NewRec(ctx *pulumi.Context,
	name string, args *RecArgs, opts ...pulumi.ResourceOption) (*Rec, error) {
	if args == nil {
		args = &RecArgs{}
	}

	var resource Rec
	err := ctx.RegisterResource("example::Rec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRec gets an existing Rec resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecState, opts ...pulumi.ResourceOption) (*Rec, error) {
	var resource Rec
	err := ctx.ReadResource("example::Rec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rec resources.
type recState struct {
}

type RecState struct {
}

func (RecState) ElementType() reflect.Type {
	return reflect.TypeOf((*recState)(nil)).Elem()
}

type recArgs struct {
}

// The set of arguments for constructing a Rec resource.
type RecArgs struct {
}

func (RecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recArgs)(nil)).Elem()
}

type RecInput interface {
	pulumi.Input

	ToRecOutput() RecOutput
	ToRecOutputWithContext(ctx context.Context) RecOutput
}

func (*Rec) ElementType() reflect.Type {
	return reflect.TypeOf((*Rec)(nil))
}

func (i *Rec) ToRecOutput() RecOutput {
	return i.ToRecOutputWithContext(context.Background())
}

func (i *Rec) ToRecOutputWithContext(ctx context.Context) RecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecOutput)
}

func (i *Rec) ToRecPtrOutput() RecPtrOutput {
	return i.ToRecPtrOutputWithContext(context.Background())
}

func (i *Rec) ToRecPtrOutputWithContext(ctx context.Context) RecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecPtrOutput)
}

type RecPtrInput interface {
	pulumi.Input

	ToRecPtrOutput() RecPtrOutput
	ToRecPtrOutputWithContext(ctx context.Context) RecPtrOutput
}

type recPtrType RecArgs

func (*recPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Rec)(nil))
}

func (i *recPtrType) ToRecPtrOutput() RecPtrOutput {
	return i.ToRecPtrOutputWithContext(context.Background())
}

func (i *recPtrType) ToRecPtrOutputWithContext(ctx context.Context) RecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecPtrOutput)
}

// RecArrayInput is an input type that accepts RecArray and RecArrayOutput values.
// You can construct a concrete instance of `RecArrayInput` via:
//
//          RecArray{ RecArgs{...} }
type RecArrayInput interface {
	pulumi.Input

	ToRecArrayOutput() RecArrayOutput
	ToRecArrayOutputWithContext(context.Context) RecArrayOutput
}

type RecArray []RecInput

func (RecArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rec)(nil)).Elem()
}

func (i RecArray) ToRecArrayOutput() RecArrayOutput {
	return i.ToRecArrayOutputWithContext(context.Background())
}

func (i RecArray) ToRecArrayOutputWithContext(ctx context.Context) RecArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecArrayOutput)
}

// RecMapInput is an input type that accepts RecMap and RecMapOutput values.
// You can construct a concrete instance of `RecMapInput` via:
//
//          RecMap{ "key": RecArgs{...} }
type RecMapInput interface {
	pulumi.Input

	ToRecMapOutput() RecMapOutput
	ToRecMapOutputWithContext(context.Context) RecMapOutput
}

type RecMap map[string]RecInput

func (RecMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rec)(nil)).Elem()
}

func (i RecMap) ToRecMapOutput() RecMapOutput {
	return i.ToRecMapOutputWithContext(context.Background())
}

func (i RecMap) ToRecMapOutputWithContext(ctx context.Context) RecMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecMapOutput)
}

type RecOutput struct{ *pulumi.OutputState }

func (RecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rec)(nil))
}

func (o RecOutput) ToRecOutput() RecOutput {
	return o
}

func (o RecOutput) ToRecOutputWithContext(ctx context.Context) RecOutput {
	return o
}

func (o RecOutput) ToRecPtrOutput() RecPtrOutput {
	return o.ToRecPtrOutputWithContext(context.Background())
}

func (o RecOutput) ToRecPtrOutputWithContext(ctx context.Context) RecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Rec) *Rec {
		return &v
	}).(RecPtrOutput)
}

type RecPtrOutput struct{ *pulumi.OutputState }

func (RecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rec)(nil))
}

func (o RecPtrOutput) ToRecPtrOutput() RecPtrOutput {
	return o
}

func (o RecPtrOutput) ToRecPtrOutputWithContext(ctx context.Context) RecPtrOutput {
	return o
}

func (o RecPtrOutput) Elem() RecOutput {
	return o.ApplyT(func(v *Rec) Rec {
		if v != nil {
			return *v
		}
		var ret Rec
		return ret
	}).(RecOutput)
}

type RecArrayOutput struct{ *pulumi.OutputState }

func (RecArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Rec)(nil))
}

func (o RecArrayOutput) ToRecArrayOutput() RecArrayOutput {
	return o
}

func (o RecArrayOutput) ToRecArrayOutputWithContext(ctx context.Context) RecArrayOutput {
	return o
}

func (o RecArrayOutput) Index(i pulumi.IntInput) RecOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Rec {
		return vs[0].([]Rec)[vs[1].(int)]
	}).(RecOutput)
}

type RecMapOutput struct{ *pulumi.OutputState }

func (RecMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Rec)(nil))
}

func (o RecMapOutput) ToRecMapOutput() RecMapOutput {
	return o
}

func (o RecMapOutput) ToRecMapOutputWithContext(ctx context.Context) RecMapOutput {
	return o
}

func (o RecMapOutput) MapIndex(k pulumi.StringInput) RecOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Rec {
		return vs[0].(map[string]Rec)[vs[1].(string)]
	}).(RecOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RecInput)(nil)).Elem(), &Rec{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecPtrInput)(nil)).Elem(), &Rec{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecArrayInput)(nil)).Elem(), RecArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecMapInput)(nil)).Elem(), RecMap{})
	pulumi.RegisterOutputType(RecOutput{})
	pulumi.RegisterOutputType(RecPtrOutput{})
	pulumi.RegisterOutputType(RecArrayOutput{})
	pulumi.RegisterOutputType(RecMapOutput{})
}
