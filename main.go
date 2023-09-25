/*
** main.go for main.go
**
** Made by gaspar_d
** Login   gaspar_d <d.gasparina@gmail.com>
**
** Started on  Wed 23 Dec 10:25:07 2015 gaspar_d
** Last update Mon  7 Mar 16:54:09 2016 gaspar_d
 */

package main

import (
	"fmt"
	"github.com/892294101/mbm/src"
)

func main() {
	option := mbm.ParseOptions()
	env := src.BackupEnv{}
	err := env.SetupBackupEnvironment(option)

	if err != nil {
		fmt.Printf("Can not setup program environment (%s)", err)
	}

	if env.Options.Operation == src.OpBackup {
		env.PerformBackup()
	} else if env.Options.Operation == src.OpRestore {
		env.PerformRestore()
	} else if env.Options.Operation == src.OpList {
		env.List(env.Options.Kind)
	} else if env.Options.Operation == src.OpDelete {
		env.PerformDeletion()
	}
}
