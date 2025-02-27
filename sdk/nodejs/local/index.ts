// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CommandArgs } from "./command";
export type Command = import("./command").Command;
export const Command: typeof import("./command").Command = null as any;
utilities.lazyLoad(exports, ["Command"], () => require("./command"));

export { RunArgs, RunResult, RunOutputArgs } from "./run";
export const run: typeof import("./run").run = null as any;
export const runOutput: typeof import("./run").runOutput = null as any;
utilities.lazyLoad(exports, ["run","runOutput"], () => require("./run"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "command:local:Command":
                return new Command(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("command", "local", _module)
