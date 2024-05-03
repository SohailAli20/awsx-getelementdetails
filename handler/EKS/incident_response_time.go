package EKS

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/Appkube-awsx/awsx-getelementdetails/comman-function"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/spf13/cobra"
)

// type IncidentResponseResult struct {
// 	RawData []struct {
// 		Timestamp time.Time
// 		Value     float64
// 	} `json:"incident response time"`
// }

var AwsxEKSIncidentResponseTimeCmd = &cobra.Command{
	Use:   "incident_response_time_panel",
	Short: "get incident response time metrics data",
	Long:  `command to get incident response time metrics data`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running from child command")
		var authFlag, clientAuth, err = authenticate.AuthenticateCommand(cmd)
		if err != nil {
			log.Printf("Error during authentication: %v\n", err)
			err := cmd.Help()
			if err != nil {

				return
			}
			return
		}
		if authFlag {

			responseType, _ := cmd.PersistentFlags().GetString("responseType")
			jsonResp, cloudwatchMetricResp, err := GetIncidentResponseTimeData(cmd, clientAuth, nil)
			if err != nil {
				log.Println("Error getting incident response time data: ", err)
				return
			}
			if responseType == "frame" {
				fmt.Println(cloudwatchMetricResp)
			} else {
				fmt.Println(jsonResp)
			}
		}
	},
}

func GetIncidentResponseTimeData(cmd *cobra.Command, clientAuth *model.Auth, cloudWatchClient *cloudwatch.CloudWatch) (string, map[string]*cloudwatch.GetMetricDataOutput, error) {

	instanceId, _ := cmd.PersistentFlags().GetString("instanceId")
	elementType, _ := cmd.PersistentFlags().GetString("elementType")
	fmt.Println(elementType)

	startTime, endTime, err := comman_function.ParseTimes(cmd)
	if err != nil {
		return "", nil, fmt.Errorf("error parsing time: %v", err)
	}

	instanceId, err = comman_function.GetCmdbData(cmd)
	if err != nil {
		return "", nil, fmt.Errorf("error getting instance ID: %v", err)
	}

	cloudwatchMetricData := map[string]*cloudwatch.GetMetricDataOutput{}
	rawData, err := comman_function.GetMetricData(clientAuth, instanceId, "ContainerInsights", "pod_status_failed", startTime, endTime, "Average", "ClusterName", cloudWatchClient)
	if err != nil {
		log.Println("Error fetching total operations raw data: ", err)
		return "", nil, err
	}

	cloudwatchMetricData["CPU_Nice"] = rawData

	return "", cloudwatchMetricData, nil
}

func init() {
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("query", "", "query")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("endTime", "", "endcl time")
	AwsxEKSIncidentResponseTimeCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
