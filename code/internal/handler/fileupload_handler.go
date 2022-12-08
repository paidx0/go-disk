package handler

import (
	"code/internal/logic"
	"code/internal/svc"
	"code/internal/types"
	"code/models"
	"code/utils"
	"crypto/md5"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"path"
	"strings"
)

func fileuploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// 文件上传验证
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		// 文件是否存在
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		has, err := models.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		// 文件已存在直接返回查询的数据
		if has {
			httpx.OkJson(w, &types.FileUploadResponse{
				Identity: rp.Identity,
				Ext:      rp.Ext,
				Name:     rp.Name,
			})
			return
		}
		// 不存在就直接上传
		var filePath string
		if fileHeader.Size < int64(1024*1024) { // 1MB小文件普通上传
			filePath, err = utils.FileUpload(r)
		} else { // 大文件分块上传
			filePath, err = utils.FileChunkUpload(r)
		}
		if err != nil {
			httpx.Error(w, err)
			return
		}
		// 赋值req 传给下级logic
		req.Ext = path.Ext(fileHeader.Filename)
		req.Name = strings.TrimSuffix(fileHeader.Filename, req.Ext)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = filePath

		l := logic.NewFileuploadLogic(r.Context(), svcCtx)
		resp, err := l.Fileupload(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
