package EC2

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-getelementdetails/comman-function"
	"log"
	"time"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/model"
	// "github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type AlarmNotification struct {
	Timestamp   time.Time
	Alert       string
	Description string
}

var AwsxEc2AlarmandNotificationcmd = &cobra.Command{
	Use:   "alerts_and_notifications_panel",
	Short: "Retrieve recent alerts and notifications related to EC2 instance availability",
	Long:  `Command to retrieve recent alerts and notifications related to EC2 instance availability`,

	Run: func(cmd *cobra.Command, args []string) {
		authFlag, clientAuth, err := handleAuth(cmd)
		if err != nil {
			log.Println("Error during authentication:", err)
			return
		}

		if authFlag {
			responseType, _ := cmd.PersistentFlags().GetString("responseType")
			notifications, err := GetAlertsAndNotificationsPanel(cmd, clientAuth)
			if err != nil {
				log.Println("Error getting alerts and notifications:", err)
				return
			}

			if responseType == "frame" {
				fmt.Println(notifications)
			} else {
				//printTable(notifications)
			}
		}
	},
}

func handleAuth(cmd *cobra.Command) (bool, *model.Auth, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateCommand(cmd)
	if err != nil {
		return false, nil, err
	}
	return authFlag, clientAuth, nil
}

func GetAlertsAndNotificationsPanel(cmd *cobra.Command, clientAuth *model.Auth) ([]AlarmNotification, error) {
	startTime, endTime, err := comman_function.ParseTimes(cmd)
	if err != nil {
		return nil, fmt.Errorf("error parsing time: %v", err)
	}

	alarms, err := comman_function.GetCloudWatchAlarms(clientAuth, startTime, endTime)
	if err != nil {
		log.Println("Error getting CloudWatch alarms:", err)
		return nil, err
	}

	notifications := make([]AlarmNotification, len(alarms))
	for i, alarm := range alarms {
		notifications[i] = AlarmNotification{
			Timestamp:   *alarm.StateUpdatedTimestamp,
			Alert:       *alarm.StateReason,
			Description: *alarm.AlarmDescription,
		}
	}

	return notifications, nil
}

//func printTable(notifications []AlarmNotification) {
//	table := tablewriter.NewWriter(os.Stdout)
//	table.SetHeader([]string{"Timestamp", "Alert", "Description"})
//
//	for _, notification := range notifications {
//		table.Append([]string{
//			notification.Timestamp.Format(time.RFC3339),
//			notification.Alert,
//			notification.Description,
//		})
//	}
//
//	table.Render()
//}

func init() {
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("rootvolumeId", "", "root volume id")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("ebsvolume1Id", "", "ebs volume 1 id")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("ebsvolume2Id", "", "ebs volume 2 id")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("elementId", "", "element id")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("ServiceName", "", "Service Name")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("elementType", "", "element type")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("clusterName", "", "cluster name")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("query", "", "query")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("startTime", "", "start time")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("endTime", "", "endcl time")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("responseType", "", "response type. json/frame")
	AwsxEc2AlarmandNotificationcmd.PersistentFlags().String("logGroupName", "", "log group name")
}
