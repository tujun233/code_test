package main

import (
	"code.byted.org/lang/gg/gptr"
	"encoding/json"
	"fmt"
	"time"
)

type CTAButtonTypeEnum int64

const (
	CTAButtonTypeEnum_WATCH_NOW       CTAButtonTypeEnum = 1
	CTAButtonTypeEnum_GET_TICKETS     CTAButtonTypeEnum = 2
	CTAButtonTypeEnum_BUY_NOW         CTAButtonTypeEnum = 3
	CTAButtonTypeEnum_RENT_NOW        CTAButtonTypeEnum = 4
	CTAButtonTypeEnum_ADD_TO_CALENDAR CTAButtonTypeEnum = 5
	CTAButtonTypeEnum_REPLAY          CTAButtonTypeEnum = 6
	CTAButtonTypeEnum_VISIT_WEBSITE   CTAButtonTypeEnum = 7
	CTAButtonTypeEnum_LEARN_MORE      CTAButtonTypeEnum = 8
	CTAButtonTypeEnum_OTHER           CTAButtonTypeEnum = 99
)

type SportsLeagueScheduleItemStatus int64

const (
	SportsLeagueScheduleItemStatus_NotStart   SportsLeagueScheduleItemStatus = 1
	SportsLeagueScheduleItemStatus_InProgress SportsLeagueScheduleItemStatus = 2
	SportsLeagueScheduleItemStatus_Finished   SportsLeagueScheduleItemStatus = 3
)

type SportsMatchBaseInfo struct {
	MatchID             *int64  `json:"MatchID"`
	MediaID             *int64  `json:"MediaID"`
	EventTime           *int64  `json:"EventTime"`
	GroupName           *string `json:"GroupName"`
	RoundName           *string `json:"RoundName"`
	PhaseName           *string `json:"PhaseName"`
	MatchStatus         *int    `json:"MatchStatus"`
	HostName            *string `json:"HostName"`
	GuestName           *string `json:"GuestName"`
	HostID              *int64  `json:"HostID"`
	GuestID             *int64  `json:"GuestID"`
	HostLogoMaterialID  *int64  `json:"HostLogoMaterialID"`
	GuestLogoMaterialID *int64  `json:"GuestLogoMaterialID"`
	HostScore           *int    `json:"HostScore"`
	GuestScore          *int    `json:"GuestScore"`
	Extra               *string `json:"Extra"`
	HostLogoURI         *string `json:"HostLogoURI"`
	GuestLogoURI        *string `json:"GuestLogoURI"`
}

type ImageURLInfoStruct struct {
	MaterialID *int64 `json:"MaterialID,omitempty" bson:"MaterialID,omitempty"`
}

type RelatedActivityCampaignInfoStruct struct {
	ShowBanner                *bool               `json:"ShowBanner,omitempty" bson:"ShowBanner,omitempty"`
	ActivityCampaignStartTime *int64              `json:"ActivityCampaignStartTime,omitempty" bson:"ActivityCampaignStartTime,omitempty"`
	ActivityCampaignEndTime   *int64              `json:"ActivityCampaignEndTime,omitempty" bson:"ActivityCampaignEndTime,omitempty"`
	BannerURLInfo             *ImageURLInfoStruct `json:"BannerURLInfo,omitempty" bson:"BannerURLInfo,omitempty"`
	ActivityCampaignURL       *string             `json:"ActivityCampaignURL,omitempty" bson:"ActivityCampaignURL,omitempty"`
	AvailableRegions          []string            `json:"AvailableRegions,omitempty" bson:"AvailableRegions,omitempty"`
}

type CTARedirectInfoStruct struct {
	RedirectURL         *string             `json:"RedirectURL,omitempty" bson:"RedirectURL,omitempty"`
	RedirectTitle       *string             `json:"RedirectTitle,omitempty" bson:"RedirectTitle,omitempty"`
	RedirectLogoURLInfo *ImageURLInfoStruct `json:"RedirectLogoURLInfo,omitempty" bson:"RedirectLogoURLInfo,omitempty"`
	AvailableRegions    []string            `json:"AvailableRegions,omitempty" bson:"AvailableRegions,omitempty"`
}

type LocalizedText struct {
	StarlingKey *string `json:"StarlingKey,omitempty" bson:"StarlingKey,omitempty"`
	Text        *string `json:"Text,omitempty" bson:"Text,omitempty"`
}

