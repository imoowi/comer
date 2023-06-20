/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"{{.moduleName}}/apps/user/models"
	"{{.moduleName}}/apps/user/services"
	"{{.moduleName}}/global"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		initDb()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initDb() {
	log.Println(`init start.`)
	global.Bootstrap()

	// var c *gin.Context
	c := gin.CreateTestContextOnly(nil, gin.New())
	c.Set(`isInit`, true)
	// 添加超管角色
	superRoleName := `超级管理员`
	var superRoleId uint
	role, err := services.Role.OneByName(c, superRoleName)
	if err != nil || role.ID <= 0 {
		_role := &models.Role{Name: `超级管理员`, Level: models.ROLE_LEVEL_SUPER}
		roleId, err := services.Role.Add(c, _role)
		if err != nil {
			log.Println(err.Error())
			return
		} else {
			superRoleId = roleId
		}
	} else {
		superRoleId = role.ID
	}

	// 添加超管
	if superRoleId > 0 {
		userAdd := &models.UserAdd{Username: `root`, Passwd: `root`, RoleId: superRoleId}
		_, err := services.User.Add(c, userAdd)
		if err != nil {
			log.Println(err.Error())
			return
		}

	}
	log.Println(`init end.`)

}
