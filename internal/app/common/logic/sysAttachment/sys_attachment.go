package sysAttachment

import (
	"IdentifyService/internal/app/common/dao"
	"IdentifyService/internal/app/common/model"
	"IdentifyService/internal/app/common/service"
	"IdentifyService/internal/app/system/consts"
	"IdentifyService/library/liberr"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

var fileKind = map[string]*gset.StrSet{
	//图片
	"image": gset.NewStrSetFrom([]string{
		"jpeg",
		"jpg",
		"png",
		"gif",
		"webp",
		"cr2",
		"tif",
		"bmp",
		"heif",
		"jxr",
		"psd",
		"ico",
		"dwg",
	}),
	//文档
	"doc": gset.NewStrSetFrom([]string{
		"doc",
		"docx",
		"dot",
		"xls",
		"xlt",
		"xlsx",
		"xltx",
		"ppt",
		"pptx",
		"pdf",
		"txt",
		"csv",
		"html",
		"xml",
		"pptm",
		"pot",
		"wpd",
		"md",
		"json",
		"yaml",
		"markdown",
		"asciidoc",
		"xsl",
		"wps",
		"sxi",
		"sti",
		"odp",
	}),
	//视频
	"video": gset.NewStrSetFrom([]string{
		"mp4",
		"m4v",
		"mkv",
		"webm",
		"mov",
		"avi",
		"wmv",
		"mpg",
		"flv",
		"3gp",
	}),
	//音频
	"audio": gset.NewStrSetFrom([]string{
		"mid",
		"mp3",
		"m4a",
		"ogg",
		"flac",
		"wav",
		"amr",
		"aac",
		"aiff",
	}),
	//压缩包
	"zip": gset.NewStrSetFrom([]string{
		"zip",
		"rar",
		"tar",
		"gz",
		"7z",
		"tar.gz",
	}),
	//其它
	"other": gset.NewStrSetFrom([]string{
		"dwf",
		"ics",
		"vcard",
		"apk",
		"ipa",
	}),
}

func init() {
	service.RegisterSysAttachment(New())
}

func New() service.ISysAttachment {
	return &sSysAttachment{}
}

type sSysAttachment struct{}
type AddHandler = func(ctx context.Context) (err error)

func (s *sSysAttachment) List(ctx context.Context, req *model.SysAttachmentSearchReq) (listRes *model.SysAttachmentSearchRes, err error) {
	listRes = new(model.SysAttachmentSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysAttachment.Ctx(ctx).WithAll()
		if req.AppId != "" {
			m = m.Where(dao.SysAttachment.Columns().AppId+" = ?", req.AppId)
		}
		if req.Drive != "" {
			m = m.Where(dao.SysAttachment.Columns().Drive+" = ?", req.Drive)
		}
		if req.Name != "" {
			m = m.Where(dao.SysAttachment.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.Kind != "" {
			m = m.Where(dao.SysAttachment.Columns().Kind+" = ?", req.Kind)
		}
		if req.MimeType != "" {
			m = m.Where(dao.SysAttachment.Columns().MimeType+" like ?", "%"+req.MimeType+"%")
		}
		if req.Status != "" {
			m = m.Where(dao.SysAttachment.Columns().Status+" = ?", gconv.Bool(req.Status))
		}
		if req.CreatedAt != nil && len(req.CreatedAt) > 0 {
			if req.CreatedAt[0] != "" {
				m = m.Where(dao.SysAttachment.Columns().UpdatedAt+" >= ?", gconv.Time(req.CreatedAt[0]))
			}
			if len(req.CreatedAt) > 1 && req.CreatedAt[1] != "" {
				m = m.Where(dao.SysAttachment.Columns().UpdatedAt+" < ?", gconv.Time(req.CreatedAt[1]))
			}
		}
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "updated_at desc,id desc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.SysAttachmentListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SysAttachmentListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SysAttachmentListRes{
				Id:        v.Id,
				AppId:     v.AppId,
				Drive:     v.Drive,
				Name:      v.Name,
				Kind:      v.Kind,
				Path:      v.Path,
				Size:      v.Size,
				Ext:       v.Ext,
				Status:    v.Status,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			}
		}
	})
	return
}

