package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/spf13/cobra"
)

type checkmarxExecuteScanOptions struct {
	AvoidDuplicateProjectScans    bool   `json:"avoidDuplicateProjectScans,omitempty"`
	FilterPattern                 string `json:"filterPattern,omitempty"`
	FullScanCycle                 string `json:"fullScanCycle,omitempty"`
	FullScansScheduled            bool   `json:"fullScansScheduled,omitempty"`
	GeneratePdfReport             bool   `json:"generatePdfReport,omitempty"`
	Incremental                   bool   `json:"incremental,omitempty"`
	Password                      string `json:"password,omitempty"`
	Preset                        string `json:"preset,omitempty"`
	ProjectName                   string `json:"projectName,omitempty"`
	PullRequestName               string `json:"pullRequestName,omitempty"`
	ServerURL                     string `json:"serverUrl,omitempty"`
	SourceEncoding                string `json:"sourceEncoding,omitempty"`
	TeamID                        string `json:"teamId,omitempty"`
	TeamName                      string `json:"teamName,omitempty"`
	Username                      string `json:"username,omitempty"`
	VulnerabilityThresholdEnabled bool   `json:"vulnerabilityThresholdEnabled,omitempty"`
	VulnerabilityThresholdHigh    int    `json:"vulnerabilityThresholdHigh,omitempty"`
	VulnerabilityThresholdLow     int    `json:"vulnerabilityThresholdLow,omitempty"`
	VulnerabilityThresholdMedium  int    `json:"vulnerabilityThresholdMedium,omitempty"`
	VulnerabilityThresholdResult  string `json:"vulnerabilityThresholdResult,omitempty"`
	VulnerabilityThresholdUnit    string `json:"vulnerabilityThresholdUnit,omitempty"`
	Verbose                       bool   `json:"verbose,omitempty"`
}

type checkmarxExecuteScanInflux struct {
	checkmarx_data struct {
		fields struct {
			high_issues                          string
			high_not_false_postive               string
			high_not_exploitable                 string
			high_confirmed                       string
			high_urgent                          string
			high_proposed_not_exploitable        string
			high_to_verify                       string
			medium_issues                        string
			medium_not_false_postive             string
			medium_not_exploitable               string
			medium_confirmed                     string
			medium_urgent                        string
			medium_proposed_not_exploitable      string
			medium_to_verify                     string
			low_issues                           string
			low_not_false_postive                string
			low_not_exploitable                  string
			low_confirmed                        string
			low_urgent                           string
			low_proposed_not_exploitable         string
			low_to_verify                        string
			information_issues                   string
			information_not_false_postive        string
			information_not_exploitable          string
			information_confirmed                string
			information_urgent                   string
			information_proposed_not_exploitable string
			information_to_verify                string
			initiator_name                       string
			owner                                string
			scan_id                              string
			project_id                           string
			project_name                         string
			team                                 string
			team_full_path_on_report_date        string
			scan_start                           string
			scan_time                            string
			lines_of_code_scanned                string
			files_scanned                        string
			checkmarx_version                    string
			scan_type                            string
			preset                               string
			deep_link                            string
			report_creation_time                 string
		}
		tags struct {
		}
	}
}

