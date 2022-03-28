/**
  author: AllenIverson
 */

package main

import (
	"log"
	"os"
	"fmt"
	"github.com/fabric-identity/sdkInit"
	"github.com/fabric-identity/service"
	"encoding/json"
	"github.com/fabric-identity/web/controller"
	"github.com/fabric-identity/web"
)

const (
	configFile = "config.yaml"
	initialized = false
	EduCC = "educc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID: "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/kongyixueyuan.com/education/fixtures/artifacts/channel.tx",

		OrgAdmin:"Admin",
		OrgName:"Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID: EduCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath: "github.com/kongyixueyuan.com/education/chaincode/",
		UserName:"User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//===========================================//

	serviceSetup := service.ServiceSetup{
		ChaincodeID:EduCC,
		Client:channelClient,
	}

	edu := service.Education{
		Name: "李小冉",
		Gender: "女",
		Nation: "汉",
		EntityID: "120225199603180672",
		Place: "重庆",
		BirthDay: "1996年03月18日",
		EnrollDate: "2021年9月",
		GraduationDate: "2025年7月",
		SchoolName: "重庆邮电大学",
		Major: "计算机科学",
		QuaType: "本科",
		Mode: "普通全日制",
		Graduation: "在读",
		CertNo: "120023456",
		Photo: "/static/photo/11.png",
	}

	edu2 := service.Education{
		Name: "范冰冰",
		Gender: "女",
		Nation: "汉",
		EntityID: "120225198303040874",
		Place: "重庆",
		BirthDay: "1983年03月04日",
		EnrollDate: "2020年9月",
		GraduationDate: "2014年7月",
		SchoolName: "重庆邮电大学",
		Major: "表演专业",
		QuaType: "普通",
		Mode: "普通全日制",
		Graduation: "在读",
		CertNo: "120025678",
		Photo: "/static/photo/22.png",
	}

	msg, err := serviceSetup.SaveEdu(edu)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		log.Printf("信息发布成功, 交易编号为: %v",msg)
	}

	msg, err = serviceSetup.SaveEdu(edu2)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		log.Printf("信息发布成功, 交易编号为: %v",msg)
	}

	// 根据证书编号与名称查询信息
	result, err := serviceSetup.FindEduByCertNoAndName("120025678","范冰冰")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		log.Printf("根据证书编号与姓名查询信息成功：%v",edu)
	}

	// 根据身份证号码查询信息
	result, err = serviceSetup.FindEduInfoByEntityID("120225199603180672")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		log.Printf("根据身份证号码查询信息成功：%v",edu)
	}


	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)

}
