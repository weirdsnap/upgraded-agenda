// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"strings"

	"github.com/weirdsnap/upgraded-agenda/cli/client"

	"github.com/spf13/cobra"
)

var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "create a new meeting",
	Long: `The createMeeting command help to create a new meeting, user should
	 offer meeting title, meeting participants, start time and end time.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participants, _ := cmd.Flags().GetString("participants")
		participantsArray := strings.Split(participants, " ")
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")

		controller.CreateMeeting(title, participantsArray, start, end)
	},
}

var modifyMeetingCmd = &cobra.Command{
	Use:   "modifyMeeting",
	Short: "add or delete  participants",
	Long: `The modifyMeeting command allow user to add participants or
	delete participants to a specific meeting.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		addStr, _ := cmd.Flags().GetString("add")
		addParticipants := strings.Split(addStr, " ")
		deleteStr, _ := cmd.Flags().GetString("delete")
		deleteParticipants := strings.Split(deleteStr, " ")

		controller.ModifyMeeting(title, addParticipants, deleteParticipants)
	},
}

var queryMeetingCmd = &cobra.Command{
	Use:   "queryMeeting",
	Short: "query meeting during a time space",
	Long: `The queryMeeting command allow user to query meetings during a time
	space, user may be host or participant of these meetings`,
	Run: func(cmd *cobra.Command, args []string) {
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")

		controller.QueryMeeting(start, end)
	},
}

var cancelMeetingCmd = &cobra.Command{
	Use:   "cancelMeeting",
	Short: "cancel the meeting as host",
	Long:  `The cancelMeeting command allow user to cancel a meeting as a host`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")

		controller.CancelMeeting(title)
	},
}

var quitMeetingCmd = &cobra.Command{
	Use:   "quitMeeting",
	Short: "quit the meeting as participant",
	Long:  `The quitMeeting command allow user to quit a meeting as a participant`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")

		controller.QuitMeeting(title)
	},
}

var clearMeetingCmd = &cobra.Command{
	Use:   "clearMeeting",
	Short: "clear all meetings as host",
	Long: `The clearMeeting command allow user to clear all the meeting in which the
	 user is host`,
	Run: func(cmd *cobra.Command, args []string) {

		controller.ClearMeeting()
	},
}

func init() {
	RootCmd.AddCommand(createMeetingCmd)
	RootCmd.AddCommand(modifyMeetingCmd)
	RootCmd.AddCommand(queryMeetingCmd)
	RootCmd.AddCommand(cancelMeetingCmd)
	RootCmd.AddCommand(quitMeetingCmd)
	RootCmd.AddCommand(clearMeetingCmd)

	createMeetingCmd.Flags().StringP("title", "t", "", "the title of the meeting")
	createMeetingCmd.Flags().StringP("participants", "p", "", "participants of the meeting")
	createMeetingCmd.Flags().StringP("start", "s", "", "start time of the meeting")
	createMeetingCmd.Flags().StringP("end", "e", "", "end time of the meeting")

	modifyMeetingCmd.Flags().StringP("title", "t", "", "the title of the meeting")
	modifyMeetingCmd.Flags().StringP("add", "a", "", "add participants")
	modifyMeetingCmd.Flags().StringP("delete", "d", "", "delete participants")

	queryMeetingCmd.Flags().StringP("start", "s", "", "start time of the query")
	queryMeetingCmd.Flags().StringP("end", "e", "", "end time of the query")

	quitMeetingCmd.Flags().StringP("title", "t", "", "the title of the meeting")

	cancelMeetingCmd.Flags().StringP("title", "t", "", "the title of the meeting")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// meetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// meetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
