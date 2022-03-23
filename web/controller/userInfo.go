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
	IsAdmin	string
}


var users []User

func init() {

	adminAccount := User{LoginName:"root", Password:"root", IsAdmin:"T"}
	stuAccount := User{LoginName:"allen", Password:"123456", IsAdmin:"F"}

	users = append(users, adminAccount)
	users = append(users, stuAccount)

}