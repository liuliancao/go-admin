package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_NEED_TOKEN:          "需要参数token",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_AUTH_GETUID_BY_TOKEN:     "Token检索不到UID",

	ERROR_CHECK_EXIST_PRODUCT_FAIL: "检查产品存在时候报错",
	ERROR_NOT_EXIST_PRODUCT:        "产品不存在",
	ERROR_GET_PRODUCT_FAIL:         "获取产品失败",
	ERROR_GET_PRODUCTS_FAIL:        "获取产品列表失败",
	ERROR_ADD_PRODUCT_FAIL:         "添加产品失败",
	ERROR_DELETE_PRODUCT_FAIL:      "删除产品失败",
	ERROR_ADD_PRODUCTUSER_FAIL:     "增加产品负责人失败",
	ERROR_GET_PRODUCTUSERS_FAIL:    "获取产品负责人失败",

	ERROR_CHECK_EXIST_APP_FAIL: "检查应用存在时候报错",
	ERROR_NOT_EXIST_APP:        "应用不存在",
	ERROR_GET_APP_FAIL:         "获取应用失败",
	ERROR_GET_APPS_FAIL:        "获取应用列表失败",
	ERROR_ADD_APP_FAIL:         "添加应用失败",
	ERROR_DELETE_APP_FAIL:      "删除应用失败",
	ERROR_ADD_APP_USER_FAIL:    "添加服务负责人失败",
	ERROR_GET_APPUSERS_FAIL:    "获取服务负责人失败",

	ERROR_CHECK_EXIST_APPENV_FAIL: "检查环境存在时候报错",
	ERROR_NOT_EXIST_APPENV:        "环境不存在",
	ERROR_GET_APPENV_FAIL:         "获取环境失败",
	ERROR_GET_APPENVS_FAIL:        "获取环境列表失败",
	ERROR_ADD_APPENV_FAIL:         "添加环境失败",
	ERROR_DELETE_APPENV_FAIL:      "删除环境失败",

	ERROR_ADD_APP_CLUSTER_FAIL: "集群添加到应用失败",

	ERROR_CHECK_EXIST_DEPARTMENT_FAIL: "检查部门存在时候报错",
	ERROR_NOT_EXIST_DEPARTMENT:        "部门不存在",
	ERROR_GET_DEPARTMENT_FAIL:         "获取部门失败",
	ERROR_GET_DEPARTMENTS_FAIL:        "获取部门列表失败",
	ERROR_ADD_DEPARTMENT_FAIL:         "添加部门失败",
	ERROR_DELETE_DEPARTMENT_FAIL:      "删除部门失败",

	ERROR_CHECK_EXIST_ROLE_FAIL: "检查角色存在时候报错",
	ERROR_NOT_EXIST_ROLE:        "角色不存在",
	ERROR_GET_ROLE_FAIL:         "获取角色失败",
	ERROR_GET_ROLES_FAIL:        "获取角色列表失败",
	ERROR_ADD_ROLE_FAIL:         "添加角色失败",
	ERROR_DELETE_ROLE_FAIL:      "删除角色失败",

	ERROR_CHECK_EXIST_HOST_FAIL:             "检查主机存在时候报错",
	ERROR_CHECK_EXIST_HOST_BY_HOSTNAME_FAIL: "通过主机名检查主机存在适合报错",
	ERROR_NOT_EXIST_HOST:                    "主机不存在",
	ERROR_GET_HOST_FAIL:                     "获取主机失败",
	ERROR_GET_HOSTS_FAIL:                    "获取主机列表失败",
	ERROR_ADD_HOST_FAIL:                     "添加主机失败",
	ERROR_DELETE_HOST_FAIL:                  "删除主机失败",
	ERROR_GET_TOTAL_FAIL:                    "获取主机总数失败",

	ERROR_CHECK_EXIST_CLUSTER_FAIL: "检查集群存在时候报错",
	ERROR_NOT_EXIST_CLUSTER:        "集群不存在",
	ERROR_GET_CLUSTER_FAIL:         "获取集群失败",
	ERROR_GET_CLUSTERS_FAIL:        "获取集群列表失败",
	ERROR_ADD_CLUSTER_FAIL:         "添加集群失败",
	ERROR_DELETE_CLUSTER_FAIL:      "删除集群失败",
	ERROR_ADD_CLUSTERUSER_FAIL:     "集群增加负责人失败",
	ERROR_GET_CLUSTERUSER_FAIL:     "集群获取负责人失败",

	ERROR_CHECK_EXIST_TAG_FAIL: "检查标签存在时候报错",
	ERROR_NOT_EXIST_TAG:        "标签不存在",
	ERROR_GET_TAG_FAIL:         "获取标签失败",
	ERROR_GET_TAGS_FAIL:        "获取标签列表失败",
	ERROR_ADD_TAG_FAIL:         "添加标签失败",
	ERROR_DELETE_TAG_FAIL:      "删除标签失败",

	ERROR_CHECK_EXIST_USER_FAIL: "检查用户存在时候报错",
	ERROR_NOT_EXIST_USER:        "用户不存在",
	ERROR_GET_USER_FAIL:         "获取用户失败",
	ERROR_GET_USERS_FAIL:        "获取用户列表失败",
	ERROR_ADD_USER_FAIL:         "添加用户失败、指定部门不存在、用户重复",
	ERROR_DELETE_USER_FAIL:      "删除用户失败",

	ERROR_CHECK_EXIST_RESOURCETYPE_FAIL: "检查资源类型存在时候报错",
	ERROR_NOT_EXIST_RESOURCETYPE:        "资源类型不存在",
	ERROR_GET_RESOURCETYPE_FAIL:         "获取资源类型失败",
	ERROR_GET_RESOURCETYPES_FAIL:        "获取资源类型列表失败",
	ERROR_ADD_RESOURCETYPE_FAIL:         "添加资源类型失败、指定部门不存在、资源类型重复",
	ERROR_DELETE_RESOURCETYPE_FAIL:      "删除资源类型失败",

	ERROR_CHECK_EXIST_RESOURCE_FAIL: "检查资源存在时候报错",
	ERROR_NOT_EXIST_RESOURCE:        "资源不存在",
	ERROR_GET_RESOURCE_FAIL:         "获取资源失败",
	ERROR_GET_RESOURCES_FAIL:        "获取资源列表失败",
	ERROR_ADD_RESOURCE_FAIL:         "添加资源失败、指定部门不存在、资源重复",
	ERROR_DELETE_RESOURCE_FAIL:      "删除资源失败",

	ERROR_CHECK_EXIST_MACRO_FAIL: "检查宏存在时候报错",
	ERROR_NOT_EXIST_MACRO:        "宏不存在",
	ERROR_GET_MACRO_FAIL:         "获取宏失败",
	ERROR_GET_MACROS_FAIL:        "获取宏列表失败",
	ERROR_ADD_MACRO_FAIL:         "添加宏失败、指定部门不存在、宏重复",
	ERROR_DELETE_MACRO_FAIL:      "删除宏失败",

	ERROR_CHECK_EXIST_HOSTDEPLOY_FAIL: "检查主机部署信息存在时候报错",
	ERROR_NOT_EXIST_HOSTDEPLOY:        "主机部署信息不存在",
	ERROR_GET_HOSTDEPLOY_FAIL:         "获取主机部署信息失败",
	ERROR_GET_HOSTDEPLOYS_FAIL:        "获取主机部署信息列表失败",
	ERROR_ADD_HOSTDEPLOY_FAIL:         "添加主机部署信息失败、主机部署信息重复",
	ERROR_DELETE_HOSTDEPLOY_FAIL:      "删除主机部署信息失败",

	ERROR_CHECK_EXIST_GUARDER_FAIL: "检查网关信息存在时候报错",
	ERROR_NOT_EXIST_GUARDER:        "网关信息不存在",
	ERROR_GET_GUARDER_FAIL:         "获取网关信息失败",
	ERROR_GET_GUARDERS_FAIL:        "获取网关信息列表失败",
	ERROR_ADD_GUARDER_FAIL:         "添加网关信息失败、网关信息重复",
	ERROR_DELETE_GUARDER_FAIL:      "删除网关信息失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
