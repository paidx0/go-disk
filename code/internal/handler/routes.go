// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"code/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: userloginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/detail",
				Handler: userdetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/mail/code/send",
				Handler: mailsendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: userregisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Middleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/file/upload",
					Handler: fileuploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/list",
					Handler: userfilelistHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/name/update",
					Handler: userfilenameupdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/delete",
					Handler: userfiledeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/owen/update",
					Handler: userfileowenupdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/share",
					Handler: sharefileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/share/list",
					Handler: sharelistHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/share/detail",
					Handler: sharefiledownloadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/share/pull",
					Handler: sharefilepullHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/download",
					Handler: filedownloadHandler(serverCtx),
				},
			}...,
		),
	)
}