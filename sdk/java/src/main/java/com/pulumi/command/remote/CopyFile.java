// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.command.remote;

import com.pulumi.command.Utilities;
import com.pulumi.command.remote.CopyFileArgs;
import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Copy a local file to a remote host.
 * 
 */
@ResourceType(type="command:remote:CopyFile")
public class CopyFile extends com.pulumi.resources.CustomResource {
    /**
     * The parameters with which to connect to the remote host.
     * 
     */
    @Export(name="connection", type=Connection.class, parameters={})
    private Output<Connection> connection;

    /**
     * @return The parameters with which to connect to the remote host.
     * 
     */
    public Output<Connection> connection() {
        return this.connection;
    }
    /**
     * The path of the file to be copied.
     * 
     */
    @Export(name="localPath", type=String.class, parameters={})
    private Output<String> localPath;

    /**
     * @return The path of the file to be copied.
     * 
     */
    public Output<String> localPath() {
        return this.localPath;
    }
    /**
     * The destination path in the remote host.
     * 
     */
    @Export(name="remotePath", type=String.class, parameters={})
    private Output<String> remotePath;

    /**
     * @return The destination path in the remote host.
     * 
     */
    public Output<String> remotePath() {
        return this.remotePath;
    }
    /**
     * Trigger replacements on changes to this input.
     * 
     */
    @Export(name="triggers", type=List.class, parameters={Object.class})
    private Output</* @Nullable */ List<Object>> triggers;

    /**
     * @return Trigger replacements on changes to this input.
     * 
     */
    public Output<Optional<List<Object>>> triggers() {
        return Codegen.optional(this.triggers);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CopyFile(String name) {
        this(name, CopyFileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CopyFile(String name, CopyFileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CopyFile(String name, CopyFileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("command:remote:CopyFile", name, args == null ? CopyFileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CopyFile(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("command:remote:CopyFile", name, null, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static CopyFile get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CopyFile(name, id, options);
    }
}
