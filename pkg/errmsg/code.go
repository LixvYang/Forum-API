package errmsg

var (
	// Common errors
	OK                  = &Errmsg{Code: 0, Message: "OK"}
	InternalServerError = &Errmsg{Code: 10001, Message: "网络错误"}
	ErrBind             = &Errmsg{Code: 10002, Message: "参数错误"}

	ErrValidation = &Errmsg{Code: 20001, Message: "参数错误"}
	ErrDatabase   = &Errmsg{Code: 20002, Message: "数据库错误"}
	ErrToken      = &Errmsg{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errmsg{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errmsg{Code: 20102, Message: "没有找到该用户"}
	ErrTokenInvalid      = &Errmsg{Code: 20103, Message: "无效token"}
	ErrPasswordIncorrect = &Errmsg{Code: 20104, Message: "密码不正确"}

	// tags errors
	ErrTagNotFound = &Errmsg{Code: 20105, Message: "标签不存在"}

	// categories errors
	ErrCategoryNotFound = &Errmsg{Code: 20106, Message: "分类不存在"}

	// article errors
	ErrArticleNotFound = &Errmsg{Code: 20107, Message: "文章不存在"}

	// permission errors
	ErrmsgtPermission = &Errmsg{Code: 20108, Message: "没有权限"}

	// casbin errors
	ErrCreateCasbin = &Errmsg{Code: 20109, Message: "创建策略失败"}

	// menu errors
	ErrMenuGet = &Errmsg{Code: 20120, Message: "菜单不存在"}

	// role errors
	ErrRoleGet = &Errmsg{Code: 20121, Message: "角色不存在"}
)