type CTAInfoStruct struct {
	CTAButton          *CTAButtonTypeEnum       `json:"CTAButton,omitempty" bson:"CTAButton,omitempty"`
	ActivateStartDate  *int64                   `json:"ActivateStartDate,omitempty" bson:"ActivateStartDate,omitempty"`
	ActivateEndDate    *int64                   `json:"ActivateEndDate,omitempty" bson:"ActivateEndDate,omitempty"`
	CTARedirectInfos   []*CTARedirectInfoStruct `json:"CTARedirectInfos,omitempty" bson:"CTARedirectInfos,omitempty"`
	CalendarEventTitle *string                  `json:"CalendarEventTitle,omitempty" bson:"CalendarEventTitle,omitempty"`
	ButtonText         *string                  `json:"ButtonText,omitempty" bson:"ButtonText,omitempty"`
	CTAButtonText      *LocalizedText           `json:"CTAButtonText,omitempty" bson:"CTAButtonText,omitempty"`
}

type TikTokAccountInfoStruct struct {
	TiktokUID        *int64   `json:"TiktokUID,omitempty" bson:"TiktokUID,omitempty"`
	AvailableRegions []string `json:"AvailableRegions,omitempty" bson:"AvailableRegions,omitempty"`
}

type SportsOfficialAccountInfoStruct struct {
	TikTokAccountInfo *TikTokAccountInfoStruct `json:"TikTokAccountInfo,omitempty" bson:"TikTokAccountInfo,omitempty"`
}

type SportsTeamInfoStruct struct {
	Name              *string                  `json:"Name,omitempty" bson:"Name,omitempty"`
	FullName          *string                  `json:"FullName,omitempty" bson:"FullName,omitempty"`
	AvatarURLInfo     *ImageURLInfoStruct      `json:"AvatarURLInfo,omitempty" bson:"AvatarURLInfo,omitempty"`
	TikTokAccountInfo *TikTokAccountInfoStruct `json:"TikTokAccountInfo,omitempty" bson:"TikTokAccountInfo,omitempty"`
}

type SportsTeamMemberInfoStruct struct {
	Name              *string                  `json:"Name,omitempty" bson:"Name,omitempty"`
	AvatarURLInfo     *ImageURLInfoStruct      `json:"AvatarURLInfo,omitempty" bson:"AvatarURLInfo,omitempty"`
	TikTokAccountInfo *TikTokAccountInfoStruct `json:"TikTokAccountInfo,omitempty" bson:"TikTokAccountInfo,omitempty"`
	TeamInfo          *SportsTeamInfoStruct    `json:"TeamInfo,omitempty" bson:"TeamInfo,omitempty"`
}

type WinnerInfoStruct struct {
	IsWinner    bool                `json:"IsWinner" bson:"IsWinner"`
	IconURLInfo *ImageURLInfoStruct `json:"IconURLInfo,omitempty" bson:"IconURLInfo,omitempty"`
}

type ScoreInfoForTeamBallStruct struct {
	TeamInfo   *SportsTeamInfoStruct `json:"TeamInfo,omitempty" bson:"TeamInfo,omitempty"`
	TotalScore *int64                `json:"TotalScore,omitempty" bson:"TotalScore,omitempty"`
}

type SportsLeagueBasicInfoStruct struct {
	PromotionEntityID *int64              `json:"PromotionEntityID,omitempty" bson:"PromotionEntityID,omitempty"`
	Title             *string             `json:"Title,omitempty" bson:"Title,omitempty"`
	Subtitle          *string             `json:"Subtitle,omitempty" bson:"Subtitle,omitempty"`
	PosterURLInfo     *ImageURLInfoStruct `json:"PosterURLInfo,omitempty" bson:"PosterURLInfo,omitempty"`
	ThumbnailURLInfo  *ImageURLInfoStruct `json:"ThumbnailURLInfo,omitempty" bson:"ThumbnailURLInfo,omitempty"`
	BackgroundColor   *string             `json:"BackgroundColor,omitempty" bson:"BackgroundColor,omitempty"`
}

type SportsLeagueAccountInfoStruct struct {
	OfficialAccountInfos []*SportsOfficialAccountInfoStruct `json:"OfficialAccountInfos,omitempty" bson:"OfficialAccountInfos,omitempty"`
	TeamInfos            []*SportsTeamInfoStruct            `json:"TeamInfos,omitempty" bson:"TeamInfos,omitempty"`
	PlayerInfos          []*SportsTeamMemberInfoStruct      `json:"PlayerInfos,omitempty" bson:"PlayerInfos,omitempty"`
}

