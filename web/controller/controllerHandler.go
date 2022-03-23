/**
  @Prject: goProjects
  @Dev_Software: GoLand
  @File : controllerHandler
  @Time : 2018/10/18 14:31 
  @Author : AllenIverson
*/

package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fabric-identity/service"
	"io"
	"net/http"
)

var cuser User

func (app *Application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	//将请求body解析到dst中
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		//如果在解析的过程中发生错误，开始分类...
		var syntaxError *json.SyntaxError
		var unmarshalTypeError * json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		//使用errors.As()函数检查错误类型是否为*json.SyntaxError。如果是，返回错误
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contain badly-formed JSON (at charcter %d)")
		//某些情况下，因为语法错误Decode()函数可能返回io.ErrUnexpectedEOF错误。
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contain badly-formed JSON")
		//同样捕获*json.UnmarshalTypeError错误，这些错误是因为JSON值和接收对象不匹配。如果错误对应到特定到字段，
		//我们可以指出哪个字段错误方便客户端debug
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		//如果解码到内容是空对象，会返回io.EOF。
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		//如果decode()函数传入一个非空的指针，将返回json.InvalidUnmarshalError。
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	return nil
}



func (app *Application) LoginView(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "login.html", nil)
}

func (app *Application) RegisterView(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "register.html", nil)
}

func (app *Application) Index(w http.ResponseWriter, r *http.Request)  {
	ShowView(w, r, "index.html", nil)
}

func (app *Application) Help(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
	}{
		CurrentUser:cuser,
	}
	ShowView(w, r, "help.html", data)
}
//用户注册
func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")
	password := r.FormValue("password")
	rule := r.FormValue("rule")

	newUser := &User{
		LoginName:		loginName,
		Password:		password,
		IsAdmin: 		rule,
	}
	fmt.Println(newUser)
	fmt.Println("初始",len(users))
	for _,k := range users {
		if k.LoginName == newUser.LoginName {
			fmt.Println("已经存在")
			//w.Write([]byte("a user with this loginName already exists"))
		}
	}
	users = append(users, *newUser)
	fmt.Println("完成",len(users))
	ShowView(w, r, "login.html", nil)

}

// 用户登录
func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")
	password := r.FormValue("password")
	fmt.Println("login:",len(users))
	var flag bool
	for _, user := range users {
		if user.LoginName == loginName && user.Password == password {
			cuser = user
			flag = true
			break
		}
	}

	data := &struct {
		CurrentUser User
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:false,
	}

	if flag {
		// 登录成功
		ShowView(w, r, "index.html", data)
	}else{
		// 登录失败
		data.Flag = true
		data.CurrentUser.LoginName = loginName
		ShowView(w, r, "login.html", data)
	}
}

// 用户登出
func (app *Application) LoginOut(w http.ResponseWriter, r *http.Request)  {
	cuser = User{}
	ShowView(w, r, "login.html", nil)
}

// 显示添加信息页面
func (app *Application) AddEduShow(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "addEdu.html", data)
}

// 添加信息
func (app *Application) AddEdu(w http.ResponseWriter, r *http.Request)  {

	edu := service.Education{
		Name:r.FormValue("name"),
		Gender:r.FormValue("gender"),
		Nation:r.FormValue("nation"),
		EntityID:r.FormValue("entityID"),
		Place:r.FormValue("place"),
		BirthDay:r.FormValue("birthDay"),
		EnrollDate:r.FormValue("enrollDate"),
		GraduationDate:r.FormValue("graduationDate"),
		SchoolName:r.FormValue("schoolName"),
		Major:r.FormValue("major"),
		QuaType:r.FormValue("quaType"),
		Mode:r.FormValue("mode"),
		Graduation:r.FormValue("graduation"),
		CertNo:r.FormValue("certNo"),
		Photo:r.FormValue("photo"),
	}

	app.Setup.SaveEdu(edu)
	/*transactionID, err := app.Setup.SaveEdu(edu)

	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
	}else{
		data.Msg = "信息添加成功:" + transactionID
	}*/

	//ShowView(w, r, "addEdu.html", data)
	r.Form.Set("certNo", edu.CertNo)
	r.Form.Set("name", edu.Name)
	app.FindCertByNoAndName(w, r)
}

func (app *Application) QueryPage(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "query.html", data)
}

// 根据证书编号与姓名查询信息
func (app *Application) FindCertByNoAndName(w http.ResponseWriter, r *http.Request)  {
	certNo := r.FormValue("certNo")
	name := r.FormValue("name")
	result, err := app.Setup.FindEduByCertNoAndName(certNo, name)
	var edu = service.Education{}
	json.Unmarshal(result, &edu)

	fmt.Println("根据证书编号与姓名查询信息成功：")
	fmt.Println(edu)

	data := &struct {
		Edu service.Education
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Edu:edu,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:false,
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "queryResult.html", data)
}

func (app *Application) QueryPage2(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "query2.html", data)
}

// 根据身份证号码查询信息
func (app *Application) FindByID(w http.ResponseWriter, r *http.Request)  {
	entityID := r.FormValue("entityID")
	result, err := app.Setup.FindEduInfoByEntityID(entityID)
	var edu = service.Education{}
	json.Unmarshal(result, &edu)

	data := &struct {
		Edu service.Education
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Edu:edu,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "queryResult.html", data)
}

// 修改/添加新信息
func (app *Application) ModifyShow(w http.ResponseWriter, r *http.Request)  {
	// 根据证书编号与姓名查询信息
	certNo := r.FormValue("certNo")
	name := r.FormValue("name")
	result, err := app.Setup.FindEduByCertNoAndName(certNo, name)

	var edu = service.Education{}
	json.Unmarshal(result, &edu)

	data := &struct {
		Edu service.Education
		CurrentUser User
		Msg string
		Flag bool
	}{
		Edu:edu,
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "modify.html", data)
}

// 修改/添加新信息
func (app *Application) Modify(w http.ResponseWriter, r *http.Request) {
	edu := service.Education{
		Name:r.FormValue("name"),
		Gender:r.FormValue("gender"),
		Nation:r.FormValue("nation"),
		EntityID:r.FormValue("entityID"),
		Place:r.FormValue("place"),
		BirthDay:r.FormValue("birthDay"),
		EnrollDate:r.FormValue("enrollDate"),
		GraduationDate:r.FormValue("graduationDate"),
		SchoolName:r.FormValue("schoolName"),
		Major:r.FormValue("major"),
		QuaType:r.FormValue("quaType"),
		Mode:r.FormValue("mode"),
		Graduation:r.FormValue("graduation"),
		CertNo:r.FormValue("certNo"),
		Photo:r.FormValue("photo"),
	}

	//transactionID, err := app.Setup.ModifyEdu(edu)
	app.Setup.ModifyEdu(edu)

	/*data := &struct {
		Edu service.Education
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
	}else{
		data.Msg = "新信息添加成功:" + transactionID
	}

	ShowView(w, r, "modify.html", data)
	*/

	r.Form.Set("certNo", edu.CertNo)
	r.Form.Set("name", edu.Name)
	app.FindCertByNoAndName(w, r)
}