func (s *sSysAttachment) GetById(ctx context.Context, id int64) (res *model.SysAttachmentInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysAttachment.Ctx(ctx).WithAll().Where(dao.SysAttachment.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSysAttachment) GetByMd5(ctx context.Context, md5 string) (res *model.SysAttachmentInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysAttachment.Ctx(ctx).WithAll().Where(dao.SysAttachment.Columns().Md5, md5).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
		if res != nil {
			_, _ = dao.SysAttachment.Ctx(ctx).Unscoped().WherePri(res.Id).Update(map[string]interface{}{
				dao.SysAttachment.Columns().UpdatedAt: gtime.Now(),
			})
		}
	})
	return
}

func (s *sSysAttachment) AddUpload(ctx context.Context, req *model.UploadResponse, attr *model.SysAttachmentAddAttribute) (err error) {
	ext := gstr.SubStrRune(req.Name, gstr.PosRRune(req.Name, ".")+1, gstr.LenRune(req.Name)-1)
	err = s.Add(ctx, &model.SysAttachmentAddReq{
		AppId:     attr.AppId,
		Drive:     attr.Driver,
		Name:      req.Name,
		Kind:      s.getFileKind(ext),
		MimeType:  req.Type,
		Path:      req.Path,
		Size:      req.Size,
		Ext:       ext,
		Md5:       attr.Md5,
		Status:    true,
		CreatedBy: attr.UserId,
	})
	return
}

func (s *sSysAttachment) Add(ctx context.Context, req *model.SysAttachmentAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAttachment.Ctx(ctx).Insert(map[string]interface{}{
			dao.SysAttachment.Columns().AppId:     req.AppId,
			dao.SysAttachment.Columns().Drive:     req.Drive,
			dao.SysAttachment.Columns().Name:      req.Name,
			dao.SysAttachment.Columns().Kind:      req.Kind,
			dao.SysAttachment.Columns().MimeType:  req.MimeType,
			dao.SysAttachment.Columns().Path:      req.Path,
			dao.SysAttachment.Columns().Size:      req.Size,
			dao.SysAttachment.Columns().Ext:       req.Ext,
			dao.SysAttachment.Columns().Md5:       req.Md5,
			dao.SysAttachment.Columns().Status:    req.Status,
			dao.SysAttachment.Columns().CreatedBy: req.CreatedBy,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sSysAttachment) Edit(ctx context.Context, req *model.SysAttachmentEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAttachment.Ctx(ctx).WherePri(req.Id).Update(map[string]interface{}{
			dao.SysAttachment.Columns().AppId:    req.AppId,
			dao.SysAttachment.Columns().Drive:    req.Drive,
			dao.SysAttachment.Columns().Name:     req.Name,
			dao.SysAttachment.Columns().Kind:     req.Kind,
			dao.SysAttachment.Columns().MimeType: req.MimeType,
			dao.SysAttachment.Columns().Path:     req.Path,
			dao.SysAttachment.Columns().Size:     req.Size,
			dao.SysAttachment.Columns().Ext:      req.Ext,
			dao.SysAttachment.Columns().Md5:      req.Md5,
			dao.SysAttachment.Columns().Status:   req.Status,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSysAttachment) Delete(ctx context.Context, ids []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAttachment.Ctx(ctx).Delete(dao.SysAttachment.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// 附件管理状态修改（状态）
func (s *sSysAttachment) ChangeStatus(ctx context.Context, id int64, status bool) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysAttachment.Ctx(ctx).WherePri(id).
			Update(map[string]interface{}{
				dao.SysAttachment.Columns().Status: status,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSysAttachment) getFileKind(ext string) string {
	for k, v := range fileKind {
		if v.ContainsI(ext) {
			return k
		}
	}
	return "other"
}