type SportsLeagueScheduleInfoStruct struct {
	ScheduleTabInfos []*SportsLeagueScheduleTabInfoStruct `json:"ScheduleTabInfos,omitempty" bson:"ScheduleTabInfos,omitempty"`
}

type SportsLeagueScheduleTabInfoStruct struct {
	TabName           *LocalizedText                        `json:"TabName,omitempty" bson:"TabName,omitempty"`
	ScheduleItemInfos []*SportsLeagueScheduleItemInfoStruct `json:"ScheduleItemInfos,omitempty" bson:"ScheduleItemInfos,omitempty"`
	SortID            *int64                                `json:"SortID,omitempty" bson:"SortID,omitempty"`
}

type SportsLeagueScheduleItemInfoStruct struct {
	StartTime  *int64                          `json:"StartTime,omitempty" bson:"StartTime,omitempty"`
	EndTime    *int64                          `json:"EndTime,omitempty" bson:"EndTime,omitempty"`
	ScoreInfos []*ScoreInfoForTeamBallStruct   `json:"ScoreInfos,omitempty" bson:"ScoreInfos,omitempty"`
	CTAInfos   []*CTAInfoStruct                `json:"CTAInfos,omitempty" bson:"CTAInfos,omitempty"`
	Name       *string                         `json:"Name,omitempty" bson:"Name,omitempty"`
	ItemID     *int64                          `json:"ItemID,omitempty" bson:"ItemID,omitempty"`
	ID         *int64                          `json:"ID,omitempty" bson:"ID,omitempty"`
	Status     *SportsLeagueScheduleItemStatus `json:"Status,omitempty" bson:"Status,omitempty"`
	SortID     *int64                          `json:"SortID,omitempty" bson:"SortID,omitempty"`
}

type DirectNoticeMqMessage struct {
	PocOpenIds           []string  `json:"poc_open_ids"`
	SubmitTime           time.Time `json:"submit_time"`
	PublisherAccountName string    `json:"publisher_account_name"`
	AccountNum           int32     `json:"account_num"`
	NoticeTimes          int8      `json:"notice_times"`
}

var d = make(map[string]string)

