package NLB

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/model"
	 "github.com/Appkube-awsx/awsx-getelementdetails/comman-function"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/spf13/cobra"
)

var AwsxTargetDeregistrationsCmd = &cobra.Command{
	Use:   "target_deregistration_panel",
	Short: "Get target deregistration panel",
	Long:  `Command to retrieve target deregistration panel`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running target deregistration command")

		var authFlag bool
		var clientAuth *model.Auth
		var err error
		authFlag, clientAuth, err = authenticate.AuthenticateCommand(cmd)

		if err != nil {
			log.Printf("Error during authentication: %v\n", err)
			err := cmd.Help()
			if err != nil {
				return
			}
			return
		}
		if authFlag {
			panel, err := GetTargetDeregistrationspanel(cmd, clientAuth, nil)
			if err != nil {
				return
			}
			fmt.Println(panel)

		}
	},
}

func GetTargetDeregistrationspanel(cmd *cobra.Command, clientAuth *model.Auth, cloudWatchLogs *cloudwatchlogs.CloudWatchLogs) ([]*cloudwatchlogs.GetQueryResultsOutput, error) {
	logGroupName, _ := cmd.PersistentFlags().GetString("logGroupName")
	startTime, endTime, err := comman_function.ParseTimes(cmd)

		if err != nil {
			return nil, fmt.Errorf("Error parsing time: %v", err)
		}
		logGroupName, err = comman_function.GetCmdbLogsData(cmd)

	
		if err != nil {
			return nil, fmt.Errorf("error getting instance ID: %v", err)
		}
		results, err := comman_function.GetLogsData(clientAuth, startTime, endTime, logGroupName, `fields @timestamp| filter eventSource=="elasticloadbalancing.amazonaws.com"| filter eventName=="DeregisterTargets"| stats count(*) as DeregistrationTargetCount by @timestamp| sort @timestamp desc`, cloudWatchLogs)
		
	if err != nil {
		return nil, nil
	}
	processedResults := comman_function.ProcessQueryResult(results)

	return processedResults, nil
}



	
	

func init() {
	comman_function.InitAwsCmdFlags(AwsxTargetDeregistrationsCmd)
}
