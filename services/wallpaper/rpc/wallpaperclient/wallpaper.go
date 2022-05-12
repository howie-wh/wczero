// Code generated by goctl. DO NOT EDIT!
// Source: wallpaper.proto

package wallpaperclient

import (
	"context"

	"wczero/services/wallpaper/rpc/wallpaper"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CategoryInfo     = wallpaper.CategoryInfo
	CategoryRequest  = wallpaper.CategoryRequest
	CategoryResponse = wallpaper.CategoryResponse
	DetailRequest    = wallpaper.DetailRequest
	DetailResponse   = wallpaper.DetailResponse
	ImportRequest    = wallpaper.ImportRequest
	ImportResponse   = wallpaper.ImportResponse
	ListRequest      = wallpaper.ListRequest
	ListResponse     = wallpaper.ListResponse
	RemoveRequest    = wallpaper.RemoveRequest
	RemoveResponse   = wallpaper.RemoveResponse
	TypeInfo         = wallpaper.TypeInfo
	WallPaperInfo    = wallpaper.WallPaperInfo

	Wallpaper interface {
		Import(ctx context.Context, in *ImportRequest, opts ...grpc.CallOption) (*ImportResponse, error)
		Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
		Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
		List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
		Category(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	}

	defaultWallpaper struct {
		cli zrpc.Client
	}
)

func NewWallpaper(cli zrpc.Client) Wallpaper {
	return &defaultWallpaper{
		cli: cli,
	}
}

func (m *defaultWallpaper) Import(ctx context.Context, in *ImportRequest, opts ...grpc.CallOption) (*ImportResponse, error) {
	client := wallpaper.NewWallpaperClient(m.cli.Conn())
	return client.Import(ctx, in, opts...)
}

func (m *defaultWallpaper) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	client := wallpaper.NewWallpaperClient(m.cli.Conn())
	return client.Remove(ctx, in, opts...)
}

func (m *defaultWallpaper) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	client := wallpaper.NewWallpaperClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}

func (m *defaultWallpaper) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	client := wallpaper.NewWallpaperClient(m.cli.Conn())
	return client.List(ctx, in, opts...)
}

func (m *defaultWallpaper) Category(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	client := wallpaper.NewWallpaperClient(m.cli.Conn())
	return client.Category(ctx, in, opts...)
}
