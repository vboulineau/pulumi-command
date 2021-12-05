// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Command extends pulumi.CustomResource {
    /**
     * Get an existing Command resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Command {
        return new Command(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'command:index:Command';

    /**
     * Returns true if the given object is an instance of Command.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Command {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Command.__pulumiType;
    }

    public readonly create!: pulumi.Output<string | undefined>;
    public readonly delete!: pulumi.Output<string | undefined>;
    public readonly dir!: pulumi.Output<string | undefined>;
    public readonly environment!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly interpreter!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly stderr!: pulumi.Output<string>;
    public /*out*/ readonly stdout!: pulumi.Output<string>;
    public readonly update!: pulumi.Output<string | undefined>;

    /**
     * Create a Command resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CommandArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["create"] = args ? args.create : undefined;
            inputs["delete"] = args ? args.delete : undefined;
            inputs["dir"] = args ? args.dir : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["interpreter"] = args ? args.interpreter : undefined;
            inputs["update"] = args ? args.update : undefined;
            inputs["stderr"] = undefined /*out*/;
            inputs["stdout"] = undefined /*out*/;
        } else {
            inputs["create"] = undefined /*out*/;
            inputs["delete"] = undefined /*out*/;
            inputs["dir"] = undefined /*out*/;
            inputs["environment"] = undefined /*out*/;
            inputs["interpreter"] = undefined /*out*/;
            inputs["stderr"] = undefined /*out*/;
            inputs["stdout"] = undefined /*out*/;
            inputs["update"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Command.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Command resource.
 */
export interface CommandArgs {
    /**
     * The command to run on create.
     */
    create?: pulumi.Input<string>;
    /**
     * The command to run on delete.
     */
    delete?: pulumi.Input<string>;
    /**
     * The contents of an SSH key to use for the connection. This takes preference over the password if provided.
     */
    dir?: pulumi.Input<string>;
    /**
     * Environment variables to set on commands.
     */
    environment?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    interpreter?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The command to run on update.
     */
    update?: pulumi.Input<string>;
}
