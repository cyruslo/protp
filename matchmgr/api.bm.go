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
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathMatchMgrPing = "/matchmgr.service.v1.MatchMgr/Ping"
var PathMatchMgrBonusPool = "/matchmgr/BonusPool"
var PathMatchMgrRankBar = "/matchmgr/RankBar"
var PathMatchMgrRealRank = "/matchmgr/RealRank"
var PathMatchMgrReward = "/matchmgr/Reward"
var PathMatchMgrMatchInfoTip = "/matchmgr/MatchInfoTip"
var PathMatchMgrMatchInfo = "/matchmgr/MatchInfo"
var PathMatchMgrHistoryRank = "/matchmgr/HistoryRank"
var PathMatchMgrLuckyInfo = "/matchmgr/LuckyInfo"
var PathMatchMgrGetAllMatchCfg = "/matchmgr/GetAllMatchCfg"
var PathMatchMgrUpdateBonus = "/matchmgr.service.v1.MatchMgr/UpdateBonus"
var PathMatchMgrMatchResult = "/matchmgr.service.v1.MatchMgr/MatchResult"
var PathMatchMgrGetAwardRecord = "/matchmgr.service.v1.MatchMgr/GetAwardRecord"
var PathMatchMgrGetAward = "/matchmgr.service.v1.MatchMgr/GetAward"
var PathMatchMgrUpdateUserInfo = "/matchmgr.service.v1.MatchMgr/UpdateUserInfo"

// MatchMgrBMServer is the server API for MatchMgr service.
type MatchMgrBMServer interface {
	Ping(ctx context.Context, req *google_protobuf1.Empty) (resp *google_protobuf1.Empty, err error)

	BonusPool(ctx context.Context, req *CSHBonusPool) (resp *SCHBonusPool, err error)

	RankBar(ctx context.Context, req *CSHRankBar) (resp *SCHRankBar, err error)

	RealRank(ctx context.Context, req *CSHRealRank) (resp *SCHRealRank, err error)

	Reward(ctx context.Context, req *CSHReward) (resp *SCHReward, err error)

	MatchInfoTip(ctx context.Context, req *CSHMatchInfoTip) (resp *SCHMatchInfoTip, err error)

	MatchInfo(ctx context.Context, req *CSHMatchInfo) (resp *SCHMatchInfo, err error)

	HistoryRank(ctx context.Context, req *CSHHistoryRank) (resp *SCHHistoryRank, err error)

	LuckyInfo(ctx context.Context, req *CSHLuckyInfo) (resp *SCHLuckyInfo, err error)

	GetAllMatchCfg(ctx context.Context, req *CSMatchCfg) (resp *SCMatchCfg, err error)

	// 以下接口内部使用
	UpdateBonus(ctx context.Context, req *CSUpdateBonus) (resp *SCUpdateBonus, err error)

	MatchResult(ctx context.Context, req *CSMatchResult) (resp *SCMatchResult, err error)

	GetAwardRecord(ctx context.Context, req *CSGetAwardRecord) (resp *SCGetAwardRecord, err error)

	GetAward(ctx context.Context, req *CSGetAward) (resp *SCGetAward, err error)

	UpdateUserInfo(ctx context.Context, req *CSUpdateUserInfo) (resp *SCUpdateUserInfo, err error)
}

var MatchMgrSvc MatchMgrBMServer

func matchMgrPing(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.Ping(c, p)
	c.JSON(resp, err)
}

func matchMgrBonusPool(c *bm.Context) {
	p := new(CSHBonusPool)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.BonusPool(c, p)
	c.JSON(resp, err)
}

func matchMgrRankBar(c *bm.Context) {
	p := new(CSHRankBar)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.RankBar(c, p)
	c.JSON(resp, err)
}

func matchMgrRealRank(c *bm.Context) {
	p := new(CSHRealRank)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.RealRank(c, p)
	c.JSON(resp, err)
}

func matchMgrReward(c *bm.Context) {
	p := new(CSHReward)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.Reward(c, p)
	c.JSON(resp, err)
}

func matchMgrMatchInfoTip(c *bm.Context) {
	p := new(CSHMatchInfoTip)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.MatchInfoTip(c, p)
	c.JSON(resp, err)
}

func matchMgrMatchInfo(c *bm.Context) {
	p := new(CSHMatchInfo)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.MatchInfo(c, p)
	c.JSON(resp, err)
}

func matchMgrHistoryRank(c *bm.Context) {
	p := new(CSHHistoryRank)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.HistoryRank(c, p)
	c.JSON(resp, err)
}

func matchMgrLuckyInfo(c *bm.Context) {
	p := new(CSHLuckyInfo)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.LuckyInfo(c, p)
	c.JSON(resp, err)
}

func matchMgrGetAllMatchCfg(c *bm.Context) {
	p := new(CSMatchCfg)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.GetAllMatchCfg(c, p)
	c.JSON(resp, err)
}

func matchMgrUpdateBonus(c *bm.Context) {
	p := new(CSUpdateBonus)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.UpdateBonus(c, p)
	c.JSON(resp, err)
}

func matchMgrMatchResult(c *bm.Context) {
	p := new(CSMatchResult)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.MatchResult(c, p)
	c.JSON(resp, err)
}

func matchMgrGetAwardRecord(c *bm.Context) {
	p := new(CSGetAwardRecord)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.GetAwardRecord(c, p)
	c.JSON(resp, err)
}

func matchMgrGetAward(c *bm.Context) {
	p := new(CSGetAward)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.GetAward(c, p)
	c.JSON(resp, err)
}

func matchMgrUpdateUserInfo(c *bm.Context) {
	p := new(CSUpdateUserInfo)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MatchMgrSvc.UpdateUserInfo(c, p)
	c.JSON(resp, err)
}

// RegisterMatchMgrBMServer Register the blademaster route
func RegisterMatchMgrBMServer(e *bm.Engine, server MatchMgrBMServer) {
	MatchMgrSvc = server
	e.GET("/matchmgr.service.v1.MatchMgr/Ping", matchMgrPing)
	e.POST("/matchmgr/BonusPool", matchMgrBonusPool)
	e.POST("/matchmgr/RankBar", matchMgrRankBar)
	e.POST("/matchmgr/RealRank", matchMgrRealRank)
	e.POST("/matchmgr/Reward", matchMgrReward)
	e.POST("/matchmgr/MatchInfoTip", matchMgrMatchInfoTip)
	e.POST("/matchmgr/MatchInfo", matchMgrMatchInfo)
	e.POST("/matchmgr/HistoryRank", matchMgrHistoryRank)
	e.POST("/matchmgr/LuckyInfo", matchMgrLuckyInfo)
	e.POST("/matchmgr/GetAllMatchCfg", matchMgrGetAllMatchCfg)
	e.GET("/matchmgr.service.v1.MatchMgr/UpdateBonus", matchMgrUpdateBonus)
	e.GET("/matchmgr.service.v1.MatchMgr/MatchResult", matchMgrMatchResult)
	e.GET("/matchmgr.service.v1.MatchMgr/GetAwardRecord", matchMgrGetAwardRecord)
	e.GET("/matchmgr.service.v1.MatchMgr/GetAward", matchMgrGetAward)
	e.GET("/matchmgr.service.v1.MatchMgr/UpdateUserInfo", matchMgrUpdateUserInfo)
}