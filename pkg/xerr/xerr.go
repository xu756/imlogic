package xerr

// 全局码
const (
	OK = 200 + iota
	SystemErrCode
	DbErrCode
	Param
	DbConnectErr
	RedisConnectErr
	RoleErr
	ExportExcelErr
)

// 用户模块码
const (
	UserExist = 210 + iota
	UserPasswordErr
	UserLoginOut
	UserExpired
	UserDeleteErr
	UserNotExist
	UserRegisterErr
	UserUpdateErr
	UserEmailExist
	UserEmailNotExist
)

// JWT 模块码
const (
	JwtCreateErr = 220 + iota
	JwtParseErr
	JwtAuthEmpty
	JwtAuthErr
	CaptchaErr
	CaptchaExpired
	NotSameUser
)

// Subject 学科
const (
	SubjectExist = 230 + iota

	SubjectCreateErr
	SubjectFindErr
	SubjectUpdateErr
)

// Mark 成绩
const (
	MarkExist = 240 + iota
	MarkCreateErr
	MarkFindErr
)

// Send 发送信息
const (
	SendEmailErr = 250 + iota
	SendSmsErr
	CaptchaNotExist
)

// File 文件
const (
	FileSaveErr = 260 + iota
	UploadFileErr
	FileNotExist
	CreatDirErr
	TeacherConNotDown
	FileTypeErr
	ExcelReadErr
	ExcelFileErr
)