func (i *checkmarxExecuteScanInflux) persist(path, resourceName string) {
	measurementContent := []struct {
		measurement string
		valType     string
		name        string
		value       string
	}{
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_issues", value: i.checkmarx_data.fields.high_issues},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_not_false_postive", value: i.checkmarx_data.fields.high_not_false_postive},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_not_exploitable", value: i.checkmarx_data.fields.high_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_confirmed", value: i.checkmarx_data.fields.high_confirmed},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_urgent", value: i.checkmarx_data.fields.high_urgent},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_proposed_not_exploitable", value: i.checkmarx_data.fields.high_proposed_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_to_verify", value: i.checkmarx_data.fields.high_to_verify},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_issues", value: i.checkmarx_data.fields.medium_issues},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_not_false_postive", value: i.checkmarx_data.fields.medium_not_false_postive},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_not_exploitable", value: i.checkmarx_data.fields.medium_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_confirmed", value: i.checkmarx_data.fields.medium_confirmed},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_urgent", value: i.checkmarx_data.fields.medium_urgent},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_proposed_not_exploitable", value: i.checkmarx_data.fields.medium_proposed_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_to_verify", value: i.checkmarx_data.fields.medium_to_verify},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_issues", value: i.checkmarx_data.fields.low_issues},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_not_false_postive", value: i.checkmarx_data.fields.low_not_false_postive},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_not_exploitable", value: i.checkmarx_data.fields.low_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_confirmed", value: i.checkmarx_data.fields.low_confirmed},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_urgent", value: i.checkmarx_data.fields.low_urgent},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_proposed_not_exploitable", value: i.checkmarx_data.fields.low_proposed_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_to_verify", value: i.checkmarx_data.fields.low_to_verify},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_issues", value: i.checkmarx_data.fields.information_issues},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_not_false_postive", value: i.checkmarx_data.fields.information_not_false_postive},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_not_exploitable", value: i.checkmarx_data.fields.information_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_confirmed", value: i.checkmarx_data.fields.information_confirmed},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_urgent", value: i.checkmarx_data.fields.information_urgent},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_proposed_not_exploitable", value: i.checkmarx_data.fields.information_proposed_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_to_verify", value: i.checkmarx_data.fields.information_to_verify},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "initiator_name", value: i.checkmarx_data.fields.initiator_name},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "owner", value: i.checkmarx_data.fields.owner},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "scan_id", value: i.checkmarx_data.fields.scan_id},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "project_id", value: i.checkmarx_data.fields.project_id},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "project_name", value: i.checkmarx_data.fields.project_name},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "team", value: i.checkmarx_data.fields.team},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "team_full_path_on_report_date", value: i.checkmarx_data.fields.team_full_path_on_report_date},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "scan_start", value: i.checkmarx_data.fields.scan_start},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "scan_time", value: i.checkmarx_data.fields.scan_time},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "lines_of_code_scanned", value: i.checkmarx_data.fields.lines_of_code_scanned},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "files_scanned", value: i.checkmarx_data.fields.files_scanned},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "checkmarx_version", value: i.checkmarx_data.fields.checkmarx_version},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "scan_type", value: i.checkmarx_data.fields.scan_type},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "preset", value: i.checkmarx_data.fields.preset},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "deep_link", value: i.checkmarx_data.fields.deep_link},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "report_creation_time", value: i.checkmarx_data.fields.report_creation_time},
	}

	errCount := 0
	for _, metric := range measurementContent {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(metric.measurement, fmt.Sprintf("%vs", metric.valType), metric.name), metric.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting influx environment.")
			errCount++
		}
	}
	if errCount > 0 {
		os.Exit(1)
	}
}

var myCheckmarxExecuteScanOptions checkmarxExecuteScanOptions

