import com.sap.piper.GenerateDocumentation
import com.sap.piper.Utils

/**
 * Can be used by other steps to unstash files. If a parameter key 'stashContent' is present containing a list
 * of contents to unstash, the 'unstash' List from the commonPipelineEnvironment's 'stageStashes' configuration for
 * the current stage is appended to it.
 *
 * @Returns The complete list of unstashed files.
 */
@GenerateDocumentation
def call(Map parameters = [:]) {

    def toUnstash = []

    handlePipelineStepErrors(stepName: 'unstashStageFiles', stepParameters: parameters) {
        def script = parameters.script
        def stage = parameters.stage
        def utils = parameters.juStabUtils ?: new Utils()

        toUnstash = parameters.stashContent ?: []

        deleteDir()
        toUnstash += script.commonPipelineEnvironment.configuration.stageStashes?.get(stage)?.unstash ?: []
        utils.unstashAll(toUnstash)
    }

    return toUnstash
}