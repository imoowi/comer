/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package global

import (
	"fmt"

	"go.uber.org/dig"
)

var DigContainer *dig.Container
var MigrateContainer *dig.Container

var DigContainerProviders []any
var DigContainerMigrateProviders []any

func init() {
	DigContainer = dig.New()
	MigrateContainer = dig.New()
	DigContainerProviders = make([]any, 0)
	DigContainerMigrateProviders = make([]any, 0)
}
func RegisterContainerProviders(provider any) {
	DigContainerProviders = append(DigContainerProviders, provider)
}
func RegisterMigrateContainerProviders(provider any) {
	DigContainerMigrateProviders = append(DigContainerMigrateProviders, provider)
}
func initContainerProviders() {
	for _, provider := range DigContainerProviders {
		err := DigContainer.Invoke(provider)
		if err != nil {
			fmt.Printf("err: %s\n", err)
		}
	}
}
func initMigrateContainerProviders() {
	for _, provider := range DigContainerMigrateProviders {
		err := MigrateContainer.Invoke(provider)
		if err != nil {
			fmt.Printf("err: %s\n", err)
		}
	}
}
func Bootstrap() {
	defer initContainerProviders()
	initLog()
	initMysql()
	initRedis()
	initCasbin()
}

func BootMigrate() {
	defer initMigrateContainerProviders()
	initMysql()
	initCasbin()
}
