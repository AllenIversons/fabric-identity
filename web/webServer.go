/**
  @Author : AllenIverson
*/

package web

import (
	"github.com/cloudflare/cfssl/log"
	"github.com/fabric-identity/web/controller"
	"net/http"
)


// 启动Web服务并指定路由信息
func WebStart(app controller.Application)  {

	fs:= http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 指定路由信息(匹配请求)
	http.HandleFunc("/", app.LoginView)
	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/loginout", app.LoginOut)
	http.HandleFunc("/registed",app.Register)
	http.HandleFunc("/register", app.RegisterView)

	//档案管理
	http.HandleFunc("/management1", app.QueryManagement)
	http.HandleFunc("/management2", app.OperateManagement)
	http.HandleFunc("/danganadd",app.DanganAdd)
	http.HandleFunc("/addArchives",app.AddArchives)


	http.HandleFunc("/queryArchives",app.DanganCheck)
	http.HandleFunc("/queryArchives2",app.RootDanganCheck)


	http.HandleFunc("/modifyArchives",app.DanganModifyShow)
	http.HandleFunc("/DanganModify",app.DanganModify)



	http.HandleFunc("/index", app.Index)
	http.HandleFunc("/help", app.Help)

	http.HandleFunc("/addEduInfo", app.AddEduShow)	// 显示添加信息页面
	http.HandleFunc("/addEdu", app.AddEdu)	// 提交信息请求
	http.HandleFunc("/addScoreInfo",app.AddScoreShow)
	http.HandleFunc("/addScore",app.AddScore)
	http.HandleFunc("/queryPage", app.QueryPage)	// 转至根据证书编号与姓名查询信息页面
	http.HandleFunc("/query", app.FindCertByNoAndName)	// 根据证书编号与姓名查询信息

	http.HandleFunc("/queryPage2", app.QueryPage2)	// 转至根据身份证号码查询信息页面
	http.HandleFunc("/query2", app.FindByID)	// 根据身份证号码查询信息
	http.HandleFunc("/query3",app.QueryPage3)
	http.HandleFunc("/queryhistory",app.FindByIDHistory)
	http.HandleFunc("/queryScore", app.ShowScore) //根据身份证号查询成绩信息
	http.HandleFunc("/queryScoreResult",app.ByNameFindScore)

	http.HandleFunc("/modifyPage", app.ModifyShow)	// 修改信息页面
	http.HandleFunc("/modify", app.Modify)	//  修改信息

	http.HandleFunc("/upload", app.UploadFile)
	log.Info("启动Web服务, 监听端口号为: 9000")
	//fmt.Println("启动Web服务, 监听端口号为: 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Errorf("Web服务启动失败: %v", err)
		//fmt.Printf("Web服务启动失败: %v", err)
	}

}



