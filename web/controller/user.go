package controller

import (
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func getUser(c echo.Context) error {
	var uid uint
	err := echo.QueryParamsBinder(c).
		MustUint("uid", &uid).
		BindError()
	if err != nil {
		return utils.ReturnBadRequestForRequiredUint(c, "uid")
	}

	user, usrErr := model.QueryUserById(uid);
	if  errors.Is(usrErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "user not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: user,
	})
}

func getAllUser(c echo.Context) error {
	users, _ := model.QueryAllUser()
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: users,
	})
}

func register(c echo.Context) error{
	var zjuId, rawPassword, name, nickname, mobile, email string
	var passportId uint

	err := echo.FormFieldBinder(c).
		MustString("zjuid",&zjuId).
		MustString("name",&name).
		MustString("password",&rawPassword).
		MustString("nickname",&nickname).
		MustString("mobile",&mobile).
		MustString("email",&email).
		MustUint("passportid",&passportId).
		BindError()
	if err != nil{
		return errors.New("some parameters are missing or wrong")
	}
	var password = []byte(rawPassword)
	err = model.CreateUser(&model.User{
		Name: name,
		ZJUid: zjuId,
		Password: password,
		Nickname: nickname,
		Mobile: mobile,
		Email: email,
		PassportId: passportId,
	})
	if err != nil{
		return utils.ReturnInternalError(c,"user creation")
	}
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "register success",
	})
}

func login(c echo.Context) error{
	var zjuId, rawPassword string

	err := echo.FormFieldBinder(c).
		MustString("zjuid",&zjuId).
		MustString("password",&rawPassword).
		BindError()
	if err != nil{
		return utils.ReturnBadRequestForRequiredString(c,"zjuid", "password");
	}
	var dbuser *model.User
	dbuser,err = model.QueryUserByZJUid(zjuId)
	if  errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "user not found",
		})
	}
	if bcrypt.CompareHashAndPassword(dbuser.Password,[]byte(rawPassword)) == bcrypt.ErrMismatchedHashAndPassword{
		return errors.New("wrong password")
	}
	jwtString, expireTime, err := utils.GenerateJwt(dbuser.ID)
	if err != nil{
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "error",
			Data: "can not generate jwt token while login",
		})
	}
	cookie := new(http.Cookie)
	cookie.Value = jwtString
	cookie.Name = "qsc_rop_jwt"
	cookie.Expires = *expireTime
	cookie.Path = "/"
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "cookie has been set",
	})
}





