// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"io/ioutil"
	"log"
	"github.com/spf13/cobra"
	PostgresqlLister "github.com/zalando/postgres-operator/pkg/generated/clientset/versioned/typed/acid.zalan.do/v1"
	v1 "github.com/zalando/postgres-operator/pkg/apis/acid.zalan.do/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme"
)

// deleteCmd represents kubectl pg delete.
var deleteCmd = &cobra.Command{
	Use:   "delete the resource of type postgresql.",
	Short: "Delete cmd to delete k8s objects by object-name/manifest -file",
	Long: `Delete cmd deletes the objects specific to a manifest file or an object provided object-name.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		deleteByName,_ :=cmd.Flags().GetString("name")
		deleteByFile,_ :=cmd.Flags().GetString("file")
		if(deleteByName!="") {
			delete(deleteByName)
		} else {
			delete(deleteByFile)
		}
	},
}

func init() {
	deleteCmd.Flags().StringP("name","n","","Delete postgresql resource by it's name.")
	deleteCmd.Flags().StringP("file","f","","using file.")
	rootCmd.AddCommand(deleteCmd)
}

func delete(deleteByName string) {
	config:=getConfig()
	ans,err:=PostgresqlLister.NewForConfig(config)
	ymlFile,err := ioutil.ReadFile("abc.yaml")
	if err != nil {
		log.Printf("%#v\n",err)
	}
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj,_,err:= decode([]byte(ymlFile),nil, &v1.Postgresql{})
	if(err!=nil){
		fmt.Println("vineeth",err)
	}
	postgresSql := obj.(*v1.Postgresql)
	fmt.Println(postgresSql)
	deleteStatus:=ans.Postgresqls("default").DeleteCollection(&metav1.DeleteOptions{},metav1.ListOptions{})
	fmt.Println(deleteStatus)
}
