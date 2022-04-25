/**
  @Author : AllenIverson
*/

package controller

import (
	"github.com/fabric-identity/service"
)

type Application struct {
	Setup *service.ServiceSetup
}


type User struct {
	LoginName	string
	Password	string
	IsAdmin		string
}

type StuScore struct {
	StuClss []*service.Score
	StuNum string
	StuName string
}

type ArchivesInfo struct {
	ArchivesID	string
	Operator string
	CurrentUser	string
	CreateTime  string
	UpdateTime  string
	InfoMsg 	string
}

var stuScores []StuScore
var stuArchives map[string][]*service.Archives

var ArchivesInfos []*ArchivesInfo

var users []User

func init() {

	adminAccount := User{LoginName:"root", Password:"root", IsAdmin:"T"}
	stuAccount := User{LoginName:"allen", Password:"123456", IsAdmin:"F"}

	users = append(users, adminAccount)
	users = append(users, stuAccount)

	stuArchives = make(map[string][]*service.Archives)

}