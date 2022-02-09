// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package api is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	api.proto
*/
package api

import (
	"context"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathArenaServerOnPMatchStatus = "/arenaServer/onPMatchStatus"
var PathArenaServerOnPShareStatus = "/arenaServer/onPShareStatus"
var PathArenaServerOnPAwardList = "/arenaServer/onPAwardList"
var PathArenaServerOnPAward = "/arenaServer/onPAward"
var PathArenaServerOnPAwardCount = "/arenaServer/onPAwardCount"
var PathArenaServerOnPRankList = "/arenaServer/onPRankList"
var PathArenaServerOnPMatchList = "/arenaServer/onPMatchList"
var PathArenaServerOnPSignUp = "/arenaServer/onPSignUp"
var PathArenaServerOnFMatchStatus = "/arenaServer/onFMatchStatus"
var PathArenaServerOnFShareStatus = "/arenaServer/onFShareStatus"
var PathArenaServerOnFAwardList = "/arenaServer/onFAwardList"
var PathArenaServerOnFAward = "/arenaServer/onFAward"
var PathArenaServerOnFAwardCount = "/arenaServer/onFAwardCount"
var PathArenaServerOnFRankList = "/arenaServer/onFRankList"
var PathArenaServerOnFMatchList = "/arenaServer/onFMatchList"
var PathArenaServerOnWaitMatch = "/arenaServer/onWaitMatch"
var PathArenaServerOnCancelMatch = "/arenaServer/onCancelMatch"
var PathArenaServerOnFSignUp = "/arenaServer/onFSignUp"
var PathArenaServerOnConfirm = "/arenaServer/onConfirm"
var PathArenaServerOnRobotLogin = "/arenaServer/onRobotLogin"
var PathArenaServerOnRobotHeart = "/arenaServer/onRobotHeart"
var PathArenaServerOnGameUserInfo = "/arenaServer/onGameUserInfo"

// ArenaServerBMServer is the server API for ArenaServer service.
type ArenaServerBMServer interface {
	OnPMatchStatus(ctx context.Context, req *MatchStatusReq) (resp *MatchStatusRsp, err error)

	OnPShareStatus(ctx context.Context, req *ShareStatusReq) (resp *ShareStatusRsp, err error)

	OnPAwardList(ctx context.Context, req *AwardListReq) (resp *AwardListRsp, err error)

	OnPAward(ctx context.Context, req *GetAwardReq) (resp *GetAwardRsp, err error)

	OnPAwardCount(ctx context.Context, req *AwardCountReq) (resp *AwardCountRsp, err error)

	OnPRankList(ctx context.Context, req *RankListReq) (resp *RankListRsp, err error)

	OnPMatchList(ctx context.Context, req *PMatchListReq) (resp *PMatchListRsp, err error)

	OnPSignUp(ctx context.Context, req *PSignUpReq) (resp *PSignUpRsp, err error)

	OnFMatchStatus(ctx context.Context, req *MatchStatusReq) (resp *MatchStatusRsp, err error)

	OnFShareStatus(ctx context.Context, req *ShareStatusReq) (resp *ShareStatusRsp, err error)

	OnFAwardList(ctx context.Context, req *AwardListReq) (resp *AwardListRsp, err error)

	OnFAward(ctx context.Context, req *GetAwardReq) (resp *GetAwardRsp, err error)

	OnFAwardCount(ctx context.Context, req *AwardCountReq) (resp *AwardCountRsp, err error)

	OnFRankList(ctx context.Context, req *RankListReq) (resp *RankListRsp, err error)

	OnFMatchList(ctx context.Context, req *FMatchListReq) (resp *FMatchListRsp, err error)

	OnWaitMatch(ctx context.Context, req *FWaitMatchReq) (resp *FWaitMatchRsp, err error)

	OnCancelMatch(ctx context.Context, req *FCancelMatchReq) (resp *FCancelMatchRsp, err error)

	OnFSignUp(ctx context.Context, req *FSignUpReq) (resp *FSignUpRsp, err error)

	OnConfirm(ctx context.Context, req *FConfirmReq) (resp *FConfirmRsp, err error)

	OnRobotLogin(ctx context.Context, req *FRobotLoginReq) (resp *FRobotLoginRsp, err error)

	OnRobotHeart(ctx context.Context, req *FRobotHeartReq) (resp *FRobotLoginRsp, err error)

	OnGameUserInfo(ctx context.Context, req *GameUserInfoReq) (resp *GameUserInfoRsp, err error)
}

var ArenaServerSvc ArenaServerBMServer

func arenaServerOnPMatchStatus(c *bm.Context) {
	p := new(MatchStatusReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnPMatchStatus(c, p)
	c.JSON(resp, err)
}

func arenaServerOnPShareStatus(c *bm.Context) {
	p := new(ShareStatusReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnPShareStatus(c, p)
	c.JSON(resp, err)
}

func arenaServerOnPAwardList(c *bm.Context) {
	p := new(AwardListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnPAwardList(c, p)
	c.JSON(resp, err)
}

func arenaServerOnPAward(c *bm.Context) {
	p := new(GetAwardReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnPAward(c, p)
	c.JSON(resp, err)
}

func arenaServerOnPAwardCount(c *bm.Context) {
	p := new(AwardCountReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnPAwardCount(c, p)
	c.JSON(resp, err)
}

func arenaServerOnPRankList(c *bm.Context) {
	p := new(RankListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnPRankList(c, p)
	c.JSON(resp, err)
}

func arenaServerOnPMatchList(c *bm.Context) {
	p := new(PMatchListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnPMatchList(c, p)
	c.JSON(resp, err)
}

func arenaServerOnPSignUp(c *bm.Context) {
	p := new(PSignUpReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnPSignUp(c, p)
	c.JSON(resp, err)
}

func arenaServerOnFMatchStatus(c *bm.Context) {
	p := new(MatchStatusReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnFMatchStatus(c, p)
	c.JSON(resp, err)
}

func arenaServerOnFShareStatus(c *bm.Context) {
	p := new(ShareStatusReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnFShareStatus(c, p)
	c.JSON(resp, err)
}

func arenaServerOnFAwardList(c *bm.Context) {
	p := new(AwardListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnFAwardList(c, p)
	c.JSON(resp, err)
}

func arenaServerOnFAward(c *bm.Context) {
	p := new(GetAwardReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnFAward(c, p)
	c.JSON(resp, err)
}

func arenaServerOnFAwardCount(c *bm.Context) {
	p := new(AwardCountReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnFAwardCount(c, p)
	c.JSON(resp, err)
}

func arenaServerOnFRankList(c *bm.Context) {
	p := new(RankListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnFRankList(c, p)
	c.JSON(resp, err)
}

func arenaServerOnFMatchList(c *bm.Context) {
	p := new(FMatchListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnFMatchList(c, p)
	c.JSON(resp, err)
}

func arenaServerOnWaitMatch(c *bm.Context) {
	p := new(FWaitMatchReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnWaitMatch(c, p)
	c.JSON(resp, err)
}

func arenaServerOnCancelMatch(c *bm.Context) {
	p := new(FCancelMatchReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnCancelMatch(c, p)
	c.JSON(resp, err)
}

func arenaServerOnFSignUp(c *bm.Context) {
	p := new(FSignUpReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnFSignUp(c, p)
	c.JSON(resp, err)
}

func arenaServerOnConfirm(c *bm.Context) {
	p := new(FConfirmReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnConfirm(c, p)
	c.JSON(resp, err)
}

func arenaServerOnRobotLogin(c *bm.Context) {
	p := new(FRobotLoginReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnRobotLogin(c, p)
	c.JSON(resp, err)
}

func arenaServerOnRobotHeart(c *bm.Context) {
	p := new(FRobotHeartReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnRobotHeart(c, p)
	c.JSON(resp, err)
}

func arenaServerOnGameUserInfo(c *bm.Context) {
	p := new(GameUserInfoReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := ArenaServerSvc.OnGameUserInfo(c, p)
	c.JSON(resp, err)
}

// RegisterArenaServerBMServer Register the blademaster route
func RegisterArenaServerBMServer(e *bm.Engine, server ArenaServerBMServer) {
	ArenaServerSvc = server
	e.POST("/arenaServer/onPMatchStatus", arenaServerOnPMatchStatus)
	e.POST("/arenaServer/onPShareStatus", arenaServerOnPShareStatus)
	e.POST("/arenaServer/onPAwardList", arenaServerOnPAwardList)
	e.POST("/arenaServer/onPAward", arenaServerOnPAward)
	e.POST("/arenaServer/onPAwardCount", arenaServerOnPAwardCount)
	e.POST("/arenaServer/onPRankList", arenaServerOnPRankList)
	e.POST("/arenaServer/onPMatchList", arenaServerOnPMatchList)
	e.POST("/arenaServer/onPSignUp", arenaServerOnPSignUp)
	e.POST("/arenaServer/onFMatchStatus", arenaServerOnFMatchStatus)
	e.POST("/arenaServer/onFShareStatus", arenaServerOnFShareStatus)
	e.POST("/arenaServer/onFAwardList", arenaServerOnFAwardList)
	e.POST("/arenaServer/onFAward", arenaServerOnFAward)
	e.POST("/arenaServer/onFAwardCount", arenaServerOnFAwardCount)
	e.POST("/arenaServer/onFRankList", arenaServerOnFRankList)
	e.POST("/arenaServer/onFMatchList", arenaServerOnFMatchList)
	e.POST("/arenaServer/onWaitMatch", arenaServerOnWaitMatch)
	e.POST("/arenaServer/onCancelMatch", arenaServerOnCancelMatch)
	e.POST("/arenaServer/onFSignUp", arenaServerOnFSignUp)
	e.POST("/arenaServer/onConfirm", arenaServerOnConfirm)
	e.POST("/arenaServer/onRobotLogin", arenaServerOnRobotLogin)
	e.POST("/arenaServer/onRobotHeart", arenaServerOnRobotHeart)
	e.POST("/arenaServer/onGameUserInfo", arenaServerOnGameUserInfo)
}
