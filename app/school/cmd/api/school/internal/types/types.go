// Code generated by goctl. DO NOT EDIT.
package types

type UserClass struct {
	ID   int64  `json:"id" description:"班级ID"`
	Name string `json:"name" description:"班级名称"`
	Room string `json:"room" description:"班级教室"`
}

type UserClassCreateRequest struct {
	Name string `json:"name" description:"名称"`
	Room string `json:"room" description:"教室"`
}

type UserClassCreateResponse struct {
	Msg string `json:"msg" tag:"返回信息"`
}

type UserClassDetailRequest struct {
	ID int64 `form:"id" description:"ID"`
}

type UserClassDetailResponse struct {
	UserClass UserClass `json:"userClass" description:"班级信息"`
}

type UserRole struct {
	ID   int64  `json:"id" description:"角色ID"`
	Name string `json:"name" description:"角色名称"`
}

type UserRoleCreateRequest struct {
	Name string `json:"name" description:"角色名称"`
}

type UserRoleCreateResponse struct {
	Msg string `json:"msg" tag:"返回信息"`
}

type UserRoleDetailRequest struct {
	ID int64 `form:"id" description:"角色ID"`
}

type UserRoleDetailResponse struct {
	UserRole UserRole `json:"userRole" description:"角色信息"`
}