func main() {
	data := `[
                        {
                            "MatchID": 7514596703178309638,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751472000,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Iceland",
                            "GuestName": "Finland",
                            "HostID": 7514510543143813176,
                            "GuestID": 7514510543143780408,
                            "HostLogoMaterialID": 7514610569014607877,
                            "GuestLogoMaterialID": 7514608433222123525,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/5ade94fef5d348b3a291a814515a727d",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/3dd8ea07a7f549b598af1c5c58587bb3"
                        },
                        {
                            "MatchID": 7514596703178326022,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751482800,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Switzerland",
                            "GuestName": "Norway",
                            "HostID": 7514510543143862328,
                            "GuestID": 7514510543143845944,
                            "HostLogoMaterialID": 7514615723180802054,
                            "GuestLogoMaterialID": 7514612489108373560,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/5b00314968a749b5a2f456d834c0edca",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/e17701f6a5c142a686e8a65555bb9e0d"
                        },
                        {
                            "MatchID": 7514596703178342406,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751558400,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Belgium",
                            "GuestName": "Italy",
                            "HostID": 7514556763879047174,
                            "GuestID": 7514510543143878712,
                            "HostLogoMaterialID": 7514614492714549254,
                            "GuestLogoMaterialID": 7514614545382473784,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/a8de265ee6b24e79979e84fc51544b17",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/d0b54a554b114d379d18b6da43b07980"
                        },
                        {
                            "MatchID": 7514596703178358790,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751569200,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Spain",
                            "GuestName": "Portugal",
                            "HostID": 7514556763879079942,
                            "GuestID": 7514556763879063558,
                            "HostLogoMaterialID": 7514616621156122630,
                            "GuestLogoMaterialID": 7514616621152124934,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/8c6b4de9ca89489f80f5ab4189a21f4c",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/9b20d13925f44b2ca41ab91af12ab26e"
                        },
                        {
                            "MatchID": 7514596703178375174,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751644800,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Denmark",
                            "GuestName": "Sweden",
                            "HostID": 7514556763879112710,
                            "GuestID": 7514556763879096326,
                            "HostLogoMaterialID": 7514611903411552262,
                            "GuestLogoMaterialID": 7514614474364600376,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/b292b3df0aee4f49bf47b9cd31572d79",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/da580920297f452da19baea2391658f2"
                        },
                        {
                            "MatchID": 7514597141218590725,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751655600,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Germany",
                            "GuestName": "Poland",
                            "HostID": 7514556763879145478,
                            "GuestID": 7514556763879129094,
                            "HostLogoMaterialID": 7514614545382604856,
                            "GuestLogoMaterialID": 7514614492714680326,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/9a60750855834132b351270c4aaabf28",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/fe81b922639b4160a25762a398f58b5e"
                        },
                        {
                            "MatchID": 7514597141218607109,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751731200,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Wales",
                            "GuestName": "Netherlands",
                            "HostID": 7514556763879178246,
                            "GuestID": 7514556763879161862,
                            "HostLogoMaterialID": 7514614474364731448,
                            "GuestLogoMaterialID": 7514615722941710392,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/f4ee57245f744e87a6e794c75d5b9206",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/a303095df66545f5a910fc22cff7ab8b"
                        },
                        {
                            "MatchID": 7514597141218623493,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751742000,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "France",
                            "GuestName": "England",
                            "HostID": 7514556763879211014,
                            "GuestID": 7514556763879194630,
                            "HostLogoMaterialID": 7514618081939865606,
                            "GuestLogoMaterialID": 7514615723180933126,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/3b891841107d4c809efc0dabc7de7285",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/955fe31920ec4394821114aa8bf4c4aa"
                        },
                        {
                            "MatchID": 7514597141218639877,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751817600,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Norway",
                            "GuestName": "Finland",
                            "HostID": 7514510543143845944,
                            "GuestID": 7514510543143780408,
                            "HostLogoMaterialID": 7514612489108373560,
                            "GuestLogoMaterialID": 7514608433222123525,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/e17701f6a5c142a686e8a65555bb9e0d",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/3dd8ea07a7f549b598af1c5c58587bb3"
                        },
                        {
                            "MatchID": 7514597141218656261,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751828400,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Switzerland",
                            "GuestName": "Iceland",
                            "HostID": 7514510543143862328,
                            "GuestID": 7514510543143813176,
                            "HostLogoMaterialID": 7514615723180802054,
                            "GuestLogoMaterialID": 7514610569014607877,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/5b00314968a749b5a2f456d834c0edca",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/5ade94fef5d348b3a291a814515a727d"
                        },
                        {
                            "MatchID": 7514597141218672645,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751904000,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Spain",
                            "GuestName": "Belgium",
                            "HostID": 7514556763879079942,
                            "GuestID": 7514556763879047174,
                            "HostLogoMaterialID": 7514616621156122630,
                            "GuestLogoMaterialID": 7514614492714549254,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/8c6b4de9ca89489f80f5ab4189a21f4c",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/a8de265ee6b24e79979e84fc51544b17"
                        },
                        {
                            "MatchID": 7514597141218689029,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751914800,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Portugal",
                            "GuestName": "Italy",
                            "HostID": 7514556763879063558,
                            "GuestID": 7514510543143878712,
                            "HostLogoMaterialID": 7514616621152124934,
                            "GuestLogoMaterialID": 7514614545382473784,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/9b20d13925f44b2ca41ab91af12ab26e",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/d0b54a554b114d379d18b6da43b07980"
                        },
                        {
                            "MatchID": 7514597141218705413,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1751990400,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Germany",
                            "GuestName": "Denmark",
                            "HostID": 7514556763879145478,
                            "GuestID": 7514556763879112710,
                            "HostLogoMaterialID": 7514614545382604856,
                            "GuestLogoMaterialID": 7514611903411552262,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/9a60750855834132b351270c4aaabf28",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/b292b3df0aee4f49bf47b9cd31572d79"
                        },
                        {
                            "MatchID": 7514597141218721797,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1752001200,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Poland",
                            "GuestName": "Sweden",
                            "HostID": 7514556763879129094,
                            "GuestID": 7514556763879096326,
                            "HostLogoMaterialID": 7514614492714680326,
                            "GuestLogoMaterialID": 7514614474364600376,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/fe81b922639b4160a25762a398f58b5e",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/da580920297f452da19baea2391658f2"
                        },
                        {
                            "MatchID": 7514597141218738181,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1752076800,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "England",
                            "GuestName": "Netherlands",
                            "HostID": 7514556763879194630,
                            "GuestID": 7514556763879161862,
                            "HostLogoMaterialID": 7514615723180933126,
                            "GuestLogoMaterialID": 7514615722941710392,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/955fe31920ec4394821114aa8bf4c4aa",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/a303095df66545f5a910fc22cff7ab8b"
                        },
                        {
                            "MatchID": 7514597141218754565,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1752087600,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "France",
                            "GuestName": "Wales",
                            "HostID": 7514556763879211014,
                            "GuestID": 7514556763879178246,
                            "HostLogoMaterialID": 7514618081939865606,
                            "GuestLogoMaterialID": 7514614474364731448,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/3b891841107d4c809efc0dabc7de7285",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/f4ee57245f744e87a6e794c75d5b9206"
                        },
                        {
                            "MatchID": 7514597141218770949,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1752174000,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Finland",
                            "GuestName": "Switzerland",
                            "HostID": 7514510543143780408,
                            "GuestID": 7514510543143862328,
                            "HostLogoMaterialID": 7514608433222123525,
                            "GuestLogoMaterialID": 7514615723180802054,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/3dd8ea07a7f549b598af1c5c58587bb3",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/5b00314968a749b5a2f456d834c0edca"
                        },
                        {
                            "MatchID": 7514597141218787333,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1752174000,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Norway",
                            "GuestName": "Iceland",
                            "HostID": 7514510543143845944,
                            "GuestID": 7514510543143813176,
                            "HostLogoMaterialID": 7514612489108373560,
                            "GuestLogoMaterialID": 7514610569014607877,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/e17701f6a5c142a686e8a65555bb9e0d",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/5ade94fef5d348b3a291a814515a727d"
                        },
                        {
                            "MatchID": 7514597141218803717,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1752260400,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Italy",
                            "GuestName": "Spain",
                            "HostID": 7514510543143878712,
                            "GuestID": 7514556763879079942,
                            "HostLogoMaterialID": 7514614545382473784,
                            "GuestLogoMaterialID": 7514616621156122630,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/d0b54a554b114d379d18b6da43b07980",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/8c6b4de9ca89489f80f5ab4189a21f4c"
                        },
                        {
                            "MatchID": 7514597141218820101,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1752260400,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Portugal",
                            "GuestName": "Belgium",
                            "HostID": 7514556763879063558,
                            "GuestID": 7514556763879047174,
                            "HostLogoMaterialID": 7514616621152124934,
                            "GuestLogoMaterialID": 7514614492714549254,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/9b20d13925f44b2ca41ab91af12ab26e",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/a8de265ee6b24e79979e84fc51544b17"
                        },
                        {
                            "MatchID": 7514597141218852869,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1752346800,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Poland",
                            "GuestName": "Denmark",
                            "HostID": 7514556763879129094,
                            "GuestID": 7514556763879112710,
                            "HostLogoMaterialID": 7514614492714680326,
                            "GuestLogoMaterialID": 7514611903411552262,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/fe81b922639b4160a25762a398f58b5e",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/b292b3df0aee4f49bf47b9cd31572d79"
                        },
                        {
                            "MatchID": 7514597141218836485,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1752346800,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Sweden",
                            "GuestName": "Germany",
                            "HostID": 7514556763879096326,
                            "GuestID": 7514556763879145478,
                            "HostLogoMaterialID": 7514614474364600376,
                            "GuestLogoMaterialID": 7514614545382604856,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/da580920297f452da19baea2391658f2",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/9a60750855834132b351270c4aaabf28"
                        },
                        {
                            "MatchID": 7514597141218885637,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1752433200,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "England",
                            "GuestName": "Wales",
                            "HostID": 7514556763879194630,
                            "GuestID": 7514556763879178246,
                            "HostLogoMaterialID": 7514615723180933126,
                            "GuestLogoMaterialID": 7514614474364731448,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/955fe31920ec4394821114aa8bf4c4aa",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/f4ee57245f744e87a6e794c75d5b9206"
                        },
                        {
                            "MatchID": 7514597141218869253,
                            "MediaID": 7514596703178145798,
                            "EventTime": 1752433200,
                            "GroupName": "",
                            "RoundName": "Group Stage Final tournament",
                            "PhaseName": "",
                            "MatchStatus": 0,
                            "HostName": "Netherlands",
                            "GuestName": "France",
                            "HostID": 7514556763879161862,
                            "GuestID": 7514556763879211014,
                            "HostLogoMaterialID": 7514615722941710392,
                            "GuestLogoMaterialID": 7514618081939865606,
                            "HostScore": 0,
                            "GuestScore": 0,
                            "Extra": "",
                            "HostLogoURI": "tos-maliva-i-f6c8vz16u4-us/a303095df66545f5a910fc22cff7ab8b",
                            "GuestLogoURI": "tos-maliva-i-f6c8vz16u4-us/3b891841107d4c809efc0dabc7de7285"
                        }
                    ]`
	var SportsMatchBaseInfos []*SportsMatchBaseInfo
	err := json.Unmarshal([]byte(data), &SportsMatchBaseInfos)
	if err != nil {
		return
	}

	groupRound := 3
	matchesPerDay := 8

	var ScheduleTabInfos []*SportsLeagueScheduleTabInfoStruct
	for i := 0; i < groupRound; i++ {
		var tab = &SportsLeagueScheduleTabInfoStruct{}
		matchDay := fmt.Sprintf("Matchday %d of 3", i+1)
		tab.TabName = &LocalizedText{Text: &matchDay}
		sortID := int64(i + 1)
		tab.SortID = &sortID
		for j := 0; j < matchesPerDay; j++ {
			index := i*matchesPerDay + j
			matchInfo := SportsMatchBaseInfos[index]
			itemInfo := &SportsLeagueScheduleItemInfoStruct{}
			itemInfo.StartTime = matchInfo.EventTime
			endTime := *itemInfo.StartTime + 9000
			itemInfo.EndTime = &endTime

			itemInfo.ID = matchInfo.MatchID

			//team1 := &ScoreInfoForTeamBallStruct{}
			//team1.TeamInfo = &SportsTeamInfoStruct{
			//	Name:     matchInfo.HostName,
			//	FullName: matchInfo.HostName,
			//	AvatarURLInfo: &ImageURLInfoStruct{
			//		MaterialID: matchInfo.HostLogoMaterialID,
			//	},
			//}
			//team2 := &ScoreInfoForTeamBallStruct{}
			//team2.TeamInfo = &SportsTeamInfoStruct{
			//	Name:     matchInfo.GuestName,
			//	FullName: matchInfo.GuestName,
			//	AvatarURLInfo: &ImageURLInfoStruct{
			//		MaterialID: matchInfo.GuestLogoMaterialID,
			//	},
			//}
			//itemInfo.ScoreInfos = append(itemInfo.ScoreInfos, team1)
			//itemInfo.ScoreInfos = append(itemInfo.ScoreInfos, team2)

			cta1 := &CTAInfoStruct{
				CTAButton:         gptr.Of(CTAButtonTypeEnum_ADD_TO_CALENDAR),
				ActivateStartDate: gptr.Of(int64(0)),
				ActivateEndDate:   matchInfo.EventTime,
				/*CTARedirectInfos: []*CTARedirectInfoStruct{
					{
						RedirectURL:
					},
				},*/
			}
			cta2 := &CTAInfoStruct{
				CTAButton:         gptr.Of(CTAButtonTypeEnum_WATCH_NOW),
				ActivateStartDate: matchInfo.EventTime,
				ActivateEndDate:   gptr.Of(*matchInfo.EventTime + 9000),
			}
			cta3 := &CTAInfoStruct{
				CTAButton:         gptr.Of(CTAButtonTypeEnum_REPLAY),
				ActivateStartDate: gptr.Of(*matchInfo.EventTime + 9000),
				ActivateEndDate:   gptr.Of(int64(1805806800)),
			}
			c := []*CTAInfoStruct{cta1, cta2, cta3}
			itemInfo.CTAInfos = c

			tab.ScheduleItemInfos = append(tab.ScheduleItemInfos, itemInfo)
		}
		ScheduleTabInfos = append(ScheduleTabInfos, tab)
	}

	marshal, err := json.Marshal(ScheduleTabInfos)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

func change(p []int) {
	p = append(p, 1)
}

func ddd(values ...interface{}) {
	fmt.Println(values)
}
