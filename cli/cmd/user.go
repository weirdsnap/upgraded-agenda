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
	"upgraded-agenda/cli/client"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "user register",
	Long: `register command is used to register a new user, you are required to offer
	username, password, email and telephone number.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		telephone, _ := cmd.Flags().GetString("telephone")

		controller.Register(username, password, email, telephone)
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "user login",
	Long:  "login command is used to login, user are required to offer username and password",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		controller.Login(username, password)
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "user logout",
	Long:  "logout command is used to logout a user.",
	Run: func(cmd *cobra.Command, args []string) {
		controller.Logout()
	},
}

var listUserCmd = &cobra.Command{
	Use:   "listUser",
	Short: "list all the registered users",
	Long:  "listUser command help user to see all the registered user.",
	Run: func(cmd *cobra.Command, args []string) {
		controller.ListUser()
	},
}

var deleteUserCmd = &cobra.Command{
	Use:   "deleteUser",
	Short: "delete current user",
	Long:  "deleteUser command is used to delete the current user.",
	Run: func(cmd *cobra.Command, args []string) {
		controller.DeleteUser()
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	RootCmd.AddCommand(loginCmd)
	RootCmd.AddCommand(logoutCmd)
	RootCmd.AddCommand(listUserCmd)
	RootCmd.AddCommand(deleteUserCmd)

	registerCmd.Flags().StringP("username", "u", "", "username to register")
	registerCmd.Flags().StringP("password", "p", "", "password of the user")
	registerCmd.Flags().StringP("email", "e", "", "email address of user")
	registerCmd.Flags().StringP("telephone", "t", "", "telephone number of the user")

	loginCmd.Flags().StringP("username", "u", "", "username to login")
	loginCmd.Flags().StringP("password", "p", "", "password to login")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
