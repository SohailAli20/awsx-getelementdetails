package RDS

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/Appkube-awsx/awsx-getelementdetails/comman-function"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/spf13/cobra"
)

// type MemoryUsage struct {
// 	Timestamp time.Time
// 	Value     float64
// }

var AwsxRDSFreeableMemoryCmd = &cobra.Command{
	Use:   "freeable_memory_panel",
	Short: "get freeable memory metrics data",
	Long:  `command to get freeable memory metrics data`,

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
			jsonResp, cloudwatchMetricResp, err := GetRDSFreeableMemoryPanel(cmd, clientAuth, nil)
			if err != nil {
				log.Println("Error getting freeable memory data: ", err)
				return
			}
			if responseType == "frame" {
				fmt.Println(cloudwatchMetricResp)
			} else {
				// default case. it prints json
				fmt.Println(jsonResp)
			}
		}

	},
}

func GetRDSFreeableMemoryPanel(cmd *cobra.Command, clientAuth *model.Auth, cloudWatchClient *cloudwatch.CloudWatch) (string, map[string]*cloudwatch.GetMetricDataOutput, error) {

	elementType, _ := cmd.PersistentFlags().GetString("elementType")
	fmt.Println(elementType)
	instanceId, _ := cmd.PersistentFlags().GetString("instanceId")
	startTime, endTime, err := comman_function.ParseTimes(cmd)

	if err != nil {
		return "", nil, fmt.Errorf("error parsing time: %v", err)
	}
	instanceId, err = comman_function.GetCmdbData(cmd)

	if err != nil {
		return "", nil, fmt.Errorf("error getting instance ID: %v", err)
	}

	cloudwatchMetricData := map[string]*cloudwatch.GetMetricDataOutput{}

	rawData, err := comman_function.GetMetricData(clientAuth, instanceId, "AWS/RDS", "FreeableMemory", startTime, endTime, "Average", "DBInstanceIdentifier", cloudWatchClient)
	if err != nil {
		log.Println("Error in getting freeable memory data: ", err)
		return "", nil, err
	}
	cloudwatchMetricData["FreeableMemory"] = rawData

	return "", cloudwatchMetricData, nil
}

// func processedRawMemoryData(result *cloudwatch.GetMetricDataOutput) []MemoryUsage {
// 	var processedData []MemoryUsage

// 	for i, timestamp := range result.MetricDataResults[0].Timestamps {
// 		value := *result.MetricDataResults[0].Values[i]
// 		processedData = append(processedData, MemoryUsage{
// 			Timestamp: *timestamp,
// 			Value:     value,
// 		})
// 	}

// 	return processedData
// }

func init() {
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("query", "", "query")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("endTime", "", "endcl time")
	AwsxRDSFreeableMemoryCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
