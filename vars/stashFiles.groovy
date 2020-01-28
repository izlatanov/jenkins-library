import com.sap.piper.GenerateDocumentation
import com.sap.piper.Utils

/**
 * Can be used by other steps to stash files. Uses the stashing configuration for the calling stage.
 */
@GenerateDocumentation
def call(Map parameters = [:]) {
    handlePipelineStepErrors(stepName: 'stashFiles', stepParameters: parameters) {
        def script = parameters.script
        def stage = parameters.stage
        def utils = parameters.juStabUtils ?: new Utils()

        List stashes = script.commonPipelineEnvironment.configuration.stageStashes?.get(stage)?.stashes ?: []

        utils.stashList(script, stashes)

        //NOTE: We do not delete the directory in case Jenkins runs on Kubernetes.
        // deleteDir() is not required in pods, but would be nice to have the same behaviour and leave a clean fileSystem.
        if (!isInsidePod(script)) {
            deleteDir()
        }
    }
}

private boolean isInsidePod(Script script){
    return script.env.POD_NAME
}