// CheckmarxExecuteScanCommand Checkmarx is the recommended tool for security scans of JavaScript, iOS, Swift and Ruby code.
func CheckmarxExecuteScanCommand() *cobra.Command {
	metadata := checkmarxExecuteScanMetadata()
	var influx checkmarxExecuteScanInflux

	var createCheckmarxExecuteScanCmd = &cobra.Command{
		Use:   "checkmarxExecuteScan",
		Short: "Checkmarx is the recommended tool for security scans of JavaScript, iOS, Swift and Ruby code.",
		Long: `Checkmarx is a Static Application Security Testing (SAST) tool to analyze i.e. Java- or TypeScript, Swift, Golang, Ruby code,
and many other programming languages for security flaws based on a set of provided rules/queries that can be customized and extended.

This step by default enforces a specific audit baseline for findings and therefore ensures that:
* No 'To Verify' High and Medium issues exist in your project
* Total number of High and Medium 'Confirmed' or 'Urgent' issues is zero
* 10% of all Low issues are 'Confirmed' or 'Not Exploitable'

You can adapt above thresholds specifically using the provided configuration parameters and i.e. check for ` + "`" + `absolute` + "`" + `
thresholds instead of ` + "`" + `percentage` + "`" + ` whereas we strongly recommend you to stay with the defaults provided.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			log.SetStepName("checkmarxExecuteScan")
			log.SetVerbose(GeneralConfig.Verbose)
			return PrepareConfig(cmd, &metadata, "checkmarxExecuteScan", &myCheckmarxExecuteScanOptions, config.OpenPiperFile)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			handler := func() {
				influx.persist(GeneralConfig.EnvRootPath, "influx")
			}
			log.DeferExitHandler(handler)
			defer handler()
			return checkmarxExecuteScan(myCheckmarxExecuteScanOptions, &influx)
		},
	}

	addCheckmarxExecuteScanFlags(createCheckmarxExecuteScanCmd)
	return createCheckmarxExecuteScanCmd
}

func addCheckmarxExecuteScanFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&myCheckmarxExecuteScanOptions.AvoidDuplicateProjectScans, "avoidDuplicateProjectScans", false, "Whether duplicate scans of the same project state shall be avoided or not")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.FilterPattern, "filterPattern", "!**/node_modules/**, !**/.xmake/**, !**/*_test.go, !**/vendor/**/*.go, **/*.html, **/*.xml, **/*.go, **/*.py, **/*.js, **/*.scala, **/*.ts", "The filter pattern used to zip the files relevant for scanning, patterns can be negated by setting an exclamation mark in front i.e. `!test/*.js` would avoid adding any javascript files located in the test directory")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.FullScanCycle, "fullScanCycle", "5", "Indicates how often a full scan should happen between the incremental scans when activated")
	cmd.Flags().BoolVar(&myCheckmarxExecuteScanOptions.FullScansScheduled, "fullScansScheduled", true, "Whether full scans are to be scheduled or not. Should be used in relation with `incremental` and `fullScanCycle`")
	cmd.Flags().BoolVar(&myCheckmarxExecuteScanOptions.GeneratePdfReport, "generatePdfReport", true, "Whether to generate a PDF report of the analysis results or not")
	cmd.Flags().BoolVar(&myCheckmarxExecuteScanOptions.Incremental, "incremental", true, "Whether incremental scans are to be applied which optimizes the scan time but might reduce detection capabilities. Therefore full scans are still required from time to time and should be scheduled via `fullScansScheduled` and `fullScanCycle`")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.Password, "password", os.Getenv("PIPER_password"), "The password to authenticate")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.Preset, "preset", os.Getenv("PIPER_preset"), "The preset to use for scanning, if not set explicitly the step will attempt to look up the project's setting based on the availability of `checkmarxCredentialsId`")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.ProjectName, "projectName", os.Getenv("PIPER_projectName"), "The name of the Checkmarx project to scan into")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.PullRequestName, "pullRequestName", os.Getenv("PIPER_pullRequestName"), "Used to supply the name for the newly created PR project branch when being used in pull request scenarios")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.ServerURL, "serverUrl", os.Getenv("PIPER_serverUrl"), "The URL pointing to the root of the Checkmarx server to be used")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.SourceEncoding, "sourceEncoding", "1", "The source encoding to be used, if not set explicitly the project's default will be used")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.TeamID, "teamId", os.Getenv("PIPER_teamId"), "The group ID related to your team which can be obtained via the Pipeline Syntax plugin as described in the `Details` section")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.TeamName, "teamName", os.Getenv("PIPER_teamName"), "The full name of the team to assign newly created projects to which is preferred to teamId")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.Username, "username", os.Getenv("PIPER_username"), "The username to authenticate")
	cmd.Flags().BoolVar(&myCheckmarxExecuteScanOptions.VulnerabilityThresholdEnabled, "vulnerabilityThresholdEnabled", true, "Whether the thresholds are enabled or not. If enabled the build will be set to `vulnerabilityThresholdResult` in case a specific threshold value is exceeded")
	cmd.Flags().IntVar(&myCheckmarxExecuteScanOptions.VulnerabilityThresholdHigh, "vulnerabilityThresholdHigh", 100, "The specific threshold for high severity findings")
	cmd.Flags().IntVar(&myCheckmarxExecuteScanOptions.VulnerabilityThresholdLow, "vulnerabilityThresholdLow", 10, "The specific threshold for low severity findings")
	cmd.Flags().IntVar(&myCheckmarxExecuteScanOptions.VulnerabilityThresholdMedium, "vulnerabilityThresholdMedium", 100, "The specific threshold for medium severity findings")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.VulnerabilityThresholdResult, "vulnerabilityThresholdResult", "FAILURE", "The result of the build in case thresholds are enabled and exceeded")
	cmd.Flags().StringVar(&myCheckmarxExecuteScanOptions.VulnerabilityThresholdUnit, "vulnerabilityThresholdUnit", "percentage", "The unit for the threshold to apply.")
	cmd.Flags().BoolVar(&myCheckmarxExecuteScanOptions.Verbose, "verbose", false, "Whether the step shall provide verbose logging output")

	cmd.MarkFlagRequired("password")
	cmd.MarkFlagRequired("projectName")
	cmd.MarkFlagRequired("serverUrl")
	cmd.MarkFlagRequired("username")
}

// retrieve step metadata
func checkmarxExecuteScanMetadata() config.StepData {
	var theMetaData = config.StepData{
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "avoidDuplicateProjectScans",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "filterPattern",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "fullScanCycle",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "fullScansScheduled",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "generatePdfReport",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "incremental",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "password",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "preset",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "projectName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "checkmarxProject"}},
					},
					{
						Name:        "pullRequestName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "serverUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "checkmarxServerUrl"}},
					},
					{
						Name:        "sourceEncoding",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "teamId",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "checkmarxGroupId"}},
					},
					{
						Name:        "teamName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "username",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "vulnerabilityThresholdEnabled",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "vulnerabilityThresholdHigh",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "vulnerabilityThresholdLow",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "vulnerabilityThresholdMedium",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "vulnerabilityThresholdResult",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "vulnerabilityThresholdUnit",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "verbose",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
