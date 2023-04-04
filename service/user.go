package service

import (
	"context"
	"gin_mall/conf"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/pkg/utils"
	"gin_mall/serializer"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
	"time"
)

type UserService struct {
	// 现阶段，前端去验证
	NickName string `json:"nick_name" form:"nick_name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"`
}

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"` // 1：绑定邮箱 2：解绑邮箱 3.改密码
}

type ValidEmailService struct {
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

// Register 注册
func (service *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	var err error
	code := e.Success
	// 验证密钥长度
	if service.Key == "" || len(service.Key) < 16 {
		code = e.ErrorUserKey
		return serializer.ErrorByCode(code, err)
	}
	// 对密钥进行对称加密
	utils.Encrypt.SetKey(service.Key)

	// 判断用户是否已存在
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.Username)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.ErrorByCode(code, err)
	}

	// 给user赋值
	user = model.User{
		Username: service.Username,
		Avatar:   "avatar.jpg",
		NickName: service.NickName,
		Status:   model.Active,
		Money:    utils.Encrypt.AesEncoding("10000"), // 初始化默认金额为10000元
	}

	// 对密码进行加密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.ErrorByCode(code, err)
	}

	// 创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error
	}
	return serializer.Success(nil)
}

// Login 登录
func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	userDao := dao.NewUserDao(ctx)

	// 检查用户是否存在
	user, exist, err := userDao.ExistOrNotByUserName(service.Username)
	if !exist || err != nil {
		code = e.ErrorUserNotFound
		return serializer.ErrorByCode(code, err)
	}

	// 查看用户密码是否正确
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.ErrorByCode(code, err)
	}

	// 签发token （认证）
	token, err := utils.GenerateToken(user.ID, service.Username, "0")
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.ErrorByCode(code, err)
	}

	return serializer.Success(serializer.TokenData{User: serializer.BuildUser(user), Token: token})
}

// Update 修改用户信息
func (service *UserService) Update(ctx context.Context, uid uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)

	// 通过ID查找用户
	user, err = userDao.GetUserByID(uid)

	// 修改昵称
	if user.NickName != "" {
		user.NickName = service.NickName
	}

	err = userDao.UpdateUserByID(uid, user)

	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	return serializer.Success(serializer.BuildUser(user))
}

// Post 上传头像
func (service *UserService) Post(ctx context.Context, uid uint, file multipart.File, size int64) serializer.Response {
	code := e.Success
	var user *model.User
	var err error
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByID(uid)
	if err != nil {
		code = e.ErrorUserNotFound
		return serializer.ErrorByCode(code, err)
	}
	// 保存图片到本地
	path, err := UploadAvatarToLocal(file, uid, user.Username)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.ErrorByCode(code, err)
	}
	user.Avatar = path
	err = userDao.UpdateUserByID(uid, user)

	return serializer.Success(serializer.BuildUser(user))
}

// Send 发送邮件
func (service *SendEmailService) Send(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var address string
	var notice *model.Notice // 绑定邮箱、修改密码时用到的模板通知
	token, err := utils.GenerateEmailToken(uid, service.Email, service.Password, service.OperationType)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.ErrorByCode(code, err)
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeByID(service.OperationType)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	address = conf.ValidEmail + token // 发送邮件的地址
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "wxshhh")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(nil)
}

// Valid 验证邮箱
func (service *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userID uint
	var email string
	var password string
	var operationType uint
	var err error
	code := e.Success
	// 验证token
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := utils.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userID = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.ErrorByCode(code, err)
	}

	// 获取该用户的信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByID(userID)
	if err != nil {
		code = e.ErrorUserNotFound
		return serializer.ErrorByCode(code, err)
	}

	switch operationType {
	case 1: // 修改邮箱
		user.Email = email
	case 2: // 解绑邮箱
		user.Email = ""
	case 3: // 修改密码
		err = user.SetPassword(password)
		if err != nil {
			code = e.Error
			return serializer.ErrorByCode(code, err)
		}
	}
	err = userDao.UpdateUserByID(userID, user)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(serializer.BuildUser(user))
}

// Show 展示用户金额
func (service *ShowMoneyService) Show(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByID(uid)
	if err != nil {
		code = e.ErrorUserNotFound
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(serializer.BuildMoney(user, service.Key))
}
