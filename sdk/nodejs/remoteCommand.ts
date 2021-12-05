// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class RemoteCommand extends pulumi.CustomResource {
    /**
     * Get an existing RemoteCommand resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RemoteCommand {
        return new RemoteCommand(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'command:index:RemoteCommand';

    /**
     * Returns true if the given object is an instance of RemoteCommand.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RemoteCommand {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RemoteCommand.__pulumiType;
    }

    public readonly connection!: pulumi.Output<outputs.RemoteConnection | undefined>;
    public readonly create!: pulumi.Output<string | undefined>;
    public readonly delete!: pulumi.Output<string | undefined>;
    public readonly environment!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly stderr!: pulumi.Output<string>;
    public /*out*/ readonly stdout!: pulumi.Output<string>;
    public readonly update!: pulumi.Output<string | undefined>;

    /**
     * Create a RemoteCommand resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RemoteCommandArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            inputs["connection"] = args ? args.connection : undefined;
            inputs["create"] = args ? args.create : undefined;
            inputs["delete"] = args ? args.delete : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["update"] = args ? args.update : undefined;
            inputs["stderr"] = undefined /*out*/;
            inputs["stdout"] = undefined /*out*/;
        } else {
            inputs["connection"] = undefined /*out*/;
            inputs["create"] = undefined /*out*/;
            inputs["delete"] = undefined /*out*/;
            inputs["environment"] = undefined /*out*/;
            inputs["stderr"] = undefined /*out*/;
            inputs["stdout"] = undefined /*out*/;
            inputs["update"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RemoteCommand.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RemoteCommand resource.
 */
export interface RemoteCommandArgs {
    connection: pulumi.Input<inputs.RemoteConnectionArgs>;
    create?: pulumi.Input<string>;
    delete?: pulumi.Input<string>;
    environment?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    update?: pulumi.Input<string>;
}
