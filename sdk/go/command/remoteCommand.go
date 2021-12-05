// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package command

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RemoteCommand struct {
	pulumi.CustomResourceState

	Connection  RemoteConnectionPtrOutput `pulumi:"connection"`
	Create      pulumi.StringPtrOutput    `pulumi:"create"`
	Delete      pulumi.StringPtrOutput    `pulumi:"delete"`
	Environment pulumi.StringMapOutput    `pulumi:"environment"`
	Stderr      pulumi.StringOutput       `pulumi:"stderr"`
	Stdout      pulumi.StringOutput       `pulumi:"stdout"`
	Update      pulumi.StringPtrOutput    `pulumi:"update"`
}

// NewRemoteCommand registers a new resource with the given unique name, arguments, and options.
func NewRemoteCommand(ctx *pulumi.Context,
	name string, args *RemoteCommandArgs, opts ...pulumi.ResourceOption) (*RemoteCommand, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	var resource RemoteCommand
	err := ctx.RegisterResource("command:index:RemoteCommand", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemoteCommand gets an existing RemoteCommand resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemoteCommand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteCommandState, opts ...pulumi.ResourceOption) (*RemoteCommand, error) {
	var resource RemoteCommand
	err := ctx.ReadResource("command:index:RemoteCommand", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemoteCommand resources.
type remoteCommandState struct {
}

type RemoteCommandState struct {
}

func (RemoteCommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteCommandState)(nil)).Elem()
}

type remoteCommandArgs struct {
	Connection  RemoteConnection  `pulumi:"connection"`
	Create      *string           `pulumi:"create"`
	Delete      *string           `pulumi:"delete"`
	Environment map[string]string `pulumi:"environment"`
	Update      *string           `pulumi:"update"`
}

// The set of arguments for constructing a RemoteCommand resource.
type RemoteCommandArgs struct {
	Connection  RemoteConnectionInput
	Create      pulumi.StringPtrInput
	Delete      pulumi.StringPtrInput
	Environment pulumi.StringMapInput
	Update      pulumi.StringPtrInput
}

func (RemoteCommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteCommandArgs)(nil)).Elem()
}

type RemoteCommandInput interface {
	pulumi.Input

	ToRemoteCommandOutput() RemoteCommandOutput
	ToRemoteCommandOutputWithContext(ctx context.Context) RemoteCommandOutput
}

func (*RemoteCommand) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteCommand)(nil))
}

func (i *RemoteCommand) ToRemoteCommandOutput() RemoteCommandOutput {
	return i.ToRemoteCommandOutputWithContext(context.Background())
}

func (i *RemoteCommand) ToRemoteCommandOutputWithContext(ctx context.Context) RemoteCommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteCommandOutput)
}

func (i *RemoteCommand) ToRemoteCommandPtrOutput() RemoteCommandPtrOutput {
	return i.ToRemoteCommandPtrOutputWithContext(context.Background())
}

func (i *RemoteCommand) ToRemoteCommandPtrOutputWithContext(ctx context.Context) RemoteCommandPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteCommandPtrOutput)
}

type RemoteCommandPtrInput interface {
	pulumi.Input

	ToRemoteCommandPtrOutput() RemoteCommandPtrOutput
	ToRemoteCommandPtrOutputWithContext(ctx context.Context) RemoteCommandPtrOutput
}

type remoteCommandPtrType RemoteCommandArgs

func (*remoteCommandPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteCommand)(nil))
}

func (i *remoteCommandPtrType) ToRemoteCommandPtrOutput() RemoteCommandPtrOutput {
	return i.ToRemoteCommandPtrOutputWithContext(context.Background())
}

func (i *remoteCommandPtrType) ToRemoteCommandPtrOutputWithContext(ctx context.Context) RemoteCommandPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteCommandPtrOutput)
}

// RemoteCommandArrayInput is an input type that accepts RemoteCommandArray and RemoteCommandArrayOutput values.
// You can construct a concrete instance of `RemoteCommandArrayInput` via:
//
//          RemoteCommandArray{ RemoteCommandArgs{...} }
type RemoteCommandArrayInput interface {
	pulumi.Input

	ToRemoteCommandArrayOutput() RemoteCommandArrayOutput
	ToRemoteCommandArrayOutputWithContext(context.Context) RemoteCommandArrayOutput
}

type RemoteCommandArray []RemoteCommandInput

func (RemoteCommandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteCommand)(nil)).Elem()
}

func (i RemoteCommandArray) ToRemoteCommandArrayOutput() RemoteCommandArrayOutput {
	return i.ToRemoteCommandArrayOutputWithContext(context.Background())
}

func (i RemoteCommandArray) ToRemoteCommandArrayOutputWithContext(ctx context.Context) RemoteCommandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteCommandArrayOutput)
}

// RemoteCommandMapInput is an input type that accepts RemoteCommandMap and RemoteCommandMapOutput values.
// You can construct a concrete instance of `RemoteCommandMapInput` via:
//
//          RemoteCommandMap{ "key": RemoteCommandArgs{...} }
type RemoteCommandMapInput interface {
	pulumi.Input

	ToRemoteCommandMapOutput() RemoteCommandMapOutput
	ToRemoteCommandMapOutputWithContext(context.Context) RemoteCommandMapOutput
}

type RemoteCommandMap map[string]RemoteCommandInput

func (RemoteCommandMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteCommand)(nil)).Elem()
}

func (i RemoteCommandMap) ToRemoteCommandMapOutput() RemoteCommandMapOutput {
	return i.ToRemoteCommandMapOutputWithContext(context.Background())
}

func (i RemoteCommandMap) ToRemoteCommandMapOutputWithContext(ctx context.Context) RemoteCommandMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteCommandMapOutput)
}

type RemoteCommandOutput struct {
	*pulumi.OutputState
}

func (RemoteCommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteCommand)(nil))
}

func (o RemoteCommandOutput) ToRemoteCommandOutput() RemoteCommandOutput {
	return o
}

func (o RemoteCommandOutput) ToRemoteCommandOutputWithContext(ctx context.Context) RemoteCommandOutput {
	return o
}

func (o RemoteCommandOutput) ToRemoteCommandPtrOutput() RemoteCommandPtrOutput {
	return o.ToRemoteCommandPtrOutputWithContext(context.Background())
}

func (o RemoteCommandOutput) ToRemoteCommandPtrOutputWithContext(ctx context.Context) RemoteCommandPtrOutput {
	return o.ApplyT(func(v RemoteCommand) *RemoteCommand {
		return &v
	}).(RemoteCommandPtrOutput)
}

type RemoteCommandPtrOutput struct {
	*pulumi.OutputState
}

func (RemoteCommandPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteCommand)(nil))
}

func (o RemoteCommandPtrOutput) ToRemoteCommandPtrOutput() RemoteCommandPtrOutput {
	return o
}

func (o RemoteCommandPtrOutput) ToRemoteCommandPtrOutputWithContext(ctx context.Context) RemoteCommandPtrOutput {
	return o
}

type RemoteCommandArrayOutput struct{ *pulumi.OutputState }

func (RemoteCommandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RemoteCommand)(nil))
}

func (o RemoteCommandArrayOutput) ToRemoteCommandArrayOutput() RemoteCommandArrayOutput {
	return o
}

func (o RemoteCommandArrayOutput) ToRemoteCommandArrayOutputWithContext(ctx context.Context) RemoteCommandArrayOutput {
	return o
}

func (o RemoteCommandArrayOutput) Index(i pulumi.IntInput) RemoteCommandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RemoteCommand {
		return vs[0].([]RemoteCommand)[vs[1].(int)]
	}).(RemoteCommandOutput)
}

type RemoteCommandMapOutput struct{ *pulumi.OutputState }

func (RemoteCommandMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RemoteCommand)(nil))
}

func (o RemoteCommandMapOutput) ToRemoteCommandMapOutput() RemoteCommandMapOutput {
	return o
}

func (o RemoteCommandMapOutput) ToRemoteCommandMapOutputWithContext(ctx context.Context) RemoteCommandMapOutput {
	return o
}

func (o RemoteCommandMapOutput) MapIndex(k pulumi.StringInput) RemoteCommandOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RemoteCommand {
		return vs[0].(map[string]RemoteCommand)[vs[1].(string)]
	}).(RemoteCommandOutput)
}

func init() {
	pulumi.RegisterOutputType(RemoteCommandOutput{})
	pulumi.RegisterOutputType(RemoteCommandPtrOutput{})
	pulumi.RegisterOutputType(RemoteCommandArrayOutput{})
	pulumi.RegisterOutputType(RemoteCommandMapOutput{})
}
