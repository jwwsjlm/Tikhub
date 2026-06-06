// Code generated from TikHub OpenAPI tags. DO NOT EDIT.
package Tikhub

import "context"

// Resources groups endpoints by OpenAPI tag, matching the official SDK style.
type Resources struct {
	// HealthCheck maps to the OpenAPI tag Health-Check.
	HealthCheck HealthCheckResource
	// TikHubDownloader maps to the OpenAPI tag TikHub-Downloader-API.
	TikHubDownloader TikHubDownloaderResource
	// TikHubUser maps to the OpenAPI tag TikHub-User-API.
	TikHubUser TikHubUserResource
	// TikTokWeb maps to the OpenAPI tag TikTok-Web-API.
	TikTokWeb TikTokWebResource
	// TikTokAppV3 maps to the OpenAPI tag TikTok-App-V3-API.
	TikTokAppV3 TikTokAppV3Resource
	// TikTokCreator maps to the OpenAPI tag TikTok-Creator-API.
	TikTokCreator TikTokCreatorResource
	// TikTokAnalytics maps to the OpenAPI tag TikTok-Analytics-API.
	TikTokAnalytics TikTokAnalyticsResource
	// TikTokAds maps to the OpenAPI tag TikTok-Ads-API.
	TikTokAds TikTokAdsResource
	// TikTokShopWeb maps to the OpenAPI tag TikTok-Shop-Web-API.
	TikTokShopWeb TikTokShopWebResource
	// TikTokInteraction maps to the OpenAPI tag TikTok-Interaction-API.
	TikTokInteraction TikTokInteractionResource
	// DouyinWeb maps to the OpenAPI tag Douyin-Web-API.
	DouyinWeb DouyinWebResource
	// DouyinAppV3 maps to the OpenAPI tag Douyin-App-V3-API.
	DouyinAppV3 DouyinAppV3Resource
	// DouyinCreator maps to the OpenAPI tag Douyin-Creator-API.
	DouyinCreator DouyinCreatorResource
	// DouyinCreatorV2 maps to the OpenAPI tag Douyin-Creator-V2-API.
	DouyinCreatorV2 DouyinCreatorV2Resource
	// DouyinIndex maps to the OpenAPI tag Douyin-Index-API.
	DouyinIndex DouyinIndexResource
	// DouyinSearch maps to the OpenAPI tag Douyin-Search-API.
	DouyinSearch DouyinSearchResource
	// DouyinBillboard maps to the OpenAPI tag Douyin-Billboard-API.
	DouyinBillboard DouyinBillboardResource
	// DouyinXingtu maps to the OpenAPI tag Douyin-Xingtu-API.
	DouyinXingtu DouyinXingtuResource
	// DouyinXingtuV2 maps to the OpenAPI tag Douyin-Xingtu-V2-API.
	DouyinXingtuV2 DouyinXingtuV2Resource
	// XiguaAppV2 maps to the OpenAPI tag Xigua-App-V2-API.
	XiguaAppV2 XiguaAppV2Resource
	// ToutiaoWeb maps to the OpenAPI tag Toutiao-Web-API.
	ToutiaoWeb ToutiaoWebResource
	// ToutiaoApp maps to the OpenAPI tag Toutiao-App-API.
	ToutiaoApp ToutiaoAppResource
	// XiaohongshuWebV3 maps to the OpenAPI tag Xiaohongshu-Web-V3-API.
	XiaohongshuWebV3 XiaohongshuWebV3Resource
	// XiaohongshuAppV2 maps to the OpenAPI tag Xiaohongshu-App-V2-API.
	XiaohongshuAppV2 XiaohongshuAppV2Resource
	// XiaohongshuApp maps to the OpenAPI tag Xiaohongshu-App-API.
	XiaohongshuApp XiaohongshuAppResource
	// XiaohongshuWebV2 maps to the OpenAPI tag Xiaohongshu-Web-V2-API.
	XiaohongshuWebV2 XiaohongshuWebV2Resource
	// XiaohongshuWeb maps to the OpenAPI tag Xiaohongshu-Web-API.
	XiaohongshuWeb XiaohongshuWebResource
	// Lemon8App maps to the OpenAPI tag Lemon8-App-API.
	Lemon8App Lemon8AppResource
	// KuaishouWeb maps to the OpenAPI tag Kuaishou-Web-API.
	KuaishouWeb KuaishouWebResource
	// KuaishouApp maps to the OpenAPI tag Kuaishou-App-API.
	KuaishouApp KuaishouAppResource
	// ZhihuWeb maps to the OpenAPI tag Zhihu-Web-API.
	ZhihuWeb ZhihuWebResource
	// PiPiXiaApp maps to the OpenAPI tag PiPiXia-App-API.
	PiPiXiaApp PiPiXiaAppResource
	// WeiboWeb maps to the OpenAPI tag Weibo-Web-API.
	WeiboWeb WeiboWebResource
	// WeiboWebV2 maps to the OpenAPI tag Weibo-Web-V2-API.
	WeiboWebV2 WeiboWebV2Resource
	// WeiboApp maps to the OpenAPI tag Weibo-App-API.
	WeiboApp WeiboAppResource
	// WeChatMediaPlatformWeb maps to the OpenAPI tag WeChat-Media-Platform-Web-API.
	WeChatMediaPlatformWeb WeChatMediaPlatformWebResource
	// WeChatChannels maps to the OpenAPI tag WeChat-Channels-API.
	WeChatChannels WeChatChannelsResource
	// InstagramV1 maps to the OpenAPI tag Instagram-V1-API.
	InstagramV1 InstagramV1Resource
	// InstagramV2 maps to the OpenAPI tag Instagram-V2-API.
	InstagramV2 InstagramV2Resource
	// InstagramV3 maps to the OpenAPI tag Instagram-V3-API.
	InstagramV3 InstagramV3Resource
	// YouTubeWeb maps to the OpenAPI tag YouTube-Web-API.
	YouTubeWeb YouTubeWebResource
	// YouTubeWebV2 maps to the OpenAPI tag YouTube-Web-V2-API.
	YouTubeWebV2 YouTubeWebV2Resource
	// LinkedInWeb maps to the OpenAPI tag LinkedIn-Web-API.
	LinkedInWeb LinkedInWebResource
	// LinkedInWebV2 maps to the OpenAPI tag LinkedIn-Web-V2-API.
	LinkedInWebV2 LinkedInWebV2Resource
	// BilibiliWeb maps to the OpenAPI tag Bilibili-Web-API.
	BilibiliWeb BilibiliWebResource
	// BilibiliApp maps to the OpenAPI tag Bilibili-App-API.
	BilibiliApp BilibiliAppResource
	// Sora2 maps to the OpenAPI tag Sora2-API.
	Sora2 Sora2Resource
	// TempMail maps to the OpenAPI tag Temp-Mail-API.
	TempMail TempMailResource
	// TwitterWeb maps to the OpenAPI tag Twitter-Web-API.
	TwitterWeb TwitterWebResource
	// ThreadsWeb maps to the OpenAPI tag Threads-Web-API.
	ThreadsWeb ThreadsWebResource
	// RedditApp maps to the OpenAPI tag Reddit-APP-API.
	RedditApp RedditAppResource
	// HybridParsing maps to the OpenAPI tag Hybrid-Parsing.
	HybridParsing HybridParsingResource
	// IOSShortcut maps to the OpenAPI tag iOS-Shortcut.
	IOSShortcut IOSShortcutResource
	// Demo maps to the OpenAPI tag Demo-API.
	Demo DemoResource
}

func (c *Client) initResources() {
	if c == nil {
		return
	}
	c.Resources = Resources{
		HealthCheck:            HealthCheckResource{client: c},
		TikHubDownloader:       TikHubDownloaderResource{client: c},
		TikHubUser:             TikHubUserResource{client: c},
		TikTokWeb:              TikTokWebResource{client: c},
		TikTokAppV3:            TikTokAppV3Resource{client: c},
		TikTokCreator:          TikTokCreatorResource{client: c},
		TikTokAnalytics:        TikTokAnalyticsResource{client: c},
		TikTokAds:              TikTokAdsResource{client: c},
		TikTokShopWeb:          TikTokShopWebResource{client: c},
		TikTokInteraction:      TikTokInteractionResource{client: c},
		DouyinWeb:              DouyinWebResource{client: c},
		DouyinAppV3:            DouyinAppV3Resource{client: c},
		DouyinCreator:          DouyinCreatorResource{client: c},
		DouyinCreatorV2:        DouyinCreatorV2Resource{client: c},
		DouyinIndex:            DouyinIndexResource{client: c},
		DouyinSearch:           DouyinSearchResource{client: c},
		DouyinBillboard:        DouyinBillboardResource{client: c},
		DouyinXingtu:           DouyinXingtuResource{client: c},
		DouyinXingtuV2:         DouyinXingtuV2Resource{client: c},
		XiguaAppV2:             XiguaAppV2Resource{client: c},
		ToutiaoWeb:             ToutiaoWebResource{client: c},
		ToutiaoApp:             ToutiaoAppResource{client: c},
		XiaohongshuWebV3:       XiaohongshuWebV3Resource{client: c},
		XiaohongshuAppV2:       XiaohongshuAppV2Resource{client: c},
		XiaohongshuApp:         XiaohongshuAppResource{client: c},
		XiaohongshuWebV2:       XiaohongshuWebV2Resource{client: c},
		XiaohongshuWeb:         XiaohongshuWebResource{client: c},
		Lemon8App:              Lemon8AppResource{client: c},
		KuaishouWeb:            KuaishouWebResource{client: c},
		KuaishouApp:            KuaishouAppResource{client: c},
		ZhihuWeb:               ZhihuWebResource{client: c},
		PiPiXiaApp:             PiPiXiaAppResource{client: c},
		WeiboWeb:               WeiboWebResource{client: c},
		WeiboWebV2:             WeiboWebV2Resource{client: c},
		WeiboApp:               WeiboAppResource{client: c},
		WeChatMediaPlatformWeb: WeChatMediaPlatformWebResource{client: c},
		WeChatChannels:         WeChatChannelsResource{client: c},
		InstagramV1:            InstagramV1Resource{client: c},
		InstagramV2:            InstagramV2Resource{client: c},
		InstagramV3:            InstagramV3Resource{client: c},
		YouTubeWeb:             YouTubeWebResource{client: c},
		YouTubeWebV2:           YouTubeWebV2Resource{client: c},
		LinkedInWeb:            LinkedInWebResource{client: c},
		LinkedInWebV2:          LinkedInWebV2Resource{client: c},
		BilibiliWeb:            BilibiliWebResource{client: c},
		BilibiliApp:            BilibiliAppResource{client: c},
		Sora2:                  Sora2Resource{client: c},
		TempMail:               TempMailResource{client: c},
		TwitterWeb:             TwitterWebResource{client: c},
		ThreadsWeb:             ThreadsWebResource{client: c},
		RedditApp:              RedditAppResource{client: c},
		HybridParsing:          HybridParsingResource{client: c},
		IOSShortcut:            IOSShortcutResource{client: c},
		Demo:                   DemoResource{client: c},
	}
}

// HealthCheckResource contains endpoints from the Health-Check tag.
type HealthCheckResource struct {
	client *Client
}

// HealthCheckCheckResponse is the response for GET /api/v1/health/check.
type HealthCheckCheckResponse = HealthCheckCheckIfTheServerRespondsToRequestsCorrectlyResponse

// Check 检查服务器是否正确响应请求 / Check if the server responds to requests correctly
//
// GET /api/v1/health/check
func (r HealthCheckResource) Check(ctx context.Context) (*HealthCheckCheckResponse, error) {
	return r.client.HealthCheckCheckIfTheServerRespondsToRequestsCorrectly(ctx)
}

// TikHubDownloaderResource contains endpoints from the TikHub-Downloader-API tag.
type TikHubDownloaderResource struct {
	client *Client
}

// TikHubDownloaderVersionResponse is the response for GET /api/v1/tikhub/downloader/version.
type TikHubDownloaderVersionResponse = TikHubDownloaderCheckForTikHubDownloaderVersionUpdatesResponse

// Version 检查TikHub下载器的版本更新 / Check for TikHub Downloader version updates
//
// GET /api/v1/tikhub/downloader/version
func (r TikHubDownloaderResource) Version(ctx context.Context) (*TikHubDownloaderVersionResponse, error) {
	return r.client.TikHubDownloaderCheckForTikHubDownloaderVersionUpdates(ctx)
}

// TikHubDownloaderRedirectDownloadResponse is the response for GET /api/v1/tikhub/downloader/redirect_download.
type TikHubDownloaderRedirectDownloadResponse = TikHubDownloaderRedirectToTheLatestVersionDownloadLinkResponse

// RedirectDownload 重定向到最新版本的下载链接 / Redirect to the latest version download link
//
// GET /api/v1/tikhub/downloader/redirect_download
func (r TikHubDownloaderResource) RedirectDownload(ctx context.Context) (*TikHubDownloaderRedirectDownloadResponse, error) {
	return r.client.TikHubDownloaderRedirectToTheLatestVersionDownloadLink(ctx)
}

// TikHubUserResource contains endpoints from the TikHub-User-API tag.
type TikHubUserResource struct {
	client *Client
}

// TikHubUserGetUserInfoResponse is the response for GET /api/v1/tikhub/user/get_user_info.
type TikHubUserGetUserInfoResponse = TikHubUserGetTikHubUserInfoResponse

// GetUserInfo 获取TikHub用户信息/Get TikHub user info
//
// GET /api/v1/tikhub/user/get_user_info
func (r TikHubUserResource) GetUserInfo(ctx context.Context) (*TikHubUserGetUserInfoResponse, error) {
	return r.client.TikHubUserGetTikHubUserInfo(ctx)
}

// GetUserDailyUsage 获取用户每日使用情况/Get user daily usage
//
// GET /api/v1/tikhub/user/get_user_daily_usage
func (r TikHubUserResource) GetUserDailyUsage(ctx context.Context) (*TikHubUserGetUserDailyUsageResponse, error) {
	return r.client.TikHubUserGetUserDailyUsage(ctx)
}

// CalculatePrice 计算价格/Calculate price
//
// GET /api/v1/tikhub/user/calculate_price
func (r TikHubUserResource) CalculatePrice(ctx context.Context, request TikHubUserCalculatePriceRequest) (*TikHubUserCalculatePriceResponse, error) {
	return r.client.TikHubUserCalculatePrice(ctx, request)
}

// TikHubUserGetTieredDiscountInfoResponse is the response for GET /api/v1/tikhub/user/get_tiered_discount_info.
type TikHubUserGetTieredDiscountInfoResponse = TikHubUserGetTieredDiscountPercentageInformationResponse

// GetTieredDiscountInfo 获取阶梯式折扣百分比信息/Get tiered discount percentage information
//
// GET /api/v1/tikhub/user/get_tiered_discount_info
func (r TikHubUserResource) GetTieredDiscountInfo(ctx context.Context) (*TikHubUserGetTieredDiscountInfoResponse, error) {
	return r.client.TikHubUserGetTieredDiscountPercentageInformation(ctx)
}

// TikHubUserGetEndpointInfoRequest is the request for GET /api/v1/tikhub/user/get_endpoint_info.
type TikHubUserGetEndpointInfoRequest = TikHubUserGetInformationOfAnEndpointRequest

// TikHubUserGetEndpointInfoResponse is the response for GET /api/v1/tikhub/user/get_endpoint_info.
type TikHubUserGetEndpointInfoResponse = TikHubUserGetInformationOfAnEndpointResponse

// GetEndpointInfo 获取一个端点的信息/Get information of an endpoint
//
// GET /api/v1/tikhub/user/get_endpoint_info
func (r TikHubUserResource) GetEndpointInfo(ctx context.Context, request TikHubUserGetEndpointInfoRequest) (*TikHubUserGetEndpointInfoResponse, error) {
	return r.client.TikHubUserGetInformationOfAnEndpoint(ctx, request)
}

// TikHubUserGetAllEndpointsInfoResponse is the response for GET /api/v1/tikhub/user/get_all_endpoints_info.
type TikHubUserGetAllEndpointsInfoResponse = TikHubUserGetAllEndpointsInformationResponse

// GetAllEndpointsInfo 获取所有端点信息/Get all endpoints information
//
// GET /api/v1/tikhub/user/get_all_endpoints_info
func (r TikHubUserResource) GetAllEndpointsInfo(ctx context.Context) (*TikHubUserGetAllEndpointsInfoResponse, error) {
	return r.client.TikHubUserGetAllEndpointsInformation(ctx)
}

// TikTokWebResource contains endpoints from the TikTok-Web-API tag.
type TikTokWebResource struct {
	client *Client
}

// TikTokWebFetchPostDetailRequest is the request for GET /api/v1/tiktok/web/fetch_post_detail.
type TikTokWebFetchPostDetailRequest = TikTokWebGetSingleVideoDataRequest

// TikTokWebFetchPostDetailResponse is the response for GET /api/v1/tiktok/web/fetch_post_detail.
type TikTokWebFetchPostDetailResponse = TikTokWebGetSingleVideoDataResponse

// FetchPostDetail 获取单个作品数据/Get single video data
//
// GET /api/v1/tiktok/web/fetch_post_detail
func (r TikTokWebResource) FetchPostDetail(ctx context.Context, request TikTokWebFetchPostDetailRequest) (*TikTokWebFetchPostDetailResponse, error) {
	return r.client.TikTokWebGetSingleVideoData(ctx, request)
}

// TikTokWebFetchPostDetailV2Request is the request for GET /api/v1/tiktok/web/fetch_post_detail_v2.
type TikTokWebFetchPostDetailV2Request = TikTokWebGetSingleVideoDataV2Request

// TikTokWebFetchPostDetailV2Response is the response for GET /api/v1/tiktok/web/fetch_post_detail_v2.
type TikTokWebFetchPostDetailV2Response = TikTokWebGetSingleVideoDataV2Response

// FetchPostDetailV2 获取单个作品数据 V2/Get single video data V2
//
// GET /api/v1/tiktok/web/fetch_post_detail_v2
func (r TikTokWebResource) FetchPostDetailV2(ctx context.Context, request TikTokWebFetchPostDetailV2Request) (*TikTokWebFetchPostDetailV2Response, error) {
	return r.client.TikTokWebGetSingleVideoDataV2(ctx, request)
}

// TikTokWebFetchExplorePostRequest is the request for GET /api/v1/tiktok/web/fetch_explore_post.
type TikTokWebFetchExplorePostRequest = TikTokWebGetExploreVideoDataRequest

// TikTokWebFetchExplorePostResponse is the response for GET /api/v1/tiktok/web/fetch_explore_post.
type TikTokWebFetchExplorePostResponse = TikTokWebGetExploreVideoDataResponse

// FetchExplorePost 获取探索作品数据/Get explore video data
//
// GET /api/v1/tiktok/web/fetch_explore_post
func (r TikTokWebResource) FetchExplorePost(ctx context.Context, request TikTokWebFetchExplorePostRequest) (*TikTokWebFetchExplorePostResponse, error) {
	return r.client.TikTokWebGetExploreVideoData(ctx, request)
}

// TikTokWebFetchTrendingPostResponse is the response for GET /api/v1/tiktok/web/fetch_trending_post.
type TikTokWebFetchTrendingPostResponse = TikTokWebGetDailyTrendingVideoDataResponse

// FetchTrendingPost 获取每日热门内容作品数据/Get daily trending video data
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/tiktok/web/fetch_trending_post
func (r TikTokWebResource) FetchTrendingPost(ctx context.Context) (*TikTokWebFetchTrendingPostResponse, error) {
	return r.client.TikTokWebGetDailyTrendingVideoData(ctx)
}

// TikTokWebFetchTrendingSearchwordsResponse is the response for GET /api/v1/tiktok/web/fetch_trending_searchwords.
type TikTokWebFetchTrendingSearchwordsResponse = TikTokWebGetDailyTrendingSearchWordsResponse

// FetchTrendingSearchwords 获取每日趋势搜索关键词/Get daily trending search words
//
// GET /api/v1/tiktok/web/fetch_trending_searchwords
func (r TikTokWebResource) FetchTrendingSearchwords(ctx context.Context) (*TikTokWebFetchTrendingSearchwordsResponse, error) {
	return r.client.TikTokWebGetDailyTrendingSearchWords(ctx)
}

// TikTokWebFetchUserProfileRequest is the request for GET /api/v1/tiktok/web/fetch_user_profile.
type TikTokWebFetchUserProfileRequest = TikTokWebGetUserProfileRequest

// TikTokWebFetchUserProfileResponse is the response for GET /api/v1/tiktok/web/fetch_user_profile.
type TikTokWebFetchUserProfileResponse = TikTokWebGetUserProfileResponse

// FetchUserProfile 获取用户的个人信息/Get user profile
//
// GET /api/v1/tiktok/web/fetch_user_profile
func (r TikTokWebResource) FetchUserProfile(ctx context.Context, request TikTokWebFetchUserProfileRequest) (*TikTokWebFetchUserProfileResponse, error) {
	return r.client.TikTokWebGetUserProfile(ctx, request)
}

// TikTokWebFetchUserPostRequest is the request for GET /api/v1/tiktok/web/fetch_user_post.
type TikTokWebFetchUserPostRequest = TikTokWebGetUserPostsRequest

// TikTokWebFetchUserPostResponse is the response for GET /api/v1/tiktok/web/fetch_user_post.
type TikTokWebFetchUserPostResponse = TikTokWebGetUserPostsResponse

// FetchUserPost 获取用户的作品列表/Get user posts
//
// GET /api/v1/tiktok/web/fetch_user_post
func (r TikTokWebResource) FetchUserPost(ctx context.Context, request TikTokWebFetchUserPostRequest) (*TikTokWebFetchUserPostResponse, error) {
	return r.client.TikTokWebGetUserPosts(ctx, request)
}

// TikTokWebFetchUserRepostRequest is the request for GET /api/v1/tiktok/web/fetch_user_repost.
type TikTokWebFetchUserRepostRequest = TikTokWebGetUserRepostsRequest

// TikTokWebFetchUserRepostResponse is the response for GET /api/v1/tiktok/web/fetch_user_repost.
type TikTokWebFetchUserRepostResponse = TikTokWebGetUserRepostsResponse

// FetchUserRepost 获取用户的转发作品列表/Get user reposts
//
// GET /api/v1/tiktok/web/fetch_user_repost
func (r TikTokWebResource) FetchUserRepost(ctx context.Context, request TikTokWebFetchUserRepostRequest) (*TikTokWebFetchUserRepostResponse, error) {
	return r.client.TikTokWebGetUserReposts(ctx, request)
}

// TikTokWebFetchUserLikeRequest is the request for GET /api/v1/tiktok/web/fetch_user_like.
type TikTokWebFetchUserLikeRequest = TikTokWebGetUserLikesRequest

// TikTokWebFetchUserLikeResponse is the response for GET /api/v1/tiktok/web/fetch_user_like.
type TikTokWebFetchUserLikeResponse = TikTokWebGetUserLikesResponse

// FetchUserLike 获取用户的点赞列表/Get user likes
//
// GET /api/v1/tiktok/web/fetch_user_like
func (r TikTokWebResource) FetchUserLike(ctx context.Context, request TikTokWebFetchUserLikeRequest) (*TikTokWebFetchUserLikeResponse, error) {
	return r.client.TikTokWebGetUserLikes(ctx, request)
}

// TikTokWebFetchUserCollectRequest is the request for GET /api/v1/tiktok/web/fetch_user_collect.
type TikTokWebFetchUserCollectRequest = TikTokWebGetUserFavoritesRequest

// TikTokWebFetchUserCollectResponse is the response for GET /api/v1/tiktok/web/fetch_user_collect.
type TikTokWebFetchUserCollectResponse = TikTokWebGetUserFavoritesResponse

// FetchUserCollect 获取用户的收藏列表/Get user favorites
//
// GET /api/v1/tiktok/web/fetch_user_collect
func (r TikTokWebResource) FetchUserCollect(ctx context.Context, request TikTokWebFetchUserCollectRequest) (*TikTokWebFetchUserCollectResponse, error) {
	return r.client.TikTokWebGetUserFavorites(ctx, request)
}

// TikTokWebFetchUserPlayListRequest is the request for GET /api/v1/tiktok/web/fetch_user_play_list.
type TikTokWebFetchUserPlayListRequest = TikTokWebGetUserPlayListRequest

// TikTokWebFetchUserPlayListResponse is the response for GET /api/v1/tiktok/web/fetch_user_play_list.
type TikTokWebFetchUserPlayListResponse = TikTokWebGetUserPlayListResponse

// FetchUserPlayList 获取用户的播放列表/Get user play list
//
// GET /api/v1/tiktok/web/fetch_user_play_list
func (r TikTokWebResource) FetchUserPlayList(ctx context.Context, request TikTokWebFetchUserPlayListRequest) (*TikTokWebFetchUserPlayListResponse, error) {
	return r.client.TikTokWebGetUserPlayList(ctx, request)
}

// TikTokWebFetchUserMixRequest is the request for GET /api/v1/tiktok/web/fetch_user_mix.
type TikTokWebFetchUserMixRequest = TikTokWebGetUserMixListRequest

// TikTokWebFetchUserMixResponse is the response for GET /api/v1/tiktok/web/fetch_user_mix.
type TikTokWebFetchUserMixResponse = TikTokWebGetUserMixListResponse

// FetchUserMix 获取用户的合辑列表/Get user mix list
//
// GET /api/v1/tiktok/web/fetch_user_mix
func (r TikTokWebResource) FetchUserMix(ctx context.Context, request TikTokWebFetchUserMixRequest) (*TikTokWebFetchUserMixResponse, error) {
	return r.client.TikTokWebGetUserMixList(ctx, request)
}

// TikTokWebFetchPostCommentRequest is the request for GET /api/v1/tiktok/web/fetch_post_comment.
type TikTokWebFetchPostCommentRequest = TikTokWebGetVideoCommentsRequest

// TikTokWebFetchPostCommentResponse is the response for GET /api/v1/tiktok/web/fetch_post_comment.
type TikTokWebFetchPostCommentResponse = TikTokWebGetVideoCommentsResponse

// FetchPostComment 获取作品的评论列表/Get video comments
//
// GET /api/v1/tiktok/web/fetch_post_comment
func (r TikTokWebResource) FetchPostComment(ctx context.Context, request TikTokWebFetchPostCommentRequest) (*TikTokWebFetchPostCommentResponse, error) {
	return r.client.TikTokWebGetVideoComments(ctx, request)
}

// TikTokWebFetchPostCommentReplyRequest is the request for GET /api/v1/tiktok/web/fetch_post_comment_reply.
type TikTokWebFetchPostCommentReplyRequest = TikTokWebGetVideoCommentRepliesRequest

// TikTokWebFetchPostCommentReplyResponse is the response for GET /api/v1/tiktok/web/fetch_post_comment_reply.
type TikTokWebFetchPostCommentReplyResponse = TikTokWebGetVideoCommentRepliesResponse

// FetchPostCommentReply 获取作品的评论回复列表/Get video comment replies
//
// GET /api/v1/tiktok/web/fetch_post_comment_reply
func (r TikTokWebResource) FetchPostCommentReply(ctx context.Context, request TikTokWebFetchPostCommentReplyRequest) (*TikTokWebFetchPostCommentReplyResponse, error) {
	return r.client.TikTokWebGetVideoCommentReplies(ctx, request)
}

// TikTokWebFetchUserFansRequest is the request for GET /api/v1/tiktok/web/fetch_user_fans.
type TikTokWebFetchUserFansRequest = TikTokWebGetUserFollowersRequest

// TikTokWebFetchUserFansResponse is the response for GET /api/v1/tiktok/web/fetch_user_fans.
type TikTokWebFetchUserFansResponse = TikTokWebGetUserFollowersResponse

// FetchUserFans 获取用户的粉丝列表/Get user followers
//
// GET /api/v1/tiktok/web/fetch_user_fans
func (r TikTokWebResource) FetchUserFans(ctx context.Context, request TikTokWebFetchUserFansRequest) (*TikTokWebFetchUserFansResponse, error) {
	return r.client.TikTokWebGetUserFollowers(ctx, request)
}

// TikTokWebFetchUserFollowRequest is the request for GET /api/v1/tiktok/web/fetch_user_follow.
type TikTokWebFetchUserFollowRequest = TikTokWebGetUserFollowingsRequest

// TikTokWebFetchUserFollowResponse is the response for GET /api/v1/tiktok/web/fetch_user_follow.
type TikTokWebFetchUserFollowResponse = TikTokWebGetUserFollowingsResponse

// FetchUserFollow 获取用户的关注列表/Get user followings
//
// GET /api/v1/tiktok/web/fetch_user_follow
func (r TikTokWebResource) FetchUserFollow(ctx context.Context, request TikTokWebFetchUserFollowRequest) (*TikTokWebFetchUserFollowResponse, error) {
	return r.client.TikTokWebGetUserFollowings(ctx, request)
}

// TikTokWebFetchUserLiveDetailRequest is the request for GET /api/v1/tiktok/web/fetch_user_live_detail.
type TikTokWebFetchUserLiveDetailRequest = TikTokWebGetUserLiveDetailsRequest

// TikTokWebFetchUserLiveDetailResponse is the response for GET /api/v1/tiktok/web/fetch_user_live_detail.
type TikTokWebFetchUserLiveDetailResponse = TikTokWebGetUserLiveDetailsResponse

// FetchUserLiveDetail 获取用户的直播详情/Get user live details
//
// GET /api/v1/tiktok/web/fetch_user_live_detail
func (r TikTokWebResource) FetchUserLiveDetail(ctx context.Context, request TikTokWebFetchUserLiveDetailRequest) (*TikTokWebFetchUserLiveDetailResponse, error) {
	return r.client.TikTokWebGetUserLiveDetails(ctx, request)
}

// TikTokWebFetchGeneralSearchRequest is the request for GET /api/v1/tiktok/web/fetch_general_search.
type TikTokWebFetchGeneralSearchRequest = TikTokWebGetGeneralSearchListRequest

// TikTokWebFetchGeneralSearchResponse is the response for GET /api/v1/tiktok/web/fetch_general_search.
type TikTokWebFetchGeneralSearchResponse = TikTokWebGetGeneralSearchListResponse

// FetchGeneralSearch 获取综合搜索列表/Get general search list
//
// GET /api/v1/tiktok/web/fetch_general_search
func (r TikTokWebResource) FetchGeneralSearch(ctx context.Context, request TikTokWebFetchGeneralSearchRequest) (*TikTokWebFetchGeneralSearchResponse, error) {
	return r.client.TikTokWebGetGeneralSearchList(ctx, request)
}

// TikTokWebFetchSearchKeywordSuggestRequest is the request for GET /api/v1/tiktok/web/fetch_search_keyword_suggest.
type TikTokWebFetchSearchKeywordSuggestRequest = TikTokWebSearchKeywordSuggestRequest

// TikTokWebFetchSearchKeywordSuggestResponse is the response for GET /api/v1/tiktok/web/fetch_search_keyword_suggest.
type TikTokWebFetchSearchKeywordSuggestResponse = TikTokWebSearchKeywordSuggestResponse

// FetchSearchKeywordSuggest 搜索关键字推荐/Search keyword suggest
//
// GET /api/v1/tiktok/web/fetch_search_keyword_suggest
func (r TikTokWebResource) FetchSearchKeywordSuggest(ctx context.Context, request TikTokWebFetchSearchKeywordSuggestRequest) (*TikTokWebFetchSearchKeywordSuggestResponse, error) {
	return r.client.TikTokWebSearchKeywordSuggest(ctx, request)
}

// TikTokWebFetchSearchUserRequest is the request for GET /api/v1/tiktok/web/fetch_search_user.
type TikTokWebFetchSearchUserRequest = TikTokWebSearchUserRequest

// TikTokWebFetchSearchUserResponse is the response for GET /api/v1/tiktok/web/fetch_search_user.
type TikTokWebFetchSearchUserResponse = TikTokWebSearchUserResponse

// FetchSearchUser 搜索用户/Search user
//
// GET /api/v1/tiktok/web/fetch_search_user
func (r TikTokWebResource) FetchSearchUser(ctx context.Context, request TikTokWebFetchSearchUserRequest) (*TikTokWebFetchSearchUserResponse, error) {
	return r.client.TikTokWebSearchUser(ctx, request)
}

// TikTokWebFetchSearchVideoRequest is the request for GET /api/v1/tiktok/web/fetch_search_video.
type TikTokWebFetchSearchVideoRequest = TikTokWebSearchVideoRequest

// TikTokWebFetchSearchVideoResponse is the response for GET /api/v1/tiktok/web/fetch_search_video.
type TikTokWebFetchSearchVideoResponse = TikTokWebSearchVideoResponse

// FetchSearchVideo 搜索视频/Search video
//
// GET /api/v1/tiktok/web/fetch_search_video
func (r TikTokWebResource) FetchSearchVideo(ctx context.Context, request TikTokWebFetchSearchVideoRequest) (*TikTokWebFetchSearchVideoResponse, error) {
	return r.client.TikTokWebSearchVideo(ctx, request)
}

// TikTokWebFetchSearchLiveRequest is the request for GET /api/v1/tiktok/web/fetch_search_live.
type TikTokWebFetchSearchLiveRequest = TikTokWebSearchLiveRequest

// TikTokWebFetchSearchLiveResponse is the response for GET /api/v1/tiktok/web/fetch_search_live.
type TikTokWebFetchSearchLiveResponse = TikTokWebSearchLiveResponse

// FetchSearchLive 搜索直播/Search live
//
// GET /api/v1/tiktok/web/fetch_search_live
func (r TikTokWebResource) FetchSearchLive(ctx context.Context, request TikTokWebFetchSearchLiveRequest) (*TikTokWebFetchSearchLiveResponse, error) {
	return r.client.TikTokWebSearchLive(ctx, request)
}

// TikTokWebFetchSearchPhotoRequest is the request for GET /api/v1/tiktok/web/fetch_search_photo.
type TikTokWebFetchSearchPhotoRequest = TikTokWebSearchPhotoRequest

// TikTokWebFetchSearchPhotoResponse is the response for GET /api/v1/tiktok/web/fetch_search_photo.
type TikTokWebFetchSearchPhotoResponse = TikTokWebSearchPhotoResponse

// FetchSearchPhoto 搜索照片/Search photo
//
// GET /api/v1/tiktok/web/fetch_search_photo
func (r TikTokWebResource) FetchSearchPhoto(ctx context.Context, request TikTokWebFetchSearchPhotoRequest) (*TikTokWebFetchSearchPhotoResponse, error) {
	return r.client.TikTokWebSearchPhoto(ctx, request)
}

// TikTokWebFetchTagDetailRequest is the request for GET /api/v1/tiktok/web/fetch_tag_detail.
type TikTokWebFetchTagDetailRequest = TikTokWebTagDetailRequest

// TikTokWebFetchTagDetailResponse is the response for GET /api/v1/tiktok/web/fetch_tag_detail.
type TikTokWebFetchTagDetailResponse = TikTokWebTagDetailResponse

// FetchTagDetail Tag详情/Tag Detail
//
// GET /api/v1/tiktok/web/fetch_tag_detail
func (r TikTokWebResource) FetchTagDetail(ctx context.Context, request TikTokWebFetchTagDetailRequest) (*TikTokWebFetchTagDetailResponse, error) {
	return r.client.TikTokWebTagDetail(ctx, request)
}

// TikTokWebFetchTagPostRequest is the request for GET /api/v1/tiktok/web/fetch_tag_post.
type TikTokWebFetchTagPostRequest = TikTokWebTagPostRequest

// TikTokWebFetchTagPostResponse is the response for GET /api/v1/tiktok/web/fetch_tag_post.
type TikTokWebFetchTagPostResponse = TikTokWebTagPostResponse

// FetchTagPost Tag作品/Tag Post
//
// GET /api/v1/tiktok/web/fetch_tag_post
func (r TikTokWebResource) FetchTagPost(ctx context.Context, request TikTokWebFetchTagPostRequest) (*TikTokWebFetchTagPostResponse, error) {
	return r.client.TikTokWebTagPost(ctx, request)
}

// TikTokWebFetchHomeFeedRequest is the request for POST /api/v1/tiktok/web/fetch_home_feed.
type TikTokWebFetchHomeFeedRequest = TikTokWebHomeFeedRequest

// TikTokWebFetchHomeFeedResponse is the response for POST /api/v1/tiktok/web/fetch_home_feed.
type TikTokWebFetchHomeFeedResponse = TikTokWebHomeFeedResponse

// FetchHomeFeed 首页推荐作品/Home Feed
//
// POST /api/v1/tiktok/web/fetch_home_feed
func (r TikTokWebResource) FetchHomeFeed(ctx context.Context, request TikTokWebFetchHomeFeedRequest) (*TikTokWebFetchHomeFeedResponse, error) {
	return r.client.TikTokWebHomeFeed(ctx, request)
}

// GenerateRealMSToken 生成真实msToken/Generate real msToken
//
// GET /api/v1/tiktok/web/generate_real_msToken
func (r TikTokWebResource) GenerateRealMSToken(ctx context.Context, request TikTokWebGenerateRealMSTokenRequest) (*TikTokWebGenerateRealMSTokenResponse, error) {
	return r.client.TikTokWebGenerateRealMSToken(ctx, request)
}

// EncryptStrData 加密strData/Encrypt strData
//
// GET /api/v1/tiktok/web/encrypt_strData
func (r TikTokWebResource) EncryptStrData(ctx context.Context, request TikTokWebEncryptStrDataRequest) (*TikTokWebEncryptStrDataResponse, error) {
	return r.client.TikTokWebEncryptStrData(ctx, request)
}

// DecryptStrData 解密strData/Decrypt strData
//
// GET /api/v1/tiktok/web/decrypt_strData
func (r TikTokWebResource) DecryptStrData(ctx context.Context, request TikTokWebDecryptStrDataRequest) (*TikTokWebDecryptStrDataResponse, error) {
	return r.client.TikTokWebDecryptStrData(ctx, request)
}

// TikTokWebGenerateFingerprintRequest is the request for GET /api/v1/tiktok/web/generate_fingerprint.
type TikTokWebGenerateFingerprintRequest = TikTokWebGenerateBrowserFingerprintRequest

// TikTokWebGenerateFingerprintResponse is the response for GET /api/v1/tiktok/web/generate_fingerprint.
type TikTokWebGenerateFingerprintResponse = TikTokWebGenerateBrowserFingerprintResponse

// GenerateFingerprint 生成浏览器指纹/Generate browser fingerprint
//
// GET /api/v1/tiktok/web/generate_fingerprint
func (r TikTokWebResource) GenerateFingerprint(ctx context.Context, request TikTokWebGenerateFingerprintRequest) (*TikTokWebGenerateFingerprintResponse, error) {
	return r.client.TikTokWebGenerateBrowserFingerprint(ctx, request)
}

// TikTokWebGenerateWebidRequest is the request for GET /api/v1/tiktok/web/generate_webid.
type TikTokWebGenerateWebidRequest = TikTokWebGenerateWebIDRequest

// TikTokWebGenerateWebidResponse is the response for GET /api/v1/tiktok/web/generate_webid.
type TikTokWebGenerateWebidResponse = TikTokWebGenerateWebIDResponse

// GenerateWebid 生成web_id/Generate web_id
//
// GET /api/v1/tiktok/web/generate_webid
func (r TikTokWebResource) GenerateWebid(ctx context.Context, request TikTokWebGenerateWebidRequest) (*TikTokWebGenerateWebidResponse, error) {
	return r.client.TikTokWebGenerateWebID(ctx, request)
}

// GenerateTtwid 生成ttwid/Generate ttwid
//
// GET /api/v1/tiktok/web/generate_ttwid
func (r TikTokWebResource) GenerateTtwid(ctx context.Context, request TikTokWebGenerateTtwidRequest) (*TikTokWebGenerateTtwidResponse, error) {
	return r.client.TikTokWebGenerateTtwid(ctx, request)
}

// GenerateXBogus 生成 XBogus/Generate XBogus
//
// POST /api/v1/tiktok/web/generate_xbogus
func (r TikTokWebResource) GenerateXBogus(ctx context.Context, request TikTokWebGenerateXBogusRequest) (*TikTokWebGenerateXBogusResponse, error) {
	return r.client.TikTokWebGenerateXBogus(ctx, request)
}

// GenerateXGnarly 生成 XGnarly /Generate XGnarly
//
// POST /api/v1/tiktok/web/generate_xgnarly
func (r TikTokWebResource) GenerateXGnarly(ctx context.Context, request TikTokWebGenerateXGnarlyRequest) (*TikTokWebGenerateXGnarlyResponse, error) {
	return r.client.TikTokWebGenerateXGnarly(ctx, request)
}

// GenerateXGnarlyAndXBogus 生成 XGnarly 和 XBogus /Generate XGnarly and XBogus
//
// POST /api/v1/tiktok/web/generate_xgnarly_and_xbogus
func (r TikTokWebResource) GenerateXGnarlyAndXBogus(ctx context.Context, request TikTokWebGenerateXGnarlyAndXBogusRequest) (*TikTokWebGenerateXGnarlyAndXBogusResponse, error) {
	return r.client.TikTokWebGenerateXGnarlyAndXBogus(ctx, request)
}

// TikTokWebGenerateXMssdkInfoRequest is the request for POST /api/v1/tiktok/web/generate_x_mssdk_info.
type TikTokWebGenerateXMssdkInfoRequest = TikTokWebGenerateXMSSDKInfoRequest

// TikTokWebGenerateXMssdkInfoResponse is the response for POST /api/v1/tiktok/web/generate_x_mssdk_info.
type TikTokWebGenerateXMssdkInfoResponse = TikTokWebGenerateXMSSDKInfoResponse

// GenerateXMssdkInfo 生成 X-Mssdk-Info /Generate X-Mssdk-Info
//
// POST /api/v1/tiktok/web/generate_x_mssdk_info
func (r TikTokWebResource) GenerateXMssdkInfo(ctx context.Context, request TikTokWebGenerateXMssdkInfoRequest) (*TikTokWebGenerateXMssdkInfoResponse, error) {
	return r.client.TikTokWebGenerateXMSSDKInfo(ctx, request)
}

// TikTokWebGetUserIDRequest is the request for GET /api/v1/tiktok/web/get_user_id.
type TikTokWebGetUserIDRequest = TikTokWebExtractUserUserIDRequest

// TikTokWebGetUserIDResponse is the response for GET /api/v1/tiktok/web/get_user_id.
type TikTokWebGetUserIDResponse = TikTokWebExtractUserUserIDResponse

// GetUserID 提取用户user_id/Extract user user_id
//
// GET /api/v1/tiktok/web/get_user_id
func (r TikTokWebResource) GetUserID(ctx context.Context, request TikTokWebGetUserIDRequest) (*TikTokWebGetUserIDResponse, error) {
	return r.client.TikTokWebExtractUserUserID(ctx, request)
}

// TikTokWebGetSecUserIDRequest is the request for GET /api/v1/tiktok/web/get_sec_user_id.
type TikTokWebGetSecUserIDRequest = TikTokWebExtractUserSecUserIDRequest

// TikTokWebGetSecUserIDResponse is the response for GET /api/v1/tiktok/web/get_sec_user_id.
type TikTokWebGetSecUserIDResponse = TikTokWebExtractUserSecUserIDResponse

// GetSecUserID 提取用户sec_user_id/Extract user sec_user_id
//
// GET /api/v1/tiktok/web/get_sec_user_id
func (r TikTokWebResource) GetSecUserID(ctx context.Context, request TikTokWebGetSecUserIDRequest) (*TikTokWebGetSecUserIDResponse, error) {
	return r.client.TikTokWebExtractUserSecUserID(ctx, request)
}

// TikTokWebGetAllSecUserIDRequest is the request for POST /api/v1/tiktok/web/get_all_sec_user_id.
type TikTokWebGetAllSecUserIDRequest = TikTokWebExtractListUserSecUserIDRequest

// TikTokWebGetAllSecUserIDResponse is the response for POST /api/v1/tiktok/web/get_all_sec_user_id.
type TikTokWebGetAllSecUserIDResponse = TikTokWebExtractListUserSecUserIDResponse

// GetAllSecUserID 提取列表用户sec_user_id/Extract list user sec_user_id
//
// POST /api/v1/tiktok/web/get_all_sec_user_id
func (r TikTokWebResource) GetAllSecUserID(ctx context.Context, request TikTokWebGetAllSecUserIDRequest) (*TikTokWebGetAllSecUserIDResponse, error) {
	return r.client.TikTokWebExtractListUserSecUserID(ctx, request)
}

// TikTokWebGetAwemeIDRequest is the request for GET /api/v1/tiktok/web/get_aweme_id.
type TikTokWebGetAwemeIDRequest = TikTokWebExtractSingleVideoIDRequest

// TikTokWebGetAwemeIDResponse is the response for GET /api/v1/tiktok/web/get_aweme_id.
type TikTokWebGetAwemeIDResponse = TikTokWebExtractSingleVideoIDResponse

// GetAwemeID 提取单个作品id/Extract single video id
//
// GET /api/v1/tiktok/web/get_aweme_id
func (r TikTokWebResource) GetAwemeID(ctx context.Context, request TikTokWebGetAwemeIDRequest) (*TikTokWebGetAwemeIDResponse, error) {
	return r.client.TikTokWebExtractSingleVideoID(ctx, request)
}

// TikTokWebGetAllAwemeIDRequest is the request for POST /api/v1/tiktok/web/get_all_aweme_id.
type TikTokWebGetAllAwemeIDRequest = TikTokWebExtractListVideoIDRequest

// TikTokWebGetAllAwemeIDResponse is the response for POST /api/v1/tiktok/web/get_all_aweme_id.
type TikTokWebGetAllAwemeIDResponse = TikTokWebExtractListVideoIDResponse

// GetAllAwemeID 提取列表作品id/Extract list video id
//
// POST /api/v1/tiktok/web/get_all_aweme_id
func (r TikTokWebResource) GetAllAwemeID(ctx context.Context, request TikTokWebGetAllAwemeIDRequest) (*TikTokWebGetAllAwemeIDResponse, error) {
	return r.client.TikTokWebExtractListVideoID(ctx, request)
}

// TikTokWebGetUniqueIDRequest is the request for GET /api/v1/tiktok/web/get_unique_id.
type TikTokWebGetUniqueIDRequest = TikTokWebGetUserUniqueIDRequest

// TikTokWebGetUniqueIDResponse is the response for GET /api/v1/tiktok/web/get_unique_id.
type TikTokWebGetUniqueIDResponse = TikTokWebGetUserUniqueIDResponse

// GetUniqueID 获取用户unique_id/Get user unique_id
//
// GET /api/v1/tiktok/web/get_unique_id
func (r TikTokWebResource) GetUniqueID(ctx context.Context, request TikTokWebGetUniqueIDRequest) (*TikTokWebGetUniqueIDResponse, error) {
	return r.client.TikTokWebGetUserUniqueID(ctx, request)
}

// TikTokWebGetAllUniqueIDRequest is the request for POST /api/v1/tiktok/web/get_all_unique_id.
type TikTokWebGetAllUniqueIDRequest = TikTokWebGetListUniqueIDRequest

// TikTokWebGetAllUniqueIDResponse is the response for POST /api/v1/tiktok/web/get_all_unique_id.
type TikTokWebGetAllUniqueIDResponse = TikTokWebGetListUniqueIDResponse

// GetAllUniqueID 获取列表unique_id/Get list unique_id
//
// POST /api/v1/tiktok/web/get_all_unique_id
func (r TikTokWebResource) GetAllUniqueID(ctx context.Context, request TikTokWebGetAllUniqueIDRequest) (*TikTokWebGetAllUniqueIDResponse, error) {
	return r.client.TikTokWebGetListUniqueID(ctx, request)
}

// TikTokWebTikTokLiveRoomRequest is the request for GET /api/v1/tiktok/web/tiktok_live_room.
type TikTokWebTikTokLiveRoomRequest = TikTokWebExtractLiveRoomDanmakuRequest

// TikTokWebTikTokLiveRoomResponse is the response for GET /api/v1/tiktok/web/tiktok_live_room.
type TikTokWebTikTokLiveRoomResponse = TikTokWebExtractLiveRoomDanmakuResponse

// TikTokLiveRoom 提取直播间弹幕/Extract live room danmaku
//
// GET /api/v1/tiktok/web/tiktok_live_room
func (r TikTokWebResource) TikTokLiveRoom(ctx context.Context, request TikTokWebTikTokLiveRoomRequest) (*TikTokWebTikTokLiveRoomResponse, error) {
	return r.client.TikTokWebExtractLiveRoomDanmaku(ctx, request)
}

// TikTokWebFetchLiveImFetchRequest is the request for GET /api/v1/tiktok/web/fetch_live_im_fetch.
type TikTokWebFetchLiveImFetchRequest = TikTokWebTiktokLiveRoomDanmakuParametersRequest

// TikTokWebFetchLiveImFetchResponse is the response for GET /api/v1/tiktok/web/fetch_live_im_fetch.
type TikTokWebFetchLiveImFetchResponse = TikTokWebTiktokLiveRoomDanmakuParametersResponse

// FetchLiveImFetch TikTok直播间弹幕参数获取/tiktok live room danmaku parameters
//
// GET /api/v1/tiktok/web/fetch_live_im_fetch
func (r TikTokWebResource) FetchLiveImFetch(ctx context.Context, request TikTokWebFetchLiveImFetchRequest) (*TikTokWebFetchLiveImFetchResponse, error) {
	return r.client.TikTokWebTiktokLiveRoomDanmakuParameters(ctx, request)
}

// TikTokWebGenerateWssXbSignatureRequest is the request for GET /api/v1/tiktok/web/generate_wss_xb_signature.
type TikTokWebGenerateWssXbSignatureRequest = TikTokWebGenerateTikTokWSSXBogusSignatureRequest

// TikTokWebGenerateWssXbSignatureResponse is the response for GET /api/v1/tiktok/web/generate_wss_xb_signature.
type TikTokWebGenerateWssXbSignatureResponse = TikTokWebGenerateTikTokWSSXBogusSignatureResponse

// GenerateWssXbSignature 生成TikTok WSS X-Bogus签名/Generate TikTok WSS X-Bogus signature
//
// GET /api/v1/tiktok/web/generate_wss_xb_signature
func (r TikTokWebResource) GenerateWssXbSignature(ctx context.Context, request TikTokWebGenerateWssXbSignatureRequest) (*TikTokWebGenerateWssXbSignatureResponse, error) {
	return r.client.TikTokWebGenerateTikTokWSSXBogusSignature(ctx, request)
}

// TikTokWebGetLiveRoomIDRequest is the request for GET /api/v1/tiktok/web/get_live_room_id.
type TikTokWebGetLiveRoomIDRequest = TikTokWebExtractLiveRoomIDFromLiveRoomLinkRequest

// TikTokWebGetLiveRoomIDResponse is the response for GET /api/v1/tiktok/web/get_live_room_id.
type TikTokWebGetLiveRoomIDResponse = TikTokWebExtractLiveRoomIDFromLiveRoomLinkResponse

// GetLiveRoomID 根据直播间链接提取直播间ID/Extract live room ID from live room link
//
// GET /api/v1/tiktok/web/get_live_room_id
func (r TikTokWebResource) GetLiveRoomID(ctx context.Context, request TikTokWebGetLiveRoomIDRequest) (*TikTokWebGetLiveRoomIDResponse, error) {
	return r.client.TikTokWebExtractLiveRoomIDFromLiveRoomLink(ctx, request)
}

// TikTokWebFetchCheckLiveAliveRequest is the request for GET /api/v1/tiktok/web/fetch_check_live_alive.
type TikTokWebFetchCheckLiveAliveRequest = TikTokWebLiveRoomStartStatusCheckRequest

// TikTokWebFetchCheckLiveAliveResponse is the response for GET /api/v1/tiktok/web/fetch_check_live_alive.
type TikTokWebFetchCheckLiveAliveResponse = TikTokWebLiveRoomStartStatusCheckResponse

// FetchCheckLiveAlive 直播间开播状态检测/Live room start status check
//
// GET /api/v1/tiktok/web/fetch_check_live_alive
func (r TikTokWebResource) FetchCheckLiveAlive(ctx context.Context, request TikTokWebFetchCheckLiveAliveRequest) (*TikTokWebFetchCheckLiveAliveResponse, error) {
	return r.client.TikTokWebLiveRoomStartStatusCheck(ctx, request)
}

// TikTokWebFetchBatchCheckLiveAliveRequest is the request for GET /api/v1/tiktok/web/fetch_batch_check_live_alive.
type TikTokWebFetchBatchCheckLiveAliveRequest = TikTokWebBatchLiveRoomStartStatusCheckRequest

// TikTokWebFetchBatchCheckLiveAliveResponse is the response for GET /api/v1/tiktok/web/fetch_batch_check_live_alive.
type TikTokWebFetchBatchCheckLiveAliveResponse = TikTokWebBatchLiveRoomStartStatusCheckResponse

// FetchBatchCheckLiveAlive 批量直播间开播状态检测/Batch live room start status check
//
// GET /api/v1/tiktok/web/fetch_batch_check_live_alive
func (r TikTokWebResource) FetchBatchCheckLiveAlive(ctx context.Context, request TikTokWebFetchBatchCheckLiveAliveRequest) (*TikTokWebFetchBatchCheckLiveAliveResponse, error) {
	return r.client.TikTokWebBatchLiveRoomStartStatusCheck(ctx, request)
}

// TikTokWebFetchTikTokLiveDataRequest is the request for GET /api/v1/tiktok/web/fetch_tiktok_live_data.
type TikTokWebFetchTikTokLiveDataRequest = TikTokWebGetLiveRoomInformationViaLiveLinkRequest

// TikTokWebFetchTikTokLiveDataResponse is the response for GET /api/v1/tiktok/web/fetch_tiktok_live_data.
type TikTokWebFetchTikTokLiveDataResponse = TikTokWebGetLiveRoomInformationViaLiveLinkResponse

// FetchTikTokLiveData 通过直播链接获取直播间信息/Get live room information via live link
//
// GET /api/v1/tiktok/web/fetch_tiktok_live_data
func (r TikTokWebResource) FetchTikTokLiveData(ctx context.Context, request TikTokWebFetchTikTokLiveDataRequest) (*TikTokWebFetchTikTokLiveDataResponse, error) {
	return r.client.TikTokWebGetLiveRoomInformationViaLiveLink(ctx, request)
}

// TikTokWebFetchLiveRecommendRequest is the request for GET /api/v1/tiktok/web/fetch_live_recommend.
type TikTokWebFetchLiveRecommendRequest = TikTokWebGetLiveRoomHomepageRecommendationListRequest

// TikTokWebFetchLiveRecommendResponse is the response for GET /api/v1/tiktok/web/fetch_live_recommend.
type TikTokWebFetchLiveRecommendResponse = TikTokWebGetLiveRoomHomepageRecommendationListResponse

// FetchLiveRecommend 获取直播间首页推荐列表/Get live room homepage recommendation list
//
// GET /api/v1/tiktok/web/fetch_live_recommend
func (r TikTokWebResource) FetchLiveRecommend(ctx context.Context, request TikTokWebFetchLiveRecommendRequest) (*TikTokWebFetchLiveRecommendResponse, error) {
	return r.client.TikTokWebGetLiveRoomHomepageRecommendationList(ctx, request)
}

// TikTokWebFetchLiveGiftListRequest is the request for GET /api/v1/tiktok/web/fetch_live_gift_list.
type TikTokWebFetchLiveGiftListRequest = TikTokWebGetLiveRoomGiftListRequest

// TikTokWebFetchLiveGiftListResponse is the response for GET /api/v1/tiktok/web/fetch_live_gift_list.
type TikTokWebFetchLiveGiftListResponse = TikTokWebGetLiveRoomGiftListResponse

// FetchLiveGiftList 获取直播间礼物列表/Get live room gift list
//
// GET /api/v1/tiktok/web/fetch_live_gift_list
func (r TikTokWebResource) FetchLiveGiftList(ctx context.Context, request TikTokWebFetchLiveGiftListRequest) (*TikTokWebFetchLiveGiftListResponse, error) {
	return r.client.TikTokWebGetLiveRoomGiftList(ctx, request)
}

// TikTokWebFetchSsoLoginQrcodeRequest is the request for GET /api/v1/tiktok/web/fetch_sso_login_qrcode.
type TikTokWebFetchSsoLoginQrcodeRequest = TikTokWebGetSSOLoginQRCodeRequest

// TikTokWebFetchSsoLoginQrcodeResponse is the response for GET /api/v1/tiktok/web/fetch_sso_login_qrcode.
type TikTokWebFetchSsoLoginQrcodeResponse = TikTokWebGetSSOLoginQRCodeResponse

// FetchSsoLoginQrcode 获取SSO登录二维码/Get SSO login QR code
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/tiktok/web/fetch_sso_login_qrcode
func (r TikTokWebResource) FetchSsoLoginQrcode(ctx context.Context, request TikTokWebFetchSsoLoginQrcodeRequest) (*TikTokWebFetchSsoLoginQrcodeResponse, error) {
	return r.client.TikTokWebGetSSOLoginQRCode(ctx, request)
}

// TikTokWebFetchSsoLoginStatusRequest is the request for GET /api/v1/tiktok/web/fetch_sso_login_status.
type TikTokWebFetchSsoLoginStatusRequest = TikTokWebGetSSOLoginStatusRequest

// TikTokWebFetchSsoLoginStatusResponse is the response for GET /api/v1/tiktok/web/fetch_sso_login_status.
type TikTokWebFetchSsoLoginStatusResponse = TikTokWebGetSSOLoginStatusResponse

// FetchSsoLoginStatus 获取SSO登录状态/Get SSO login status
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/tiktok/web/fetch_sso_login_status
func (r TikTokWebResource) FetchSsoLoginStatus(ctx context.Context, request TikTokWebFetchSsoLoginStatusRequest) (*TikTokWebFetchSsoLoginStatusResponse, error) {
	return r.client.TikTokWebGetSSOLoginStatus(ctx, request)
}

// TikTokWebFetchSsoLoginAuthRequest is the request for GET /api/v1/tiktok/web/fetch_sso_login_auth.
type TikTokWebFetchSsoLoginAuthRequest = TikTokWebAuthenticateSSOLoginRequest

// TikTokWebFetchSsoLoginAuthResponse is the response for GET /api/v1/tiktok/web/fetch_sso_login_auth.
type TikTokWebFetchSsoLoginAuthResponse = TikTokWebAuthenticateSSOLoginResponse

// FetchSsoLoginAuth 认证SSO登录/Authenticate SSO login
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/tiktok/web/fetch_sso_login_auth
func (r TikTokWebResource) FetchSsoLoginAuth(ctx context.Context, request TikTokWebFetchSsoLoginAuthRequest) (*TikTokWebFetchSsoLoginAuthResponse, error) {
	return r.client.TikTokWebAuthenticateSSOLogin(ctx, request)
}

// GenerateHashedID 生成哈希ID/Generate hashed ID
//
// GET /api/v1/tiktok/web/generate_hashed_id
func (r TikTokWebResource) GenerateHashedID(ctx context.Context, request TikTokWebGenerateHashedIDRequest) (*TikTokWebGenerateHashedIDResponse, error) {
	return r.client.TikTokWebGenerateHashedID(ctx, request)
}

// TikTokWebFetchGiftNameByIDRequest is the request for POST /api/v1/tiktok/web/fetch_gift_name_by_id.
type TikTokWebFetchGiftNameByIDRequest = TikTokWebGetGiftNameByGiftIDRequest

// TikTokWebFetchGiftNameByIDResponse is the response for POST /api/v1/tiktok/web/fetch_gift_name_by_id.
type TikTokWebFetchGiftNameByIDResponse = TikTokWebGetGiftNameByGiftIDResponse

// FetchGiftNameByID 根据Gift ID查询礼物名称/Get gift name by gift ID
//
// POST /api/v1/tiktok/web/fetch_gift_name_by_id
func (r TikTokWebResource) FetchGiftNameByID(ctx context.Context, request TikTokWebFetchGiftNameByIDRequest) (*TikTokWebFetchGiftNameByIDResponse, error) {
	return r.client.TikTokWebGetGiftNameByGiftID(ctx, request)
}

// TikTokWebFetchGiftNamesByIdsRequest is the request for POST /api/v1/tiktok/web/fetch_gift_names_by_ids.
type TikTokWebFetchGiftNamesByIdsRequest = TikTokWebN50BatchGetGiftNamesByGiftIDsRequest

// TikTokWebFetchGiftNamesByIdsResponse is the response for POST /api/v1/tiktok/web/fetch_gift_names_by_ids.
type TikTokWebFetchGiftNamesByIdsResponse = TikTokWebN50BatchGetGiftNamesByGiftIDsResponse

// FetchGiftNamesByIds 批量查询Gift ID对应的礼物名称($0.025/次,建议50个)/Batch get gift names by gift IDs ($0.025/call, suggest 50)
//
// POST /api/v1/tiktok/web/fetch_gift_names_by_ids
func (r TikTokWebResource) FetchGiftNamesByIds(ctx context.Context, request TikTokWebFetchGiftNamesByIdsRequest) (*TikTokWebFetchGiftNamesByIdsResponse, error) {
	return r.client.TikTokWebN50BatchGetGiftNamesByGiftIDs(ctx, request)
}

// TikTokWebFetchTikTokWebGuestCookieRequest is the request for GET /api/v1/tiktok/web/fetch_tiktok_web_guest_cookie.
type TikTokWebFetchTikTokWebGuestCookieRequest = TikTokWebGetTheGuestCookieRequest

// TikTokWebFetchTikTokWebGuestCookieResponse is the response for GET /api/v1/tiktok/web/fetch_tiktok_web_guest_cookie.
type TikTokWebFetchTikTokWebGuestCookieResponse = TikTokWebGetTheGuestCookieResponse

// FetchTikTokWebGuestCookie 获取游客 Cookie/Get the guest Cookie
//
// GET /api/v1/tiktok/web/fetch_tiktok_web_guest_cookie
func (r TikTokWebResource) FetchTikTokWebGuestCookie(ctx context.Context, request TikTokWebFetchTikTokWebGuestCookieRequest) (*TikTokWebFetchTikTokWebGuestCookieResponse, error) {
	return r.client.TikTokWebGetTheGuestCookie(ctx, request)
}

// TikTokWebDeviceRegisterResponse is the response for GET /api/v1/tiktok/web/device_register.
type TikTokWebDeviceRegisterResponse = TikTokWebRegisterDeviceForTikTokWebResponse

// DeviceRegister 设备注册/Register device for TikTok Web
//
// GET /api/v1/tiktok/web/device_register
func (r TikTokWebResource) DeviceRegister(ctx context.Context) (*TikTokWebDeviceRegisterResponse, error) {
	return r.client.TikTokWebRegisterDeviceForTikTokWeb(ctx)
}

// TikTokAppV3Resource contains endpoints from the TikTok-App-V3-API tag.
type TikTokAppV3Resource struct {
	client *Client
}

// TikTokAppV3FetchOneVideoRequest is the request for GET /api/v1/tiktok/app/v3/fetch_one_video.
type TikTokAppV3FetchOneVideoRequest = TikTokAppV3GetSingleVideoDataRequest

// TikTokAppV3FetchOneVideoResponse is the response for GET /api/v1/tiktok/app/v3/fetch_one_video.
type TikTokAppV3FetchOneVideoResponse = TikTokAppV3GetSingleVideoDataResponse

// FetchOneVideo 获取单个作品数据/Get single video data
//
// GET /api/v1/tiktok/app/v3/fetch_one_video
func (r TikTokAppV3Resource) FetchOneVideo(ctx context.Context, request TikTokAppV3FetchOneVideoRequest) (*TikTokAppV3FetchOneVideoResponse, error) {
	return r.client.TikTokAppV3GetSingleVideoData(ctx, request)
}

// TikTokAppV3FetchOneVideoV2Request is the request for GET /api/v1/tiktok/app/v3/fetch_one_video_v2.
type TikTokAppV3FetchOneVideoV2Request = TikTokAppV3GetSingleVideoDataV2Request

// TikTokAppV3FetchOneVideoV2Response is the response for GET /api/v1/tiktok/app/v3/fetch_one_video_v2.
type TikTokAppV3FetchOneVideoV2Response = TikTokAppV3GetSingleVideoDataV2Response

// FetchOneVideoV2 获取单个作品数据 V2/Get single video data V2
//
// GET /api/v1/tiktok/app/v3/fetch_one_video_v2
func (r TikTokAppV3Resource) FetchOneVideoV2(ctx context.Context, request TikTokAppV3FetchOneVideoV2Request) (*TikTokAppV3FetchOneVideoV2Response, error) {
	return r.client.TikTokAppV3GetSingleVideoDataV2(ctx, request)
}

// TikTokAppV3FetchOneVideoV3Request is the request for GET /api/v1/tiktok/app/v3/fetch_one_video_v3.
type TikTokAppV3FetchOneVideoV3Request = TikTokAppV3GetSingleVideoDataV3Request

// TikTokAppV3FetchOneVideoV3Response is the response for GET /api/v1/tiktok/app/v3/fetch_one_video_v3.
type TikTokAppV3FetchOneVideoV3Response = TikTokAppV3GetSingleVideoDataV3Response

// FetchOneVideoV3 获取单个作品数据 V3(支持国家参数)/Get single video data V3 (support country parameter)
//
// GET /api/v1/tiktok/app/v3/fetch_one_video_v3
func (r TikTokAppV3Resource) FetchOneVideoV3(ctx context.Context, request TikTokAppV3FetchOneVideoV3Request) (*TikTokAppV3FetchOneVideoV3Response, error) {
	return r.client.TikTokAppV3GetSingleVideoDataV3(ctx, request)
}

// TikTokAppV3FetchMultiVideoRequest is the request for POST /api/v1/tiktok/app/v3/fetch_multi_video.
type TikTokAppV3FetchMultiVideoRequest = TikTokAppV3BatchGetVideoInformationRequest

// TikTokAppV3FetchMultiVideoResponse is the response for POST /api/v1/tiktok/app/v3/fetch_multi_video.
type TikTokAppV3FetchMultiVideoResponse = TikTokAppV3BatchGetVideoInformationResponse

// FetchMultiVideo 批量获取视频信息/Batch Get Video Information
//
// POST /api/v1/tiktok/app/v3/fetch_multi_video
func (r TikTokAppV3Resource) FetchMultiVideo(ctx context.Context, request TikTokAppV3FetchMultiVideoRequest) (*TikTokAppV3FetchMultiVideoResponse, error) {
	return r.client.TikTokAppV3BatchGetVideoInformation(ctx, request)
}

// TikTokAppV3FetchMultiVideoV2Request is the request for POST /api/v1/tiktok/app/v3/fetch_multi_video_v2.
type TikTokAppV3FetchMultiVideoV2Request = TikTokAppV3BatchGetVideoInformationV2Request

// TikTokAppV3FetchMultiVideoV2Response is the response for POST /api/v1/tiktok/app/v3/fetch_multi_video_v2.
type TikTokAppV3FetchMultiVideoV2Response = TikTokAppV3BatchGetVideoInformationV2Response

// FetchMultiVideoV2 批量获取视频信息 V2/Batch Get Video Information V2
//
// POST /api/v1/tiktok/app/v3/fetch_multi_video_v2
func (r TikTokAppV3Resource) FetchMultiVideoV2(ctx context.Context, request TikTokAppV3FetchMultiVideoV2Request) (*TikTokAppV3FetchMultiVideoV2Response, error) {
	return r.client.TikTokAppV3BatchGetVideoInformationV2(ctx, request)
}

// TikTokAppV3FetchOneVideoByShareURLV2Request is the request for GET /api/v1/tiktok/app/v3/fetch_one_video_by_share_url_v2.
type TikTokAppV3FetchOneVideoByShareURLV2Request = TikTokAppV3GetSingleVideoDataBySharingLinkRequest

// TikTokAppV3FetchOneVideoByShareURLV2Response is the response for GET /api/v1/tiktok/app/v3/fetch_one_video_by_share_url_v2.
type TikTokAppV3FetchOneVideoByShareURLV2Response = TikTokAppV3GetSingleVideoDataBySharingLinkResponse

// FetchOneVideoByShareURLV2 根据分享链接获取单个作品数据/Get single video data by sharing link
//
// GET /api/v1/tiktok/app/v3/fetch_one_video_by_share_url_v2
func (r TikTokAppV3Resource) FetchOneVideoByShareURLV2(ctx context.Context, request TikTokAppV3FetchOneVideoByShareURLV2Request) (*TikTokAppV3FetchOneVideoByShareURLV2Response, error) {
	return r.client.TikTokAppV3GetSingleVideoDataBySharingLink(ctx, request)
}

// TikTokAppV3FetchOneVideoByShareURLRequest is the request for GET /api/v1/tiktok/app/v3/fetch_one_video_by_share_url.
type TikTokAppV3FetchOneVideoByShareURLRequest = TikTokAppV3GetSingleVideoDataBySharingLinkV3FetchOneVideoByShareURLRequest

// TikTokAppV3FetchOneVideoByShareURLResponse is the response for GET /api/v1/tiktok/app/v3/fetch_one_video_by_share_url.
type TikTokAppV3FetchOneVideoByShareURLResponse = TikTokAppV3GetSingleVideoDataBySharingLinkV3FetchOneVideoByShareURLResponse

// FetchOneVideoByShareURL 根据分享链接获取单个作品数据/Get single video data by sharing link
//
// GET /api/v1/tiktok/app/v3/fetch_one_video_by_share_url
func (r TikTokAppV3Resource) FetchOneVideoByShareURL(ctx context.Context, request TikTokAppV3FetchOneVideoByShareURLRequest) (*TikTokAppV3FetchOneVideoByShareURLResponse, error) {
	return r.client.TikTokAppV3GetSingleVideoDataBySharingLinkV3FetchOneVideoByShareURL(ctx, request)
}

// GetUserIDAndSecUserIDByUsername 使用用户名获取用户 user_id 和 sec_user_id/Get user_id and sec_user_id by Username
//
// GET /api/v1/tiktok/app/v3/get_user_id_and_sec_user_id_by_username
func (r TikTokAppV3Resource) GetUserIDAndSecUserIDByUsername(ctx context.Context, request TikTokAppV3GetUserIDAndSecUserIDByUsernameRequest) (*TikTokAppV3GetUserIDAndSecUserIDByUsernameResponse, error) {
	return r.client.TikTokAppV3GetUserIDAndSecUserIDByUsername(ctx, request)
}

// TikTokAppV3HandlerUserProfileRequest is the request for GET /api/v1/tiktok/app/v3/handler_user_profile.
type TikTokAppV3HandlerUserProfileRequest = TikTokAppV3GetInformationOfSpecifiedUserRequest

// TikTokAppV3HandlerUserProfileResponse is the response for GET /api/v1/tiktok/app/v3/handler_user_profile.
type TikTokAppV3HandlerUserProfileResponse = TikTokAppV3GetInformationOfSpecifiedUserResponse

// HandlerUserProfile 获取指定用户的信息/Get information of specified user
//
// GET /api/v1/tiktok/app/v3/handler_user_profile
func (r TikTokAppV3Resource) HandlerUserProfile(ctx context.Context, request TikTokAppV3HandlerUserProfileRequest) (*TikTokAppV3HandlerUserProfileResponse, error) {
	return r.client.TikTokAppV3GetInformationOfSpecifiedUser(ctx, request)
}

// TikTokAppV3FetchWebcastUserInfoRequest is the request for GET /api/v1/tiktok/app/v3/fetch_webcast_user_info.
type TikTokAppV3FetchWebcastUserInfoRequest = TikTokAppV3GetInformationOfSpecifiedWebcastUserRequest

// TikTokAppV3FetchWebcastUserInfoResponse is the response for GET /api/v1/tiktok/app/v3/fetch_webcast_user_info.
type TikTokAppV3FetchWebcastUserInfoResponse = TikTokAppV3GetInformationOfSpecifiedWebcastUserResponse

// FetchWebcastUserInfo 获取指定 Webcast 用户的信息/Get information of specified Webcast user
//
// GET /api/v1/tiktok/app/v3/fetch_webcast_user_info
func (r TikTokAppV3Resource) FetchWebcastUserInfo(ctx context.Context, request TikTokAppV3FetchWebcastUserInfoRequest) (*TikTokAppV3FetchWebcastUserInfoResponse, error) {
	return r.client.TikTokAppV3GetInformationOfSpecifiedWebcastUser(ctx, request)
}

// TikTokAppV3FetchUserCountryByUsernameRequest is the request for GET /api/v1/tiktok/app/v3/fetch_user_country_by_username.
type TikTokAppV3FetchUserCountryByUsernameRequest = TikTokAppV3GetUserAccountCountryByUsernameRequest

// TikTokAppV3FetchUserCountryByUsernameResponse is the response for GET /api/v1/tiktok/app/v3/fetch_user_country_by_username.
type TikTokAppV3FetchUserCountryByUsernameResponse = TikTokAppV3GetUserAccountCountryByUsernameResponse

// FetchUserCountryByUsername 通过用户名获取用户账号国家地区/Get user account country by username
//
// GET /api/v1/tiktok/app/v3/fetch_user_country_by_username
func (r TikTokAppV3Resource) FetchUserCountryByUsername(ctx context.Context, request TikTokAppV3FetchUserCountryByUsernameRequest) (*TikTokAppV3FetchUserCountryByUsernameResponse, error) {
	return r.client.TikTokAppV3GetUserAccountCountryByUsername(ctx, request)
}

// TikTokAppV3FetchSimilarUserRecommendationsRequest is the request for GET /api/v1/tiktok/app/v3/fetch_similar_user_recommendations.
type TikTokAppV3FetchSimilarUserRecommendationsRequest = TikTokAppV3SimilarUserRecommendationsRequest

// TikTokAppV3FetchSimilarUserRecommendationsResponse is the response for GET /api/v1/tiktok/app/v3/fetch_similar_user_recommendations.
type TikTokAppV3FetchSimilarUserRecommendationsResponse = TikTokAppV3SimilarUserRecommendationsResponse

// FetchSimilarUserRecommendations 获取类似用户推荐/Similar User Recommendations
//
// GET /api/v1/tiktok/app/v3/fetch_similar_user_recommendations
func (r TikTokAppV3Resource) FetchSimilarUserRecommendations(ctx context.Context, request TikTokAppV3FetchSimilarUserRecommendationsRequest) (*TikTokAppV3FetchSimilarUserRecommendationsResponse, error) {
	return r.client.TikTokAppV3SimilarUserRecommendations(ctx, request)
}

// TikTokAppV3FetchUserRepostVideosRequest is the request for GET /api/v1/tiktok/app/v3/fetch_user_repost_videos.
type TikTokAppV3FetchUserRepostVideosRequest = TikTokAppV3GetUserRepostVideoDataRequest

// TikTokAppV3FetchUserRepostVideosResponse is the response for GET /api/v1/tiktok/app/v3/fetch_user_repost_videos.
type TikTokAppV3FetchUserRepostVideosResponse = TikTokAppV3GetUserRepostVideoDataResponse

// FetchUserRepostVideos 获取用户转发的作品数据/Get user repost video data
//
// GET /api/v1/tiktok/app/v3/fetch_user_repost_videos
func (r TikTokAppV3Resource) FetchUserRepostVideos(ctx context.Context, request TikTokAppV3FetchUserRepostVideosRequest) (*TikTokAppV3FetchUserRepostVideosResponse, error) {
	return r.client.TikTokAppV3GetUserRepostVideoData(ctx, request)
}

// TikTokAppV3FetchUserPostVideosRequest is the request for GET /api/v1/tiktok/app/v3/fetch_user_post_videos.
type TikTokAppV3FetchUserPostVideosRequest = TikTokAppV3GetUserHomepageVideoDataV1Request

// TikTokAppV3FetchUserPostVideosResponse is the response for GET /api/v1/tiktok/app/v3/fetch_user_post_videos.
type TikTokAppV3FetchUserPostVideosResponse = TikTokAppV3GetUserHomepageVideoDataV1Response

// FetchUserPostVideos 获取用户主页作品数据 V1/Get user homepage video data V1
//
// GET /api/v1/tiktok/app/v3/fetch_user_post_videos
func (r TikTokAppV3Resource) FetchUserPostVideos(ctx context.Context, request TikTokAppV3FetchUserPostVideosRequest) (*TikTokAppV3FetchUserPostVideosResponse, error) {
	return r.client.TikTokAppV3GetUserHomepageVideoDataV1(ctx, request)
}

// TikTokAppV3FetchUserPostVideosV2Request is the request for GET /api/v1/tiktok/app/v3/fetch_user_post_videos_v2.
type TikTokAppV3FetchUserPostVideosV2Request = TikTokAppV3GetUserHomepageVideoDataV2Request

// TikTokAppV3FetchUserPostVideosV2Response is the response for GET /api/v1/tiktok/app/v3/fetch_user_post_videos_v2.
type TikTokAppV3FetchUserPostVideosV2Response = TikTokAppV3GetUserHomepageVideoDataV2Response

// FetchUserPostVideosV2 获取用户主页作品数据 V2/Get user homepage video data V2
//
// GET /api/v1/tiktok/app/v3/fetch_user_post_videos_v2
func (r TikTokAppV3Resource) FetchUserPostVideosV2(ctx context.Context, request TikTokAppV3FetchUserPostVideosV2Request) (*TikTokAppV3FetchUserPostVideosV2Response, error) {
	return r.client.TikTokAppV3GetUserHomepageVideoDataV2(ctx, request)
}

// TikTokAppV3FetchUserPostVideosV3Request is the request for GET /api/v1/tiktok/app/v3/fetch_user_post_videos_v3.
type TikTokAppV3FetchUserPostVideosV3Request = TikTokAppV3GetUserHomepageVideoDataV3Request

// TikTokAppV3FetchUserPostVideosV3Response is the response for GET /api/v1/tiktok/app/v3/fetch_user_post_videos_v3.
type TikTokAppV3FetchUserPostVideosV3Response = TikTokAppV3GetUserHomepageVideoDataV3Response

// FetchUserPostVideosV3 获取用户主页作品数据 V3（精简数据-更快速）/Get user homepage video data V3 (simplified data - faster)
//
// GET /api/v1/tiktok/app/v3/fetch_user_post_videos_v3
func (r TikTokAppV3Resource) FetchUserPostVideosV3(ctx context.Context, request TikTokAppV3FetchUserPostVideosV3Request) (*TikTokAppV3FetchUserPostVideosV3Response, error) {
	return r.client.TikTokAppV3GetUserHomepageVideoDataV3(ctx, request)
}

// TikTokAppV3FetchUserLikeVideosRequest is the request for GET /api/v1/tiktok/app/v3/fetch_user_like_videos.
type TikTokAppV3FetchUserLikeVideosRequest = TikTokAppV3GetUserLikeVideoDataRequest

// TikTokAppV3FetchUserLikeVideosResponse is the response for GET /api/v1/tiktok/app/v3/fetch_user_like_videos.
type TikTokAppV3FetchUserLikeVideosResponse = TikTokAppV3GetUserLikeVideoDataResponse

// FetchUserLikeVideos 获取用户喜欢作品数据/Get user like video data
//
// GET /api/v1/tiktok/app/v3/fetch_user_like_videos
func (r TikTokAppV3Resource) FetchUserLikeVideos(ctx context.Context, request TikTokAppV3FetchUserLikeVideosRequest) (*TikTokAppV3FetchUserLikeVideosResponse, error) {
	return r.client.TikTokAppV3GetUserLikeVideoData(ctx, request)
}

// TikTokAppV3FetchVideoCommentsRequest is the request for GET /api/v1/tiktok/app/v3/fetch_video_comments.
type TikTokAppV3FetchVideoCommentsRequest = TikTokAppV3GetSingleVideoCommentsDataRequest

// TikTokAppV3FetchVideoCommentsResponse is the response for GET /api/v1/tiktok/app/v3/fetch_video_comments.
type TikTokAppV3FetchVideoCommentsResponse = TikTokAppV3GetSingleVideoCommentsDataResponse

// FetchVideoComments 获取单个视频评论数据/Get single video comments data
//
// GET /api/v1/tiktok/app/v3/fetch_video_comments
func (r TikTokAppV3Resource) FetchVideoComments(ctx context.Context, request TikTokAppV3FetchVideoCommentsRequest) (*TikTokAppV3FetchVideoCommentsResponse, error) {
	return r.client.TikTokAppV3GetSingleVideoCommentsData(ctx, request)
}

// TikTokAppV3FetchVideoCommentRepliesRequest is the request for GET /api/v1/tiktok/app/v3/fetch_video_comment_replies.
type TikTokAppV3FetchVideoCommentRepliesRequest = TikTokAppV3GetCommentRepliesDataOfSpecifiedVideoRequest

// TikTokAppV3FetchVideoCommentRepliesResponse is the response for GET /api/v1/tiktok/app/v3/fetch_video_comment_replies.
type TikTokAppV3FetchVideoCommentRepliesResponse = TikTokAppV3GetCommentRepliesDataOfSpecifiedVideoResponse

// FetchVideoCommentReplies 获取指定视频的评论回复数据/Get comment replies data of specified video
//
// GET /api/v1/tiktok/app/v3/fetch_video_comment_replies
func (r TikTokAppV3Resource) FetchVideoCommentReplies(ctx context.Context, request TikTokAppV3FetchVideoCommentRepliesRequest) (*TikTokAppV3FetchVideoCommentRepliesResponse, error) {
	return r.client.TikTokAppV3GetCommentRepliesDataOfSpecifiedVideo(ctx, request)
}

// TikTokAppV3FetchGeneralSearchResultRequest is the request for GET /api/v1/tiktok/app/v3/fetch_general_search_result.
type TikTokAppV3FetchGeneralSearchResultRequest = TikTokAppV3GetComprehensiveSearchResultsOfSpecifiedKeywordsRequest

// TikTokAppV3FetchGeneralSearchResultResponse is the response for GET /api/v1/tiktok/app/v3/fetch_general_search_result.
type TikTokAppV3FetchGeneralSearchResultResponse = TikTokAppV3GetComprehensiveSearchResultsOfSpecifiedKeywordsResponse

// FetchGeneralSearchResult 获取指定关键词的综合搜索结果/Get comprehensive search results of specified keywords
//
// GET /api/v1/tiktok/app/v3/fetch_general_search_result
func (r TikTokAppV3Resource) FetchGeneralSearchResult(ctx context.Context, request TikTokAppV3FetchGeneralSearchResultRequest) (*TikTokAppV3FetchGeneralSearchResultResponse, error) {
	return r.client.TikTokAppV3GetComprehensiveSearchResultsOfSpecifiedKeywords(ctx, request)
}

// TikTokAppV3FetchVideoSearchResultRequest is the request for GET /api/v1/tiktok/app/v3/fetch_video_search_result.
type TikTokAppV3FetchVideoSearchResultRequest = TikTokAppV3GetVideoSearchResultsOfSpecifiedKeywordsRequest

// TikTokAppV3FetchVideoSearchResultResponse is the response for GET /api/v1/tiktok/app/v3/fetch_video_search_result.
type TikTokAppV3FetchVideoSearchResultResponse = TikTokAppV3GetVideoSearchResultsOfSpecifiedKeywordsResponse

// FetchVideoSearchResult 获取指定关键词的视频搜索结果/Get video search results of specified keywords
//
// GET /api/v1/tiktok/app/v3/fetch_video_search_result
func (r TikTokAppV3Resource) FetchVideoSearchResult(ctx context.Context, request TikTokAppV3FetchVideoSearchResultRequest) (*TikTokAppV3FetchVideoSearchResultResponse, error) {
	return r.client.TikTokAppV3GetVideoSearchResultsOfSpecifiedKeywords(ctx, request)
}

// TikTokAppV3FetchUserSearchResultRequest is the request for GET /api/v1/tiktok/app/v3/fetch_user_search_result.
type TikTokAppV3FetchUserSearchResultRequest = TikTokAppV3GetUserSearchResultsOfSpecifiedKeywordsRequest

// TikTokAppV3FetchUserSearchResultResponse is the response for GET /api/v1/tiktok/app/v3/fetch_user_search_result.
type TikTokAppV3FetchUserSearchResultResponse = TikTokAppV3GetUserSearchResultsOfSpecifiedKeywordsResponse

// FetchUserSearchResult 获取指定关键词的用户搜索结果/Get user search results of specified keywords
//
// GET /api/v1/tiktok/app/v3/fetch_user_search_result
func (r TikTokAppV3Resource) FetchUserSearchResult(ctx context.Context, request TikTokAppV3FetchUserSearchResultRequest) (*TikTokAppV3FetchUserSearchResultResponse, error) {
	return r.client.TikTokAppV3GetUserSearchResultsOfSpecifiedKeywords(ctx, request)
}

// TikTokAppV3FetchMusicSearchResultRequest is the request for GET /api/v1/tiktok/app/v3/fetch_music_search_result.
type TikTokAppV3FetchMusicSearchResultRequest = TikTokAppV3GetMusicSearchResultsOfSpecifiedKeywordsRequest

// TikTokAppV3FetchMusicSearchResultResponse is the response for GET /api/v1/tiktok/app/v3/fetch_music_search_result.
type TikTokAppV3FetchMusicSearchResultResponse = TikTokAppV3GetMusicSearchResultsOfSpecifiedKeywordsResponse

// FetchMusicSearchResult 获取指定关键词的音乐搜索结果/Get music search results of specified keywords
//
// GET /api/v1/tiktok/app/v3/fetch_music_search_result
func (r TikTokAppV3Resource) FetchMusicSearchResult(ctx context.Context, request TikTokAppV3FetchMusicSearchResultRequest) (*TikTokAppV3FetchMusicSearchResultResponse, error) {
	return r.client.TikTokAppV3GetMusicSearchResultsOfSpecifiedKeywords(ctx, request)
}

// TikTokAppV3FetchHashtagSearchResultRequest is the request for GET /api/v1/tiktok/app/v3/fetch_hashtag_search_result.
type TikTokAppV3FetchHashtagSearchResultRequest = TikTokAppV3GetHashtagSearchResultsOfSpecifiedKeywordsRequest

// TikTokAppV3FetchHashtagSearchResultResponse is the response for GET /api/v1/tiktok/app/v3/fetch_hashtag_search_result.
type TikTokAppV3FetchHashtagSearchResultResponse = TikTokAppV3GetHashtagSearchResultsOfSpecifiedKeywordsResponse

// FetchHashtagSearchResult 获取指定关键词的话题搜索结果/Get hashtag search results of specified keywords
//
// GET /api/v1/tiktok/app/v3/fetch_hashtag_search_result
func (r TikTokAppV3Resource) FetchHashtagSearchResult(ctx context.Context, request TikTokAppV3FetchHashtagSearchResultRequest) (*TikTokAppV3FetchHashtagSearchResultResponse, error) {
	return r.client.TikTokAppV3GetHashtagSearchResultsOfSpecifiedKeywords(ctx, request)
}

// TikTokAppV3FetchLiveSearchResultRequest is the request for GET /api/v1/tiktok/app/v3/fetch_live_search_result.
type TikTokAppV3FetchLiveSearchResultRequest = TikTokAppV3GetLiveSearchResultsOfSpecifiedKeywordsRequest

// TikTokAppV3FetchLiveSearchResultResponse is the response for GET /api/v1/tiktok/app/v3/fetch_live_search_result.
type TikTokAppV3FetchLiveSearchResultResponse = TikTokAppV3GetLiveSearchResultsOfSpecifiedKeywordsResponse

// FetchLiveSearchResult 获取指定关键词的直播搜索结果/Get live search results of specified keywords
//
// GET /api/v1/tiktok/app/v3/fetch_live_search_result
func (r TikTokAppV3Resource) FetchLiveSearchResult(ctx context.Context, request TikTokAppV3FetchLiveSearchResultRequest) (*TikTokAppV3FetchLiveSearchResultResponse, error) {
	return r.client.TikTokAppV3GetLiveSearchResultsOfSpecifiedKeywords(ctx, request)
}

// TikTokAppV3FetchLocationSearchRequest is the request for GET /api/v1/tiktok/app/v3/fetch_location_search.
type TikTokAppV3FetchLocationSearchRequest = TikTokAppV3GetLocationSearchResultsRequest

// TikTokAppV3FetchLocationSearchResponse is the response for GET /api/v1/tiktok/app/v3/fetch_location_search.
type TikTokAppV3FetchLocationSearchResponse = TikTokAppV3GetLocationSearchResultsResponse

// FetchLocationSearch 获取地点搜索结果/Get location search results
//
// GET /api/v1/tiktok/app/v3/fetch_location_search
func (r TikTokAppV3Resource) FetchLocationSearch(ctx context.Context, request TikTokAppV3FetchLocationSearchRequest) (*TikTokAppV3FetchLocationSearchResponse, error) {
	return r.client.TikTokAppV3GetLocationSearchResults(ctx, request)
}

// TikTokAppV3FetchMusicDetailRequest is the request for GET /api/v1/tiktok/app/v3/fetch_music_detail.
type TikTokAppV3FetchMusicDetailRequest = TikTokAppV3GetDetailsOfSpecifiedMusicRequest

// TikTokAppV3FetchMusicDetailResponse is the response for GET /api/v1/tiktok/app/v3/fetch_music_detail.
type TikTokAppV3FetchMusicDetailResponse = TikTokAppV3GetDetailsOfSpecifiedMusicResponse

// FetchMusicDetail 获取指定音乐的详情数据/Get details of specified music
//
// GET /api/v1/tiktok/app/v3/fetch_music_detail
func (r TikTokAppV3Resource) FetchMusicDetail(ctx context.Context, request TikTokAppV3FetchMusicDetailRequest) (*TikTokAppV3FetchMusicDetailResponse, error) {
	return r.client.TikTokAppV3GetDetailsOfSpecifiedMusic(ctx, request)
}

// TikTokAppV3FetchMusicVideoListRequest is the request for GET /api/v1/tiktok/app/v3/fetch_music_video_list.
type TikTokAppV3FetchMusicVideoListRequest = TikTokAppV3GetVideoListOfSpecifiedMusicRequest

// TikTokAppV3FetchMusicVideoListResponse is the response for GET /api/v1/tiktok/app/v3/fetch_music_video_list.
type TikTokAppV3FetchMusicVideoListResponse = TikTokAppV3GetVideoListOfSpecifiedMusicResponse

// FetchMusicVideoList 获取指定音乐的视频列表数据/Get video list of specified music
//
// GET /api/v1/tiktok/app/v3/fetch_music_video_list
func (r TikTokAppV3Resource) FetchMusicVideoList(ctx context.Context, request TikTokAppV3FetchMusicVideoListRequest) (*TikTokAppV3FetchMusicVideoListResponse, error) {
	return r.client.TikTokAppV3GetVideoListOfSpecifiedMusic(ctx, request)
}

// TikTokAppV3FetchHashtagDetailRequest is the request for GET /api/v1/tiktok/app/v3/fetch_hashtag_detail.
type TikTokAppV3FetchHashtagDetailRequest = TikTokAppV3GetDetailsOfSpecifiedHashtagRequest

// TikTokAppV3FetchHashtagDetailResponse is the response for GET /api/v1/tiktok/app/v3/fetch_hashtag_detail.
type TikTokAppV3FetchHashtagDetailResponse = TikTokAppV3GetDetailsOfSpecifiedHashtagResponse

// FetchHashtagDetail 获取指定话题的详情数据/Get details of specified hashtag
//
// GET /api/v1/tiktok/app/v3/fetch_hashtag_detail
func (r TikTokAppV3Resource) FetchHashtagDetail(ctx context.Context, request TikTokAppV3FetchHashtagDetailRequest) (*TikTokAppV3FetchHashtagDetailResponse, error) {
	return r.client.TikTokAppV3GetDetailsOfSpecifiedHashtag(ctx, request)
}

// TikTokAppV3FetchHashtagVideoListRequest is the request for GET /api/v1/tiktok/app/v3/fetch_hashtag_video_list.
type TikTokAppV3FetchHashtagVideoListRequest = TikTokAppV3GetVideoListOfSpecifiedHashtagRequest

// TikTokAppV3FetchHashtagVideoListResponse is the response for GET /api/v1/tiktok/app/v3/fetch_hashtag_video_list.
type TikTokAppV3FetchHashtagVideoListResponse = TikTokAppV3GetVideoListOfSpecifiedHashtagResponse

// FetchHashtagVideoList 获取指定话题的作品数据/Get video list of specified hashtag
//
// GET /api/v1/tiktok/app/v3/fetch_hashtag_video_list
func (r TikTokAppV3Resource) FetchHashtagVideoList(ctx context.Context, request TikTokAppV3FetchHashtagVideoListRequest) (*TikTokAppV3FetchHashtagVideoListResponse, error) {
	return r.client.TikTokAppV3GetVideoListOfSpecifiedHashtag(ctx, request)
}

// TikTokAppV3FetchUserFollowerListRequest is the request for GET /api/v1/tiktok/app/v3/fetch_user_follower_list.
type TikTokAppV3FetchUserFollowerListRequest = TikTokAppV3GetFollowerListOfSpecifiedUserRequest

// TikTokAppV3FetchUserFollowerListResponse is the response for GET /api/v1/tiktok/app/v3/fetch_user_follower_list.
type TikTokAppV3FetchUserFollowerListResponse = TikTokAppV3GetFollowerListOfSpecifiedUserResponse

// FetchUserFollowerList 获取指定用户的粉丝列表数据/Get follower list of specified user
//
// GET /api/v1/tiktok/app/v3/fetch_user_follower_list
func (r TikTokAppV3Resource) FetchUserFollowerList(ctx context.Context, request TikTokAppV3FetchUserFollowerListRequest) (*TikTokAppV3FetchUserFollowerListResponse, error) {
	return r.client.TikTokAppV3GetFollowerListOfSpecifiedUser(ctx, request)
}

// TikTokAppV3FetchUserFollowingListRequest is the request for GET /api/v1/tiktok/app/v3/fetch_user_following_list.
type TikTokAppV3FetchUserFollowingListRequest = TikTokAppV3GetFollowingListOfSpecifiedUserRequest

// TikTokAppV3FetchUserFollowingListResponse is the response for GET /api/v1/tiktok/app/v3/fetch_user_following_list.
type TikTokAppV3FetchUserFollowingListResponse = TikTokAppV3GetFollowingListOfSpecifiedUserResponse

// FetchUserFollowingList 获取指定用户的关注列表数据/Get following list of specified user
//
// GET /api/v1/tiktok/app/v3/fetch_user_following_list
func (r TikTokAppV3Resource) FetchUserFollowingList(ctx context.Context, request TikTokAppV3FetchUserFollowingListRequest) (*TikTokAppV3FetchUserFollowingListResponse, error) {
	return r.client.TikTokAppV3GetFollowingListOfSpecifiedUser(ctx, request)
}

// TikTokAppV3FetchCreatorSearchInsightsRequest is the request for GET /api/v1/tiktok/app/v3/fetch_creator_search_insights.
type TikTokAppV3FetchCreatorSearchInsightsRequest = TikTokAppV3CreatorSearchInsightsRequest

// TikTokAppV3FetchCreatorSearchInsightsResponse is the response for GET /api/v1/tiktok/app/v3/fetch_creator_search_insights.
type TikTokAppV3FetchCreatorSearchInsightsResponse = TikTokAppV3CreatorSearchInsightsResponse

// FetchCreatorSearchInsights 创作者搜索洞察/Creator Search Insights
//
// GET /api/v1/tiktok/app/v3/fetch_creator_search_insights
func (r TikTokAppV3Resource) FetchCreatorSearchInsights(ctx context.Context, request TikTokAppV3FetchCreatorSearchInsightsRequest) (*TikTokAppV3FetchCreatorSearchInsightsResponse, error) {
	return r.client.TikTokAppV3CreatorSearchInsights(ctx, request)
}

// TikTokAppV3FetchCreatorSearchInsightsDetailRequest is the request for GET /api/v1/tiktok/app/v3/fetch_creator_search_insights_detail.
type TikTokAppV3FetchCreatorSearchInsightsDetailRequest = TikTokAppV3CreatorSearchInsightsDetailRequest

// TikTokAppV3FetchCreatorSearchInsightsDetailResponse is the response for GET /api/v1/tiktok/app/v3/fetch_creator_search_insights_detail.
type TikTokAppV3FetchCreatorSearchInsightsDetailResponse = TikTokAppV3CreatorSearchInsightsDetailResponse

// FetchCreatorSearchInsightsDetail 创作者搜索洞察详情/Creator Search Insights Detail
//
// GET /api/v1/tiktok/app/v3/fetch_creator_search_insights_detail
func (r TikTokAppV3Resource) FetchCreatorSearchInsightsDetail(ctx context.Context, request TikTokAppV3FetchCreatorSearchInsightsDetailRequest) (*TikTokAppV3FetchCreatorSearchInsightsDetailResponse, error) {
	return r.client.TikTokAppV3CreatorSearchInsightsDetail(ctx, request)
}

// TikTokAppV3FetchCreatorSearchInsightsTrendRequest is the request for GET /api/v1/tiktok/app/v3/fetch_creator_search_insights_trend.
type TikTokAppV3FetchCreatorSearchInsightsTrendRequest = TikTokAppV3CreatorSearchInsightsTrendRequest

// TikTokAppV3FetchCreatorSearchInsightsTrendResponse is the response for GET /api/v1/tiktok/app/v3/fetch_creator_search_insights_trend.
type TikTokAppV3FetchCreatorSearchInsightsTrendResponse = TikTokAppV3CreatorSearchInsightsTrendResponse

// FetchCreatorSearchInsightsTrend 创作者搜索洞察趋势/Creator Search Insights Trend
//
// GET /api/v1/tiktok/app/v3/fetch_creator_search_insights_trend
func (r TikTokAppV3Resource) FetchCreatorSearchInsightsTrend(ctx context.Context, request TikTokAppV3FetchCreatorSearchInsightsTrendRequest) (*TikTokAppV3FetchCreatorSearchInsightsTrendResponse, error) {
	return r.client.TikTokAppV3CreatorSearchInsightsTrend(ctx, request)
}

// TikTokAppV3FetchCreatorSearchInsightsVideosRequest is the request for GET /api/v1/tiktok/app/v3/fetch_creator_search_insights_videos.
type TikTokAppV3FetchCreatorSearchInsightsVideosRequest = TikTokAppV3CreatorSearchInsightsVideosRequest

// TikTokAppV3FetchCreatorSearchInsightsVideosResponse is the response for GET /api/v1/tiktok/app/v3/fetch_creator_search_insights_videos.
type TikTokAppV3FetchCreatorSearchInsightsVideosResponse = TikTokAppV3CreatorSearchInsightsVideosResponse

// FetchCreatorSearchInsightsVideos 创作者搜索洞察相关视频/Creator Search Insights Videos
//
// GET /api/v1/tiktok/app/v3/fetch_creator_search_insights_videos
func (r TikTokAppV3Resource) FetchCreatorSearchInsightsVideos(ctx context.Context, request TikTokAppV3FetchCreatorSearchInsightsVideosRequest) (*TikTokAppV3FetchCreatorSearchInsightsVideosResponse, error) {
	return r.client.TikTokAppV3CreatorSearchInsightsVideos(ctx, request)
}

// TikTokAppV3FetchMusicChartListRequest is the request for GET /api/v1/tiktok/app/v3/fetch_music_chart_list.
type TikTokAppV3FetchMusicChartListRequest = TikTokAppV3MusicChartListRequest

// TikTokAppV3FetchMusicChartListResponse is the response for GET /api/v1/tiktok/app/v3/fetch_music_chart_list.
type TikTokAppV3FetchMusicChartListResponse = TikTokAppV3MusicChartListResponse

// FetchMusicChartList 音乐排行榜/Music Chart List
//
// GET /api/v1/tiktok/app/v3/fetch_music_chart_list
func (r TikTokAppV3Resource) FetchMusicChartList(ctx context.Context, request TikTokAppV3FetchMusicChartListRequest) (*TikTokAppV3FetchMusicChartListResponse, error) {
	return r.client.TikTokAppV3MusicChartList(ctx, request)
}

// SearchFollowerList 搜索粉丝列表/Search follower list
//
// GET /api/v1/tiktok/app/v3/search_follower_list
func (r TikTokAppV3Resource) SearchFollowerList(ctx context.Context, request TikTokAppV3SearchFollowerListRequest) (*TikTokAppV3SearchFollowerListResponse, error) {
	return r.client.TikTokAppV3SearchFollowerList(ctx, request)
}

// SearchFollowingList 搜索关注列表/Search following list
//
// GET /api/v1/tiktok/app/v3/search_following_list
func (r TikTokAppV3Resource) SearchFollowingList(ctx context.Context, request TikTokAppV3SearchFollowingListRequest) (*TikTokAppV3SearchFollowingListResponse, error) {
	return r.client.TikTokAppV3SearchFollowingList(ctx, request)
}

// TikTokAppV3FetchLiveRoomInfoRequest is the request for GET /api/v1/tiktok/app/v3/fetch_live_room_info.
type TikTokAppV3FetchLiveRoomInfoRequest = TikTokAppV3GetDataOfSpecifiedLiveRoomRequest

// TikTokAppV3FetchLiveRoomInfoResponse is the response for GET /api/v1/tiktok/app/v3/fetch_live_room_info.
type TikTokAppV3FetchLiveRoomInfoResponse = TikTokAppV3GetDataOfSpecifiedLiveRoomResponse

// FetchLiveRoomInfo 获取指定直播间的数据/Get data of specified live room
//
// GET /api/v1/tiktok/app/v3/fetch_live_room_info
func (r TikTokAppV3Resource) FetchLiveRoomInfo(ctx context.Context, request TikTokAppV3FetchLiveRoomInfoRequest) (*TikTokAppV3FetchLiveRoomInfoResponse, error) {
	return r.client.TikTokAppV3GetDataOfSpecifiedLiveRoom(ctx, request)
}

// TikTokAppV3FetchLiveRankingListRequest is the request for GET /api/v1/tiktok/app/v3/fetch_live_ranking_list.
type TikTokAppV3FetchLiveRankingListRequest = TikTokAppV3GetLiveRoomRankingListRequest

// TikTokAppV3FetchLiveRankingListResponse is the response for GET /api/v1/tiktok/app/v3/fetch_live_ranking_list.
type TikTokAppV3FetchLiveRankingListResponse = TikTokAppV3GetLiveRoomRankingListResponse

// FetchLiveRankingList 获取直播间排行榜数据/Get live room ranking list
//
// GET /api/v1/tiktok/app/v3/fetch_live_ranking_list
func (r TikTokAppV3Resource) FetchLiveRankingList(ctx context.Context, request TikTokAppV3FetchLiveRankingListRequest) (*TikTokAppV3FetchLiveRankingListResponse, error) {
	return r.client.TikTokAppV3GetLiveRoomRankingList(ctx, request)
}

// TikTokAppV3CheckLiveRoomOnlineRequest is the request for GET /api/v1/tiktok/app/v3/check_live_room_online.
type TikTokAppV3CheckLiveRoomOnlineRequest = TikTokAppV3CheckIfLiveRoomIsOnlineRequest

// TikTokAppV3CheckLiveRoomOnlineResponse is the response for GET /api/v1/tiktok/app/v3/check_live_room_online.
type TikTokAppV3CheckLiveRoomOnlineResponse = TikTokAppV3CheckIfLiveRoomIsOnlineResponse

// CheckLiveRoomOnline 检测直播间是否在线/Check if live room is online
//
// GET /api/v1/tiktok/app/v3/check_live_room_online
func (r TikTokAppV3Resource) CheckLiveRoomOnline(ctx context.Context, request TikTokAppV3CheckLiveRoomOnlineRequest) (*TikTokAppV3CheckLiveRoomOnlineResponse, error) {
	return r.client.TikTokAppV3CheckIfLiveRoomIsOnline(ctx, request)
}

// TikTokAppV3CheckLiveRoomOnlineBatchRequest is the request for POST /api/v1/tiktok/app/v3/check_live_room_online_batch.
type TikTokAppV3CheckLiveRoomOnlineBatchRequest = TikTokAppV3BatchCheckIfLiveRoomsAreOnlineRequest

// TikTokAppV3CheckLiveRoomOnlineBatchResponse is the response for POST /api/v1/tiktok/app/v3/check_live_room_online_batch.
type TikTokAppV3CheckLiveRoomOnlineBatchResponse = TikTokAppV3BatchCheckIfLiveRoomsAreOnlineResponse

// CheckLiveRoomOnlineBatch 批量检测直播间是否在线/Batch check if live rooms are online
//
// POST /api/v1/tiktok/app/v3/check_live_room_online_batch
func (r TikTokAppV3Resource) CheckLiveRoomOnlineBatch(ctx context.Context, request TikTokAppV3CheckLiveRoomOnlineBatchRequest) (*TikTokAppV3CheckLiveRoomOnlineBatchResponse, error) {
	return r.client.TikTokAppV3BatchCheckIfLiveRoomsAreOnline(ctx, request)
}

// TikTokAppV3FetchShareShortLinkRequest is the request for GET /api/v1/tiktok/app/v3/fetch_share_short_link.
type TikTokAppV3FetchShareShortLinkRequest = TikTokAppV3GetShareShortLinkRequest

// TikTokAppV3FetchShareShortLinkResponse is the response for GET /api/v1/tiktok/app/v3/fetch_share_short_link.
type TikTokAppV3FetchShareShortLinkResponse = TikTokAppV3GetShareShortLinkResponse

// FetchShareShortLink 获取分享短链接/Get share short link
//
// GET /api/v1/tiktok/app/v3/fetch_share_short_link
func (r TikTokAppV3Resource) FetchShareShortLink(ctx context.Context, request TikTokAppV3FetchShareShortLinkRequest) (*TikTokAppV3FetchShareShortLinkResponse, error) {
	return r.client.TikTokAppV3GetShareShortLink(ctx, request)
}

// TikTokAppV3FetchShareQrCodeRequest is the request for GET /api/v1/tiktok/app/v3/fetch_share_qr_code.
type TikTokAppV3FetchShareQrCodeRequest = TikTokAppV3GetShareQRCodeRequest

// TikTokAppV3FetchShareQrCodeResponse is the response for GET /api/v1/tiktok/app/v3/fetch_share_qr_code.
type TikTokAppV3FetchShareQrCodeResponse = TikTokAppV3GetShareQRCodeResponse

// FetchShareQrCode 获取分享二维码/Get share QR code
//
// GET /api/v1/tiktok/app/v3/fetch_share_qr_code
func (r TikTokAppV3Resource) FetchShareQrCode(ctx context.Context, request TikTokAppV3FetchShareQrCodeRequest) (*TikTokAppV3FetchShareQrCodeResponse, error) {
	return r.client.TikTokAppV3GetShareQRCode(ctx, request)
}

// TikTokAppV3FetchProductSearchRequest is the request for GET /api/v1/tiktok/app/v3/fetch_product_search.
type TikTokAppV3FetchProductSearchRequest = TikTokAppV3GetProductSearchResultsRequest

// TikTokAppV3FetchProductSearchResponse is the response for GET /api/v1/tiktok/app/v3/fetch_product_search.
type TikTokAppV3FetchProductSearchResponse = TikTokAppV3GetProductSearchResultsResponse

// FetchProductSearch 获取商品搜索结果/Get product search results
//
// GET /api/v1/tiktok/app/v3/fetch_product_search
func (r TikTokAppV3Resource) FetchProductSearch(ctx context.Context, request TikTokAppV3FetchProductSearchRequest) (*TikTokAppV3FetchProductSearchResponse, error) {
	return r.client.TikTokAppV3GetProductSearchResults(ctx, request)
}

// TikTokAppV3FetchCreatorInfoRequest is the request for GET /api/v1/tiktok/app/v3/fetch_creator_info.
type TikTokAppV3FetchCreatorInfoRequest = TikTokAppV3GetShoppingCreatorInformationRequest

// TikTokAppV3FetchCreatorInfoResponse is the response for GET /api/v1/tiktok/app/v3/fetch_creator_info.
type TikTokAppV3FetchCreatorInfoResponse = TikTokAppV3GetShoppingCreatorInformationResponse

// FetchCreatorInfo 获取带货创作者信息/Get shopping creator information
//
// GET /api/v1/tiktok/app/v3/fetch_creator_info
func (r TikTokAppV3Resource) FetchCreatorInfo(ctx context.Context, request TikTokAppV3FetchCreatorInfoRequest) (*TikTokAppV3FetchCreatorInfoResponse, error) {
	return r.client.TikTokAppV3GetShoppingCreatorInformation(ctx, request)
}

// TikTokAppV3FetchCreatorShowcaseProductListRequest is the request for GET /api/v1/tiktok/app/v3/fetch_creator_showcase_product_list.
type TikTokAppV3FetchCreatorShowcaseProductListRequest = TikTokAppV3GetCreatorShowcaseProductListRequest

// TikTokAppV3FetchCreatorShowcaseProductListResponse is the response for GET /api/v1/tiktok/app/v3/fetch_creator_showcase_product_list.
type TikTokAppV3FetchCreatorShowcaseProductListResponse = TikTokAppV3GetCreatorShowcaseProductListResponse

// FetchCreatorShowcaseProductList 获取创作者橱窗商品列表/Get creator showcase product list
//
// GET /api/v1/tiktok/app/v3/fetch_creator_showcase_product_list
func (r TikTokAppV3Resource) FetchCreatorShowcaseProductList(ctx context.Context, request TikTokAppV3FetchCreatorShowcaseProductListRequest) (*TikTokAppV3FetchCreatorShowcaseProductListResponse, error) {
	return r.client.TikTokAppV3GetCreatorShowcaseProductList(ctx, request)
}

// TikTokAppV3FetchShopIDByShareLinkRequest is the request for GET /api/v1/tiktok/app/v3/fetch_shop_id_by_share_link.
type TikTokAppV3FetchShopIDByShareLinkRequest = TikTokAppV3GetShopIDByShareLinkRequest

// TikTokAppV3FetchShopIDByShareLinkResponse is the response for GET /api/v1/tiktok/app/v3/fetch_shop_id_by_share_link.
type TikTokAppV3FetchShopIDByShareLinkResponse = TikTokAppV3GetShopIDByShareLinkResponse

// FetchShopIDByShareLink 通过分享链接获取店铺ID/Get Shop ID by Share Link
//
// GET /api/v1/tiktok/app/v3/fetch_shop_id_by_share_link
func (r TikTokAppV3Resource) FetchShopIDByShareLink(ctx context.Context, request TikTokAppV3FetchShopIDByShareLinkRequest) (*TikTokAppV3FetchShopIDByShareLinkResponse, error) {
	return r.client.TikTokAppV3GetShopIDByShareLink(ctx, request)
}

// TikTokAppV3FetchProductIDByShareLinkRequest is the request for GET /api/v1/tiktok/app/v3/fetch_product_id_by_share_link.
type TikTokAppV3FetchProductIDByShareLinkRequest = TikTokAppV3GetProductIDByShareLinkRequest

// TikTokAppV3FetchProductIDByShareLinkResponse is the response for GET /api/v1/tiktok/app/v3/fetch_product_id_by_share_link.
type TikTokAppV3FetchProductIDByShareLinkResponse = TikTokAppV3GetProductIDByShareLinkResponse

// FetchProductIDByShareLink 通过分享链接获取商品ID/Get Product ID by Share Link
//
// GET /api/v1/tiktok/app/v3/fetch_product_id_by_share_link
func (r TikTokAppV3Resource) FetchProductIDByShareLink(ctx context.Context, request TikTokAppV3FetchProductIDByShareLinkRequest) (*TikTokAppV3FetchProductIDByShareLinkResponse, error) {
	return r.client.TikTokAppV3GetProductIDByShareLink(ctx, request)
}

// TikTokAppV3FetchProductDetailRequest is the request for GET /api/v1/tiktok/app/v3/fetch_product_detail.
type TikTokAppV3FetchProductDetailRequest = TikTokAppV3GetProductDetailDataRequest

// TikTokAppV3FetchProductDetailResponse is the response for GET /api/v1/tiktok/app/v3/fetch_product_detail.
type TikTokAppV3FetchProductDetailResponse = TikTokAppV3GetProductDetailDataResponse

// FetchProductDetail 获取商品详情数据（即将弃用，使用 fetch_product_detail_v2 代替）/Get product detail data (will be deprecated, use fetch_product_detail_v2 instead)
//
// GET /api/v1/tiktok/app/v3/fetch_product_detail
func (r TikTokAppV3Resource) FetchProductDetail(ctx context.Context, request TikTokAppV3FetchProductDetailRequest) (*TikTokAppV3FetchProductDetailResponse, error) {
	return r.client.TikTokAppV3GetProductDetailData(ctx, request)
}

// TikTokAppV3FetchProductDetailV2Request is the request for GET /api/v1/tiktok/app/v3/fetch_product_detail_v2.
type TikTokAppV3FetchProductDetailV2Request = TikTokAppV3GetProductDetailDataV2Request

// TikTokAppV3FetchProductDetailV2Response is the response for GET /api/v1/tiktok/app/v3/fetch_product_detail_v2.
type TikTokAppV3FetchProductDetailV2Response = TikTokAppV3GetProductDetailDataV2Response

// FetchProductDetailV2 获取商品详情数据V2/Get product detail data V2
//
// GET /api/v1/tiktok/app/v3/fetch_product_detail_v2
func (r TikTokAppV3Resource) FetchProductDetailV2(ctx context.Context, request TikTokAppV3FetchProductDetailV2Request) (*TikTokAppV3FetchProductDetailV2Response, error) {
	return r.client.TikTokAppV3GetProductDetailDataV2(ctx, request)
}

// TikTokAppV3FetchProductDetailV3Request is the request for GET /api/v1/tiktok/app/v3/fetch_product_detail_v3.
type TikTokAppV3FetchProductDetailV3Request = TikTokAppV3GetProductDetailDataV3Request

// TikTokAppV3FetchProductDetailV3Response is the response for GET /api/v1/tiktok/app/v3/fetch_product_detail_v3.
type TikTokAppV3FetchProductDetailV3Response = TikTokAppV3GetProductDetailDataV3Response

// FetchProductDetailV3 获取商品详情数据V3 / Get product detail data V3
//
// GET /api/v1/tiktok/app/v3/fetch_product_detail_v3
func (r TikTokAppV3Resource) FetchProductDetailV3(ctx context.Context, request TikTokAppV3FetchProductDetailV3Request) (*TikTokAppV3FetchProductDetailV3Response, error) {
	return r.client.TikTokAppV3GetProductDetailDataV3(ctx, request)
}

// TikTokAppV3FetchProductDetailV4Request is the request for GET /api/v1/tiktok/app/v3/fetch_product_detail_v4.
type TikTokAppV3FetchProductDetailV4Request = TikTokAppV3GetProductDetailDataV4Request

// TikTokAppV3FetchProductDetailV4Response is the response for GET /api/v1/tiktok/app/v3/fetch_product_detail_v4.
type TikTokAppV3FetchProductDetailV4Response = TikTokAppV3GetProductDetailDataV4Response

// FetchProductDetailV4 获取商品详情数据V4 / Get product detail data V4
//
// GET /api/v1/tiktok/app/v3/fetch_product_detail_v4
func (r TikTokAppV3Resource) FetchProductDetailV4(ctx context.Context, request TikTokAppV3FetchProductDetailV4Request) (*TikTokAppV3FetchProductDetailV4Response, error) {
	return r.client.TikTokAppV3GetProductDetailDataV4(ctx, request)
}

// TikTokAppV3FetchProductReviewRequest is the request for GET /api/v1/tiktok/app/v3/fetch_product_review.
type TikTokAppV3FetchProductReviewRequest = TikTokAppV3GetProductReviewDataRequest

// TikTokAppV3FetchProductReviewResponse is the response for GET /api/v1/tiktok/app/v3/fetch_product_review.
type TikTokAppV3FetchProductReviewResponse = TikTokAppV3GetProductReviewDataResponse

// FetchProductReview 获取商品评价数据/Get product review data
//
// GET /api/v1/tiktok/app/v3/fetch_product_review
func (r TikTokAppV3Resource) FetchProductReview(ctx context.Context, request TikTokAppV3FetchProductReviewRequest) (*TikTokAppV3FetchProductReviewResponse, error) {
	return r.client.TikTokAppV3GetProductReviewData(ctx, request)
}

// TikTokAppV3FetchShopHomePageListRequest is the request for GET /api/v1/tiktok/app/v3/fetch_shop_home_page_list.
type TikTokAppV3FetchShopHomePageListRequest = TikTokAppV3GetShopHomePageListDataRequest

// TikTokAppV3FetchShopHomePageListResponse is the response for GET /api/v1/tiktok/app/v3/fetch_shop_home_page_list.
type TikTokAppV3FetchShopHomePageListResponse = TikTokAppV3GetShopHomePageListDataResponse

// FetchShopHomePageList 获取商家主页Page列表数据/Get shop home page list data
//
// GET /api/v1/tiktok/app/v3/fetch_shop_home_page_list
func (r TikTokAppV3Resource) FetchShopHomePageList(ctx context.Context, request TikTokAppV3FetchShopHomePageListRequest) (*TikTokAppV3FetchShopHomePageListResponse, error) {
	return r.client.TikTokAppV3GetShopHomePageListData(ctx, request)
}

// TikTokAppV3FetchShopHomeRequest is the request for GET /api/v1/tiktok/app/v3/fetch_shop_home.
type TikTokAppV3FetchShopHomeRequest = TikTokAppV3GetShopHomePageDataRequest

// TikTokAppV3FetchShopHomeResponse is the response for GET /api/v1/tiktok/app/v3/fetch_shop_home.
type TikTokAppV3FetchShopHomeResponse = TikTokAppV3GetShopHomePageDataResponse

// FetchShopHome 获取商家主页数据/Get shop home page data
//
// GET /api/v1/tiktok/app/v3/fetch_shop_home
func (r TikTokAppV3Resource) FetchShopHome(ctx context.Context, request TikTokAppV3FetchShopHomeRequest) (*TikTokAppV3FetchShopHomeResponse, error) {
	return r.client.TikTokAppV3GetShopHomePageData(ctx, request)
}

// TikTokAppV3FetchShopProductRecommendRequest is the request for GET /api/v1/tiktok/app/v3/fetch_shop_product_recommend.
type TikTokAppV3FetchShopProductRecommendRequest = TikTokAppV3GetShopProductRecommendDataRequest

// TikTokAppV3FetchShopProductRecommendResponse is the response for GET /api/v1/tiktok/app/v3/fetch_shop_product_recommend.
type TikTokAppV3FetchShopProductRecommendResponse = TikTokAppV3GetShopProductRecommendDataResponse

// FetchShopProductRecommend 获取商家商品推荐数据/Get shop product recommend data
//
// GET /api/v1/tiktok/app/v3/fetch_shop_product_recommend
func (r TikTokAppV3Resource) FetchShopProductRecommend(ctx context.Context, request TikTokAppV3FetchShopProductRecommendRequest) (*TikTokAppV3FetchShopProductRecommendResponse, error) {
	return r.client.TikTokAppV3GetShopProductRecommendData(ctx, request)
}

// TikTokAppV3FetchShopProductListRequest is the request for GET /api/v1/tiktok/app/v3/fetch_shop_product_list.
type TikTokAppV3FetchShopProductListRequest = TikTokAppV3GetShopProductListDataRequest

// TikTokAppV3FetchShopProductListResponse is the response for GET /api/v1/tiktok/app/v3/fetch_shop_product_list.
type TikTokAppV3FetchShopProductListResponse = TikTokAppV3GetShopProductListDataResponse

// FetchShopProductList 获取商家商品列表数据/Get shop product list data
//
// GET /api/v1/tiktok/app/v3/fetch_shop_product_list
func (r TikTokAppV3Resource) FetchShopProductList(ctx context.Context, request TikTokAppV3FetchShopProductListRequest) (*TikTokAppV3FetchShopProductListResponse, error) {
	return r.client.TikTokAppV3GetShopProductListData(ctx, request)
}

// TikTokAppV3FetchShopProductListV2Request is the request for GET /api/v1/tiktok/app/v3/fetch_shop_product_list_v2.
type TikTokAppV3FetchShopProductListV2Request = TikTokAppV3GetShopProductListDataV2Request

// TikTokAppV3FetchShopProductListV2Response is the response for GET /api/v1/tiktok/app/v3/fetch_shop_product_list_v2.
type TikTokAppV3FetchShopProductListV2Response = TikTokAppV3GetShopProductListDataV2Response

// FetchShopProductListV2 获取商家商品列表数据 V2/Get shop product list data V2
//
// GET /api/v1/tiktok/app/v3/fetch_shop_product_list_v2
func (r TikTokAppV3Resource) FetchShopProductListV2(ctx context.Context, request TikTokAppV3FetchShopProductListV2Request) (*TikTokAppV3FetchShopProductListV2Response, error) {
	return r.client.TikTokAppV3GetShopProductListDataV2(ctx, request)
}

// TikTokAppV3FetchShopInfoRequest is the request for GET /api/v1/tiktok/app/v3/fetch_shop_info.
type TikTokAppV3FetchShopInfoRequest = TikTokAppV3GetShopInformationDataRequest

// TikTokAppV3FetchShopInfoResponse is the response for GET /api/v1/tiktok/app/v3/fetch_shop_info.
type TikTokAppV3FetchShopInfoResponse = TikTokAppV3GetShopInformationDataResponse

// FetchShopInfo 获取商家信息数据/Get shop information data
//
// GET /api/v1/tiktok/app/v3/fetch_shop_info
func (r TikTokAppV3Resource) FetchShopInfo(ctx context.Context, request TikTokAppV3FetchShopInfoRequest) (*TikTokAppV3FetchShopInfoResponse, error) {
	return r.client.TikTokAppV3GetShopInformationData(ctx, request)
}

// TikTokAppV3FetchShopProductCategoryRequest is the request for GET /api/v1/tiktok/app/v3/fetch_shop_product_category.
type TikTokAppV3FetchShopProductCategoryRequest = TikTokAppV3GetShopProductCategoryDataRequest

// TikTokAppV3FetchShopProductCategoryResponse is the response for GET /api/v1/tiktok/app/v3/fetch_shop_product_category.
type TikTokAppV3FetchShopProductCategoryResponse = TikTokAppV3GetShopProductCategoryDataResponse

// FetchShopProductCategory 获取商家产品分类数据/Get shop product category data
//
// GET /api/v1/tiktok/app/v3/fetch_shop_product_category
func (r TikTokAppV3Resource) FetchShopProductCategory(ctx context.Context, request TikTokAppV3FetchShopProductCategoryRequest) (*TikTokAppV3FetchShopProductCategoryResponse, error) {
	return r.client.TikTokAppV3GetShopProductCategoryData(ctx, request)
}

// TikTokAppV3FetchLiveDailyRankRequest is the request for GET /api/v1/tiktok/app/v3/fetch_live_daily_rank.
type TikTokAppV3FetchLiveDailyRankRequest = TikTokAppV3GetLiveDailyRankDataRequest

// TikTokAppV3FetchLiveDailyRankResponse is the response for GET /api/v1/tiktok/app/v3/fetch_live_daily_rank.
type TikTokAppV3FetchLiveDailyRankResponse = TikTokAppV3GetLiveDailyRankDataResponse

// FetchLiveDailyRank 获取直播每日榜单数据/Get live daily rank data
//
// GET /api/v1/tiktok/app/v3/fetch_live_daily_rank
func (r TikTokAppV3Resource) FetchLiveDailyRank(ctx context.Context, request TikTokAppV3FetchLiveDailyRankRequest) (*TikTokAppV3FetchLiveDailyRankResponse, error) {
	return r.client.TikTokAppV3GetLiveDailyRankData(ctx, request)
}

// TikTokAppV3FetchUserMusicListRequest is the request for GET /api/v1/tiktok/app/v3/fetch_user_music_list.
type TikTokAppV3FetchUserMusicListRequest = TikTokAppV3GetUserMusicListDataRequest

// TikTokAppV3FetchUserMusicListResponse is the response for GET /api/v1/tiktok/app/v3/fetch_user_music_list.
type TikTokAppV3FetchUserMusicListResponse = TikTokAppV3GetUserMusicListDataResponse

// FetchUserMusicList 获取用户音乐列表数据/Get user music list data
//
// GET /api/v1/tiktok/app/v3/fetch_user_music_list
func (r TikTokAppV3Resource) FetchUserMusicList(ctx context.Context, request TikTokAppV3FetchUserMusicListRequest) (*TikTokAppV3FetchUserMusicListResponse, error) {
	return r.client.TikTokAppV3GetUserMusicListData(ctx, request)
}

// TikTokAppV3FetchContentTranslateRequest is the request for POST /api/v1/tiktok/app/v3/fetch_content_translate.
type TikTokAppV3FetchContentTranslateRequest = TikTokAppV3GetContentTranslationDataRequest

// TikTokAppV3FetchContentTranslateResponse is the response for POST /api/v1/tiktok/app/v3/fetch_content_translate.
type TikTokAppV3FetchContentTranslateResponse = TikTokAppV3GetContentTranslationDataResponse

// FetchContentTranslate 获取内容翻译数据/Get content translation data
//
// POST /api/v1/tiktok/app/v3/fetch_content_translate
func (r TikTokAppV3Resource) FetchContentTranslate(ctx context.Context, request TikTokAppV3FetchContentTranslateRequest) (*TikTokAppV3FetchContentTranslateResponse, error) {
	return r.client.TikTokAppV3GetContentTranslationData(ctx, request)
}

// TikTokAppV3FetchHomeFeedRequest is the request for POST /api/v1/tiktok/app/v3/fetch_home_feed.
type TikTokAppV3FetchHomeFeedRequest = TikTokAppV3GetHomeFeedVideoDataRequest

// TikTokAppV3FetchHomeFeedResponse is the response for POST /api/v1/tiktok/app/v3/fetch_home_feed.
type TikTokAppV3FetchHomeFeedResponse = TikTokAppV3GetHomeFeedVideoDataResponse

// FetchHomeFeed 获取主页视频推荐数据/Get home feed(recommend) video data
//
// POST /api/v1/tiktok/app/v3/fetch_home_feed
func (r TikTokAppV3Resource) FetchHomeFeed(ctx context.Context, request TikTokAppV3FetchHomeFeedRequest) (*TikTokAppV3FetchHomeFeedResponse, error) {
	return r.client.TikTokAppV3GetHomeFeedVideoData(ctx, request)
}

// TikTokAppV3TTencryptAlgorithmRequest is the request for POST /api/v1/tiktok/app/v3/TTencrypt_algorithm.
type TikTokAppV3TTencryptAlgorithmRequest = TikTokAppV3TikTokAppEncryptionAlgorithmRequest

// TikTokAppV3TTencryptAlgorithmResponse is the response for POST /api/v1/tiktok/app/v3/TTencrypt_algorithm.
type TikTokAppV3TTencryptAlgorithmResponse = TikTokAppV3TikTokAppEncryptionAlgorithmResponse

// TTencryptAlgorithm TikTok APP加密算法/TikTok APP encryption algorithm
//
// POST /api/v1/tiktok/app/v3/TTencrypt_algorithm
func (r TikTokAppV3Resource) TTencryptAlgorithm(ctx context.Context, request TikTokAppV3TTencryptAlgorithmRequest) (*TikTokAppV3TTencryptAlgorithmResponse, error) {
	return r.client.TikTokAppV3TikTokAppEncryptionAlgorithm(ctx, request)
}

// TikTokAppV3FetchLiveRoomProductListRequest is the request for GET /api/v1/tiktok/app/v3/fetch_live_room_product_list.
type TikTokAppV3FetchLiveRoomProductListRequest = TikTokAppV3GetLiveRoomProductListDataRequest

// TikTokAppV3FetchLiveRoomProductListResponse is the response for GET /api/v1/tiktok/app/v3/fetch_live_room_product_list.
type TikTokAppV3FetchLiveRoomProductListResponse = TikTokAppV3GetLiveRoomProductListDataResponse

// FetchLiveRoomProductList 获取直播间商品列表数据/Get live room product list data
//
// GET /api/v1/tiktok/app/v3/fetch_live_room_product_list
func (r TikTokAppV3Resource) FetchLiveRoomProductList(ctx context.Context, request TikTokAppV3FetchLiveRoomProductListRequest) (*TikTokAppV3FetchLiveRoomProductListResponse, error) {
	return r.client.TikTokAppV3GetLiveRoomProductListData(ctx, request)
}

// TikTokAppV3FetchLiveRoomProductListV2Request is the request for GET /api/v1/tiktok/app/v3/fetch_live_room_product_list_v2.
type TikTokAppV3FetchLiveRoomProductListV2Request = TikTokAppV3GetLiveRoomProductListDataV2Request

// TikTokAppV3FetchLiveRoomProductListV2Response is the response for GET /api/v1/tiktok/app/v3/fetch_live_room_product_list_v2.
type TikTokAppV3FetchLiveRoomProductListV2Response = TikTokAppV3GetLiveRoomProductListDataV2Response

// FetchLiveRoomProductListV2 获取直播间商品列表数据 V2 /Get live room product list data V2
//
// GET /api/v1/tiktok/app/v3/fetch_live_room_product_list_v2
func (r TikTokAppV3Resource) FetchLiveRoomProductListV2(ctx context.Context, request TikTokAppV3FetchLiveRoomProductListV2Request) (*TikTokAppV3FetchLiveRoomProductListV2Response, error) {
	return r.client.TikTokAppV3GetLiveRoomProductListDataV2(ctx, request)
}

// TikTokAppV3AddVideoPlayCountRequest is the request for GET /api/v1/tiktok/app/v3/add_video_play_count.
type TikTokAppV3AddVideoPlayCountRequest = TikTokAppV3IncreaseTheNumberOfPlaysOfTheWorkAccordingToTheVideoIDRequest

// TikTokAppV3AddVideoPlayCountResponse is the response for GET /api/v1/tiktok/app/v3/add_video_play_count.
type TikTokAppV3AddVideoPlayCountResponse = TikTokAppV3IncreaseTheNumberOfPlaysOfTheWorkAccordingToTheVideoIDResponse

// AddVideoPlayCount 根据视频ID来增加作品的播放数/Increase the number of plays of the work according to the video ID
//
// GET /api/v1/tiktok/app/v3/add_video_play_count
func (r TikTokAppV3Resource) AddVideoPlayCount(ctx context.Context, request TikTokAppV3AddVideoPlayCountRequest) (*TikTokAppV3AddVideoPlayCountResponse, error) {
	return r.client.TikTokAppV3IncreaseTheNumberOfPlaysOfTheWorkAccordingToTheVideoID(ctx, request)
}

// TikTokAppV3EncryptDecryptLoginRequestRequest is the request for POST /api/v1/tiktok/app/v3/encrypt_decrypt_login_request.
type TikTokAppV3EncryptDecryptLoginRequestRequest = TikTokAppV3EncryptOrDecryptTikTokAppLoginRequestBodyRequest

// TikTokAppV3EncryptDecryptLoginRequestResponse is the response for POST /api/v1/tiktok/app/v3/encrypt_decrypt_login_request.
type TikTokAppV3EncryptDecryptLoginRequestResponse = TikTokAppV3EncryptOrDecryptTikTokAppLoginRequestBodyResponse

// EncryptDecryptLoginRequest 加密或解密 TikTok APP 登录请求体/Encrypt or Decrypt TikTok APP login request body
//
// POST /api/v1/tiktok/app/v3/encrypt_decrypt_login_request
func (r TikTokAppV3Resource) EncryptDecryptLoginRequest(ctx context.Context, request TikTokAppV3EncryptDecryptLoginRequestRequest) (*TikTokAppV3EncryptDecryptLoginRequestResponse, error) {
	return r.client.TikTokAppV3EncryptOrDecryptTikTokAppLoginRequestBody(ctx, request)
}

// TikTokAppV3OpenTikTokAppToVideoDetailRequest is the request for GET /api/v1/tiktok/app/v3/open_tiktok_app_to_video_detail.
type TikTokAppV3OpenTikTokAppToVideoDetailRequest = TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndJumpToTheSpecifiedVideoDetailsPageRequest

// TikTokAppV3OpenTikTokAppToVideoDetailResponse is the response for GET /api/v1/tiktok/app/v3/open_tiktok_app_to_video_detail.
type TikTokAppV3OpenTikTokAppToVideoDetailResponse = TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndJumpToTheSpecifiedVideoDetailsPageResponse

// OpenTikTokAppToVideoDetail 生成TikTok分享链接，唤起TikTok APP，跳转指定作品详情页/Generate TikTok share link, call TikTok APP, and jump to the specified video details page
//
// GET /api/v1/tiktok/app/v3/open_tiktok_app_to_video_detail
func (r TikTokAppV3Resource) OpenTikTokAppToVideoDetail(ctx context.Context, request TikTokAppV3OpenTikTokAppToVideoDetailRequest) (*TikTokAppV3OpenTikTokAppToVideoDetailResponse, error) {
	return r.client.TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndJumpToTheSpecifiedVideoDetailsPage(ctx, request)
}

// TikTokAppV3OpenTikTokAppToUserProfileRequest is the request for GET /api/v1/tiktok/app/v3/open_tiktok_app_to_user_profile.
type TikTokAppV3OpenTikTokAppToUserProfileRequest = TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndJumpToTheSpecifiedUserProfileRequest

// TikTokAppV3OpenTikTokAppToUserProfileResponse is the response for GET /api/v1/tiktok/app/v3/open_tiktok_app_to_user_profile.
type TikTokAppV3OpenTikTokAppToUserProfileResponse = TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndJumpToTheSpecifiedUserProfileResponse

// OpenTikTokAppToUserProfile 生成TikTok分享链接，唤起TikTok APP，跳转指定用户主页/Generate TikTok share link, call TikTok APP, and jump to the specified user profile
//
// GET /api/v1/tiktok/app/v3/open_tiktok_app_to_user_profile
func (r TikTokAppV3Resource) OpenTikTokAppToUserProfile(ctx context.Context, request TikTokAppV3OpenTikTokAppToUserProfileRequest) (*TikTokAppV3OpenTikTokAppToUserProfileResponse, error) {
	return r.client.TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndJumpToTheSpecifiedUserProfile(ctx, request)
}

// TikTokAppV3OpenTikTokAppToKeywordSearchRequest is the request for GET /api/v1/tiktok/app/v3/open_tiktok_app_to_keyword_search.
type TikTokAppV3OpenTikTokAppToKeywordSearchRequest = TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndJumpToTheSpecifiedKeywordSearchResultRequest

// TikTokAppV3OpenTikTokAppToKeywordSearchResponse is the response for GET /api/v1/tiktok/app/v3/open_tiktok_app_to_keyword_search.
type TikTokAppV3OpenTikTokAppToKeywordSearchResponse = TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndJumpToTheSpecifiedKeywordSearchResultResponse

// OpenTikTokAppToKeywordSearch 生成TikTok分享链接，唤起TikTok APP，跳转指定关键词搜索结果/Generate TikTok share link, call TikTok APP, and jump to the specified keyword search result
//
// GET /api/v1/tiktok/app/v3/open_tiktok_app_to_keyword_search
func (r TikTokAppV3Resource) OpenTikTokAppToKeywordSearch(ctx context.Context, request TikTokAppV3OpenTikTokAppToKeywordSearchRequest) (*TikTokAppV3OpenTikTokAppToKeywordSearchResponse, error) {
	return r.client.TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndJumpToTheSpecifiedKeywordSearchResult(ctx, request)
}

// TikTokAppV3OpenTikTokAppToSendPrivateMessageRequest is the request for GET /api/v1/tiktok/app/v3/open_tiktok_app_to_send_private_message.
type TikTokAppV3OpenTikTokAppToSendPrivateMessageRequest = TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndSendPrivateMessagesToSpecifiedUsersRequest

// TikTokAppV3OpenTikTokAppToSendPrivateMessageResponse is the response for GET /api/v1/tiktok/app/v3/open_tiktok_app_to_send_private_message.
type TikTokAppV3OpenTikTokAppToSendPrivateMessageResponse = TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndSendPrivateMessagesToSpecifiedUsersResponse

// OpenTikTokAppToSendPrivateMessage 生成TikTok分享链接，唤起TikTok APP，给指定用户发送私信/Generate TikTok share link, call TikTok APP, and send private messages to specified users
//
// GET /api/v1/tiktok/app/v3/open_tiktok_app_to_send_private_message
func (r TikTokAppV3Resource) OpenTikTokAppToSendPrivateMessage(ctx context.Context, request TikTokAppV3OpenTikTokAppToSendPrivateMessageRequest) (*TikTokAppV3OpenTikTokAppToSendPrivateMessageResponse, error) {
	return r.client.TikTokAppV3GenerateTikTokShareLinkCallTikTokAppAndSendPrivateMessagesToSpecifiedUsers(ctx, request)
}

// TikTokCreatorResource contains endpoints from the TikTok-Creator-API tag.
type TikTokCreatorResource struct {
	client *Client
}

// TikTokCreatorGetAccountHealthStatusRequest is the request for POST /api/v1/tiktok/creator/get_account_health_status.
type TikTokCreatorGetAccountHealthStatusRequest = TikTokCreatorGetCreatorAccountHealthStatusRequest

// TikTokCreatorGetAccountHealthStatusResponse is the response for POST /api/v1/tiktok/creator/get_account_health_status.
type TikTokCreatorGetAccountHealthStatusResponse = TikTokCreatorGetCreatorAccountHealthStatusResponse

// GetAccountHealthStatus 获取创作者账号健康状态/Get Creator Account Health Status
//
// POST /api/v1/tiktok/creator/get_account_health_status
func (r TikTokCreatorResource) GetAccountHealthStatus(ctx context.Context, request TikTokCreatorGetAccountHealthStatusRequest) (*TikTokCreatorGetAccountHealthStatusResponse, error) {
	return r.client.TikTokCreatorGetCreatorAccountHealthStatus(ctx, request)
}

// TikTokCreatorGetAccountViolationListRequest is the request for POST /api/v1/tiktok/creator/get_account_violation_list.
type TikTokCreatorGetAccountViolationListRequest = TikTokCreatorGetCreatorAccountViolationRecordListRequest

// TikTokCreatorGetAccountViolationListResponse is the response for POST /api/v1/tiktok/creator/get_account_violation_list.
type TikTokCreatorGetAccountViolationListResponse = TikTokCreatorGetCreatorAccountViolationRecordListResponse

// GetAccountViolationList 获取创作者账号违规记录列表/Get Creator Account Violation Record List
//
// POST /api/v1/tiktok/creator/get_account_violation_list
func (r TikTokCreatorResource) GetAccountViolationList(ctx context.Context, request TikTokCreatorGetAccountViolationListRequest) (*TikTokCreatorGetAccountViolationListResponse, error) {
	return r.client.TikTokCreatorGetCreatorAccountViolationRecordList(ctx, request)
}

// TikTokCreatorGetAccountInsightsOverviewRequest is the request for POST /api/v1/tiktok/creator/get_account_insights_overview.
type TikTokCreatorGetAccountInsightsOverviewRequest = TikTokCreatorGetCreatorAccountOverviewRequest

// TikTokCreatorGetAccountInsightsOverviewResponse is the response for POST /api/v1/tiktok/creator/get_account_insights_overview.
type TikTokCreatorGetAccountInsightsOverviewResponse = TikTokCreatorGetCreatorAccountOverviewResponse

// GetAccountInsightsOverview 获取创作者账号概览/Get Creator Account Overview
//
// POST /api/v1/tiktok/creator/get_account_insights_overview
func (r TikTokCreatorResource) GetAccountInsightsOverview(ctx context.Context, request TikTokCreatorGetAccountInsightsOverviewRequest) (*TikTokCreatorGetAccountInsightsOverviewResponse, error) {
	return r.client.TikTokCreatorGetCreatorAccountOverview(ctx, request)
}

// TikTokCreatorGetLiveAnalyticsSummaryRequest is the request for POST /api/v1/tiktok/creator/get_live_analytics_summary.
type TikTokCreatorGetLiveAnalyticsSummaryRequest = TikTokCreatorGetCreatorLiveOverviewRequest

// TikTokCreatorGetLiveAnalyticsSummaryResponse is the response for POST /api/v1/tiktok/creator/get_live_analytics_summary.
type TikTokCreatorGetLiveAnalyticsSummaryResponse = TikTokCreatorGetCreatorLiveOverviewResponse

// GetLiveAnalyticsSummary 获取创作者直播概览/Get Creator Live Overview
//
// POST /api/v1/tiktok/creator/get_live_analytics_summary
func (r TikTokCreatorResource) GetLiveAnalyticsSummary(ctx context.Context, request TikTokCreatorGetLiveAnalyticsSummaryRequest) (*TikTokCreatorGetLiveAnalyticsSummaryResponse, error) {
	return r.client.TikTokCreatorGetCreatorLiveOverview(ctx, request)
}

// TikTokCreatorGetVideoAnalyticsSummaryRequest is the request for POST /api/v1/tiktok/creator/get_video_analytics_summary.
type TikTokCreatorGetVideoAnalyticsSummaryRequest = TikTokCreatorGetCreatorVideoOverviewRequest

// TikTokCreatorGetVideoAnalyticsSummaryResponse is the response for POST /api/v1/tiktok/creator/get_video_analytics_summary.
type TikTokCreatorGetVideoAnalyticsSummaryResponse = TikTokCreatorGetCreatorVideoOverviewResponse

// GetVideoAnalyticsSummary 获取创作者视频概览/Get Creator Video Overview
//
// POST /api/v1/tiktok/creator/get_video_analytics_summary
func (r TikTokCreatorResource) GetVideoAnalyticsSummary(ctx context.Context, request TikTokCreatorGetVideoAnalyticsSummaryRequest) (*TikTokCreatorGetVideoAnalyticsSummaryResponse, error) {
	return r.client.TikTokCreatorGetCreatorVideoOverview(ctx, request)
}

// TikTokCreatorGetVideoListAnalyticsRequest is the request for POST /api/v1/tiktok/creator/get_video_list_analytics.
type TikTokCreatorGetVideoListAnalyticsRequest = TikTokCreatorGetCreatorVideoListAnalyticsRequest

// TikTokCreatorGetVideoListAnalyticsResponse is the response for POST /api/v1/tiktok/creator/get_video_list_analytics.
type TikTokCreatorGetVideoListAnalyticsResponse = TikTokCreatorGetCreatorVideoListAnalyticsResponse

// GetVideoListAnalytics 获取创作者视频列表分析/Get Creator Video List Analytics
//
// POST /api/v1/tiktok/creator/get_video_list_analytics
func (r TikTokCreatorResource) GetVideoListAnalytics(ctx context.Context, request TikTokCreatorGetVideoListAnalyticsRequest) (*TikTokCreatorGetVideoListAnalyticsResponse, error) {
	return r.client.TikTokCreatorGetCreatorVideoListAnalytics(ctx, request)
}

// TikTokCreatorGetProductAnalyticsListRequest is the request for POST /api/v1/tiktok/creator/get_product_analytics_list.
type TikTokCreatorGetProductAnalyticsListRequest = TikTokCreatorGetCreatorProductListAnalyticsRequest

// TikTokCreatorGetProductAnalyticsListResponse is the response for POST /api/v1/tiktok/creator/get_product_analytics_list.
type TikTokCreatorGetProductAnalyticsListResponse = TikTokCreatorGetCreatorProductListAnalyticsResponse

// GetProductAnalyticsList 获取创作者商品列表分析/Get Creator Product List Analytics
//
// POST /api/v1/tiktok/creator/get_product_analytics_list
func (r TikTokCreatorResource) GetProductAnalyticsList(ctx context.Context, request TikTokCreatorGetProductAnalyticsListRequest) (*TikTokCreatorGetProductAnalyticsListResponse, error) {
	return r.client.TikTokCreatorGetCreatorProductListAnalytics(ctx, request)
}

// GetCreatorAccountInfo 获取创作者账号信息/Get Creator Account Info
//
// POST /api/v1/tiktok/creator/get_creator_account_info
func (r TikTokCreatorResource) GetCreatorAccountInfo(ctx context.Context, request TikTokCreatorGetCreatorAccountInfoRequest) (*TikTokCreatorGetCreatorAccountInfoResponse, error) {
	return r.client.TikTokCreatorGetCreatorAccountInfo(ctx, request)
}

// GetShowcaseProductList 获取橱窗商品列表/Get Showcase Product List
//
// POST /api/v1/tiktok/creator/get_showcase_product_list
func (r TikTokCreatorResource) GetShowcaseProductList(ctx context.Context, request TikTokCreatorGetShowcaseProductListRequest) (*TikTokCreatorGetShowcaseProductListResponse, error) {
	return r.client.TikTokCreatorGetShowcaseProductList(ctx, request)
}

// GetVideoAssociatedProductList 获取视频关联商品列表/Get Video Associated Product List
//
// POST /api/v1/tiktok/creator/get_video_associated_product_list
func (r TikTokCreatorResource) GetVideoAssociatedProductList(ctx context.Context, request TikTokCreatorGetVideoAssociatedProductListRequest) (*TikTokCreatorGetVideoAssociatedProductListResponse, error) {
	return r.client.TikTokCreatorGetVideoAssociatedProductList(ctx, request)
}

// TikTokCreatorGetVideoDetailedStatsRequest is the request for POST /api/v1/tiktok/creator/get_video_detailed_stats.
type TikTokCreatorGetVideoDetailedStatsRequest = TikTokCreatorGetVideoDetailedStatisticsRequest

// TikTokCreatorGetVideoDetailedStatsResponse is the response for POST /api/v1/tiktok/creator/get_video_detailed_stats.
type TikTokCreatorGetVideoDetailedStatsResponse = TikTokCreatorGetVideoDetailedStatisticsResponse

// GetVideoDetailedStats 获取视频详细分段统计数据/Get Video Detailed Statistics
//
// POST /api/v1/tiktok/creator/get_video_detailed_stats
func (r TikTokCreatorResource) GetVideoDetailedStats(ctx context.Context, request TikTokCreatorGetVideoDetailedStatsRequest) (*TikTokCreatorGetVideoDetailedStatsResponse, error) {
	return r.client.TikTokCreatorGetVideoDetailedStatistics(ctx, request)
}

// TikTokCreatorGetVideoToProductStatsRequest is the request for POST /api/v1/tiktok/creator/get_video_to_product_stats.
type TikTokCreatorGetVideoToProductStatsRequest = TikTokCreatorGetVideoProductAssociationStatisticsRequest

// TikTokCreatorGetVideoToProductStatsResponse is the response for POST /api/v1/tiktok/creator/get_video_to_product_stats.
type TikTokCreatorGetVideoToProductStatsResponse = TikTokCreatorGetVideoProductAssociationStatisticsResponse

// GetVideoToProductStats 获取视频与商品关联统计数据/Get Video-Product Association Statistics
//
// POST /api/v1/tiktok/creator/get_video_to_product_stats
func (r TikTokCreatorResource) GetVideoToProductStats(ctx context.Context, request TikTokCreatorGetVideoToProductStatsRequest) (*TikTokCreatorGetVideoToProductStatsResponse, error) {
	return r.client.TikTokCreatorGetVideoProductAssociationStatistics(ctx, request)
}

// GetProductRelatedVideos 获取同款商品关联视频/Get Product Related Videos
//
// POST /api/v1/tiktok/creator/get_product_related_videos
func (r TikTokCreatorResource) GetProductRelatedVideos(ctx context.Context, request TikTokCreatorGetProductRelatedVideosRequest) (*TikTokCreatorGetProductRelatedVideosResponse, error) {
	return r.client.TikTokCreatorGetProductRelatedVideos(ctx, request)
}

// TikTokCreatorGetVideoAudienceStatsRequest is the request for POST /api/v1/tiktok/creator/get_video_audience_stats.
type TikTokCreatorGetVideoAudienceStatsRequest = TikTokCreatorGetVideoAudienceAnalysisDataRequest

// TikTokCreatorGetVideoAudienceStatsResponse is the response for POST /api/v1/tiktok/creator/get_video_audience_stats.
type TikTokCreatorGetVideoAudienceStatsResponse = TikTokCreatorGetVideoAudienceAnalysisDataResponse

// GetVideoAudienceStats 获取视频受众分析数据/Get Video Audience Analysis Data
//
// POST /api/v1/tiktok/creator/get_video_audience_stats
func (r TikTokCreatorResource) GetVideoAudienceStats(ctx context.Context, request TikTokCreatorGetVideoAudienceStatsRequest) (*TikTokCreatorGetVideoAudienceStatsResponse, error) {
	return r.client.TikTokCreatorGetVideoAudienceAnalysisData(ctx, request)
}

// TikTokAnalyticsResource contains endpoints from the TikTok-Analytics-API tag.
type TikTokAnalyticsResource struct {
	client *Client
}

// TikTokAnalyticsFetchVideoMetricsRequest is the request for GET /api/v1/tiktok/analytics/fetch_video_metrics.
type TikTokAnalyticsFetchVideoMetricsRequest = TikTokAnalyticsGetVideoMetricsRequest

// TikTokAnalyticsFetchVideoMetricsResponse is the response for GET /api/v1/tiktok/analytics/fetch_video_metrics.
type TikTokAnalyticsFetchVideoMetricsResponse = TikTokAnalyticsGetVideoMetricsResponse

// FetchVideoMetrics 获取作品的统计数据/Get video metrics
//
// GET /api/v1/tiktok/analytics/fetch_video_metrics
func (r TikTokAnalyticsResource) FetchVideoMetrics(ctx context.Context, request TikTokAnalyticsFetchVideoMetricsRequest) (*TikTokAnalyticsFetchVideoMetricsResponse, error) {
	return r.client.TikTokAnalyticsGetVideoMetrics(ctx, request)
}

// TikTokAnalyticsDetectFakeViewsRequest is the request for GET /api/v1/tiktok/analytics/detect_fake_views.
type TikTokAnalyticsDetectFakeViewsRequest = TikTokAnalyticsDetectFakeViewsInVideoRequest

// TikTokAnalyticsDetectFakeViewsResponse is the response for GET /api/v1/tiktok/analytics/detect_fake_views.
type TikTokAnalyticsDetectFakeViewsResponse = TikTokAnalyticsDetectFakeViewsInVideoResponse

// DetectFakeViews 检测视频虚假流量分析/Detect fake views in video
//
// GET /api/v1/tiktok/analytics/detect_fake_views
func (r TikTokAnalyticsResource) DetectFakeViews(ctx context.Context, request TikTokAnalyticsDetectFakeViewsRequest) (*TikTokAnalyticsDetectFakeViewsResponse, error) {
	return r.client.TikTokAnalyticsDetectFakeViewsInVideo(ctx, request)
}

// TikTokAnalyticsFetchCommentKeywordsRequest is the request for GET /api/v1/tiktok/analytics/fetch_comment_keywords.
type TikTokAnalyticsFetchCommentKeywordsRequest = TikTokAnalyticsGetCommentKeywordsAnalysisRequest

// TikTokAnalyticsFetchCommentKeywordsResponse is the response for GET /api/v1/tiktok/analytics/fetch_comment_keywords.
type TikTokAnalyticsFetchCommentKeywordsResponse = TikTokAnalyticsGetCommentKeywordsAnalysisResponse

// FetchCommentKeywords 获取视频评论关键词分析/Get comment keywords analysis
//
// GET /api/v1/tiktok/analytics/fetch_comment_keywords
func (r TikTokAnalyticsResource) FetchCommentKeywords(ctx context.Context, request TikTokAnalyticsFetchCommentKeywordsRequest) (*TikTokAnalyticsFetchCommentKeywordsResponse, error) {
	return r.client.TikTokAnalyticsGetCommentKeywordsAnalysis(ctx, request)
}

// TikTokAnalyticsFetchCreatorInfoAndMilestonesRequest is the request for GET /api/v1/tiktok/analytics/fetch_creator_info_and_milestones.
type TikTokAnalyticsFetchCreatorInfoAndMilestonesRequest = TikTokAnalyticsGetCreatorInfoAndMilestonesRequest

// TikTokAnalyticsFetchCreatorInfoAndMilestonesResponse is the response for GET /api/v1/tiktok/analytics/fetch_creator_info_and_milestones.
type TikTokAnalyticsFetchCreatorInfoAndMilestonesResponse = TikTokAnalyticsGetCreatorInfoAndMilestonesResponse

// FetchCreatorInfoAndMilestones 获取创作者信息和里程碑数据/Get creator info and milestones
//
// GET /api/v1/tiktok/analytics/fetch_creator_info_and_milestones
func (r TikTokAnalyticsResource) FetchCreatorInfoAndMilestones(ctx context.Context, request TikTokAnalyticsFetchCreatorInfoAndMilestonesRequest) (*TikTokAnalyticsFetchCreatorInfoAndMilestonesResponse, error) {
	return r.client.TikTokAnalyticsGetCreatorInfoAndMilestones(ctx, request)
}

// TikTokAdsResource contains endpoints from the TikTok-Ads-API tag.
type TikTokAdsResource struct {
	client *Client
}

// TikTokAdsGetAdsDetailRequest is the request for GET /api/v1/tiktok/ads/get_ads_detail.
type TikTokAdsGetAdsDetailRequest = TikTokAdsGetSingleAdDetailRequest

// TikTokAdsGetAdsDetailResponse is the response for GET /api/v1/tiktok/ads/get_ads_detail.
type TikTokAdsGetAdsDetailResponse = TikTokAdsGetSingleAdDetailResponse

// GetAdsDetail 获取单个广告详情/Get single ad detail
//
// GET /api/v1/tiktok/ads/get_ads_detail
func (r TikTokAdsResource) GetAdsDetail(ctx context.Context, request TikTokAdsGetAdsDetailRequest) (*TikTokAdsGetAdsDetailResponse, error) {
	return r.client.TikTokAdsGetSingleAdDetail(ctx, request)
}

// SearchAds 搜索广告/Search ads
//
// GET /api/v1/tiktok/ads/search_ads
func (r TikTokAdsResource) SearchAds(ctx context.Context, request TikTokAdsSearchAdsRequest) (*TikTokAdsSearchAdsResponse, error) {
	return r.client.TikTokAdsSearchAds(ctx, request)
}

// TikTokAdsGetKeywordInsightsRequest is the request for GET /api/v1/tiktok/ads/get_keyword_insights.
type TikTokAdsGetKeywordInsightsRequest = TikTokAdsGetKeywordInsightsDataRequest

// TikTokAdsGetKeywordInsightsResponse is the response for GET /api/v1/tiktok/ads/get_keyword_insights.
type TikTokAdsGetKeywordInsightsResponse = TikTokAdsGetKeywordInsightsDataResponse

// GetKeywordInsights 获取关键词洞察数据/Get keyword insights data
//
// GET /api/v1/tiktok/ads/get_keyword_insights
func (r TikTokAdsResource) GetKeywordInsights(ctx context.Context, request TikTokAdsGetKeywordInsightsRequest) (*TikTokAdsGetKeywordInsightsResponse, error) {
	return r.client.TikTokAdsGetKeywordInsightsData(ctx, request)
}

// TikTokAdsGetTopProductsRequest is the request for GET /api/v1/tiktok/ads/get_top_products.
type TikTokAdsGetTopProductsRequest = TikTokAdsGetTopProductsListRequest

// TikTokAdsGetTopProductsResponse is the response for GET /api/v1/tiktok/ads/get_top_products.
type TikTokAdsGetTopProductsResponse = TikTokAdsGetTopProductsListResponse

// GetTopProducts 获取热门产品列表/Get top products list
//
// GET /api/v1/tiktok/ads/get_top_products
func (r TikTokAdsResource) GetTopProducts(ctx context.Context, request TikTokAdsGetTopProductsRequest) (*TikTokAdsGetTopProductsResponse, error) {
	return r.client.TikTokAdsGetTopProductsList(ctx, request)
}

// TikTokAdsGetHashtagListRequest is the request for GET /api/v1/tiktok/ads/get_hashtag_list.
type TikTokAdsGetHashtagListRequest = TikTokAdsGetPopularHashtagsListRequest

// TikTokAdsGetHashtagListResponse is the response for GET /api/v1/tiktok/ads/get_hashtag_list.
type TikTokAdsGetHashtagListResponse = TikTokAdsGetPopularHashtagsListResponse

// GetHashtagList 获取热门标签列表/Get popular hashtags list
//
// GET /api/v1/tiktok/ads/get_hashtag_list
func (r TikTokAdsResource) GetHashtagList(ctx context.Context, request TikTokAdsGetHashtagListRequest) (*TikTokAdsGetHashtagListResponse, error) {
	return r.client.TikTokAdsGetPopularHashtagsList(ctx, request)
}

// TikTokAdsGetSoundRankListRequest is the request for GET /api/v1/tiktok/ads/get_sound_rank_list.
type TikTokAdsGetSoundRankListRequest = TikTokAdsGetPopularSoundRankingsRequest

// TikTokAdsGetSoundRankListResponse is the response for GET /api/v1/tiktok/ads/get_sound_rank_list.
type TikTokAdsGetSoundRankListResponse = TikTokAdsGetPopularSoundRankingsResponse

// GetSoundRankList 获取热门音乐排行榜/Get popular sound rankings
//
// GET /api/v1/tiktok/ads/get_sound_rank_list
func (r TikTokAdsResource) GetSoundRankList(ctx context.Context, request TikTokAdsGetSoundRankListRequest) (*TikTokAdsGetSoundRankListResponse, error) {
	return r.client.TikTokAdsGetPopularSoundRankings(ctx, request)
}

// GetKeywordList 获取关键词列表/Get keyword list
//
// GET /api/v1/tiktok/ads/get_keyword_list
func (r TikTokAdsResource) GetKeywordList(ctx context.Context, request TikTokAdsGetKeywordListRequest) (*TikTokAdsGetKeywordListResponse, error) {
	return r.client.TikTokAdsGetKeywordList(ctx, request)
}

// GetTopAdsSpotlight 获取热门广告聚光灯/Get top ads spotlight
//
// GET /api/v1/tiktok/ads/get_top_ads_spotlight
func (r TikTokAdsResource) GetTopAdsSpotlight(ctx context.Context, request TikTokAdsGetTopAdsSpotlightRequest) (*TikTokAdsGetTopAdsSpotlightResponse, error) {
	return r.client.TikTokAdsGetTopAdsSpotlight(ctx, request)
}

// GetAdKeyframeAnalysis 获取广告关键帧分析/Get ad keyframe analysis
//
// GET /api/v1/tiktok/ads/get_ad_keyframe_analysis
func (r TikTokAdsResource) GetAdKeyframeAnalysis(ctx context.Context, request TikTokAdsGetAdKeyframeAnalysisRequest) (*TikTokAdsGetAdKeyframeAnalysisResponse, error) {
	return r.client.TikTokAdsGetAdKeyframeAnalysis(ctx, request)
}

// TikTokAdsGetAdPercentileRequest is the request for GET /api/v1/tiktok/ads/get_ad_percentile.
type TikTokAdsGetAdPercentileRequest = TikTokAdsGetAdPercentileDataRequest

// TikTokAdsGetAdPercentileResponse is the response for GET /api/v1/tiktok/ads/get_ad_percentile.
type TikTokAdsGetAdPercentileResponse = TikTokAdsGetAdPercentileDataResponse

// GetAdPercentile 获取广告百分位数据/Get ad percentile data
//
// GET /api/v1/tiktok/ads/get_ad_percentile
func (r TikTokAdsResource) GetAdPercentile(ctx context.Context, request TikTokAdsGetAdPercentileRequest) (*TikTokAdsGetAdPercentileResponse, error) {
	return r.client.TikTokAdsGetAdPercentileData(ctx, request)
}

// GetAdInteractiveAnalysis 获取广告互动分析/Get ad interactive analysis
//
// GET /api/v1/tiktok/ads/get_ad_interactive_analysis
func (r TikTokAdsResource) GetAdInteractiveAnalysis(ctx context.Context, request TikTokAdsGetAdInteractiveAnalysisRequest) (*TikTokAdsGetAdInteractiveAnalysisResponse, error) {
	return r.client.TikTokAdsGetAdInteractiveAnalysis(ctx, request)
}

// GetRecommendedAds 获取推荐广告/Get recommended ads
//
// GET /api/v1/tiktok/ads/get_recommended_ads
func (r TikTokAdsResource) GetRecommendedAds(ctx context.Context, request TikTokAdsGetRecommendedAdsRequest) (*TikTokAdsGetRecommendedAdsResponse, error) {
	return r.client.TikTokAdsGetRecommendedAds(ctx, request)
}

// GetQuerySuggestions 获取查询建议/Get query suggestions
//
// GET /api/v1/tiktok/ads/get_query_suggestions
func (r TikTokAdsResource) GetQuerySuggestions(ctx context.Context, request TikTokAdsGetQuerySuggestionsRequest) (*TikTokAdsGetQuerySuggestionsResponse, error) {
	return r.client.TikTokAdsGetQuerySuggestions(ctx, request)
}

// GetKeywordFilters 获取关键词筛选器/Get keyword filters
//
// GET /api/v1/tiktok/ads/get_keyword_filters
func (r TikTokAdsResource) GetKeywordFilters(ctx context.Context) (*TikTokAdsGetKeywordFiltersResponse, error) {
	return r.client.TikTokAdsGetKeywordFilters(ctx)
}

// GetRelatedKeywords 获取相关关键词/Get related keywords
//
// GET /api/v1/tiktok/ads/get_related_keywords
func (r TikTokAdsResource) GetRelatedKeywords(ctx context.Context, request TikTokAdsGetRelatedKeywordsRequest) (*TikTokAdsGetRelatedKeywordsResponse, error) {
	return r.client.TikTokAdsGetRelatedKeywords(ctx, request)
}

// GetKeywordDetails 获取关键词详细信息/Get keyword details
//
// GET /api/v1/tiktok/ads/get_keyword_details
func (r TikTokAdsResource) GetKeywordDetails(ctx context.Context, request TikTokAdsGetKeywordDetailsRequest) (*TikTokAdsGetKeywordDetailsResponse, error) {
	return r.client.TikTokAdsGetKeywordDetails(ctx, request)
}

// TikTokAdsGetCreativePatternsRequest is the request for GET /api/v1/tiktok/ads/get_creative_patterns.
type TikTokAdsGetCreativePatternsRequest = TikTokAdsGetCreativePatternRankingsRequest

// TikTokAdsGetCreativePatternsResponse is the response for GET /api/v1/tiktok/ads/get_creative_patterns.
type TikTokAdsGetCreativePatternsResponse = TikTokAdsGetCreativePatternRankingsResponse

// GetCreativePatterns 获取创意模式排行榜/Get creative pattern rankings
//
// GET /api/v1/tiktok/ads/get_creative_patterns
func (r TikTokAdsResource) GetCreativePatterns(ctx context.Context, request TikTokAdsGetCreativePatternsRequest) (*TikTokAdsGetCreativePatternsResponse, error) {
	return r.client.TikTokAdsGetCreativePatternRankings(ctx, request)
}

// GetProductFilters 获取产品筛选器/Get product filters
//
// GET /api/v1/tiktok/ads/get_product_filters
func (r TikTokAdsResource) GetProductFilters(ctx context.Context) (*TikTokAdsGetProductFiltersResponse, error) {
	return r.client.TikTokAdsGetProductFilters(ctx)
}

// GetProductMetrics 获取产品指标数据/Get product metrics
//
// GET /api/v1/tiktok/ads/get_product_metrics
func (r TikTokAdsResource) GetProductMetrics(ctx context.Context, request TikTokAdsGetProductMetricsRequest) (*TikTokAdsGetProductMetricsResponse, error) {
	return r.client.TikTokAdsGetProductMetrics(ctx, request)
}

// GetProductDetail 获取产品详细信息/Get product detail
//
// GET /api/v1/tiktok/ads/get_product_detail
func (r TikTokAdsResource) GetProductDetail(ctx context.Context, request TikTokAdsGetProductDetailRequest) (*TikTokAdsGetProductDetailResponse, error) {
	return r.client.TikTokAdsGetProductDetail(ctx, request)
}

// GetHashtagFilters 获取标签筛选器/Get hashtag filters
//
// GET /api/v1/tiktok/ads/get_hashtag_filters
func (r TikTokAdsResource) GetHashtagFilters(ctx context.Context) (*TikTokAdsGetHashtagFiltersResponse, error) {
	return r.client.TikTokAdsGetHashtagFilters(ctx)
}

// TikTokAdsGetHashtagCreatorRequest is the request for GET /api/v1/tiktok/ads/get_hashtag_creator.
type TikTokAdsGetHashtagCreatorRequest = TikTokAdsGetHashtagCreatorInfoRequest

// TikTokAdsGetHashtagCreatorResponse is the response for GET /api/v1/tiktok/ads/get_hashtag_creator.
type TikTokAdsGetHashtagCreatorResponse = TikTokAdsGetHashtagCreatorInfoResponse

// GetHashtagCreator 获取标签创作者信息/Get hashtag creator info
//
// GET /api/v1/tiktok/ads/get_hashtag_creator
func (r TikTokAdsResource) GetHashtagCreator(ctx context.Context, request TikTokAdsGetHashtagCreatorRequest) (*TikTokAdsGetHashtagCreatorResponse, error) {
	return r.client.TikTokAdsGetHashtagCreatorInfo(ctx, request)
}

// GetSoundFilters 获取音乐筛选器/Get sound filters
//
// GET /api/v1/tiktok/ads/get_sound_filters
func (r TikTokAdsResource) GetSoundFilters(ctx context.Context, request TikTokAdsGetSoundFiltersRequest) (*TikTokAdsGetSoundFiltersResponse, error) {
	return r.client.TikTokAdsGetSoundFilters(ctx, request)
}

// GetSoundDetail 获取音乐详情/Get sound detail
//
// GET /api/v1/tiktok/ads/get_sound_detail
func (r TikTokAdsResource) GetSoundDetail(ctx context.Context, request TikTokAdsGetSoundDetailRequest) (*TikTokAdsGetSoundDetailResponse, error) {
	return r.client.TikTokAdsGetSoundDetail(ctx, request)
}

// TikTokAdsSearchSoundHintRequest is the request for GET /api/v1/tiktok/ads/search_sound_hint.
type TikTokAdsSearchSoundHintRequest = TikTokAdsSearchSoundHintsRequest

// TikTokAdsSearchSoundHintResponse is the response for GET /api/v1/tiktok/ads/search_sound_hint.
type TikTokAdsSearchSoundHintResponse = TikTokAdsSearchSoundHintsResponse

// SearchSoundHint 搜索音乐提示/Search sound hints
//
// GET /api/v1/tiktok/ads/search_sound_hint
func (r TikTokAdsResource) SearchSoundHint(ctx context.Context, request TikTokAdsSearchSoundHintRequest) (*TikTokAdsSearchSoundHintResponse, error) {
	return r.client.TikTokAdsSearchSoundHints(ctx, request)
}

// TikTokAdsSearchSoundRequest is the request for GET /api/v1/tiktok/ads/search_sound.
type TikTokAdsSearchSoundRequest = TikTokAdsSearchSoundsRequest

// TikTokAdsSearchSoundResponse is the response for GET /api/v1/tiktok/ads/search_sound.
type TikTokAdsSearchSoundResponse = TikTokAdsSearchSoundsResponse

// SearchSound 搜索音乐/Search sounds
//
// GET /api/v1/tiktok/ads/search_sound
func (r TikTokAdsResource) SearchSound(ctx context.Context, request TikTokAdsSearchSoundRequest) (*TikTokAdsSearchSoundResponse, error) {
	return r.client.TikTokAdsSearchSounds(ctx, request)
}

// GetSoundRecommendations 获取音乐推荐/Get sound recommendations
//
// GET /api/v1/tiktok/ads/get_sound_recommendations
func (r TikTokAdsResource) GetSoundRecommendations(ctx context.Context, request TikTokAdsGetSoundRecommendationsRequest) (*TikTokAdsGetSoundRecommendationsResponse, error) {
	return r.client.TikTokAdsGetSoundRecommendations(ctx, request)
}

// GetCreatorFilters 获取创作者筛选器/Get creator filters
//
// GET /api/v1/tiktok/ads/get_creator_filters
func (r TikTokAdsResource) GetCreatorFilters(ctx context.Context) (*TikTokAdsGetCreatorFiltersResponse, error) {
	return r.client.TikTokAdsGetCreatorFilters(ctx)
}

// GetCreatorList 获取创作者列表/Get creator list
//
// GET /api/v1/tiktok/ads/get_creator_list
func (r TikTokAdsResource) GetCreatorList(ctx context.Context, request TikTokAdsGetCreatorListRequest) (*TikTokAdsGetCreatorListResponse, error) {
	return r.client.TikTokAdsGetCreatorList(ctx, request)
}

// SearchCreators 搜索创作者/Search creators
//
// GET /api/v1/tiktok/ads/search_creators
func (r TikTokAdsResource) SearchCreators(ctx context.Context, request TikTokAdsSearchCreatorsRequest) (*TikTokAdsSearchCreatorsResponse, error) {
	return r.client.TikTokAdsSearchCreators(ctx, request)
}

// TikTokAdsGetPopularTrendsRequest is the request for GET /api/v1/tiktok/ads/get_popular_trends.
type TikTokAdsGetPopularTrendsRequest = TikTokAdsGetPopularTrendVideosRequest

// TikTokAdsGetPopularTrendsResponse is the response for GET /api/v1/tiktok/ads/get_popular_trends.
type TikTokAdsGetPopularTrendsResponse = TikTokAdsGetPopularTrendVideosResponse

// GetPopularTrends 获取流行趋势视频/Get popular trend videos
//
// GET /api/v1/tiktok/ads/get_popular_trends
func (r TikTokAdsResource) GetPopularTrends(ctx context.Context, request TikTokAdsGetPopularTrendsRequest) (*TikTokAdsGetPopularTrendsResponse, error) {
	return r.client.TikTokAdsGetPopularTrendVideos(ctx, request)
}

// TikTokShopWebResource contains endpoints from the TikTok-Shop-Web-API tag.
type TikTokShopWebResource struct {
	client *Client
}

// TikTokShopWebFetchProductDetailRequest is the request for GET /api/v1/tiktok/shop/web/fetch_product_detail.
type TikTokShopWebFetchProductDetailRequest = TikTokShopWebGetProductDetailV1Request

// TikTokShopWebFetchProductDetailResponse is the response for GET /api/v1/tiktok/shop/web/fetch_product_detail.
type TikTokShopWebFetchProductDetailResponse = TikTokShopWebGetProductDetailV1Response

// FetchProductDetail 获取商品详情V1(桌面端-数据完整)/Get product detail V1(Full data)
//
// GET /api/v1/tiktok/shop/web/fetch_product_detail
func (r TikTokShopWebResource) FetchProductDetail(ctx context.Context, request TikTokShopWebFetchProductDetailRequest) (*TikTokShopWebFetchProductDetailResponse, error) {
	return r.client.TikTokShopWebGetProductDetailV1(ctx, request)
}

// TikTokShopWebFetchProductDetailV2Request is the request for GET /api/v1/tiktok/shop/web/fetch_product_detail_v2.
type TikTokShopWebFetchProductDetailV2Request = TikTokShopWebGetProductDetailV2Request

// TikTokShopWebFetchProductDetailV2Response is the response for GET /api/v1/tiktok/shop/web/fetch_product_detail_v2.
type TikTokShopWebFetchProductDetailV2Response = TikTokShopWebGetProductDetailV2Response

// FetchProductDetailV2 获取商品详情V2(移动端-数据少)/Get product detail V2 (Less Data)
//
// GET /api/v1/tiktok/shop/web/fetch_product_detail_v2
func (r TikTokShopWebResource) FetchProductDetailV2(ctx context.Context, request TikTokShopWebFetchProductDetailV2Request) (*TikTokShopWebFetchProductDetailV2Response, error) {
	return r.client.TikTokShopWebGetProductDetailV2(ctx, request)
}

// TikTokShopWebFetchProductDetailV3Request is the request for GET /api/v1/tiktok/shop/web/fetch_product_detail_v3.
type TikTokShopWebFetchProductDetailV3Request = TikTokShopWebGetProductDetailV3Request

// TikTokShopWebFetchProductDetailV3Response is the response for GET /api/v1/tiktok/shop/web/fetch_product_detail_v3.
type TikTokShopWebFetchProductDetailV3Response = TikTokShopWebGetProductDetailV3Response

// FetchProductDetailV3 获取商品详情V3(移动端-数据完整)/Get product detail V3 (Full Data)
//
// GET /api/v1/tiktok/shop/web/fetch_product_detail_v3
func (r TikTokShopWebResource) FetchProductDetailV3(ctx context.Context, request TikTokShopWebFetchProductDetailV3Request) (*TikTokShopWebFetchProductDetailV3Response, error) {
	return r.client.TikTokShopWebGetProductDetailV3(ctx, request)
}

// TikTokShopWebFetchProductReviewsV2Request is the request for GET /api/v1/tiktok/shop/web/fetch_product_reviews_v2.
type TikTokShopWebFetchProductReviewsV2Request = TikTokShopWebGetProductReviewsV2Request

// TikTokShopWebFetchProductReviewsV2Response is the response for GET /api/v1/tiktok/shop/web/fetch_product_reviews_v2.
type TikTokShopWebFetchProductReviewsV2Response = TikTokShopWebGetProductReviewsV2Response

// FetchProductReviewsV2 获取商品评论V2/Get product reviews V2
//
// GET /api/v1/tiktok/shop/web/fetch_product_reviews_v2
func (r TikTokShopWebResource) FetchProductReviewsV2(ctx context.Context, request TikTokShopWebFetchProductReviewsV2Request) (*TikTokShopWebFetchProductReviewsV2Response, error) {
	return r.client.TikTokShopWebGetProductReviewsV2(ctx, request)
}

// TikTokShopWebFetchSellerProductsListRequest is the request for GET /api/v1/tiktok/shop/web/fetch_seller_products_list.
type TikTokShopWebFetchSellerProductsListRequest = TikTokShopWebGetSellerProductsListV1Request

// TikTokShopWebFetchSellerProductsListResponse is the response for GET /api/v1/tiktok/shop/web/fetch_seller_products_list.
type TikTokShopWebFetchSellerProductsListResponse = TikTokShopWebGetSellerProductsListV1Response

// FetchSellerProductsList 获取商家商品列表V1/Get seller products list V1
//
// GET /api/v1/tiktok/shop/web/fetch_seller_products_list
func (r TikTokShopWebResource) FetchSellerProductsList(ctx context.Context, request TikTokShopWebFetchSellerProductsListRequest) (*TikTokShopWebFetchSellerProductsListResponse, error) {
	return r.client.TikTokShopWebGetSellerProductsListV1(ctx, request)
}

// TikTokShopWebFetchSellerProductsListV2Request is the request for GET /api/v1/tiktok/shop/web/fetch_seller_products_list_v2.
type TikTokShopWebFetchSellerProductsListV2Request = TikTokShopWebGetSellerProductsListV2Request

// TikTokShopWebFetchSellerProductsListV2Response is the response for GET /api/v1/tiktok/shop/web/fetch_seller_products_list_v2.
type TikTokShopWebFetchSellerProductsListV2Response = TikTokShopWebGetSellerProductsListV2Response

// FetchSellerProductsListV2 获取商家商品列表V2(移动端)/Get seller products list V2 (Mobile)
//
// GET /api/v1/tiktok/shop/web/fetch_seller_products_list_v2
func (r TikTokShopWebResource) FetchSellerProductsListV2(ctx context.Context, request TikTokShopWebFetchSellerProductsListV2Request) (*TikTokShopWebFetchSellerProductsListV2Response, error) {
	return r.client.TikTokShopWebGetSellerProductsListV2(ctx, request)
}

// TikTokShopWebFetchSearchWordSuggestionRequest is the request for GET /api/v1/tiktok/shop/web/fetch_search_word_suggestion.
type TikTokShopWebFetchSearchWordSuggestionRequest = TikTokShopWebGetSearchKeywordSuggestionsV1Request

// TikTokShopWebFetchSearchWordSuggestionResponse is the response for GET /api/v1/tiktok/shop/web/fetch_search_word_suggestion.
type TikTokShopWebFetchSearchWordSuggestionResponse = TikTokShopWebGetSearchKeywordSuggestionsV1Response

// FetchSearchWordSuggestion 获取搜索关键词建议V1/Get search keyword suggestions V1
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/tiktok/shop/web/fetch_search_word_suggestion
func (r TikTokShopWebResource) FetchSearchWordSuggestion(ctx context.Context, request TikTokShopWebFetchSearchWordSuggestionRequest) (*TikTokShopWebFetchSearchWordSuggestionResponse, error) {
	return r.client.TikTokShopWebGetSearchKeywordSuggestionsV1(ctx, request)
}

// TikTokShopWebFetchSearchWordSuggestionV2Request is the request for GET /api/v1/tiktok/shop/web/fetch_search_word_suggestion_v2.
type TikTokShopWebFetchSearchWordSuggestionV2Request = TikTokShopWebGetSearchKeywordSuggestionsV2Request

// TikTokShopWebFetchSearchWordSuggestionV2Response is the response for GET /api/v1/tiktok/shop/web/fetch_search_word_suggestion_v2.
type TikTokShopWebFetchSearchWordSuggestionV2Response = TikTokShopWebGetSearchKeywordSuggestionsV2Response

// FetchSearchWordSuggestionV2 获取搜索关键词建议V2(移动端)/Get search keyword suggestions V2 (Mobile)
//
// GET /api/v1/tiktok/shop/web/fetch_search_word_suggestion_v2
func (r TikTokShopWebResource) FetchSearchWordSuggestionV2(ctx context.Context, request TikTokShopWebFetchSearchWordSuggestionV2Request) (*TikTokShopWebFetchSearchWordSuggestionV2Response, error) {
	return r.client.TikTokShopWebGetSearchKeywordSuggestionsV2(ctx, request)
}

// TikTokShopWebFetchSearchProductsListRequest is the request for GET /api/v1/tiktok/shop/web/fetch_search_products_list.
type TikTokShopWebFetchSearchProductsListRequest = TikTokShopWebSearchProductsListV1Request

// TikTokShopWebFetchSearchProductsListResponse is the response for GET /api/v1/tiktok/shop/web/fetch_search_products_list.
type TikTokShopWebFetchSearchProductsListResponse = TikTokShopWebSearchProductsListV1Response

// FetchSearchProductsList 搜索商品列表V1/Search products list V1
//
// GET /api/v1/tiktok/shop/web/fetch_search_products_list
func (r TikTokShopWebResource) FetchSearchProductsList(ctx context.Context, request TikTokShopWebFetchSearchProductsListRequest) (*TikTokShopWebFetchSearchProductsListResponse, error) {
	return r.client.TikTokShopWebSearchProductsListV1(ctx, request)
}

// TikTokShopWebFetchSearchProductsListV2Request is the request for GET /api/v1/tiktok/shop/web/fetch_search_products_list_v2.
type TikTokShopWebFetchSearchProductsListV2Request = TikTokShopWebSearchProductsListV2Request

// TikTokShopWebFetchSearchProductsListV2Response is the response for GET /api/v1/tiktok/shop/web/fetch_search_products_list_v2.
type TikTokShopWebFetchSearchProductsListV2Response = TikTokShopWebSearchProductsListV2Response

// FetchSearchProductsListV2 搜索商品列表V2(移动端)/Search products list V2 (Mobile)
//
// GET /api/v1/tiktok/shop/web/fetch_search_products_list_v2
func (r TikTokShopWebResource) FetchSearchProductsListV2(ctx context.Context, request TikTokShopWebFetchSearchProductsListV2Request) (*TikTokShopWebFetchSearchProductsListV2Response, error) {
	return r.client.TikTokShopWebSearchProductsListV2(ctx, request)
}

// TikTokShopWebFetchProductsCategoryListRequest is the request for GET /api/v1/tiktok/shop/web/fetch_products_category_list.
type TikTokShopWebFetchProductsCategoryListRequest = TikTokShopWebGetProductCategoryListRequest

// TikTokShopWebFetchProductsCategoryListResponse is the response for GET /api/v1/tiktok/shop/web/fetch_products_category_list.
type TikTokShopWebFetchProductsCategoryListResponse = TikTokShopWebGetProductCategoryListResponse

// FetchProductsCategoryList 获取商品分类列表/Get product category list
//
// GET /api/v1/tiktok/shop/web/fetch_products_category_list
func (r TikTokShopWebResource) FetchProductsCategoryList(ctx context.Context, request TikTokShopWebFetchProductsCategoryListRequest) (*TikTokShopWebFetchProductsCategoryListResponse, error) {
	return r.client.TikTokShopWebGetProductCategoryList(ctx, request)
}

// TikTokShopWebFetchProductsByCategoryIDRequest is the request for GET /api/v1/tiktok/shop/web/fetch_products_by_category_id.
type TikTokShopWebFetchProductsByCategoryIDRequest = TikTokShopWebGetProductsByCategoryIDRequest

// TikTokShopWebFetchProductsByCategoryIDResponse is the response for GET /api/v1/tiktok/shop/web/fetch_products_by_category_id.
type TikTokShopWebFetchProductsByCategoryIDResponse = TikTokShopWebGetProductsByCategoryIDResponse

// FetchProductsByCategoryID 根据分类ID获取商品列表/Get products by category ID
//
// GET /api/v1/tiktok/shop/web/fetch_products_by_category_id
func (r TikTokShopWebResource) FetchProductsByCategoryID(ctx context.Context, request TikTokShopWebFetchProductsByCategoryIDRequest) (*TikTokShopWebFetchProductsByCategoryIDResponse, error) {
	return r.client.TikTokShopWebGetProductsByCategoryID(ctx, request)
}

// TikTokShopWebFetchHotSellingProductsListRequest is the request for GET /api/v1/tiktok/shop/web/fetch_hot_selling_products_list.
type TikTokShopWebFetchHotSellingProductsListRequest = TikTokShopWebGetHotSellingProductsListRequest

// TikTokShopWebFetchHotSellingProductsListResponse is the response for GET /api/v1/tiktok/shop/web/fetch_hot_selling_products_list.
type TikTokShopWebFetchHotSellingProductsListResponse = TikTokShopWebGetHotSellingProductsListResponse

// FetchHotSellingProductsList 获取热卖商品列表/Get hot selling products list
//
// GET /api/v1/tiktok/shop/web/fetch_hot_selling_products_list
func (r TikTokShopWebResource) FetchHotSellingProductsList(ctx context.Context, request TikTokShopWebFetchHotSellingProductsListRequest) (*TikTokShopWebFetchHotSellingProductsListResponse, error) {
	return r.client.TikTokShopWebGetHotSellingProductsList(ctx, request)
}

// TikTokInteractionResource contains endpoints from the TikTok-Interaction-API tag.
type TikTokInteractionResource struct {
	client *Client
}

// TikTokInteractionApplyRequest is the request for GET /api/v1/tiktok/interaction/apply.
type TikTokInteractionApplyRequest = TikTokInteractionApplyForTikTokInteractionAPIPermissionRequest

// TikTokInteractionApplyResponse is the response for GET /api/v1/tiktok/interaction/apply.
type TikTokInteractionApplyResponse = TikTokInteractionApplyForTikTokInteractionAPIPermissionResponse

// Apply 申请使用TikTok交互API权限（Scope）/Apply for TikTok Interaction API permission (Scope)
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/tiktok/interaction/apply
func (r TikTokInteractionResource) Apply(ctx context.Context, request TikTokInteractionApplyRequest) (*TikTokInteractionApplyResponse, error) {
	return r.client.TikTokInteractionApplyForTikTokInteractionAPIPermission(ctx, request)
}

// PostComment 发送评论/Post comment
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// POST /api/v1/tiktok/interaction/post_comment
func (r TikTokInteractionResource) PostComment(ctx context.Context, request TikTokInteractionPostCommentRequest) (*TikTokInteractionPostCommentResponse, error) {
	return r.client.TikTokInteractionPostComment(ctx, request)
}

// TikTokInteractionReplyCommentRequest is the request for POST /api/v1/tiktok/interaction/reply_comment.
type TikTokInteractionReplyCommentRequest = TikTokInteractionReplyToCommentRequest

// TikTokInteractionReplyCommentResponse is the response for POST /api/v1/tiktok/interaction/reply_comment.
type TikTokInteractionReplyCommentResponse = TikTokInteractionReplyToCommentResponse

// ReplyComment 回复评论/Reply to comment
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// POST /api/v1/tiktok/interaction/reply_comment
func (r TikTokInteractionResource) ReplyComment(ctx context.Context, request TikTokInteractionReplyCommentRequest) (*TikTokInteractionReplyCommentResponse, error) {
	return r.client.TikTokInteractionReplyToComment(ctx, request)
}

// Like 点赞/Like
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// POST /api/v1/tiktok/interaction/like
func (r TikTokInteractionResource) Like(ctx context.Context, request TikTokInteractionLikeRequest) (*TikTokInteractionLikeResponse, error) {
	return r.client.TikTokInteractionLike(ctx, request)
}

// Follow 关注/Follow
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// POST /api/v1/tiktok/interaction/follow
func (r TikTokInteractionResource) Follow(ctx context.Context, request TikTokInteractionFollowRequest) (*TikTokInteractionFollowResponse, error) {
	return r.client.TikTokInteractionFollow(ctx, request)
}

// Collect 收藏/Collect
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// POST /api/v1/tiktok/interaction/collect
func (r TikTokInteractionResource) Collect(ctx context.Context, request TikTokInteractionCollectRequest) (*TikTokInteractionCollectResponse, error) {
	return r.client.TikTokInteractionCollect(ctx, request)
}

// Forward 转发/Forward
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// POST /api/v1/tiktok/interaction/forward
func (r TikTokInteractionResource) Forward(ctx context.Context, request TikTokInteractionForwardRequest) (*TikTokInteractionForwardResponse, error) {
	return r.client.TikTokInteractionForward(ctx, request)
}

// DouyinWebResource contains endpoints from the Douyin-Web-API tag.
type DouyinWebResource struct {
	client *Client
}

// DouyinWebFetchOneVideoRequest is the request for GET /api/v1/douyin/web/fetch_one_video.
type DouyinWebFetchOneVideoRequest = DouyinWebGetSingleVideoDataRequest

// DouyinWebFetchOneVideoResponse is the response for GET /api/v1/douyin/web/fetch_one_video.
type DouyinWebFetchOneVideoResponse = DouyinWebGetSingleVideoDataResponse

// FetchOneVideo 获取单个作品数据/Get single video data
//
// GET /api/v1/douyin/web/fetch_one_video
func (r DouyinWebResource) FetchOneVideo(ctx context.Context, request DouyinWebFetchOneVideoRequest) (*DouyinWebFetchOneVideoResponse, error) {
	return r.client.DouyinWebGetSingleVideoData(ctx, request)
}

// DouyinWebFetchOneVideoV2Request is the request for GET /api/v1/douyin/web/fetch_one_video_v2.
type DouyinWebFetchOneVideoV2Request = DouyinWebGetSingleVideoDataV2Request

// DouyinWebFetchOneVideoV2Response is the response for GET /api/v1/douyin/web/fetch_one_video_v2.
type DouyinWebFetchOneVideoV2Response = DouyinWebGetSingleVideoDataV2Response

// FetchOneVideoV2 获取单个作品数据 V2/Get single video data V2
//
// GET /api/v1/douyin/web/fetch_one_video_v2
func (r DouyinWebResource) FetchOneVideoV2(ctx context.Context, request DouyinWebFetchOneVideoV2Request) (*DouyinWebFetchOneVideoV2Response, error) {
	return r.client.DouyinWebGetSingleVideoDataV2(ctx, request)
}

// DouyinWebFetchOneVideoByShareURLRequest is the request for GET /api/v1/douyin/web/fetch_one_video_by_share_url.
type DouyinWebFetchOneVideoByShareURLRequest = DouyinWebGetSingleVideoDataBySharingLinkRequest

// DouyinWebFetchOneVideoByShareURLResponse is the response for GET /api/v1/douyin/web/fetch_one_video_by_share_url.
type DouyinWebFetchOneVideoByShareURLResponse = DouyinWebGetSingleVideoDataBySharingLinkResponse

// FetchOneVideoByShareURL 根据分享链接获取单个作品数据/Get single video data by sharing link
//
// GET /api/v1/douyin/web/fetch_one_video_by_share_url
func (r DouyinWebResource) FetchOneVideoByShareURL(ctx context.Context, request DouyinWebFetchOneVideoByShareURLRequest) (*DouyinWebFetchOneVideoByShareURLResponse, error) {
	return r.client.DouyinWebGetSingleVideoDataBySharingLink(ctx, request)
}

// DouyinWebFetchVideoHighQualityPlayURLRequest is the request for GET /api/v1/douyin/web/fetch_video_high_quality_play_url.
type DouyinWebFetchVideoHighQualityPlayURLRequest = DouyinWebGetTheHighestQualityPlayURLOfTheVideoRequest

// DouyinWebFetchVideoHighQualityPlayURLResponse is the response for GET /api/v1/douyin/web/fetch_video_high_quality_play_url.
type DouyinWebFetchVideoHighQualityPlayURLResponse = DouyinWebGetTheHighestQualityPlayURLOfTheVideoResponse

// FetchVideoHighQualityPlayURL 获取视频的最高画质播放链接/Get the highest quality play URL of the video
//
// GET /api/v1/douyin/web/fetch_video_high_quality_play_url
func (r DouyinWebResource) FetchVideoHighQualityPlayURL(ctx context.Context, request DouyinWebFetchVideoHighQualityPlayURLRequest) (*DouyinWebFetchVideoHighQualityPlayURLResponse, error) {
	return r.client.DouyinWebGetTheHighestQualityPlayURLOfTheVideo(ctx, request)
}

// DouyinWebFetchMultiVideoHighQualityPlayURLRequest is the request for POST /api/v1/douyin/web/fetch_multi_video_high_quality_play_url.
type DouyinWebFetchMultiVideoHighQualityPlayURLRequest = DouyinWebBatchGetTheHighestQualityPlayURLOfVideosRequest

// DouyinWebFetchMultiVideoHighQualityPlayURLResponse is the response for POST /api/v1/douyin/web/fetch_multi_video_high_quality_play_url.
type DouyinWebFetchMultiVideoHighQualityPlayURLResponse = DouyinWebBatchGetTheHighestQualityPlayURLOfVideosResponse

// FetchMultiVideoHighQualityPlayURL 批量获取视频的最高画质播放链接/Batch get the highest quality play URL of videos
//
// POST /api/v1/douyin/web/fetch_multi_video_high_quality_play_url
func (r DouyinWebResource) FetchMultiVideoHighQualityPlayURL(ctx context.Context, request DouyinWebFetchMultiVideoHighQualityPlayURLRequest) (*DouyinWebFetchMultiVideoHighQualityPlayURLResponse, error) {
	return r.client.DouyinWebBatchGetTheHighestQualityPlayURLOfVideos(ctx, request)
}

// DouyinWebFetchMultiVideoRequest is the request for POST /api/v1/douyin/web/fetch_multi_video.
type DouyinWebFetchMultiVideoRequest = DouyinWebBatchGetVideoInformationRequest

// DouyinWebFetchMultiVideoResponse is the response for POST /api/v1/douyin/web/fetch_multi_video.
type DouyinWebFetchMultiVideoResponse = DouyinWebBatchGetVideoInformationResponse

// FetchMultiVideo 批量获取视频信息/Batch Get Video Information
//
// POST /api/v1/douyin/web/fetch_multi_video
func (r DouyinWebResource) FetchMultiVideo(ctx context.Context, request DouyinWebFetchMultiVideoRequest) (*DouyinWebFetchMultiVideoResponse, error) {
	return r.client.DouyinWebBatchGetVideoInformation(ctx, request)
}

// DouyinWebFetchOneVideoDanmakuRequest is the request for GET /api/v1/douyin/web/fetch_one_video_danmaku.
type DouyinWebFetchOneVideoDanmakuRequest = DouyinWebGetSingleVideoDanmakuDataRequest

// DouyinWebFetchOneVideoDanmakuResponse is the response for GET /api/v1/douyin/web/fetch_one_video_danmaku.
type DouyinWebFetchOneVideoDanmakuResponse = DouyinWebGetSingleVideoDanmakuDataResponse

// FetchOneVideoDanmaku 获取单个作品视频弹幕数据/Get single video danmaku data
//
// GET /api/v1/douyin/web/fetch_one_video_danmaku
func (r DouyinWebResource) FetchOneVideoDanmaku(ctx context.Context, request DouyinWebFetchOneVideoDanmakuRequest) (*DouyinWebFetchOneVideoDanmakuResponse, error) {
	return r.client.DouyinWebGetSingleVideoDanmakuData(ctx, request)
}

// DouyinWebFetchHomeFeedRequest is the request for GET /api/v1/douyin/web/fetch_home_feed.
type DouyinWebFetchHomeFeedRequest = DouyinWebGetHomeFeedDataRequest

// DouyinWebFetchHomeFeedResponse is the response for GET /api/v1/douyin/web/fetch_home_feed.
type DouyinWebFetchHomeFeedResponse = DouyinWebGetHomeFeedDataResponse

// FetchHomeFeed 获取首页推荐数据/Get home feed data
//
// GET /api/v1/douyin/web/fetch_home_feed
func (r DouyinWebResource) FetchHomeFeed(ctx context.Context, request DouyinWebFetchHomeFeedRequest) (*DouyinWebFetchHomeFeedResponse, error) {
	return r.client.DouyinWebGetHomeFeedData(ctx, request)
}

// DouyinWebFetchRelatedPostsRequest is the request for GET /api/v1/douyin/web/fetch_related_posts.
type DouyinWebFetchRelatedPostsRequest = DouyinWebGetRelatedPostsRecommendationDataRequest

// DouyinWebFetchRelatedPostsResponse is the response for GET /api/v1/douyin/web/fetch_related_posts.
type DouyinWebFetchRelatedPostsResponse = DouyinWebGetRelatedPostsRecommendationDataResponse

// FetchRelatedPosts 获取相关作品推荐数据/Get related posts recommendation data
//
// GET /api/v1/douyin/web/fetch_related_posts
func (r DouyinWebResource) FetchRelatedPosts(ctx context.Context, request DouyinWebFetchRelatedPostsRequest) (*DouyinWebFetchRelatedPostsResponse, error) {
	return r.client.DouyinWebGetRelatedPostsRecommendationData(ctx, request)
}

// DouyinWebFetchUserPostVideosRequest is the request for GET /api/v1/douyin/web/fetch_user_post_videos.
type DouyinWebFetchUserPostVideosRequest = DouyinWebGetUserHomepageVideoDataRequest

// DouyinWebFetchUserPostVideosResponse is the response for GET /api/v1/douyin/web/fetch_user_post_videos.
type DouyinWebFetchUserPostVideosResponse = DouyinWebGetUserHomepageVideoDataResponse

// FetchUserPostVideos 获取用户主页作品数据/Get user homepage video data
//
// GET /api/v1/douyin/web/fetch_user_post_videos
func (r DouyinWebResource) FetchUserPostVideos(ctx context.Context, request DouyinWebFetchUserPostVideosRequest) (*DouyinWebFetchUserPostVideosResponse, error) {
	return r.client.DouyinWebGetUserHomepageVideoData(ctx, request)
}

// DouyinWebFetchUserLikeVideosRequest is the request for POST /api/v1/douyin/web/fetch_user_like_videos.
type DouyinWebFetchUserLikeVideosRequest = DouyinWebGetUserLikeVideoDataRequest

// DouyinWebFetchUserLikeVideosResponse is the response for POST /api/v1/douyin/web/fetch_user_like_videos.
type DouyinWebFetchUserLikeVideosResponse = DouyinWebGetUserLikeVideoDataResponse

// FetchUserLikeVideos 获取用户喜欢作品数据/Get user like video data
//
// POST /api/v1/douyin/web/fetch_user_like_videos
func (r DouyinWebResource) FetchUserLikeVideos(ctx context.Context, request DouyinWebFetchUserLikeVideosRequest) (*DouyinWebFetchUserLikeVideosResponse, error) {
	return r.client.DouyinWebGetUserLikeVideoData(ctx, request)
}

// DouyinWebFetchUserCollectionVideosRequest is the request for POST /api/v1/douyin/web/fetch_user_collection_videos.
type DouyinWebFetchUserCollectionVideosRequest = DouyinWebGetUserCollectionVideoDataRequest

// DouyinWebFetchUserCollectionVideosResponse is the response for POST /api/v1/douyin/web/fetch_user_collection_videos.
type DouyinWebFetchUserCollectionVideosResponse = DouyinWebGetUserCollectionVideoDataResponse

// FetchUserCollectionVideos 获取用户收藏作品数据/Get user collection video data
//
// POST /api/v1/douyin/web/fetch_user_collection_videos
func (r DouyinWebResource) FetchUserCollectionVideos(ctx context.Context, request DouyinWebFetchUserCollectionVideosRequest) (*DouyinWebFetchUserCollectionVideosResponse, error) {
	return r.client.DouyinWebGetUserCollectionVideoData(ctx, request)
}

// DouyinWebFetchUserCollectsRequest is the request for POST /api/v1/douyin/web/fetch_user_collects.
type DouyinWebFetchUserCollectsRequest = DouyinWebGetUserCollectionRequest

// DouyinWebFetchUserCollectsResponse is the response for POST /api/v1/douyin/web/fetch_user_collects.
type DouyinWebFetchUserCollectsResponse = DouyinWebGetUserCollectionResponse

// FetchUserCollects 获取用户收藏夹/Get user collection
//
// POST /api/v1/douyin/web/fetch_user_collects
func (r DouyinWebResource) FetchUserCollects(ctx context.Context, request DouyinWebFetchUserCollectsRequest) (*DouyinWebFetchUserCollectsResponse, error) {
	return r.client.DouyinWebGetUserCollection(ctx, request)
}

// DouyinWebFetchUserCollectsVideosRequest is the request for GET /api/v1/douyin/web/fetch_user_collects_videos.
type DouyinWebFetchUserCollectsVideosRequest = DouyinWebGetUserCollectionDataRequest

// DouyinWebFetchUserCollectsVideosResponse is the response for GET /api/v1/douyin/web/fetch_user_collects_videos.
type DouyinWebFetchUserCollectsVideosResponse = DouyinWebGetUserCollectionDataResponse

// FetchUserCollectsVideos 获取用户收藏夹数据/Get user collection data
//
// GET /api/v1/douyin/web/fetch_user_collects_videos
func (r DouyinWebResource) FetchUserCollectsVideos(ctx context.Context, request DouyinWebFetchUserCollectsVideosRequest) (*DouyinWebFetchUserCollectsVideosResponse, error) {
	return r.client.DouyinWebGetUserCollectionData(ctx, request)
}

// DouyinWebFetchUserMixVideosRequest is the request for GET /api/v1/douyin/web/fetch_user_mix_videos.
type DouyinWebFetchUserMixVideosRequest = DouyinWebGetUserMixVideoDataRequest

// DouyinWebFetchUserMixVideosResponse is the response for GET /api/v1/douyin/web/fetch_user_mix_videos.
type DouyinWebFetchUserMixVideosResponse = DouyinWebGetUserMixVideoDataResponse

// FetchUserMixVideos 获取用户合辑作品数据/Get user mix video data
//
// GET /api/v1/douyin/web/fetch_user_mix_videos
func (r DouyinWebResource) FetchUserMixVideos(ctx context.Context, request DouyinWebFetchUserMixVideosRequest) (*DouyinWebFetchUserMixVideosResponse, error) {
	return r.client.DouyinWebGetUserMixVideoData(ctx, request)
}

// DouyinWebFetchUserLiveVideosRequest is the request for GET /api/v1/douyin/web/fetch_user_live_videos.
type DouyinWebFetchUserLiveVideosRequest = DouyinWebGetUserLiveVideoDataRequest

// DouyinWebFetchUserLiveVideosResponse is the response for GET /api/v1/douyin/web/fetch_user_live_videos.
type DouyinWebFetchUserLiveVideosResponse = DouyinWebGetUserLiveVideoDataResponse

// FetchUserLiveVideos 获取用户直播流数据/Get user live video data
//
// GET /api/v1/douyin/web/fetch_user_live_videos
func (r DouyinWebResource) FetchUserLiveVideos(ctx context.Context, request DouyinWebFetchUserLiveVideosRequest) (*DouyinWebFetchUserLiveVideosResponse, error) {
	return r.client.DouyinWebGetUserLiveVideoData(ctx, request)
}

// DouyinWebFetchUserLiveVideosBySecUIDRequest is the request for GET /api/v1/douyin/web/fetch_user_live_videos_by_sec_uid.
type DouyinWebFetchUserLiveVideosBySecUIDRequest = DouyinWebGetLiveVideoDataOfSpecifiedUserBySecUIDRequest

// DouyinWebFetchUserLiveVideosBySecUIDResponse is the response for GET /api/v1/douyin/web/fetch_user_live_videos_by_sec_uid.
type DouyinWebFetchUserLiveVideosBySecUIDResponse = DouyinWebGetLiveVideoDataOfSpecifiedUserBySecUIDResponse

// FetchUserLiveVideosBySecUID 通过sec_uid获取指定用户的直播流数据/Get live video data of specified user by sec_uid
//
// GET /api/v1/douyin/web/fetch_user_live_videos_by_sec_uid
func (r DouyinWebResource) FetchUserLiveVideosBySecUID(ctx context.Context, request DouyinWebFetchUserLiveVideosBySecUIDRequest) (*DouyinWebFetchUserLiveVideosBySecUIDResponse, error) {
	return r.client.DouyinWebGetLiveVideoDataOfSpecifiedUserBySecUID(ctx, request)
}

// DouyinWebFetchUserLiveVideosByRoomIDRequest is the request for GET /api/v1/douyin/web/fetch_user_live_videos_by_room_id.
type DouyinWebFetchUserLiveVideosByRoomIDRequest = DouyinWebGetLiveVideoDataOfSpecifiedUserByRoomIDV1Request

// DouyinWebFetchUserLiveVideosByRoomIDResponse is the response for GET /api/v1/douyin/web/fetch_user_live_videos_by_room_id.
type DouyinWebFetchUserLiveVideosByRoomIDResponse = DouyinWebGetLiveVideoDataOfSpecifiedUserByRoomIDV1Response

// FetchUserLiveVideosByRoomID 通过room_id获取指定用户的直播流数据 V1/Get live video data of specified user by room_id V1
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/douyin/web/fetch_user_live_videos_by_room_id
func (r DouyinWebResource) FetchUserLiveVideosByRoomID(ctx context.Context, request DouyinWebFetchUserLiveVideosByRoomIDRequest) (*DouyinWebFetchUserLiveVideosByRoomIDResponse, error) {
	return r.client.DouyinWebGetLiveVideoDataOfSpecifiedUserByRoomIDV1(ctx, request)
}

// DouyinWebFetchUserLiveVideosByRoomIDV2Request is the request for GET /api/v1/douyin/web/fetch_user_live_videos_by_room_id_v2.
type DouyinWebFetchUserLiveVideosByRoomIDV2Request = DouyinWebGetsTheLiveStreamDataOfTheSpecifiedUserByRoomIDV2Request

// DouyinWebFetchUserLiveVideosByRoomIDV2Response is the response for GET /api/v1/douyin/web/fetch_user_live_videos_by_room_id_v2.
type DouyinWebFetchUserLiveVideosByRoomIDV2Response = DouyinWebGetsTheLiveStreamDataOfTheSpecifiedUserByRoomIDV2Response

// FetchUserLiveVideosByRoomIDV2 通过room_id获取指定用户的直播流数据 V2/Gets the live stream data of the specified user by room_id V2
//
// GET /api/v1/douyin/web/fetch_user_live_videos_by_room_id_v2
func (r DouyinWebResource) FetchUserLiveVideosByRoomIDV2(ctx context.Context, request DouyinWebFetchUserLiveVideosByRoomIDV2Request) (*DouyinWebFetchUserLiveVideosByRoomIDV2Response, error) {
	return r.client.DouyinWebGetsTheLiveStreamDataOfTheSpecifiedUserByRoomIDV2(ctx, request)
}

// DouyinWebFetchLiveGiftRankingRequest is the request for GET /api/v1/douyin/web/fetch_live_gift_ranking.
type DouyinWebFetchLiveGiftRankingRequest = DouyinWebGetLiveRoomGiftUserRankingRequest

// DouyinWebFetchLiveGiftRankingResponse is the response for GET /api/v1/douyin/web/fetch_live_gift_ranking.
type DouyinWebFetchLiveGiftRankingResponse = DouyinWebGetLiveRoomGiftUserRankingResponse

// FetchLiveGiftRanking 获取直播间送礼用户排行榜/Get live room gift user ranking
//
// GET /api/v1/douyin/web/fetch_live_gift_ranking
func (r DouyinWebResource) FetchLiveGiftRanking(ctx context.Context, request DouyinWebFetchLiveGiftRankingRequest) (*DouyinWebFetchLiveGiftRankingResponse, error) {
	return r.client.DouyinWebGetLiveRoomGiftUserRanking(ctx, request)
}

// DouyinWebFetchLiveRoomProductResultRequest is the request for GET /api/v1/douyin/web/fetch_live_room_product_result.
type DouyinWebFetchLiveRoomProductResultRequest = DouyinWebDouyinLiveRoomProductInformationRequest

// DouyinWebFetchLiveRoomProductResultResponse is the response for GET /api/v1/douyin/web/fetch_live_room_product_result.
type DouyinWebFetchLiveRoomProductResultResponse = DouyinWebDouyinLiveRoomProductInformationResponse

// FetchLiveRoomProductResult 抖音直播间商品信息/Douyin live room product information
//
// GET /api/v1/douyin/web/fetch_live_room_product_result
func (r DouyinWebResource) FetchLiveRoomProductResult(ctx context.Context, request DouyinWebFetchLiveRoomProductResultRequest) (*DouyinWebFetchLiveRoomProductResultResponse, error) {
	return r.client.DouyinWebDouyinLiveRoomProductInformation(ctx, request)
}

// DouyinWebFetchProductSkuListRequest is the request for GET /api/v1/douyin/web/fetch_product_sku_list.
type DouyinWebFetchProductSkuListRequest = DouyinWebGetProductSkuListRequest

// DouyinWebFetchProductSkuListResponse is the response for GET /api/v1/douyin/web/fetch_product_sku_list.
type DouyinWebFetchProductSkuListResponse = DouyinWebGetProductSkuListResponse

// FetchProductSkuList 获取商品SKU列表/Get product SKU list
//
// GET /api/v1/douyin/web/fetch_product_sku_list
func (r DouyinWebResource) FetchProductSkuList(ctx context.Context, request DouyinWebFetchProductSkuListRequest) (*DouyinWebFetchProductSkuListResponse, error) {
	return r.client.DouyinWebGetProductSkuList(ctx, request)
}

// DouyinWebFetchProductCouponRequest is the request for GET /api/v1/douyin/web/fetch_product_coupon.
type DouyinWebFetchProductCouponRequest = DouyinWebGetProductCouponInformationRequest

// DouyinWebFetchProductCouponResponse is the response for GET /api/v1/douyin/web/fetch_product_coupon.
type DouyinWebFetchProductCouponResponse = DouyinWebGetProductCouponInformationResponse

// FetchProductCoupon 获取商品优惠券信息/Get product coupon information
//
// GET /api/v1/douyin/web/fetch_product_coupon
func (r DouyinWebResource) FetchProductCoupon(ctx context.Context, request DouyinWebFetchProductCouponRequest) (*DouyinWebFetchProductCouponResponse, error) {
	return r.client.DouyinWebGetProductCouponInformation(ctx, request)
}

// DouyinWebFetchProductReviewScoreRequest is the request for GET /api/v1/douyin/web/fetch_product_review_score.
type DouyinWebFetchProductReviewScoreRequest = DouyinWebGetProductReviewScoreRequest

// DouyinWebFetchProductReviewScoreResponse is the response for GET /api/v1/douyin/web/fetch_product_review_score.
type DouyinWebFetchProductReviewScoreResponse = DouyinWebGetProductReviewScoreResponse

// FetchProductReviewScore 获取商品评价评分/Get product review score
//
// GET /api/v1/douyin/web/fetch_product_review_score
func (r DouyinWebResource) FetchProductReviewScore(ctx context.Context, request DouyinWebFetchProductReviewScoreRequest) (*DouyinWebFetchProductReviewScoreResponse, error) {
	return r.client.DouyinWebGetProductReviewScore(ctx, request)
}

// DouyinWebFetchProductReviewListRequest is the request for GET /api/v1/douyin/web/fetch_product_review_list.
type DouyinWebFetchProductReviewListRequest = DouyinWebGetProductReviewListRequest

// DouyinWebFetchProductReviewListResponse is the response for GET /api/v1/douyin/web/fetch_product_review_list.
type DouyinWebFetchProductReviewListResponse = DouyinWebGetProductReviewListResponse

// FetchProductReviewList 获取商品评价列表/Get product review list
//
// GET /api/v1/douyin/web/fetch_product_review_list
func (r DouyinWebResource) FetchProductReviewList(ctx context.Context, request DouyinWebFetchProductReviewListRequest) (*DouyinWebFetchProductReviewListResponse, error) {
	return r.client.DouyinWebGetProductReviewList(ctx, request)
}

// DouyinWebFetchUserProfileByUIDRequest is the request for GET /api/v1/douyin/web/fetch_user_profile_by_uid.
type DouyinWebFetchUserProfileByUIDRequest = DouyinWebGetUserInformationByUIDRequest

// DouyinWebFetchUserProfileByUIDResponse is the response for GET /api/v1/douyin/web/fetch_user_profile_by_uid.
type DouyinWebFetchUserProfileByUIDResponse = DouyinWebGetUserInformationByUIDResponse

// FetchUserProfileByUID 使用UID获取用户信息/Get user information by UID
//
// GET /api/v1/douyin/web/fetch_user_profile_by_uid
func (r DouyinWebResource) FetchUserProfileByUID(ctx context.Context, request DouyinWebFetchUserProfileByUIDRequest) (*DouyinWebFetchUserProfileByUIDResponse, error) {
	return r.client.DouyinWebGetUserInformationByUID(ctx, request)
}

// DouyinWebFetchBatchUserProfileV1Request is the request for GET /api/v1/douyin/web/fetch_batch_user_profile_v1.
type DouyinWebFetchBatchUserProfileV1Request = DouyinWebGetBatchUserProfileRequest

// DouyinWebFetchBatchUserProfileV1Response is the response for GET /api/v1/douyin/web/fetch_batch_user_profile_v1.
type DouyinWebFetchBatchUserProfileV1Response = DouyinWebGetBatchUserProfileResponse

// FetchBatchUserProfileV1 获取批量用户信息(最多10个)/Get batch user profile (up to 10)
//
// GET /api/v1/douyin/web/fetch_batch_user_profile_v1
func (r DouyinWebResource) FetchBatchUserProfileV1(ctx context.Context, request DouyinWebFetchBatchUserProfileV1Request) (*DouyinWebFetchBatchUserProfileV1Response, error) {
	return r.client.DouyinWebGetBatchUserProfile(ctx, request)
}

// DouyinWebFetchBatchUserProfileV2Request is the request for GET /api/v1/douyin/web/fetch_batch_user_profile_v2.
type DouyinWebFetchBatchUserProfileV2Request = DouyinWebGetBatchUserProfileWebFetchBatchUserProfileV2Request

// DouyinWebFetchBatchUserProfileV2Response is the response for GET /api/v1/douyin/web/fetch_batch_user_profile_v2.
type DouyinWebFetchBatchUserProfileV2Response = DouyinWebGetBatchUserProfileWebFetchBatchUserProfileV2Response

// FetchBatchUserProfileV2 获取批量用户信息(最多50个)/Get batch user profile (up to 50)
//
// GET /api/v1/douyin/web/fetch_batch_user_profile_v2
func (r DouyinWebResource) FetchBatchUserProfileV2(ctx context.Context, request DouyinWebFetchBatchUserProfileV2Request) (*DouyinWebFetchBatchUserProfileV2Response, error) {
	return r.client.DouyinWebGetBatchUserProfileWebFetchBatchUserProfileV2(ctx, request)
}

// DouyinWebFetchUserLiveInfoByUIDRequest is the request for GET /api/v1/douyin/web/fetch_user_live_info_by_uid.
type DouyinWebFetchUserLiveInfoByUIDRequest = DouyinWebGetUserLiveInformationByUIDRequest

// DouyinWebFetchUserLiveInfoByUIDResponse is the response for GET /api/v1/douyin/web/fetch_user_live_info_by_uid.
type DouyinWebFetchUserLiveInfoByUIDResponse = DouyinWebGetUserLiveInformationByUIDResponse

// FetchUserLiveInfoByUID 使用UID获取用户开播信息/Get user live information by UID
//
// GET /api/v1/douyin/web/fetch_user_live_info_by_uid
func (r DouyinWebResource) FetchUserLiveInfoByUID(ctx context.Context, request DouyinWebFetchUserLiveInfoByUIDRequest) (*DouyinWebFetchUserLiveInfoByUIDResponse, error) {
	return r.client.DouyinWebGetUserLiveInformationByUID(ctx, request)
}

// DouyinWebFetchUserProfileByShortIDRequest is the request for GET /api/v1/douyin/web/fetch_user_profile_by_short_id.
type DouyinWebFetchUserProfileByShortIDRequest = DouyinWebGetUserInformationByShortIDRequest

// DouyinWebFetchUserProfileByShortIDResponse is the response for GET /api/v1/douyin/web/fetch_user_profile_by_short_id.
type DouyinWebFetchUserProfileByShortIDResponse = DouyinWebGetUserInformationByShortIDResponse

// FetchUserProfileByShortID 使用Short ID获取用户信息/Get user information by Short ID
//
// GET /api/v1/douyin/web/fetch_user_profile_by_short_id
func (r DouyinWebResource) FetchUserProfileByShortID(ctx context.Context, request DouyinWebFetchUserProfileByShortIDRequest) (*DouyinWebFetchUserProfileByShortIDResponse, error) {
	return r.client.DouyinWebGetUserInformationByShortID(ctx, request)
}

// DouyinWebHandlerShortenURLRequest is the request for GET /api/v1/douyin/web/handler_shorten_url.
type DouyinWebHandlerShortenURLRequest = DouyinWebValueRequest

// DouyinWebHandlerShortenURLResponse is the response for GET /api/v1/douyin/web/handler_shorten_url.
type DouyinWebHandlerShortenURLResponse = DouyinWebValueResponse

// HandlerShortenURL 生成短链接
//
// GET /api/v1/douyin/web/handler_shorten_url
func (r DouyinWebResource) HandlerShortenURL(ctx context.Context, request DouyinWebHandlerShortenURLRequest) (*DouyinWebHandlerShortenURLResponse, error) {
	return r.client.DouyinWebValue(ctx, request)
}

// DouyinWebHandlerUserProfileRequest is the request for GET /api/v1/douyin/web/handler_user_profile.
type DouyinWebHandlerUserProfileRequest = DouyinWebGetInformationOfSpecifiedUserBySecUserIDRequest

// DouyinWebHandlerUserProfileResponse is the response for GET /api/v1/douyin/web/handler_user_profile.
type DouyinWebHandlerUserProfileResponse = DouyinWebGetInformationOfSpecifiedUserBySecUserIDResponse

// HandlerUserProfile 使用sec_user_id获取指定用户的信息/Get information of specified user by sec_user_id
//
// GET /api/v1/douyin/web/handler_user_profile
func (r DouyinWebResource) HandlerUserProfile(ctx context.Context, request DouyinWebHandlerUserProfileRequest) (*DouyinWebHandlerUserProfileResponse, error) {
	return r.client.DouyinWebGetInformationOfSpecifiedUserBySecUserID(ctx, request)
}

// DouyinWebHandlerUserProfileV2Request is the request for GET /api/v1/douyin/web/handler_user_profile_v2.
type DouyinWebHandlerUserProfileV2Request = DouyinWebGetInformationOfSpecifiedUserByUniqueIDRequest

// DouyinWebHandlerUserProfileV2Response is the response for GET /api/v1/douyin/web/handler_user_profile_v2.
type DouyinWebHandlerUserProfileV2Response = DouyinWebGetInformationOfSpecifiedUserByUniqueIDResponse

// HandlerUserProfileV2 使用unique_id（抖音号）获取指定用户的信息/Get information of specified user by unique_id
//
// GET /api/v1/douyin/web/handler_user_profile_v2
func (r DouyinWebResource) HandlerUserProfileV2(ctx context.Context, request DouyinWebHandlerUserProfileV2Request) (*DouyinWebHandlerUserProfileV2Response, error) {
	return r.client.DouyinWebGetInformationOfSpecifiedUserByUniqueID(ctx, request)
}

// DouyinWebEncryptUIDToSecUserIDRequest is the request for GET /api/v1/douyin/web/encrypt_uid_to_sec_user_id.
type DouyinWebEncryptUIDToSecUserIDRequest = DouyinWebEncryptUserUIDToSecUserIDRequest

// DouyinWebEncryptUIDToSecUserIDResponse is the response for GET /api/v1/douyin/web/encrypt_uid_to_sec_user_id.
type DouyinWebEncryptUIDToSecUserIDResponse = DouyinWebEncryptUserUIDToSecUserIDResponse

// EncryptUIDToSecUserID 加密用户uid到sec_user_id/Encrypt user uid to sec_user_id
//
// GET /api/v1/douyin/web/encrypt_uid_to_sec_user_id
func (r DouyinWebResource) EncryptUIDToSecUserID(ctx context.Context, request DouyinWebEncryptUIDToSecUserIDRequest) (*DouyinWebEncryptUIDToSecUserIDResponse, error) {
	return r.client.DouyinWebEncryptUserUIDToSecUserID(ctx, request)
}

// DouyinWebHandlerUserProfileV3Request is the request for GET /api/v1/douyin/web/handler_user_profile_v3.
type DouyinWebHandlerUserProfileV3Request = DouyinWebGetInformationOfSpecifiedUserByUIDRequest

// DouyinWebHandlerUserProfileV3Response is the response for GET /api/v1/douyin/web/handler_user_profile_v3.
type DouyinWebHandlerUserProfileV3Response = DouyinWebGetInformationOfSpecifiedUserByUIDResponse

// HandlerUserProfileV3 根据抖音uid获取指定用户的信息/Get information of specified user by uid
//
// GET /api/v1/douyin/web/handler_user_profile_v3
func (r DouyinWebResource) HandlerUserProfileV3(ctx context.Context, request DouyinWebHandlerUserProfileV3Request) (*DouyinWebHandlerUserProfileV3Response, error) {
	return r.client.DouyinWebGetInformationOfSpecifiedUserByUID(ctx, request)
}

// DouyinWebHandlerUserProfileV4Request is the request for GET /api/v1/douyin/web/handler_user_profile_v4.
type DouyinWebHandlerUserProfileV4Request = DouyinWebGetInformationOfSpecifiedUserBySecUserIDWebHandlerUserProfileV4Request

// DouyinWebHandlerUserProfileV4Response is the response for GET /api/v1/douyin/web/handler_user_profile_v4.
type DouyinWebHandlerUserProfileV4Response = DouyinWebGetInformationOfSpecifiedUserBySecUserIDWebHandlerUserProfileV4Response

// HandlerUserProfileV4 根据sec_user_id获取指定用户的信息（性别，年龄，直播等级、牌子）/Get information of specified user by sec_user_id (gender, age, live level、brand)
//
// GET /api/v1/douyin/web/handler_user_profile_v4
func (r DouyinWebResource) HandlerUserProfileV4(ctx context.Context, request DouyinWebHandlerUserProfileV4Request) (*DouyinWebHandlerUserProfileV4Response, error) {
	return r.client.DouyinWebGetInformationOfSpecifiedUserBySecUserIDWebHandlerUserProfileV4(ctx, request)
}

// DouyinWebFetchUserFansListRequest is the request for GET /api/v1/douyin/web/fetch_user_fans_list.
type DouyinWebFetchUserFansListRequest = DouyinWebGetUserFansListRequest

// DouyinWebFetchUserFansListResponse is the response for GET /api/v1/douyin/web/fetch_user_fans_list.
type DouyinWebFetchUserFansListResponse = DouyinWebGetUserFansListResponse

// FetchUserFansList 获取用户粉丝列表/Get user fans list
//
// GET /api/v1/douyin/web/fetch_user_fans_list
func (r DouyinWebResource) FetchUserFansList(ctx context.Context, request DouyinWebFetchUserFansListRequest) (*DouyinWebFetchUserFansListResponse, error) {
	return r.client.DouyinWebGetUserFansList(ctx, request)
}

// DouyinWebFetchUserFollowingListRequest is the request for GET /api/v1/douyin/web/fetch_user_following_list.
type DouyinWebFetchUserFollowingListRequest = DouyinWebGetUserFollowingListRequest

// DouyinWebFetchUserFollowingListResponse is the response for GET /api/v1/douyin/web/fetch_user_following_list.
type DouyinWebFetchUserFollowingListResponse = DouyinWebGetUserFollowingListResponse

// FetchUserFollowingList 获取用户关注列表/Get user following list
//
// GET /api/v1/douyin/web/fetch_user_following_list
func (r DouyinWebResource) FetchUserFollowingList(ctx context.Context, request DouyinWebFetchUserFollowingListRequest) (*DouyinWebFetchUserFollowingListResponse, error) {
	return r.client.DouyinWebGetUserFollowingList(ctx, request)
}

// DouyinWebFetchVideoCommentsRequest is the request for GET /api/v1/douyin/web/fetch_video_comments.
type DouyinWebFetchVideoCommentsRequest = DouyinWebGetSingleVideoCommentsDataRequest

// DouyinWebFetchVideoCommentsResponse is the response for GET /api/v1/douyin/web/fetch_video_comments.
type DouyinWebFetchVideoCommentsResponse = DouyinWebGetSingleVideoCommentsDataResponse

// FetchVideoComments 获取单个视频评论数据/Get single video comments data
//
// GET /api/v1/douyin/web/fetch_video_comments
func (r DouyinWebResource) FetchVideoComments(ctx context.Context, request DouyinWebFetchVideoCommentsRequest) (*DouyinWebFetchVideoCommentsResponse, error) {
	return r.client.DouyinWebGetSingleVideoCommentsData(ctx, request)
}

// DouyinWebFetchVideoCommentRepliesRequest is the request for GET /api/v1/douyin/web/fetch_video_comment_replies.
type DouyinWebFetchVideoCommentRepliesRequest = DouyinWebGetCommentRepliesDataOfSpecifiedVideoRequest

// DouyinWebFetchVideoCommentRepliesResponse is the response for GET /api/v1/douyin/web/fetch_video_comment_replies.
type DouyinWebFetchVideoCommentRepliesResponse = DouyinWebGetCommentRepliesDataOfSpecifiedVideoResponse

// FetchVideoCommentReplies 获取指定视频的评论回复数据/Get comment replies data of specified video
//
// GET /api/v1/douyin/web/fetch_video_comment_replies
func (r DouyinWebResource) FetchVideoCommentReplies(ctx context.Context, request DouyinWebFetchVideoCommentRepliesRequest) (*DouyinWebFetchVideoCommentRepliesResponse, error) {
	return r.client.DouyinWebGetCommentRepliesDataOfSpecifiedVideo(ctx, request)
}

// DouyinWebFetchUserSearchResultV3Request is the request for GET /api/v1/douyin/web/fetch_user_search_result_v3.
type DouyinWebFetchUserSearchResultV3Request = DouyinWebGetUserSearchResultsOfSpecifiedKeywordsV3Request

// DouyinWebFetchUserSearchResultV3Response is the response for GET /api/v1/douyin/web/fetch_user_search_result_v3.
type DouyinWebFetchUserSearchResultV3Response = DouyinWebGetUserSearchResultsOfSpecifiedKeywordsV3Response

// FetchUserSearchResultV3 获取指定关键词的用户搜索结果 V3 (已弃用，替代接口请参考下方文档)/Get user search results of specified keywords V3 (deprecated, please refer to the following document for replacement interface)
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/douyin/web/fetch_user_search_result_v3
func (r DouyinWebResource) FetchUserSearchResultV3(ctx context.Context, request DouyinWebFetchUserSearchResultV3Request) (*DouyinWebFetchUserSearchResultV3Response, error) {
	return r.client.DouyinWebGetUserSearchResultsOfSpecifiedKeywordsV3(ctx, request)
}

// DouyinWebFetchChallengePostsRequest is the request for POST /api/v1/douyin/web/fetch_challenge_posts.
type DouyinWebFetchChallengePostsRequest = DouyinWebChallengePostsRequest

// DouyinWebFetchChallengePostsResponse is the response for POST /api/v1/douyin/web/fetch_challenge_posts.
type DouyinWebFetchChallengePostsResponse = DouyinWebChallengePostsResponse

// FetchChallengePosts 话题作品/Challenge Posts
//
// POST /api/v1/douyin/web/fetch_challenge_posts
func (r DouyinWebResource) FetchChallengePosts(ctx context.Context, request DouyinWebFetchChallengePostsRequest) (*DouyinWebFetchChallengePostsResponse, error) {
	return r.client.DouyinWebChallengePosts(ctx, request)
}

// DouyinWebFetchHotSearchResultResponse is the response for GET /api/v1/douyin/web/fetch_hot_search_result.
type DouyinWebFetchHotSearchResultResponse = DouyinWebGetDouyinHotSearchResultsResponse

// FetchHotSearchResult 获取抖音热榜数据/Get Douyin hot search results
//
// GET /api/v1/douyin/web/fetch_hot_search_result
func (r DouyinWebResource) FetchHotSearchResult(ctx context.Context) (*DouyinWebFetchHotSearchResultResponse, error) {
	return r.client.DouyinWebGetDouyinHotSearchResults(ctx)
}

// DouyinWebFetchVideoChannelResultRequest is the request for GET /api/v1/douyin/web/fetch_video_channel_result.
type DouyinWebFetchVideoChannelResultRequest = DouyinWebDouyinVideoChannelDataRequest

// DouyinWebFetchVideoChannelResultResponse is the response for GET /api/v1/douyin/web/fetch_video_channel_result.
type DouyinWebFetchVideoChannelResultResponse = DouyinWebDouyinVideoChannelDataResponse

// FetchVideoChannelResult 抖音视频频道数据/Douyin video channel data
//
// GET /api/v1/douyin/web/fetch_video_channel_result
func (r DouyinWebResource) FetchVideoChannelResult(ctx context.Context, request DouyinWebFetchVideoChannelResultRequest) (*DouyinWebFetchVideoChannelResultResponse, error) {
	return r.client.DouyinWebDouyinVideoChannelData(ctx, request)
}

// DouyinWebFetchDouyinWebGuestCookieRequest is the request for GET /api/v1/douyin/web/fetch_douyin_web_guest_cookie.
type DouyinWebFetchDouyinWebGuestCookieRequest = DouyinWebGetTheGuestCookieOfDouyinWebRequest

// DouyinWebFetchDouyinWebGuestCookieResponse is the response for GET /api/v1/douyin/web/fetch_douyin_web_guest_cookie.
type DouyinWebFetchDouyinWebGuestCookieResponse = DouyinWebGetTheGuestCookieOfDouyinWebResponse

// FetchDouyinWebGuestCookie 获取抖音Web的游客Cookie/Get the guest Cookie of Douyin Web
//
// GET /api/v1/douyin/web/fetch_douyin_web_guest_cookie
func (r DouyinWebResource) FetchDouyinWebGuestCookie(ctx context.Context, request DouyinWebFetchDouyinWebGuestCookieRequest) (*DouyinWebFetchDouyinWebGuestCookieResponse, error) {
	return r.client.DouyinWebGetTheGuestCookieOfDouyinWeb(ctx, request)
}

// GenerateRealMSToken 生成真实msToken/Generate real msToken
//
// GET /api/v1/douyin/web/generate_real_msToken
func (r DouyinWebResource) GenerateRealMSToken(ctx context.Context) (*DouyinWebGenerateRealMSTokenResponse, error) {
	return r.client.DouyinWebGenerateRealMSToken(ctx)
}

// GenerateTtwid 生成ttwid/Generate ttwid
//
// GET /api/v1/douyin/web/generate_ttwid
func (r DouyinWebResource) GenerateTtwid(ctx context.Context, request DouyinWebGenerateTtwidRequest) (*DouyinWebGenerateTtwidResponse, error) {
	return r.client.DouyinWebGenerateTtwid(ctx, request)
}

// DouyinWebFetchQueryUserRequest is the request for POST /api/v1/douyin/web/fetch_query_user.
type DouyinWebFetchQueryUserRequest = DouyinWebQueryDouyinUserBasicInformationRequest

// DouyinWebFetchQueryUserResponse is the response for POST /api/v1/douyin/web/fetch_query_user.
type DouyinWebFetchQueryUserResponse = DouyinWebQueryDouyinUserBasicInformationResponse

// FetchQueryUser 查询抖音用户基本信息/Query Douyin user basic information
//
// POST /api/v1/douyin/web/fetch_query_user
func (r DouyinWebResource) FetchQueryUser(ctx context.Context, request DouyinWebFetchQueryUserRequest) (*DouyinWebFetchQueryUserResponse, error) {
	return r.client.DouyinWebQueryDouyinUserBasicInformation(ctx, request)
}

// GenerateVerifyFp 生成verify_fp/Generate verify_fp
//
// GET /api/v1/douyin/web/generate_verify_fp
func (r DouyinWebResource) GenerateVerifyFp(ctx context.Context) (*DouyinWebGenerateVerifyFpResponse, error) {
	return r.client.DouyinWebGenerateVerifyFp(ctx)
}

// GenerateSVWebID 生成s_v_web_id/Generate s_v_web_id
//
// GET /api/v1/douyin/web/generate_s_v_web_id
func (r DouyinWebResource) GenerateSVWebID(ctx context.Context) (*DouyinWebGenerateSVWebIDResponse, error) {
	return r.client.DouyinWebGenerateSVWebID(ctx)
}

// DouyinWebGenerateWssXbSignatureRequest is the request for GET /api/v1/douyin/web/generate_wss_xb_signature.
type DouyinWebGenerateWssXbSignatureRequest = DouyinWebGenerateBarrageXbSignatureRequest

// DouyinWebGenerateWssXbSignatureResponse is the response for GET /api/v1/douyin/web/generate_wss_xb_signature.
type DouyinWebGenerateWssXbSignatureResponse = DouyinWebGenerateBarrageXbSignatureResponse

// GenerateWssXbSignature 生成弹幕xb签名/Generate barrage xb signature
//
// GET /api/v1/douyin/web/generate_wss_xb_signature
func (r DouyinWebResource) GenerateWssXbSignature(ctx context.Context, request DouyinWebGenerateWssXbSignatureRequest) (*DouyinWebGenerateWssXbSignatureResponse, error) {
	return r.client.DouyinWebGenerateBarrageXbSignature(ctx, request)
}

// DouyinWebGenerateXBogusRequest is the request for POST /api/v1/douyin/web/generate_x_bogus.
type DouyinWebGenerateXBogusRequest = DouyinWebGenerateXBogusParameterUsingAPIURLRequest

// DouyinWebGenerateXBogusResponse is the response for POST /api/v1/douyin/web/generate_x_bogus.
type DouyinWebGenerateXBogusResponse = DouyinWebGenerateXBogusParameterUsingAPIURLResponse

// GenerateXBogus 使用接口网址生成X-Bogus参数/Generate X-Bogus parameter using API URL
//
// POST /api/v1/douyin/web/generate_x_bogus
func (r DouyinWebResource) GenerateXBogus(ctx context.Context, request DouyinWebGenerateXBogusRequest) (*DouyinWebGenerateXBogusResponse, error) {
	return r.client.DouyinWebGenerateXBogusParameterUsingAPIURL(ctx, request)
}

// DouyinWebGenerateABogusRequest is the request for POST /api/v1/douyin/web/generate_a_bogus.
type DouyinWebGenerateABogusRequest = DouyinWebGenerateABogusParameterUsingAPIURLRequest

// DouyinWebGenerateABogusResponse is the response for POST /api/v1/douyin/web/generate_a_bogus.
type DouyinWebGenerateABogusResponse = DouyinWebGenerateABogusParameterUsingAPIURLResponse

// GenerateABogus 使用接口网址生成A-Bogus参数/Generate A-Bogus parameter using API URL
//
// POST /api/v1/douyin/web/generate_a_bogus
func (r DouyinWebResource) GenerateABogus(ctx context.Context, request DouyinWebGenerateABogusRequest) (*DouyinWebGenerateABogusResponse, error) {
	return r.client.DouyinWebGenerateABogusParameterUsingAPIURL(ctx, request)
}

// DouyinWebGetSecUserIDRequest is the request for GET /api/v1/douyin/web/get_sec_user_id.
type DouyinWebGetSecUserIDRequest = DouyinWebExtractSingleUserIDRequest

// DouyinWebGetSecUserIDResponse is the response for GET /api/v1/douyin/web/get_sec_user_id.
type DouyinWebGetSecUserIDResponse = DouyinWebExtractSingleUserIDResponse

// GetSecUserID 提取单个用户id/Extract single user id
//
// GET /api/v1/douyin/web/get_sec_user_id
func (r DouyinWebResource) GetSecUserID(ctx context.Context, request DouyinWebGetSecUserIDRequest) (*DouyinWebGetSecUserIDResponse, error) {
	return r.client.DouyinWebExtractSingleUserID(ctx, request)
}

// DouyinWebGetAllSecUserIDRequest is the request for POST /api/v1/douyin/web/get_all_sec_user_id.
type DouyinWebGetAllSecUserIDRequest = DouyinWebExtractListUserIDRequest

// DouyinWebGetAllSecUserIDResponse is the response for POST /api/v1/douyin/web/get_all_sec_user_id.
type DouyinWebGetAllSecUserIDResponse = DouyinWebExtractListUserIDResponse

// GetAllSecUserID 提取列表用户id/Extract list user id
//
// POST /api/v1/douyin/web/get_all_sec_user_id
func (r DouyinWebResource) GetAllSecUserID(ctx context.Context, request DouyinWebGetAllSecUserIDRequest) (*DouyinWebGetAllSecUserIDResponse, error) {
	return r.client.DouyinWebExtractListUserID(ctx, request)
}

// DouyinWebGetAwemeIDRequest is the request for GET /api/v1/douyin/web/get_aweme_id.
type DouyinWebGetAwemeIDRequest = DouyinWebExtractSingleVideoIDRequest

// DouyinWebGetAwemeIDResponse is the response for GET /api/v1/douyin/web/get_aweme_id.
type DouyinWebGetAwemeIDResponse = DouyinWebExtractSingleVideoIDResponse

// GetAwemeID 提取单个作品id/Extract single video id
//
// GET /api/v1/douyin/web/get_aweme_id
func (r DouyinWebResource) GetAwemeID(ctx context.Context, request DouyinWebGetAwemeIDRequest) (*DouyinWebGetAwemeIDResponse, error) {
	return r.client.DouyinWebExtractSingleVideoID(ctx, request)
}

// DouyinWebGetAllAwemeIDRequest is the request for POST /api/v1/douyin/web/get_all_aweme_id.
type DouyinWebGetAllAwemeIDRequest = DouyinWebExtractListVideoIDRequest

// DouyinWebGetAllAwemeIDResponse is the response for POST /api/v1/douyin/web/get_all_aweme_id.
type DouyinWebGetAllAwemeIDResponse = DouyinWebExtractListVideoIDResponse

// GetAllAwemeID 提取列表作品id/Extract list video id
//
// POST /api/v1/douyin/web/get_all_aweme_id
func (r DouyinWebResource) GetAllAwemeID(ctx context.Context, request DouyinWebGetAllAwemeIDRequest) (*DouyinWebGetAllAwemeIDResponse, error) {
	return r.client.DouyinWebExtractListVideoID(ctx, request)
}

// DouyinWebGetWebcastIDRequest is the request for GET /api/v1/douyin/web/get_webcast_id.
type DouyinWebGetWebcastIDRequest = DouyinWebExtractWebcastIDRequest

// DouyinWebGetWebcastIDResponse is the response for GET /api/v1/douyin/web/get_webcast_id.
type DouyinWebGetWebcastIDResponse = DouyinWebExtractWebcastIDResponse

// GetWebcastID 提取直播间号/Extract webcast id
//
// GET /api/v1/douyin/web/get_webcast_id
func (r DouyinWebResource) GetWebcastID(ctx context.Context, request DouyinWebGetWebcastIDRequest) (*DouyinWebGetWebcastIDResponse, error) {
	return r.client.DouyinWebExtractWebcastID(ctx, request)
}

// DouyinWebGetAllWebcastIDRequest is the request for POST /api/v1/douyin/web/get_all_webcast_id.
type DouyinWebGetAllWebcastIDRequest = DouyinWebExtractListWebcastIDRequest

// DouyinWebGetAllWebcastIDResponse is the response for POST /api/v1/douyin/web/get_all_webcast_id.
type DouyinWebGetAllWebcastIDResponse = DouyinWebExtractListWebcastIDResponse

// GetAllWebcastID 提取列表直播间号/Extract list webcast id
//
// POST /api/v1/douyin/web/get_all_webcast_id
func (r DouyinWebResource) GetAllWebcastID(ctx context.Context, request DouyinWebGetAllWebcastIDRequest) (*DouyinWebGetAllWebcastIDResponse, error) {
	return r.client.DouyinWebExtractListWebcastID(ctx, request)
}

// DouyinWebWebcastID2RoomIDRequest is the request for GET /api/v1/douyin/web/webcast_id_2_room_id.
type DouyinWebWebcastID2RoomIDRequest = DouyinWebWebcastIDToRoomIDRequest

// DouyinWebWebcastID2RoomIDResponse is the response for GET /api/v1/douyin/web/webcast_id_2_room_id.
type DouyinWebWebcastID2RoomIDResponse = DouyinWebWebcastIDToRoomIDResponse

// WebcastID2RoomID 直播间号转房间号/Webcast id to room id
//
// GET /api/v1/douyin/web/webcast_id_2_room_id
func (r DouyinWebResource) WebcastID2RoomID(ctx context.Context, request DouyinWebWebcastID2RoomIDRequest) (*DouyinWebWebcastID2RoomIDResponse, error) {
	return r.client.DouyinWebWebcastIDToRoomID(ctx, request)
}

// DouyinWebDouyinLiveRoomRequest is the request for GET /api/v1/douyin/web/douyin_live_room.
type DouyinWebDouyinLiveRoomRequest = DouyinWebExtractLiveRoomDanmakuRequest

// DouyinWebDouyinLiveRoomResponse is the response for GET /api/v1/douyin/web/douyin_live_room.
type DouyinWebDouyinLiveRoomResponse = DouyinWebExtractLiveRoomDanmakuResponse

// DouyinLiveRoom 提取直播间弹幕/Extract live room danmaku
//
// GET /api/v1/douyin/web/douyin_live_room
func (r DouyinWebResource) DouyinLiveRoom(ctx context.Context, request DouyinWebDouyinLiveRoomRequest) (*DouyinWebDouyinLiveRoomResponse, error) {
	return r.client.DouyinWebExtractLiveRoomDanmaku(ctx, request)
}

// DouyinWebFetchLiveImFetchRequest is the request for GET /api/v1/douyin/web/fetch_live_im_fetch.
type DouyinWebFetchLiveImFetchRequest = DouyinWebDouyinLiveRoomDanmakuParametersRequest

// DouyinWebFetchLiveImFetchResponse is the response for GET /api/v1/douyin/web/fetch_live_im_fetch.
type DouyinWebFetchLiveImFetchResponse = DouyinWebDouyinLiveRoomDanmakuParametersResponse

// FetchLiveImFetch 抖音直播间弹幕参数获取/Douyin live room danmaku parameters
//
// GET /api/v1/douyin/web/fetch_live_im_fetch
func (r DouyinWebResource) FetchLiveImFetch(ctx context.Context, request DouyinWebFetchLiveImFetchRequest) (*DouyinWebFetchLiveImFetchResponse, error) {
	return r.client.DouyinWebDouyinLiveRoomDanmakuParameters(ctx, request)
}

// DouyinWebFetchSeriesAwemeRequest is the request for GET /api/v1/douyin/web/fetch_series_aweme.
type DouyinWebFetchSeriesAwemeRequest = DouyinWebSeriesVideoRequest

// DouyinWebFetchSeriesAwemeResponse is the response for GET /api/v1/douyin/web/fetch_series_aweme.
type DouyinWebFetchSeriesAwemeResponse = DouyinWebSeriesVideoResponse

// FetchSeriesAweme 短剧作品/Series Video
//
// GET /api/v1/douyin/web/fetch_series_aweme
func (r DouyinWebResource) FetchSeriesAweme(ctx context.Context, request DouyinWebFetchSeriesAwemeRequest) (*DouyinWebFetchSeriesAwemeResponse, error) {
	return r.client.DouyinWebSeriesVideo(ctx, request)
}

// DouyinWebFetchKnowledgeAwemeRequest is the request for GET /api/v1/douyin/web/fetch_knowledge_aweme.
type DouyinWebFetchKnowledgeAwemeRequest = DouyinWebKnowledgeVideoRequest

// DouyinWebFetchKnowledgeAwemeResponse is the response for GET /api/v1/douyin/web/fetch_knowledge_aweme.
type DouyinWebFetchKnowledgeAwemeResponse = DouyinWebKnowledgeVideoResponse

// FetchKnowledgeAweme 知识作品推荐/Knowledge Video
//
// GET /api/v1/douyin/web/fetch_knowledge_aweme
func (r DouyinWebResource) FetchKnowledgeAweme(ctx context.Context, request DouyinWebFetchKnowledgeAwemeRequest) (*DouyinWebFetchKnowledgeAwemeResponse, error) {
	return r.client.DouyinWebKnowledgeVideo(ctx, request)
}

// DouyinWebFetchGameAwemeRequest is the request for GET /api/v1/douyin/web/fetch_game_aweme.
type DouyinWebFetchGameAwemeRequest = DouyinWebGameVideoRequest

// DouyinWebFetchGameAwemeResponse is the response for GET /api/v1/douyin/web/fetch_game_aweme.
type DouyinWebFetchGameAwemeResponse = DouyinWebGameVideoResponse

// FetchGameAweme 游戏作品推荐/Game Video
//
// GET /api/v1/douyin/web/fetch_game_aweme
func (r DouyinWebResource) FetchGameAweme(ctx context.Context, request DouyinWebFetchGameAwemeRequest) (*DouyinWebFetchGameAwemeResponse, error) {
	return r.client.DouyinWebGameVideo(ctx, request)
}

// DouyinWebFetchCartoonAwemeRequest is the request for GET /api/v1/douyin/web/fetch_cartoon_aweme.
type DouyinWebFetchCartoonAwemeRequest = DouyinWebAnimeVideoRequest

// DouyinWebFetchCartoonAwemeResponse is the response for GET /api/v1/douyin/web/fetch_cartoon_aweme.
type DouyinWebFetchCartoonAwemeResponse = DouyinWebAnimeVideoResponse

// FetchCartoonAweme 二次元作品推荐/Anime Video
//
// GET /api/v1/douyin/web/fetch_cartoon_aweme
func (r DouyinWebResource) FetchCartoonAweme(ctx context.Context, request DouyinWebFetchCartoonAwemeRequest) (*DouyinWebFetchCartoonAwemeResponse, error) {
	return r.client.DouyinWebAnimeVideo(ctx, request)
}

// DouyinWebFetchMusicAwemeRequest is the request for GET /api/v1/douyin/web/fetch_music_aweme.
type DouyinWebFetchMusicAwemeRequest = DouyinWebMusicVideoRequest

// DouyinWebFetchMusicAwemeResponse is the response for GET /api/v1/douyin/web/fetch_music_aweme.
type DouyinWebFetchMusicAwemeResponse = DouyinWebMusicVideoResponse

// FetchMusicAweme 音乐作品推荐/Music Video
//
// GET /api/v1/douyin/web/fetch_music_aweme
func (r DouyinWebResource) FetchMusicAweme(ctx context.Context, request DouyinWebFetchMusicAwemeRequest) (*DouyinWebFetchMusicAwemeResponse, error) {
	return r.client.DouyinWebMusicVideo(ctx, request)
}

// DouyinWebFetchFoodAwemeRequest is the request for GET /api/v1/douyin/web/fetch_food_aweme.
type DouyinWebFetchFoodAwemeRequest = DouyinWebFoodVideoRequest

// DouyinWebFetchFoodAwemeResponse is the response for GET /api/v1/douyin/web/fetch_food_aweme.
type DouyinWebFetchFoodAwemeResponse = DouyinWebFoodVideoResponse

// FetchFoodAweme 美食作品推荐/Food Video
//
// GET /api/v1/douyin/web/fetch_food_aweme
func (r DouyinWebResource) FetchFoodAweme(ctx context.Context, request DouyinWebFetchFoodAwemeRequest) (*DouyinWebFetchFoodAwemeResponse, error) {
	return r.client.DouyinWebFoodVideo(ctx, request)
}

// DouyinAppV3Resource contains endpoints from the Douyin-App-V3-API tag.
type DouyinAppV3Resource struct {
	client *Client
}

// DouyinAppV3FetchOneVideoRequest is the request for GET /api/v1/douyin/app/v3/fetch_one_video.
type DouyinAppV3FetchOneVideoRequest = DouyinAppV3GetSingleVideoDataRequest

// DouyinAppV3FetchOneVideoResponse is the response for GET /api/v1/douyin/app/v3/fetch_one_video.
type DouyinAppV3FetchOneVideoResponse = DouyinAppV3GetSingleVideoDataResponse

// FetchOneVideo 获取单个作品数据/Get single video data
//
// GET /api/v1/douyin/app/v3/fetch_one_video
func (r DouyinAppV3Resource) FetchOneVideo(ctx context.Context, request DouyinAppV3FetchOneVideoRequest) (*DouyinAppV3FetchOneVideoResponse, error) {
	return r.client.DouyinAppV3GetSingleVideoData(ctx, request)
}

// DouyinAppV3FetchOneVideoV2Request is the request for GET /api/v1/douyin/app/v3/fetch_one_video_v2.
type DouyinAppV3FetchOneVideoV2Request = DouyinAppV3GetSingleVideoDataV2Request

// DouyinAppV3FetchOneVideoV2Response is the response for GET /api/v1/douyin/app/v3/fetch_one_video_v2.
type DouyinAppV3FetchOneVideoV2Response = DouyinAppV3GetSingleVideoDataV2Response

// FetchOneVideoV2 获取单个作品数据 V2/Get single video data V2
//
// GET /api/v1/douyin/app/v3/fetch_one_video_v2
func (r DouyinAppV3Resource) FetchOneVideoV2(ctx context.Context, request DouyinAppV3FetchOneVideoV2Request) (*DouyinAppV3FetchOneVideoV2Response, error) {
	return r.client.DouyinAppV3GetSingleVideoDataV2(ctx, request)
}

// DouyinAppV3FetchOneVideoV3Request is the request for GET /api/v1/douyin/app/v3/fetch_one_video_v3.
type DouyinAppV3FetchOneVideoV3Request = DouyinAppV3GetSingleVideoDataV3Request

// DouyinAppV3FetchOneVideoV3Response is the response for GET /api/v1/douyin/app/v3/fetch_one_video_v3.
type DouyinAppV3FetchOneVideoV3Response = DouyinAppV3GetSingleVideoDataV3Response

// FetchOneVideoV3 获取单个作品数据 V3 (无版权限制)/Get single video data V3 (No copyright restrictions)
//
// GET /api/v1/douyin/app/v3/fetch_one_video_v3
func (r DouyinAppV3Resource) FetchOneVideoV3(ctx context.Context, request DouyinAppV3FetchOneVideoV3Request) (*DouyinAppV3FetchOneVideoV3Response, error) {
	return r.client.DouyinAppV3GetSingleVideoDataV3(ctx, request)
}

// DouyinAppV3FetchShareInfoByShareCodeRequest is the request for GET /api/v1/douyin/app/v3/fetch_share_info_by_share_code.
type DouyinAppV3FetchShareInfoByShareCodeRequest = DouyinAppV3GetShareInfoByShareCodeRequest

// DouyinAppV3FetchShareInfoByShareCodeResponse is the response for GET /api/v1/douyin/app/v3/fetch_share_info_by_share_code.
type DouyinAppV3FetchShareInfoByShareCodeResponse = DouyinAppV3GetShareInfoByShareCodeResponse

// FetchShareInfoByShareCode 根据分享口令获取分享信息/Get share info by share code
//
// GET /api/v1/douyin/app/v3/fetch_share_info_by_share_code
func (r DouyinAppV3Resource) FetchShareInfoByShareCode(ctx context.Context, request DouyinAppV3FetchShareInfoByShareCodeRequest) (*DouyinAppV3FetchShareInfoByShareCodeResponse, error) {
	return r.client.DouyinAppV3GetShareInfoByShareCode(ctx, request)
}

// DouyinAppV3FetchMultiVideoRequest is the request for POST /api/v1/douyin/app/v3/fetch_multi_video.
type DouyinAppV3FetchMultiVideoRequest = DouyinAppV3BatchGetVideoInformationV1Request

// DouyinAppV3FetchMultiVideoResponse is the response for POST /api/v1/douyin/app/v3/fetch_multi_video.
type DouyinAppV3FetchMultiVideoResponse = DouyinAppV3BatchGetVideoInformationV1Response

// FetchMultiVideo 批量获取视频信息 V1/Batch Get Video Information V1
//
// POST /api/v1/douyin/app/v3/fetch_multi_video
func (r DouyinAppV3Resource) FetchMultiVideo(ctx context.Context, request DouyinAppV3FetchMultiVideoRequest) (*DouyinAppV3FetchMultiVideoResponse, error) {
	return r.client.DouyinAppV3BatchGetVideoInformationV1(ctx, request)
}

// DouyinAppV3FetchMultiVideoV2Request is the request for POST /api/v1/douyin/app/v3/fetch_multi_video_v2.
type DouyinAppV3FetchMultiVideoV2Request = DouyinAppV3BatchGetVideoInformationV2Request

// DouyinAppV3FetchMultiVideoV2Response is the response for POST /api/v1/douyin/app/v3/fetch_multi_video_v2.
type DouyinAppV3FetchMultiVideoV2Response = DouyinAppV3BatchGetVideoInformationV2Response

// FetchMultiVideoV2 批量获取视频信息 V2/Batch Get Video Information V2
//
// POST /api/v1/douyin/app/v3/fetch_multi_video_v2
func (r DouyinAppV3Resource) FetchMultiVideoV2(ctx context.Context, request DouyinAppV3FetchMultiVideoV2Request) (*DouyinAppV3FetchMultiVideoV2Response, error) {
	return r.client.DouyinAppV3BatchGetVideoInformationV2(ctx, request)
}

// DouyinAppV3FetchOneVideoByShareURLRequest is the request for GET /api/v1/douyin/app/v3/fetch_one_video_by_share_url.
type DouyinAppV3FetchOneVideoByShareURLRequest = DouyinAppV3GetSingleVideoDataBySharingLinkRequest

// DouyinAppV3FetchOneVideoByShareURLResponse is the response for GET /api/v1/douyin/app/v3/fetch_one_video_by_share_url.
type DouyinAppV3FetchOneVideoByShareURLResponse = DouyinAppV3GetSingleVideoDataBySharingLinkResponse

// FetchOneVideoByShareURL 根据分享链接获取单个作品数据/Get single video data by sharing link
//
// GET /api/v1/douyin/app/v3/fetch_one_video_by_share_url
func (r DouyinAppV3Resource) FetchOneVideoByShareURL(ctx context.Context, request DouyinAppV3FetchOneVideoByShareURLRequest) (*DouyinAppV3FetchOneVideoByShareURLResponse, error) {
	return r.client.DouyinAppV3GetSingleVideoDataBySharingLink(ctx, request)
}

// DouyinAppV3FetchVideoHighQualityPlayURLRequest is the request for GET /api/v1/douyin/app/v3/fetch_video_high_quality_play_url.
type DouyinAppV3FetchVideoHighQualityPlayURLRequest = DouyinAppV3GetTheHighestQualityPlayURLOfTheVideoRequest

// DouyinAppV3FetchVideoHighQualityPlayURLResponse is the response for GET /api/v1/douyin/app/v3/fetch_video_high_quality_play_url.
type DouyinAppV3FetchVideoHighQualityPlayURLResponse = DouyinAppV3GetTheHighestQualityPlayURLOfTheVideoResponse

// FetchVideoHighQualityPlayURL 获取视频的最高画质播放链接/Get the highest quality play URL of the video
//
// GET /api/v1/douyin/app/v3/fetch_video_high_quality_play_url
func (r DouyinAppV3Resource) FetchVideoHighQualityPlayURL(ctx context.Context, request DouyinAppV3FetchVideoHighQualityPlayURLRequest) (*DouyinAppV3FetchVideoHighQualityPlayURLResponse, error) {
	return r.client.DouyinAppV3GetTheHighestQualityPlayURLOfTheVideo(ctx, request)
}

// DouyinAppV3FetchMultiVideoHighQualityPlayURLRequest is the request for POST /api/v1/douyin/app/v3/fetch_multi_video_high_quality_play_url.
type DouyinAppV3FetchMultiVideoHighQualityPlayURLRequest = DouyinAppV3BatchGetTheHighestQualityPlayURLOfVideosRequest

// DouyinAppV3FetchMultiVideoHighQualityPlayURLResponse is the response for POST /api/v1/douyin/app/v3/fetch_multi_video_high_quality_play_url.
type DouyinAppV3FetchMultiVideoHighQualityPlayURLResponse = DouyinAppV3BatchGetTheHighestQualityPlayURLOfVideosResponse

// FetchMultiVideoHighQualityPlayURL 批量获取视频的最高画质播放链接/Batch get the highest quality play URL of videos
//
// POST /api/v1/douyin/app/v3/fetch_multi_video_high_quality_play_url
func (r DouyinAppV3Resource) FetchMultiVideoHighQualityPlayURL(ctx context.Context, request DouyinAppV3FetchMultiVideoHighQualityPlayURLRequest) (*DouyinAppV3FetchMultiVideoHighQualityPlayURLResponse, error) {
	return r.client.DouyinAppV3BatchGetTheHighestQualityPlayURLOfVideos(ctx, request)
}

// DouyinAppV3FetchVideoStatisticsRequest is the request for GET /api/v1/douyin/app/v3/fetch_video_statistics.
type DouyinAppV3FetchVideoStatisticsRequest = DouyinAppV3GetTheStatisticalDataOfThePostAccordingToTheVideoIDRequest

// DouyinAppV3FetchVideoStatisticsResponse is the response for GET /api/v1/douyin/app/v3/fetch_video_statistics.
type DouyinAppV3FetchVideoStatisticsResponse = DouyinAppV3GetTheStatisticalDataOfThePostAccordingToTheVideoIDResponse

// FetchVideoStatistics 根据视频ID获取作品的统计数据（点赞数、下载数、播放数、分享数）/Get the statistical data of the Post according to the video ID (like count, download count, play count, share count)
//
// GET /api/v1/douyin/app/v3/fetch_video_statistics
func (r DouyinAppV3Resource) FetchVideoStatistics(ctx context.Context, request DouyinAppV3FetchVideoStatisticsRequest) (*DouyinAppV3FetchVideoStatisticsResponse, error) {
	return r.client.DouyinAppV3GetTheStatisticalDataOfThePostAccordingToTheVideoID(ctx, request)
}

// DouyinAppV3FetchMultiVideoStatisticsRequest is the request for GET /api/v1/douyin/app/v3/fetch_multi_video_statistics.
type DouyinAppV3FetchMultiVideoStatisticsRequest = DouyinAppV3GetTheStatisticalDataOfThePostAccordingToTheVideoIDV3FetchMultiVideoStatisticsRequest

// DouyinAppV3FetchMultiVideoStatisticsResponse is the response for GET /api/v1/douyin/app/v3/fetch_multi_video_statistics.
type DouyinAppV3FetchMultiVideoStatisticsResponse = DouyinAppV3GetTheStatisticalDataOfThePostAccordingToTheVideoIDV3FetchMultiVideoStatisticsResponse

// FetchMultiVideoStatistics 根据视频ID批量获取作品的统计数据（点赞数、下载数、播放数、分享数）/Get the statistical data of the Post according to the video ID (like count, download count, play count, share count)
//
// GET /api/v1/douyin/app/v3/fetch_multi_video_statistics
func (r DouyinAppV3Resource) FetchMultiVideoStatistics(ctx context.Context, request DouyinAppV3FetchMultiVideoStatisticsRequest) (*DouyinAppV3FetchMultiVideoStatisticsResponse, error) {
	return r.client.DouyinAppV3GetTheStatisticalDataOfThePostAccordingToTheVideoIDV3FetchMultiVideoStatistics(ctx, request)
}

// DouyinAppV3AddVideoPlayCountRequest is the request for GET /api/v1/douyin/app/v3/add_video_play_count.
type DouyinAppV3AddVideoPlayCountRequest = DouyinAppV3IncreaseTheNumberOfPlaysOfTheWorkAccordingToTheVideoIDRequest

// DouyinAppV3AddVideoPlayCountResponse is the response for GET /api/v1/douyin/app/v3/add_video_play_count.
type DouyinAppV3AddVideoPlayCountResponse = DouyinAppV3IncreaseTheNumberOfPlaysOfTheWorkAccordingToTheVideoIDResponse

// AddVideoPlayCount 根据视频ID来增加作品的播放数/Increase the number of plays of the work according to the video ID
//
// GET /api/v1/douyin/app/v3/add_video_play_count
func (r DouyinAppV3Resource) AddVideoPlayCount(ctx context.Context, request DouyinAppV3AddVideoPlayCountRequest) (*DouyinAppV3AddVideoPlayCountResponse, error) {
	return r.client.DouyinAppV3IncreaseTheNumberOfPlaysOfTheWorkAccordingToTheVideoID(ctx, request)
}

// DouyinAppV3HandlerUserProfileRequest is the request for GET /api/v1/douyin/app/v3/handler_user_profile.
type DouyinAppV3HandlerUserProfileRequest = DouyinAppV3GetInformationOfSpecifiedUserRequest

// DouyinAppV3HandlerUserProfileResponse is the response for GET /api/v1/douyin/app/v3/handler_user_profile.
type DouyinAppV3HandlerUserProfileResponse = DouyinAppV3GetInformationOfSpecifiedUserResponse

// HandlerUserProfile 获取指定用户的信息/Get information of specified user
//
// GET /api/v1/douyin/app/v3/handler_user_profile
func (r DouyinAppV3Resource) HandlerUserProfile(ctx context.Context, request DouyinAppV3HandlerUserProfileRequest) (*DouyinAppV3HandlerUserProfileResponse, error) {
	return r.client.DouyinAppV3GetInformationOfSpecifiedUser(ctx, request)
}

// DouyinAppV3FetchUserFansListRequest is the request for GET /api/v1/douyin/app/v3/fetch_user_fans_list.
type DouyinAppV3FetchUserFansListRequest = DouyinAppV3GetUserFansListRequest

// DouyinAppV3FetchUserFansListResponse is the response for GET /api/v1/douyin/app/v3/fetch_user_fans_list.
type DouyinAppV3FetchUserFansListResponse = DouyinAppV3GetUserFansListResponse

// FetchUserFansList 获取用户粉丝列表/Get user fans list
//
// GET /api/v1/douyin/app/v3/fetch_user_fans_list
func (r DouyinAppV3Resource) FetchUserFansList(ctx context.Context, request DouyinAppV3FetchUserFansListRequest) (*DouyinAppV3FetchUserFansListResponse, error) {
	return r.client.DouyinAppV3GetUserFansList(ctx, request)
}

// DouyinAppV3FetchUserFollowingListRequest is the request for GET /api/v1/douyin/app/v3/fetch_user_following_list.
type DouyinAppV3FetchUserFollowingListRequest = DouyinAppV3APIV1DouyinWebFetchUserFollowingListGetUserFollowingListRequest

// DouyinAppV3FetchUserFollowingListResponse is the response for GET /api/v1/douyin/app/v3/fetch_user_following_list.
type DouyinAppV3FetchUserFollowingListResponse = DouyinAppV3APIV1DouyinWebFetchUserFollowingListGetUserFollowingListResponse

// FetchUserFollowingList 获取用户关注列表 (弃用，使用 /api/v1/douyin/web/fetch_user_following_list 替代)/Get user following list (Deprecated, use /api/v1/douyin/web/fetch_user_following_list instead)
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/douyin/app/v3/fetch_user_following_list
func (r DouyinAppV3Resource) FetchUserFollowingList(ctx context.Context, request DouyinAppV3FetchUserFollowingListRequest) (*DouyinAppV3FetchUserFollowingListResponse, error) {
	return r.client.DouyinAppV3APIV1DouyinWebFetchUserFollowingListGetUserFollowingList(ctx, request)
}

// DouyinAppV3FetchUserPostVideosRequest is the request for GET /api/v1/douyin/app/v3/fetch_user_post_videos.
type DouyinAppV3FetchUserPostVideosRequest = DouyinAppV3GetUserHomepageVideoDataRequest

// DouyinAppV3FetchUserPostVideosResponse is the response for GET /api/v1/douyin/app/v3/fetch_user_post_videos.
type DouyinAppV3FetchUserPostVideosResponse = DouyinAppV3GetUserHomepageVideoDataResponse

// FetchUserPostVideos 获取用户主页作品数据/Get user homepage video data
//
// GET /api/v1/douyin/app/v3/fetch_user_post_videos
func (r DouyinAppV3Resource) FetchUserPostVideos(ctx context.Context, request DouyinAppV3FetchUserPostVideosRequest) (*DouyinAppV3FetchUserPostVideosResponse, error) {
	return r.client.DouyinAppV3GetUserHomepageVideoData(ctx, request)
}

// DouyinAppV3FetchUserLikeVideosRequest is the request for GET /api/v1/douyin/app/v3/fetch_user_like_videos.
type DouyinAppV3FetchUserLikeVideosRequest = DouyinAppV3GetUserLikeVideoDataRequest

// DouyinAppV3FetchUserLikeVideosResponse is the response for GET /api/v1/douyin/app/v3/fetch_user_like_videos.
type DouyinAppV3FetchUserLikeVideosResponse = DouyinAppV3GetUserLikeVideoDataResponse

// FetchUserLikeVideos 获取用户喜欢作品数据/Get user like video data
//
// GET /api/v1/douyin/app/v3/fetch_user_like_videos
func (r DouyinAppV3Resource) FetchUserLikeVideos(ctx context.Context, request DouyinAppV3FetchUserLikeVideosRequest) (*DouyinAppV3FetchUserLikeVideosResponse, error) {
	return r.client.DouyinAppV3GetUserLikeVideoData(ctx, request)
}

// DouyinAppV3FetchVideoCommentsRequest is the request for GET /api/v1/douyin/app/v3/fetch_video_comments.
type DouyinAppV3FetchVideoCommentsRequest = DouyinAppV3GetSingleVideoCommentsDataRequest

// DouyinAppV3FetchVideoCommentsResponse is the response for GET /api/v1/douyin/app/v3/fetch_video_comments.
type DouyinAppV3FetchVideoCommentsResponse = DouyinAppV3GetSingleVideoCommentsDataResponse

// FetchVideoComments 获取单个视频评论数据/Get single video comments data
//
// GET /api/v1/douyin/app/v3/fetch_video_comments
func (r DouyinAppV3Resource) FetchVideoComments(ctx context.Context, request DouyinAppV3FetchVideoCommentsRequest) (*DouyinAppV3FetchVideoCommentsResponse, error) {
	return r.client.DouyinAppV3GetSingleVideoCommentsData(ctx, request)
}

// DouyinAppV3FetchVideoCommentRepliesRequest is the request for GET /api/v1/douyin/app/v3/fetch_video_comment_replies.
type DouyinAppV3FetchVideoCommentRepliesRequest = DouyinAppV3GetCommentRepliesDataOfSpecifiedVideoRequest

// DouyinAppV3FetchVideoCommentRepliesResponse is the response for GET /api/v1/douyin/app/v3/fetch_video_comment_replies.
type DouyinAppV3FetchVideoCommentRepliesResponse = DouyinAppV3GetCommentRepliesDataOfSpecifiedVideoResponse

// FetchVideoCommentReplies 获取指定视频的评论回复数据/Get comment replies data of specified video
//
// GET /api/v1/douyin/app/v3/fetch_video_comment_replies
func (r DouyinAppV3Resource) FetchVideoCommentReplies(ctx context.Context, request DouyinAppV3FetchVideoCommentRepliesRequest) (*DouyinAppV3FetchVideoCommentRepliesResponse, error) {
	return r.client.DouyinAppV3GetCommentRepliesDataOfSpecifiedVideo(ctx, request)
}

// DouyinAppV3FetchVideoMixDetailRequest is the request for GET /api/v1/douyin/app/v3/fetch_video_mix_detail.
type DouyinAppV3FetchVideoMixDetailRequest = DouyinAppV3GetDouyinVideoMixDetailDataRequest

// DouyinAppV3FetchVideoMixDetailResponse is the response for GET /api/v1/douyin/app/v3/fetch_video_mix_detail.
type DouyinAppV3FetchVideoMixDetailResponse = DouyinAppV3GetDouyinVideoMixDetailDataResponse

// FetchVideoMixDetail 获取抖音视频合集详情数据/Get Douyin video mix detail data
//
// GET /api/v1/douyin/app/v3/fetch_video_mix_detail
func (r DouyinAppV3Resource) FetchVideoMixDetail(ctx context.Context, request DouyinAppV3FetchVideoMixDetailRequest) (*DouyinAppV3FetchVideoMixDetailResponse, error) {
	return r.client.DouyinAppV3GetDouyinVideoMixDetailData(ctx, request)
}

// DouyinAppV3FetchVideoMixPostListRequest is the request for GET /api/v1/douyin/app/v3/fetch_video_mix_post_list.
type DouyinAppV3FetchVideoMixPostListRequest = DouyinAppV3GetDouyinVideoMixPostListDataRequest

// DouyinAppV3FetchVideoMixPostListResponse is the response for GET /api/v1/douyin/app/v3/fetch_video_mix_post_list.
type DouyinAppV3FetchVideoMixPostListResponse = DouyinAppV3GetDouyinVideoMixPostListDataResponse

// FetchVideoMixPostList 获取抖音视频合集作品列表数据/Get Douyin video mix post list data
//
// GET /api/v1/douyin/app/v3/fetch_video_mix_post_list
func (r DouyinAppV3Resource) FetchVideoMixPostList(ctx context.Context, request DouyinAppV3FetchVideoMixPostListRequest) (*DouyinAppV3FetchVideoMixPostListResponse, error) {
	return r.client.DouyinAppV3GetDouyinVideoMixPostListData(ctx, request)
}

// DouyinAppV3FetchUserSeriesListRequest is the request for GET /api/v1/douyin/app/v3/fetch_user_series_list.
type DouyinAppV3FetchUserSeriesListRequest = DouyinAppV3GetUserSeriesListRequest

// DouyinAppV3FetchUserSeriesListResponse is the response for GET /api/v1/douyin/app/v3/fetch_user_series_list.
type DouyinAppV3FetchUserSeriesListResponse = DouyinAppV3GetUserSeriesListResponse

// FetchUserSeriesList 获取用户短剧合集列表/Get user series list
//
// GET /api/v1/douyin/app/v3/fetch_user_series_list
func (r DouyinAppV3Resource) FetchUserSeriesList(ctx context.Context, request DouyinAppV3FetchUserSeriesListRequest) (*DouyinAppV3FetchUserSeriesListResponse, error) {
	return r.client.DouyinAppV3GetUserSeriesList(ctx, request)
}

// DouyinAppV3FetchSeriesVideoListRequest is the request for GET /api/v1/douyin/app/v3/fetch_series_video_list.
type DouyinAppV3FetchSeriesVideoListRequest = DouyinAppV3GetSeriesVideoListRequest

// DouyinAppV3FetchSeriesVideoListResponse is the response for GET /api/v1/douyin/app/v3/fetch_series_video_list.
type DouyinAppV3FetchSeriesVideoListResponse = DouyinAppV3GetSeriesVideoListResponse

// FetchSeriesVideoList 获取短剧视频列表/Get series video list
//
// GET /api/v1/douyin/app/v3/fetch_series_video_list
func (r DouyinAppV3Resource) FetchSeriesVideoList(ctx context.Context, request DouyinAppV3FetchSeriesVideoListRequest) (*DouyinAppV3FetchSeriesVideoListResponse, error) {
	return r.client.DouyinAppV3GetSeriesVideoList(ctx, request)
}

// DouyinAppV3FetchSeriesDetailRequest is the request for GET /api/v1/douyin/app/v3/fetch_series_detail.
type DouyinAppV3FetchSeriesDetailRequest = DouyinAppV3GetSeriesDetailRequest

// DouyinAppV3FetchSeriesDetailResponse is the response for GET /api/v1/douyin/app/v3/fetch_series_detail.
type DouyinAppV3FetchSeriesDetailResponse = DouyinAppV3GetSeriesDetailResponse

// FetchSeriesDetail 获取短剧详情信息/Get series detail
//
// GET /api/v1/douyin/app/v3/fetch_series_detail
func (r DouyinAppV3Resource) FetchSeriesDetail(ctx context.Context, request DouyinAppV3FetchSeriesDetailRequest) (*DouyinAppV3FetchSeriesDetailResponse, error) {
	return r.client.DouyinAppV3GetSeriesDetail(ctx, request)
}

// DouyinAppV3FetchGeneralSearchResultRequest is the request for GET /api/v1/douyin/app/v3/fetch_general_search_result.
type DouyinAppV3FetchGeneralSearchResultRequest = DouyinAppV3GetComprehensiveSearchResultsOfSpecifiedKeywordsRequest

// DouyinAppV3FetchGeneralSearchResultResponse is the response for GET /api/v1/douyin/app/v3/fetch_general_search_result.
type DouyinAppV3FetchGeneralSearchResultResponse = DouyinAppV3GetComprehensiveSearchResultsOfSpecifiedKeywordsResponse

// FetchGeneralSearchResult 获取指定关键词的综合搜索结果（弃用，替代接口见下方文档说明）/Get comprehensive search results of specified keywords (deprecated, see the documentation below for alternative interfaces)
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/douyin/app/v3/fetch_general_search_result
func (r DouyinAppV3Resource) FetchGeneralSearchResult(ctx context.Context, request DouyinAppV3FetchGeneralSearchResultRequest) (*DouyinAppV3FetchGeneralSearchResultResponse, error) {
	return r.client.DouyinAppV3GetComprehensiveSearchResultsOfSpecifiedKeywords(ctx, request)
}

// DouyinAppV3FetchVideoSearchResultRequest is the request for GET /api/v1/douyin/app/v3/fetch_video_search_result.
type DouyinAppV3FetchVideoSearchResultRequest = DouyinAppV3GetVideoSearchResultsOfSpecifiedKeywordsRequest

// DouyinAppV3FetchVideoSearchResultResponse is the response for GET /api/v1/douyin/app/v3/fetch_video_search_result.
type DouyinAppV3FetchVideoSearchResultResponse = DouyinAppV3GetVideoSearchResultsOfSpecifiedKeywordsResponse

// FetchVideoSearchResult 获取指定关键词的视频搜索结果（弃用，替代接口见下方文档说明）/Get video search results of specified keywords (deprecated, see the documentation below for alternative interfaces)
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/douyin/app/v3/fetch_video_search_result
func (r DouyinAppV3Resource) FetchVideoSearchResult(ctx context.Context, request DouyinAppV3FetchVideoSearchResultRequest) (*DouyinAppV3FetchVideoSearchResultResponse, error) {
	return r.client.DouyinAppV3GetVideoSearchResultsOfSpecifiedKeywords(ctx, request)
}

// DouyinAppV3FetchUserSearchResultRequest is the request for GET /api/v1/douyin/app/v3/fetch_user_search_result.
type DouyinAppV3FetchUserSearchResultRequest = DouyinAppV3GetUserSearchResultsOfSpecifiedKeywordsRequest

// DouyinAppV3FetchUserSearchResultResponse is the response for GET /api/v1/douyin/app/v3/fetch_user_search_result.
type DouyinAppV3FetchUserSearchResultResponse = DouyinAppV3GetUserSearchResultsOfSpecifiedKeywordsResponse

// FetchUserSearchResult 获取指定关键词的用户搜索结果（弃用，替代接口见下方文档说明）/Get user search results of specified keywords (deprecated, see the documentation below for alternative interfaces)
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/douyin/app/v3/fetch_user_search_result
func (r DouyinAppV3Resource) FetchUserSearchResult(ctx context.Context, request DouyinAppV3FetchUserSearchResultRequest) (*DouyinAppV3FetchUserSearchResultResponse, error) {
	return r.client.DouyinAppV3GetUserSearchResultsOfSpecifiedKeywords(ctx, request)
}

// DouyinAppV3FetchLiveSearchResultRequest is the request for GET /api/v1/douyin/app/v3/fetch_live_search_result.
type DouyinAppV3FetchLiveSearchResultRequest = DouyinAppV3GetLiveSearchResultsOfSpecifiedKeywordsRequest

// DouyinAppV3FetchLiveSearchResultResponse is the response for GET /api/v1/douyin/app/v3/fetch_live_search_result.
type DouyinAppV3FetchLiveSearchResultResponse = DouyinAppV3GetLiveSearchResultsOfSpecifiedKeywordsResponse

// FetchLiveSearchResult 获取指定关键词的直播搜索结果（弃用，替代接口见下方文档说明）/Get live search results of specified keywords (deprecated, see the documentation below for alternative interfaces)
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/douyin/app/v3/fetch_live_search_result
func (r DouyinAppV3Resource) FetchLiveSearchResult(ctx context.Context, request DouyinAppV3FetchLiveSearchResultRequest) (*DouyinAppV3FetchLiveSearchResultResponse, error) {
	return r.client.DouyinAppV3GetLiveSearchResultsOfSpecifiedKeywords(ctx, request)
}

// DouyinAppV3FetchMusicSearchResultRequest is the request for GET /api/v1/douyin/app/v3/fetch_music_search_result.
type DouyinAppV3FetchMusicSearchResultRequest = DouyinAppV3GetMusicSearchResultsOfSpecifiedKeywordsRequest

// DouyinAppV3FetchMusicSearchResultResponse is the response for GET /api/v1/douyin/app/v3/fetch_music_search_result.
type DouyinAppV3FetchMusicSearchResultResponse = DouyinAppV3GetMusicSearchResultsOfSpecifiedKeywordsResponse

// FetchMusicSearchResult 获取指定关键词的音乐搜索结果（弃用，替代接口见下方文档说明）/Get music search results of specified keywords (deprecated, see the documentation below for alternative interfaces)
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/douyin/app/v3/fetch_music_search_result
func (r DouyinAppV3Resource) FetchMusicSearchResult(ctx context.Context, request DouyinAppV3FetchMusicSearchResultRequest) (*DouyinAppV3FetchMusicSearchResultResponse, error) {
	return r.client.DouyinAppV3GetMusicSearchResultsOfSpecifiedKeywords(ctx, request)
}

// DouyinAppV3FetchHashtagSearchResultRequest is the request for GET /api/v1/douyin/app/v3/fetch_hashtag_search_result.
type DouyinAppV3FetchHashtagSearchResultRequest = DouyinAppV3GetHashtagSearchResultsOfSpecifiedKeywordsRequest

// DouyinAppV3FetchHashtagSearchResultResponse is the response for GET /api/v1/douyin/app/v3/fetch_hashtag_search_result.
type DouyinAppV3FetchHashtagSearchResultResponse = DouyinAppV3GetHashtagSearchResultsOfSpecifiedKeywordsResponse

// FetchHashtagSearchResult 获取指定关键词的话题搜索结果（弃用，替代接口见下方文档说明）/Get hashtag search results of specified keywords (deprecated, see the documentation below for alternative interfaces)
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/douyin/app/v3/fetch_hashtag_search_result
func (r DouyinAppV3Resource) FetchHashtagSearchResult(ctx context.Context, request DouyinAppV3FetchHashtagSearchResultRequest) (*DouyinAppV3FetchHashtagSearchResultResponse, error) {
	return r.client.DouyinAppV3GetHashtagSearchResultsOfSpecifiedKeywords(ctx, request)
}

// DouyinAppV3FetchMusicDetailRequest is the request for GET /api/v1/douyin/app/v3/fetch_music_detail.
type DouyinAppV3FetchMusicDetailRequest = DouyinAppV3GetDetailsOfSpecifiedMusicRequest

// DouyinAppV3FetchMusicDetailResponse is the response for GET /api/v1/douyin/app/v3/fetch_music_detail.
type DouyinAppV3FetchMusicDetailResponse = DouyinAppV3GetDetailsOfSpecifiedMusicResponse

// FetchMusicDetail 获取指定音乐的详情数据/Get details of specified music
//
// GET /api/v1/douyin/app/v3/fetch_music_detail
func (r DouyinAppV3Resource) FetchMusicDetail(ctx context.Context, request DouyinAppV3FetchMusicDetailRequest) (*DouyinAppV3FetchMusicDetailResponse, error) {
	return r.client.DouyinAppV3GetDetailsOfSpecifiedMusic(ctx, request)
}

// DouyinAppV3FetchMusicVideoListRequest is the request for GET /api/v1/douyin/app/v3/fetch_music_video_list.
type DouyinAppV3FetchMusicVideoListRequest = DouyinAppV3GetVideoListOfSpecifiedMusicRequest

// DouyinAppV3FetchMusicVideoListResponse is the response for GET /api/v1/douyin/app/v3/fetch_music_video_list.
type DouyinAppV3FetchMusicVideoListResponse = DouyinAppV3GetVideoListOfSpecifiedMusicResponse

// FetchMusicVideoList 获取指定音乐的视频列表数据/Get video list of specified music
//
// GET /api/v1/douyin/app/v3/fetch_music_video_list
func (r DouyinAppV3Resource) FetchMusicVideoList(ctx context.Context, request DouyinAppV3FetchMusicVideoListRequest) (*DouyinAppV3FetchMusicVideoListResponse, error) {
	return r.client.DouyinAppV3GetVideoListOfSpecifiedMusic(ctx, request)
}

// DouyinAppV3FetchHashtagDetailRequest is the request for GET /api/v1/douyin/app/v3/fetch_hashtag_detail.
type DouyinAppV3FetchHashtagDetailRequest = DouyinAppV3GetDetailsOfSpecifiedHashtagRequest

// DouyinAppV3FetchHashtagDetailResponse is the response for GET /api/v1/douyin/app/v3/fetch_hashtag_detail.
type DouyinAppV3FetchHashtagDetailResponse = DouyinAppV3GetDetailsOfSpecifiedHashtagResponse

// FetchHashtagDetail 获取指定话题的详情数据/Get details of specified hashtag
//
// GET /api/v1/douyin/app/v3/fetch_hashtag_detail
func (r DouyinAppV3Resource) FetchHashtagDetail(ctx context.Context, request DouyinAppV3FetchHashtagDetailRequest) (*DouyinAppV3FetchHashtagDetailResponse, error) {
	return r.client.DouyinAppV3GetDetailsOfSpecifiedHashtag(ctx, request)
}

// DouyinAppV3FetchHashtagVideoListRequest is the request for GET /api/v1/douyin/app/v3/fetch_hashtag_video_list.
type DouyinAppV3FetchHashtagVideoListRequest = DouyinAppV3GetVideoListOfSpecifiedHashtagRequest

// DouyinAppV3FetchHashtagVideoListResponse is the response for GET /api/v1/douyin/app/v3/fetch_hashtag_video_list.
type DouyinAppV3FetchHashtagVideoListResponse = DouyinAppV3GetVideoListOfSpecifiedHashtagResponse

// FetchHashtagVideoList 获取指定话题的作品数据/Get video list of specified hashtag
//
// GET /api/v1/douyin/app/v3/fetch_hashtag_video_list
func (r DouyinAppV3Resource) FetchHashtagVideoList(ctx context.Context, request DouyinAppV3FetchHashtagVideoListRequest) (*DouyinAppV3FetchHashtagVideoListResponse, error) {
	return r.client.DouyinAppV3GetVideoListOfSpecifiedHashtag(ctx, request)
}

// DouyinAppV3FetchHotSearchListRequest is the request for GET /api/v1/douyin/app/v3/fetch_hot_search_list.
type DouyinAppV3FetchHotSearchListRequest = DouyinAppV3GetDouyinHotSearchListDataRequest

// DouyinAppV3FetchHotSearchListResponse is the response for GET /api/v1/douyin/app/v3/fetch_hot_search_list.
type DouyinAppV3FetchHotSearchListResponse = DouyinAppV3GetDouyinHotSearchListDataResponse

// FetchHotSearchList 获取抖音热搜榜数据/Get Douyin hot search list data
//
// GET /api/v1/douyin/app/v3/fetch_hot_search_list
func (r DouyinAppV3Resource) FetchHotSearchList(ctx context.Context, request DouyinAppV3FetchHotSearchListRequest) (*DouyinAppV3FetchHotSearchListResponse, error) {
	return r.client.DouyinAppV3GetDouyinHotSearchListData(ctx, request)
}

// DouyinAppV3FetchLiveHotSearchListResponse is the response for GET /api/v1/douyin/app/v3/fetch_live_hot_search_list.
type DouyinAppV3FetchLiveHotSearchListResponse = DouyinAppV3GetDouyinLiveHotSearchListDataResponse

// FetchLiveHotSearchList 获取抖音直播热搜榜数据/Get Douyin live hot search list data
//
// GET /api/v1/douyin/app/v3/fetch_live_hot_search_list
func (r DouyinAppV3Resource) FetchLiveHotSearchList(ctx context.Context) (*DouyinAppV3FetchLiveHotSearchListResponse, error) {
	return r.client.DouyinAppV3GetDouyinLiveHotSearchListData(ctx)
}

// DouyinAppV3FetchMusicHotSearchListRequest is the request for GET /api/v1/douyin/app/v3/fetch_music_hot_search_list.
type DouyinAppV3FetchMusicHotSearchListRequest = DouyinAppV3GetDouyinMusicHotSearchListDataRequest

// DouyinAppV3FetchMusicHotSearchListResponse is the response for GET /api/v1/douyin/app/v3/fetch_music_hot_search_list.
type DouyinAppV3FetchMusicHotSearchListResponse = DouyinAppV3GetDouyinMusicHotSearchListDataResponse

// FetchMusicHotSearchList 获取抖音音乐榜数据/Get Douyin music hot search list data
//
// GET /api/v1/douyin/app/v3/fetch_music_hot_search_list
func (r DouyinAppV3Resource) FetchMusicHotSearchList(ctx context.Context, request DouyinAppV3FetchMusicHotSearchListRequest) (*DouyinAppV3FetchMusicHotSearchListResponse, error) {
	return r.client.DouyinAppV3GetDouyinMusicHotSearchListData(ctx, request)
}

// DouyinAppV3FetchBrandHotSearchListResponse is the response for GET /api/v1/douyin/app/v3/fetch_brand_hot_search_list.
type DouyinAppV3FetchBrandHotSearchListResponse = DouyinAppV3GetDouyinBrandHotSearchListDataResponse

// FetchBrandHotSearchList 获取抖音品牌热榜分类数据/Get Douyin brand hot search list data
//
// GET /api/v1/douyin/app/v3/fetch_brand_hot_search_list
func (r DouyinAppV3Resource) FetchBrandHotSearchList(ctx context.Context) (*DouyinAppV3FetchBrandHotSearchListResponse, error) {
	return r.client.DouyinAppV3GetDouyinBrandHotSearchListData(ctx)
}

// DouyinAppV3FetchBrandHotSearchListDetailRequest is the request for GET /api/v1/douyin/app/v3/fetch_brand_hot_search_list_detail.
type DouyinAppV3FetchBrandHotSearchListDetailRequest = DouyinAppV3GetDouyinBrandHotSearchListDetailDataRequest

// DouyinAppV3FetchBrandHotSearchListDetailResponse is the response for GET /api/v1/douyin/app/v3/fetch_brand_hot_search_list_detail.
type DouyinAppV3FetchBrandHotSearchListDetailResponse = DouyinAppV3GetDouyinBrandHotSearchListDetailDataResponse

// FetchBrandHotSearchListDetail 获取抖音品牌热榜具体分类数据/Get Douyin brand hot search list detail data
//
// GET /api/v1/douyin/app/v3/fetch_brand_hot_search_list_detail
func (r DouyinAppV3Resource) FetchBrandHotSearchListDetail(ctx context.Context, request DouyinAppV3FetchBrandHotSearchListDetailRequest) (*DouyinAppV3FetchBrandHotSearchListDetailResponse, error) {
	return r.client.DouyinAppV3GetDouyinBrandHotSearchListDetailData(ctx, request)
}

// DouyinAppV3GenerateDouyinShortURLRequest is the request for GET /api/v1/douyin/app/v3/generate_douyin_short_url.
type DouyinAppV3GenerateDouyinShortURLRequest = DouyinAppV3GenerateDouyinShortLinkRequest

// DouyinAppV3GenerateDouyinShortURLResponse is the response for GET /api/v1/douyin/app/v3/generate_douyin_short_url.
type DouyinAppV3GenerateDouyinShortURLResponse = DouyinAppV3GenerateDouyinShortLinkResponse

// GenerateDouyinShortURL 生成抖音短链接/Generate Douyin short link
//
// GET /api/v1/douyin/app/v3/generate_douyin_short_url
func (r DouyinAppV3Resource) GenerateDouyinShortURL(ctx context.Context, request DouyinAppV3GenerateDouyinShortURLRequest) (*DouyinAppV3GenerateDouyinShortURLResponse, error) {
	return r.client.DouyinAppV3GenerateDouyinShortLink(ctx, request)
}

// DouyinAppV3GenerateDouyinVideoShareQrcodeRequest is the request for GET /api/v1/douyin/app/v3/generate_douyin_video_share_qrcode.
type DouyinAppV3GenerateDouyinVideoShareQrcodeRequest = DouyinAppV3GenerateDouyinVideoShareQRCodeRequest

// DouyinAppV3GenerateDouyinVideoShareQrcodeResponse is the response for GET /api/v1/douyin/app/v3/generate_douyin_video_share_qrcode.
type DouyinAppV3GenerateDouyinVideoShareQrcodeResponse = DouyinAppV3GenerateDouyinVideoShareQRCodeResponse

// GenerateDouyinVideoShareQrcode 生成抖音视频分享二维码/Generate Douyin video share QR code
//
// GET /api/v1/douyin/app/v3/generate_douyin_video_share_qrcode
func (r DouyinAppV3Resource) GenerateDouyinVideoShareQrcode(ctx context.Context, request DouyinAppV3GenerateDouyinVideoShareQrcodeRequest) (*DouyinAppV3GenerateDouyinVideoShareQrcodeResponse, error) {
	return r.client.DouyinAppV3GenerateDouyinVideoShareQRCode(ctx, request)
}

// DouyinAppV3RegisterDeviceRequest is the request for GET /api/v1/douyin/app/v3/register_device.
type DouyinAppV3RegisterDeviceRequest = DouyinAppV3DouyinAppRegisterDeviceRequest

// DouyinAppV3RegisterDeviceResponse is the response for GET /api/v1/douyin/app/v3/register_device.
type DouyinAppV3RegisterDeviceResponse = DouyinAppV3DouyinAppRegisterDeviceResponse

// RegisterDevice 抖音APP注册设备/Douyin APP register device
//
// GET /api/v1/douyin/app/v3/register_device
func (r DouyinAppV3Resource) RegisterDevice(ctx context.Context, request DouyinAppV3RegisterDeviceRequest) (*DouyinAppV3RegisterDeviceResponse, error) {
	return r.client.DouyinAppV3DouyinAppRegisterDevice(ctx, request)
}

// DouyinAppV3OpenDouyinAppToVideoDetailRequest is the request for GET /api/v1/douyin/app/v3/open_douyin_app_to_video_detail.
type DouyinAppV3OpenDouyinAppToVideoDetailRequest = DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndJumpToTheSpecifiedVideoDetailsPageRequest

// DouyinAppV3OpenDouyinAppToVideoDetailResponse is the response for GET /api/v1/douyin/app/v3/open_douyin_app_to_video_detail.
type DouyinAppV3OpenDouyinAppToVideoDetailResponse = DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndJumpToTheSpecifiedVideoDetailsPageResponse

// OpenDouyinAppToVideoDetail 生成抖音分享链接，唤起抖音APP，跳转指定作品详情页/Generate Douyin share link, call Douyin APP, and jump to the specified video details page
//
// GET /api/v1/douyin/app/v3/open_douyin_app_to_video_detail
func (r DouyinAppV3Resource) OpenDouyinAppToVideoDetail(ctx context.Context, request DouyinAppV3OpenDouyinAppToVideoDetailRequest) (*DouyinAppV3OpenDouyinAppToVideoDetailResponse, error) {
	return r.client.DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndJumpToTheSpecifiedVideoDetailsPage(ctx, request)
}

// DouyinAppV3OpenDouyinAppToUserProfileRequest is the request for GET /api/v1/douyin/app/v3/open_douyin_app_to_user_profile.
type DouyinAppV3OpenDouyinAppToUserProfileRequest = DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndJumpToTheSpecifiedUserProfileRequest

// DouyinAppV3OpenDouyinAppToUserProfileResponse is the response for GET /api/v1/douyin/app/v3/open_douyin_app_to_user_profile.
type DouyinAppV3OpenDouyinAppToUserProfileResponse = DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndJumpToTheSpecifiedUserProfileResponse

// OpenDouyinAppToUserProfile 生成抖音分享链接，唤起抖音APP，跳转指定用户主页/Generate Douyin share link, call Douyin APP, and jump to the specified user profile
//
// GET /api/v1/douyin/app/v3/open_douyin_app_to_user_profile
func (r DouyinAppV3Resource) OpenDouyinAppToUserProfile(ctx context.Context, request DouyinAppV3OpenDouyinAppToUserProfileRequest) (*DouyinAppV3OpenDouyinAppToUserProfileResponse, error) {
	return r.client.DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndJumpToTheSpecifiedUserProfile(ctx, request)
}

// DouyinAppV3OpenDouyinAppToKeywordSearchRequest is the request for GET /api/v1/douyin/app/v3/open_douyin_app_to_keyword_search.
type DouyinAppV3OpenDouyinAppToKeywordSearchRequest = DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndJumpToTheSpecifiedKeywordSearchResultRequest

// DouyinAppV3OpenDouyinAppToKeywordSearchResponse is the response for GET /api/v1/douyin/app/v3/open_douyin_app_to_keyword_search.
type DouyinAppV3OpenDouyinAppToKeywordSearchResponse = DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndJumpToTheSpecifiedKeywordSearchResultResponse

// OpenDouyinAppToKeywordSearch 生成抖音分享链接，唤起抖音APP，跳转指定关键词搜索结果/Generate Douyin share link, call Douyin APP, and jump to the specified keyword search result
//
// GET /api/v1/douyin/app/v3/open_douyin_app_to_keyword_search
func (r DouyinAppV3Resource) OpenDouyinAppToKeywordSearch(ctx context.Context, request DouyinAppV3OpenDouyinAppToKeywordSearchRequest) (*DouyinAppV3OpenDouyinAppToKeywordSearchResponse, error) {
	return r.client.DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndJumpToTheSpecifiedKeywordSearchResult(ctx, request)
}

// DouyinAppV3OpenDouyinAppToSendPrivateMessageRequest is the request for GET /api/v1/douyin/app/v3/open_douyin_app_to_send_private_message.
type DouyinAppV3OpenDouyinAppToSendPrivateMessageRequest = DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndSendPrivateMessagesToSpecifiedUsersRequest

// DouyinAppV3OpenDouyinAppToSendPrivateMessageResponse is the response for GET /api/v1/douyin/app/v3/open_douyin_app_to_send_private_message.
type DouyinAppV3OpenDouyinAppToSendPrivateMessageResponse = DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndSendPrivateMessagesToSpecifiedUsersResponse

// OpenDouyinAppToSendPrivateMessage 生成抖音分享链接，唤起抖音APP，给指定用户发送私信/Generate Douyin share link, call Douyin APP, and send private messages to specified users
//
// GET /api/v1/douyin/app/v3/open_douyin_app_to_send_private_message
func (r DouyinAppV3Resource) OpenDouyinAppToSendPrivateMessage(ctx context.Context, request DouyinAppV3OpenDouyinAppToSendPrivateMessageRequest) (*DouyinAppV3OpenDouyinAppToSendPrivateMessageResponse, error) {
	return r.client.DouyinAppV3GenerateDouyinShareLinkCallDouyinAppAndSendPrivateMessagesToSpecifiedUsers(ctx, request)
}

// DouyinCreatorResource contains endpoints from the Douyin-Creator-API tag.
type DouyinCreatorResource struct {
	client *Client
}

// DouyinCreatorFetchCreatorActivityListRequest is the request for GET /api/v1/douyin/creator/fetch_creator_activity_list.
type DouyinCreatorFetchCreatorActivityListRequest = DouyinCreatorGetCreatorActivityListRequest

// DouyinCreatorFetchCreatorActivityListResponse is the response for GET /api/v1/douyin/creator/fetch_creator_activity_list.
type DouyinCreatorFetchCreatorActivityListResponse = DouyinCreatorGetCreatorActivityListResponse

// FetchCreatorActivityList 获取创作者活动列表/Get creator activity list
//
// GET /api/v1/douyin/creator/fetch_creator_activity_list
func (r DouyinCreatorResource) FetchCreatorActivityList(ctx context.Context, request DouyinCreatorFetchCreatorActivityListRequest) (*DouyinCreatorFetchCreatorActivityListResponse, error) {
	return r.client.DouyinCreatorGetCreatorActivityList(ctx, request)
}

// DouyinCreatorFetchCreatorActivityDetailRequest is the request for GET /api/v1/douyin/creator/fetch_creator_activity_detail.
type DouyinCreatorFetchCreatorActivityDetailRequest = DouyinCreatorGetCreatorActivityDetailRequest

// DouyinCreatorFetchCreatorActivityDetailResponse is the response for GET /api/v1/douyin/creator/fetch_creator_activity_detail.
type DouyinCreatorFetchCreatorActivityDetailResponse = DouyinCreatorGetCreatorActivityDetailResponse

// FetchCreatorActivityDetail 获取创作者活动详情/Get creator activity detail
//
// GET /api/v1/douyin/creator/fetch_creator_activity_detail
func (r DouyinCreatorResource) FetchCreatorActivityDetail(ctx context.Context, request DouyinCreatorFetchCreatorActivityDetailRequest) (*DouyinCreatorFetchCreatorActivityDetailResponse, error) {
	return r.client.DouyinCreatorGetCreatorActivityDetail(ctx, request)
}

// DouyinCreatorFetchCreatorMaterialCenterConfigResponse is the response for GET /api/v1/douyin/creator/fetch_creator_material_center_config.
type DouyinCreatorFetchCreatorMaterialCenterConfigResponse = DouyinCreatorGetCreatorMaterialCenterConfigResponse

// FetchCreatorMaterialCenterConfig 获取创作者中心配置/Get creator material center config
//
// GET /api/v1/douyin/creator/fetch_creator_material_center_config
func (r DouyinCreatorResource) FetchCreatorMaterialCenterConfig(ctx context.Context) (*DouyinCreatorFetchCreatorMaterialCenterConfigResponse, error) {
	return r.client.DouyinCreatorGetCreatorMaterialCenterConfig(ctx)
}

// DouyinCreatorFetchCreatorMaterialCenterBillboardRequest is the request for GET /api/v1/douyin/creator/fetch_creator_material_center_billboard.
type DouyinCreatorFetchCreatorMaterialCenterBillboardRequest = DouyinCreatorGetCreatorMaterialCenterBillboardRequest

// DouyinCreatorFetchCreatorMaterialCenterBillboardResponse is the response for GET /api/v1/douyin/creator/fetch_creator_material_center_billboard.
type DouyinCreatorFetchCreatorMaterialCenterBillboardResponse = DouyinCreatorGetCreatorMaterialCenterBillboardResponse

// FetchCreatorMaterialCenterBillboard 获取创作者中心热门视频榜单/Get creator material center billboard
//
// GET /api/v1/douyin/creator/fetch_creator_material_center_billboard
func (r DouyinCreatorResource) FetchCreatorMaterialCenterBillboard(ctx context.Context, request DouyinCreatorFetchCreatorMaterialCenterBillboardRequest) (*DouyinCreatorFetchCreatorMaterialCenterBillboardResponse, error) {
	return r.client.DouyinCreatorGetCreatorMaterialCenterBillboard(ctx, request)
}

// DouyinCreatorFetchCreatorMaterialCenterRelatedRequest is the request for GET /api/v1/douyin/creator/fetch_creator_material_center_related.
type DouyinCreatorFetchCreatorMaterialCenterRelatedRequest = DouyinCreatorGetTopicOrHotSpotRelatedVideosRequest

// DouyinCreatorFetchCreatorMaterialCenterRelatedResponse is the response for GET /api/v1/douyin/creator/fetch_creator_material_center_related.
type DouyinCreatorFetchCreatorMaterialCenterRelatedResponse = DouyinCreatorGetTopicOrHotSpotRelatedVideosResponse

// FetchCreatorMaterialCenterRelated 获取话题/热点相关视频/Get topic or hot spot related videos
//
// GET /api/v1/douyin/creator/fetch_creator_material_center_related
func (r DouyinCreatorResource) FetchCreatorMaterialCenterRelated(ctx context.Context, request DouyinCreatorFetchCreatorMaterialCenterRelatedRequest) (*DouyinCreatorFetchCreatorMaterialCenterRelatedResponse, error) {
	return r.client.DouyinCreatorGetTopicOrHotSpotRelatedVideos(ctx, request)
}

// DouyinCreatorFetchCreatorHotSpotBillboardRequest is the request for GET /api/v1/douyin/creator/fetch_creator_hot_spot_billboard.
type DouyinCreatorFetchCreatorHotSpotBillboardRequest = DouyinCreatorGetCreatorHotSpotBillboardRequest

// DouyinCreatorFetchCreatorHotSpotBillboardResponse is the response for GET /api/v1/douyin/creator/fetch_creator_hot_spot_billboard.
type DouyinCreatorFetchCreatorHotSpotBillboardResponse = DouyinCreatorGetCreatorHotSpotBillboardResponse

// FetchCreatorHotSpotBillboard 获取创作者中心创作热点/Get creator hot spot billboard
//
// GET /api/v1/douyin/creator/fetch_creator_hot_spot_billboard
func (r DouyinCreatorResource) FetchCreatorHotSpotBillboard(ctx context.Context, request DouyinCreatorFetchCreatorHotSpotBillboardRequest) (*DouyinCreatorFetchCreatorHotSpotBillboardResponse, error) {
	return r.client.DouyinCreatorGetCreatorHotSpotBillboard(ctx, request)
}

// DouyinCreatorFetchCreatorHotTopicBillboardRequest is the request for GET /api/v1/douyin/creator/fetch_creator_hot_topic_billboard.
type DouyinCreatorFetchCreatorHotTopicBillboardRequest = DouyinCreatorGetCreatorHotTopicBillboardRequest

// DouyinCreatorFetchCreatorHotTopicBillboardResponse is the response for GET /api/v1/douyin/creator/fetch_creator_hot_topic_billboard.
type DouyinCreatorFetchCreatorHotTopicBillboardResponse = DouyinCreatorGetCreatorHotTopicBillboardResponse

// FetchCreatorHotTopicBillboard 获取创作者热门话题榜单/Get creator hot topic billboard
//
// GET /api/v1/douyin/creator/fetch_creator_hot_topic_billboard
func (r DouyinCreatorResource) FetchCreatorHotTopicBillboard(ctx context.Context, request DouyinCreatorFetchCreatorHotTopicBillboardRequest) (*DouyinCreatorFetchCreatorHotTopicBillboardResponse, error) {
	return r.client.DouyinCreatorGetCreatorHotTopicBillboard(ctx, request)
}

// DouyinCreatorFetchCreatorHotPropsBillboardRequest is the request for GET /api/v1/douyin/creator/fetch_creator_hot_props_billboard.
type DouyinCreatorFetchCreatorHotPropsBillboardRequest = DouyinCreatorGetCreatorHotPropsBillboardRequest

// DouyinCreatorFetchCreatorHotPropsBillboardResponse is the response for GET /api/v1/douyin/creator/fetch_creator_hot_props_billboard.
type DouyinCreatorFetchCreatorHotPropsBillboardResponse = DouyinCreatorGetCreatorHotPropsBillboardResponse

// FetchCreatorHotPropsBillboard 获取创作者热门道具榜单/Get creator hot props billboard
//
// GET /api/v1/douyin/creator/fetch_creator_hot_props_billboard
func (r DouyinCreatorResource) FetchCreatorHotPropsBillboard(ctx context.Context, request DouyinCreatorFetchCreatorHotPropsBillboardRequest) (*DouyinCreatorFetchCreatorHotPropsBillboardResponse, error) {
	return r.client.DouyinCreatorGetCreatorHotPropsBillboard(ctx, request)
}

// DouyinCreatorFetchCreatorHotChallengeBillboardResponse is the response for GET /api/v1/douyin/creator/fetch_creator_hot_challenge_billboard.
type DouyinCreatorFetchCreatorHotChallengeBillboardResponse = DouyinCreatorGetCreatorHotChallengeBillboardResponse

// FetchCreatorHotChallengeBillboard 获取创作者热门挑战榜单/Get creator hot challenge billboard
//
// GET /api/v1/douyin/creator/fetch_creator_hot_challenge_billboard
func (r DouyinCreatorResource) FetchCreatorHotChallengeBillboard(ctx context.Context) (*DouyinCreatorFetchCreatorHotChallengeBillboardResponse, error) {
	return r.client.DouyinCreatorGetCreatorHotChallengeBillboard(ctx)
}

// DouyinCreatorFetchCreatorHotMusicBillboardRequest is the request for GET /api/v1/douyin/creator/fetch_creator_hot_music_billboard.
type DouyinCreatorFetchCreatorHotMusicBillboardRequest = DouyinCreatorGetCreatorHotMusicBillboardRequest

// DouyinCreatorFetchCreatorHotMusicBillboardResponse is the response for GET /api/v1/douyin/creator/fetch_creator_hot_music_billboard.
type DouyinCreatorFetchCreatorHotMusicBillboardResponse = DouyinCreatorGetCreatorHotMusicBillboardResponse

// FetchCreatorHotMusicBillboard 获取创作者热门音乐榜单/Get creator hot music billboard
//
// GET /api/v1/douyin/creator/fetch_creator_hot_music_billboard
func (r DouyinCreatorResource) FetchCreatorHotMusicBillboard(ctx context.Context, request DouyinCreatorFetchCreatorHotMusicBillboardRequest) (*DouyinCreatorFetchCreatorHotMusicBillboardResponse, error) {
	return r.client.DouyinCreatorGetCreatorHotMusicBillboard(ctx, request)
}

// DouyinCreatorFetchCreatorHotCourseRequest is the request for GET /api/v1/douyin/creator/fetch_creator_hot_course.
type DouyinCreatorFetchCreatorHotCourseRequest = DouyinCreatorGetCreatorHotCourseRequest

// DouyinCreatorFetchCreatorHotCourseResponse is the response for GET /api/v1/douyin/creator/fetch_creator_hot_course.
type DouyinCreatorFetchCreatorHotCourseResponse = DouyinCreatorGetCreatorHotCourseResponse

// FetchCreatorHotCourse 获取创作者热门课程/Get creator hot course
//
// GET /api/v1/douyin/creator/fetch_creator_hot_course
func (r DouyinCreatorResource) FetchCreatorHotCourse(ctx context.Context, request DouyinCreatorFetchCreatorHotCourseRequest) (*DouyinCreatorFetchCreatorHotCourseResponse, error) {
	return r.client.DouyinCreatorGetCreatorHotCourse(ctx, request)
}

// DouyinCreatorFetchCreatorContentCategoryResponse is the response for GET /api/v1/douyin/creator/fetch_creator_content_category.
type DouyinCreatorFetchCreatorContentCategoryResponse = DouyinCreatorGetCreatorContentCreationCategoryResponse

// FetchCreatorContentCategory 获取创作者内容创作合集分类/Get creator content creation category
//
// GET /api/v1/douyin/creator/fetch_creator_content_category
func (r DouyinCreatorResource) FetchCreatorContentCategory(ctx context.Context) (*DouyinCreatorFetchCreatorContentCategoryResponse, error) {
	return r.client.DouyinCreatorGetCreatorContentCreationCategory(ctx)
}

// DouyinCreatorFetchCreatorContentCourseRequest is the request for GET /api/v1/douyin/creator/fetch_creator_content_course.
type DouyinCreatorFetchCreatorContentCourseRequest = DouyinCreatorGetCreatorContentCreationCourseRequest

// DouyinCreatorFetchCreatorContentCourseResponse is the response for GET /api/v1/douyin/creator/fetch_creator_content_course.
type DouyinCreatorFetchCreatorContentCourseResponse = DouyinCreatorGetCreatorContentCreationCourseResponse

// FetchCreatorContentCourse 获取创作者内容创作课程/Get creator content creation course
//
// GET /api/v1/douyin/creator/fetch_creator_content_course
func (r DouyinCreatorResource) FetchCreatorContentCourse(ctx context.Context, request DouyinCreatorFetchCreatorContentCourseRequest) (*DouyinCreatorFetchCreatorContentCourseResponse, error) {
	return r.client.DouyinCreatorGetCreatorContentCreationCourse(ctx, request)
}

// DouyinCreatorFetchVideoDanmakuListRequest is the request for GET /api/v1/douyin/creator/fetch_video_danmaku_list.
type DouyinCreatorFetchVideoDanmakuListRequest = DouyinCreatorGetVideoDanmakuListRequest

// DouyinCreatorFetchVideoDanmakuListResponse is the response for GET /api/v1/douyin/creator/fetch_video_danmaku_list.
type DouyinCreatorFetchVideoDanmakuListResponse = DouyinCreatorGetVideoDanmakuListResponse

// FetchVideoDanmakuList 获取作品弹幕列表/Get video danmaku list
//
// GET /api/v1/douyin/creator/fetch_video_danmaku_list
func (r DouyinCreatorResource) FetchVideoDanmakuList(ctx context.Context, request DouyinCreatorFetchVideoDanmakuListRequest) (*DouyinCreatorFetchVideoDanmakuListResponse, error) {
	return r.client.DouyinCreatorGetVideoDanmakuList(ctx, request)
}

// DouyinCreatorFetchUserSearchRequest is the request for GET /api/v1/douyin/creator/fetch_user_search.
type DouyinCreatorFetchUserSearchRequest = DouyinCreatorSearchUsersRequest

// DouyinCreatorFetchUserSearchResponse is the response for GET /api/v1/douyin/creator/fetch_user_search.
type DouyinCreatorFetchUserSearchResponse = DouyinCreatorSearchUsersResponse

// FetchUserSearch 搜索用户/Search users
//
// GET /api/v1/douyin/creator/fetch_user_search
func (r DouyinCreatorResource) FetchUserSearch(ctx context.Context, request DouyinCreatorFetchUserSearchRequest) (*DouyinCreatorFetchUserSearchResponse, error) {
	return r.client.DouyinCreatorSearchUsers(ctx, request)
}

// DouyinCreatorFetchMissionTaskListRequest is the request for GET /api/v1/douyin/creator/fetch_mission_task_list.
type DouyinCreatorFetchMissionTaskListRequest = DouyinCreatorGetMissionTaskListRequest

// DouyinCreatorFetchMissionTaskListResponse is the response for GET /api/v1/douyin/creator/fetch_mission_task_list.
type DouyinCreatorFetchMissionTaskListResponse = DouyinCreatorGetMissionTaskListResponse

// FetchMissionTaskList 获取商单任务列表/Get mission task list
//
// GET /api/v1/douyin/creator/fetch_mission_task_list
func (r DouyinCreatorResource) FetchMissionTaskList(ctx context.Context, request DouyinCreatorFetchMissionTaskListRequest) (*DouyinCreatorFetchMissionTaskListResponse, error) {
	return r.client.DouyinCreatorGetMissionTaskList(ctx, request)
}

// DouyinCreatorFetchIndustryCategoryConfigResponse is the response for GET /api/v1/douyin/creator/fetch_industry_category_config.
type DouyinCreatorFetchIndustryCategoryConfigResponse = DouyinCreatorGetIndustryCategoryConfigResponse

// FetchIndustryCategoryConfig 获取行业分类配置/Get industry category config
//
// GET /api/v1/douyin/creator/fetch_industry_category_config
func (r DouyinCreatorResource) FetchIndustryCategoryConfig(ctx context.Context) (*DouyinCreatorFetchIndustryCategoryConfigResponse, error) {
	return r.client.DouyinCreatorGetIndustryCategoryConfig(ctx)
}

// DouyinCreatorV2Resource contains endpoints from the Douyin-Creator-V2-API tag.
type DouyinCreatorV2Resource struct {
	client *Client
}

// FetchItemOverviewData 获取作品总览数据/Fetch item overview data
//
// POST /api/v1/douyin/creator_v2/fetch_item_overview_data
func (r DouyinCreatorV2Resource) FetchItemOverviewData(ctx context.Context, request DouyinCreatorV2FetchItemOverviewDataRequest) (*DouyinCreatorV2FetchItemOverviewDataResponse, error) {
	return r.client.DouyinCreatorV2FetchItemOverviewData(ctx, request)
}

// DouyinCreatorV2FetchItemPlaySourceRequest is the request for POST /api/v1/douyin/creator_v2/fetch_item_play_source.
type DouyinCreatorV2FetchItemPlaySourceRequest = DouyinCreatorV2FetchItemPlaySourceStatisticsRequest

// DouyinCreatorV2FetchItemPlaySourceResponse is the response for POST /api/v1/douyin/creator_v2/fetch_item_play_source.
type DouyinCreatorV2FetchItemPlaySourceResponse = DouyinCreatorV2FetchItemPlaySourceStatisticsResponse

// FetchItemPlaySource 获取作品流量来源统计/Fetch item play source statistics
//
// POST /api/v1/douyin/creator_v2/fetch_item_play_source
func (r DouyinCreatorV2Resource) FetchItemPlaySource(ctx context.Context, request DouyinCreatorV2FetchItemPlaySourceRequest) (*DouyinCreatorV2FetchItemPlaySourceResponse, error) {
	return r.client.DouyinCreatorV2FetchItemPlaySourceStatistics(ctx, request)
}

// DouyinCreatorV2FetchItemSearchKeywordRequest is the request for POST /api/v1/douyin/creator_v2/fetch_item_search_keyword.
type DouyinCreatorV2FetchItemSearchKeywordRequest = DouyinCreatorV2FetchItemSearchKeywordsStatisticsRequest

// DouyinCreatorV2FetchItemSearchKeywordResponse is the response for POST /api/v1/douyin/creator_v2/fetch_item_search_keyword.
type DouyinCreatorV2FetchItemSearchKeywordResponse = DouyinCreatorV2FetchItemSearchKeywordsStatisticsResponse

// FetchItemSearchKeyword 获取作品搜索关键词统计/Fetch item search keywords statistics
//
// POST /api/v1/douyin/creator_v2/fetch_item_search_keyword
func (r DouyinCreatorV2Resource) FetchItemSearchKeyword(ctx context.Context, request DouyinCreatorV2FetchItemSearchKeywordRequest) (*DouyinCreatorV2FetchItemSearchKeywordResponse, error) {
	return r.client.DouyinCreatorV2FetchItemSearchKeywordsStatistics(ctx, request)
}

// DouyinCreatorV2FetchItemWatchTrendRequest is the request for POST /api/v1/douyin/creator_v2/fetch_item_watch_trend.
type DouyinCreatorV2FetchItemWatchTrendRequest = DouyinCreatorV2FetchItemWatchTrendAnalysisRequest

// DouyinCreatorV2FetchItemWatchTrendResponse is the response for POST /api/v1/douyin/creator_v2/fetch_item_watch_trend.
type DouyinCreatorV2FetchItemWatchTrendResponse = DouyinCreatorV2FetchItemWatchTrendAnalysisResponse

// FetchItemWatchTrend 获取作品观看趋势分析/Fetch item watch trend analysis
//
// POST /api/v1/douyin/creator_v2/fetch_item_watch_trend
func (r DouyinCreatorV2Resource) FetchItemWatchTrend(ctx context.Context, request DouyinCreatorV2FetchItemWatchTrendRequest) (*DouyinCreatorV2FetchItemWatchTrendResponse, error) {
	return r.client.DouyinCreatorV2FetchItemWatchTrendAnalysis(ctx, request)
}

// DouyinCreatorV2FetchItemDanmakuAnalysisRequest is the request for POST /api/v1/douyin/creator_v2/fetch_item_danmaku_analysis.
type DouyinCreatorV2FetchItemDanmakuAnalysisRequest = DouyinCreatorV2FetchItemBulletAnalysisRequest

// DouyinCreatorV2FetchItemDanmakuAnalysisResponse is the response for POST /api/v1/douyin/creator_v2/fetch_item_danmaku_analysis.
type DouyinCreatorV2FetchItemDanmakuAnalysisResponse = DouyinCreatorV2FetchItemBulletAnalysisResponse

// FetchItemDanmakuAnalysis 获取作品弹幕分析/Fetch item bullet analysis
//
// POST /api/v1/douyin/creator_v2/fetch_item_danmaku_analysis
func (r DouyinCreatorV2Resource) FetchItemDanmakuAnalysis(ctx context.Context, request DouyinCreatorV2FetchItemDanmakuAnalysisRequest) (*DouyinCreatorV2FetchItemDanmakuAnalysisResponse, error) {
	return r.client.DouyinCreatorV2FetchItemBulletAnalysis(ctx, request)
}

// FetchItemAudiencePortrait 获取作品观众数据分析/Fetch item audience portrait
//
// POST /api/v1/douyin/creator_v2/fetch_item_audience_portrait
func (r DouyinCreatorV2Resource) FetchItemAudiencePortrait(ctx context.Context, request DouyinCreatorV2FetchItemAudiencePortraitRequest) (*DouyinCreatorV2FetchItemAudiencePortraitResponse, error) {
	return r.client.DouyinCreatorV2FetchItemAudiencePortrait(ctx, request)
}

// DouyinCreatorV2FetchItemAudienceOthersRequest is the request for POST /api/v1/douyin/creator_v2/fetch_item_audience_others.
type DouyinCreatorV2FetchItemAudienceOthersRequest = DouyinCreatorV2FetchItemAudienceOthersAnalysisRequest

// DouyinCreatorV2FetchItemAudienceOthersResponse is the response for POST /api/v1/douyin/creator_v2/fetch_item_audience_others.
type DouyinCreatorV2FetchItemAudienceOthersResponse = DouyinCreatorV2FetchItemAudienceOthersAnalysisResponse

// FetchItemAudienceOthers 获取作品观众其他数据分析/Fetch item audience others analysis
//
// POST /api/v1/douyin/creator_v2/fetch_item_audience_others
func (r DouyinCreatorV2Resource) FetchItemAudienceOthers(ctx context.Context, request DouyinCreatorV2FetchItemAudienceOthersRequest) (*DouyinCreatorV2FetchItemAudienceOthersResponse, error) {
	return r.client.DouyinCreatorV2FetchItemAudienceOthersAnalysis(ctx, request)
}

// FetchItemAnalysisInvolvedVertical 获取作品垂类标签/Fetch item analysis involved vertical
//
// POST /api/v1/douyin/creator_v2/fetch_item_analysis_involved_vertical
func (r DouyinCreatorV2Resource) FetchItemAnalysisInvolvedVertical(ctx context.Context, request DouyinCreatorV2FetchItemAnalysisInvolvedVerticalRequest) (*DouyinCreatorV2FetchItemAnalysisInvolvedVerticalResponse, error) {
	return r.client.DouyinCreatorV2FetchItemAnalysisInvolvedVertical(ctx, request)
}

// FetchItemAnalysisOverview 获取投稿分析概览/Fetch item analysis overview
//
// POST /api/v1/douyin/creator_v2/fetch_item_analysis_overview
func (r DouyinCreatorV2Resource) FetchItemAnalysisOverview(ctx context.Context, request DouyinCreatorV2FetchItemAnalysisOverviewRequest) (*DouyinCreatorV2FetchItemAnalysisOverviewResponse, error) {
	return r.client.DouyinCreatorV2FetchItemAnalysisOverview(ctx, request)
}

// FetchItemAnalysisItemPerformance 获取投稿表现数据/Fetch item analysis item performance
//
// POST /api/v1/douyin/creator_v2/fetch_item_analysis_item_performance
func (r DouyinCreatorV2Resource) FetchItemAnalysisItemPerformance(ctx context.Context, request DouyinCreatorV2FetchItemAnalysisItemPerformanceRequest) (*DouyinCreatorV2FetchItemAnalysisItemPerformanceResponse, error) {
	return r.client.DouyinCreatorV2FetchItemAnalysisItemPerformance(ctx, request)
}

// FetchItemList 获取投稿作品列表/Fetch item list
//
// POST /api/v1/douyin/creator_v2/fetch_item_list
func (r DouyinCreatorV2Resource) FetchItemList(ctx context.Context, request DouyinCreatorV2FetchItemListRequest) (*DouyinCreatorV2FetchItemListResponse, error) {
	return r.client.DouyinCreatorV2FetchItemList(ctx, request)
}

// DouyinCreatorV2FetchItemListDownloadRequest is the request for POST /api/v1/douyin/creator_v2/fetch_item_list_download.
type DouyinCreatorV2FetchItemListDownloadRequest = DouyinCreatorV2DownloadItemListRequest

// DouyinCreatorV2FetchItemListDownloadResponse is the response for POST /api/v1/douyin/creator_v2/fetch_item_list_download.
type DouyinCreatorV2FetchItemListDownloadResponse = DouyinCreatorV2DownloadItemListResponse

// FetchItemListDownload 导出投稿作品列表/Download item list
//
// POST /api/v1/douyin/creator_v2/fetch_item_list_download
func (r DouyinCreatorV2Resource) FetchItemListDownload(ctx context.Context, request DouyinCreatorV2FetchItemListDownloadRequest) (*DouyinCreatorV2FetchItemListDownloadResponse, error) {
	return r.client.DouyinCreatorV2DownloadItemList(ctx, request)
}

// FetchLiveRoomHistoryList 获取直播场次历史记录/Fetch live room history list
//
// POST /api/v1/douyin/creator_v2/fetch_live_room_history_list
func (r DouyinCreatorV2Resource) FetchLiveRoomHistoryList(ctx context.Context, request DouyinCreatorV2FetchLiveRoomHistoryListRequest) (*DouyinCreatorV2FetchLiveRoomHistoryListResponse, error) {
	return r.client.DouyinCreatorV2FetchLiveRoomHistoryList(ctx, request)
}

// FetchAuthorDiagnosis 获取创作者账号诊断/Fetch author diagnosis
//
// POST /api/v1/douyin/creator_v2/fetch_author_diagnosis
func (r DouyinCreatorV2Resource) FetchAuthorDiagnosis(ctx context.Context, request DouyinCreatorV2FetchAuthorDiagnosisRequest) (*DouyinCreatorV2FetchAuthorDiagnosisResponse, error) {
	return r.client.DouyinCreatorV2FetchAuthorDiagnosis(ctx, request)
}

// DouyinIndexResource contains endpoints from the Douyin-Index-API tag.
type DouyinIndexResource struct {
	client *Client
}

// DouyinIndexFetchAllValidDateResponse is the response for GET /api/v1/douyin/index/fetch_all_valid_date.
type DouyinIndexFetchAllValidDateResponse = DouyinIndexGetAllValidDatesResponse

// FetchAllValidDate 获取所有有效日期/Get all valid dates
//
// GET /api/v1/douyin/index/fetch_all_valid_date
func (r DouyinIndexResource) FetchAllValidDate(ctx context.Context) (*DouyinIndexFetchAllValidDateResponse, error) {
	return r.client.DouyinIndexGetAllValidDates(ctx)
}

// DouyinIndexFetchValidDateForRelationResponse is the response for GET /api/v1/douyin/index/fetch_valid_date_for_relation.
type DouyinIndexFetchValidDateForRelationResponse = DouyinIndexGetValidDateForRelationResponse

// FetchValidDateForRelation 获取关联分析有效日期/Get valid date for relation
//
// GET /api/v1/douyin/index/fetch_valid_date_for_relation
func (r DouyinIndexResource) FetchValidDateForRelation(ctx context.Context) (*DouyinIndexFetchValidDateForRelationResponse, error) {
	return r.client.DouyinIndexGetValidDateForRelation(ctx)
}

// DouyinIndexFetchAllAreaResponse is the response for GET /api/v1/douyin/index/fetch_all_area.
type DouyinIndexFetchAllAreaResponse = DouyinIndexGetAllAreaListResponse

// FetchAllArea 获取所有地区列表/Get all area list
//
// GET /api/v1/douyin/index/fetch_all_area
func (r DouyinIndexResource) FetchAllArea(ctx context.Context) (*DouyinIndexFetchAllAreaResponse, error) {
	return r.client.DouyinIndexGetAllAreaList(ctx)
}

// DouyinIndexFetchCurrentHotTopicResponse is the response for GET /api/v1/douyin/index/fetch_current_hot_topic.
type DouyinIndexFetchCurrentHotTopicResponse = DouyinIndexGetCurrentHotTopicsResponse

// FetchCurrentHotTopic 获取实时热点排行/Get current hot topics
//
// GET /api/v1/douyin/index/fetch_current_hot_topic
func (r DouyinIndexResource) FetchCurrentHotTopic(ctx context.Context) (*DouyinIndexFetchCurrentHotTopicResponse, error) {
	return r.client.DouyinIndexGetCurrentHotTopics(ctx)
}

// DouyinIndexFetchHotWordsRequest is the request for GET /api/v1/douyin/index/fetch_hot_words.
type DouyinIndexFetchHotWordsRequest = DouyinIndexGetHotWordsRequest

// DouyinIndexFetchHotWordsResponse is the response for GET /api/v1/douyin/index/fetch_hot_words.
type DouyinIndexFetchHotWordsResponse = DouyinIndexGetHotWordsResponse

// FetchHotWords 获取热门关键词/Get hot words
//
// GET /api/v1/douyin/index/fetch_hot_words
func (r DouyinIndexResource) FetchHotWords(ctx context.Context, request DouyinIndexFetchHotWordsRequest) (*DouyinIndexFetchHotWordsResponse, error) {
	return r.client.DouyinIndexGetHotWords(ctx, request)
}

// DouyinIndexFetchKeywordValidDateRequest is the request for POST /api/v1/douyin/index/fetch_keyword_valid_date.
type DouyinIndexFetchKeywordValidDateRequest = DouyinIndexGetKeywordValidDateRequest

// DouyinIndexFetchKeywordValidDateResponse is the response for POST /api/v1/douyin/index/fetch_keyword_valid_date.
type DouyinIndexFetchKeywordValidDateResponse = DouyinIndexGetKeywordValidDateResponse

// FetchKeywordValidDate 获取关键词有效日期/Get keyword valid date
//
// POST /api/v1/douyin/index/fetch_keyword_valid_date
func (r DouyinIndexResource) FetchKeywordValidDate(ctx context.Context, request DouyinIndexFetchKeywordValidDateRequest) (*DouyinIndexFetchKeywordValidDateResponse, error) {
	return r.client.DouyinIndexGetKeywordValidDate(ctx, request)
}

// DouyinIndexFetchMultiKeywordHotTrendRequest is the request for POST /api/v1/douyin/index/fetch_multi_keyword_hot_trend.
type DouyinIndexFetchMultiKeywordHotTrendRequest = DouyinIndexGetMultiKeywordHotTrendRequest

// DouyinIndexFetchMultiKeywordHotTrendResponse is the response for POST /api/v1/douyin/index/fetch_multi_keyword_hot_trend.
type DouyinIndexFetchMultiKeywordHotTrendResponse = DouyinIndexGetMultiKeywordHotTrendResponse

// FetchMultiKeywordHotTrend 获取多关键词热度趋势/Get multi-keyword hot trend
//
// POST /api/v1/douyin/index/fetch_multi_keyword_hot_trend
func (r DouyinIndexResource) FetchMultiKeywordHotTrend(ctx context.Context, request DouyinIndexFetchMultiKeywordHotTrendRequest) (*DouyinIndexFetchMultiKeywordHotTrendResponse, error) {
	return r.client.DouyinIndexGetMultiKeywordHotTrend(ctx, request)
}

// DouyinIndexFetchMultiKeywordInterpretationRequest is the request for POST /api/v1/douyin/index/fetch_multi_keyword_interpretation.
type DouyinIndexFetchMultiKeywordInterpretationRequest = DouyinIndexGetMultiKeywordInterpretationRequest

// DouyinIndexFetchMultiKeywordInterpretationResponse is the response for POST /api/v1/douyin/index/fetch_multi_keyword_interpretation.
type DouyinIndexFetchMultiKeywordInterpretationResponse = DouyinIndexGetMultiKeywordInterpretationResponse

// FetchMultiKeywordInterpretation 获取多关键词解读/Get multi-keyword interpretation
//
// POST /api/v1/douyin/index/fetch_multi_keyword_interpretation
func (r DouyinIndexResource) FetchMultiKeywordInterpretation(ctx context.Context, request DouyinIndexFetchMultiKeywordInterpretationRequest) (*DouyinIndexFetchMultiKeywordInterpretationResponse, error) {
	return r.client.DouyinIndexGetMultiKeywordInterpretation(ctx, request)
}

// DouyinIndexFetchRelationWordRequest is the request for POST /api/v1/douyin/index/fetch_relation_word.
type DouyinIndexFetchRelationWordRequest = DouyinIndexGetRelationWordAnalysisRequest

// DouyinIndexFetchRelationWordResponse is the response for POST /api/v1/douyin/index/fetch_relation_word.
type DouyinIndexFetchRelationWordResponse = DouyinIndexGetRelationWordAnalysisResponse

// FetchRelationWord 获取关联词分析/Get relation word analysis
//
// POST /api/v1/douyin/index/fetch_relation_word
func (r DouyinIndexResource) FetchRelationWord(ctx context.Context, request DouyinIndexFetchRelationWordRequest) (*DouyinIndexFetchRelationWordResponse, error) {
	return r.client.DouyinIndexGetRelationWordAnalysis(ctx, request)
}

// DouyinIndexFetchPortraitRequest is the request for POST /api/v1/douyin/index/fetch_portrait.
type DouyinIndexFetchPortraitRequest = DouyinIndexGetCrowdPortraitRequest

// DouyinIndexFetchPortraitResponse is the response for POST /api/v1/douyin/index/fetch_portrait.
type DouyinIndexFetchPortraitResponse = DouyinIndexGetCrowdPortraitResponse

// FetchPortrait 获取人群画像/Get crowd portrait
//
// POST /api/v1/douyin/index/fetch_portrait
func (r DouyinIndexResource) FetchPortrait(ctx context.Context, request DouyinIndexFetchPortraitRequest) (*DouyinIndexFetchPortraitResponse, error) {
	return r.client.DouyinIndexGetCrowdPortrait(ctx, request)
}

// DouyinIndexFetchGetUserSubWordResponse is the response for POST /api/v1/douyin/index/fetch_get_user_sub_word.
type DouyinIndexFetchGetUserSubWordResponse = DouyinIndexGetUserSubscribedKeywordsResponse

// FetchGetUserSubWord 获取用户订阅关键词/Get user subscribed keywords
//
// POST /api/v1/douyin/index/fetch_get_user_sub_word
func (r DouyinIndexResource) FetchGetUserSubWord(ctx context.Context) (*DouyinIndexFetchGetUserSubWordResponse, error) {
	return r.client.DouyinIndexGetUserSubscribedKeywords(ctx)
}

// DouyinIndexFetchEncryptUserIDRequest is the request for GET /api/v1/douyin/index/fetch_encrypt_user_id.
type DouyinIndexFetchEncryptUserIDRequest = DouyinIndexEncryptDouyinUIDToUserIDRequest

// DouyinIndexFetchEncryptUserIDResponse is the response for GET /api/v1/douyin/index/fetch_encrypt_user_id.
type DouyinIndexFetchEncryptUserIDResponse = DouyinIndexEncryptDouyinUIDToUserIDResponse

// FetchEncryptUserID 抖音 uid 转加密 user_id/Encrypt Douyin uid to user_id
//
// GET /api/v1/douyin/index/fetch_encrypt_user_id
func (r DouyinIndexResource) FetchEncryptUserID(ctx context.Context, request DouyinIndexFetchEncryptUserIDRequest) (*DouyinIndexFetchEncryptUserIDResponse, error) {
	return r.client.DouyinIndexEncryptDouyinUIDToUserID(ctx, request)
}

// DouyinIndexFetchDarenSugGreatUserListRequest is the request for POST /api/v1/douyin/index/fetch_daren_sug_great_user_list.
type DouyinIndexFetchDarenSugGreatUserListRequest = DouyinIndexDarenSearchSuggestRequest

// DouyinIndexFetchDarenSugGreatUserListResponse is the response for POST /api/v1/douyin/index/fetch_daren_sug_great_user_list.
type DouyinIndexFetchDarenSugGreatUserListResponse = DouyinIndexDarenSearchSuggestResponse

// FetchDarenSugGreatUserList 达人搜索建议/Daren search suggest
//
// POST /api/v1/douyin/index/fetch_daren_sug_great_user_list
func (r DouyinIndexResource) FetchDarenSugGreatUserList(ctx context.Context, request DouyinIndexFetchDarenSugGreatUserListRequest) (*DouyinIndexFetchDarenSugGreatUserListResponse, error) {
	return r.client.DouyinIndexDarenSearchSuggest(ctx, request)
}

// DouyinIndexFetchDarenCompareUsersStableRequest is the request for POST /api/v1/douyin/index/fetch_daren_compare_users_stable.
type DouyinIndexFetchDarenCompareUsersStableRequest = DouyinIndexDarenCompareUsersRequest

// DouyinIndexFetchDarenCompareUsersStableResponse is the response for POST /api/v1/douyin/index/fetch_daren_compare_users_stable.
type DouyinIndexFetchDarenCompareUsersStableResponse = DouyinIndexDarenCompareUsersResponse

// FetchDarenCompareUsersStable 达人趋势对比/Daren compare users
//
// POST /api/v1/douyin/index/fetch_daren_compare_users_stable
func (r DouyinIndexResource) FetchDarenCompareUsersStable(ctx context.Context, request DouyinIndexFetchDarenCompareUsersStableRequest) (*DouyinIndexFetchDarenCompareUsersStableResponse, error) {
	return r.client.DouyinIndexDarenCompareUsers(ctx, request)
}

// DouyinIndexFetchDarenSimilarUsersRequest is the request for POST /api/v1/douyin/index/fetch_daren_similar_users.
type DouyinIndexFetchDarenSimilarUsersRequest = DouyinIndexGetSimilarDarenRequest

// DouyinIndexFetchDarenSimilarUsersResponse is the response for POST /api/v1/douyin/index/fetch_daren_similar_users.
type DouyinIndexFetchDarenSimilarUsersResponse = DouyinIndexGetSimilarDarenResponse

// FetchDarenSimilarUsers 获取相似达人/Get similar daren
//
// POST /api/v1/douyin/index/fetch_daren_similar_users
func (r DouyinIndexResource) FetchDarenSimilarUsers(ctx context.Context, request DouyinIndexFetchDarenSimilarUsersRequest) (*DouyinIndexFetchDarenSimilarUsersResponse, error) {
	return r.client.DouyinIndexGetSimilarDaren(ctx, request)
}

// DouyinIndexFetchDarenGreatUserTopVideoRequest is the request for POST /api/v1/douyin/index/fetch_daren_great_user_top_video.
type DouyinIndexFetchDarenGreatUserTopVideoRequest = DouyinIndexGetDarenTopVideosRequest

// DouyinIndexFetchDarenGreatUserTopVideoResponse is the response for POST /api/v1/douyin/index/fetch_daren_great_user_top_video.
type DouyinIndexFetchDarenGreatUserTopVideoResponse = DouyinIndexGetDarenTopVideosResponse

// FetchDarenGreatUserTopVideo 获取达人视频/Get daren top videos
//
// POST /api/v1/douyin/index/fetch_daren_great_user_top_video
func (r DouyinIndexResource) FetchDarenGreatUserTopVideo(ctx context.Context, request DouyinIndexFetchDarenGreatUserTopVideoRequest) (*DouyinIndexFetchDarenGreatUserTopVideoResponse, error) {
	return r.client.DouyinIndexGetDarenTopVideos(ctx, request)
}

// DouyinIndexFetchDarenGreatItemMileInfoRequest is the request for POST /api/v1/douyin/index/fetch_daren_great_item_mile_info.
type DouyinIndexFetchDarenGreatItemMileInfoRequest = DouyinIndexGetDarenCoreMetricsRequest

// DouyinIndexFetchDarenGreatItemMileInfoResponse is the response for POST /api/v1/douyin/index/fetch_daren_great_item_mile_info.
type DouyinIndexFetchDarenGreatItemMileInfoResponse = DouyinIndexGetDarenCoreMetricsResponse

// FetchDarenGreatItemMileInfo 获取达人核心指标/Get daren core metrics
//
// POST /api/v1/douyin/index/fetch_daren_great_item_mile_info
func (r DouyinIndexResource) FetchDarenGreatItemMileInfo(ctx context.Context, request DouyinIndexFetchDarenGreatItemMileInfoRequest) (*DouyinIndexFetchDarenGreatItemMileInfoResponse, error) {
	return r.client.DouyinIndexGetDarenCoreMetrics(ctx, request)
}

// DouyinIndexFetchDarenGreatUserFansInfoRequest is the request for POST /api/v1/douyin/index/fetch_daren_great_user_fans_info.
type DouyinIndexFetchDarenGreatUserFansInfoRequest = DouyinIndexGetDarenFansAnalysisRequest

// DouyinIndexFetchDarenGreatUserFansInfoResponse is the response for POST /api/v1/douyin/index/fetch_daren_great_user_fans_info.
type DouyinIndexFetchDarenGreatUserFansInfoResponse = DouyinIndexGetDarenFansAnalysisResponse

// FetchDarenGreatUserFansInfo 获取达人粉丝分析/Get daren fans analysis
//
// POST /api/v1/douyin/index/fetch_daren_great_user_fans_info
func (r DouyinIndexResource) FetchDarenGreatUserFansInfo(ctx context.Context, request DouyinIndexFetchDarenGreatUserFansInfoRequest) (*DouyinIndexFetchDarenGreatUserFansInfoResponse, error) {
	return r.client.DouyinIndexGetDarenFansAnalysis(ctx, request)
}

// DouyinIndexFetchItemFilterOptionsResponse is the response for GET /api/v1/douyin/index/fetch_item_filter_options.
type DouyinIndexFetchItemFilterOptionsResponse = DouyinIndexGetVideoSearchFilterOptionsResponse

// FetchItemFilterOptions 获取视频搜索筛选选项/Get video search filter options
//
// GET /api/v1/douyin/index/fetch_item_filter_options
func (r DouyinIndexResource) FetchItemFilterOptions(ctx context.Context) (*DouyinIndexFetchItemFilterOptionsResponse, error) {
	return r.client.DouyinIndexGetVideoSearchFilterOptions(ctx)
}

// DouyinIndexFetchItemSugRequest is the request for POST /api/v1/douyin/index/fetch_item_sug.
type DouyinIndexFetchItemSugRequest = DouyinIndexVideoSearchSuggestRequest

// DouyinIndexFetchItemSugResponse is the response for POST /api/v1/douyin/index/fetch_item_sug.
type DouyinIndexFetchItemSugResponse = DouyinIndexVideoSearchSuggestResponse

// FetchItemSug 视频搜索建议/Video search suggest
//
// POST /api/v1/douyin/index/fetch_item_sug
func (r DouyinIndexResource) FetchItemSug(ctx context.Context, request DouyinIndexFetchItemSugRequest) (*DouyinIndexFetchItemSugResponse, error) {
	return r.client.DouyinIndexVideoSearchSuggest(ctx, request)
}

// DouyinIndexFetchItemQueryRequest is the request for POST /api/v1/douyin/index/fetch_item_query.
type DouyinIndexFetchItemQueryRequest = DouyinIndexVideoSearchResultsRequest

// DouyinIndexFetchItemQueryResponse is the response for POST /api/v1/douyin/index/fetch_item_query.
type DouyinIndexFetchItemQueryResponse = DouyinIndexVideoSearchResultsResponse

// FetchItemQuery 视频搜索结果/Video search results
//
// POST /api/v1/douyin/index/fetch_item_query
func (r DouyinIndexResource) FetchItemQuery(ctx context.Context, request DouyinIndexFetchItemQueryRequest) (*DouyinIndexFetchItemQueryResponse, error) {
	return r.client.DouyinIndexVideoSearchResults(ctx, request)
}

// DouyinIndexFetchBrandSuggestRequest is the request for POST /api/v1/douyin/index/fetch_brand_suggest.
type DouyinIndexFetchBrandSuggestRequest = DouyinIndexBrandSearchSuggestRequest

// DouyinIndexFetchBrandSuggestResponse is the response for POST /api/v1/douyin/index/fetch_brand_suggest.
type DouyinIndexFetchBrandSuggestResponse = DouyinIndexBrandSearchSuggestResponse

// FetchBrandSuggest 品牌搜索建议/Brand search suggest
//
// POST /api/v1/douyin/index/fetch_brand_suggest
func (r DouyinIndexResource) FetchBrandSuggest(ctx context.Context, request DouyinIndexFetchBrandSuggestRequest) (*DouyinIndexFetchBrandSuggestResponse, error) {
	return r.client.DouyinIndexBrandSearchSuggest(ctx, request)
}

// DouyinIndexFetchBrandValidInfoRequest is the request for POST /api/v1/douyin/index/fetch_brand_valid_info.
type DouyinIndexFetchBrandValidInfoRequest = DouyinIndexGetBrandIndexRequest

// DouyinIndexFetchBrandValidInfoResponse is the response for POST /api/v1/douyin/index/fetch_brand_valid_info.
type DouyinIndexFetchBrandValidInfoResponse = DouyinIndexGetBrandIndexResponse

// FetchBrandValidInfo 获取品牌指数/Get brand index
//
// POST /api/v1/douyin/index/fetch_brand_valid_info
func (r DouyinIndexResource) FetchBrandValidInfo(ctx context.Context, request DouyinIndexFetchBrandValidInfoRequest) (*DouyinIndexFetchBrandValidInfoResponse, error) {
	return r.client.DouyinIndexGetBrandIndex(ctx, request)
}

// DouyinIndexFetchBrandRadarChartRequest is the request for POST /api/v1/douyin/index/fetch_brand_radar_chart.
type DouyinIndexFetchBrandRadarChartRequest = DouyinIndexGetBrandRadarChartRequest

// DouyinIndexFetchBrandRadarChartResponse is the response for POST /api/v1/douyin/index/fetch_brand_radar_chart.
type DouyinIndexFetchBrandRadarChartResponse = DouyinIndexGetBrandRadarChartResponse

// FetchBrandRadarChart 获取品牌雷达图/Get brand radar chart
//
// POST /api/v1/douyin/index/fetch_brand_radar_chart
func (r DouyinIndexResource) FetchBrandRadarChart(ctx context.Context, request DouyinIndexFetchBrandRadarChartRequest) (*DouyinIndexFetchBrandRadarChartResponse, error) {
	return r.client.DouyinIndexGetBrandRadarChart(ctx, request)
}

// DouyinIndexFetchBrandLinesRequest is the request for POST /api/v1/douyin/index/fetch_brand_lines.
type DouyinIndexFetchBrandLinesRequest = DouyinIndexGetBrandTrendLinesRequest

// DouyinIndexFetchBrandLinesResponse is the response for POST /api/v1/douyin/index/fetch_brand_lines.
type DouyinIndexFetchBrandLinesResponse = DouyinIndexGetBrandTrendLinesResponse

// FetchBrandLines 获取品牌趋势线/Get brand trend lines
//
// POST /api/v1/douyin/index/fetch_brand_lines
func (r DouyinIndexResource) FetchBrandLines(ctx context.Context, request DouyinIndexFetchBrandLinesRequest) (*DouyinIndexFetchBrandLinesResponse, error) {
	return r.client.DouyinIndexGetBrandTrendLines(ctx, request)
}

// DouyinIndexFetchBrandCyclesRequest is the request for POST /api/v1/douyin/index/fetch_brand_cycles.
type DouyinIndexFetchBrandCyclesRequest = DouyinIndexGetBrandCyclesRequest

// DouyinIndexFetchBrandCyclesResponse is the response for POST /api/v1/douyin/index/fetch_brand_cycles.
type DouyinIndexFetchBrandCyclesResponse = DouyinIndexGetBrandCyclesResponse

// FetchBrandCycles 获取品牌周期数据/Get brand cycles
//
// POST /api/v1/douyin/index/fetch_brand_cycles
func (r DouyinIndexResource) FetchBrandCycles(ctx context.Context, request DouyinIndexFetchBrandCyclesRequest) (*DouyinIndexFetchBrandCyclesResponse, error) {
	return r.client.DouyinIndexGetBrandCycles(ctx, request)
}

// DouyinIndexFetchBrandInitiativeRankWeeklyRequest is the request for POST /api/v1/douyin/index/fetch_brand_initiative_rank_weekly.
type DouyinIndexFetchBrandInitiativeRankWeeklyRequest = DouyinIndexGetBrandInitiativeRankWeeklyRequest

// DouyinIndexFetchBrandInitiativeRankWeeklyResponse is the response for POST /api/v1/douyin/index/fetch_brand_initiative_rank_weekly.
type DouyinIndexFetchBrandInitiativeRankWeeklyResponse = DouyinIndexGetBrandInitiativeRankWeeklyResponse

// FetchBrandInitiativeRankWeekly 获取品牌主动排行周榜/Get brand initiative rank weekly
//
// POST /api/v1/douyin/index/fetch_brand_initiative_rank_weekly
func (r DouyinIndexResource) FetchBrandInitiativeRankWeekly(ctx context.Context, request DouyinIndexFetchBrandInitiativeRankWeeklyRequest) (*DouyinIndexFetchBrandInitiativeRankWeeklyResponse, error) {
	return r.client.DouyinIndexGetBrandInitiativeRankWeekly(ctx, request)
}

// DouyinIndexFetchTopicSuggestRequest is the request for POST /api/v1/douyin/index/fetch_topic_suggest.
type DouyinIndexFetchTopicSuggestRequest = DouyinIndexTopicSearchSuggestRequest

// DouyinIndexFetchTopicSuggestResponse is the response for POST /api/v1/douyin/index/fetch_topic_suggest.
type DouyinIndexFetchTopicSuggestResponse = DouyinIndexTopicSearchSuggestResponse

// FetchTopicSuggest 话题搜索建议/Topic search suggest
//
// POST /api/v1/douyin/index/fetch_topic_suggest
func (r DouyinIndexResource) FetchTopicSuggest(ctx context.Context, request DouyinIndexFetchTopicSuggestRequest) (*DouyinIndexFetchTopicSuggestResponse, error) {
	return r.client.DouyinIndexTopicSearchSuggest(ctx, request)
}

// DouyinIndexFetchTopicQueryRequest is the request for POST /api/v1/douyin/index/fetch_topic_query.
type DouyinIndexFetchTopicQueryRequest = DouyinIndexTopicSearchResultsRequest

// DouyinIndexFetchTopicQueryResponse is the response for POST /api/v1/douyin/index/fetch_topic_query.
type DouyinIndexFetchTopicQueryResponse = DouyinIndexTopicSearchResultsResponse

// FetchTopicQuery 话题搜索结果/Topic search results
//
// POST /api/v1/douyin/index/fetch_topic_query
func (r DouyinIndexResource) FetchTopicQuery(ctx context.Context, request DouyinIndexFetchTopicQueryRequest) (*DouyinIndexFetchTopicQueryResponse, error) {
	return r.client.DouyinIndexTopicSearchResults(ctx, request)
}

// DouyinIndexFetchContentValidDateResponse is the response for GET /api/v1/douyin/index/fetch_content_valid_date.
type DouyinIndexFetchContentValidDateResponse = DouyinIndexGetContentValidDateResponse

// FetchContentValidDate 创作指南有效日期/Get content valid date
//
// GET /api/v1/douyin/index/fetch_content_valid_date
func (r DouyinIndexResource) FetchContentValidDate(ctx context.Context) (*DouyinIndexFetchContentValidDateResponse, error) {
	return r.client.DouyinIndexGetContentValidDate(ctx)
}

// DouyinIndexFetchBrandHotVideosTimeScopeResponse is the response for POST /api/v1/douyin/index/fetch_brand_hot_videos_time_scope.
type DouyinIndexFetchBrandHotVideosTimeScopeResponse = DouyinIndexBrandHotVideosTimeScopeResponse

// FetchBrandHotVideosTimeScope 热门视频时间范围/Brand hot videos time scope
//
// POST /api/v1/douyin/index/fetch_brand_hot_videos_time_scope
func (r DouyinIndexResource) FetchBrandHotVideosTimeScope(ctx context.Context) (*DouyinIndexFetchBrandHotVideosTimeScopeResponse, error) {
	return r.client.DouyinIndexBrandHotVideosTimeScope(ctx)
}

// DouyinIndexFetchContentCreativeKeywordsRequest is the request for POST /api/v1/douyin/index/fetch_content_creative_keywords.
type DouyinIndexFetchContentCreativeKeywordsRequest = DouyinIndexContentCreativeKeywordsRequest

// DouyinIndexFetchContentCreativeKeywordsResponse is the response for POST /api/v1/douyin/index/fetch_content_creative_keywords.
type DouyinIndexFetchContentCreativeKeywordsResponse = DouyinIndexContentCreativeKeywordsResponse

// FetchContentCreativeKeywords 创作热门关键词/Content creative keywords
//
// POST /api/v1/douyin/index/fetch_content_creative_keywords
func (r DouyinIndexResource) FetchContentCreativeKeywords(ctx context.Context, request DouyinIndexFetchContentCreativeKeywordsRequest) (*DouyinIndexFetchContentCreativeKeywordsResponse, error) {
	return r.client.DouyinIndexContentCreativeKeywords(ctx, request)
}

// DouyinIndexFetchContentCreativeKeywordItemsRequest is the request for POST /api/v1/douyin/index/fetch_content_creative_keyword_items.
type DouyinIndexFetchContentCreativeKeywordItemsRequest = DouyinIndexCreativeKeywordRelatedItemsRequest

// DouyinIndexFetchContentCreativeKeywordItemsResponse is the response for POST /api/v1/douyin/index/fetch_content_creative_keyword_items.
type DouyinIndexFetchContentCreativeKeywordItemsResponse = DouyinIndexCreativeKeywordRelatedItemsResponse

// FetchContentCreativeKeywordItems 关键词相关视频/Creative keyword related items
//
// POST /api/v1/douyin/index/fetch_content_creative_keyword_items
func (r DouyinIndexResource) FetchContentCreativeKeywordItems(ctx context.Context, request DouyinIndexFetchContentCreativeKeywordItemsRequest) (*DouyinIndexFetchContentCreativeKeywordItemsResponse, error) {
	return r.client.DouyinIndexCreativeKeywordRelatedItems(ctx, request)
}

// DouyinIndexFetchContentCreativeTopicRequest is the request for POST /api/v1/douyin/index/fetch_content_creative_topic.
type DouyinIndexFetchContentCreativeTopicRequest = DouyinIndexContentCreativeTopicRequest

// DouyinIndexFetchContentCreativeTopicResponse is the response for POST /api/v1/douyin/index/fetch_content_creative_topic.
type DouyinIndexFetchContentCreativeTopicResponse = DouyinIndexContentCreativeTopicResponse

// FetchContentCreativeTopic 创作热门话题/Content creative topic
//
// POST /api/v1/douyin/index/fetch_content_creative_topic
func (r DouyinIndexResource) FetchContentCreativeTopic(ctx context.Context, request DouyinIndexFetchContentCreativeTopicRequest) (*DouyinIndexFetchContentCreativeTopicResponse, error) {
	return r.client.DouyinIndexContentCreativeTopic(ctx, request)
}

// DouyinIndexFetchContentPublishTrendRequest is the request for GET /api/v1/douyin/index/fetch_content_publish_trend.
type DouyinIndexFetchContentPublishTrendRequest = DouyinIndexContentPublishTrendRequest

// DouyinIndexFetchContentPublishTrendResponse is the response for GET /api/v1/douyin/index/fetch_content_publish_trend.
type DouyinIndexFetchContentPublishTrendResponse = DouyinIndexContentPublishTrendResponse

// FetchContentPublishTrend 内容发布趋势/Content publish trend
//
// GET /api/v1/douyin/index/fetch_content_publish_trend
func (r DouyinIndexResource) FetchContentPublishTrend(ctx context.Context, request DouyinIndexFetchContentPublishTrendRequest) (*DouyinIndexFetchContentPublishTrendResponse, error) {
	return r.client.DouyinIndexContentPublishTrend(ctx, request)
}

// DouyinIndexFetchContentCreativeDurationRequest is the request for POST /api/v1/douyin/index/fetch_content_creative_duration.
type DouyinIndexFetchContentCreativeDurationRequest = DouyinIndexContentCreativeDurationRequest

// DouyinIndexFetchContentCreativeDurationResponse is the response for POST /api/v1/douyin/index/fetch_content_creative_duration.
type DouyinIndexFetchContentCreativeDurationResponse = DouyinIndexContentCreativeDurationResponse

// FetchContentCreativeDuration 创作时长分布/Content creative duration
//
// POST /api/v1/douyin/index/fetch_content_creative_duration
func (r DouyinIndexResource) FetchContentCreativeDuration(ctx context.Context, request DouyinIndexFetchContentCreativeDurationRequest) (*DouyinIndexFetchContentCreativeDurationResponse, error) {
	return r.client.DouyinIndexContentCreativeDuration(ctx, request)
}

// DouyinIndexFetchContentAuthorPortraitRequest is the request for POST /api/v1/douyin/index/fetch_content_author_portrait.
type DouyinIndexFetchContentAuthorPortraitRequest = DouyinIndexContentAuthorPortraitRequest

// DouyinIndexFetchContentAuthorPortraitResponse is the response for POST /api/v1/douyin/index/fetch_content_author_portrait.
type DouyinIndexFetchContentAuthorPortraitResponse = DouyinIndexContentAuthorPortraitResponse

// FetchContentAuthorPortrait 创作者画像/Content author portrait
//
// POST /api/v1/douyin/index/fetch_content_author_portrait
func (r DouyinIndexResource) FetchContentAuthorPortrait(ctx context.Context, request DouyinIndexFetchContentAuthorPortraitRequest) (*DouyinIndexFetchContentAuthorPortraitResponse, error) {
	return r.client.DouyinIndexContentAuthorPortrait(ctx, request)
}

// DouyinIndexFetchContentConsumerPortraitRequest is the request for POST /api/v1/douyin/index/fetch_content_consumer_portrait.
type DouyinIndexFetchContentConsumerPortraitRequest = DouyinIndexContentConsumerPortraitRequest

// DouyinIndexFetchContentConsumerPortraitResponse is the response for POST /api/v1/douyin/index/fetch_content_consumer_portrait.
type DouyinIndexFetchContentConsumerPortraitResponse = DouyinIndexContentConsumerPortraitResponse

// FetchContentConsumerPortrait 消费者画像/Content consumer portrait
//
// POST /api/v1/douyin/index/fetch_content_consumer_portrait
func (r DouyinIndexResource) FetchContentConsumerPortrait(ctx context.Context, request DouyinIndexFetchContentConsumerPortraitRequest) (*DouyinIndexFetchContentConsumerPortraitResponse, error) {
	return r.client.DouyinIndexContentConsumerPortrait(ctx, request)
}

// DouyinIndexFetchContentInteractTrendRequest is the request for POST /api/v1/douyin/index/fetch_content_interact_trend.
type DouyinIndexFetchContentInteractTrendRequest = DouyinIndexContentInteractTrendRequest

// DouyinIndexFetchContentInteractTrendResponse is the response for POST /api/v1/douyin/index/fetch_content_interact_trend.
type DouyinIndexFetchContentInteractTrendResponse = DouyinIndexContentInteractTrendResponse

// FetchContentInteractTrend 互动趋势/Content interact trend
//
// POST /api/v1/douyin/index/fetch_content_interact_trend
func (r DouyinIndexResource) FetchContentInteractTrend(ctx context.Context, request DouyinIndexFetchContentInteractTrendRequest) (*DouyinIndexFetchContentInteractTrendResponse, error) {
	return r.client.DouyinIndexContentInteractTrend(ctx, request)
}

// DouyinIndexFetchContentConsumeTrendRequest is the request for POST /api/v1/douyin/index/fetch_content_consume_trend.
type DouyinIndexFetchContentConsumeTrendRequest = DouyinIndexContentConsumeTrendRequest

// DouyinIndexFetchContentConsumeTrendResponse is the response for POST /api/v1/douyin/index/fetch_content_consume_trend.
type DouyinIndexFetchContentConsumeTrendResponse = DouyinIndexContentConsumeTrendResponse

// FetchContentConsumeTrend 消费趋势/Content consume trend
//
// POST /api/v1/douyin/index/fetch_content_consume_trend
func (r DouyinIndexResource) FetchContentConsumeTrend(ctx context.Context, request DouyinIndexFetchContentConsumeTrendRequest) (*DouyinIndexFetchContentConsumeTrendResponse, error) {
	return r.client.DouyinIndexContentConsumeTrend(ctx, request)
}

// DouyinIndexFetchInsightRecommendResponse is the response for GET /api/v1/douyin/index/fetch_insight_recommend.
type DouyinIndexFetchInsightRecommendResponse = DouyinIndexGetRecommendedInsightReportsResponse

// FetchInsightRecommend 获取推荐报告/Get recommended insight reports
//
// GET /api/v1/douyin/index/fetch_insight_recommend
func (r DouyinIndexResource) FetchInsightRecommend(ctx context.Context) (*DouyinIndexFetchInsightRecommendResponse, error) {
	return r.client.DouyinIndexGetRecommendedInsightReports(ctx)
}

// DouyinIndexFetchReportSearchRequest is the request for POST /api/v1/douyin/index/fetch_report_search.
type DouyinIndexFetchReportSearchRequest = DouyinIndexSearchTrendReportsRequest

// DouyinIndexFetchReportSearchResponse is the response for POST /api/v1/douyin/index/fetch_report_search.
type DouyinIndexFetchReportSearchResponse = DouyinIndexSearchTrendReportsResponse

// FetchReportSearch 搜索趋势报告/Search trend reports
//
// POST /api/v1/douyin/index/fetch_report_search
func (r DouyinIndexResource) FetchReportSearch(ctx context.Context, request DouyinIndexFetchReportSearchRequest) (*DouyinIndexFetchReportSearchResponse, error) {
	return r.client.DouyinIndexSearchTrendReports(ctx, request)
}

// DouyinIndexFetchReportDetailRequest is the request for GET /api/v1/douyin/index/fetch_report_detail.
type DouyinIndexFetchReportDetailRequest = DouyinIndexGetReportDetailRequest

// DouyinIndexFetchReportDetailResponse is the response for GET /api/v1/douyin/index/fetch_report_detail.
type DouyinIndexFetchReportDetailResponse = DouyinIndexGetReportDetailResponse

// FetchReportDetail 获取报告详情/Get report detail
//
// GET /api/v1/douyin/index/fetch_report_detail
func (r DouyinIndexResource) FetchReportDetail(ctx context.Context, request DouyinIndexFetchReportDetailRequest) (*DouyinIndexFetchReportDetailResponse, error) {
	return r.client.DouyinIndexGetReportDetail(ctx, request)
}

// DouyinIndexFetchInsightGetRecRequest is the request for GET /api/v1/douyin/index/fetch_insight_get_rec.
type DouyinIndexFetchInsightGetRecRequest = DouyinIndexGetRelatedInsightRecommendationsRequest

// DouyinIndexFetchInsightGetRecResponse is the response for GET /api/v1/douyin/index/fetch_insight_get_rec.
type DouyinIndexFetchInsightGetRecResponse = DouyinIndexGetRelatedInsightRecommendationsResponse

// FetchInsightGetRec 获取报告相关推荐/Get related insight recommendations
//
// GET /api/v1/douyin/index/fetch_insight_get_rec
func (r DouyinIndexResource) FetchInsightGetRec(ctx context.Context, request DouyinIndexFetchInsightGetRecRequest) (*DouyinIndexFetchInsightGetRecResponse, error) {
	return r.client.DouyinIndexGetRelatedInsightRecommendations(ctx, request)
}

// DouyinSearchResource contains endpoints from the Douyin-Search-API tag.
type DouyinSearchResource struct {
	client *Client
}

// FetchGeneralSearchV1 获取综合搜索 V1/Fetch general search V1
//
// POST /api/v1/douyin/search/fetch_general_search_v1
func (r DouyinSearchResource) FetchGeneralSearchV1(ctx context.Context, request DouyinSearchFetchGeneralSearchV1Request) (*DouyinSearchFetchGeneralSearchV1Response, error) {
	return r.client.DouyinSearchFetchGeneralSearchV1(ctx, request)
}

// FetchGeneralSearchV2 获取综合搜索 V2/Fetch general search V2
//
// POST /api/v1/douyin/search/fetch_general_search_v2
func (r DouyinSearchResource) FetchGeneralSearchV2(ctx context.Context, request DouyinSearchFetchGeneralSearchV2Request) (*DouyinSearchFetchGeneralSearchV2Response, error) {
	return r.client.DouyinSearchFetchGeneralSearchV2(ctx, request)
}

// DouyinSearchFetchSearchSuggestRequest is the request for POST /api/v1/douyin/search/fetch_search_suggest.
type DouyinSearchFetchSearchSuggestRequest = DouyinSearchFetchSearchKeywordSuggestionsRequest

// DouyinSearchFetchSearchSuggestResponse is the response for POST /api/v1/douyin/search/fetch_search_suggest.
type DouyinSearchFetchSearchSuggestResponse = DouyinSearchFetchSearchKeywordSuggestionsResponse

// FetchSearchSuggest 获取搜索关键词推荐/Fetch search keyword suggestions
//
// POST /api/v1/douyin/search/fetch_search_suggest
func (r DouyinSearchResource) FetchSearchSuggest(ctx context.Context, request DouyinSearchFetchSearchSuggestRequest) (*DouyinSearchFetchSearchSuggestResponse, error) {
	return r.client.DouyinSearchFetchSearchKeywordSuggestions(ctx, request)
}

// FetchVideoSearchV1 获取视频搜索 V1/Fetch video search V1
//
// POST /api/v1/douyin/search/fetch_video_search_v1
func (r DouyinSearchResource) FetchVideoSearchV1(ctx context.Context, request DouyinSearchFetchVideoSearchV1Request) (*DouyinSearchFetchVideoSearchV1Response, error) {
	return r.client.DouyinSearchFetchVideoSearchV1(ctx, request)
}

// FetchVideoSearchV2 获取视频搜索 V2/Fetch video search V2
//
// POST /api/v1/douyin/search/fetch_video_search_v2
func (r DouyinSearchResource) FetchVideoSearchV2(ctx context.Context, request DouyinSearchFetchVideoSearchV2Request) (*DouyinSearchFetchVideoSearchV2Response, error) {
	return r.client.DouyinSearchFetchVideoSearchV2(ctx, request)
}

// DouyinSearchFetchMultiSearchRequest is the request for POST /api/v1/douyin/search/fetch_multi_search.
type DouyinSearchFetchMultiSearchRequest = DouyinSearchFetchMultiTypeSearchRequest

// DouyinSearchFetchMultiSearchResponse is the response for POST /api/v1/douyin/search/fetch_multi_search.
type DouyinSearchFetchMultiSearchResponse = DouyinSearchFetchMultiTypeSearchResponse

// FetchMultiSearch 获取多重搜索/Fetch multi-type search
//
// POST /api/v1/douyin/search/fetch_multi_search
func (r DouyinSearchResource) FetchMultiSearch(ctx context.Context, request DouyinSearchFetchMultiSearchRequest) (*DouyinSearchFetchMultiSearchResponse, error) {
	return r.client.DouyinSearchFetchMultiTypeSearch(ctx, request)
}

// FetchUserSearch 获取用户搜索/Fetch user search
//
// POST /api/v1/douyin/search/fetch_user_search
func (r DouyinSearchResource) FetchUserSearch(ctx context.Context, request DouyinSearchFetchUserSearchRequest) (*DouyinSearchFetchUserSearchResponse, error) {
	return r.client.DouyinSearchFetchUserSearch(ctx, request)
}

// FetchUserSearchV2 获取用户搜索 V2/Fetch user search V2
//
// POST /api/v1/douyin/search/fetch_user_search_v2
func (r DouyinSearchResource) FetchUserSearchV2(ctx context.Context, request DouyinSearchFetchUserSearchV2Request) (*DouyinSearchFetchUserSearchV2Response, error) {
	return r.client.DouyinSearchFetchUserSearchV2(ctx, request)
}

// FetchImageSearch 获取图片搜索/Fetch image search
//
// POST /api/v1/douyin/search/fetch_image_search
func (r DouyinSearchResource) FetchImageSearch(ctx context.Context, request DouyinSearchFetchImageSearchRequest) (*DouyinSearchFetchImageSearchResponse, error) {
	return r.client.DouyinSearchFetchImageSearch(ctx, request)
}

// DouyinSearchFetchImageSearchV3Request is the request for POST /api/v1/douyin/search/fetch_image_search_v3.
type DouyinSearchFetchImageSearchV3Request = DouyinSearchFetchImageTextSearchV3Request

// DouyinSearchFetchImageSearchV3Response is the response for POST /api/v1/douyin/search/fetch_image_search_v3.
type DouyinSearchFetchImageSearchV3Response = DouyinSearchFetchImageTextSearchV3Response

// FetchImageSearchV3 获取图文搜索 V3/Fetch image-text search V3
//
// POST /api/v1/douyin/search/fetch_image_search_v3
func (r DouyinSearchResource) FetchImageSearchV3(ctx context.Context, request DouyinSearchFetchImageSearchV3Request) (*DouyinSearchFetchImageSearchV3Response, error) {
	return r.client.DouyinSearchFetchImageTextSearchV3(ctx, request)
}

// FetchLiveSearchV1 获取直播搜索 V1/Fetch live search V1
//
// POST /api/v1/douyin/search/fetch_live_search_v1
func (r DouyinSearchResource) FetchLiveSearchV1(ctx context.Context, request DouyinSearchFetchLiveSearchV1Request) (*DouyinSearchFetchLiveSearchV1Response, error) {
	return r.client.DouyinSearchFetchLiveSearchV1(ctx, request)
}

// DouyinSearchFetchChallengeSearchV1Request is the request for POST /api/v1/douyin/search/fetch_challenge_search_v1.
type DouyinSearchFetchChallengeSearchV1Request = DouyinSearchFetchHashtagSearchV1Request

// DouyinSearchFetchChallengeSearchV1Response is the response for POST /api/v1/douyin/search/fetch_challenge_search_v1.
type DouyinSearchFetchChallengeSearchV1Response = DouyinSearchFetchHashtagSearchV1Response

// FetchChallengeSearchV1 获取话题搜索 V1/Fetch hashtag search V1
//
// POST /api/v1/douyin/search/fetch_challenge_search_v1
func (r DouyinSearchResource) FetchChallengeSearchV1(ctx context.Context, request DouyinSearchFetchChallengeSearchV1Request) (*DouyinSearchFetchChallengeSearchV1Response, error) {
	return r.client.DouyinSearchFetchHashtagSearchV1(ctx, request)
}

// DouyinSearchFetchChallengeSearchV2Request is the request for POST /api/v1/douyin/search/fetch_challenge_search_v2.
type DouyinSearchFetchChallengeSearchV2Request = DouyinSearchFetchHashtagSearchV2Request

// DouyinSearchFetchChallengeSearchV2Response is the response for POST /api/v1/douyin/search/fetch_challenge_search_v2.
type DouyinSearchFetchChallengeSearchV2Response = DouyinSearchFetchHashtagSearchV2Response

// FetchChallengeSearchV2 获取话题搜索 V2/Fetch hashtag search V2
//
// POST /api/v1/douyin/search/fetch_challenge_search_v2
func (r DouyinSearchResource) FetchChallengeSearchV2(ctx context.Context, request DouyinSearchFetchChallengeSearchV2Request) (*DouyinSearchFetchChallengeSearchV2Response, error) {
	return r.client.DouyinSearchFetchHashtagSearchV2(ctx, request)
}

// DouyinSearchFetchChallengeSuggestRequest is the request for POST /api/v1/douyin/search/fetch_challenge_suggest.
type DouyinSearchFetchChallengeSuggestRequest = DouyinSearchFetchHashtagSuggestionsRequest

// DouyinSearchFetchChallengeSuggestResponse is the response for POST /api/v1/douyin/search/fetch_challenge_suggest.
type DouyinSearchFetchChallengeSuggestResponse = DouyinSearchFetchHashtagSuggestionsResponse

// FetchChallengeSuggest 获取话题推荐搜索/Fetch hashtag suggestions
//
// POST /api/v1/douyin/search/fetch_challenge_suggest
func (r DouyinSearchResource) FetchChallengeSuggest(ctx context.Context, request DouyinSearchFetchChallengeSuggestRequest) (*DouyinSearchFetchChallengeSuggestResponse, error) {
	return r.client.DouyinSearchFetchHashtagSuggestions(ctx, request)
}

// FetchExperienceSearch 获取经验搜索/Fetch experience search
//
// POST /api/v1/douyin/search/fetch_experience_search
func (r DouyinSearchResource) FetchExperienceSearch(ctx context.Context, request DouyinSearchFetchExperienceSearchRequest) (*DouyinSearchFetchExperienceSearchResponse, error) {
	return r.client.DouyinSearchFetchExperienceSearch(ctx, request)
}

// FetchMusicSearch 获取音乐搜索/Fetch music search
//
// POST /api/v1/douyin/search/fetch_music_search
func (r DouyinSearchResource) FetchMusicSearch(ctx context.Context, request DouyinSearchFetchMusicSearchRequest) (*DouyinSearchFetchMusicSearchResponse, error) {
	return r.client.DouyinSearchFetchMusicSearch(ctx, request)
}

// DouyinSearchFetchDiscussSearchRequest is the request for POST /api/v1/douyin/search/fetch_discuss_search.
type DouyinSearchFetchDiscussSearchRequest = DouyinSearchFetchDiscussionSearchRequest

// DouyinSearchFetchDiscussSearchResponse is the response for POST /api/v1/douyin/search/fetch_discuss_search.
type DouyinSearchFetchDiscussSearchResponse = DouyinSearchFetchDiscussionSearchResponse

// FetchDiscussSearch 获取讨论搜索/Fetch discussion search
//
// POST /api/v1/douyin/search/fetch_discuss_search
func (r DouyinSearchResource) FetchDiscussSearch(ctx context.Context, request DouyinSearchFetchDiscussSearchRequest) (*DouyinSearchFetchDiscussSearchResponse, error) {
	return r.client.DouyinSearchFetchDiscussionSearch(ctx, request)
}

// FetchSchoolSearch 获取学校搜索/Fetch school search
//
// POST /api/v1/douyin/search/fetch_school_search
func (r DouyinSearchResource) FetchSchoolSearch(ctx context.Context, request DouyinSearchFetchSchoolSearchRequest) (*DouyinSearchFetchSchoolSearchResponse, error) {
	return r.client.DouyinSearchFetchSchoolSearch(ctx, request)
}

// FetchVisionSearch 获取图像识别搜索/Fetch vision search (image-based search)
//
// POST /api/v1/douyin/search/fetch_vision_search
func (r DouyinSearchResource) FetchVisionSearch(ctx context.Context, request DouyinSearchFetchVisionSearchRequest) (*DouyinSearchFetchVisionSearchResponse, error) {
	return r.client.DouyinSearchFetchVisionSearch(ctx, request)
}

// DouyinBillboardResource contains endpoints from the Douyin-Billboard-API tag.
type DouyinBillboardResource struct {
	client *Client
}

// DouyinBillboardFetchCityListResponse is the response for GET /api/v1/douyin/billboard/fetch_city_list.
type DouyinBillboardFetchCityListResponse = DouyinBillboardFetchChineseCityListResponse

// FetchCityList 获取中国城市列表/Fetch Chinese city list
//
// GET /api/v1/douyin/billboard/fetch_city_list
func (r DouyinBillboardResource) FetchCityList(ctx context.Context) (*DouyinBillboardFetchCityListResponse, error) {
	return r.client.DouyinBillboardFetchChineseCityList(ctx)
}

// DouyinBillboardFetchContentTagResponse is the response for GET /api/v1/douyin/billboard/fetch_content_tag.
type DouyinBillboardFetchContentTagResponse = DouyinBillboardFetchVerticalContentTagsResponse

// FetchContentTag 获取垂类内容标签/Fetch vertical content tags
//
// GET /api/v1/douyin/billboard/fetch_content_tag
func (r DouyinBillboardResource) FetchContentTag(ctx context.Context) (*DouyinBillboardFetchContentTagResponse, error) {
	return r.client.DouyinBillboardFetchVerticalContentTags(ctx)
}

// DouyinBillboardFetchHotCategoryListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_category_list.
type DouyinBillboardFetchHotCategoryListRequest = DouyinBillboardFetchHotListCategoryRequest

// DouyinBillboardFetchHotCategoryListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_category_list.
type DouyinBillboardFetchHotCategoryListResponse = DouyinBillboardFetchHotListCategoryResponse

// FetchHotCategoryList 获取热点榜分类/Fetch hot list category
//
// GET /api/v1/douyin/billboard/fetch_hot_category_list
func (r DouyinBillboardResource) FetchHotCategoryList(ctx context.Context, request DouyinBillboardFetchHotCategoryListRequest) (*DouyinBillboardFetchHotCategoryListResponse, error) {
	return r.client.DouyinBillboardFetchHotListCategory(ctx, request)
}

// DouyinBillboardFetchHotRiseListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_rise_list.
type DouyinBillboardFetchHotRiseListRequest = DouyinBillboardFetchRisingHotListRequest

// DouyinBillboardFetchHotRiseListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_rise_list.
type DouyinBillboardFetchHotRiseListResponse = DouyinBillboardFetchRisingHotListResponse

// FetchHotRiseList 获取上升热点榜/Fetch rising hot list
//
// GET /api/v1/douyin/billboard/fetch_hot_rise_list
func (r DouyinBillboardResource) FetchHotRiseList(ctx context.Context, request DouyinBillboardFetchHotRiseListRequest) (*DouyinBillboardFetchHotRiseListResponse, error) {
	return r.client.DouyinBillboardFetchRisingHotList(ctx, request)
}

// DouyinBillboardFetchHotCityListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_city_list.
type DouyinBillboardFetchHotCityListRequest = DouyinBillboardFetchCityHotListRequest

// DouyinBillboardFetchHotCityListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_city_list.
type DouyinBillboardFetchHotCityListResponse = DouyinBillboardFetchCityHotListResponse

// FetchHotCityList 获取同城热点榜/Fetch city hot list
//
// GET /api/v1/douyin/billboard/fetch_hot_city_list
func (r DouyinBillboardResource) FetchHotCityList(ctx context.Context, request DouyinBillboardFetchHotCityListRequest) (*DouyinBillboardFetchHotCityListResponse, error) {
	return r.client.DouyinBillboardFetchCityHotList(ctx, request)
}

// FetchHotChallengeList 获取挑战热榜/Fetch hot challenge list
//
// GET /api/v1/douyin/billboard/fetch_hot_challenge_list
func (r DouyinBillboardResource) FetchHotChallengeList(ctx context.Context, request DouyinBillboardFetchHotChallengeListRequest) (*DouyinBillboardFetchHotChallengeListResponse, error) {
	return r.client.DouyinBillboardFetchHotChallengeList(ctx, request)
}

// DouyinBillboardFetchHotTotalListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_total_list.
type DouyinBillboardFetchHotTotalListRequest = DouyinBillboardFetchTotalHotListRequest

// DouyinBillboardFetchHotTotalListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_total_list.
type DouyinBillboardFetchHotTotalListResponse = DouyinBillboardFetchTotalHotListResponse

// FetchHotTotalList 获取热点总榜/Fetch total hot list
//
// GET /api/v1/douyin/billboard/fetch_hot_total_list
func (r DouyinBillboardResource) FetchHotTotalList(ctx context.Context, request DouyinBillboardFetchHotTotalListRequest) (*DouyinBillboardFetchHotTotalListResponse, error) {
	return r.client.DouyinBillboardFetchTotalHotList(ctx, request)
}

// DouyinBillboardFetchHotCalendarListRequest is the request for POST /api/v1/douyin/billboard/fetch_hot_calendar_list.
type DouyinBillboardFetchHotCalendarListRequest = DouyinBillboardFetchActivityCalendarRequest

// DouyinBillboardFetchHotCalendarListResponse is the response for POST /api/v1/douyin/billboard/fetch_hot_calendar_list.
type DouyinBillboardFetchHotCalendarListResponse = DouyinBillboardFetchActivityCalendarResponse

// FetchHotCalendarList 获取活动日历/Fetch activity calendar
//
// POST /api/v1/douyin/billboard/fetch_hot_calendar_list
func (r DouyinBillboardResource) FetchHotCalendarList(ctx context.Context, request DouyinBillboardFetchHotCalendarListRequest) (*DouyinBillboardFetchHotCalendarListResponse, error) {
	return r.client.DouyinBillboardFetchActivityCalendar(ctx, request)
}

// DouyinBillboardFetchHotCalendarDetailRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_calendar_detail.
type DouyinBillboardFetchHotCalendarDetailRequest = DouyinBillboardFetchActivityCalendarDetailRequest

// DouyinBillboardFetchHotCalendarDetailResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_calendar_detail.
type DouyinBillboardFetchHotCalendarDetailResponse = DouyinBillboardFetchActivityCalendarDetailResponse

// FetchHotCalendarDetail 获取活动日历详情/Fetch activity calendar detail
//
// GET /api/v1/douyin/billboard/fetch_hot_calendar_detail
func (r DouyinBillboardResource) FetchHotCalendarDetail(ctx context.Context, request DouyinBillboardFetchHotCalendarDetailRequest) (*DouyinBillboardFetchHotCalendarDetailResponse, error) {
	return r.client.DouyinBillboardFetchActivityCalendarDetail(ctx, request)
}

// DouyinBillboardFetchHotUserPortraitListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_user_portrait_list.
type DouyinBillboardFetchHotUserPortraitListRequest = DouyinBillboardFetchWorkLikeAudiencePortraitHotListOnlyRequest

// DouyinBillboardFetchHotUserPortraitListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_user_portrait_list.
type DouyinBillboardFetchHotUserPortraitListResponse = DouyinBillboardFetchWorkLikeAudiencePortraitHotListOnlyResponse

// FetchHotUserPortraitList 获取作品点赞观众画像-仅限热门榜/Fetch work like audience portrait - hot list only
//
// GET /api/v1/douyin/billboard/fetch_hot_user_portrait_list
func (r DouyinBillboardResource) FetchHotUserPortraitList(ctx context.Context, request DouyinBillboardFetchHotUserPortraitListRequest) (*DouyinBillboardFetchHotUserPortraitListResponse, error) {
	return r.client.DouyinBillboardFetchWorkLikeAudiencePortraitHotListOnly(ctx, request)
}

// DouyinBillboardFetchHotCommentWordListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_comment_word_list.
type DouyinBillboardFetchHotCommentWordListRequest = DouyinBillboardFetchWorkCommentAnalysisWordCloudWeightRequest

// DouyinBillboardFetchHotCommentWordListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_comment_word_list.
type DouyinBillboardFetchHotCommentWordListResponse = DouyinBillboardFetchWorkCommentAnalysisWordCloudWeightResponse

// FetchHotCommentWordList 获取作品评论分析-词云权重/Fetch work comment analysis word cloud weight
//
// GET /api/v1/douyin/billboard/fetch_hot_comment_word_list
func (r DouyinBillboardResource) FetchHotCommentWordList(ctx context.Context, request DouyinBillboardFetchHotCommentWordListRequest) (*DouyinBillboardFetchHotCommentWordListResponse, error) {
	return r.client.DouyinBillboardFetchWorkCommentAnalysisWordCloudWeight(ctx, request)
}

// DouyinBillboardFetchHotItemTrendsListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_item_trends_list.
type DouyinBillboardFetchHotItemTrendsListRequest = DouyinBillboardFetchPostDataTrendRequest

// DouyinBillboardFetchHotItemTrendsListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_item_trends_list.
type DouyinBillboardFetchHotItemTrendsListResponse = DouyinBillboardFetchPostDataTrendResponse

// FetchHotItemTrendsList 获取作品数据趋势/Fetch post data trend
//
// GET /api/v1/douyin/billboard/fetch_hot_item_trends_list
func (r DouyinBillboardResource) FetchHotItemTrendsList(ctx context.Context, request DouyinBillboardFetchHotItemTrendsListRequest) (*DouyinBillboardFetchHotItemTrendsListResponse, error) {
	return r.client.DouyinBillboardFetchPostDataTrend(ctx, request)
}

// FetchHotAccountList 获取热门账号/Fetch hot account list
//
// POST /api/v1/douyin/billboard/fetch_hot_account_list
func (r DouyinBillboardResource) FetchHotAccountList(ctx context.Context, request DouyinBillboardFetchHotAccountListRequest) (*DouyinBillboardFetchHotAccountListResponse, error) {
	return r.client.DouyinBillboardFetchHotAccountList(ctx, request)
}

// DouyinBillboardFetchHotAccountSearchListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_account_search_list.
type DouyinBillboardFetchHotAccountSearchListRequest = DouyinBillboardFetchAccountSearchListRequest

// DouyinBillboardFetchHotAccountSearchListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_account_search_list.
type DouyinBillboardFetchHotAccountSearchListResponse = DouyinBillboardFetchAccountSearchListResponse

// FetchHotAccountSearchList 搜索用户名或抖音号/Fetch account search list
//
// GET /api/v1/douyin/billboard/fetch_hot_account_search_list
func (r DouyinBillboardResource) FetchHotAccountSearchList(ctx context.Context, request DouyinBillboardFetchHotAccountSearchListRequest) (*DouyinBillboardFetchHotAccountSearchListResponse, error) {
	return r.client.DouyinBillboardFetchAccountSearchList(ctx, request)
}

// DouyinBillboardFetchHotAccountTrendsListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_account_trends_list.
type DouyinBillboardFetchHotAccountTrendsListRequest = DouyinBillboardFetchAccountFanDataTrendRequest

// DouyinBillboardFetchHotAccountTrendsListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_account_trends_list.
type DouyinBillboardFetchHotAccountTrendsListResponse = DouyinBillboardFetchAccountFanDataTrendResponse

// FetchHotAccountTrendsList 获取账号粉丝数据趋势/Fetch account fan data trend
//
// GET /api/v1/douyin/billboard/fetch_hot_account_trends_list
func (r DouyinBillboardResource) FetchHotAccountTrendsList(ctx context.Context, request DouyinBillboardFetchHotAccountTrendsListRequest) (*DouyinBillboardFetchHotAccountTrendsListResponse, error) {
	return r.client.DouyinBillboardFetchAccountFanDataTrend(ctx, request)
}

// DouyinBillboardFetchHotAccountItemAnalysisListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_account_item_analysis_list.
type DouyinBillboardFetchHotAccountItemAnalysisListRequest = DouyinBillboardFetchAccountWorkAnalysisLastWeekRequest

// DouyinBillboardFetchHotAccountItemAnalysisListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_account_item_analysis_list.
type DouyinBillboardFetchHotAccountItemAnalysisListResponse = DouyinBillboardFetchAccountWorkAnalysisLastWeekResponse

// FetchHotAccountItemAnalysisList 获取账号作品分析-上周/Fetch account work analysis - last week
//
// GET /api/v1/douyin/billboard/fetch_hot_account_item_analysis_list
func (r DouyinBillboardResource) FetchHotAccountItemAnalysisList(ctx context.Context, request DouyinBillboardFetchHotAccountItemAnalysisListRequest) (*DouyinBillboardFetchHotAccountItemAnalysisListResponse, error) {
	return r.client.DouyinBillboardFetchAccountWorkAnalysisLastWeek(ctx, request)
}

// DouyinBillboardFetchHotAccountFansPortraitListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_account_fans_portrait_list.
type DouyinBillboardFetchHotAccountFansPortraitListRequest = DouyinBillboardFetchFanPortraitRequest

// DouyinBillboardFetchHotAccountFansPortraitListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_account_fans_portrait_list.
type DouyinBillboardFetchHotAccountFansPortraitListResponse = DouyinBillboardFetchFanPortraitResponse

// FetchHotAccountFansPortraitList 获取粉丝画像/Fetch fan portrait
//
// GET /api/v1/douyin/billboard/fetch_hot_account_fans_portrait_list
func (r DouyinBillboardResource) FetchHotAccountFansPortraitList(ctx context.Context, request DouyinBillboardFetchHotAccountFansPortraitListRequest) (*DouyinBillboardFetchHotAccountFansPortraitListResponse, error) {
	return r.client.DouyinBillboardFetchFanPortrait(ctx, request)
}

// DouyinBillboardFetchHotAccountFansInterestAccountListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_account_fans_interest_account_list.
type DouyinBillboardFetchHotAccountFansInterestAccountListRequest = DouyinBillboardFetchFanInterestAuthor20UsersRequest

// DouyinBillboardFetchHotAccountFansInterestAccountListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_account_fans_interest_account_list.
type DouyinBillboardFetchHotAccountFansInterestAccountListResponse = DouyinBillboardFetchFanInterestAuthor20UsersResponse

// FetchHotAccountFansInterestAccountList 获取粉丝兴趣作者 20个用户/Fetch fan interest author 20 users
//
// GET /api/v1/douyin/billboard/fetch_hot_account_fans_interest_account_list
func (r DouyinBillboardResource) FetchHotAccountFansInterestAccountList(ctx context.Context, request DouyinBillboardFetchHotAccountFansInterestAccountListRequest) (*DouyinBillboardFetchHotAccountFansInterestAccountListResponse, error) {
	return r.client.DouyinBillboardFetchFanInterestAuthor20Users(ctx, request)
}

// DouyinBillboardFetchHotAccountFansInterestTopicListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_account_fans_interest_topic_list.
type DouyinBillboardFetchHotAccountFansInterestTopicListRequest = DouyinBillboardFetchFanInterestTopicInTheLast3Days10TopicsRequest

// DouyinBillboardFetchHotAccountFansInterestTopicListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_account_fans_interest_topic_list.
type DouyinBillboardFetchHotAccountFansInterestTopicListResponse = DouyinBillboardFetchFanInterestTopicInTheLast3Days10TopicsResponse

// FetchHotAccountFansInterestTopicList 获取粉丝近3天感兴趣的话题 10个话题/Fetch fan interest topic in the last 3 days 10 topics
//
// GET /api/v1/douyin/billboard/fetch_hot_account_fans_interest_topic_list
func (r DouyinBillboardResource) FetchHotAccountFansInterestTopicList(ctx context.Context, request DouyinBillboardFetchHotAccountFansInterestTopicListRequest) (*DouyinBillboardFetchHotAccountFansInterestTopicListResponse, error) {
	return r.client.DouyinBillboardFetchFanInterestTopicInTheLast3Days10Topics(ctx, request)
}

// DouyinBillboardFetchHotAccountFansInterestSearchListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_account_fans_interest_search_list.
type DouyinBillboardFetchHotAccountFansInterestSearchListRequest = DouyinBillboardFetchFanInterestSearchTermInTheLast3Days10SearchTermsRequest

// DouyinBillboardFetchHotAccountFansInterestSearchListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_account_fans_interest_search_list.
type DouyinBillboardFetchHotAccountFansInterestSearchListResponse = DouyinBillboardFetchFanInterestSearchTermInTheLast3Days10SearchTermsResponse

// FetchHotAccountFansInterestSearchList 获取粉丝近3天搜索词 10个搜索词/Fetch fan interest search term in the last 3 days 10 search terms
//
// GET /api/v1/douyin/billboard/fetch_hot_account_fans_interest_search_list
func (r DouyinBillboardResource) FetchHotAccountFansInterestSearchList(ctx context.Context, request DouyinBillboardFetchHotAccountFansInterestSearchListRequest) (*DouyinBillboardFetchHotAccountFansInterestSearchListResponse, error) {
	return r.client.DouyinBillboardFetchFanInterestSearchTermInTheLast3Days10SearchTerms(ctx, request)
}

// DouyinBillboardFetchHotTotalVideoListRequest is the request for POST /api/v1/douyin/billboard/fetch_hot_total_video_list.
type DouyinBillboardFetchHotTotalVideoListRequest = DouyinBillboardFetchVideoHotListRequest

// DouyinBillboardFetchHotTotalVideoListResponse is the response for POST /api/v1/douyin/billboard/fetch_hot_total_video_list.
type DouyinBillboardFetchHotTotalVideoListResponse = DouyinBillboardFetchVideoHotListResponse

// FetchHotTotalVideoList 获取视频热榜/Fetch video hot list
//
// POST /api/v1/douyin/billboard/fetch_hot_total_video_list
func (r DouyinBillboardResource) FetchHotTotalVideoList(ctx context.Context, request DouyinBillboardFetchHotTotalVideoListRequest) (*DouyinBillboardFetchHotTotalVideoListResponse, error) {
	return r.client.DouyinBillboardFetchVideoHotList(ctx, request)
}

// DouyinBillboardFetchHotTotalLowFanListRequest is the request for POST /api/v1/douyin/billboard/fetch_hot_total_low_fan_list.
type DouyinBillboardFetchHotTotalLowFanListRequest = DouyinBillboardFetchLowFanExplosionListRequest

// DouyinBillboardFetchHotTotalLowFanListResponse is the response for POST /api/v1/douyin/billboard/fetch_hot_total_low_fan_list.
type DouyinBillboardFetchHotTotalLowFanListResponse = DouyinBillboardFetchLowFanExplosionListResponse

// FetchHotTotalLowFanList 获取低粉爆款榜/Fetch low fan explosion list
//
// POST /api/v1/douyin/billboard/fetch_hot_total_low_fan_list
func (r DouyinBillboardResource) FetchHotTotalLowFanList(ctx context.Context, request DouyinBillboardFetchHotTotalLowFanListRequest) (*DouyinBillboardFetchHotTotalLowFanListResponse, error) {
	return r.client.DouyinBillboardFetchLowFanExplosionList(ctx, request)
}

// DouyinBillboardFetchHotTotalHighPlayListRequest is the request for POST /api/v1/douyin/billboard/fetch_hot_total_high_play_list.
type DouyinBillboardFetchHotTotalHighPlayListRequest = DouyinBillboardFetchHighCompletionRateListRequest

// DouyinBillboardFetchHotTotalHighPlayListResponse is the response for POST /api/v1/douyin/billboard/fetch_hot_total_high_play_list.
type DouyinBillboardFetchHotTotalHighPlayListResponse = DouyinBillboardFetchHighCompletionRateListResponse

// FetchHotTotalHighPlayList 获取高完播率榜/Fetch high completion rate list
//
// POST /api/v1/douyin/billboard/fetch_hot_total_high_play_list
func (r DouyinBillboardResource) FetchHotTotalHighPlayList(ctx context.Context, request DouyinBillboardFetchHotTotalHighPlayListRequest) (*DouyinBillboardFetchHotTotalHighPlayListResponse, error) {
	return r.client.DouyinBillboardFetchHighCompletionRateList(ctx, request)
}

// DouyinBillboardFetchHotTotalHighLikeListRequest is the request for POST /api/v1/douyin/billboard/fetch_hot_total_high_like_list.
type DouyinBillboardFetchHotTotalHighLikeListRequest = DouyinBillboardFetchHighLikeRateListRequest

// DouyinBillboardFetchHotTotalHighLikeListResponse is the response for POST /api/v1/douyin/billboard/fetch_hot_total_high_like_list.
type DouyinBillboardFetchHotTotalHighLikeListResponse = DouyinBillboardFetchHighLikeRateListResponse

// FetchHotTotalHighLikeList 获取高点赞率榜/Fetch high like rate list
//
// POST /api/v1/douyin/billboard/fetch_hot_total_high_like_list
func (r DouyinBillboardResource) FetchHotTotalHighLikeList(ctx context.Context, request DouyinBillboardFetchHotTotalHighLikeListRequest) (*DouyinBillboardFetchHotTotalHighLikeListResponse, error) {
	return r.client.DouyinBillboardFetchHighLikeRateList(ctx, request)
}

// DouyinBillboardFetchHotTotalHighFanListRequest is the request for POST /api/v1/douyin/billboard/fetch_hot_total_high_fan_list.
type DouyinBillboardFetchHotTotalHighFanListRequest = DouyinBillboardFetchHighFanRateListRequest

// DouyinBillboardFetchHotTotalHighFanListResponse is the response for POST /api/v1/douyin/billboard/fetch_hot_total_high_fan_list.
type DouyinBillboardFetchHotTotalHighFanListResponse = DouyinBillboardFetchHighFanRateListResponse

// FetchHotTotalHighFanList 获取高涨粉率榜/Fetch high fan rate list
//
// POST /api/v1/douyin/billboard/fetch_hot_total_high_fan_list
func (r DouyinBillboardResource) FetchHotTotalHighFanList(ctx context.Context, request DouyinBillboardFetchHotTotalHighFanListRequest) (*DouyinBillboardFetchHotTotalHighFanListResponse, error) {
	return r.client.DouyinBillboardFetchHighFanRateList(ctx, request)
}

// DouyinBillboardFetchHotTotalTopicListRequest is the request for POST /api/v1/douyin/billboard/fetch_hot_total_topic_list.
type DouyinBillboardFetchHotTotalTopicListRequest = DouyinBillboardFetchTopicHotListRequest

// DouyinBillboardFetchHotTotalTopicListResponse is the response for POST /api/v1/douyin/billboard/fetch_hot_total_topic_list.
type DouyinBillboardFetchHotTotalTopicListResponse = DouyinBillboardFetchTopicHotListResponse

// FetchHotTotalTopicList 获取话题热榜/Fetch topic hot list
//
// POST /api/v1/douyin/billboard/fetch_hot_total_topic_list
func (r DouyinBillboardResource) FetchHotTotalTopicList(ctx context.Context, request DouyinBillboardFetchHotTotalTopicListRequest) (*DouyinBillboardFetchHotTotalTopicListResponse, error) {
	return r.client.DouyinBillboardFetchTopicHotList(ctx, request)
}

// DouyinBillboardFetchHotTotalHighTopicListRequest is the request for POST /api/v1/douyin/billboard/fetch_hot_total_high_topic_list.
type DouyinBillboardFetchHotTotalHighTopicListRequest = DouyinBillboardFetchTopicListWithRisingPopularityRequest

// DouyinBillboardFetchHotTotalHighTopicListResponse is the response for POST /api/v1/douyin/billboard/fetch_hot_total_high_topic_list.
type DouyinBillboardFetchHotTotalHighTopicListResponse = DouyinBillboardFetchTopicListWithRisingPopularityResponse

// FetchHotTotalHighTopicList 获取热度飙升的话题榜/Fetch topic list with rising popularity
//
// POST /api/v1/douyin/billboard/fetch_hot_total_high_topic_list
func (r DouyinBillboardResource) FetchHotTotalHighTopicList(ctx context.Context, request DouyinBillboardFetchHotTotalHighTopicListRequest) (*DouyinBillboardFetchHotTotalHighTopicListResponse, error) {
	return r.client.DouyinBillboardFetchTopicListWithRisingPopularity(ctx, request)
}

// DouyinBillboardFetchHotTotalSearchListRequest is the request for POST /api/v1/douyin/billboard/fetch_hot_total_search_list.
type DouyinBillboardFetchHotTotalSearchListRequest = DouyinBillboardFetchSearchHotListRequest

// DouyinBillboardFetchHotTotalSearchListResponse is the response for POST /api/v1/douyin/billboard/fetch_hot_total_search_list.
type DouyinBillboardFetchHotTotalSearchListResponse = DouyinBillboardFetchSearchHotListResponse

// FetchHotTotalSearchList 获取搜索热榜/Fetch search hot list
//
// POST /api/v1/douyin/billboard/fetch_hot_total_search_list
func (r DouyinBillboardResource) FetchHotTotalSearchList(ctx context.Context, request DouyinBillboardFetchHotTotalSearchListRequest) (*DouyinBillboardFetchHotTotalSearchListResponse, error) {
	return r.client.DouyinBillboardFetchSearchHotList(ctx, request)
}

// DouyinBillboardFetchHotTotalHighSearchListRequest is the request for POST /api/v1/douyin/billboard/fetch_hot_total_high_search_list.
type DouyinBillboardFetchHotTotalHighSearchListRequest = DouyinBillboardFetchSearchListWithRisingPopularityRequest

// DouyinBillboardFetchHotTotalHighSearchListResponse is the response for POST /api/v1/douyin/billboard/fetch_hot_total_high_search_list.
type DouyinBillboardFetchHotTotalHighSearchListResponse = DouyinBillboardFetchSearchListWithRisingPopularityResponse

// FetchHotTotalHighSearchList 获取热度飙升的搜索榜/Fetch search list with rising popularity
//
// POST /api/v1/douyin/billboard/fetch_hot_total_high_search_list
func (r DouyinBillboardResource) FetchHotTotalHighSearchList(ctx context.Context, request DouyinBillboardFetchHotTotalHighSearchListRequest) (*DouyinBillboardFetchHotTotalHighSearchListResponse, error) {
	return r.client.DouyinBillboardFetchSearchListWithRisingPopularity(ctx, request)
}

// DouyinBillboardFetchHotTotalHotWordListRequest is the request for POST /api/v1/douyin/billboard/fetch_hot_total_hot_word_list.
type DouyinBillboardFetchHotTotalHotWordListRequest = DouyinBillboardFetchAllHotContentWordsRequest

// DouyinBillboardFetchHotTotalHotWordListResponse is the response for POST /api/v1/douyin/billboard/fetch_hot_total_hot_word_list.
type DouyinBillboardFetchHotTotalHotWordListResponse = DouyinBillboardFetchAllHotContentWordsResponse

// FetchHotTotalHotWordList 获取全部热门内容词/Fetch all hot content words
//
// POST /api/v1/douyin/billboard/fetch_hot_total_hot_word_list
func (r DouyinBillboardResource) FetchHotTotalHotWordList(ctx context.Context, request DouyinBillboardFetchHotTotalHotWordListRequest) (*DouyinBillboardFetchHotTotalHotWordListResponse, error) {
	return r.client.DouyinBillboardFetchAllHotContentWords(ctx, request)
}

// DouyinBillboardFetchHotTotalHotWordDetailListRequest is the request for GET /api/v1/douyin/billboard/fetch_hot_total_hot_word_detail_list.
type DouyinBillboardFetchHotTotalHotWordDetailListRequest = DouyinBillboardFetchContentWordDetailsRequest

// DouyinBillboardFetchHotTotalHotWordDetailListResponse is the response for GET /api/v1/douyin/billboard/fetch_hot_total_hot_word_detail_list.
type DouyinBillboardFetchHotTotalHotWordDetailListResponse = DouyinBillboardFetchContentWordDetailsResponse

// FetchHotTotalHotWordDetailList 获取内容词详情/Fetch content word details
//
// GET /api/v1/douyin/billboard/fetch_hot_total_hot_word_detail_list
func (r DouyinBillboardResource) FetchHotTotalHotWordDetailList(ctx context.Context, request DouyinBillboardFetchHotTotalHotWordDetailListRequest) (*DouyinBillboardFetchHotTotalHotWordDetailListResponse, error) {
	return r.client.DouyinBillboardFetchContentWordDetails(ctx, request)
}

// DouyinXingtuResource contains endpoints from the Douyin-Xingtu-API tag.
type DouyinXingtuResource struct {
	client *Client
}

// GetSignImage 获取加密图片解析/Get Sign Image
//
// GET /api/v1/douyin/xingtu/get_sign_image
func (r DouyinXingtuResource) GetSignImage(ctx context.Context, request DouyinXingtuGetSignImageRequest) (*DouyinXingtuGetSignImageResponse, error) {
	return r.client.DouyinXingtuGetSignImage(ctx, request)
}

// DouyinXingtuGetXingtuKolidByUIDRequest is the request for GET /api/v1/douyin/xingtu/get_xingtu_kolid_by_uid.
type DouyinXingtuGetXingtuKolidByUIDRequest = DouyinXingtuGetXingTuKolidByDouyinUserIDRequest

// DouyinXingtuGetXingtuKolidByUIDResponse is the response for GET /api/v1/douyin/xingtu/get_xingtu_kolid_by_uid.
type DouyinXingtuGetXingtuKolidByUIDResponse = DouyinXingtuGetXingTuKolidByDouyinUserIDResponse

// GetXingtuKolidByUID 根据抖音用户ID获取游客星图kolid/Get XingTu kolid by Douyin User ID
//
// GET /api/v1/douyin/xingtu/get_xingtu_kolid_by_uid
func (r DouyinXingtuResource) GetXingtuKolidByUID(ctx context.Context, request DouyinXingtuGetXingtuKolidByUIDRequest) (*DouyinXingtuGetXingtuKolidByUIDResponse, error) {
	return r.client.DouyinXingtuGetXingTuKolidByDouyinUserID(ctx, request)
}

// DouyinXingtuGetXingtuKolidBySecUserIDRequest is the request for GET /api/v1/douyin/xingtu/get_xingtu_kolid_by_sec_user_id.
type DouyinXingtuGetXingtuKolidBySecUserIDRequest = DouyinXingtuGetXingTuKolidByDouyinSecUserIDRequest

// DouyinXingtuGetXingtuKolidBySecUserIDResponse is the response for GET /api/v1/douyin/xingtu/get_xingtu_kolid_by_sec_user_id.
type DouyinXingtuGetXingtuKolidBySecUserIDResponse = DouyinXingtuGetXingTuKolidByDouyinSecUserIDResponse

// GetXingtuKolidBySecUserID 根据抖音sec_user_id获取游客星图kolid/Get XingTu kolid by Douyin sec_user_id
//
// GET /api/v1/douyin/xingtu/get_xingtu_kolid_by_sec_user_id
func (r DouyinXingtuResource) GetXingtuKolidBySecUserID(ctx context.Context, request DouyinXingtuGetXingtuKolidBySecUserIDRequest) (*DouyinXingtuGetXingtuKolidBySecUserIDResponse, error) {
	return r.client.DouyinXingtuGetXingTuKolidByDouyinSecUserID(ctx, request)
}

// DouyinXingtuGetXingtuKolidByUniqueIDRequest is the request for GET /api/v1/douyin/xingtu/get_xingtu_kolid_by_unique_id.
type DouyinXingtuGetXingtuKolidByUniqueIDRequest = DouyinXingtuGetXingTuKolidByDouyinUniqueIDRequest

// DouyinXingtuGetXingtuKolidByUniqueIDResponse is the response for GET /api/v1/douyin/xingtu/get_xingtu_kolid_by_unique_id.
type DouyinXingtuGetXingtuKolidByUniqueIDResponse = DouyinXingtuGetXingTuKolidByDouyinUniqueIDResponse

// GetXingtuKolidByUniqueID 根据抖音号获取游客星图kolid/Get XingTu kolid by Douyin unique_id
//
// GET /api/v1/douyin/xingtu/get_xingtu_kolid_by_unique_id
func (r DouyinXingtuResource) GetXingtuKolidByUniqueID(ctx context.Context, request DouyinXingtuGetXingtuKolidByUniqueIDRequest) (*DouyinXingtuGetXingtuKolidByUniqueIDResponse, error) {
	return r.client.DouyinXingtuGetXingTuKolidByDouyinUniqueID(ctx, request)
}

// DouyinXingtuKolBaseInfoV1Request is the request for GET /api/v1/douyin/xingtu/kol_base_info_v1.
type DouyinXingtuKolBaseInfoV1Request = DouyinXingtuGetKolBaseInfoV1Request

// DouyinXingtuKolBaseInfoV1Response is the response for GET /api/v1/douyin/xingtu/kol_base_info_v1.
type DouyinXingtuKolBaseInfoV1Response = DouyinXingtuGetKolBaseInfoV1Response

// KolBaseInfoV1 获取kol基本信息V1/Get kol Base Info V1
//
// GET /api/v1/douyin/xingtu/kol_base_info_v1
func (r DouyinXingtuResource) KolBaseInfoV1(ctx context.Context, request DouyinXingtuKolBaseInfoV1Request) (*DouyinXingtuKolBaseInfoV1Response, error) {
	return r.client.DouyinXingtuGetKolBaseInfoV1(ctx, request)
}

// DouyinXingtuKolAudiencePortraitV1Request is the request for GET /api/v1/douyin/xingtu/kol_audience_portrait_v1.
type DouyinXingtuKolAudiencePortraitV1Request = DouyinXingtuGetKolAudiencePortraitV1Request

// DouyinXingtuKolAudiencePortraitV1Response is the response for GET /api/v1/douyin/xingtu/kol_audience_portrait_v1.
type DouyinXingtuKolAudiencePortraitV1Response = DouyinXingtuGetKolAudiencePortraitV1Response

// KolAudiencePortraitV1 获取kol观众画像V1/Get kol Audience Portrait V1
//
// GET /api/v1/douyin/xingtu/kol_audience_portrait_v1
func (r DouyinXingtuResource) KolAudiencePortraitV1(ctx context.Context, request DouyinXingtuKolAudiencePortraitV1Request) (*DouyinXingtuKolAudiencePortraitV1Response, error) {
	return r.client.DouyinXingtuGetKolAudiencePortraitV1(ctx, request)
}

// DouyinXingtuKolFansPortraitV1Request is the request for GET /api/v1/douyin/xingtu/kol_fans_portrait_v1.
type DouyinXingtuKolFansPortraitV1Request = DouyinXingtuGetKolFansPortraitV1Request

// DouyinXingtuKolFansPortraitV1Response is the response for GET /api/v1/douyin/xingtu/kol_fans_portrait_v1.
type DouyinXingtuKolFansPortraitV1Response = DouyinXingtuGetKolFansPortraitV1Response

// KolFansPortraitV1 获取kol粉丝画像V1/Get kol Fans Portrait V1
//
// GET /api/v1/douyin/xingtu/kol_fans_portrait_v1
func (r DouyinXingtuResource) KolFansPortraitV1(ctx context.Context, request DouyinXingtuKolFansPortraitV1Request) (*DouyinXingtuKolFansPortraitV1Response, error) {
	return r.client.DouyinXingtuGetKolFansPortraitV1(ctx, request)
}

// DouyinXingtuKolServicePriceV1Request is the request for GET /api/v1/douyin/xingtu/kol_service_price_v1.
type DouyinXingtuKolServicePriceV1Request = DouyinXingtuGetKolServicePriceV1Request

// DouyinXingtuKolServicePriceV1Response is the response for GET /api/v1/douyin/xingtu/kol_service_price_v1.
type DouyinXingtuKolServicePriceV1Response = DouyinXingtuGetKolServicePriceV1Response

// KolServicePriceV1 获取kol服务报价V1/Get kol Service Price V1
//
// GET /api/v1/douyin/xingtu/kol_service_price_v1
func (r DouyinXingtuResource) KolServicePriceV1(ctx context.Context, request DouyinXingtuKolServicePriceV1Request) (*DouyinXingtuKolServicePriceV1Response, error) {
	return r.client.DouyinXingtuGetKolServicePriceV1(ctx, request)
}

// DouyinXingtuKolDataOverviewV1Request is the request for GET /api/v1/douyin/xingtu/kol_data_overview_v1.
type DouyinXingtuKolDataOverviewV1Request = DouyinXingtuGetKolDataOverviewV1Request

// DouyinXingtuKolDataOverviewV1Response is the response for GET /api/v1/douyin/xingtu/kol_data_overview_v1.
type DouyinXingtuKolDataOverviewV1Response = DouyinXingtuGetKolDataOverviewV1Response

// KolDataOverviewV1 获取kol数据概览V1/Get kol Data Overview V1
//
// GET /api/v1/douyin/xingtu/kol_data_overview_v1
func (r DouyinXingtuResource) KolDataOverviewV1(ctx context.Context, request DouyinXingtuKolDataOverviewV1Request) (*DouyinXingtuKolDataOverviewV1Response, error) {
	return r.client.DouyinXingtuGetKolDataOverviewV1(ctx, request)
}

// SearchKolV1 关键词搜索kol V1/Search Kol V1
//
// GET /api/v1/douyin/xingtu/search_kol_v1
func (r DouyinXingtuResource) SearchKolV1(ctx context.Context, request DouyinXingtuSearchKolV1Request) (*DouyinXingtuSearchKolV1Response, error) {
	return r.client.DouyinXingtuSearchKolV1(ctx, request)
}

// DouyinXingtuSearchKolV2Request is the request for GET /api/v1/douyin/xingtu/search_kol_v2.
type DouyinXingtuSearchKolV2Request = DouyinXingtuSearchKolAdvancedV2Request

// DouyinXingtuSearchKolV2Response is the response for GET /api/v1/douyin/xingtu/search_kol_v2.
type DouyinXingtuSearchKolV2Response = DouyinXingtuSearchKolAdvancedV2Response

// SearchKolV2 高级搜索kol V2/Search Kol Advanced V2
//
// GET /api/v1/douyin/xingtu/search_kol_v2
func (r DouyinXingtuResource) SearchKolV2(ctx context.Context, request DouyinXingtuSearchKolV2Request) (*DouyinXingtuSearchKolV2Response, error) {
	return r.client.DouyinXingtuSearchKolAdvancedV2(ctx, request)
}

// DouyinXingtuKolConversionAbilityAnalysisV1Request is the request for GET /api/v1/douyin/xingtu/kol_conversion_ability_analysis_v1.
type DouyinXingtuKolConversionAbilityAnalysisV1Request = DouyinXingtuGetKolConversionAbilityAnalysisV1Request

// DouyinXingtuKolConversionAbilityAnalysisV1Response is the response for GET /api/v1/douyin/xingtu/kol_conversion_ability_analysis_v1.
type DouyinXingtuKolConversionAbilityAnalysisV1Response = DouyinXingtuGetKolConversionAbilityAnalysisV1Response

// KolConversionAbilityAnalysisV1 获取kol转化能力分析V1/Get kol Conversion Ability Analysis V1
//
// GET /api/v1/douyin/xingtu/kol_conversion_ability_analysis_v1
func (r DouyinXingtuResource) KolConversionAbilityAnalysisV1(ctx context.Context, request DouyinXingtuKolConversionAbilityAnalysisV1Request) (*DouyinXingtuKolConversionAbilityAnalysisV1Response, error) {
	return r.client.DouyinXingtuGetKolConversionAbilityAnalysisV1(ctx, request)
}

// DouyinXingtuKolVideoPerformanceV1Request is the request for GET /api/v1/douyin/xingtu/kol_video_performance_v1.
type DouyinXingtuKolVideoPerformanceV1Request = DouyinXingtuGetKolVideoPerformanceV1Request

// DouyinXingtuKolVideoPerformanceV1Response is the response for GET /api/v1/douyin/xingtu/kol_video_performance_v1.
type DouyinXingtuKolVideoPerformanceV1Response = DouyinXingtuGetKolVideoPerformanceV1Response

// KolVideoPerformanceV1 获取kol视频表现V1/Get kol Video Performance V1
//
// GET /api/v1/douyin/xingtu/kol_video_performance_v1
func (r DouyinXingtuResource) KolVideoPerformanceV1(ctx context.Context, request DouyinXingtuKolVideoPerformanceV1Request) (*DouyinXingtuKolVideoPerformanceV1Response, error) {
	return r.client.DouyinXingtuGetKolVideoPerformanceV1(ctx, request)
}

// DouyinXingtuKolXingtuIndexV1Request is the request for GET /api/v1/douyin/xingtu/kol_xingtu_index_v1.
type DouyinXingtuKolXingtuIndexV1Request = DouyinXingtuGetKolXingtuIndexV1Request

// DouyinXingtuKolXingtuIndexV1Response is the response for GET /api/v1/douyin/xingtu/kol_xingtu_index_v1.
type DouyinXingtuKolXingtuIndexV1Response = DouyinXingtuGetKolXingtuIndexV1Response

// KolXingtuIndexV1 获取kol星图指数V1/Get kol Xingtu Index V1
//
// GET /api/v1/douyin/xingtu/kol_xingtu_index_v1
func (r DouyinXingtuResource) KolXingtuIndexV1(ctx context.Context, request DouyinXingtuKolXingtuIndexV1Request) (*DouyinXingtuKolXingtuIndexV1Response, error) {
	return r.client.DouyinXingtuGetKolXingtuIndexV1(ctx, request)
}

// DouyinXingtuKolConvertVideoDisplayV1Request is the request for GET /api/v1/douyin/xingtu/kol_convert_video_display_v1.
type DouyinXingtuKolConvertVideoDisplayV1Request = DouyinXingtuGetKolConvertVideoDisplayV1Request

// DouyinXingtuKolConvertVideoDisplayV1Response is the response for GET /api/v1/douyin/xingtu/kol_convert_video_display_v1.
type DouyinXingtuKolConvertVideoDisplayV1Response = DouyinXingtuGetKolConvertVideoDisplayV1Response

// KolConvertVideoDisplayV1 获取kol转化视频展示V1/Get kol Convert Video Display V1
//
// GET /api/v1/douyin/xingtu/kol_convert_video_display_v1
func (r DouyinXingtuResource) KolConvertVideoDisplayV1(ctx context.Context, request DouyinXingtuKolConvertVideoDisplayV1Request) (*DouyinXingtuKolConvertVideoDisplayV1Response, error) {
	return r.client.DouyinXingtuGetKolConvertVideoDisplayV1(ctx, request)
}

// DouyinXingtuKolLinkStructV1Request is the request for GET /api/v1/douyin/xingtu/kol_link_struct_v1.
type DouyinXingtuKolLinkStructV1Request = DouyinXingtuGetKolLinkStructV1Request

// DouyinXingtuKolLinkStructV1Response is the response for GET /api/v1/douyin/xingtu/kol_link_struct_v1.
type DouyinXingtuKolLinkStructV1Response = DouyinXingtuGetKolLinkStructV1Response

// KolLinkStructV1 获取kol连接用户V1/Get kol Link Struct V1
//
// GET /api/v1/douyin/xingtu/kol_link_struct_v1
func (r DouyinXingtuResource) KolLinkStructV1(ctx context.Context, request DouyinXingtuKolLinkStructV1Request) (*DouyinXingtuKolLinkStructV1Response, error) {
	return r.client.DouyinXingtuGetKolLinkStructV1(ctx, request)
}

// DouyinXingtuKolTouchDistributionV1Request is the request for GET /api/v1/douyin/xingtu/kol_touch_distribution_v1.
type DouyinXingtuKolTouchDistributionV1Request = DouyinXingtuGetKolTouchDistributionV1Request

// DouyinXingtuKolTouchDistributionV1Response is the response for GET /api/v1/douyin/xingtu/kol_touch_distribution_v1.
type DouyinXingtuKolTouchDistributionV1Response = DouyinXingtuGetKolTouchDistributionV1Response

// KolTouchDistributionV1 获取kol连接用户来源V1/Get kol Touch Distribution V1
//
// GET /api/v1/douyin/xingtu/kol_touch_distribution_v1
func (r DouyinXingtuResource) KolTouchDistributionV1(ctx context.Context, request DouyinXingtuKolTouchDistributionV1Request) (*DouyinXingtuKolTouchDistributionV1Response, error) {
	return r.client.DouyinXingtuGetKolTouchDistributionV1(ctx, request)
}

// DouyinXingtuKolCpInfoV1Request is the request for GET /api/v1/douyin/xingtu/kol_cp_info_v1.
type DouyinXingtuKolCpInfoV1Request = DouyinXingtuGetKolCpInfoV1Request

// DouyinXingtuKolCpInfoV1Response is the response for GET /api/v1/douyin/xingtu/kol_cp_info_v1.
type DouyinXingtuKolCpInfoV1Response = DouyinXingtuGetKolCpInfoV1Response

// KolCpInfoV1 获取kol性价比能力分析V1/Get kol Cp Info V1
//
// GET /api/v1/douyin/xingtu/kol_cp_info_v1
func (r DouyinXingtuResource) KolCpInfoV1(ctx context.Context, request DouyinXingtuKolCpInfoV1Request) (*DouyinXingtuKolCpInfoV1Response, error) {
	return r.client.DouyinXingtuGetKolCpInfoV1(ctx, request)
}

// DouyinXingtuKolRecVideosV1Request is the request for GET /api/v1/douyin/xingtu/kol_rec_videos_v1.
type DouyinXingtuKolRecVideosV1Request = DouyinXingtuGetKolRecVideosV1Request

// DouyinXingtuKolRecVideosV1Response is the response for GET /api/v1/douyin/xingtu/kol_rec_videos_v1.
type DouyinXingtuKolRecVideosV1Response = DouyinXingtuGetKolRecVideosV1Response

// KolRecVideosV1 获取kol内容表现V1/Get kol Rec Videos V1
//
// GET /api/v1/douyin/xingtu/kol_rec_videos_v1
func (r DouyinXingtuResource) KolRecVideosV1(ctx context.Context, request DouyinXingtuKolRecVideosV1Request) (*DouyinXingtuKolRecVideosV1Response, error) {
	return r.client.DouyinXingtuGetKolRecVideosV1(ctx, request)
}

// DouyinXingtuKolDailyFansV1Request is the request for GET /api/v1/douyin/xingtu/kol_daily_fans_v1.
type DouyinXingtuKolDailyFansV1Request = DouyinXingtuGetKolDailyFansV1Request

// DouyinXingtuKolDailyFansV1Response is the response for GET /api/v1/douyin/xingtu/kol_daily_fans_v1.
type DouyinXingtuKolDailyFansV1Response = DouyinXingtuGetKolDailyFansV1Response

// KolDailyFansV1 获取kol粉丝趋势V1/Get kol Daily Fans V1
//
// GET /api/v1/douyin/xingtu/kol_daily_fans_v1
func (r DouyinXingtuResource) KolDailyFansV1(ctx context.Context, request DouyinXingtuKolDailyFansV1Request) (*DouyinXingtuKolDailyFansV1Response, error) {
	return r.client.DouyinXingtuGetKolDailyFansV1(ctx, request)
}

// DouyinXingtuAuthorHotCommentTokensV1Request is the request for GET /api/v1/douyin/xingtu/author_hot_comment_tokens_v1.
type DouyinXingtuAuthorHotCommentTokensV1Request = DouyinXingtuGetAuthorHotCommentTokensV1Request

// DouyinXingtuAuthorHotCommentTokensV1Response is the response for GET /api/v1/douyin/xingtu/author_hot_comment_tokens_v1.
type DouyinXingtuAuthorHotCommentTokensV1Response = DouyinXingtuGetAuthorHotCommentTokensV1Response

// AuthorHotCommentTokensV1 获取kol热词分析评论V1/Get Author Hot Comment Tokens V1
//
// GET /api/v1/douyin/xingtu/author_hot_comment_tokens_v1
func (r DouyinXingtuResource) AuthorHotCommentTokensV1(ctx context.Context, request DouyinXingtuAuthorHotCommentTokensV1Request) (*DouyinXingtuAuthorHotCommentTokensV1Response, error) {
	return r.client.DouyinXingtuGetAuthorHotCommentTokensV1(ctx, request)
}

// DouyinXingtuAuthorContentHotCommentKeywordsV1Request is the request for GET /api/v1/douyin/xingtu/author_content_hot_comment_keywords_v1.
type DouyinXingtuAuthorContentHotCommentKeywordsV1Request = DouyinXingtuGetAuthorContentHotCommentKeywordsV1Request

// DouyinXingtuAuthorContentHotCommentKeywordsV1Response is the response for GET /api/v1/douyin/xingtu/author_content_hot_comment_keywords_v1.
type DouyinXingtuAuthorContentHotCommentKeywordsV1Response = DouyinXingtuGetAuthorContentHotCommentKeywordsV1Response

// AuthorContentHotCommentKeywordsV1 获取kol热词分析内容V1/Get Author Content Hot Comment Keywords V1
//
// GET /api/v1/douyin/xingtu/author_content_hot_comment_keywords_v1
func (r DouyinXingtuResource) AuthorContentHotCommentKeywordsV1(ctx context.Context, request DouyinXingtuAuthorContentHotCommentKeywordsV1Request) (*DouyinXingtuAuthorContentHotCommentKeywordsV1Response, error) {
	return r.client.DouyinXingtuGetAuthorContentHotCommentKeywordsV1(ctx, request)
}

// DouyinXingtuV2Resource contains endpoints from the Douyin-Xingtu-V2-API tag.
type DouyinXingtuV2Resource struct {
	client *Client
}

// GetRankingListCatalog 获取星图热榜分类/Get Ranking List Catalog
//
// GET /api/v1/douyin/xingtu_v2/get_ranking_list_catalog
func (r DouyinXingtuV2Resource) GetRankingListCatalog(ctx context.Context, request DouyinXingtuV2GetRankingListCatalogRequest) (*DouyinXingtuV2GetRankingListCatalogResponse, error) {
	return r.client.DouyinXingtuV2GetRankingListCatalog(ctx, request)
}

// GetRankingListData 获取星图达人商业榜数据/Get Ranking List Data
//
// GET /api/v1/douyin/xingtu_v2/get_ranking_list_data
func (r DouyinXingtuV2Resource) GetRankingListData(ctx context.Context, request DouyinXingtuV2GetRankingListDataRequest) (*DouyinXingtuV2GetRankingListDataResponse, error) {
	return r.client.DouyinXingtuV2GetRankingListData(ctx, request)
}

// GetPlayletActorRankCatalog 获取短剧演员热榜分类/Get Playlet Actor Rank Catalog
//
// POST /api/v1/douyin/xingtu_v2/get_playlet_actor_rank_catalog
func (r DouyinXingtuV2Resource) GetPlayletActorRankCatalog(ctx context.Context) (*DouyinXingtuV2GetPlayletActorRankCatalogResponse, error) {
	return r.client.DouyinXingtuV2GetPlayletActorRankCatalog(ctx)
}

// GetPlayletActorRankList 获取短剧演员热榜/Get Playlet Actor Rank List
//
// GET /api/v1/douyin/xingtu_v2/get_playlet_actor_rank_list
func (r DouyinXingtuV2Resource) GetPlayletActorRankList(ctx context.Context, request DouyinXingtuV2GetPlayletActorRankListRequest) (*DouyinXingtuV2GetPlayletActorRankListResponse, error) {
	return r.client.DouyinXingtuV2GetPlayletActorRankList(ctx, request)
}

// GetAuthorMarketFields 获取达人广场筛选字段/Get Author Market Fields
//
// GET /api/v1/douyin/xingtu_v2/get_author_market_fields
func (r DouyinXingtuV2Resource) GetAuthorMarketFields(ctx context.Context, request DouyinXingtuV2GetAuthorMarketFieldsRequest) (*DouyinXingtuV2GetAuthorMarketFieldsResponse, error) {
	return r.client.DouyinXingtuV2GetAuthorMarketFields(ctx, request)
}

// GetAuthorBaseInfo 获取创作者基本信息/Get Author Base Info
//
// GET /api/v1/douyin/xingtu_v2/get_author_base_info
func (r DouyinXingtuV2Resource) GetAuthorBaseInfo(ctx context.Context, request DouyinXingtuV2GetAuthorBaseInfoRequest) (*DouyinXingtuV2GetAuthorBaseInfoResponse, error) {
	return r.client.DouyinXingtuV2GetAuthorBaseInfo(ctx, request)
}

// GetAuthorBusinessCardInfo 获取创作者商业卡片信息/Get Author Business Card Info
//
// GET /api/v1/douyin/xingtu_v2/get_author_business_card_info
func (r DouyinXingtuV2Resource) GetAuthorBusinessCardInfo(ctx context.Context, request DouyinXingtuV2GetAuthorBusinessCardInfoRequest) (*DouyinXingtuV2GetAuthorBusinessCardInfoResponse, error) {
	return r.client.DouyinXingtuV2GetAuthorBusinessCardInfo(ctx, request)
}

// GetAuthorLocalInfo 获取创作者位置信息/Get Author Local Info
//
// GET /api/v1/douyin/xingtu_v2/get_author_local_info
func (r DouyinXingtuV2Resource) GetAuthorLocalInfo(ctx context.Context, request DouyinXingtuV2GetAuthorLocalInfoRequest) (*DouyinXingtuV2GetAuthorLocalInfoResponse, error) {
	return r.client.DouyinXingtuV2GetAuthorLocalInfo(ctx, request)
}

// GetAuthorShowItems 获取创作者视频列表/Get Author Show Items
//
// GET /api/v1/douyin/xingtu_v2/get_author_show_items
func (r DouyinXingtuV2Resource) GetAuthorShowItems(ctx context.Context, request DouyinXingtuV2GetAuthorShowItemsRequest) (*DouyinXingtuV2GetAuthorShowItemsResponse, error) {
	return r.client.DouyinXingtuV2GetAuthorShowItems(ctx, request)
}

// GetAuthorHotCommentTokens 获取创作者评论热词/Get Author Hot Comment Tokens
//
// GET /api/v1/douyin/xingtu_v2/get_author_hot_comment_tokens
func (r DouyinXingtuV2Resource) GetAuthorHotCommentTokens(ctx context.Context, request DouyinXingtuV2GetAuthorHotCommentTokensRequest) (*DouyinXingtuV2GetAuthorHotCommentTokensResponse, error) {
	return r.client.DouyinXingtuV2GetAuthorHotCommentTokens(ctx, request)
}

// GetAuthorContentHotKeywords 获取创作者内容热词/Get Author Content Hot Keywords
//
// GET /api/v1/douyin/xingtu_v2/get_author_content_hot_keywords
func (r DouyinXingtuV2Resource) GetAuthorContentHotKeywords(ctx context.Context, request DouyinXingtuV2GetAuthorContentHotKeywordsRequest) (*DouyinXingtuV2GetAuthorContentHotKeywordsResponse, error) {
	return r.client.DouyinXingtuV2GetAuthorContentHotKeywords(ctx, request)
}

// DouyinXingtuV2GetRecommendForStarAuthorsRequest is the request for POST /api/v1/douyin/xingtu_v2/get_recommend_for_star_authors.
type DouyinXingtuV2GetRecommendForStarAuthorsRequest = DouyinXingtuV2GetRecommendSimilarStarAuthorsRequest

// DouyinXingtuV2GetRecommendForStarAuthorsResponse is the response for POST /api/v1/douyin/xingtu_v2/get_recommend_for_star_authors.
type DouyinXingtuV2GetRecommendForStarAuthorsResponse = DouyinXingtuV2GetRecommendSimilarStarAuthorsResponse

// GetRecommendForStarAuthors 获取相似创作者推荐/Get Recommend Similar Star Authors
//
// POST /api/v1/douyin/xingtu_v2/get_recommend_for_star_authors
func (r DouyinXingtuV2Resource) GetRecommendForStarAuthors(ctx context.Context, request DouyinXingtuV2GetRecommendForStarAuthorsRequest) (*DouyinXingtuV2GetRecommendForStarAuthorsResponse, error) {
	return r.client.DouyinXingtuV2GetRecommendSimilarStarAuthors(ctx, request)
}

// GetExcellentCaseCategoryList 获取优秀行业分类列表/Get Excellent Case Category List
//
// GET /api/v1/douyin/xingtu_v2/get_excellent_case_category_list
func (r DouyinXingtuV2Resource) GetExcellentCaseCategoryList(ctx context.Context, request DouyinXingtuV2GetExcellentCaseCategoryListRequest) (*DouyinXingtuV2GetExcellentCaseCategoryListResponse, error) {
	return r.client.DouyinXingtuV2GetExcellentCaseCategoryList(ctx, request)
}

// GetAuthorSpreadInfo 获取创作者传播价值/Get Author Spread Info
//
// GET /api/v1/douyin/xingtu_v2/get_author_spread_info
func (r DouyinXingtuV2Resource) GetAuthorSpreadInfo(ctx context.Context, request DouyinXingtuV2GetAuthorSpreadInfoRequest) (*DouyinXingtuV2GetAuthorSpreadInfoResponse, error) {
	return r.client.DouyinXingtuV2GetAuthorSpreadInfo(ctx, request)
}

// DouyinXingtuV2GetUserProfileQrcodeRequest is the request for GET /api/v1/douyin/xingtu_v2/get_user_profile_qrcode.
type DouyinXingtuV2GetUserProfileQrcodeRequest = DouyinXingtuV2GetUserProfileQRCodeRequest

// DouyinXingtuV2GetUserProfileQrcodeResponse is the response for GET /api/v1/douyin/xingtu_v2/get_user_profile_qrcode.
type DouyinXingtuV2GetUserProfileQrcodeResponse = DouyinXingtuV2GetUserProfileQRCodeResponse

// GetUserProfileQrcode 获取用户主页二维码/Get User Profile QRCode
//
// GET /api/v1/douyin/xingtu_v2/get_user_profile_qrcode
func (r DouyinXingtuV2Resource) GetUserProfileQrcode(ctx context.Context, request DouyinXingtuV2GetUserProfileQrcodeRequest) (*DouyinXingtuV2GetUserProfileQrcodeResponse, error) {
	return r.client.DouyinXingtuV2GetUserProfileQRCode(ctx, request)
}

// GetContentTrendGuide 获取内容趋势指南/Get Content Trend Guide
//
// GET /api/v1/douyin/xingtu_v2/get_content_trend_guide
func (r DouyinXingtuV2Resource) GetContentTrendGuide(ctx context.Context) (*DouyinXingtuV2GetContentTrendGuideResponse, error) {
	return r.client.DouyinXingtuV2GetContentTrendGuide(ctx)
}

// DouyinXingtuV2GetIpActivityIndustryListResponse is the response for GET /api/v1/douyin/xingtu_v2/get_ip_activity_industry_list.
type DouyinXingtuV2GetIpActivityIndustryListResponse = DouyinXingtuV2GetIPActivityIndustryListResponse

// GetIpActivityIndustryList 获取星图IP日历行业列表/Get IP Activity Industry List
//
// GET /api/v1/douyin/xingtu_v2/get_ip_activity_industry_list
func (r DouyinXingtuV2Resource) GetIpActivityIndustryList(ctx context.Context) (*DouyinXingtuV2GetIpActivityIndustryListResponse, error) {
	return r.client.DouyinXingtuV2GetIPActivityIndustryList(ctx)
}

// DouyinXingtuV2GetIpActivityListRequest is the request for POST /api/v1/douyin/xingtu_v2/get_ip_activity_list.
type DouyinXingtuV2GetIpActivityListRequest = DouyinXingtuV2GetIPActivityListRequest

// DouyinXingtuV2GetIpActivityListResponse is the response for POST /api/v1/douyin/xingtu_v2/get_ip_activity_list.
type DouyinXingtuV2GetIpActivityListResponse = DouyinXingtuV2GetIPActivityListResponse

// GetIpActivityList 获取星图IP日历活动列表/Get IP Activity List
//
// POST /api/v1/douyin/xingtu_v2/get_ip_activity_list
func (r DouyinXingtuV2Resource) GetIpActivityList(ctx context.Context, request DouyinXingtuV2GetIpActivityListRequest) (*DouyinXingtuV2GetIpActivityListResponse, error) {
	return r.client.DouyinXingtuV2GetIPActivityList(ctx, request)
}

// DouyinXingtuV2GetIpActivityDetailRequest is the request for GET /api/v1/douyin/xingtu_v2/get_ip_activity_detail.
type DouyinXingtuV2GetIpActivityDetailRequest = DouyinXingtuV2GetIPActivityDetailRequest

// DouyinXingtuV2GetIpActivityDetailResponse is the response for GET /api/v1/douyin/xingtu_v2/get_ip_activity_detail.
type DouyinXingtuV2GetIpActivityDetailResponse = DouyinXingtuV2GetIPActivityDetailResponse

// GetIpActivityDetail 获取星图IP活动详情/Get IP Activity Detail
//
// GET /api/v1/douyin/xingtu_v2/get_ip_activity_detail
func (r DouyinXingtuV2Resource) GetIpActivityDetail(ctx context.Context, request DouyinXingtuV2GetIpActivityDetailRequest) (*DouyinXingtuV2GetIpActivityDetailResponse, error) {
	return r.client.DouyinXingtuV2GetIPActivityDetail(ctx, request)
}

// GetResourceList 获取营销活动案例/Get Resource List
//
// GET /api/v1/douyin/xingtu_v2/get_resource_list
func (r DouyinXingtuV2Resource) GetResourceList(ctx context.Context, request DouyinXingtuV2GetResourceListRequest) (*DouyinXingtuV2GetResourceListResponse, error) {
	return r.client.DouyinXingtuV2GetResourceList(ctx, request)
}

// GetDemanderMcnList 搜索MCN机构列表/Get Demander MCN List
//
// GET /api/v1/douyin/xingtu_v2/get_demander_mcn_list
func (r DouyinXingtuV2Resource) GetDemanderMcnList(ctx context.Context, request DouyinXingtuV2GetDemanderMcnListRequest) (*DouyinXingtuV2GetDemanderMcnListResponse, error) {
	return r.client.DouyinXingtuV2GetDemanderMcnList(ctx, request)
}

// XiguaAppV2Resource contains endpoints from the Xigua-App-V2-API tag.
type XiguaAppV2Resource struct {
	client *Client
}

// XiguaAppV2FetchOneVideoRequest is the request for GET /api/v1/xigua/app/v2/fetch_one_video.
type XiguaAppV2FetchOneVideoRequest = XiguaAppV2GetSingleVideoDataRequest

// XiguaAppV2FetchOneVideoResponse is the response for GET /api/v1/xigua/app/v2/fetch_one_video.
type XiguaAppV2FetchOneVideoResponse = XiguaAppV2GetSingleVideoDataResponse

// FetchOneVideo 获取单个作品数据/Get single video data
//
// GET /api/v1/xigua/app/v2/fetch_one_video
func (r XiguaAppV2Resource) FetchOneVideo(ctx context.Context, request XiguaAppV2FetchOneVideoRequest) (*XiguaAppV2FetchOneVideoResponse, error) {
	return r.client.XiguaAppV2GetSingleVideoData(ctx, request)
}

// XiguaAppV2FetchOneVideoV2Request is the request for GET /api/v1/xigua/app/v2/fetch_one_video_v2.
type XiguaAppV2FetchOneVideoV2Request = XiguaAppV2GetSingleVideoDataV2Request

// XiguaAppV2FetchOneVideoV2Response is the response for GET /api/v1/xigua/app/v2/fetch_one_video_v2.
type XiguaAppV2FetchOneVideoV2Response = XiguaAppV2GetSingleVideoDataV2Response

// FetchOneVideoV2 获取单个作品数据 V2/Get single video data V2
//
// GET /api/v1/xigua/app/v2/fetch_one_video_v2
func (r XiguaAppV2Resource) FetchOneVideoV2(ctx context.Context, request XiguaAppV2FetchOneVideoV2Request) (*XiguaAppV2FetchOneVideoV2Response, error) {
	return r.client.XiguaAppV2GetSingleVideoDataV2(ctx, request)
}

// XiguaAppV2FetchOneVideoPlayURLRequest is the request for GET /api/v1/xigua/app/v2/fetch_one_video_play_url.
type XiguaAppV2FetchOneVideoPlayURLRequest = XiguaAppV2GetSingleVideoPlayURLRequest

// XiguaAppV2FetchOneVideoPlayURLResponse is the response for GET /api/v1/xigua/app/v2/fetch_one_video_play_url.
type XiguaAppV2FetchOneVideoPlayURLResponse = XiguaAppV2GetSingleVideoPlayURLResponse

// FetchOneVideoPlayURL 获取单个作品的播放链接/Get single video play URL
//
// GET /api/v1/xigua/app/v2/fetch_one_video_play_url
func (r XiguaAppV2Resource) FetchOneVideoPlayURL(ctx context.Context, request XiguaAppV2FetchOneVideoPlayURLRequest) (*XiguaAppV2FetchOneVideoPlayURLResponse, error) {
	return r.client.XiguaAppV2GetSingleVideoPlayURL(ctx, request)
}

// XiguaAppV2FetchVideoCommentListRequest is the request for GET /api/v1/xigua/app/v2/fetch_video_comment_list.
type XiguaAppV2FetchVideoCommentListRequest = XiguaAppV2VideoCommentListRequest

// XiguaAppV2FetchVideoCommentListResponse is the response for GET /api/v1/xigua/app/v2/fetch_video_comment_list.
type XiguaAppV2FetchVideoCommentListResponse = XiguaAppV2VideoCommentListResponse

// FetchVideoCommentList 视频评论列表/Video comment list
//
// GET /api/v1/xigua/app/v2/fetch_video_comment_list
func (r XiguaAppV2Resource) FetchVideoCommentList(ctx context.Context, request XiguaAppV2FetchVideoCommentListRequest) (*XiguaAppV2FetchVideoCommentListResponse, error) {
	return r.client.XiguaAppV2VideoCommentList(ctx, request)
}

// SearchVideo 搜索视频/Search video
//
// GET /api/v1/xigua/app/v2/search_video
func (r XiguaAppV2Resource) SearchVideo(ctx context.Context, request XiguaAppV2SearchVideoRequest) (*XiguaAppV2SearchVideoResponse, error) {
	return r.client.XiguaAppV2SearchVideo(ctx, request)
}

// XiguaAppV2FetchUserInfoRequest is the request for GET /api/v1/xigua/app/v2/fetch_user_info.
type XiguaAppV2FetchUserInfoRequest = XiguaAppV2PersonalInformationRequest

// XiguaAppV2FetchUserInfoResponse is the response for GET /api/v1/xigua/app/v2/fetch_user_info.
type XiguaAppV2FetchUserInfoResponse = XiguaAppV2PersonalInformationResponse

// FetchUserInfo 个人信息/Personal information
//
// GET /api/v1/xigua/app/v2/fetch_user_info
func (r XiguaAppV2Resource) FetchUserInfo(ctx context.Context, request XiguaAppV2FetchUserInfoRequest) (*XiguaAppV2FetchUserInfoResponse, error) {
	return r.client.XiguaAppV2PersonalInformation(ctx, request)
}

// XiguaAppV2FetchUserPostListRequest is the request for GET /api/v1/xigua/app/v2/fetch_user_post_list.
type XiguaAppV2FetchUserPostListRequest = XiguaAppV2GetUserPostListRequest

// XiguaAppV2FetchUserPostListResponse is the response for GET /api/v1/xigua/app/v2/fetch_user_post_list.
type XiguaAppV2FetchUserPostListResponse = XiguaAppV2GetUserPostListResponse

// FetchUserPostList 获取个人作品列表/Get user post list
//
// GET /api/v1/xigua/app/v2/fetch_user_post_list
func (r XiguaAppV2Resource) FetchUserPostList(ctx context.Context, request XiguaAppV2FetchUserPostListRequest) (*XiguaAppV2FetchUserPostListResponse, error) {
	return r.client.XiguaAppV2GetUserPostList(ctx, request)
}

// ToutiaoWebResource contains endpoints from the Toutiao-Web-API tag.
type ToutiaoWebResource struct {
	client *Client
}

// ToutiaoWebGetArticleInfoRequest is the request for GET /api/v1/toutiao/web/get_article_info.
type ToutiaoWebGetArticleInfoRequest = ToutiaoWebGetInformationOfSpecifiedArticleRequest

// ToutiaoWebGetArticleInfoResponse is the response for GET /api/v1/toutiao/web/get_article_info.
type ToutiaoWebGetArticleInfoResponse = ToutiaoWebGetInformationOfSpecifiedArticleResponse

// GetArticleInfo 获取指定文章的信息/Get information of specified article
//
// GET /api/v1/toutiao/web/get_article_info
func (r ToutiaoWebResource) GetArticleInfo(ctx context.Context, request ToutiaoWebGetArticleInfoRequest) (*ToutiaoWebGetArticleInfoResponse, error) {
	return r.client.ToutiaoWebGetInformationOfSpecifiedArticle(ctx, request)
}

// ToutiaoWebGetVideoInfoRequest is the request for GET /api/v1/toutiao/web/get_video_info.
type ToutiaoWebGetVideoInfoRequest = ToutiaoWebGetInformationOfSpecifiedVideoRequest

// ToutiaoWebGetVideoInfoResponse is the response for GET /api/v1/toutiao/web/get_video_info.
type ToutiaoWebGetVideoInfoResponse = ToutiaoWebGetInformationOfSpecifiedVideoResponse

// GetVideoInfo 获取指定视频的信息/Get information of specified video
//
// GET /api/v1/toutiao/web/get_video_info
func (r ToutiaoWebResource) GetVideoInfo(ctx context.Context, request ToutiaoWebGetVideoInfoRequest) (*ToutiaoWebGetVideoInfoResponse, error) {
	return r.client.ToutiaoWebGetInformationOfSpecifiedVideo(ctx, request)
}

// ToutiaoAppResource contains endpoints from the Toutiao-App-API tag.
type ToutiaoAppResource struct {
	client *Client
}

// ToutiaoAppGetArticleInfoRequest is the request for GET /api/v1/toutiao/app/get_article_info.
type ToutiaoAppGetArticleInfoRequest = ToutiaoAppGetInformationOfSpecifiedArticleRequest

// ToutiaoAppGetArticleInfoResponse is the response for GET /api/v1/toutiao/app/get_article_info.
type ToutiaoAppGetArticleInfoResponse = ToutiaoAppGetInformationOfSpecifiedArticleResponse

// GetArticleInfo 获取指定文章的信息/Get information of specified article
//
// GET /api/v1/toutiao/app/get_article_info
func (r ToutiaoAppResource) GetArticleInfo(ctx context.Context, request ToutiaoAppGetArticleInfoRequest) (*ToutiaoAppGetArticleInfoResponse, error) {
	return r.client.ToutiaoAppGetInformationOfSpecifiedArticle(ctx, request)
}

// ToutiaoAppGetVideoInfoRequest is the request for GET /api/v1/toutiao/app/get_video_info.
type ToutiaoAppGetVideoInfoRequest = ToutiaoAppGetInformationOfSpecifiedVideoRequest

// ToutiaoAppGetVideoInfoResponse is the response for GET /api/v1/toutiao/app/get_video_info.
type ToutiaoAppGetVideoInfoResponse = ToutiaoAppGetInformationOfSpecifiedVideoResponse

// GetVideoInfo 获取指定视频的信息/Get information of specified video
//
// GET /api/v1/toutiao/app/get_video_info
func (r ToutiaoAppResource) GetVideoInfo(ctx context.Context, request ToutiaoAppGetVideoInfoRequest) (*ToutiaoAppGetVideoInfoResponse, error) {
	return r.client.ToutiaoAppGetInformationOfSpecifiedVideo(ctx, request)
}

// ToutiaoAppGetCommentsRequest is the request for GET /api/v1/toutiao/app/get_comments.
type ToutiaoAppGetCommentsRequest = ToutiaoAppGetCommentsOfSpecifiedPostRequest

// ToutiaoAppGetCommentsResponse is the response for GET /api/v1/toutiao/app/get_comments.
type ToutiaoAppGetCommentsResponse = ToutiaoAppGetCommentsOfSpecifiedPostResponse

// GetComments 获取指定作品的评论/Get comments of specified post
//
// GET /api/v1/toutiao/app/get_comments
func (r ToutiaoAppResource) GetComments(ctx context.Context, request ToutiaoAppGetCommentsRequest) (*ToutiaoAppGetCommentsResponse, error) {
	return r.client.ToutiaoAppGetCommentsOfSpecifiedPost(ctx, request)
}

// ToutiaoAppGetUserInfoRequest is the request for GET /api/v1/toutiao/app/get_user_info.
type ToutiaoAppGetUserInfoRequest = ToutiaoAppGetInformationOfSpecifiedUserRequest

// ToutiaoAppGetUserInfoResponse is the response for GET /api/v1/toutiao/app/get_user_info.
type ToutiaoAppGetUserInfoResponse = ToutiaoAppGetInformationOfSpecifiedUserResponse

// GetUserInfo 获取指定用户的信息/Get information of specified user
//
// GET /api/v1/toutiao/app/get_user_info
func (r ToutiaoAppResource) GetUserInfo(ctx context.Context, request ToutiaoAppGetUserInfoRequest) (*ToutiaoAppGetUserInfoResponse, error) {
	return r.client.ToutiaoAppGetInformationOfSpecifiedUser(ctx, request)
}

// ToutiaoAppGetUserIDRequest is the request for GET /api/v1/toutiao/app/get_user_id.
type ToutiaoAppGetUserIDRequest = ToutiaoAppGetUserIDFromUserProfileRequest

// ToutiaoAppGetUserIDResponse is the response for GET /api/v1/toutiao/app/get_user_id.
type ToutiaoAppGetUserIDResponse = ToutiaoAppGetUserIDFromUserProfileResponse

// GetUserID 从头条用户主页获取用户user_id/Get user_id from user profile
//
// GET /api/v1/toutiao/app/get_user_id
func (r ToutiaoAppResource) GetUserID(ctx context.Context, request ToutiaoAppGetUserIDRequest) (*ToutiaoAppGetUserIDResponse, error) {
	return r.client.ToutiaoAppGetUserIDFromUserProfile(ctx, request)
}

// XiaohongshuWebV3Resource contains endpoints from the Xiaohongshu-Web-V3-API tag.
type XiaohongshuWebV3Resource struct {
	client *Client
}

// FetchNoteDetail 获取笔记详情/Fetch note detail
//
// GET /api/v1/xiaohongshu/web_v3/fetch_note_detail
func (r XiaohongshuWebV3Resource) FetchNoteDetail(ctx context.Context, request XiaohongshuWebV3FetchNoteDetailRequest) (*XiaohongshuWebV3FetchNoteDetailResponse, error) {
	return r.client.XiaohongshuWebV3FetchNoteDetail(ctx, request)
}

// FetchNoteComments 获取笔记评论/Fetch note comments
//
// GET /api/v1/xiaohongshu/web_v3/fetch_note_comments
func (r XiaohongshuWebV3Resource) FetchNoteComments(ctx context.Context, request XiaohongshuWebV3FetchNoteCommentsRequest) (*XiaohongshuWebV3FetchNoteCommentsResponse, error) {
	return r.client.XiaohongshuWebV3FetchNoteComments(ctx, request)
}

// FetchSubComments 获取子评论/Fetch sub comments
//
// GET /api/v1/xiaohongshu/web_v3/fetch_sub_comments
func (r XiaohongshuWebV3Resource) FetchSubComments(ctx context.Context, request XiaohongshuWebV3FetchSubCommentsRequest) (*XiaohongshuWebV3FetchSubCommentsResponse, error) {
	return r.client.XiaohongshuWebV3FetchSubComments(ctx, request)
}

// XiaohongshuWebV3FetchSearchNotesRequest is the request for GET /api/v1/xiaohongshu/web_v3/fetch_search_notes.
type XiaohongshuWebV3FetchSearchNotesRequest = XiaohongshuWebV3SearchNotesRequest

// XiaohongshuWebV3FetchSearchNotesResponse is the response for GET /api/v1/xiaohongshu/web_v3/fetch_search_notes.
type XiaohongshuWebV3FetchSearchNotesResponse = XiaohongshuWebV3SearchNotesResponse

// FetchSearchNotes 搜索笔记/Search notes
//
// GET /api/v1/xiaohongshu/web_v3/fetch_search_notes
func (r XiaohongshuWebV3Resource) FetchSearchNotes(ctx context.Context, request XiaohongshuWebV3FetchSearchNotesRequest) (*XiaohongshuWebV3FetchSearchNotesResponse, error) {
	return r.client.XiaohongshuWebV3SearchNotes(ctx, request)
}

// XiaohongshuWebV3FetchSearchUsersRequest is the request for GET /api/v1/xiaohongshu/web_v3/fetch_search_users.
type XiaohongshuWebV3FetchSearchUsersRequest = XiaohongshuWebV3SearchUsersRequest

// XiaohongshuWebV3FetchSearchUsersResponse is the response for GET /api/v1/xiaohongshu/web_v3/fetch_search_users.
type XiaohongshuWebV3FetchSearchUsersResponse = XiaohongshuWebV3SearchUsersResponse

// FetchSearchUsers 搜索用户/Search users
//
// GET /api/v1/xiaohongshu/web_v3/fetch_search_users
func (r XiaohongshuWebV3Resource) FetchSearchUsers(ctx context.Context, request XiaohongshuWebV3FetchSearchUsersRequest) (*XiaohongshuWebV3FetchSearchUsersResponse, error) {
	return r.client.XiaohongshuWebV3SearchUsers(ctx, request)
}

// XiaohongshuWebV3FetchTrendingResponse is the response for GET /api/v1/xiaohongshu/web_v3/fetch_trending.
type XiaohongshuWebV3FetchTrendingResponse = XiaohongshuWebV3FetchTrendingKeywordsResponse

// FetchTrending 获取热搜词/Fetch trending keywords
//
// GET /api/v1/xiaohongshu/web_v3/fetch_trending
func (r XiaohongshuWebV3Resource) FetchTrending(ctx context.Context) (*XiaohongshuWebV3FetchTrendingResponse, error) {
	return r.client.XiaohongshuWebV3FetchTrendingKeywords(ctx)
}

// XiaohongshuWebV3FetchSearchSuggestRequest is the request for GET /api/v1/xiaohongshu/web_v3/fetch_search_suggest.
type XiaohongshuWebV3FetchSearchSuggestRequest = XiaohongshuWebV3FetchSearchSuggestionsRequest

// XiaohongshuWebV3FetchSearchSuggestResponse is the response for GET /api/v1/xiaohongshu/web_v3/fetch_search_suggest.
type XiaohongshuWebV3FetchSearchSuggestResponse = XiaohongshuWebV3FetchSearchSuggestionsResponse

// FetchSearchSuggest 获取搜索联想词/Fetch search suggestions
//
// GET /api/v1/xiaohongshu/web_v3/fetch_search_suggest
func (r XiaohongshuWebV3Resource) FetchSearchSuggest(ctx context.Context, request XiaohongshuWebV3FetchSearchSuggestRequest) (*XiaohongshuWebV3FetchSearchSuggestResponse, error) {
	return r.client.XiaohongshuWebV3FetchSearchSuggestions(ctx, request)
}

// XiaohongshuWebV3FetchHomefeedRequest is the request for GET /api/v1/xiaohongshu/web_v3/fetch_homefeed.
type XiaohongshuWebV3FetchHomefeedRequest = XiaohongshuWebV3FetchHomepageFeedRequest

// XiaohongshuWebV3FetchHomefeedResponse is the response for GET /api/v1/xiaohongshu/web_v3/fetch_homefeed.
type XiaohongshuWebV3FetchHomefeedResponse = XiaohongshuWebV3FetchHomepageFeedResponse

// FetchHomefeed 获取首页推荐/Fetch homepage feed
//
// GET /api/v1/xiaohongshu/web_v3/fetch_homefeed
func (r XiaohongshuWebV3Resource) FetchHomefeed(ctx context.Context, request XiaohongshuWebV3FetchHomefeedRequest) (*XiaohongshuWebV3FetchHomefeedResponse, error) {
	return r.client.XiaohongshuWebV3FetchHomepageFeed(ctx, request)
}

// XiaohongshuWebV3FetchHomefeedCategoriesResponse is the response for GET /api/v1/xiaohongshu/web_v3/fetch_homefeed_categories.
type XiaohongshuWebV3FetchHomefeedCategoriesResponse = XiaohongshuWebV3FetchHomepageCategoriesResponse

// FetchHomefeedCategories 获取首页分类列表/Fetch homepage categories
//
// GET /api/v1/xiaohongshu/web_v3/fetch_homefeed_categories
func (r XiaohongshuWebV3Resource) FetchHomefeedCategories(ctx context.Context) (*XiaohongshuWebV3FetchHomefeedCategoriesResponse, error) {
	return r.client.XiaohongshuWebV3FetchHomepageCategories(ctx)
}

// FetchUserInfo 获取用户信息/Fetch user info
//
// GET /api/v1/xiaohongshu/web_v3/fetch_user_info
func (r XiaohongshuWebV3Resource) FetchUserInfo(ctx context.Context, request XiaohongshuWebV3FetchUserInfoRequest) (*XiaohongshuWebV3FetchUserInfoResponse, error) {
	return r.client.XiaohongshuWebV3FetchUserInfo(ctx, request)
}

// FetchUserNotes 获取用户笔记列表/Fetch user notes
//
// GET /api/v1/xiaohongshu/web_v3/fetch_user_notes
func (r XiaohongshuWebV3Resource) FetchUserNotes(ctx context.Context, request XiaohongshuWebV3FetchUserNotesRequest) (*XiaohongshuWebV3FetchUserNotesResponse, error) {
	return r.client.XiaohongshuWebV3FetchUserNotes(ctx, request)
}

// XiaohongshuAppV2Resource contains endpoints from the Xiaohongshu-App-V2-API tag.
type XiaohongshuAppV2Resource struct {
	client *Client
}

// GetImageNoteDetail 获取图文笔记详情/Get image note detail
//
// GET /api/v1/xiaohongshu/app_v2/get_image_note_detail
func (r XiaohongshuAppV2Resource) GetImageNoteDetail(ctx context.Context, request XiaohongshuAppV2GetImageNoteDetailRequest) (*XiaohongshuAppV2GetImageNoteDetailResponse, error) {
	return r.client.XiaohongshuAppV2GetImageNoteDetail(ctx, request)
}

// GetVideoNoteDetail 获取视频笔记详情/Get video note detail
//
// GET /api/v1/xiaohongshu/app_v2/get_video_note_detail
func (r XiaohongshuAppV2Resource) GetVideoNoteDetail(ctx context.Context, request XiaohongshuAppV2GetVideoNoteDetailRequest) (*XiaohongshuAppV2GetVideoNoteDetailResponse, error) {
	return r.client.XiaohongshuAppV2GetVideoNoteDetail(ctx, request)
}

// GetNoteComments 获取笔记评论列表/Get note comments
//
// GET /api/v1/xiaohongshu/app_v2/get_note_comments
func (r XiaohongshuAppV2Resource) GetNoteComments(ctx context.Context, request XiaohongshuAppV2GetNoteCommentsRequest) (*XiaohongshuAppV2GetNoteCommentsResponse, error) {
	return r.client.XiaohongshuAppV2GetNoteComments(ctx, request)
}

// GetNoteSubComments 获取笔记二级评论列表/Get note sub comments
//
// GET /api/v1/xiaohongshu/app_v2/get_note_sub_comments
func (r XiaohongshuAppV2Resource) GetNoteSubComments(ctx context.Context, request XiaohongshuAppV2GetNoteSubCommentsRequest) (*XiaohongshuAppV2GetNoteSubCommentsResponse, error) {
	return r.client.XiaohongshuAppV2GetNoteSubComments(ctx, request)
}

// GetUserInfo 获取用户信息/Get user info
//
// GET /api/v1/xiaohongshu/app_v2/get_user_info
func (r XiaohongshuAppV2Resource) GetUserInfo(ctx context.Context, request XiaohongshuAppV2GetUserInfoRequest) (*XiaohongshuAppV2GetUserInfoResponse, error) {
	return r.client.XiaohongshuAppV2GetUserInfo(ctx, request)
}

// GetUserPostedNotes 获取用户笔记列表/Get user posted notes
//
// GET /api/v1/xiaohongshu/app_v2/get_user_posted_notes
func (r XiaohongshuAppV2Resource) GetUserPostedNotes(ctx context.Context, request XiaohongshuAppV2GetUserPostedNotesRequest) (*XiaohongshuAppV2GetUserPostedNotesResponse, error) {
	return r.client.XiaohongshuAppV2GetUserPostedNotes(ctx, request)
}

// GetUserFavedNotes 获取用户收藏笔记列表/Get user faved notes
//
// GET /api/v1/xiaohongshu/app_v2/get_user_faved_notes
func (r XiaohongshuAppV2Resource) GetUserFavedNotes(ctx context.Context, request XiaohongshuAppV2GetUserFavedNotesRequest) (*XiaohongshuAppV2GetUserFavedNotesResponse, error) {
	return r.client.XiaohongshuAppV2GetUserFavedNotes(ctx, request)
}

// SearchNotes 搜索笔记/Search notes
//
// GET /api/v1/xiaohongshu/app_v2/search_notes
func (r XiaohongshuAppV2Resource) SearchNotes(ctx context.Context, request XiaohongshuAppV2SearchNotesRequest) (*XiaohongshuAppV2SearchNotesResponse, error) {
	return r.client.XiaohongshuAppV2SearchNotes(ctx, request)
}

// SearchUsers 搜索用户/Search users
//
// GET /api/v1/xiaohongshu/app_v2/search_users
func (r XiaohongshuAppV2Resource) SearchUsers(ctx context.Context, request XiaohongshuAppV2SearchUsersRequest) (*XiaohongshuAppV2SearchUsersResponse, error) {
	return r.client.XiaohongshuAppV2SearchUsers(ctx, request)
}

// SearchImages 搜索图片/Search images
//
// GET /api/v1/xiaohongshu/app_v2/search_images
func (r XiaohongshuAppV2Resource) SearchImages(ctx context.Context, request XiaohongshuAppV2SearchImagesRequest) (*XiaohongshuAppV2SearchImagesResponse, error) {
	return r.client.XiaohongshuAppV2SearchImages(ctx, request)
}

// SearchProducts 搜索商品/Search products
//
// GET /api/v1/xiaohongshu/app_v2/search_products
func (r XiaohongshuAppV2Resource) SearchProducts(ctx context.Context, request XiaohongshuAppV2SearchProductsRequest) (*XiaohongshuAppV2SearchProductsResponse, error) {
	return r.client.XiaohongshuAppV2SearchProducts(ctx, request)
}

// SearchGroups 搜索群聊/Search groups
//
// GET /api/v1/xiaohongshu/app_v2/search_groups
func (r XiaohongshuAppV2Resource) SearchGroups(ctx context.Context, request XiaohongshuAppV2SearchGroupsRequest) (*XiaohongshuAppV2SearchGroupsResponse, error) {
	return r.client.XiaohongshuAppV2SearchGroups(ctx, request)
}

// GetProductDetail 获取商品详情/Get product detail
//
// GET /api/v1/xiaohongshu/app_v2/get_product_detail
func (r XiaohongshuAppV2Resource) GetProductDetail(ctx context.Context, request XiaohongshuAppV2GetProductDetailRequest) (*XiaohongshuAppV2GetProductDetailResponse, error) {
	return r.client.XiaohongshuAppV2GetProductDetail(ctx, request)
}

// GetProductReviewOverview 获取商品评论总览/Get product review overview
//
// GET /api/v1/xiaohongshu/app_v2/get_product_review_overview
func (r XiaohongshuAppV2Resource) GetProductReviewOverview(ctx context.Context, request XiaohongshuAppV2GetProductReviewOverviewRequest) (*XiaohongshuAppV2GetProductReviewOverviewResponse, error) {
	return r.client.XiaohongshuAppV2GetProductReviewOverview(ctx, request)
}

// GetProductReviews 获取商品评论列表/Get product reviews
//
// GET /api/v1/xiaohongshu/app_v2/get_product_reviews
func (r XiaohongshuAppV2Resource) GetProductReviews(ctx context.Context, request XiaohongshuAppV2GetProductReviewsRequest) (*XiaohongshuAppV2GetProductReviewsResponse, error) {
	return r.client.XiaohongshuAppV2GetProductReviews(ctx, request)
}

// GetProductRecommendations 获取商品推荐列表/Get product recommendations
//
// GET /api/v1/xiaohongshu/app_v2/get_product_recommendations
func (r XiaohongshuAppV2Resource) GetProductRecommendations(ctx context.Context, request XiaohongshuAppV2GetProductRecommendationsRequest) (*XiaohongshuAppV2GetProductRecommendationsResponse, error) {
	return r.client.XiaohongshuAppV2GetProductRecommendations(ctx, request)
}

// GetTopicInfo 获取话题详情/Get topic info
//
// GET /api/v1/xiaohongshu/app_v2/get_topic_info
func (r XiaohongshuAppV2Resource) GetTopicInfo(ctx context.Context, request XiaohongshuAppV2GetTopicInfoRequest) (*XiaohongshuAppV2GetTopicInfoResponse, error) {
	return r.client.XiaohongshuAppV2GetTopicInfo(ctx, request)
}

// GetTopicFeed 获取话题笔记列表/Get topic feed
//
// GET /api/v1/xiaohongshu/app_v2/get_topic_feed
func (r XiaohongshuAppV2Resource) GetTopicFeed(ctx context.Context, request XiaohongshuAppV2GetTopicFeedRequest) (*XiaohongshuAppV2GetTopicFeedResponse, error) {
	return r.client.XiaohongshuAppV2GetTopicFeed(ctx, request)
}

// GetCreatorInspirationFeed 获取创作者推荐灵感列表/Get creator inspiration feed
//
// GET /api/v1/xiaohongshu/app_v2/get_creator_inspiration_feed
func (r XiaohongshuAppV2Resource) GetCreatorInspirationFeed(ctx context.Context, request XiaohongshuAppV2GetCreatorInspirationFeedRequest) (*XiaohongshuAppV2GetCreatorInspirationFeedResponse, error) {
	return r.client.XiaohongshuAppV2GetCreatorInspirationFeed(ctx, request)
}

// GetCreatorHotInspirationFeed 获取创作者热点灵感列表/Get creator hot inspiration feed
//
// GET /api/v1/xiaohongshu/app_v2/get_creator_hot_inspiration_feed
func (r XiaohongshuAppV2Resource) GetCreatorHotInspirationFeed(ctx context.Context, request XiaohongshuAppV2GetCreatorHotInspirationFeedRequest) (*XiaohongshuAppV2GetCreatorHotInspirationFeedResponse, error) {
	return r.client.XiaohongshuAppV2GetCreatorHotInspirationFeed(ctx, request)
}

// XiaohongshuAppResource contains endpoints from the Xiaohongshu-App-API tag.
type XiaohongshuAppResource struct {
	client *Client
}

// XiaohongshuAppGetNoteInfoRequest is the request for GET /api/v1/xiaohongshu/app/get_note_info.
type XiaohongshuAppGetNoteInfoRequest = XiaohongshuAppGetNoteInfoV1Request

// XiaohongshuAppGetNoteInfoResponse is the response for GET /api/v1/xiaohongshu/app/get_note_info.
type XiaohongshuAppGetNoteInfoResponse = XiaohongshuAppGetNoteInfoV1Response

// GetNoteInfo 获取笔记信息 V1/Get note info V1
//
// GET /api/v1/xiaohongshu/app/get_note_info
func (r XiaohongshuAppResource) GetNoteInfo(ctx context.Context, request XiaohongshuAppGetNoteInfoRequest) (*XiaohongshuAppGetNoteInfoResponse, error) {
	return r.client.XiaohongshuAppGetNoteInfoV1(ctx, request)
}

// GetNoteInfoV2 获取笔记信息 V2 (蒲公英商家后台)/Get note info V2 (Pugongying Business Backend)
//
// GET /api/v1/xiaohongshu/app/get_note_info_v2
func (r XiaohongshuAppResource) GetNoteInfoV2(ctx context.Context, request XiaohongshuAppGetNoteInfoV2Request) (*XiaohongshuAppGetNoteInfoV2Response, error) {
	return r.client.XiaohongshuAppGetNoteInfoV2(ctx, request)
}

// GetNoteComments 获取笔记评论/Get note comments
//
// GET /api/v1/xiaohongshu/app/get_note_comments
func (r XiaohongshuAppResource) GetNoteComments(ctx context.Context, request XiaohongshuAppGetNoteCommentsRequest) (*XiaohongshuAppGetNoteCommentsResponse, error) {
	return r.client.XiaohongshuAppGetNoteComments(ctx, request)
}

// GetSubComments 获取子评论/Get sub comments
//
// GET /api/v1/xiaohongshu/app/get_sub_comments
func (r XiaohongshuAppResource) GetSubComments(ctx context.Context, request XiaohongshuAppGetSubCommentsRequest) (*XiaohongshuAppGetSubCommentsResponse, error) {
	return r.client.XiaohongshuAppGetSubComments(ctx, request)
}

// XiaohongshuAppGetTopicNotesRequest is the request for GET /api/v1/xiaohongshu/app/get_topic_notes.
type XiaohongshuAppGetTopicNotesRequest = XiaohongshuAppGetNotesByTopicRequest

// XiaohongshuAppGetTopicNotesResponse is the response for GET /api/v1/xiaohongshu/app/get_topic_notes.
type XiaohongshuAppGetTopicNotesResponse = XiaohongshuAppGetNotesByTopicResponse

// GetTopicNotes 根据话题标签获取作品/Get notes by topic
//
// GET /api/v1/xiaohongshu/app/get_topic_notes
func (r XiaohongshuAppResource) GetTopicNotes(ctx context.Context, request XiaohongshuAppGetTopicNotesRequest) (*XiaohongshuAppGetTopicNotesResponse, error) {
	return r.client.XiaohongshuAppGetNotesByTopic(ctx, request)
}

// GetNotesByTopic [已弃用/Deprecated] 根据话题标签获取作品/Get notes by topic
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/xiaohongshu/app/get_notes_by_topic
func (r XiaohongshuAppResource) GetNotesByTopic(ctx context.Context, request XiaohongshuAppDeprecatedGetNotesByTopicRequest) (*XiaohongshuAppDeprecatedGetNotesByTopicResponse, error) {
	return r.client.XiaohongshuAppDeprecatedGetNotesByTopic(ctx, request)
}

// SearchNotes 搜索笔记/Search notes
//
// GET /api/v1/xiaohongshu/app/search_notes
func (r XiaohongshuAppResource) SearchNotes(ctx context.Context, request XiaohongshuAppSearchNotesRequest) (*XiaohongshuAppSearchNotesResponse, error) {
	return r.client.XiaohongshuAppSearchNotes(ctx, request)
}

// GetUserInfo 获取用户信息/Get user info
//
// GET /api/v1/xiaohongshu/app/get_user_info
func (r XiaohongshuAppResource) GetUserInfo(ctx context.Context, request XiaohongshuAppGetUserInfoRequest) (*XiaohongshuAppGetUserInfoResponse, error) {
	return r.client.XiaohongshuAppGetUserInfo(ctx, request)
}

// GetUserNotes 获取用户作品列表/Get user notes
//
// GET /api/v1/xiaohongshu/app/get_user_notes
func (r XiaohongshuAppResource) GetUserNotes(ctx context.Context, request XiaohongshuAppGetUserNotesRequest) (*XiaohongshuAppGetUserNotesResponse, error) {
	return r.client.XiaohongshuAppGetUserNotes(ctx, request)
}

// XiaohongshuAppExtractShareInfoRequest is the request for GET /api/v1/xiaohongshu/app/extract_share_info.
type XiaohongshuAppExtractShareInfoRequest = XiaohongshuAppExtractShareLinkInfoRequest

// XiaohongshuAppExtractShareInfoResponse is the response for GET /api/v1/xiaohongshu/app/extract_share_info.
type XiaohongshuAppExtractShareInfoResponse = XiaohongshuAppExtractShareLinkInfoResponse

// ExtractShareInfo 提取分享链接信息/Extract share link info
//
// GET /api/v1/xiaohongshu/app/extract_share_info
func (r XiaohongshuAppResource) ExtractShareInfo(ctx context.Context, request XiaohongshuAppExtractShareInfoRequest) (*XiaohongshuAppExtractShareInfoResponse, error) {
	return r.client.XiaohongshuAppExtractShareLinkInfo(ctx, request)
}

// XiaohongshuAppGetUserIDAndXsecTokenRequest is the request for GET /api/v1/xiaohongshu/app/get_user_id_and_xsec_token.
type XiaohongshuAppGetUserIDAndXsecTokenRequest = XiaohongshuAppExtractUserIDAndXsecTokenFromShareLinkRequest

// XiaohongshuAppGetUserIDAndXsecTokenResponse is the response for GET /api/v1/xiaohongshu/app/get_user_id_and_xsec_token.
type XiaohongshuAppGetUserIDAndXsecTokenResponse = XiaohongshuAppExtractUserIDAndXsecTokenFromShareLinkResponse

// GetUserIDAndXsecToken 从分享链接中提取用户ID和xsec_token/Extract user ID and xsec_token from share link
//
// GET /api/v1/xiaohongshu/app/get_user_id_and_xsec_token
func (r XiaohongshuAppResource) GetUserIDAndXsecToken(ctx context.Context, request XiaohongshuAppGetUserIDAndXsecTokenRequest) (*XiaohongshuAppGetUserIDAndXsecTokenResponse, error) {
	return r.client.XiaohongshuAppExtractUserIDAndXsecTokenFromShareLink(ctx, request)
}

// GetProductDetail 获取商品详情/Get product detail
//
// GET /api/v1/xiaohongshu/app/get_product_detail
func (r XiaohongshuAppResource) GetProductDetail(ctx context.Context, request XiaohongshuAppGetProductDetailRequest) (*XiaohongshuAppGetProductDetailResponse, error) {
	return r.client.XiaohongshuAppGetProductDetail(ctx, request)
}

// SearchProducts 搜索商品/Search products
//
// GET /api/v1/xiaohongshu/app/search_products
func (r XiaohongshuAppResource) SearchProducts(ctx context.Context, request XiaohongshuAppSearchProductsRequest) (*XiaohongshuAppSearchProductsResponse, error) {
	return r.client.XiaohongshuAppSearchProducts(ctx, request)
}

// XiaohongshuWebV2Resource contains endpoints from the Xiaohongshu-Web-V2-API tag.
type XiaohongshuWebV2Resource struct {
	client *Client
}

// XiaohongshuWebV2FetchFeedNotesRequest is the request for GET /api/v1/xiaohongshu/web_v2/fetch_feed_notes.
type XiaohongshuWebV2FetchFeedNotesRequest = XiaohongshuWebV2GetImageNoteDetailV1Request

// XiaohongshuWebV2FetchFeedNotesResponse is the response for GET /api/v1/xiaohongshu/web_v2/fetch_feed_notes.
type XiaohongshuWebV2FetchFeedNotesResponse = XiaohongshuWebV2GetImageNoteDetailV1Response

// FetchFeedNotes 获取图文笔记详情 V1/Get image note detail V1
//
// GET /api/v1/xiaohongshu/web_v2/fetch_feed_notes
func (r XiaohongshuWebV2Resource) FetchFeedNotes(ctx context.Context, request XiaohongshuWebV2FetchFeedNotesRequest) (*XiaohongshuWebV2FetchFeedNotesResponse, error) {
	return r.client.XiaohongshuWebV2GetImageNoteDetailV1(ctx, request)
}

// XiaohongshuWebV2FetchFeedNotesV2Request is the request for GET /api/v1/xiaohongshu/web_v2/fetch_feed_notes_v2.
type XiaohongshuWebV2FetchFeedNotesV2Request = XiaohongshuWebV2GetImageNoteDetailV2Request

// XiaohongshuWebV2FetchFeedNotesV2Response is the response for GET /api/v1/xiaohongshu/web_v2/fetch_feed_notes_v2.
type XiaohongshuWebV2FetchFeedNotesV2Response = XiaohongshuWebV2GetImageNoteDetailV2Response

// FetchFeedNotesV2 获取图文笔记详情 V2/Get image note detail V2
//
// GET /api/v1/xiaohongshu/web_v2/fetch_feed_notes_v2
func (r XiaohongshuWebV2Resource) FetchFeedNotesV2(ctx context.Context, request XiaohongshuWebV2FetchFeedNotesV2Request) (*XiaohongshuWebV2FetchFeedNotesV2Response, error) {
	return r.client.XiaohongshuWebV2GetImageNoteDetailV2(ctx, request)
}

// XiaohongshuWebV2FetchHomeNotesAppRequest is the request for GET /api/v1/xiaohongshu/web_v2/fetch_home_notes_app.
type XiaohongshuWebV2FetchHomeNotesAppRequest = XiaohongshuWebV2FetchUserNotesRequest

// XiaohongshuWebV2FetchHomeNotesAppResponse is the response for GET /api/v1/xiaohongshu/web_v2/fetch_home_notes_app.
type XiaohongshuWebV2FetchHomeNotesAppResponse = XiaohongshuWebV2FetchUserNotesResponse

// FetchHomeNotesApp 获取用户笔记/Fetch user notes
//
// GET /api/v1/xiaohongshu/web_v2/fetch_home_notes_app
func (r XiaohongshuWebV2Resource) FetchHomeNotesApp(ctx context.Context, request XiaohongshuWebV2FetchHomeNotesAppRequest) (*XiaohongshuWebV2FetchHomeNotesAppResponse, error) {
	return r.client.XiaohongshuWebV2FetchUserNotes(ctx, request)
}

// FetchNoteComments 获取笔记评论/Fetch note comments
//
// GET /api/v1/xiaohongshu/web_v2/fetch_note_comments
func (r XiaohongshuWebV2Resource) FetchNoteComments(ctx context.Context, request XiaohongshuWebV2FetchNoteCommentsRequest) (*XiaohongshuWebV2FetchNoteCommentsResponse, error) {
	return r.client.XiaohongshuWebV2FetchNoteComments(ctx, request)
}

// FetchSubComments 获取子评论/Fetch sub comments
//
// GET /api/v1/xiaohongshu/web_v2/fetch_sub_comments
func (r XiaohongshuWebV2Resource) FetchSubComments(ctx context.Context, request XiaohongshuWebV2FetchSubCommentsRequest) (*XiaohongshuWebV2FetchSubCommentsResponse, error) {
	return r.client.XiaohongshuWebV2FetchSubComments(ctx, request)
}

// XiaohongshuWebV2FetchUserInfoAppRequest is the request for GET /api/v1/xiaohongshu/web_v2/fetch_user_info_app.
type XiaohongshuWebV2FetchUserInfoAppRequest = XiaohongshuWebV2FetchAppUserInfoRequest

// XiaohongshuWebV2FetchUserInfoAppResponse is the response for GET /api/v1/xiaohongshu/web_v2/fetch_user_info_app.
type XiaohongshuWebV2FetchUserInfoAppResponse = XiaohongshuWebV2FetchAppUserInfoResponse

// FetchUserInfoApp 获取App用户信息/Fetch App user info
//
// GET /api/v1/xiaohongshu/web_v2/fetch_user_info_app
func (r XiaohongshuWebV2Resource) FetchUserInfoApp(ctx context.Context, request XiaohongshuWebV2FetchUserInfoAppRequest) (*XiaohongshuWebV2FetchUserInfoAppResponse, error) {
	return r.client.XiaohongshuWebV2FetchAppUserInfo(ctx, request)
}

// XiaohongshuWebV2FetchHotListResponse is the response for GET /api/v1/xiaohongshu/web_v2/fetch_hot_list.
type XiaohongshuWebV2FetchHotListResponse = XiaohongshuWebV2FetchXiaohongshuHotListResponse

// FetchHotList 获取小红书热榜/Fetch Xiaohongshu hot list
//
// GET /api/v1/xiaohongshu/web_v2/fetch_hot_list
func (r XiaohongshuWebV2Resource) FetchHotList(ctx context.Context) (*XiaohongshuWebV2FetchHotListResponse, error) {
	return r.client.XiaohongshuWebV2FetchXiaohongshuHotList(ctx)
}

// XiaohongshuWebResource contains endpoints from the Xiaohongshu-Web-API tag.
type XiaohongshuWebResource struct {
	client *Client
}

// GetHomeRecommend 获取首页推荐/Get home recommend
//
// POST /api/v1/xiaohongshu/web/get_home_recommend
func (r XiaohongshuWebResource) GetHomeRecommend(ctx context.Context, request XiaohongshuWebGetHomeRecommendRequest) (*XiaohongshuWebGetHomeRecommendResponse, error) {
	return r.client.XiaohongshuWebGetHomeRecommend(ctx, request)
}

// GetNoteInfoV4 获取笔记信息 V4/Get note info V4
//
// GET /api/v1/xiaohongshu/web/get_note_info_v4
func (r XiaohongshuWebResource) GetNoteInfoV4(ctx context.Context, request XiaohongshuWebGetNoteInfoV4Request) (*XiaohongshuWebGetNoteInfoV4Response, error) {
	return r.client.XiaohongshuWebGetNoteInfoV4(ctx, request)
}

// GetNoteInfoV5 获取笔记信息 V5 (自带Cookie)/Get note info V5 (Self-provided Cookie)
//
// POST /api/v1/xiaohongshu/web/get_note_info_v5
func (r XiaohongshuWebResource) GetNoteInfoV5(ctx context.Context, request XiaohongshuWebGetNoteInfoV5Request) (*XiaohongshuWebGetNoteInfoV5Response, error) {
	return r.client.XiaohongshuWebGetNoteInfoV5(ctx, request)
}

// GetNoteInfoV7 获取笔记信息 V7/Get note info V7
//
// GET /api/v1/xiaohongshu/web/get_note_info_v7
func (r XiaohongshuWebResource) GetNoteInfoV7(ctx context.Context, request XiaohongshuWebGetNoteInfoV7Request) (*XiaohongshuWebGetNoteInfoV7Response, error) {
	return r.client.XiaohongshuWebGetNoteInfoV7(ctx, request)
}

// XiaohongshuWebGetUserInfoRequest is the request for GET /api/v1/xiaohongshu/web/get_user_info.
type XiaohongshuWebGetUserInfoRequest = XiaohongshuWebGetUserInfoV1Request

// XiaohongshuWebGetUserInfoResponse is the response for GET /api/v1/xiaohongshu/web/get_user_info.
type XiaohongshuWebGetUserInfoResponse = XiaohongshuWebGetUserInfoV1Response

// GetUserInfo 获取用户信息 V1/Get user info V1
//
// GET /api/v1/xiaohongshu/web/get_user_info
func (r XiaohongshuWebResource) GetUserInfo(ctx context.Context, request XiaohongshuWebGetUserInfoRequest) (*XiaohongshuWebGetUserInfoResponse, error) {
	return r.client.XiaohongshuWebGetUserInfoV1(ctx, request)
}

// GetVisitorCookie 获取游客Cookie/Get visitor cookie
//
// GET /api/v1/xiaohongshu/web/get_visitor_cookie
func (r XiaohongshuWebResource) GetVisitorCookie(ctx context.Context, request XiaohongshuWebGetVisitorCookieRequest) (*XiaohongshuWebGetVisitorCookieResponse, error) {
	return r.client.XiaohongshuWebGetVisitorCookie(ctx, request)
}

// XiaohongshuWebSignRequest is the request for POST /api/v1/xiaohongshu/web/sign.
type XiaohongshuWebSignRequest = XiaohongshuWebXiaohongshuWebSignRequest

// XiaohongshuWebSignResponse is the response for POST /api/v1/xiaohongshu/web/sign.
type XiaohongshuWebSignResponse = XiaohongshuWebXiaohongshuWebSignResponse

// Sign 小红书Web签名/Xiaohongshu Web sign
//
// POST /api/v1/xiaohongshu/web/sign
func (r XiaohongshuWebResource) Sign(ctx context.Context, request XiaohongshuWebSignRequest) (*XiaohongshuWebSignResponse, error) {
	return r.client.XiaohongshuWebXiaohongshuWebSign(ctx, request)
}

// XiaohongshuWebGetNoteIDAndXsecTokenRequest is the request for GET /api/v1/xiaohongshu/web/get_note_id_and_xsec_token.
type XiaohongshuWebGetNoteIDAndXsecTokenRequest = XiaohongshuWebGetXiaohongshuNoteIDAndXsecTokenByShareLinkRequest

// XiaohongshuWebGetNoteIDAndXsecTokenResponse is the response for GET /api/v1/xiaohongshu/web/get_note_id_and_xsec_token.
type XiaohongshuWebGetNoteIDAndXsecTokenResponse = XiaohongshuWebGetXiaohongshuNoteIDAndXsecTokenByShareLinkResponse

// GetNoteIDAndXsecToken 通过分享链接获取小红书的Note ID 和 xsec_token/Get Xiaohongshu Note ID and xsec_token by share link
//
// GET /api/v1/xiaohongshu/web/get_note_id_and_xsec_token
func (r XiaohongshuWebResource) GetNoteIDAndXsecToken(ctx context.Context, request XiaohongshuWebGetNoteIDAndXsecTokenRequest) (*XiaohongshuWebGetNoteIDAndXsecTokenResponse, error) {
	return r.client.XiaohongshuWebGetXiaohongshuNoteIDAndXsecTokenByShareLink(ctx, request)
}

// XiaohongshuWebGetProductInfoRequest is the request for GET /api/v1/xiaohongshu/web/get_product_info.
type XiaohongshuWebGetProductInfoRequest = XiaohongshuWebGetXiaohongshuProductInfoRequest

// XiaohongshuWebGetProductInfoResponse is the response for GET /api/v1/xiaohongshu/web/get_product_info.
type XiaohongshuWebGetProductInfoResponse = XiaohongshuWebGetXiaohongshuProductInfoResponse

// GetProductInfo 获取小红书商品信息/Get Xiaohongshu product info
//
// GET /api/v1/xiaohongshu/web/get_product_info
func (r XiaohongshuWebResource) GetProductInfo(ctx context.Context, request XiaohongshuWebGetProductInfoRequest) (*XiaohongshuWebGetProductInfoResponse, error) {
	return r.client.XiaohongshuWebGetXiaohongshuProductInfo(ctx, request)
}

// Lemon8AppResource contains endpoints from the Lemon8-App-API tag.
type Lemon8AppResource struct {
	client *Client
}

// Lemon8AppFetchUserProfileRequest is the request for GET /api/v1/lemon8/app/fetch_user_profile.
type Lemon8AppFetchUserProfileRequest = Lemon8AppGetInformationOfSpecifiedUserRequest

// Lemon8AppFetchUserProfileResponse is the response for GET /api/v1/lemon8/app/fetch_user_profile.
type Lemon8AppFetchUserProfileResponse = Lemon8AppGetInformationOfSpecifiedUserResponse

// FetchUserProfile 获取指定用户的信息/Get information of specified user
//
// GET /api/v1/lemon8/app/fetch_user_profile
func (r Lemon8AppResource) FetchUserProfile(ctx context.Context, request Lemon8AppFetchUserProfileRequest) (*Lemon8AppFetchUserProfileResponse, error) {
	return r.client.Lemon8AppGetInformationOfSpecifiedUser(ctx, request)
}

// Lemon8AppFetchPostDetailRequest is the request for GET /api/v1/lemon8/app/fetch_post_detail.
type Lemon8AppFetchPostDetailRequest = Lemon8AppGetInformationOfSpecifiedPostRequest

// Lemon8AppFetchPostDetailResponse is the response for GET /api/v1/lemon8/app/fetch_post_detail.
type Lemon8AppFetchPostDetailResponse = Lemon8AppGetInformationOfSpecifiedPostResponse

// FetchPostDetail 获取指定作品的信息/Get information of specified post
//
// GET /api/v1/lemon8/app/fetch_post_detail
func (r Lemon8AppResource) FetchPostDetail(ctx context.Context, request Lemon8AppFetchPostDetailRequest) (*Lemon8AppFetchPostDetailResponse, error) {
	return r.client.Lemon8AppGetInformationOfSpecifiedPost(ctx, request)
}

// Lemon8AppFetchUserFollowerListRequest is the request for GET /api/v1/lemon8/app/fetch_user_follower_list.
type Lemon8AppFetchUserFollowerListRequest = Lemon8AppGetFansListOfSpecifiedUserRequest

// Lemon8AppFetchUserFollowerListResponse is the response for GET /api/v1/lemon8/app/fetch_user_follower_list.
type Lemon8AppFetchUserFollowerListResponse = Lemon8AppGetFansListOfSpecifiedUserResponse

// FetchUserFollowerList 获取指定用户的粉丝列表/Get fans list of specified user
//
// GET /api/v1/lemon8/app/fetch_user_follower_list
func (r Lemon8AppResource) FetchUserFollowerList(ctx context.Context, request Lemon8AppFetchUserFollowerListRequest) (*Lemon8AppFetchUserFollowerListResponse, error) {
	return r.client.Lemon8AppGetFansListOfSpecifiedUser(ctx, request)
}

// Lemon8AppFetchUserFollowingListRequest is the request for GET /api/v1/lemon8/app/fetch_user_following_list.
type Lemon8AppFetchUserFollowingListRequest = Lemon8AppGetFollowingListOfSpecifiedUserRequest

// Lemon8AppFetchUserFollowingListResponse is the response for GET /api/v1/lemon8/app/fetch_user_following_list.
type Lemon8AppFetchUserFollowingListResponse = Lemon8AppGetFollowingListOfSpecifiedUserResponse

// FetchUserFollowingList 获取指定用户的关注列表/Get following list of specified user
//
// GET /api/v1/lemon8/app/fetch_user_following_list
func (r Lemon8AppResource) FetchUserFollowingList(ctx context.Context, request Lemon8AppFetchUserFollowingListRequest) (*Lemon8AppFetchUserFollowingListResponse, error) {
	return r.client.Lemon8AppGetFollowingListOfSpecifiedUser(ctx, request)
}

// Lemon8AppFetchPostCommentListRequest is the request for GET /api/v1/lemon8/app/fetch_post_comment_list.
type Lemon8AppFetchPostCommentListRequest = Lemon8AppGetCommentsListOfSpecifiedPostRequest

// Lemon8AppFetchPostCommentListResponse is the response for GET /api/v1/lemon8/app/fetch_post_comment_list.
type Lemon8AppFetchPostCommentListResponse = Lemon8AppGetCommentsListOfSpecifiedPostResponse

// FetchPostCommentList 获取指定作品的评论列表/Get comments list of specified post
//
// GET /api/v1/lemon8/app/fetch_post_comment_list
func (r Lemon8AppResource) FetchPostCommentList(ctx context.Context, request Lemon8AppFetchPostCommentListRequest) (*Lemon8AppFetchPostCommentListResponse, error) {
	return r.client.Lemon8AppGetCommentsListOfSpecifiedPost(ctx, request)
}

// Lemon8AppFetchDiscoverBannersResponse is the response for GET /api/v1/lemon8/app/fetch_discover_banners.
type Lemon8AppFetchDiscoverBannersResponse = Lemon8AppGetBannersOfDiscoverPageResponse

// FetchDiscoverBanners 获取发现页Banner/Get banners of discover page
//
// GET /api/v1/lemon8/app/fetch_discover_banners
func (r Lemon8AppResource) FetchDiscoverBanners(ctx context.Context) (*Lemon8AppFetchDiscoverBannersResponse, error) {
	return r.client.Lemon8AppGetBannersOfDiscoverPage(ctx)
}

// Lemon8AppFetchDiscoverTabResponse is the response for GET /api/v1/lemon8/app/fetch_discover_tab.
type Lemon8AppFetchDiscoverTabResponse = Lemon8AppGetMainContentOfDiscoverPageResponse

// FetchDiscoverTab 获取发现页主体内容/Get main content of discover page
//
// GET /api/v1/lemon8/app/fetch_discover_tab
func (r Lemon8AppResource) FetchDiscoverTab(ctx context.Context) (*Lemon8AppFetchDiscoverTabResponse, error) {
	return r.client.Lemon8AppGetMainContentOfDiscoverPage(ctx)
}

// Lemon8AppFetchDiscoverTabInformationTabsResponse is the response for GET /api/v1/lemon8/app/fetch_discover_tab_information_tabs.
type Lemon8AppFetchDiscoverTabInformationTabsResponse = Lemon8AppGetEditorSPicksOfDiscoverPageResponse

// FetchDiscoverTabInformationTabs 获取发现页的 Editor's Picks/Get Editor's Picks of discover page
//
// GET /api/v1/lemon8/app/fetch_discover_tab_information_tabs
func (r Lemon8AppResource) FetchDiscoverTabInformationTabs(ctx context.Context) (*Lemon8AppFetchDiscoverTabInformationTabsResponse, error) {
	return r.client.Lemon8AppGetEditorSPicksOfDiscoverPage(ctx)
}

// Lemon8AppFetchHotSearchKeywordsResponse is the response for GET /api/v1/lemon8/app/fetch_hot_search_keywords.
type Lemon8AppFetchHotSearchKeywordsResponse = Lemon8AppGetHotSearchKeywordsResponse

// FetchHotSearchKeywords 获取热搜关键词/Get hot search keywords
//
// GET /api/v1/lemon8/app/fetch_hot_search_keywords
func (r Lemon8AppResource) FetchHotSearchKeywords(ctx context.Context) (*Lemon8AppFetchHotSearchKeywordsResponse, error) {
	return r.client.Lemon8AppGetHotSearchKeywords(ctx)
}

// Lemon8AppFetchTopicInfoRequest is the request for GET /api/v1/lemon8/app/fetch_topic_info.
type Lemon8AppFetchTopicInfoRequest = Lemon8AppGetTopicInformationRequest

// Lemon8AppFetchTopicInfoResponse is the response for GET /api/v1/lemon8/app/fetch_topic_info.
type Lemon8AppFetchTopicInfoResponse = Lemon8AppGetTopicInformationResponse

// FetchTopicInfo 获取话题信息/Get topic information
//
// GET /api/v1/lemon8/app/fetch_topic_info
func (r Lemon8AppResource) FetchTopicInfo(ctx context.Context, request Lemon8AppFetchTopicInfoRequest) (*Lemon8AppFetchTopicInfoResponse, error) {
	return r.client.Lemon8AppGetTopicInformation(ctx, request)
}

// Lemon8AppFetchTopicPostListRequest is the request for GET /api/v1/lemon8/app/fetch_topic_post_list.
type Lemon8AppFetchTopicPostListRequest = Lemon8AppGetTopicPostListRequest

// Lemon8AppFetchTopicPostListResponse is the response for GET /api/v1/lemon8/app/fetch_topic_post_list.
type Lemon8AppFetchTopicPostListResponse = Lemon8AppGetTopicPostListResponse

// FetchTopicPostList 获取话题作品列表/Get topic post list
//
// GET /api/v1/lemon8/app/fetch_topic_post_list
func (r Lemon8AppResource) FetchTopicPostList(ctx context.Context, request Lemon8AppFetchTopicPostListRequest) (*Lemon8AppFetchTopicPostListResponse, error) {
	return r.client.Lemon8AppGetTopicPostList(ctx, request)
}

// Lemon8AppFetchSearchRequest is the request for GET /api/v1/lemon8/app/fetch_search.
type Lemon8AppFetchSearchRequest = Lemon8AppSearchAPIRequest

// Lemon8AppFetchSearchResponse is the response for GET /api/v1/lemon8/app/fetch_search.
type Lemon8AppFetchSearchResponse = Lemon8AppSearchAPIResponse

// FetchSearch 搜索接口/Search API
//
// GET /api/v1/lemon8/app/fetch_search
func (r Lemon8AppResource) FetchSearch(ctx context.Context, request Lemon8AppFetchSearchRequest) (*Lemon8AppFetchSearchResponse, error) {
	return r.client.Lemon8AppSearchAPI(ctx, request)
}

// Lemon8AppGetItemIDRequest is the request for GET /api/v1/lemon8/app/get_item_id.
type Lemon8AppGetItemIDRequest = Lemon8AppGetPostIDThroughSharingLinkRequest

// Lemon8AppGetItemIDResponse is the response for GET /api/v1/lemon8/app/get_item_id.
type Lemon8AppGetItemIDResponse = Lemon8AppGetPostIDThroughSharingLinkResponse

// GetItemID 通过分享链接获取作品ID/Get post ID through sharing link
//
// GET /api/v1/lemon8/app/get_item_id
func (r Lemon8AppResource) GetItemID(ctx context.Context, request Lemon8AppGetItemIDRequest) (*Lemon8AppGetItemIDResponse, error) {
	return r.client.Lemon8AppGetPostIDThroughSharingLink(ctx, request)
}

// Lemon8AppGetUserIDRequest is the request for GET /api/v1/lemon8/app/get_user_id.
type Lemon8AppGetUserIDRequest = Lemon8AppGetUserIDThroughSharingLinkRequest

// Lemon8AppGetUserIDResponse is the response for GET /api/v1/lemon8/app/get_user_id.
type Lemon8AppGetUserIDResponse = Lemon8AppGetUserIDThroughSharingLinkResponse

// GetUserID 通过分享链接获取用户ID/Get user ID through sharing link
//
// GET /api/v1/lemon8/app/get_user_id
func (r Lemon8AppResource) GetUserID(ctx context.Context, request Lemon8AppGetUserIDRequest) (*Lemon8AppGetUserIDResponse, error) {
	return r.client.Lemon8AppGetUserIDThroughSharingLink(ctx, request)
}

// Lemon8AppGetItemIdsRequest is the request for POST /api/v1/lemon8/app/get_item_ids.
type Lemon8AppGetItemIdsRequest = Lemon8AppGetPostIDsInBatchThroughSharingLinksRequest

// Lemon8AppGetItemIdsResponse is the response for POST /api/v1/lemon8/app/get_item_ids.
type Lemon8AppGetItemIdsResponse = Lemon8AppGetPostIDsInBatchThroughSharingLinksResponse

// GetItemIds 通过分享链接批量获取作品ID/Get post IDs in batch through sharing links
//
// POST /api/v1/lemon8/app/get_item_ids
func (r Lemon8AppResource) GetItemIds(ctx context.Context, request Lemon8AppGetItemIdsRequest) (*Lemon8AppGetItemIdsResponse, error) {
	return r.client.Lemon8AppGetPostIDsInBatchThroughSharingLinks(ctx, request)
}

// Lemon8AppGetUserIdsRequest is the request for POST /api/v1/lemon8/app/get_user_ids.
type Lemon8AppGetUserIdsRequest = Lemon8AppGetUserIDsInBatchThroughSharingLinksRequest

// Lemon8AppGetUserIdsResponse is the response for POST /api/v1/lemon8/app/get_user_ids.
type Lemon8AppGetUserIdsResponse = Lemon8AppGetUserIDsInBatchThroughSharingLinksResponse

// GetUserIds 通过分享链接批量获取用户ID/Get user IDs in batch through sharing links
//
// POST /api/v1/lemon8/app/get_user_ids
func (r Lemon8AppResource) GetUserIds(ctx context.Context, request Lemon8AppGetUserIdsRequest) (*Lemon8AppGetUserIdsResponse, error) {
	return r.client.Lemon8AppGetUserIDsInBatchThroughSharingLinks(ctx, request)
}

// KuaishouWebResource contains endpoints from the Kuaishou-Web-API tag.
type KuaishouWebResource struct {
	client *Client
}

// KuaishouWebFetchOneVideoRequest is the request for GET /api/v1/kuaishou/web/fetch_one_video.
type KuaishouWebFetchOneVideoRequest = KuaishouWebGetSingleVideoDataV1Request

// KuaishouWebFetchOneVideoResponse is the response for GET /api/v1/kuaishou/web/fetch_one_video.
type KuaishouWebFetchOneVideoResponse = KuaishouWebGetSingleVideoDataV1Response

// FetchOneVideo 获取单个作品数据 V1/Get single video data V1
//
// GET /api/v1/kuaishou/web/fetch_one_video
func (r KuaishouWebResource) FetchOneVideo(ctx context.Context, request KuaishouWebFetchOneVideoRequest) (*KuaishouWebFetchOneVideoResponse, error) {
	return r.client.KuaishouWebGetSingleVideoDataV1(ctx, request)
}

// KuaishouWebFetchOneVideoV2Request is the request for GET /api/v1/kuaishou/web/fetch_one_video_v2.
type KuaishouWebFetchOneVideoV2Request = KuaishouWebGetSingleVideoDataV2Request

// KuaishouWebFetchOneVideoV2Response is the response for GET /api/v1/kuaishou/web/fetch_one_video_v2.
type KuaishouWebFetchOneVideoV2Response = KuaishouWebGetSingleVideoDataV2Response

// FetchOneVideoV2 获取单个作品数据 V2/Get single video data V2
//
// GET /api/v1/kuaishou/web/fetch_one_video_v2
func (r KuaishouWebResource) FetchOneVideoV2(ctx context.Context, request KuaishouWebFetchOneVideoV2Request) (*KuaishouWebFetchOneVideoV2Response, error) {
	return r.client.KuaishouWebGetSingleVideoDataV2(ctx, request)
}

// KuaishouWebFetchOneVideoByURLRequest is the request for GET /api/v1/kuaishou/web/fetch_one_video_by_url.
type KuaishouWebFetchOneVideoByURLRequest = KuaishouWebFetchSingleVideoByURLRequest

// KuaishouWebFetchOneVideoByURLResponse is the response for GET /api/v1/kuaishou/web/fetch_one_video_by_url.
type KuaishouWebFetchOneVideoByURLResponse = KuaishouWebFetchSingleVideoByURLResponse

// FetchOneVideoByURL 链接获取作品数据/Fetch single video by URL
//
// GET /api/v1/kuaishou/web/fetch_one_video_by_url
func (r KuaishouWebResource) FetchOneVideoByURL(ctx context.Context, request KuaishouWebFetchOneVideoByURLRequest) (*KuaishouWebFetchOneVideoByURLResponse, error) {
	return r.client.KuaishouWebFetchSingleVideoByURL(ctx, request)
}

// KuaishouWebFetchOneVideoCommentRequest is the request for GET /api/v1/kuaishou/web/fetch_one_video_comment.
type KuaishouWebFetchOneVideoCommentRequest = KuaishouWebFetchVideoCommentsRequest

// KuaishouWebFetchOneVideoCommentResponse is the response for GET /api/v1/kuaishou/web/fetch_one_video_comment.
type KuaishouWebFetchOneVideoCommentResponse = KuaishouWebFetchVideoCommentsResponse

// FetchOneVideoComment 获取作品一级评论/Fetch video comments
//
// GET /api/v1/kuaishou/web/fetch_one_video_comment
func (r KuaishouWebResource) FetchOneVideoComment(ctx context.Context, request KuaishouWebFetchOneVideoCommentRequest) (*KuaishouWebFetchOneVideoCommentResponse, error) {
	return r.client.KuaishouWebFetchVideoComments(ctx, request)
}

// KuaishouWebFetchOneVideoSubCommentRequest is the request for GET /api/v1/kuaishou/web/fetch_one_video_sub_comment.
type KuaishouWebFetchOneVideoSubCommentRequest = KuaishouWebFetchVideoSubCommentsRequest

// KuaishouWebFetchOneVideoSubCommentResponse is the response for GET /api/v1/kuaishou/web/fetch_one_video_sub_comment.
type KuaishouWebFetchOneVideoSubCommentResponse = KuaishouWebFetchVideoSubCommentsResponse

// FetchOneVideoSubComment 获取作品二级评论/Fetch video sub comments
//
// GET /api/v1/kuaishou/web/fetch_one_video_sub_comment
func (r KuaishouWebResource) FetchOneVideoSubComment(ctx context.Context, request KuaishouWebFetchOneVideoSubCommentRequest) (*KuaishouWebFetchOneVideoSubCommentResponse, error) {
	return r.client.KuaishouWebFetchVideoSubComments(ctx, request)
}

// GenerateShareShortURL 生成分享短连接/Generate share short URL
//
// GET /api/v1/kuaishou/web/generate_share_short_url
func (r KuaishouWebResource) GenerateShareShortURL(ctx context.Context, request KuaishouWebGenerateShareShortURLRequest) (*KuaishouWebGenerateShareShortURLResponse, error) {
	return r.client.KuaishouWebGenerateShareShortURL(ctx, request)
}

// FetchUserInfo 获取用户信息/Fetch user info
//
// GET /api/v1/kuaishou/web/fetch_user_info
func (r KuaishouWebResource) FetchUserInfo(ctx context.Context, request KuaishouWebFetchUserInfoRequest) (*KuaishouWebFetchUserInfoResponse, error) {
	return r.client.KuaishouWebFetchUserInfo(ctx, request)
}

// KuaishouWebFetchUserPostRequest is the request for GET /api/v1/kuaishou/web/fetch_user_post.
type KuaishouWebFetchUserPostRequest = KuaishouWebFetchUserPostsRequest

// KuaishouWebFetchUserPostResponse is the response for GET /api/v1/kuaishou/web/fetch_user_post.
type KuaishouWebFetchUserPostResponse = KuaishouWebFetchUserPostsResponse

// FetchUserPost 获取用户发布作品/Fetch user posts
//
// GET /api/v1/kuaishou/web/fetch_user_post
func (r KuaishouWebResource) FetchUserPost(ctx context.Context, request KuaishouWebFetchUserPostRequest) (*KuaishouWebFetchUserPostResponse, error) {
	return r.client.KuaishouWebFetchUserPosts(ctx, request)
}

// FetchUserLiveReplay 获取用户直播回放/Fetch user live replay
//
// GET /api/v1/kuaishou/web/fetch_user_live_replay
func (r KuaishouWebResource) FetchUserLiveReplay(ctx context.Context, request KuaishouWebFetchUserLiveReplayRequest) (*KuaishouWebFetchUserLiveReplayResponse, error) {
	return r.client.KuaishouWebFetchUserLiveReplay(ctx, request)
}

// FetchUserCollect 获取用户收藏作品/Fetch user collect
//
// GET /api/v1/kuaishou/web/fetch_user_collect
func (r KuaishouWebResource) FetchUserCollect(ctx context.Context, request KuaishouWebFetchUserCollectRequest) (*KuaishouWebFetchUserCollectResponse, error) {
	return r.client.KuaishouWebFetchUserCollect(ctx, request)
}

// FetchKuaishouHotListV1 获取快手热榜 V1/Fetch Kuaishou Hot List V1
//
// GET /api/v1/kuaishou/web/fetch_kuaishou_hot_list_v1
func (r KuaishouWebResource) FetchKuaishouHotListV1(ctx context.Context) (*KuaishouWebFetchKuaishouHotListV1Response, error) {
	return r.client.KuaishouWebFetchKuaishouHotListV1(ctx)
}

// FetchKuaishouHotListV2 获取快手热榜 V2/Fetch Kuaishou Hot List V2
//
// GET /api/v1/kuaishou/web/fetch_kuaishou_hot_list_v2
func (r KuaishouWebResource) FetchKuaishouHotListV2(ctx context.Context, request KuaishouWebFetchKuaishouHotListV2Request) (*KuaishouWebFetchKuaishouHotListV2Response, error) {
	return r.client.KuaishouWebFetchKuaishouHotListV2(ctx, request)
}

// KuaishouWebFetchGetUserIDRequest is the request for GET /api/v1/kuaishou/web/fetch_get_user_id.
type KuaishouWebFetchGetUserIDRequest = KuaishouWebFetchUserIDRequest

// KuaishouWebFetchGetUserIDResponse is the response for GET /api/v1/kuaishou/web/fetch_get_user_id.
type KuaishouWebFetchGetUserIDResponse = KuaishouWebFetchUserIDResponse

// FetchGetUserID 获取用户ID/Fetch user ID
//
// GET /api/v1/kuaishou/web/fetch_get_user_id
func (r KuaishouWebResource) FetchGetUserID(ctx context.Context, request KuaishouWebFetchGetUserIDRequest) (*KuaishouWebFetchGetUserIDResponse, error) {
	return r.client.KuaishouWebFetchUserID(ctx, request)
}

// KuaishouAppResource contains endpoints from the Kuaishou-App-API tag.
type KuaishouAppResource struct {
	client *Client
}

// KuaishouAppFetchOneVideoRequest is the request for GET /api/v1/kuaishou/app/fetch_one_video.
type KuaishouAppFetchOneVideoRequest = KuaishouAppVideoDetailsV1Request

// KuaishouAppFetchOneVideoResponse is the response for GET /api/v1/kuaishou/app/fetch_one_video.
type KuaishouAppFetchOneVideoResponse = KuaishouAppVideoDetailsV1Response

// FetchOneVideo 视频详情V1/Video detailsV1
//
// GET /api/v1/kuaishou/app/fetch_one_video
func (r KuaishouAppResource) FetchOneVideo(ctx context.Context, request KuaishouAppFetchOneVideoRequest) (*KuaishouAppFetchOneVideoResponse, error) {
	return r.client.KuaishouAppVideoDetailsV1(ctx, request)
}

// KuaishouAppFetchVideosBatchRequest is the request for GET /api/v1/kuaishou/app/fetch_videos_batch.
type KuaishouAppFetchVideosBatchRequest = KuaishouAppKuaishouBatchVideoQueryAPIRequest

// KuaishouAppFetchVideosBatchResponse is the response for GET /api/v1/kuaishou/app/fetch_videos_batch.
type KuaishouAppFetchVideosBatchResponse = KuaishouAppKuaishouBatchVideoQueryAPIResponse

// FetchVideosBatch 快手批量视频查询接口/Kuaishou batch video query API
//
// GET /api/v1/kuaishou/app/fetch_videos_batch
func (r KuaishouAppResource) FetchVideosBatch(ctx context.Context, request KuaishouAppFetchVideosBatchRequest) (*KuaishouAppFetchVideosBatchResponse, error) {
	return r.client.KuaishouAppKuaishouBatchVideoQueryAPI(ctx, request)
}

// KuaishouAppFetchOneVideoByURLRequest is the request for GET /api/v1/kuaishou/app/fetch_one_video_by_url.
type KuaishouAppFetchOneVideoByURLRequest = KuaishouAppFetchSingleVideoByURLRequest

// KuaishouAppFetchOneVideoByURLResponse is the response for GET /api/v1/kuaishou/app/fetch_one_video_by_url.
type KuaishouAppFetchOneVideoByURLResponse = KuaishouAppFetchSingleVideoByURLResponse

// FetchOneVideoByURL 根据链接获取单个作品数据/Fetch single video by URL
//
// GET /api/v1/kuaishou/app/fetch_one_video_by_url
func (r KuaishouAppResource) FetchOneVideoByURL(ctx context.Context, request KuaishouAppFetchOneVideoByURLRequest) (*KuaishouAppFetchOneVideoByURLResponse, error) {
	return r.client.KuaishouAppFetchSingleVideoByURL(ctx, request)
}

// KuaishouAppFetchSelectionFeedResponse is the response for GET /api/v1/kuaishou/app/fetch_selection_feed.
type KuaishouAppFetchSelectionFeedResponse = KuaishouAppFeedSelectionFeedResponse

// FetchSelectionFeed 精选/推荐Feed流/Selection feed
//
// GET /api/v1/kuaishou/app/fetch_selection_feed
func (r KuaishouAppResource) FetchSelectionFeed(ctx context.Context) (*KuaishouAppFetchSelectionFeedResponse, error) {
	return r.client.KuaishouAppFeedSelectionFeed(ctx)
}

// GenerateKuaishouShareLink 生成快手分享链接/Generate Kuaishou share link
//
// GET /api/v1/kuaishou/app/generate_kuaishou_share_link
func (r KuaishouAppResource) GenerateKuaishouShareLink(ctx context.Context, request KuaishouAppGenerateKuaishouShareLinkRequest) (*KuaishouAppGenerateKuaishouShareLinkResponse, error) {
	return r.client.KuaishouAppGenerateKuaishouShareLink(ctx, request)
}

// KuaishouAppFetchVideoCommentRequest is the request for GET /api/v1/kuaishou/app/fetch_video_comment.
type KuaishouAppFetchVideoCommentRequest = KuaishouAppGetSingleVideoCommentDataRequest

// KuaishouAppFetchVideoCommentResponse is the response for GET /api/v1/kuaishou/app/fetch_video_comment.
type KuaishouAppFetchVideoCommentResponse = KuaishouAppGetSingleVideoCommentDataResponse

// FetchVideoComment 获取单个作品评论数据/Get single video comment data
//
// GET /api/v1/kuaishou/app/fetch_video_comment
func (r KuaishouAppResource) FetchVideoComment(ctx context.Context, request KuaishouAppFetchVideoCommentRequest) (*KuaishouAppFetchVideoCommentResponse, error) {
	return r.client.KuaishouAppGetSingleVideoCommentData(ctx, request)
}

// KuaishouAppFetchVideoSubCommentsRequest is the request for GET /api/v1/kuaishou/app/fetch_video_sub_comments.
type KuaishouAppFetchVideoSubCommentsRequest = KuaishouAppVideoSubCommentsRequest

// KuaishouAppFetchVideoSubCommentsResponse is the response for GET /api/v1/kuaishou/app/fetch_video_sub_comments.
type KuaishouAppFetchVideoSubCommentsResponse = KuaishouAppVideoSubCommentsResponse

// FetchVideoSubComments 评论二级回复/Video sub comments
//
// GET /api/v1/kuaishou/app/fetch_video_sub_comments
func (r KuaishouAppResource) FetchVideoSubComments(ctx context.Context, request KuaishouAppFetchVideoSubCommentsRequest) (*KuaishouAppFetchVideoSubCommentsResponse, error) {
	return r.client.KuaishouAppVideoSubComments(ctx, request)
}

// KuaishouAppFetchOneUserV2Request is the request for GET /api/v1/kuaishou/app/fetch_one_user_v2.
type KuaishouAppFetchOneUserV2Request = KuaishouAppGetSingleUserDataV2Request

// KuaishouAppFetchOneUserV2Response is the response for GET /api/v1/kuaishou/app/fetch_one_user_v2.
type KuaishouAppFetchOneUserV2Response = KuaishouAppGetSingleUserDataV2Response

// FetchOneUserV2 获取单个用户数据V2/Get single user data V2
//
// GET /api/v1/kuaishou/app/fetch_one_user_v2
func (r KuaishouAppResource) FetchOneUserV2(ctx context.Context, request KuaishouAppFetchOneUserV2Request) (*KuaishouAppFetchOneUserV2Response, error) {
	return r.client.KuaishouAppGetSingleUserDataV2(ctx, request)
}

// KuaishouAppFetchUserPostV2Request is the request for GET /api/v1/kuaishou/app/fetch_user_post_v2.
type KuaishouAppFetchUserPostV2Request = KuaishouAppUserVideoListV2Request

// KuaishouAppFetchUserPostV2Response is the response for GET /api/v1/kuaishou/app/fetch_user_post_v2.
type KuaishouAppFetchUserPostV2Response = KuaishouAppUserVideoListV2Response

// FetchUserPostV2 用户视频列表V2/User video list V2
//
// GET /api/v1/kuaishou/app/fetch_user_post_v2
func (r KuaishouAppResource) FetchUserPostV2(ctx context.Context, request KuaishouAppFetchUserPostV2Request) (*KuaishouAppFetchUserPostV2Response, error) {
	return r.client.KuaishouAppUserVideoListV2(ctx, request)
}

// KuaishouAppFetchUserHotPostRequest is the request for GET /api/v1/kuaishou/app/fetch_user_hot_post.
type KuaishouAppFetchUserHotPostRequest = KuaishouAppGetUserHotPostDataRequest

// KuaishouAppFetchUserHotPostResponse is the response for GET /api/v1/kuaishou/app/fetch_user_hot_post.
type KuaishouAppFetchUserHotPostResponse = KuaishouAppGetUserHotPostDataResponse

// FetchUserHotPost 获取用户热门作品数据/Get user hot post data
//
// GET /api/v1/kuaishou/app/fetch_user_hot_post
func (r KuaishouAppResource) FetchUserHotPost(ctx context.Context, request KuaishouAppFetchUserHotPostRequest) (*KuaishouAppFetchUserHotPostResponse, error) {
	return r.client.KuaishouAppGetUserHotPostData(ctx, request)
}

// KuaishouAppFetchUserLiveInfoRequest is the request for GET /api/v1/kuaishou/app/fetch_user_live_info.
type KuaishouAppFetchUserLiveInfoRequest = KuaishouAppGetUserLiveInfoRequest

// KuaishouAppFetchUserLiveInfoResponse is the response for GET /api/v1/kuaishou/app/fetch_user_live_info.
type KuaishouAppFetchUserLiveInfoResponse = KuaishouAppGetUserLiveInfoResponse

// FetchUserLiveInfo 获取用户直播信息/Get user live info
//
// GET /api/v1/kuaishou/app/fetch_user_live_info
func (r KuaishouAppResource) FetchUserLiveInfo(ctx context.Context, request KuaishouAppFetchUserLiveInfoRequest) (*KuaishouAppFetchUserLiveInfoResponse, error) {
	return r.client.KuaishouAppGetUserLiveInfo(ctx, request)
}

// KuaishouAppSearchComprehensiveRequest is the request for GET /api/v1/kuaishou/app/search_comprehensive.
type KuaishouAppSearchComprehensiveRequest = KuaishouAppComprehensiveSearchRequest

// KuaishouAppSearchComprehensiveResponse is the response for GET /api/v1/kuaishou/app/search_comprehensive.
type KuaishouAppSearchComprehensiveResponse = KuaishouAppComprehensiveSearchResponse

// SearchComprehensive 综合搜索/Comprehensive search
//
// GET /api/v1/kuaishou/app/search_comprehensive
func (r KuaishouAppResource) SearchComprehensive(ctx context.Context, request KuaishouAppSearchComprehensiveRequest) (*KuaishouAppSearchComprehensiveResponse, error) {
	return r.client.KuaishouAppComprehensiveSearch(ctx, request)
}

// SearchVideoV2 搜索视频V2/Search video V2
//
// GET /api/v1/kuaishou/app/search_video_v2
func (r KuaishouAppResource) SearchVideoV2(ctx context.Context, request KuaishouAppSearchVideoV2Request) (*KuaishouAppSearchVideoV2Response, error) {
	return r.client.KuaishouAppSearchVideoV2(ctx, request)
}

// SearchUserV2 搜索用户V2/Search user V2
//
// GET /api/v1/kuaishou/app/search_user_v2
func (r KuaishouAppResource) SearchUserV2(ctx context.Context, request KuaishouAppSearchUserV2Request) (*KuaishouAppSearchUserV2Response, error) {
	return r.client.KuaishouAppSearchUserV2(ctx, request)
}

// SearchImage 搜索图片作品/Search image
//
// GET /api/v1/kuaishou/app/search_image
func (r KuaishouAppResource) SearchImage(ctx context.Context, request KuaishouAppSearchImageRequest) (*KuaishouAppSearchImageResponse, error) {
	return r.client.KuaishouAppSearchImage(ctx, request)
}

// SearchLive 搜索直播间/Search live
//
// GET /api/v1/kuaishou/app/search_live
func (r KuaishouAppResource) SearchLive(ctx context.Context, request KuaishouAppSearchLiveRequest) (*KuaishouAppSearchLiveResponse, error) {
	return r.client.KuaishouAppSearchLive(ctx, request)
}

// SearchMusic 搜索音乐/Search music
//
// GET /api/v1/kuaishou/app/search_music
func (r KuaishouAppResource) SearchMusic(ctx context.Context, request KuaishouAppSearchMusicRequest) (*KuaishouAppSearchMusicResponse, error) {
	return r.client.KuaishouAppSearchMusic(ctx, request)
}

// SearchTag 搜索话题标签/Search tag
//
// GET /api/v1/kuaishou/app/search_tag
func (r KuaishouAppResource) SearchTag(ctx context.Context, request KuaishouAppSearchTagRequest) (*KuaishouAppSearchTagResponse, error) {
	return r.client.KuaishouAppSearchTag(ctx, request)
}

// KuaishouAppFetchTagFeedRequest is the request for GET /api/v1/kuaishou/app/fetch_tag_feed.
type KuaishouAppFetchTagFeedRequest = KuaishouAppTagFeedRequest

// KuaishouAppFetchTagFeedResponse is the response for GET /api/v1/kuaishou/app/fetch_tag_feed.
type KuaishouAppFetchTagFeedResponse = KuaishouAppTagFeedResponse

// FetchTagFeed 话题标签聚合页/Tag feed
//
// GET /api/v1/kuaishou/app/fetch_tag_feed
func (r KuaishouAppResource) FetchTagFeed(ctx context.Context, request KuaishouAppFetchTagFeedRequest) (*KuaishouAppFetchTagFeedResponse, error) {
	return r.client.KuaishouAppTagFeed(ctx, request)
}

// KuaishouAppFetchLiveTopListRequest is the request for GET /api/v1/kuaishou/app/fetch_live_top_list.
type KuaishouAppFetchLiveTopListRequest = KuaishouAppKuaishouLiveTopListRequest

// KuaishouAppFetchLiveTopListResponse is the response for GET /api/v1/kuaishou/app/fetch_live_top_list.
type KuaishouAppFetchLiveTopListResponse = KuaishouAppKuaishouLiveTopListResponse

// FetchLiveTopList 快手直播榜单/Kuaishou live top list
//
// GET /api/v1/kuaishou/app/fetch_live_top_list
func (r KuaishouAppResource) FetchLiveTopList(ctx context.Context, request KuaishouAppFetchLiveTopListRequest) (*KuaishouAppFetchLiveTopListResponse, error) {
	return r.client.KuaishouAppKuaishouLiveTopList(ctx, request)
}

// KuaishouAppFetchHotBoardCategoriesResponse is the response for GET /api/v1/kuaishou/app/fetch_hot_board_categories.
type KuaishouAppFetchHotBoardCategoriesResponse = KuaishouAppKuaishouHotCategoriesResponse

// FetchHotBoardCategories 快手热榜分类/Kuaishou hot categories
//
// GET /api/v1/kuaishou/app/fetch_hot_board_categories
func (r KuaishouAppResource) FetchHotBoardCategories(ctx context.Context) (*KuaishouAppFetchHotBoardCategoriesResponse, error) {
	return r.client.KuaishouAppKuaishouHotCategories(ctx)
}

// KuaishouAppFetchHotBoardDetailRequest is the request for GET /api/v1/kuaishou/app/fetch_hot_board_detail.
type KuaishouAppFetchHotBoardDetailRequest = KuaishouAppKuaishouHotBoardDetailRequest

// KuaishouAppFetchHotBoardDetailResponse is the response for GET /api/v1/kuaishou/app/fetch_hot_board_detail.
type KuaishouAppFetchHotBoardDetailResponse = KuaishouAppKuaishouHotBoardDetailResponse

// FetchHotBoardDetail 快手热榜详情/Kuaishou hot board detail
//
// GET /api/v1/kuaishou/app/fetch_hot_board_detail
func (r KuaishouAppResource) FetchHotBoardDetail(ctx context.Context, request KuaishouAppFetchHotBoardDetailRequest) (*KuaishouAppFetchHotBoardDetailResponse, error) {
	return r.client.KuaishouAppKuaishouHotBoardDetail(ctx, request)
}

// KuaishouAppFetchHotSearchPersonResponse is the response for GET /api/v1/kuaishou/app/fetch_hot_search_person.
type KuaishouAppFetchHotSearchPersonResponse = KuaishouAppKuaishouHotSearchPersonBoardResponse

// FetchHotSearchPerson 快手热搜人物榜单/Kuaishou hot search person board
//
// GET /api/v1/kuaishou/app/fetch_hot_search_person
func (r KuaishouAppResource) FetchHotSearchPerson(ctx context.Context) (*KuaishouAppFetchHotSearchPersonResponse, error) {
	return r.client.KuaishouAppKuaishouHotSearchPersonBoard(ctx)
}

// KuaishouAppFetchShoppingTopListRequest is the request for GET /api/v1/kuaishou/app/fetch_shopping_top_list.
type KuaishouAppFetchShoppingTopListRequest = KuaishouAppKuaishouShoppingTopListRequest

// KuaishouAppFetchShoppingTopListResponse is the response for GET /api/v1/kuaishou/app/fetch_shopping_top_list.
type KuaishouAppFetchShoppingTopListResponse = KuaishouAppKuaishouShoppingTopListResponse

// FetchShoppingTopList 快手购物榜单/Kuaishou shopping top list
//
// GET /api/v1/kuaishou/app/fetch_shopping_top_list
func (r KuaishouAppResource) FetchShoppingTopList(ctx context.Context, request KuaishouAppFetchShoppingTopListRequest) (*KuaishouAppFetchShoppingTopListResponse, error) {
	return r.client.KuaishouAppKuaishouShoppingTopList(ctx, request)
}

// KuaishouAppFetchBrandTopListRequest is the request for GET /api/v1/kuaishou/app/fetch_brand_top_list.
type KuaishouAppFetchBrandTopListRequest = KuaishouAppKuaishouBrandTopListRequest

// KuaishouAppFetchBrandTopListResponse is the response for GET /api/v1/kuaishou/app/fetch_brand_top_list.
type KuaishouAppFetchBrandTopListResponse = KuaishouAppKuaishouBrandTopListResponse

// FetchBrandTopList 快手品牌榜单/Kuaishou brand top list
//
// GET /api/v1/kuaishou/app/fetch_brand_top_list
func (r KuaishouAppResource) FetchBrandTopList(ctx context.Context, request KuaishouAppFetchBrandTopListRequest) (*KuaishouAppFetchBrandTopListResponse, error) {
	return r.client.KuaishouAppKuaishouBrandTopList(ctx, request)
}

// KuaishouAppFetchMusicRankingRequest is the request for GET /api/v1/kuaishou/app/fetch_music_ranking.
type KuaishouAppFetchMusicRankingRequest = KuaishouAppMusicRankingRequest

// KuaishouAppFetchMusicRankingResponse is the response for GET /api/v1/kuaishou/app/fetch_music_ranking.
type KuaishouAppFetchMusicRankingResponse = KuaishouAppMusicRankingResponse

// FetchMusicRanking 音乐榜单/Music ranking
//
// GET /api/v1/kuaishou/app/fetch_music_ranking
func (r KuaishouAppResource) FetchMusicRanking(ctx context.Context, request KuaishouAppFetchMusicRankingRequest) (*KuaishouAppFetchMusicRankingResponse, error) {
	return r.client.KuaishouAppMusicRanking(ctx, request)
}

// ZhihuWebResource contains endpoints from the Zhihu-Web-API tag.
type ZhihuWebResource struct {
	client *Client
}

// ZhihuWebFetchColumnArticlesRequest is the request for GET /api/v1/zhihu/web/fetch_column_articles.
type ZhihuWebFetchColumnArticlesRequest = ZhihuWebGetZhihuColumnArticlesRequest

// ZhihuWebFetchColumnArticlesResponse is the response for GET /api/v1/zhihu/web/fetch_column_articles.
type ZhihuWebFetchColumnArticlesResponse = ZhihuWebGetZhihuColumnArticlesResponse

// FetchColumnArticles 获取知乎专栏文章列表/Get Zhihu Column Articles
//
// GET /api/v1/zhihu/web/fetch_column_articles
func (r ZhihuWebResource) FetchColumnArticles(ctx context.Context, request ZhihuWebFetchColumnArticlesRequest) (*ZhihuWebFetchColumnArticlesResponse, error) {
	return r.client.ZhihuWebGetZhihuColumnArticles(ctx, request)
}

// ZhihuWebFetchColumnArticleDetailRequest is the request for GET /api/v1/zhihu/web/fetch_column_article_detail.
type ZhihuWebFetchColumnArticleDetailRequest = ZhihuWebGetZhihuColumnArticleDetailRequest

// ZhihuWebFetchColumnArticleDetailResponse is the response for GET /api/v1/zhihu/web/fetch_column_article_detail.
type ZhihuWebFetchColumnArticleDetailResponse = ZhihuWebGetZhihuColumnArticleDetailResponse

// FetchColumnArticleDetail 获取知乎专栏文章详情/Get Zhihu Column Article Detail
//
// GET /api/v1/zhihu/web/fetch_column_article_detail
func (r ZhihuWebResource) FetchColumnArticleDetail(ctx context.Context, request ZhihuWebFetchColumnArticleDetailRequest) (*ZhihuWebFetchColumnArticleDetailResponse, error) {
	return r.client.ZhihuWebGetZhihuColumnArticleDetail(ctx, request)
}

// ZhihuWebFetchColumnRecommendRequest is the request for GET /api/v1/zhihu/web/fetch_column_recommend.
type ZhihuWebFetchColumnRecommendRequest = ZhihuWebGetZhihuSimilarColumnRecommendRequest

// ZhihuWebFetchColumnRecommendResponse is the response for GET /api/v1/zhihu/web/fetch_column_recommend.
type ZhihuWebFetchColumnRecommendResponse = ZhihuWebGetZhihuSimilarColumnRecommendResponse

// FetchColumnRecommend 获取知乎相似专栏推荐/Get Zhihu Similar Column Recommend
//
// GET /api/v1/zhihu/web/fetch_column_recommend
func (r ZhihuWebResource) FetchColumnRecommend(ctx context.Context, request ZhihuWebFetchColumnRecommendRequest) (*ZhihuWebFetchColumnRecommendResponse, error) {
	return r.client.ZhihuWebGetZhihuSimilarColumnRecommend(ctx, request)
}

// ZhihuWebFetchColumnRelationshipRequest is the request for GET /api/v1/zhihu/web/fetch_column_relationship.
type ZhihuWebFetchColumnRelationshipRequest = ZhihuWebGetZhihuColumnArticleRelationshipRequest

// ZhihuWebFetchColumnRelationshipResponse is the response for GET /api/v1/zhihu/web/fetch_column_relationship.
type ZhihuWebFetchColumnRelationshipResponse = ZhihuWebGetZhihuColumnArticleRelationshipResponse

// FetchColumnRelationship 获取知乎专栏文章互动关系/Get Zhihu Column Article Relationship
//
// GET /api/v1/zhihu/web/fetch_column_relationship
func (r ZhihuWebResource) FetchColumnRelationship(ctx context.Context, request ZhihuWebFetchColumnRelationshipRequest) (*ZhihuWebFetchColumnRelationshipResponse, error) {
	return r.client.ZhihuWebGetZhihuColumnArticleRelationship(ctx, request)
}

// ZhihuWebFetchColumnCommentConfigRequest is the request for GET /api/v1/zhihu/web/fetch_column_comment_config.
type ZhihuWebFetchColumnCommentConfigRequest = ZhihuWebGetZhihuColumnCommentConfigRequest

// ZhihuWebFetchColumnCommentConfigResponse is the response for GET /api/v1/zhihu/web/fetch_column_comment_config.
type ZhihuWebFetchColumnCommentConfigResponse = ZhihuWebGetZhihuColumnCommentConfigResponse

// FetchColumnCommentConfig 获取知乎专栏评论区配置/Get Zhihu Column Comment Config
//
// GET /api/v1/zhihu/web/fetch_column_comment_config
func (r ZhihuWebResource) FetchColumnCommentConfig(ctx context.Context, request ZhihuWebFetchColumnCommentConfigRequest) (*ZhihuWebFetchColumnCommentConfigResponse, error) {
	return r.client.ZhihuWebGetZhihuColumnCommentConfig(ctx, request)
}

// ZhihuWebFetchHotRecommendRequest is the request for GET /api/v1/zhihu/web/fetch_hot_recommend.
type ZhihuWebFetchHotRecommendRequest = ZhihuWebGetZhihuHotRecommendRequest

// ZhihuWebFetchHotRecommendResponse is the response for GET /api/v1/zhihu/web/fetch_hot_recommend.
type ZhihuWebFetchHotRecommendResponse = ZhihuWebGetZhihuHotRecommendResponse

// FetchHotRecommend 获取知乎首页推荐/Get Zhihu Hot Recommend
//
// GET /api/v1/zhihu/web/fetch_hot_recommend
func (r ZhihuWebResource) FetchHotRecommend(ctx context.Context, request ZhihuWebFetchHotRecommendRequest) (*ZhihuWebFetchHotRecommendResponse, error) {
	return r.client.ZhihuWebGetZhihuHotRecommend(ctx, request)
}

// ZhihuWebFetchHotListRequest is the request for GET /api/v1/zhihu/web/fetch_hot_list.
type ZhihuWebFetchHotListRequest = ZhihuWebGetZhihuHotListRequest

// ZhihuWebFetchHotListResponse is the response for GET /api/v1/zhihu/web/fetch_hot_list.
type ZhihuWebFetchHotListResponse = ZhihuWebGetZhihuHotListResponse

// FetchHotList 获取知乎首页热榜/Get Zhihu Hot List
//
// GET /api/v1/zhihu/web/fetch_hot_list
func (r ZhihuWebResource) FetchHotList(ctx context.Context, request ZhihuWebFetchHotListRequest) (*ZhihuWebFetchHotListResponse, error) {
	return r.client.ZhihuWebGetZhihuHotList(ctx, request)
}

// ZhihuWebFetchVideoListRequest is the request for GET /api/v1/zhihu/web/fetch_video_list.
type ZhihuWebFetchVideoListRequest = ZhihuWebGetZhihuVideoListRequest

// ZhihuWebFetchVideoListResponse is the response for GET /api/v1/zhihu/web/fetch_video_list.
type ZhihuWebFetchVideoListResponse = ZhihuWebGetZhihuVideoListResponse

// FetchVideoList 获取知乎首页视频榜/Get Zhihu Video List
//
// GET /api/v1/zhihu/web/fetch_video_list
func (r ZhihuWebResource) FetchVideoList(ctx context.Context, request ZhihuWebFetchVideoListRequest) (*ZhihuWebFetchVideoListResponse, error) {
	return r.client.ZhihuWebGetZhihuVideoList(ctx, request)
}

// ZhihuWebFetchArticleSearchV3Request is the request for GET /api/v1/zhihu/web/fetch_article_search_v3.
type ZhihuWebFetchArticleSearchV3Request = ZhihuWebGetZhihuArticleSearchV3Request

// ZhihuWebFetchArticleSearchV3Response is the response for GET /api/v1/zhihu/web/fetch_article_search_v3.
type ZhihuWebFetchArticleSearchV3Response = ZhihuWebGetZhihuArticleSearchV3Response

// FetchArticleSearchV3 获取知乎文章搜索V3/Get Zhihu Article Search V3
//
// GET /api/v1/zhihu/web/fetch_article_search_v3
func (r ZhihuWebResource) FetchArticleSearchV3(ctx context.Context, request ZhihuWebFetchArticleSearchV3Request) (*ZhihuWebFetchArticleSearchV3Response, error) {
	return r.client.ZhihuWebGetZhihuArticleSearchV3(ctx, request)
}

// ZhihuWebFetchUserSearchV3Request is the request for GET /api/v1/zhihu/web/fetch_user_search_v3.
type ZhihuWebFetchUserSearchV3Request = ZhihuWebGetZhihuUserSearchV3Request

// ZhihuWebFetchUserSearchV3Response is the response for GET /api/v1/zhihu/web/fetch_user_search_v3.
type ZhihuWebFetchUserSearchV3Response = ZhihuWebGetZhihuUserSearchV3Response

// FetchUserSearchV3 获取知乎用户搜索V3/Get Zhihu User Search V3
//
// GET /api/v1/zhihu/web/fetch_user_search_v3
func (r ZhihuWebResource) FetchUserSearchV3(ctx context.Context, request ZhihuWebFetchUserSearchV3Request) (*ZhihuWebFetchUserSearchV3Response, error) {
	return r.client.ZhihuWebGetZhihuUserSearchV3(ctx, request)
}

// ZhihuWebFetchTopicSearchV3Request is the request for GET /api/v1/zhihu/web/fetch_topic_search_v3.
type ZhihuWebFetchTopicSearchV3Request = ZhihuWebGetZhihuTopicSearchV3Request

// ZhihuWebFetchTopicSearchV3Response is the response for GET /api/v1/zhihu/web/fetch_topic_search_v3.
type ZhihuWebFetchTopicSearchV3Response = ZhihuWebGetZhihuTopicSearchV3Response

// FetchTopicSearchV3 获取知乎话题搜索V3/Get Zhihu Topic Search V3
//
// GET /api/v1/zhihu/web/fetch_topic_search_v3
func (r ZhihuWebResource) FetchTopicSearchV3(ctx context.Context, request ZhihuWebFetchTopicSearchV3Request) (*ZhihuWebFetchTopicSearchV3Response, error) {
	return r.client.ZhihuWebGetZhihuTopicSearchV3(ctx, request)
}

// ZhihuWebFetchScholarSearchV3Request is the request for POST /api/v1/zhihu/web/fetch_scholar_search_v3.
type ZhihuWebFetchScholarSearchV3Request = ZhihuWebGetZhihuScholarSearchV3Request

// ZhihuWebFetchScholarSearchV3Response is the response for POST /api/v1/zhihu/web/fetch_scholar_search_v3.
type ZhihuWebFetchScholarSearchV3Response = ZhihuWebGetZhihuScholarSearchV3Response

// FetchScholarSearchV3 获取知乎论文搜索V3/Get Zhihu Scholar Search V3
//
// POST /api/v1/zhihu/web/fetch_scholar_search_v3
func (r ZhihuWebResource) FetchScholarSearchV3(ctx context.Context, request ZhihuWebFetchScholarSearchV3Request) (*ZhihuWebFetchScholarSearchV3Response, error) {
	return r.client.ZhihuWebGetZhihuScholarSearchV3(ctx, request)
}

// ZhihuWebFetchAiSearchRequest is the request for GET /api/v1/zhihu/web/fetch_ai_search.
type ZhihuWebFetchAiSearchRequest = ZhihuWebGetZhihuAiSearchRequest

// ZhihuWebFetchAiSearchResponse is the response for GET /api/v1/zhihu/web/fetch_ai_search.
type ZhihuWebFetchAiSearchResponse = ZhihuWebGetZhihuAiSearchResponse

// FetchAiSearch 获取知乎AI搜索/Get Zhihu AI Search
//
// GET /api/v1/zhihu/web/fetch_ai_search
func (r ZhihuWebResource) FetchAiSearch(ctx context.Context, request ZhihuWebFetchAiSearchRequest) (*ZhihuWebFetchAiSearchResponse, error) {
	return r.client.ZhihuWebGetZhihuAiSearch(ctx, request)
}

// ZhihuWebFetchAiSearchResultRequest is the request for GET /api/v1/zhihu/web/fetch_ai_search_result.
type ZhihuWebFetchAiSearchResultRequest = ZhihuWebGetZhihuAiSearchResultRequest

// ZhihuWebFetchAiSearchResultResponse is the response for GET /api/v1/zhihu/web/fetch_ai_search_result.
type ZhihuWebFetchAiSearchResultResponse = ZhihuWebGetZhihuAiSearchResultResponse

// FetchAiSearchResult 获取知乎AI搜索结果/Get Zhihu AI Search Result
//
// GET /api/v1/zhihu/web/fetch_ai_search_result
func (r ZhihuWebResource) FetchAiSearchResult(ctx context.Context, request ZhihuWebFetchAiSearchResultRequest) (*ZhihuWebFetchAiSearchResultResponse, error) {
	return r.client.ZhihuWebGetZhihuAiSearchResult(ctx, request)
}

// ZhihuWebFetchVideoSearchV3Request is the request for GET /api/v1/zhihu/web/fetch_video_search_v3.
type ZhihuWebFetchVideoSearchV3Request = ZhihuWebGetZhihuVideoSearchV3Request

// ZhihuWebFetchVideoSearchV3Response is the response for GET /api/v1/zhihu/web/fetch_video_search_v3.
type ZhihuWebFetchVideoSearchV3Response = ZhihuWebGetZhihuVideoSearchV3Response

// FetchVideoSearchV3 获取知乎视频搜索V3/Get Zhihu Video Search V3
//
// GET /api/v1/zhihu/web/fetch_video_search_v3
func (r ZhihuWebResource) FetchVideoSearchV3(ctx context.Context, request ZhihuWebFetchVideoSearchV3Request) (*ZhihuWebFetchVideoSearchV3Response, error) {
	return r.client.ZhihuWebGetZhihuVideoSearchV3(ctx, request)
}

// ZhihuWebFetchColumnSearchV3Request is the request for GET /api/v1/zhihu/web/fetch_column_search_v3.
type ZhihuWebFetchColumnSearchV3Request = ZhihuWebGetZhihuColumnSearchV3Request

// ZhihuWebFetchColumnSearchV3Response is the response for GET /api/v1/zhihu/web/fetch_column_search_v3.
type ZhihuWebFetchColumnSearchV3Response = ZhihuWebGetZhihuColumnSearchV3Response

// FetchColumnSearchV3 获取知乎专栏搜索V3/Get Zhihu Column Search V3
//
// GET /api/v1/zhihu/web/fetch_column_search_v3
func (r ZhihuWebResource) FetchColumnSearchV3(ctx context.Context, request ZhihuWebFetchColumnSearchV3Request) (*ZhihuWebFetchColumnSearchV3Response, error) {
	return r.client.ZhihuWebGetZhihuColumnSearchV3(ctx, request)
}

// ZhihuWebFetchSaltSearchV3Request is the request for GET /api/v1/zhihu/web/fetch_salt_search_v3.
type ZhihuWebFetchSaltSearchV3Request = ZhihuWebGetZhihuSaltSearchV3Request

// ZhihuWebFetchSaltSearchV3Response is the response for GET /api/v1/zhihu/web/fetch_salt_search_v3.
type ZhihuWebFetchSaltSearchV3Response = ZhihuWebGetZhihuSaltSearchV3Response

// FetchSaltSearchV3 获取知乎盐选内容搜索V3/Get Zhihu Salt Search V3
//
// GET /api/v1/zhihu/web/fetch_salt_search_v3
func (r ZhihuWebResource) FetchSaltSearchV3(ctx context.Context, request ZhihuWebFetchSaltSearchV3Request) (*ZhihuWebFetchSaltSearchV3Response, error) {
	return r.client.ZhihuWebGetZhihuSaltSearchV3(ctx, request)
}

// ZhihuWebFetchEbookSearchV3Request is the request for GET /api/v1/zhihu/web/fetch_ebook_search_v3.
type ZhihuWebFetchEbookSearchV3Request = ZhihuWebGetZhihuEbookSearchV3Request

// ZhihuWebFetchEbookSearchV3Response is the response for GET /api/v1/zhihu/web/fetch_ebook_search_v3.
type ZhihuWebFetchEbookSearchV3Response = ZhihuWebGetZhihuEbookSearchV3Response

// FetchEbookSearchV3 获取知乎电子书搜索V3/Get Zhihu Ebook Search V3
//
// GET /api/v1/zhihu/web/fetch_ebook_search_v3
func (r ZhihuWebResource) FetchEbookSearchV3(ctx context.Context, request ZhihuWebFetchEbookSearchV3Request) (*ZhihuWebFetchEbookSearchV3Response, error) {
	return r.client.ZhihuWebGetZhihuEbookSearchV3(ctx, request)
}

// ZhihuWebFetchPresetSearchResponse is the response for GET /api/v1/zhihu/web/fetch_preset_search.
type ZhihuWebFetchPresetSearchResponse = ZhihuWebGetZhihuPresetSearchResponse

// FetchPresetSearch 获取知乎搜索预设词/Get Zhihu Preset Search
//
// GET /api/v1/zhihu/web/fetch_preset_search
func (r ZhihuWebResource) FetchPresetSearch(ctx context.Context) (*ZhihuWebFetchPresetSearchResponse, error) {
	return r.client.ZhihuWebGetZhihuPresetSearch(ctx)
}

// ZhihuWebFetchSearchRecommendResponse is the response for GET /api/v1/zhihu/web/fetch_search_recommend.
type ZhihuWebFetchSearchRecommendResponse = ZhihuWebGetZhihuSearchRecommendResponse

// FetchSearchRecommend 获取知乎搜索发现/Get Zhihu Search Recommend
//
// GET /api/v1/zhihu/web/fetch_search_recommend
func (r ZhihuWebResource) FetchSearchRecommend(ctx context.Context) (*ZhihuWebFetchSearchRecommendResponse, error) {
	return r.client.ZhihuWebGetZhihuSearchRecommend(ctx)
}

// ZhihuWebFetchSearchSuggestRequest is the request for GET /api/v1/zhihu/web/fetch_search_suggest.
type ZhihuWebFetchSearchSuggestRequest = ZhihuWebGetZhihuSearchSuggestRequest

// ZhihuWebFetchSearchSuggestResponse is the response for GET /api/v1/zhihu/web/fetch_search_suggest.
type ZhihuWebFetchSearchSuggestResponse = ZhihuWebGetZhihuSearchSuggestResponse

// FetchSearchSuggest 知乎搜索预测词/Get Zhihu Search Suggest
//
// GET /api/v1/zhihu/web/fetch_search_suggest
func (r ZhihuWebResource) FetchSearchSuggest(ctx context.Context, request ZhihuWebFetchSearchSuggestRequest) (*ZhihuWebFetchSearchSuggestResponse, error) {
	return r.client.ZhihuWebGetZhihuSearchSuggest(ctx, request)
}

// ZhihuWebFetchCommentV5Request is the request for GET /api/v1/zhihu/web/fetch_comment_v5.
type ZhihuWebFetchCommentV5Request = ZhihuWebGetZhihuCommentV5Request

// ZhihuWebFetchCommentV5Response is the response for GET /api/v1/zhihu/web/fetch_comment_v5.
type ZhihuWebFetchCommentV5Response = ZhihuWebGetZhihuCommentV5Response

// FetchCommentV5 获取知乎评论区V5/Get Zhihu Comment V5
//
// GET /api/v1/zhihu/web/fetch_comment_v5
func (r ZhihuWebResource) FetchCommentV5(ctx context.Context, request ZhihuWebFetchCommentV5Request) (*ZhihuWebFetchCommentV5Response, error) {
	return r.client.ZhihuWebGetZhihuCommentV5(ctx, request)
}

// ZhihuWebFetchSubCommentV5Request is the request for GET /api/v1/zhihu/web/fetch_sub_comment_v5.
type ZhihuWebFetchSubCommentV5Request = ZhihuWebGetZhihuSubCommentV5Request

// ZhihuWebFetchSubCommentV5Response is the response for GET /api/v1/zhihu/web/fetch_sub_comment_v5.
type ZhihuWebFetchSubCommentV5Response = ZhihuWebGetZhihuSubCommentV5Response

// FetchSubCommentV5 获取知乎子评论区V5/Get Zhihu Sub Comment V5
//
// GET /api/v1/zhihu/web/fetch_sub_comment_v5
func (r ZhihuWebResource) FetchSubCommentV5(ctx context.Context, request ZhihuWebFetchSubCommentV5Request) (*ZhihuWebFetchSubCommentV5Response, error) {
	return r.client.ZhihuWebGetZhihuSubCommentV5(ctx, request)
}

// ZhihuWebFetchUserInfoRequest is the request for GET /api/v1/zhihu/web/fetch_user_info.
type ZhihuWebFetchUserInfoRequest = ZhihuWebGetZhihuUserInfoRequest

// ZhihuWebFetchUserInfoResponse is the response for GET /api/v1/zhihu/web/fetch_user_info.
type ZhihuWebFetchUserInfoResponse = ZhihuWebGetZhihuUserInfoResponse

// FetchUserInfo 获取知乎用户信息/Get Zhihu User Info
//
// GET /api/v1/zhihu/web/fetch_user_info
func (r ZhihuWebResource) FetchUserInfo(ctx context.Context, request ZhihuWebFetchUserInfoRequest) (*ZhihuWebFetchUserInfoResponse, error) {
	return r.client.ZhihuWebGetZhihuUserInfo(ctx, request)
}

// ZhihuWebFetchUserFolloweesRequest is the request for GET /api/v1/zhihu/web/fetch_user_followees.
type ZhihuWebFetchUserFolloweesRequest = ZhihuWebGetZhihuUserFollowingRequest

// ZhihuWebFetchUserFolloweesResponse is the response for GET /api/v1/zhihu/web/fetch_user_followees.
type ZhihuWebFetchUserFolloweesResponse = ZhihuWebGetZhihuUserFollowingResponse

// FetchUserFollowees 获取知乎用户关注列表/Get Zhihu User Following
//
// GET /api/v1/zhihu/web/fetch_user_followees
func (r ZhihuWebResource) FetchUserFollowees(ctx context.Context, request ZhihuWebFetchUserFolloweesRequest) (*ZhihuWebFetchUserFolloweesResponse, error) {
	return r.client.ZhihuWebGetZhihuUserFollowing(ctx, request)
}

// ZhihuWebFetchUserFollowersRequest is the request for GET /api/v1/zhihu/web/fetch_user_followers.
type ZhihuWebFetchUserFollowersRequest = ZhihuWebGetZhihuUserFollowersRequest

// ZhihuWebFetchUserFollowersResponse is the response for GET /api/v1/zhihu/web/fetch_user_followers.
type ZhihuWebFetchUserFollowersResponse = ZhihuWebGetZhihuUserFollowersResponse

// FetchUserFollowers 获取知乎用户粉丝列表/Get Zhihu User Followers
//
// GET /api/v1/zhihu/web/fetch_user_followers
func (r ZhihuWebResource) FetchUserFollowers(ctx context.Context, request ZhihuWebFetchUserFollowersRequest) (*ZhihuWebFetchUserFollowersResponse, error) {
	return r.client.ZhihuWebGetZhihuUserFollowers(ctx, request)
}

// ZhihuWebFetchUserArticlesRequest is the request for GET /api/v1/zhihu/web/fetch_user_articles.
type ZhihuWebFetchUserArticlesRequest = ZhihuWebGetZhihuUserArticlesRequest

// ZhihuWebFetchUserArticlesResponse is the response for GET /api/v1/zhihu/web/fetch_user_articles.
type ZhihuWebFetchUserArticlesResponse = ZhihuWebGetZhihuUserArticlesResponse

// FetchUserArticles 获取知乎用户的文章列表/Get Zhihu User Articles
//
// GET /api/v1/zhihu/web/fetch_user_articles
func (r ZhihuWebResource) FetchUserArticles(ctx context.Context, request ZhihuWebFetchUserArticlesRequest) (*ZhihuWebFetchUserArticlesResponse, error) {
	return r.client.ZhihuWebGetZhihuUserArticles(ctx, request)
}

// ZhihuWebFetchUserIncludedArticlesRequest is the request for GET /api/v1/zhihu/web/fetch_user_included_articles.
type ZhihuWebFetchUserIncludedArticlesRequest = ZhihuWebGetZhihuUserIncludedArticlesRequest

// ZhihuWebFetchUserIncludedArticlesResponse is the response for GET /api/v1/zhihu/web/fetch_user_included_articles.
type ZhihuWebFetchUserIncludedArticlesResponse = ZhihuWebGetZhihuUserIncludedArticlesResponse

// FetchUserIncludedArticles 获取知乎用户的被收录文章列表/Get Zhihu User Included Articles
//
// GET /api/v1/zhihu/web/fetch_user_included_articles
func (r ZhihuWebResource) FetchUserIncludedArticles(ctx context.Context, request ZhihuWebFetchUserIncludedArticlesRequest) (*ZhihuWebFetchUserIncludedArticlesResponse, error) {
	return r.client.ZhihuWebGetZhihuUserIncludedArticles(ctx, request)
}

// ZhihuWebFetchUserFollowColumnsRequest is the request for GET /api/v1/zhihu/web/fetch_user_follow_columns.
type ZhihuWebFetchUserFollowColumnsRequest = ZhihuWebGetZhihuUserColumnsRequest

// ZhihuWebFetchUserFollowColumnsResponse is the response for GET /api/v1/zhihu/web/fetch_user_follow_columns.
type ZhihuWebFetchUserFollowColumnsResponse = ZhihuWebGetZhihuUserColumnsResponse

// FetchUserFollowColumns 获取知乎用户订阅的专栏/Get Zhihu User Columns
//
// GET /api/v1/zhihu/web/fetch_user_follow_columns
func (r ZhihuWebResource) FetchUserFollowColumns(ctx context.Context, request ZhihuWebFetchUserFollowColumnsRequest) (*ZhihuWebFetchUserFollowColumnsResponse, error) {
	return r.client.ZhihuWebGetZhihuUserColumns(ctx, request)
}

// ZhihuWebFetchUserFollowQuestionsRequest is the request for GET /api/v1/zhihu/web/fetch_user_follow_questions.
type ZhihuWebFetchUserFollowQuestionsRequest = ZhihuWebGetZhihuUserFollowQuestionsRequest

// ZhihuWebFetchUserFollowQuestionsResponse is the response for GET /api/v1/zhihu/web/fetch_user_follow_questions.
type ZhihuWebFetchUserFollowQuestionsResponse = ZhihuWebGetZhihuUserFollowQuestionsResponse

// FetchUserFollowQuestions 获取知乎用户关注的问题/Get Zhihu User Follow Questions
//
// GET /api/v1/zhihu/web/fetch_user_follow_questions
func (r ZhihuWebResource) FetchUserFollowQuestions(ctx context.Context, request ZhihuWebFetchUserFollowQuestionsRequest) (*ZhihuWebFetchUserFollowQuestionsResponse, error) {
	return r.client.ZhihuWebGetZhihuUserFollowQuestions(ctx, request)
}

// ZhihuWebFetchUserFollowCollectionsRequest is the request for GET /api/v1/zhihu/web/fetch_user_follow_collections.
type ZhihuWebFetchUserFollowCollectionsRequest = ZhihuWebGetZhihuUserFollowCollectionsRequest

// ZhihuWebFetchUserFollowCollectionsResponse is the response for GET /api/v1/zhihu/web/fetch_user_follow_collections.
type ZhihuWebFetchUserFollowCollectionsResponse = ZhihuWebGetZhihuUserFollowCollectionsResponse

// FetchUserFollowCollections 获取知乎用户关注的收藏/Get Zhihu User Follow Collections
//
// GET /api/v1/zhihu/web/fetch_user_follow_collections
func (r ZhihuWebResource) FetchUserFollowCollections(ctx context.Context, request ZhihuWebFetchUserFollowCollectionsRequest) (*ZhihuWebFetchUserFollowCollectionsResponse, error) {
	return r.client.ZhihuWebGetZhihuUserFollowCollections(ctx, request)
}

// ZhihuWebFetchUserFollowTopicsRequest is the request for GET /api/v1/zhihu/web/fetch_user_follow_topics.
type ZhihuWebFetchUserFollowTopicsRequest = ZhihuWebGetZhihuUserFollowTopicsRequest

// ZhihuWebFetchUserFollowTopicsResponse is the response for GET /api/v1/zhihu/web/fetch_user_follow_topics.
type ZhihuWebFetchUserFollowTopicsResponse = ZhihuWebGetZhihuUserFollowTopicsResponse

// FetchUserFollowTopics 获取知乎用户关注的话题/Get Zhihu User Follow Topics
//
// GET /api/v1/zhihu/web/fetch_user_follow_topics
func (r ZhihuWebResource) FetchUserFollowTopics(ctx context.Context, request ZhihuWebFetchUserFollowTopicsRequest) (*ZhihuWebFetchUserFollowTopicsResponse, error) {
	return r.client.ZhihuWebGetZhihuUserFollowTopics(ctx, request)
}

// ZhihuWebFetchRecommendFolloweesResponse is the response for GET /api/v1/zhihu/web/fetch_recommend_followees.
type ZhihuWebFetchRecommendFolloweesResponse = ZhihuWebGetZhihuRecommendFolloweesResponse

// FetchRecommendFollowees 获取知乎推荐关注列表/Get Zhihu Recommend Followees
//
// GET /api/v1/zhihu/web/fetch_recommend_followees
func (r ZhihuWebResource) FetchRecommendFollowees(ctx context.Context) (*ZhihuWebFetchRecommendFolloweesResponse, error) {
	return r.client.ZhihuWebGetZhihuRecommendFollowees(ctx)
}

// ZhihuWebFetchQuestionAnswersRequest is the request for GET /api/v1/zhihu/web/fetch_question_answers.
type ZhihuWebFetchQuestionAnswersRequest = ZhihuWebGetZhihuQuestionAnswersRequest

// ZhihuWebFetchQuestionAnswersResponse is the response for GET /api/v1/zhihu/web/fetch_question_answers.
type ZhihuWebFetchQuestionAnswersResponse = ZhihuWebGetZhihuQuestionAnswersResponse

// FetchQuestionAnswers 获取知乎问题回答列表/Get Zhihu Question Answers
//
// GET /api/v1/zhihu/web/fetch_question_answers
func (r ZhihuWebResource) FetchQuestionAnswers(ctx context.Context, request ZhihuWebFetchQuestionAnswersRequest) (*ZhihuWebFetchQuestionAnswersResponse, error) {
	return r.client.ZhihuWebGetZhihuQuestionAnswers(ctx, request)
}

// PiPiXiaAppResource contains endpoints from the PiPiXia-App-API tag.
type PiPiXiaAppResource struct {
	client *Client
}

// PiPiXiaAppFetchPostDetailRequest is the request for GET /api/v1/pipixia/app/fetch_post_detail.
type PiPiXiaAppFetchPostDetailRequest = PiPiXiaAppGetSingleVideoDataRequest

// PiPiXiaAppFetchPostDetailResponse is the response for GET /api/v1/pipixia/app/fetch_post_detail.
type PiPiXiaAppFetchPostDetailResponse = PiPiXiaAppGetSingleVideoDataResponse

// FetchPostDetail 获取单个作品数据/Get single video data
//
// GET /api/v1/pipixia/app/fetch_post_detail
func (r PiPiXiaAppResource) FetchPostDetail(ctx context.Context, request PiPiXiaAppFetchPostDetailRequest) (*PiPiXiaAppFetchPostDetailResponse, error) {
	return r.client.PiPiXiaAppGetSingleVideoData(ctx, request)
}

// PiPiXiaAppFetchIncreasePostViewCountRequest is the request for GET /api/v1/pipixia/app/fetch_increase_post_view_count.
type PiPiXiaAppFetchIncreasePostViewCountRequest = PiPiXiaAppIncreasePostViewCountRequest

// PiPiXiaAppFetchIncreasePostViewCountResponse is the response for GET /api/v1/pipixia/app/fetch_increase_post_view_count.
type PiPiXiaAppFetchIncreasePostViewCountResponse = PiPiXiaAppIncreasePostViewCountResponse

// FetchIncreasePostViewCount 增加作品浏览数/Increase post view count
//
// GET /api/v1/pipixia/app/fetch_increase_post_view_count
func (r PiPiXiaAppResource) FetchIncreasePostViewCount(ctx context.Context, request PiPiXiaAppFetchIncreasePostViewCountRequest) (*PiPiXiaAppFetchIncreasePostViewCountResponse, error) {
	return r.client.PiPiXiaAppIncreasePostViewCount(ctx, request)
}

// PiPiXiaAppFetchPostStatisticsRequest is the request for GET /api/v1/pipixia/app/fetch_post_statistics.
type PiPiXiaAppFetchPostStatisticsRequest = PiPiXiaAppGetPostStatisticsRequest

// PiPiXiaAppFetchPostStatisticsResponse is the response for GET /api/v1/pipixia/app/fetch_post_statistics.
type PiPiXiaAppFetchPostStatisticsResponse = PiPiXiaAppGetPostStatisticsResponse

// FetchPostStatistics 获取作品统计数据/Get post statistics
//
// GET /api/v1/pipixia/app/fetch_post_statistics
func (r PiPiXiaAppResource) FetchPostStatistics(ctx context.Context, request PiPiXiaAppFetchPostStatisticsRequest) (*PiPiXiaAppFetchPostStatisticsResponse, error) {
	return r.client.PiPiXiaAppGetPostStatistics(ctx, request)
}

// PiPiXiaAppFetchUserInfoRequest is the request for GET /api/v1/pipixia/app/fetch_user_info.
type PiPiXiaAppFetchUserInfoRequest = PiPiXiaAppGetUserInformationRequest

// PiPiXiaAppFetchUserInfoResponse is the response for GET /api/v1/pipixia/app/fetch_user_info.
type PiPiXiaAppFetchUserInfoResponse = PiPiXiaAppGetUserInformationResponse

// FetchUserInfo 获取用户信息/Get user information
//
// GET /api/v1/pipixia/app/fetch_user_info
func (r PiPiXiaAppResource) FetchUserInfo(ctx context.Context, request PiPiXiaAppFetchUserInfoRequest) (*PiPiXiaAppFetchUserInfoResponse, error) {
	return r.client.PiPiXiaAppGetUserInformation(ctx, request)
}

// PiPiXiaAppFetchUserPostListRequest is the request for GET /api/v1/pipixia/app/fetch_user_post_list.
type PiPiXiaAppFetchUserPostListRequest = PiPiXiaAppGetUserPostListRequest

// PiPiXiaAppFetchUserPostListResponse is the response for GET /api/v1/pipixia/app/fetch_user_post_list.
type PiPiXiaAppFetchUserPostListResponse = PiPiXiaAppGetUserPostListResponse

// FetchUserPostList 获取用户作品列表/Get user post list
//
// GET /api/v1/pipixia/app/fetch_user_post_list
func (r PiPiXiaAppResource) FetchUserPostList(ctx context.Context, request PiPiXiaAppFetchUserPostListRequest) (*PiPiXiaAppFetchUserPostListResponse, error) {
	return r.client.PiPiXiaAppGetUserPostList(ctx, request)
}

// PiPiXiaAppFetchUserFollowerListRequest is the request for GET /api/v1/pipixia/app/fetch_user_follower_list.
type PiPiXiaAppFetchUserFollowerListRequest = PiPiXiaAppGetUserFollowerListRequest

// PiPiXiaAppFetchUserFollowerListResponse is the response for GET /api/v1/pipixia/app/fetch_user_follower_list.
type PiPiXiaAppFetchUserFollowerListResponse = PiPiXiaAppGetUserFollowerListResponse

// FetchUserFollowerList 获取用户粉丝列表/Get user follower list
//
// GET /api/v1/pipixia/app/fetch_user_follower_list
func (r PiPiXiaAppResource) FetchUserFollowerList(ctx context.Context, request PiPiXiaAppFetchUserFollowerListRequest) (*PiPiXiaAppFetchUserFollowerListResponse, error) {
	return r.client.PiPiXiaAppGetUserFollowerList(ctx, request)
}

// PiPiXiaAppFetchUserFollowingListRequest is the request for GET /api/v1/pipixia/app/fetch_user_following_list.
type PiPiXiaAppFetchUserFollowingListRequest = PiPiXiaAppGetUserFollowingListRequest

// PiPiXiaAppFetchUserFollowingListResponse is the response for GET /api/v1/pipixia/app/fetch_user_following_list.
type PiPiXiaAppFetchUserFollowingListResponse = PiPiXiaAppGetUserFollowingListResponse

// FetchUserFollowingList 获取用户关注列表/Get user following list
//
// GET /api/v1/pipixia/app/fetch_user_following_list
func (r PiPiXiaAppResource) FetchUserFollowingList(ctx context.Context, request PiPiXiaAppFetchUserFollowingListRequest) (*PiPiXiaAppFetchUserFollowingListResponse, error) {
	return r.client.PiPiXiaAppGetUserFollowingList(ctx, request)
}

// PiPiXiaAppFetchPostCommentListRequest is the request for GET /api/v1/pipixia/app/fetch_post_comment_list.
type PiPiXiaAppFetchPostCommentListRequest = PiPiXiaAppGetPostCommentListRequest

// PiPiXiaAppFetchPostCommentListResponse is the response for GET /api/v1/pipixia/app/fetch_post_comment_list.
type PiPiXiaAppFetchPostCommentListResponse = PiPiXiaAppGetPostCommentListResponse

// FetchPostCommentList 获取作品评论列表/Get post comment list
//
// GET /api/v1/pipixia/app/fetch_post_comment_list
func (r PiPiXiaAppResource) FetchPostCommentList(ctx context.Context, request PiPiXiaAppFetchPostCommentListRequest) (*PiPiXiaAppFetchPostCommentListResponse, error) {
	return r.client.PiPiXiaAppGetPostCommentList(ctx, request)
}

// PiPiXiaAppFetchShortURLRequest is the request for GET /api/v1/pipixia/app/fetch_short_url.
type PiPiXiaAppFetchShortURLRequest = PiPiXiaAppGenerateShortURLRequest

// PiPiXiaAppFetchShortURLResponse is the response for GET /api/v1/pipixia/app/fetch_short_url.
type PiPiXiaAppFetchShortURLResponse = PiPiXiaAppGenerateShortURLResponse

// FetchShortURL 生成短连接/Generate short URL
//
// GET /api/v1/pipixia/app/fetch_short_url
func (r PiPiXiaAppResource) FetchShortURL(ctx context.Context, request PiPiXiaAppFetchShortURLRequest) (*PiPiXiaAppFetchShortURLResponse, error) {
	return r.client.PiPiXiaAppGenerateShortURL(ctx, request)
}

// PiPiXiaAppFetchHomeFeedRequest is the request for GET /api/v1/pipixia/app/fetch_home_feed.
type PiPiXiaAppFetchHomeFeedRequest = PiPiXiaAppGetHomeFeedRequest

// PiPiXiaAppFetchHomeFeedResponse is the response for GET /api/v1/pipixia/app/fetch_home_feed.
type PiPiXiaAppFetchHomeFeedResponse = PiPiXiaAppGetHomeFeedResponse

// FetchHomeFeed 获取首页推荐/Get home feed
//
// GET /api/v1/pipixia/app/fetch_home_feed
func (r PiPiXiaAppResource) FetchHomeFeed(ctx context.Context, request PiPiXiaAppFetchHomeFeedRequest) (*PiPiXiaAppFetchHomeFeedResponse, error) {
	return r.client.PiPiXiaAppGetHomeFeed(ctx, request)
}

// PiPiXiaAppFetchHotSearchWordsResponse is the response for GET /api/v1/pipixia/app/fetch_hot_search_words.
type PiPiXiaAppFetchHotSearchWordsResponse = PiPiXiaAppGetHotSearchWordsResponse

// FetchHotSearchWords 获取热搜词条/Get hot search words
//
// GET /api/v1/pipixia/app/fetch_hot_search_words
func (r PiPiXiaAppResource) FetchHotSearchWords(ctx context.Context) (*PiPiXiaAppFetchHotSearchWordsResponse, error) {
	return r.client.PiPiXiaAppGetHotSearchWords(ctx)
}

// PiPiXiaAppFetchHotSearchBoardListResponse is the response for GET /api/v1/pipixia/app/fetch_hot_search_board_list.
type PiPiXiaAppFetchHotSearchBoardListResponse = PiPiXiaAppGetHotSearchBoardListResponse

// FetchHotSearchBoardList 获取热搜榜单列表/Get hot search board list
//
// GET /api/v1/pipixia/app/fetch_hot_search_board_list
func (r PiPiXiaAppResource) FetchHotSearchBoardList(ctx context.Context) (*PiPiXiaAppFetchHotSearchBoardListResponse, error) {
	return r.client.PiPiXiaAppGetHotSearchBoardList(ctx)
}

// PiPiXiaAppFetchHotSearchBoardDetailRequest is the request for GET /api/v1/pipixia/app/fetch_hot_search_board_detail.
type PiPiXiaAppFetchHotSearchBoardDetailRequest = PiPiXiaAppGetHotSearchBoardDetailRequest

// PiPiXiaAppFetchHotSearchBoardDetailResponse is the response for GET /api/v1/pipixia/app/fetch_hot_search_board_detail.
type PiPiXiaAppFetchHotSearchBoardDetailResponse = PiPiXiaAppGetHotSearchBoardDetailResponse

// FetchHotSearchBoardDetail 获取热搜榜单详情/Get hot search board detail
//
// GET /api/v1/pipixia/app/fetch_hot_search_board_detail
func (r PiPiXiaAppResource) FetchHotSearchBoardDetail(ctx context.Context, request PiPiXiaAppFetchHotSearchBoardDetailRequest) (*PiPiXiaAppFetchHotSearchBoardDetailResponse, error) {
	return r.client.PiPiXiaAppGetHotSearchBoardDetail(ctx, request)
}

// PiPiXiaAppFetchSearchRequest is the request for GET /api/v1/pipixia/app/fetch_search.
type PiPiXiaAppFetchSearchRequest = PiPiXiaAppSearchAPIRequest

// PiPiXiaAppFetchSearchResponse is the response for GET /api/v1/pipixia/app/fetch_search.
type PiPiXiaAppFetchSearchResponse = PiPiXiaAppSearchAPIResponse

// FetchSearch 搜索接口/Search API
//
// GET /api/v1/pipixia/app/fetch_search
func (r PiPiXiaAppResource) FetchSearch(ctx context.Context, request PiPiXiaAppFetchSearchRequest) (*PiPiXiaAppFetchSearchResponse, error) {
	return r.client.PiPiXiaAppSearchAPI(ctx, request)
}

// PiPiXiaAppFetchHashtagDetailRequest is the request for GET /api/v1/pipixia/app/fetch_hashtag_detail.
type PiPiXiaAppFetchHashtagDetailRequest = PiPiXiaAppGetHashtagDetailRequest

// PiPiXiaAppFetchHashtagDetailResponse is the response for GET /api/v1/pipixia/app/fetch_hashtag_detail.
type PiPiXiaAppFetchHashtagDetailResponse = PiPiXiaAppGetHashtagDetailResponse

// FetchHashtagDetail 获取话题详情/Get hashtag detail
//
// GET /api/v1/pipixia/app/fetch_hashtag_detail
func (r PiPiXiaAppResource) FetchHashtagDetail(ctx context.Context, request PiPiXiaAppFetchHashtagDetailRequest) (*PiPiXiaAppFetchHashtagDetailResponse, error) {
	return r.client.PiPiXiaAppGetHashtagDetail(ctx, request)
}

// PiPiXiaAppFetchHashtagPostListRequest is the request for GET /api/v1/pipixia/app/fetch_hashtag_post_list.
type PiPiXiaAppFetchHashtagPostListRequest = PiPiXiaAppGetHashtagPostListRequest

// PiPiXiaAppFetchHashtagPostListResponse is the response for GET /api/v1/pipixia/app/fetch_hashtag_post_list.
type PiPiXiaAppFetchHashtagPostListResponse = PiPiXiaAppGetHashtagPostListResponse

// FetchHashtagPostList 获取话题作品列表/Get hashtag post list
//
// GET /api/v1/pipixia/app/fetch_hashtag_post_list
func (r PiPiXiaAppResource) FetchHashtagPostList(ctx context.Context, request PiPiXiaAppFetchHashtagPostListRequest) (*PiPiXiaAppFetchHashtagPostListResponse, error) {
	return r.client.PiPiXiaAppGetHashtagPostList(ctx, request)
}

// PiPiXiaAppFetchHomeShortDramaFeedRequest is the request for GET /api/v1/pipixia/app/fetch_home_short_drama_feed.
type PiPiXiaAppFetchHomeShortDramaFeedRequest = PiPiXiaAppGetHomeShortDramaFeedRequest

// PiPiXiaAppFetchHomeShortDramaFeedResponse is the response for GET /api/v1/pipixia/app/fetch_home_short_drama_feed.
type PiPiXiaAppFetchHomeShortDramaFeedResponse = PiPiXiaAppGetHomeShortDramaFeedResponse

// FetchHomeShortDramaFeed 获取首页短剧推荐/Get home short drama feed
//
// GET /api/v1/pipixia/app/fetch_home_short_drama_feed
func (r PiPiXiaAppResource) FetchHomeShortDramaFeed(ctx context.Context, request PiPiXiaAppFetchHomeShortDramaFeedRequest) (*PiPiXiaAppFetchHomeShortDramaFeedResponse, error) {
	return r.client.PiPiXiaAppGetHomeShortDramaFeed(ctx, request)
}

// WeiboWebResource contains endpoints from the Weibo-Web-API tag.
type WeiboWebResource struct {
	client *Client
}

// WeiboWebFetchConfigListResponse is the response for GET /api/v1/weibo/web/fetch_config_list.
type WeiboWebFetchConfigListResponse = WeiboWebGetChannelConfigListResponse

// FetchConfigList 获取频道配置列表/Get channel config list
//
// GET /api/v1/weibo/web/fetch_config_list
func (r WeiboWebResource) FetchConfigList(ctx context.Context) (*WeiboWebFetchConfigListResponse, error) {
	return r.client.WeiboWebGetChannelConfigList(ctx)
}

// WeiboWebFetchTrendTopRequest is the request for GET /api/v1/weibo/web/fetch_trend_top.
type WeiboWebFetchTrendTopRequest = WeiboWebGetChannelTrendTopRequest

// WeiboWebFetchTrendTopResponse is the response for GET /api/v1/weibo/web/fetch_trend_top.
type WeiboWebFetchTrendTopResponse = WeiboWebGetChannelTrendTopResponse

// FetchTrendTop 获取频道热门趋势/Get channel trend top
//
// GET /api/v1/weibo/web/fetch_trend_top
func (r WeiboWebResource) FetchTrendTop(ctx context.Context, request WeiboWebFetchTrendTopRequest) (*WeiboWebFetchTrendTopResponse, error) {
	return r.client.WeiboWebGetChannelTrendTop(ctx, request)
}

// WeiboWebFetchChannelFeedRequest is the request for GET /api/v1/weibo/web/fetch_channel_feed.
type WeiboWebFetchChannelFeedRequest = WeiboWebGetChannelFeedByNameRequest

// WeiboWebFetchChannelFeedResponse is the response for GET /api/v1/weibo/web/fetch_channel_feed.
type WeiboWebFetchChannelFeedResponse = WeiboWebGetChannelFeedByNameResponse

// FetchChannelFeed 根据频道名称获取热门内容/Get channel feed by name
//
// GET /api/v1/weibo/web/fetch_channel_feed
func (r WeiboWebResource) FetchChannelFeed(ctx context.Context, request WeiboWebFetchChannelFeedRequest) (*WeiboWebFetchChannelFeedResponse, error) {
	return r.client.WeiboWebGetChannelFeedByName(ctx, request)
}

// WeiboWebFetchUserInfoRequest is the request for GET /api/v1/weibo/web/fetch_user_info.
type WeiboWebFetchUserInfoRequest = WeiboWebGetUserInformationRequest

// WeiboWebFetchUserInfoResponse is the response for GET /api/v1/weibo/web/fetch_user_info.
type WeiboWebFetchUserInfoResponse = WeiboWebGetUserInformationResponse

// FetchUserInfo 获取用户信息/Get user information
//
// GET /api/v1/weibo/web/fetch_user_info
func (r WeiboWebResource) FetchUserInfo(ctx context.Context, request WeiboWebFetchUserInfoRequest) (*WeiboWebFetchUserInfoResponse, error) {
	return r.client.WeiboWebGetUserInformation(ctx, request)
}

// WeiboWebFetchUserPostsRequest is the request for GET /api/v1/weibo/web/fetch_user_posts.
type WeiboWebFetchUserPostsRequest = WeiboWebGetUserPostsRequest

// WeiboWebFetchUserPostsResponse is the response for GET /api/v1/weibo/web/fetch_user_posts.
type WeiboWebFetchUserPostsResponse = WeiboWebGetUserPostsResponse

// FetchUserPosts 获取用户微博列表/Get user posts
//
// GET /api/v1/weibo/web/fetch_user_posts
func (r WeiboWebResource) FetchUserPosts(ctx context.Context, request WeiboWebFetchUserPostsRequest) (*WeiboWebFetchUserPostsResponse, error) {
	return r.client.WeiboWebGetUserPosts(ctx, request)
}

// WeiboWebFetchPostDetailRequest is the request for GET /api/v1/weibo/web/fetch_post_detail.
type WeiboWebFetchPostDetailRequest = WeiboWebGetPostDetailRequest

// WeiboWebFetchPostDetailResponse is the response for GET /api/v1/weibo/web/fetch_post_detail.
type WeiboWebFetchPostDetailResponse = WeiboWebGetPostDetailResponse

// FetchPostDetail 获取微博详情/Get post detail
//
// GET /api/v1/weibo/web/fetch_post_detail
func (r WeiboWebResource) FetchPostDetail(ctx context.Context, request WeiboWebFetchPostDetailRequest) (*WeiboWebFetchPostDetailResponse, error) {
	return r.client.WeiboWebGetPostDetail(ctx, request)
}

// WeiboWebFetchPostCommentsRequest is the request for GET /api/v1/weibo/web/fetch_post_comments.
type WeiboWebFetchPostCommentsRequest = WeiboWebGetPostCommentsRequest

// WeiboWebFetchPostCommentsResponse is the response for GET /api/v1/weibo/web/fetch_post_comments.
type WeiboWebFetchPostCommentsResponse = WeiboWebGetPostCommentsResponse

// FetchPostComments 获取微博评论/Get post comments
//
// GET /api/v1/weibo/web/fetch_post_comments
func (r WeiboWebResource) FetchPostComments(ctx context.Context, request WeiboWebFetchPostCommentsRequest) (*WeiboWebFetchPostCommentsResponse, error) {
	return r.client.WeiboWebGetPostComments(ctx, request)
}

// WeiboWebFetchCommentRepliesRequest is the request for GET /api/v1/weibo/web/fetch_comment_replies.
type WeiboWebFetchCommentRepliesRequest = WeiboWebGetCommentRepliesRequest

// WeiboWebFetchCommentRepliesResponse is the response for GET /api/v1/weibo/web/fetch_comment_replies.
type WeiboWebFetchCommentRepliesResponse = WeiboWebGetCommentRepliesResponse

// FetchCommentReplies 获取评论子评论/Get comment replies
//
// GET /api/v1/weibo/web/fetch_comment_replies
func (r WeiboWebResource) FetchCommentReplies(ctx context.Context, request WeiboWebFetchCommentRepliesRequest) (*WeiboWebFetchCommentRepliesResponse, error) {
	return r.client.WeiboWebGetCommentReplies(ctx, request)
}

// WeiboWebFetchSearchRequest is the request for GET /api/v1/weibo/web/fetch_search.
type WeiboWebFetchSearchRequest = WeiboWebSearchWeiboRequest

// WeiboWebFetchSearchResponse is the response for GET /api/v1/weibo/web/fetch_search.
type WeiboWebFetchSearchResponse = WeiboWebSearchWeiboResponse

// FetchSearch 搜索微博/Search Weibo
//
// GET /api/v1/weibo/web/fetch_search
func (r WeiboWebResource) FetchSearch(ctx context.Context, request WeiboWebFetchSearchRequest) (*WeiboWebFetchSearchResponse, error) {
	return r.client.WeiboWebSearchWeibo(ctx, request)
}

// WeiboWebFetchHotSearchResponse is the response for GET /api/v1/weibo/web/fetch_hot_search.
type WeiboWebFetchHotSearchResponse = WeiboWebGetHotSearchRankingResponse

// FetchHotSearch 获取热搜榜/Get hot search ranking
//
// GET /api/v1/weibo/web/fetch_hot_search
func (r WeiboWebResource) FetchHotSearch(ctx context.Context) (*WeiboWebFetchHotSearchResponse, error) {
	return r.client.WeiboWebGetHotSearchRanking(ctx)
}

// WeiboWebFetchSearchTopicsResponse is the response for GET /api/v1/weibo/web/fetch_search_topics.
type WeiboWebFetchSearchTopicsResponse = WeiboWebGetSearchPageHotTopicsResponse

// FetchSearchTopics 获取搜索页热搜词/Get search page hot topics
//
// GET /api/v1/weibo/web/fetch_search_topics
func (r WeiboWebResource) FetchSearchTopics(ctx context.Context) (*WeiboWebFetchSearchTopicsResponse, error) {
	return r.client.WeiboWebGetSearchPageHotTopics(ctx)
}

// WeiboWebV2Resource contains endpoints from the Weibo-Web-V2-API tag.
type WeiboWebV2Resource struct {
	client *Client
}

// WeiboWebV2CheckAllowCommentWithPicRequest is the request for GET /api/v1/weibo/web_v2/check_allow_comment_with_pic.
type WeiboWebV2CheckAllowCommentWithPicRequest = WeiboWebV2CheckIfWeiboAllowsImageCommentsRequest

// WeiboWebV2CheckAllowCommentWithPicResponse is the response for GET /api/v1/weibo/web_v2/check_allow_comment_with_pic.
type WeiboWebV2CheckAllowCommentWithPicResponse = WeiboWebV2CheckIfWeiboAllowsImageCommentsResponse

// CheckAllowCommentWithPic 检查微博是否允许带图评论/Check if Weibo allows image comments
//
// GET /api/v1/weibo/web_v2/check_allow_comment_with_pic
func (r WeiboWebV2Resource) CheckAllowCommentWithPic(ctx context.Context, request WeiboWebV2CheckAllowCommentWithPicRequest) (*WeiboWebV2CheckAllowCommentWithPicResponse, error) {
	return r.client.WeiboWebV2CheckIfWeiboAllowsImageComments(ctx, request)
}

// WeiboWebV2FetchPostDetailRequest is the request for GET /api/v1/weibo/web_v2/fetch_post_detail.
type WeiboWebV2FetchPostDetailRequest = WeiboWebV2GetSinglePostDataRequest

// WeiboWebV2FetchPostDetailResponse is the response for GET /api/v1/weibo/web_v2/fetch_post_detail.
type WeiboWebV2FetchPostDetailResponse = WeiboWebV2GetSinglePostDataResponse

// FetchPostDetail 获取单个作品数据/Get single post data
//
// GET /api/v1/weibo/web_v2/fetch_post_detail
func (r WeiboWebV2Resource) FetchPostDetail(ctx context.Context, request WeiboWebV2FetchPostDetailRequest) (*WeiboWebV2FetchPostDetailResponse, error) {
	return r.client.WeiboWebV2GetSinglePostData(ctx, request)
}

// WeiboWebV2FetchUserInfoRequest is the request for GET /api/v1/weibo/web_v2/fetch_user_info.
type WeiboWebV2FetchUserInfoRequest = WeiboWebV2GetUserInformationRequest

// WeiboWebV2FetchUserInfoResponse is the response for GET /api/v1/weibo/web_v2/fetch_user_info.
type WeiboWebV2FetchUserInfoResponse = WeiboWebV2GetUserInformationResponse

// FetchUserInfo 获取用户信息/Get user information
//
// GET /api/v1/weibo/web_v2/fetch_user_info
func (r WeiboWebV2Resource) FetchUserInfo(ctx context.Context, request WeiboWebV2FetchUserInfoRequest) (*WeiboWebV2FetchUserInfoResponse, error) {
	return r.client.WeiboWebV2GetUserInformation(ctx, request)
}

// WeiboWebV2FetchUserBasicInfoRequest is the request for GET /api/v1/weibo/web_v2/fetch_user_basic_info.
type WeiboWebV2FetchUserBasicInfoRequest = WeiboWebV2GetUserBasicInformationRequest

// WeiboWebV2FetchUserBasicInfoResponse is the response for GET /api/v1/weibo/web_v2/fetch_user_basic_info.
type WeiboWebV2FetchUserBasicInfoResponse = WeiboWebV2GetUserBasicInformationResponse

// FetchUserBasicInfo 获取用户基本信息/Get user basic information
//
// GET /api/v1/weibo/web_v2/fetch_user_basic_info
func (r WeiboWebV2Resource) FetchUserBasicInfo(ctx context.Context, request WeiboWebV2FetchUserBasicInfoRequest) (*WeiboWebV2FetchUserBasicInfoResponse, error) {
	return r.client.WeiboWebV2GetUserBasicInformation(ctx, request)
}

// WeiboWebV2FetchUserPostsRequest is the request for GET /api/v1/weibo/web_v2/fetch_user_posts.
type WeiboWebV2FetchUserPostsRequest = WeiboWebV2GetWeiboUserPostsRequest

// WeiboWebV2FetchUserPostsResponse is the response for GET /api/v1/weibo/web_v2/fetch_user_posts.
type WeiboWebV2FetchUserPostsResponse = WeiboWebV2GetWeiboUserPostsResponse

// FetchUserPosts 获取微博用户文章数据/Get Weibo user posts
//
// GET /api/v1/weibo/web_v2/fetch_user_posts
func (r WeiboWebV2Resource) FetchUserPosts(ctx context.Context, request WeiboWebV2FetchUserPostsRequest) (*WeiboWebV2FetchUserPostsResponse, error) {
	return r.client.WeiboWebV2GetWeiboUserPosts(ctx, request)
}

// WeiboWebV2FetchUserOriginalPostsRequest is the request for GET /api/v1/weibo/web_v2/fetch_user_original_posts.
type WeiboWebV2FetchUserOriginalPostsRequest = WeiboWebV2GetWeiboUserOriginalPostsRequest

// WeiboWebV2FetchUserOriginalPostsResponse is the response for GET /api/v1/weibo/web_v2/fetch_user_original_posts.
type WeiboWebV2FetchUserOriginalPostsResponse = WeiboWebV2GetWeiboUserOriginalPostsResponse

// FetchUserOriginalPosts 获取微博用户原创微博数据/Get Weibo user original posts
//
// GET /api/v1/weibo/web_v2/fetch_user_original_posts
func (r WeiboWebV2Resource) FetchUserOriginalPosts(ctx context.Context, request WeiboWebV2FetchUserOriginalPostsRequest) (*WeiboWebV2FetchUserOriginalPostsResponse, error) {
	return r.client.WeiboWebV2GetWeiboUserOriginalPosts(ctx, request)
}

// WeiboWebV2FetchPostCommentsRequest is the request for GET /api/v1/weibo/web_v2/fetch_post_comments.
type WeiboWebV2FetchPostCommentsRequest = WeiboWebV2GetWeiboCommentsRequest

// WeiboWebV2FetchPostCommentsResponse is the response for GET /api/v1/weibo/web_v2/fetch_post_comments.
type WeiboWebV2FetchPostCommentsResponse = WeiboWebV2GetWeiboCommentsResponse

// FetchPostComments 获取微博评论/Get Weibo comments
//
// GET /api/v1/weibo/web_v2/fetch_post_comments
func (r WeiboWebV2Resource) FetchPostComments(ctx context.Context, request WeiboWebV2FetchPostCommentsRequest) (*WeiboWebV2FetchPostCommentsResponse, error) {
	return r.client.WeiboWebV2GetWeiboComments(ctx, request)
}

// WeiboWebV2FetchPostSubCommentsRequest is the request for GET /api/v1/weibo/web_v2/fetch_post_sub_comments.
type WeiboWebV2FetchPostSubCommentsRequest = WeiboWebV2GetWeiboSubCommentsRequest

// WeiboWebV2FetchPostSubCommentsResponse is the response for GET /api/v1/weibo/web_v2/fetch_post_sub_comments.
type WeiboWebV2FetchPostSubCommentsResponse = WeiboWebV2GetWeiboSubCommentsResponse

// FetchPostSubComments 获取微博子评论/Get Weibo sub-comments
//
// GET /api/v1/weibo/web_v2/fetch_post_sub_comments
func (r WeiboWebV2Resource) FetchPostSubComments(ctx context.Context, request WeiboWebV2FetchPostSubCommentsRequest) (*WeiboWebV2FetchPostSubCommentsResponse, error) {
	return r.client.WeiboWebV2GetWeiboSubComments(ctx, request)
}

// SearchUserPosts 搜索用户微博/Search user posts
//
// GET /api/v1/weibo/web_v2/search_user_posts
func (r WeiboWebV2Resource) SearchUserPosts(ctx context.Context, request WeiboWebV2SearchUserPostsRequest) (*WeiboWebV2SearchUserPostsResponse, error) {
	return r.client.WeiboWebV2SearchUserPosts(ctx, request)
}

// WeiboWebV2FetchUserVideoCollectionListRequest is the request for GET /api/v1/weibo/web_v2/fetch_user_video_collection_list.
type WeiboWebV2FetchUserVideoCollectionListRequest = WeiboWebV2GetUserVideoCollectionListRequest

// WeiboWebV2FetchUserVideoCollectionListResponse is the response for GET /api/v1/weibo/web_v2/fetch_user_video_collection_list.
type WeiboWebV2FetchUserVideoCollectionListResponse = WeiboWebV2GetUserVideoCollectionListResponse

// FetchUserVideoCollectionList 获取用户微博视频收藏夹列表/Get user video collection list
//
// GET /api/v1/weibo/web_v2/fetch_user_video_collection_list
func (r WeiboWebV2Resource) FetchUserVideoCollectionList(ctx context.Context, request WeiboWebV2FetchUserVideoCollectionListRequest) (*WeiboWebV2FetchUserVideoCollectionListResponse, error) {
	return r.client.WeiboWebV2GetUserVideoCollectionList(ctx, request)
}

// WeiboWebV2FetchUserVideoCollectionDetailRequest is the request for GET /api/v1/weibo/web_v2/fetch_user_video_collection_detail.
type WeiboWebV2FetchUserVideoCollectionDetailRequest = WeiboWebV2GetUserVideoCollectionDetailRequest

// WeiboWebV2FetchUserVideoCollectionDetailResponse is the response for GET /api/v1/weibo/web_v2/fetch_user_video_collection_detail.
type WeiboWebV2FetchUserVideoCollectionDetailResponse = WeiboWebV2GetUserVideoCollectionDetailResponse

// FetchUserVideoCollectionDetail 获取用户微博视频收藏夹详情/Get user video collection detail
//
// GET /api/v1/weibo/web_v2/fetch_user_video_collection_detail
func (r WeiboWebV2Resource) FetchUserVideoCollectionDetail(ctx context.Context, request WeiboWebV2FetchUserVideoCollectionDetailRequest) (*WeiboWebV2FetchUserVideoCollectionDetailResponse, error) {
	return r.client.WeiboWebV2GetUserVideoCollectionDetail(ctx, request)
}

// WeiboWebV2FetchUserVideoListRequest is the request for GET /api/v1/weibo/web_v2/fetch_user_video_list.
type WeiboWebV2FetchUserVideoListRequest = WeiboWebV2GetUserAllVideosRequest

// WeiboWebV2FetchUserVideoListResponse is the response for GET /api/v1/weibo/web_v2/fetch_user_video_list.
type WeiboWebV2FetchUserVideoListResponse = WeiboWebV2GetUserAllVideosResponse

// FetchUserVideoList 获取微博用户全部视频/Get user all videos
//
// GET /api/v1/weibo/web_v2/fetch_user_video_list
func (r WeiboWebV2Resource) FetchUserVideoList(ctx context.Context, request WeiboWebV2FetchUserVideoListRequest) (*WeiboWebV2FetchUserVideoListResponse, error) {
	return r.client.WeiboWebV2GetUserAllVideos(ctx, request)
}

// WeiboWebV2FetchUserFollowingRequest is the request for GET /api/v1/weibo/web_v2/fetch_user_following.
type WeiboWebV2FetchUserFollowingRequest = WeiboWebV2GetUserFollowingListRequest

// WeiboWebV2FetchUserFollowingResponse is the response for GET /api/v1/weibo/web_v2/fetch_user_following.
type WeiboWebV2FetchUserFollowingResponse = WeiboWebV2GetUserFollowingListResponse

// FetchUserFollowing 获取用户关注列表/Get user following list
//
// GET /api/v1/weibo/web_v2/fetch_user_following
func (r WeiboWebV2Resource) FetchUserFollowing(ctx context.Context, request WeiboWebV2FetchUserFollowingRequest) (*WeiboWebV2FetchUserFollowingResponse, error) {
	return r.client.WeiboWebV2GetUserFollowingList(ctx, request)
}

// WeiboWebV2FetchUserFansRequest is the request for GET /api/v1/weibo/web_v2/fetch_user_fans.
type WeiboWebV2FetchUserFansRequest = WeiboWebV2GetUserFansListRequest

// WeiboWebV2FetchUserFansResponse is the response for GET /api/v1/weibo/web_v2/fetch_user_fans.
type WeiboWebV2FetchUserFansResponse = WeiboWebV2GetUserFansListResponse

// FetchUserFans 获取用户粉丝列表/Get user fans list
//
// GET /api/v1/weibo/web_v2/fetch_user_fans
func (r WeiboWebV2Resource) FetchUserFans(ctx context.Context, request WeiboWebV2FetchUserFansRequest) (*WeiboWebV2FetchUserFansResponse, error) {
	return r.client.WeiboWebV2GetUserFansList(ctx, request)
}

// WeiboWebV2FetchAllGroupsResponse is the response for GET /api/v1/weibo/web_v2/fetch_all_groups.
type WeiboWebV2FetchAllGroupsResponse = WeiboWebV2GetAllGroupsInformationResponse

// FetchAllGroups 获取所有分组信息/Get all groups information
//
// GET /api/v1/weibo/web_v2/fetch_all_groups
func (r WeiboWebV2Resource) FetchAllGroups(ctx context.Context) (*WeiboWebV2FetchAllGroupsResponse, error) {
	return r.client.WeiboWebV2GetAllGroupsInformation(ctx)
}

// WeiboWebV2FetchUserRecommendTimelineRequest is the request for GET /api/v1/weibo/web_v2/fetch_user_recommend_timeline.
type WeiboWebV2FetchUserRecommendTimelineRequest = WeiboWebV2GetUserRecommendTimelineRequest

// WeiboWebV2FetchUserRecommendTimelineResponse is the response for GET /api/v1/weibo/web_v2/fetch_user_recommend_timeline.
type WeiboWebV2FetchUserRecommendTimelineResponse = WeiboWebV2GetUserRecommendTimelineResponse

// FetchUserRecommendTimeline 获取微博主页推荐时间轴/Get user recommend timeline
//
// GET /api/v1/weibo/web_v2/fetch_user_recommend_timeline
func (r WeiboWebV2Resource) FetchUserRecommendTimeline(ctx context.Context, request WeiboWebV2FetchUserRecommendTimelineRequest) (*WeiboWebV2FetchUserRecommendTimelineResponse, error) {
	return r.client.WeiboWebV2GetUserRecommendTimeline(ctx, request)
}

// WeiboWebV2FetchHotRankingTimelineRequest is the request for GET /api/v1/weibo/web_v2/fetch_hot_ranking_timeline.
type WeiboWebV2FetchHotRankingTimelineRequest = WeiboWebV2GetHotRankingTimelineRequest

// WeiboWebV2FetchHotRankingTimelineResponse is the response for GET /api/v1/weibo/web_v2/fetch_hot_ranking_timeline.
type WeiboWebV2FetchHotRankingTimelineResponse = WeiboWebV2GetHotRankingTimelineResponse

// FetchHotRankingTimeline 获取微博热门榜单时间轴/Get hot ranking timeline
//
// GET /api/v1/weibo/web_v2/fetch_hot_ranking_timeline
func (r WeiboWebV2Resource) FetchHotRankingTimeline(ctx context.Context, request WeiboWebV2FetchHotRankingTimelineRequest) (*WeiboWebV2FetchHotRankingTimelineResponse, error) {
	return r.client.WeiboWebV2GetHotRankingTimeline(ctx, request)
}

// WeiboWebV2FetchHotSearchIndexResponse is the response for GET /api/v1/weibo/web_v2/fetch_hot_search_index.
type WeiboWebV2FetchHotSearchIndexResponse = WeiboWebV2GetWeiboHotSearchIndexResponse

// FetchHotSearchIndex 获取微博热搜词条(10条)/Get Weibo hot search index (10 items)
//
// GET /api/v1/weibo/web_v2/fetch_hot_search_index
func (r WeiboWebV2Resource) FetchHotSearchIndex(ctx context.Context) (*WeiboWebV2FetchHotSearchIndexResponse, error) {
	return r.client.WeiboWebV2GetWeiboHotSearchIndex(ctx)
}

// WeiboWebV2FetchHotSearchSummaryResponse is the response for GET /api/v1/weibo/web_v2/fetch_hot_search_summary.
type WeiboWebV2FetchHotSearchSummaryResponse = WeiboWebV2GetWeiboCompleteHotSearchRankingResponse

// FetchHotSearchSummary 获取微博完整热搜榜单(50条)/Get Weibo complete hot search ranking (50 items)
//
// GET /api/v1/weibo/web_v2/fetch_hot_search_summary
func (r WeiboWebV2Resource) FetchHotSearchSummary(ctx context.Context) (*WeiboWebV2FetchHotSearchSummaryResponse, error) {
	return r.client.WeiboWebV2GetWeiboCompleteHotSearchRanking(ctx)
}

// WeiboWebV2FetchHotSearchResponse is the response for GET /api/v1/weibo/web_v2/fetch_hot_search.
type WeiboWebV2FetchHotSearchResponse = WeiboWebV2GetWeiboHotSearchRankingResponse

// FetchHotSearch 获取微博热搜榜单/Get Weibo hot search ranking
//
// GET /api/v1/weibo/web_v2/fetch_hot_search
func (r WeiboWebV2Resource) FetchHotSearch(ctx context.Context) (*WeiboWebV2FetchHotSearchResponse, error) {
	return r.client.WeiboWebV2GetWeiboHotSearchRanking(ctx)
}

// WeiboWebV2FetchEntertainmentRankingResponse is the response for GET /api/v1/weibo/web_v2/fetch_entertainment_ranking.
type WeiboWebV2FetchEntertainmentRankingResponse = WeiboWebV2GetWeiboEntertainmentRankingResponse

// FetchEntertainmentRanking 获取微博文娱榜单/Get Weibo entertainment ranking
//
// GET /api/v1/weibo/web_v2/fetch_entertainment_ranking
func (r WeiboWebV2Resource) FetchEntertainmentRanking(ctx context.Context) (*WeiboWebV2FetchEntertainmentRankingResponse, error) {
	return r.client.WeiboWebV2GetWeiboEntertainmentRanking(ctx)
}

// WeiboWebV2FetchLifeRankingResponse is the response for GET /api/v1/weibo/web_v2/fetch_life_ranking.
type WeiboWebV2FetchLifeRankingResponse = WeiboWebV2GetWeiboLifeRankingResponse

// FetchLifeRanking 获取微博生活榜单/Get Weibo life ranking
//
// GET /api/v1/weibo/web_v2/fetch_life_ranking
func (r WeiboWebV2Resource) FetchLifeRanking(ctx context.Context) (*WeiboWebV2FetchLifeRankingResponse, error) {
	return r.client.WeiboWebV2GetWeiboLifeRanking(ctx)
}

// WeiboWebV2FetchSocialRankingResponse is the response for GET /api/v1/weibo/web_v2/fetch_social_ranking.
type WeiboWebV2FetchSocialRankingResponse = WeiboWebV2GetWeiboSocialRankingResponse

// FetchSocialRanking 获取微博社会榜单/Get Weibo social ranking
//
// GET /api/v1/weibo/web_v2/fetch_social_ranking
func (r WeiboWebV2Resource) FetchSocialRanking(ctx context.Context) (*WeiboWebV2FetchSocialRankingResponse, error) {
	return r.client.WeiboWebV2GetWeiboSocialRanking(ctx)
}

// WeiboWebV2FetchSimilarSearchRequest is the request for GET /api/v1/weibo/web_v2/fetch_similar_search.
type WeiboWebV2FetchSimilarSearchRequest = WeiboWebV2GetWeiboSimilarSearchRecommendationsRequest

// WeiboWebV2FetchSimilarSearchResponse is the response for GET /api/v1/weibo/web_v2/fetch_similar_search.
type WeiboWebV2FetchSimilarSearchResponse = WeiboWebV2GetWeiboSimilarSearchRecommendationsResponse

// FetchSimilarSearch 获取微博相似搜索词推荐/Get Weibo similar search recommendations
//
// GET /api/v1/weibo/web_v2/fetch_similar_search
func (r WeiboWebV2Resource) FetchSimilarSearch(ctx context.Context, request WeiboWebV2FetchSimilarSearchRequest) (*WeiboWebV2FetchSimilarSearchResponse, error) {
	return r.client.WeiboWebV2GetWeiboSimilarSearchRecommendations(ctx, request)
}

// WeiboWebV2FetchAiSearchRequest is the request for GET /api/v1/weibo/web_v2/fetch_ai_search.
type WeiboWebV2FetchAiSearchRequest = WeiboWebV2WeiboAiSearchRequest

// WeiboWebV2FetchAiSearchResponse is the response for GET /api/v1/weibo/web_v2/fetch_ai_search.
type WeiboWebV2FetchAiSearchResponse = WeiboWebV2WeiboAiSearchResponse

// FetchAiSearch 微博智能搜索/Weibo AI Search
//
// GET /api/v1/weibo/web_v2/fetch_ai_search
func (r WeiboWebV2Resource) FetchAiSearch(ctx context.Context, request WeiboWebV2FetchAiSearchRequest) (*WeiboWebV2FetchAiSearchResponse, error) {
	return r.client.WeiboWebV2WeiboAiSearch(ctx, request)
}

// WeiboWebV2FetchAiRelatedSearchRequest is the request for GET /api/v1/weibo/web_v2/fetch_ai_related_search.
type WeiboWebV2FetchAiRelatedSearchRequest = WeiboWebV2WeiboAiSearchContentExtensionRequest

// WeiboWebV2FetchAiRelatedSearchResponse is the response for GET /api/v1/weibo/web_v2/fetch_ai_related_search.
type WeiboWebV2FetchAiRelatedSearchResponse = WeiboWebV2WeiboAiSearchContentExtensionResponse

// FetchAiRelatedSearch 微博AI搜索内容扩展/Weibo AI Search Content Extension
//
// GET /api/v1/weibo/web_v2/fetch_ai_related_search
func (r WeiboWebV2Resource) FetchAiRelatedSearch(ctx context.Context, request WeiboWebV2FetchAiRelatedSearchRequest) (*WeiboWebV2FetchAiRelatedSearchResponse, error) {
	return r.client.WeiboWebV2WeiboAiSearchContentExtension(ctx, request)
}

// WeiboWebV2FetchAdvancedSearchRequest is the request for GET /api/v1/weibo/web_v2/fetch_advanced_search.
type WeiboWebV2FetchAdvancedSearchRequest = WeiboWebV2WeiboAdvancedSearchRequest

// WeiboWebV2FetchAdvancedSearchResponse is the response for GET /api/v1/weibo/web_v2/fetch_advanced_search.
type WeiboWebV2FetchAdvancedSearchResponse = WeiboWebV2WeiboAdvancedSearchResponse

// FetchAdvancedSearch 微博高级搜索/Weibo Advanced Search
//
// GET /api/v1/weibo/web_v2/fetch_advanced_search
func (r WeiboWebV2Resource) FetchAdvancedSearch(ctx context.Context, request WeiboWebV2FetchAdvancedSearchRequest) (*WeiboWebV2FetchAdvancedSearchResponse, error) {
	return r.client.WeiboWebV2WeiboAdvancedSearch(ctx, request)
}

// WeiboWebV2FetchCityListRequest is the request for GET /api/v1/weibo/web_v2/fetch_city_list.
type WeiboWebV2FetchCityListRequest = WeiboWebV2RegionCityListRequest

// WeiboWebV2FetchCityListResponse is the response for GET /api/v1/weibo/web_v2/fetch_city_list.
type WeiboWebV2FetchCityListResponse = WeiboWebV2RegionCityListResponse

// FetchCityList 地区省市映射/Region City List
//
// GET /api/v1/weibo/web_v2/fetch_city_list
func (r WeiboWebV2Resource) FetchCityList(ctx context.Context, request WeiboWebV2FetchCityListRequest) (*WeiboWebV2FetchCityListResponse, error) {
	return r.client.WeiboWebV2RegionCityList(ctx, request)
}

// WeiboWebV2FetchRealtimeSearchRequest is the request for GET /api/v1/weibo/web_v2/fetch_realtime_search.
type WeiboWebV2FetchRealtimeSearchRequest = WeiboWebV2WeiboRealtimeSearchRequest

// WeiboWebV2FetchRealtimeSearchResponse is the response for GET /api/v1/weibo/web_v2/fetch_realtime_search.
type WeiboWebV2FetchRealtimeSearchResponse = WeiboWebV2WeiboRealtimeSearchResponse

// FetchRealtimeSearch 实时搜索/Weibo Realtime Search
//
// GET /api/v1/weibo/web_v2/fetch_realtime_search
func (r WeiboWebV2Resource) FetchRealtimeSearch(ctx context.Context, request WeiboWebV2FetchRealtimeSearchRequest) (*WeiboWebV2FetchRealtimeSearchResponse, error) {
	return r.client.WeiboWebV2WeiboRealtimeSearch(ctx, request)
}

// WeiboWebV2FetchUserSearchRequest is the request for GET /api/v1/weibo/web_v2/fetch_user_search.
type WeiboWebV2FetchUserSearchRequest = WeiboWebV2UserSearchRequest

// WeiboWebV2FetchUserSearchResponse is the response for GET /api/v1/weibo/web_v2/fetch_user_search.
type WeiboWebV2FetchUserSearchResponse = WeiboWebV2UserSearchResponse

// FetchUserSearch 用户搜索/User search
//
// GET /api/v1/weibo/web_v2/fetch_user_search
func (r WeiboWebV2Resource) FetchUserSearch(ctx context.Context, request WeiboWebV2FetchUserSearchRequest) (*WeiboWebV2FetchUserSearchResponse, error) {
	return r.client.WeiboWebV2UserSearch(ctx, request)
}

// WeiboWebV2FetchVideoSearchRequest is the request for GET /api/v1/weibo/web_v2/fetch_video_search.
type WeiboWebV2FetchVideoSearchRequest = WeiboWebV2WeiboVideoSearchRequest

// WeiboWebV2FetchVideoSearchResponse is the response for GET /api/v1/weibo/web_v2/fetch_video_search.
type WeiboWebV2FetchVideoSearchResponse = WeiboWebV2WeiboVideoSearchResponse

// FetchVideoSearch 视频搜索（热门/全部）/Weibo video search (hot/all)
//
// GET /api/v1/weibo/web_v2/fetch_video_search
func (r WeiboWebV2Resource) FetchVideoSearch(ctx context.Context, request WeiboWebV2FetchVideoSearchRequest) (*WeiboWebV2FetchVideoSearchResponse, error) {
	return r.client.WeiboWebV2WeiboVideoSearch(ctx, request)
}

// WeiboWebV2FetchPicSearchRequest is the request for GET /api/v1/weibo/web_v2/fetch_pic_search.
type WeiboWebV2FetchPicSearchRequest = WeiboWebV2WeiboPictureSearchRequest

// WeiboWebV2FetchPicSearchResponse is the response for GET /api/v1/weibo/web_v2/fetch_pic_search.
type WeiboWebV2FetchPicSearchResponse = WeiboWebV2WeiboPictureSearchResponse

// FetchPicSearch 图片搜索/Weibo picture search
//
// GET /api/v1/weibo/web_v2/fetch_pic_search
func (r WeiboWebV2Resource) FetchPicSearch(ctx context.Context, request WeiboWebV2FetchPicSearchRequest) (*WeiboWebV2FetchPicSearchResponse, error) {
	return r.client.WeiboWebV2WeiboPictureSearch(ctx, request)
}

// WeiboWebV2FetchTopicSearchRequest is the request for GET /api/v1/weibo/web_v2/fetch_topic_search.
type WeiboWebV2FetchTopicSearchRequest = WeiboWebV2WeiboTopicSearchRequest

// WeiboWebV2FetchTopicSearchResponse is the response for GET /api/v1/weibo/web_v2/fetch_topic_search.
type WeiboWebV2FetchTopicSearchResponse = WeiboWebV2WeiboTopicSearchResponse

// FetchTopicSearch 话题搜索/Weibo topic search
//
// GET /api/v1/weibo/web_v2/fetch_topic_search
func (r WeiboWebV2Resource) FetchTopicSearch(ctx context.Context, request WeiboWebV2FetchTopicSearchRequest) (*WeiboWebV2FetchTopicSearchResponse, error) {
	return r.client.WeiboWebV2WeiboTopicSearch(ctx, request)
}

// WeiboAppResource contains endpoints from the Weibo-App-API tag.
type WeiboAppResource struct {
	client *Client
}

// WeiboAppFetchUserInfoRequest is the request for GET /api/v1/weibo/app/fetch_user_info.
type WeiboAppFetchUserInfoRequest = WeiboAppGetUserInformationRequest

// WeiboAppFetchUserInfoResponse is the response for GET /api/v1/weibo/app/fetch_user_info.
type WeiboAppFetchUserInfoResponse = WeiboAppGetUserInformationResponse

// FetchUserInfo 获取用户信息/Get user information
//
// GET /api/v1/weibo/app/fetch_user_info
func (r WeiboAppResource) FetchUserInfo(ctx context.Context, request WeiboAppFetchUserInfoRequest) (*WeiboAppFetchUserInfoResponse, error) {
	return r.client.WeiboAppGetUserInformation(ctx, request)
}

// WeiboAppFetchUserInfoDetailRequest is the request for GET /api/v1/weibo/app/fetch_user_info_detail.
type WeiboAppFetchUserInfoDetailRequest = WeiboAppGetUserDetailInformationRequest

// WeiboAppFetchUserInfoDetailResponse is the response for GET /api/v1/weibo/app/fetch_user_info_detail.
type WeiboAppFetchUserInfoDetailResponse = WeiboAppGetUserDetailInformationResponse

// FetchUserInfoDetail 获取用户详细信息/Get user detail information
//
// GET /api/v1/weibo/app/fetch_user_info_detail
func (r WeiboAppResource) FetchUserInfoDetail(ctx context.Context, request WeiboAppFetchUserInfoDetailRequest) (*WeiboAppFetchUserInfoDetailResponse, error) {
	return r.client.WeiboAppGetUserDetailInformation(ctx, request)
}

// WeiboAppFetchUserTimelineRequest is the request for GET /api/v1/weibo/app/fetch_user_timeline.
type WeiboAppFetchUserTimelineRequest = WeiboAppGetUserTimelineRequest

// WeiboAppFetchUserTimelineResponse is the response for GET /api/v1/weibo/app/fetch_user_timeline.
type WeiboAppFetchUserTimelineResponse = WeiboAppGetUserTimelineResponse

// FetchUserTimeline 获取用户发布的微博/Get user timeline
//
// GET /api/v1/weibo/app/fetch_user_timeline
func (r WeiboAppResource) FetchUserTimeline(ctx context.Context, request WeiboAppFetchUserTimelineRequest) (*WeiboAppFetchUserTimelineResponse, error) {
	return r.client.WeiboAppGetUserTimeline(ctx, request)
}

// WeiboAppFetchUserVideosRequest is the request for GET /api/v1/weibo/app/fetch_user_videos.
type WeiboAppFetchUserVideosRequest = WeiboAppGetUserVideosRequest

// WeiboAppFetchUserVideosResponse is the response for GET /api/v1/weibo/app/fetch_user_videos.
type WeiboAppFetchUserVideosResponse = WeiboAppGetUserVideosResponse

// FetchUserVideos 获取用户视频列表/Get user videos
//
// GET /api/v1/weibo/app/fetch_user_videos
func (r WeiboAppResource) FetchUserVideos(ctx context.Context, request WeiboAppFetchUserVideosRequest) (*WeiboAppFetchUserVideosResponse, error) {
	return r.client.WeiboAppGetUserVideos(ctx, request)
}

// WeiboAppFetchUserSuperTopicsRequest is the request for GET /api/v1/weibo/app/fetch_user_super_topics.
type WeiboAppFetchUserSuperTopicsRequest = WeiboAppGetUserSuperTopicsRequest

// WeiboAppFetchUserSuperTopicsResponse is the response for GET /api/v1/weibo/app/fetch_user_super_topics.
type WeiboAppFetchUserSuperTopicsResponse = WeiboAppGetUserSuperTopicsResponse

// FetchUserSuperTopics 获取用户参与的超话列表/Get user super topics
//
// GET /api/v1/weibo/app/fetch_user_super_topics
func (r WeiboAppResource) FetchUserSuperTopics(ctx context.Context, request WeiboAppFetchUserSuperTopicsRequest) (*WeiboAppFetchUserSuperTopicsResponse, error) {
	return r.client.WeiboAppGetUserSuperTopics(ctx, request)
}

// WeiboAppFetchUserAlbumRequest is the request for GET /api/v1/weibo/app/fetch_user_album.
type WeiboAppFetchUserAlbumRequest = WeiboAppGetUserAlbumRequest

// WeiboAppFetchUserAlbumResponse is the response for GET /api/v1/weibo/app/fetch_user_album.
type WeiboAppFetchUserAlbumResponse = WeiboAppGetUserAlbumResponse

// FetchUserAlbum 获取用户相册/Get user album
//
// GET /api/v1/weibo/app/fetch_user_album
func (r WeiboAppResource) FetchUserAlbum(ctx context.Context, request WeiboAppFetchUserAlbumRequest) (*WeiboAppFetchUserAlbumResponse, error) {
	return r.client.WeiboAppGetUserAlbum(ctx, request)
}

// WeiboAppFetchUserArticlesRequest is the request for GET /api/v1/weibo/app/fetch_user_articles.
type WeiboAppFetchUserArticlesRequest = WeiboAppGetUserArticlesRequest

// WeiboAppFetchUserArticlesResponse is the response for GET /api/v1/weibo/app/fetch_user_articles.
type WeiboAppFetchUserArticlesResponse = WeiboAppGetUserArticlesResponse

// FetchUserArticles 获取用户文章列表/Get user articles
//
// GET /api/v1/weibo/app/fetch_user_articles
func (r WeiboAppResource) FetchUserArticles(ctx context.Context, request WeiboAppFetchUserArticlesRequest) (*WeiboAppFetchUserArticlesResponse, error) {
	return r.client.WeiboAppGetUserArticles(ctx, request)
}

// WeiboAppFetchUserAudiosRequest is the request for GET /api/v1/weibo/app/fetch_user_audios.
type WeiboAppFetchUserAudiosRequest = WeiboAppGetUserAudiosRequest

// WeiboAppFetchUserAudiosResponse is the response for GET /api/v1/weibo/app/fetch_user_audios.
type WeiboAppFetchUserAudiosResponse = WeiboAppGetUserAudiosResponse

// FetchUserAudios 获取用户音频列表/Get user audios
//
// GET /api/v1/weibo/app/fetch_user_audios
func (r WeiboAppResource) FetchUserAudios(ctx context.Context, request WeiboAppFetchUserAudiosRequest) (*WeiboAppFetchUserAudiosResponse, error) {
	return r.client.WeiboAppGetUserAudios(ctx, request)
}

// WeiboAppFetchUserProfileFeedRequest is the request for GET /api/v1/weibo/app/fetch_user_profile_feed.
type WeiboAppFetchUserProfileFeedRequest = WeiboAppGetUserProfileFeedRequest

// WeiboAppFetchUserProfileFeedResponse is the response for GET /api/v1/weibo/app/fetch_user_profile_feed.
type WeiboAppFetchUserProfileFeedResponse = WeiboAppGetUserProfileFeedResponse

// FetchUserProfileFeed 获取用户主页动态/Get user profile feed
//
// GET /api/v1/weibo/app/fetch_user_profile_feed
func (r WeiboAppResource) FetchUserProfileFeed(ctx context.Context, request WeiboAppFetchUserProfileFeedRequest) (*WeiboAppFetchUserProfileFeedResponse, error) {
	return r.client.WeiboAppGetUserProfileFeed(ctx, request)
}

// WeiboAppFetchStatusDetailRequest is the request for GET /api/v1/weibo/app/fetch_status_detail.
type WeiboAppFetchStatusDetailRequest = WeiboAppGetPostDetailRequest

// WeiboAppFetchStatusDetailResponse is the response for GET /api/v1/weibo/app/fetch_status_detail.
type WeiboAppFetchStatusDetailResponse = WeiboAppGetPostDetailResponse

// FetchStatusDetail 获取微博详情/Get post detail
//
// GET /api/v1/weibo/app/fetch_status_detail
func (r WeiboAppResource) FetchStatusDetail(ctx context.Context, request WeiboAppFetchStatusDetailRequest) (*WeiboAppFetchStatusDetailResponse, error) {
	return r.client.WeiboAppGetPostDetail(ctx, request)
}

// WeiboAppFetchStatusCommentsRequest is the request for GET /api/v1/weibo/app/fetch_status_comments.
type WeiboAppFetchStatusCommentsRequest = WeiboAppGetPostCommentsRequest

// WeiboAppFetchStatusCommentsResponse is the response for GET /api/v1/weibo/app/fetch_status_comments.
type WeiboAppFetchStatusCommentsResponse = WeiboAppGetPostCommentsResponse

// FetchStatusComments 获取微博评论/Get post comments
//
// GET /api/v1/weibo/app/fetch_status_comments
func (r WeiboAppResource) FetchStatusComments(ctx context.Context, request WeiboAppFetchStatusCommentsRequest) (*WeiboAppFetchStatusCommentsResponse, error) {
	return r.client.WeiboAppGetPostComments(ctx, request)
}

// WeiboAppFetchStatusRepostsRequest is the request for GET /api/v1/weibo/app/fetch_status_reposts.
type WeiboAppFetchStatusRepostsRequest = WeiboAppGetPostRepostsRequest

// WeiboAppFetchStatusRepostsResponse is the response for GET /api/v1/weibo/app/fetch_status_reposts.
type WeiboAppFetchStatusRepostsResponse = WeiboAppGetPostRepostsResponse

// FetchStatusReposts 获取微博转发列表/Get post reposts
//
// GET /api/v1/weibo/app/fetch_status_reposts
func (r WeiboAppResource) FetchStatusReposts(ctx context.Context, request WeiboAppFetchStatusRepostsRequest) (*WeiboAppFetchStatusRepostsResponse, error) {
	return r.client.WeiboAppGetPostReposts(ctx, request)
}

// WeiboAppFetchStatusLikesRequest is the request for GET /api/v1/weibo/app/fetch_status_likes.
type WeiboAppFetchStatusLikesRequest = WeiboAppGetPostLikesRequest

// WeiboAppFetchStatusLikesResponse is the response for GET /api/v1/weibo/app/fetch_status_likes.
type WeiboAppFetchStatusLikesResponse = WeiboAppGetPostLikesResponse

// FetchStatusLikes 获取微博点赞列表/Get post likes
//
// GET /api/v1/weibo/app/fetch_status_likes
func (r WeiboAppResource) FetchStatusLikes(ctx context.Context, request WeiboAppFetchStatusLikesRequest) (*WeiboAppFetchStatusLikesResponse, error) {
	return r.client.WeiboAppGetPostLikes(ctx, request)
}

// WeiboAppFetchVideoDetailRequest is the request for GET /api/v1/weibo/app/fetch_video_detail.
type WeiboAppFetchVideoDetailRequest = WeiboAppGetVideoDetailRequest

// WeiboAppFetchVideoDetailResponse is the response for GET /api/v1/weibo/app/fetch_video_detail.
type WeiboAppFetchVideoDetailResponse = WeiboAppGetVideoDetailResponse

// FetchVideoDetail 获取视频详情/Get video detail
//
// GET /api/v1/weibo/app/fetch_video_detail
func (r WeiboAppResource) FetchVideoDetail(ctx context.Context, request WeiboAppFetchVideoDetailRequest) (*WeiboAppFetchVideoDetailResponse, error) {
	return r.client.WeiboAppGetVideoDetail(ctx, request)
}

// WeiboAppFetchVideoFeaturedFeedRequest is the request for GET /api/v1/weibo/app/fetch_video_featured_feed.
type WeiboAppFetchVideoFeaturedFeedRequest = WeiboAppGetVideoFeaturedFeedRequest

// WeiboAppFetchVideoFeaturedFeedResponse is the response for GET /api/v1/weibo/app/fetch_video_featured_feed.
type WeiboAppFetchVideoFeaturedFeedResponse = WeiboAppGetVideoFeaturedFeedResponse

// FetchVideoFeaturedFeed 获取短视频精选Feed流/Get video featured feed
//
// GET /api/v1/weibo/app/fetch_video_featured_feed
func (r WeiboAppResource) FetchVideoFeaturedFeed(ctx context.Context, request WeiboAppFetchVideoFeaturedFeedRequest) (*WeiboAppFetchVideoFeaturedFeedResponse, error) {
	return r.client.WeiboAppGetVideoFeaturedFeed(ctx, request)
}

// WeiboAppFetchSearchAllRequest is the request for GET /api/v1/weibo/app/fetch_search_all.
type WeiboAppFetchSearchAllRequest = WeiboAppComprehensiveSearchRequest

// WeiboAppFetchSearchAllResponse is the response for GET /api/v1/weibo/app/fetch_search_all.
type WeiboAppFetchSearchAllResponse = WeiboAppComprehensiveSearchResponse

// FetchSearchAll 综合搜索/Comprehensive search
//
// GET /api/v1/weibo/app/fetch_search_all
func (r WeiboAppResource) FetchSearchAll(ctx context.Context, request WeiboAppFetchSearchAllRequest) (*WeiboAppFetchSearchAllResponse, error) {
	return r.client.WeiboAppComprehensiveSearch(ctx, request)
}

// WeiboAppFetchAiSmartSearchRequest is the request for GET /api/v1/weibo/app/fetch_ai_smart_search.
type WeiboAppFetchAiSmartSearchRequest = WeiboAppAiSmartSearchRequest

// WeiboAppFetchAiSmartSearchResponse is the response for GET /api/v1/weibo/app/fetch_ai_smart_search.
type WeiboAppFetchAiSmartSearchResponse = WeiboAppAiSmartSearchResponse

// FetchAiSmartSearch AI智搜/AI Smart Search
//
// GET /api/v1/weibo/app/fetch_ai_smart_search
func (r WeiboAppResource) FetchAiSmartSearch(ctx context.Context, request WeiboAppFetchAiSmartSearchRequest) (*WeiboAppFetchAiSmartSearchResponse, error) {
	return r.client.WeiboAppAiSmartSearch(ctx, request)
}

// WeiboAppFetchHomeRecommendFeedRequest is the request for GET /api/v1/weibo/app/fetch_home_recommend_feed.
type WeiboAppFetchHomeRecommendFeedRequest = WeiboAppGetHomeRecommendFeedRequest

// WeiboAppFetchHomeRecommendFeedResponse is the response for GET /api/v1/weibo/app/fetch_home_recommend_feed.
type WeiboAppFetchHomeRecommendFeedResponse = WeiboAppGetHomeRecommendFeedResponse

// FetchHomeRecommendFeed 获取首页推荐Feed流/Get home recommend feed
//
// GET /api/v1/weibo/app/fetch_home_recommend_feed
func (r WeiboAppResource) FetchHomeRecommendFeed(ctx context.Context, request WeiboAppFetchHomeRecommendFeedRequest) (*WeiboAppFetchHomeRecommendFeedResponse, error) {
	return r.client.WeiboAppGetHomeRecommendFeed(ctx, request)
}

// WeiboAppFetchHotSearchRequest is the request for GET /api/v1/weibo/app/fetch_hot_search.
type WeiboAppFetchHotSearchRequest = WeiboAppGetHotSearchRequest

// WeiboAppFetchHotSearchResponse is the response for GET /api/v1/weibo/app/fetch_hot_search.
type WeiboAppFetchHotSearchResponse = WeiboAppGetHotSearchResponse

// FetchHotSearch 获取热搜榜/Get hot search
//
// GET /api/v1/weibo/app/fetch_hot_search
func (r WeiboAppResource) FetchHotSearch(ctx context.Context, request WeiboAppFetchHotSearchRequest) (*WeiboAppFetchHotSearchResponse, error) {
	return r.client.WeiboAppGetHotSearch(ctx, request)
}

// WeiboAppFetchHotSearchCategoriesResponse is the response for GET /api/v1/weibo/app/fetch_hot_search_categories.
type WeiboAppFetchHotSearchCategoriesResponse = WeiboAppGetHotSearchCategoriesResponse

// FetchHotSearchCategories 获取热搜分类列表/Get hot search categories
//
// GET /api/v1/weibo/app/fetch_hot_search_categories
func (r WeiboAppResource) FetchHotSearchCategories(ctx context.Context) (*WeiboAppFetchHotSearchCategoriesResponse, error) {
	return r.client.WeiboAppGetHotSearchCategories(ctx)
}

// WeChatMediaPlatformWebResource contains endpoints from the WeChat-Media-Platform-Web-API tag.
type WeChatMediaPlatformWebResource struct {
	client *Client
}

// WeChatMediaPlatformWebFetchMpArticleDetailJSONRequest is the request for GET /api/v1/wechat_mp/web/fetch_mp_article_detail_json.
type WeChatMediaPlatformWebFetchMpArticleDetailJSONRequest = WeChatMediaPlatformWebGetWechatMpArticleDetailJSONRequest

// WeChatMediaPlatformWebFetchMpArticleDetailJSONResponse is the response for GET /api/v1/wechat_mp/web/fetch_mp_article_detail_json.
type WeChatMediaPlatformWebFetchMpArticleDetailJSONResponse = WeChatMediaPlatformWebGetWechatMpArticleDetailJSONResponse

// FetchMpArticleDetailJSON 获取微信公众号文章详情的JSON/Get Wechat MP Article Detail JSON
//
// GET /api/v1/wechat_mp/web/fetch_mp_article_detail_json
func (r WeChatMediaPlatformWebResource) FetchMpArticleDetailJSON(ctx context.Context, request WeChatMediaPlatformWebFetchMpArticleDetailJSONRequest) (*WeChatMediaPlatformWebFetchMpArticleDetailJSONResponse, error) {
	return r.client.WeChatMediaPlatformWebGetWechatMpArticleDetailJSON(ctx, request)
}

// WeChatMediaPlatformWebFetchMpArticleDetailHTMLRequest is the request for GET /api/v1/wechat_mp/web/fetch_mp_article_detail_html.
type WeChatMediaPlatformWebFetchMpArticleDetailHTMLRequest = WeChatMediaPlatformWebGetWechatMpArticleDetailHtmlRequest

// WeChatMediaPlatformWebFetchMpArticleDetailHTMLResponse is the response for GET /api/v1/wechat_mp/web/fetch_mp_article_detail_html.
type WeChatMediaPlatformWebFetchMpArticleDetailHTMLResponse = WeChatMediaPlatformWebGetWechatMpArticleDetailHtmlResponse

// FetchMpArticleDetailHTML 获取微信公众号文章详情的HTML/Get Wechat MP Article Detail HTML
//
// GET /api/v1/wechat_mp/web/fetch_mp_article_detail_html
func (r WeChatMediaPlatformWebResource) FetchMpArticleDetailHTML(ctx context.Context, request WeChatMediaPlatformWebFetchMpArticleDetailHTMLRequest) (*WeChatMediaPlatformWebFetchMpArticleDetailHTMLResponse, error) {
	return r.client.WeChatMediaPlatformWebGetWechatMpArticleDetailHtml(ctx, request)
}

// WeChatMediaPlatformWebFetchMpArticleListRequest is the request for GET /api/v1/wechat_mp/web/fetch_mp_article_list.
type WeChatMediaPlatformWebFetchMpArticleListRequest = WeChatMediaPlatformWebGetWechatMpArticleListRequest

// WeChatMediaPlatformWebFetchMpArticleListResponse is the response for GET /api/v1/wechat_mp/web/fetch_mp_article_list.
type WeChatMediaPlatformWebFetchMpArticleListResponse = WeChatMediaPlatformWebGetWechatMpArticleListResponse

// FetchMpArticleList 获取微信公众号文章列表/Get Wechat MP Article List
//
// GET /api/v1/wechat_mp/web/fetch_mp_article_list
func (r WeChatMediaPlatformWebResource) FetchMpArticleList(ctx context.Context, request WeChatMediaPlatformWebFetchMpArticleListRequest) (*WeChatMediaPlatformWebFetchMpArticleListResponse, error) {
	return r.client.WeChatMediaPlatformWebGetWechatMpArticleList(ctx, request)
}

// WeChatMediaPlatformWebFetchMpArticleReadCountRequest is the request for GET /api/v1/wechat_mp/web/fetch_mp_article_read_count.
type WeChatMediaPlatformWebFetchMpArticleReadCountRequest = WeChatMediaPlatformWebGetWechatMpArticleReadCountRequest

// WeChatMediaPlatformWebFetchMpArticleReadCountResponse is the response for GET /api/v1/wechat_mp/web/fetch_mp_article_read_count.
type WeChatMediaPlatformWebFetchMpArticleReadCountResponse = WeChatMediaPlatformWebGetWechatMpArticleReadCountResponse

// FetchMpArticleReadCount 获取微信公众号文章阅读量/Get Wechat MP Article Read Count
//
// GET /api/v1/wechat_mp/web/fetch_mp_article_read_count
func (r WeChatMediaPlatformWebResource) FetchMpArticleReadCount(ctx context.Context, request WeChatMediaPlatformWebFetchMpArticleReadCountRequest) (*WeChatMediaPlatformWebFetchMpArticleReadCountResponse, error) {
	return r.client.WeChatMediaPlatformWebGetWechatMpArticleReadCount(ctx, request)
}

// WeChatMediaPlatformWebFetchMpArticleURLRequest is the request for GET /api/v1/wechat_mp/web/fetch_mp_article_url.
type WeChatMediaPlatformWebFetchMpArticleURLRequest = WeChatMediaPlatformWebGetWechatMpArticleURLRequest

// WeChatMediaPlatformWebFetchMpArticleURLResponse is the response for GET /api/v1/wechat_mp/web/fetch_mp_article_url.
type WeChatMediaPlatformWebFetchMpArticleURLResponse = WeChatMediaPlatformWebGetWechatMpArticleURLResponse

// FetchMpArticleURL 获取微信公众号文章永久链接/Get Wechat MP Article URL
//
// GET /api/v1/wechat_mp/web/fetch_mp_article_url
func (r WeChatMediaPlatformWebResource) FetchMpArticleURL(ctx context.Context, request WeChatMediaPlatformWebFetchMpArticleURLRequest) (*WeChatMediaPlatformWebFetchMpArticleURLResponse, error) {
	return r.client.WeChatMediaPlatformWebGetWechatMpArticleURL(ctx, request)
}

// WeChatMediaPlatformWebFetchMpArticleCommentListRequest is the request for GET /api/v1/wechat_mp/web/fetch_mp_article_comment_list.
type WeChatMediaPlatformWebFetchMpArticleCommentListRequest = WeChatMediaPlatformWebGetWechatMpArticleCommentListRequest

// WeChatMediaPlatformWebFetchMpArticleCommentListResponse is the response for GET /api/v1/wechat_mp/web/fetch_mp_article_comment_list.
type WeChatMediaPlatformWebFetchMpArticleCommentListResponse = WeChatMediaPlatformWebGetWechatMpArticleCommentListResponse

// FetchMpArticleCommentList 获取微信公众号文章评论列表/Get Wechat MP Article Comment List
//
// GET /api/v1/wechat_mp/web/fetch_mp_article_comment_list
func (r WeChatMediaPlatformWebResource) FetchMpArticleCommentList(ctx context.Context, request WeChatMediaPlatformWebFetchMpArticleCommentListRequest) (*WeChatMediaPlatformWebFetchMpArticleCommentListResponse, error) {
	return r.client.WeChatMediaPlatformWebGetWechatMpArticleCommentList(ctx, request)
}

// WeChatMediaPlatformWebFetchMpArticleCommentReplyListRequest is the request for GET /api/v1/wechat_mp/web/fetch_mp_article_comment_reply_list.
type WeChatMediaPlatformWebFetchMpArticleCommentReplyListRequest = WeChatMediaPlatformWebGetWechatMpArticleCommentReplyListRequest

// WeChatMediaPlatformWebFetchMpArticleCommentReplyListResponse is the response for GET /api/v1/wechat_mp/web/fetch_mp_article_comment_reply_list.
type WeChatMediaPlatformWebFetchMpArticleCommentReplyListResponse = WeChatMediaPlatformWebGetWechatMpArticleCommentReplyListResponse

// FetchMpArticleCommentReplyList 获取微信公众号文章评论回复列表/Get Wechat MP Article Comment Reply List
//
// GET /api/v1/wechat_mp/web/fetch_mp_article_comment_reply_list
func (r WeChatMediaPlatformWebResource) FetchMpArticleCommentReplyList(ctx context.Context, request WeChatMediaPlatformWebFetchMpArticleCommentReplyListRequest) (*WeChatMediaPlatformWebFetchMpArticleCommentReplyListResponse, error) {
	return r.client.WeChatMediaPlatformWebGetWechatMpArticleCommentReplyList(ctx, request)
}

// WeChatMediaPlatformWebFetchMpArticleAdRequest is the request for GET /api/v1/wechat_mp/web/fetch_mp_article_ad.
type WeChatMediaPlatformWebFetchMpArticleAdRequest = WeChatMediaPlatformWebGetWechatMpArticleAdRequest

// WeChatMediaPlatformWebFetchMpArticleAdResponse is the response for GET /api/v1/wechat_mp/web/fetch_mp_article_ad.
type WeChatMediaPlatformWebFetchMpArticleAdResponse = WeChatMediaPlatformWebGetWechatMpArticleAdResponse

// FetchMpArticleAd 获取微信公众号广告/Get Wechat MP Article Ad
//
// GET /api/v1/wechat_mp/web/fetch_mp_article_ad
func (r WeChatMediaPlatformWebResource) FetchMpArticleAd(ctx context.Context, request WeChatMediaPlatformWebFetchMpArticleAdRequest) (*WeChatMediaPlatformWebFetchMpArticleAdResponse, error) {
	return r.client.WeChatMediaPlatformWebGetWechatMpArticleAd(ctx, request)
}

// WeChatMediaPlatformWebFetchMpArticleURLConversionRequest is the request for GET /api/v1/wechat_mp/web/fetch_mp_article_url_conversion.
type WeChatMediaPlatformWebFetchMpArticleURLConversionRequest = WeChatMediaPlatformWebGetWechatMpLongURLToShortURLRequest

// WeChatMediaPlatformWebFetchMpArticleURLConversionResponse is the response for GET /api/v1/wechat_mp/web/fetch_mp_article_url_conversion.
type WeChatMediaPlatformWebFetchMpArticleURLConversionResponse = WeChatMediaPlatformWebGetWechatMpLongURLToShortURLResponse

// FetchMpArticleURLConversion 获取微信公众号长链接转短链接/Get Wechat MP Long URL to Short URL
//
// GET /api/v1/wechat_mp/web/fetch_mp_article_url_conversion
func (r WeChatMediaPlatformWebResource) FetchMpArticleURLConversion(ctx context.Context, request WeChatMediaPlatformWebFetchMpArticleURLConversionRequest) (*WeChatMediaPlatformWebFetchMpArticleURLConversionResponse, error) {
	return r.client.WeChatMediaPlatformWebGetWechatMpLongURLToShortURL(ctx, request)
}

// WeChatMediaPlatformWebFetchMpRelatedArticlesRequest is the request for GET /api/v1/wechat_mp/web/fetch_mp_related_articles.
type WeChatMediaPlatformWebFetchMpRelatedArticlesRequest = WeChatMediaPlatformWebGetWechatMpRelatedArticlesRequest

// WeChatMediaPlatformWebFetchMpRelatedArticlesResponse is the response for GET /api/v1/wechat_mp/web/fetch_mp_related_articles.
type WeChatMediaPlatformWebFetchMpRelatedArticlesResponse = WeChatMediaPlatformWebGetWechatMpRelatedArticlesResponse

// FetchMpRelatedArticles 获取微信公众号关联文章/Get Wechat MP Related Articles
//
// GET /api/v1/wechat_mp/web/fetch_mp_related_articles
func (r WeChatMediaPlatformWebResource) FetchMpRelatedArticles(ctx context.Context, request WeChatMediaPlatformWebFetchMpRelatedArticlesRequest) (*WeChatMediaPlatformWebFetchMpRelatedArticlesResponse, error) {
	return r.client.WeChatMediaPlatformWebGetWechatMpRelatedArticles(ctx, request)
}

// WeChatMediaPlatformWebFetchSearchOfficialAccountRequest is the request for GET /api/v1/wechat_mp/web/fetch_search_official_account.
type WeChatMediaPlatformWebFetchSearchOfficialAccountRequest = WeChatMediaPlatformWebSearchWechatOfficialAccountRequest

// WeChatMediaPlatformWebFetchSearchOfficialAccountResponse is the response for GET /api/v1/wechat_mp/web/fetch_search_official_account.
type WeChatMediaPlatformWebFetchSearchOfficialAccountResponse = WeChatMediaPlatformWebSearchWechatOfficialAccountResponse

// FetchSearchOfficialAccount 搜索微信公众号/Search Wechat Official Account
//
// GET /api/v1/wechat_mp/web/fetch_search_official_account
func (r WeChatMediaPlatformWebResource) FetchSearchOfficialAccount(ctx context.Context, request WeChatMediaPlatformWebFetchSearchOfficialAccountRequest) (*WeChatMediaPlatformWebFetchSearchOfficialAccountResponse, error) {
	return r.client.WeChatMediaPlatformWebSearchWechatOfficialAccount(ctx, request)
}

// WeChatMediaPlatformWebFetchSearchArticleRequest is the request for GET /api/v1/wechat_mp/web/fetch_search_article.
type WeChatMediaPlatformWebFetchSearchArticleRequest = WeChatMediaPlatformWebSearchWechatMpArticleRequest

// WeChatMediaPlatformWebFetchSearchArticleResponse is the response for GET /api/v1/wechat_mp/web/fetch_search_article.
type WeChatMediaPlatformWebFetchSearchArticleResponse = WeChatMediaPlatformWebSearchWechatMpArticleResponse

// FetchSearchArticle 搜索微信公众号文章/Search Wechat MP Article
//
// GET /api/v1/wechat_mp/web/fetch_search_article
func (r WeChatMediaPlatformWebResource) FetchSearchArticle(ctx context.Context, request WeChatMediaPlatformWebFetchSearchArticleRequest) (*WeChatMediaPlatformWebFetchSearchArticleResponse, error) {
	return r.client.WeChatMediaPlatformWebSearchWechatMpArticle(ctx, request)
}

// WeChatChannelsResource contains endpoints from the WeChat-Channels-API tag.
type WeChatChannelsResource struct {
	client *Client
}

// WeChatChannelsFetchDefaultSearchRequest is the request for POST /api/v1/wechat_channels/fetch_default_search.
type WeChatChannelsFetchDefaultSearchRequest = WeChatChannelsWeChatChannelsDefaultSearchRequest

// WeChatChannelsFetchDefaultSearchResponse is the response for POST /api/v1/wechat_channels/fetch_default_search.
type WeChatChannelsFetchDefaultSearchResponse = WeChatChannelsWeChatChannelsDefaultSearchResponse

// FetchDefaultSearch 微信视频号默认搜索/WeChat Channels Default Search
//
// POST /api/v1/wechat_channels/fetch_default_search
func (r WeChatChannelsResource) FetchDefaultSearch(ctx context.Context, request WeChatChannelsFetchDefaultSearchRequest) (*WeChatChannelsFetchDefaultSearchResponse, error) {
	return r.client.WeChatChannelsWeChatChannelsDefaultSearch(ctx, request)
}

// WeChatChannelsFetchSearchLatestRequest is the request for GET /api/v1/wechat_channels/fetch_search_latest.
type WeChatChannelsFetchSearchLatestRequest = WeChatChannelsWeChatChannelsSearchLatestVideosRequest

// WeChatChannelsFetchSearchLatestResponse is the response for GET /api/v1/wechat_channels/fetch_search_latest.
type WeChatChannelsFetchSearchLatestResponse = WeChatChannelsWeChatChannelsSearchLatestVideosResponse

// FetchSearchLatest 微信视频号搜索最新视频/WeChat Channels Search Latest Videos
//
// GET /api/v1/wechat_channels/fetch_search_latest
func (r WeChatChannelsResource) FetchSearchLatest(ctx context.Context, request WeChatChannelsFetchSearchLatestRequest) (*WeChatChannelsFetchSearchLatestResponse, error) {
	return r.client.WeChatChannelsWeChatChannelsSearchLatestVideos(ctx, request)
}

// WeChatChannelsFetchSearchOrdinaryRequest is the request for GET /api/v1/wechat_channels/fetch_search_ordinary.
type WeChatChannelsFetchSearchOrdinaryRequest = WeChatChannelsWeChatChannelsComprehensiveSearchRequest

// WeChatChannelsFetchSearchOrdinaryResponse is the response for GET /api/v1/wechat_channels/fetch_search_ordinary.
type WeChatChannelsFetchSearchOrdinaryResponse = WeChatChannelsWeChatChannelsComprehensiveSearchResponse

// FetchSearchOrdinary 微信视频号综合搜索/WeChat Channels Comprehensive Search
//
// GET /api/v1/wechat_channels/fetch_search_ordinary
func (r WeChatChannelsResource) FetchSearchOrdinary(ctx context.Context, request WeChatChannelsFetchSearchOrdinaryRequest) (*WeChatChannelsFetchSearchOrdinaryResponse, error) {
	return r.client.WeChatChannelsWeChatChannelsComprehensiveSearch(ctx, request)
}

// WeChatChannelsFetchUserSearchRequest is the request for GET /api/v1/wechat_channels/fetch_user_search.
type WeChatChannelsFetchUserSearchRequest = WeChatChannelsWeChatChannelsUserSearchRequest

// WeChatChannelsFetchUserSearchResponse is the response for GET /api/v1/wechat_channels/fetch_user_search.
type WeChatChannelsFetchUserSearchResponse = WeChatChannelsWeChatChannelsUserSearchResponse

// FetchUserSearch 微信视频号用户搜索/WeChat Channels User Search
//
// GET /api/v1/wechat_channels/fetch_user_search
func (r WeChatChannelsResource) FetchUserSearch(ctx context.Context, request WeChatChannelsFetchUserSearchRequest) (*WeChatChannelsFetchUserSearchResponse, error) {
	return r.client.WeChatChannelsWeChatChannelsUserSearch(ctx, request)
}

// WeChatChannelsFetchUserSearchV2Request is the request for GET /api/v1/wechat_channels/fetch_user_search_v2.
type WeChatChannelsFetchUserSearchV2Request = WeChatChannelsWeChatChannelsUserSearchV2Request

// WeChatChannelsFetchUserSearchV2Response is the response for GET /api/v1/wechat_channels/fetch_user_search_v2.
type WeChatChannelsFetchUserSearchV2Response = WeChatChannelsWeChatChannelsUserSearchV2Response

// FetchUserSearchV2 微信视频号用户搜索V2/WeChat Channels User Search V2
//
// GET /api/v1/wechat_channels/fetch_user_search_v2
func (r WeChatChannelsResource) FetchUserSearchV2(ctx context.Context, request WeChatChannelsFetchUserSearchV2Request) (*WeChatChannelsFetchUserSearchV2Response, error) {
	return r.client.WeChatChannelsWeChatChannelsUserSearchV2(ctx, request)
}

// WeChatChannelsFetchVideoDetailRequest is the request for GET /api/v1/wechat_channels/fetch_video_detail.
type WeChatChannelsFetchVideoDetailRequest = WeChatChannelsWeChatChannelsVideoDetailRequest

// WeChatChannelsFetchVideoDetailResponse is the response for GET /api/v1/wechat_channels/fetch_video_detail.
type WeChatChannelsFetchVideoDetailResponse = WeChatChannelsWeChatChannelsVideoDetailResponse

// FetchVideoDetail 微信视频号视频详情/WeChat Channels Video Detail
//
// GET /api/v1/wechat_channels/fetch_video_detail
func (r WeChatChannelsResource) FetchVideoDetail(ctx context.Context, request WeChatChannelsFetchVideoDetailRequest) (*WeChatChannelsFetchVideoDetailResponse, error) {
	return r.client.WeChatChannelsWeChatChannelsVideoDetail(ctx, request)
}

// WeChatChannelsFetchVideoByShareURLRequest is the request for GET /api/v1/wechat_channels/fetch_video_by_share_url.
type WeChatChannelsFetchVideoByShareURLRequest = WeChatChannelsWeChatChannelsShareDetailRequest

// WeChatChannelsFetchVideoByShareURLResponse is the response for GET /api/v1/wechat_channels/fetch_video_by_share_url.
type WeChatChannelsFetchVideoByShareURLResponse = WeChatChannelsWeChatChannelsShareDetailResponse

// FetchVideoByShareURL 微信视频号分享详情/WeChat Channels Share Detail
//
// GET /api/v1/wechat_channels/fetch_video_by_share_url
func (r WeChatChannelsResource) FetchVideoByShareURL(ctx context.Context, request WeChatChannelsFetchVideoByShareURLRequest) (*WeChatChannelsFetchVideoByShareURLResponse, error) {
	return r.client.WeChatChannelsWeChatChannelsShareDetail(ctx, request)
}

// WeChatChannelsFetchHomePageRequest is the request for POST /api/v1/wechat_channels/fetch_home_page.
type WeChatChannelsFetchHomePageRequest = WeChatChannelsWeChatChannelsHomePageRequest

// WeChatChannelsFetchHomePageResponse is the response for POST /api/v1/wechat_channels/fetch_home_page.
type WeChatChannelsFetchHomePageResponse = WeChatChannelsWeChatChannelsHomePageResponse

// FetchHomePage 微信视频号主页/WeChat Channels Home Page
//
// POST /api/v1/wechat_channels/fetch_home_page
func (r WeChatChannelsResource) FetchHomePage(ctx context.Context, request WeChatChannelsFetchHomePageRequest) (*WeChatChannelsFetchHomePageResponse, error) {
	return r.client.WeChatChannelsWeChatChannelsHomePage(ctx, request)
}

// WeChatChannelsFetchCommentsRequest is the request for POST /api/v1/wechat_channels/fetch_comments.
type WeChatChannelsFetchCommentsRequest = WeChatChannelsWeChatChannelsCommentsRequest

// WeChatChannelsFetchCommentsResponse is the response for POST /api/v1/wechat_channels/fetch_comments.
type WeChatChannelsFetchCommentsResponse = WeChatChannelsWeChatChannelsCommentsResponse

// FetchComments 微信视频号评论/WeChat Channels Comments
//
// POST /api/v1/wechat_channels/fetch_comments
func (r WeChatChannelsResource) FetchComments(ctx context.Context, request WeChatChannelsFetchCommentsRequest) (*WeChatChannelsFetchCommentsResponse, error) {
	return r.client.WeChatChannelsWeChatChannelsComments(ctx, request)
}

// WeChatChannelsFetchLiveHistoryRequest is the request for GET /api/v1/wechat_channels/fetch_live_history.
type WeChatChannelsFetchLiveHistoryRequest = WeChatChannelsWeChatChannelsLiveHistoryRequest

// WeChatChannelsFetchLiveHistoryResponse is the response for GET /api/v1/wechat_channels/fetch_live_history.
type WeChatChannelsFetchLiveHistoryResponse = WeChatChannelsWeChatChannelsLiveHistoryResponse

// FetchLiveHistory 微信视频号直播回放/WeChat Channels Live History
//
// GET /api/v1/wechat_channels/fetch_live_history
func (r WeChatChannelsResource) FetchLiveHistory(ctx context.Context, request WeChatChannelsFetchLiveHistoryRequest) (*WeChatChannelsFetchLiveHistoryResponse, error) {
	return r.client.WeChatChannelsWeChatChannelsLiveHistory(ctx, request)
}

// WeChatChannelsFetchSearchChannelsRequest is the request for GET /api/v1/wechat_channels/fetch_search_channels.
type WeChatChannelsFetchSearchChannelsRequest = WeChatChannelsSearchWeChatChannelsRequest

// WeChatChannelsFetchSearchChannelsResponse is the response for GET /api/v1/wechat_channels/fetch_search_channels.
type WeChatChannelsFetchSearchChannelsResponse = WeChatChannelsSearchWeChatChannelsResponse

// FetchSearchChannels 微信视频号搜索/Search WeChat Channels
//
// GET /api/v1/wechat_channels/fetch_search_channels
func (r WeChatChannelsResource) FetchSearchChannels(ctx context.Context, request WeChatChannelsFetchSearchChannelsRequest) (*WeChatChannelsFetchSearchChannelsResponse, error) {
	return r.client.WeChatChannelsSearchWeChatChannels(ctx, request)
}

// WeChatChannelsFetchHotWordsResponse is the response for GET /api/v1/wechat_channels/fetch_hot_words.
type WeChatChannelsFetchHotWordsResponse = WeChatChannelsWeChatChannelsHotTopicsResponse

// FetchHotWords 微信视频号热门话题/WeChat Channels Hot Topics
//
// GET /api/v1/wechat_channels/fetch_hot_words
func (r WeChatChannelsResource) FetchHotWords(ctx context.Context) (*WeChatChannelsFetchHotWordsResponse, error) {
	return r.client.WeChatChannelsWeChatChannelsHotTopics(ctx)
}

// InstagramV1Resource contains endpoints from the Instagram-V1-API tag.
type InstagramV1Resource struct {
	client *Client
}

// InstagramV1ShortcodeToMediaIDRequest is the request for GET /api/v1/instagram/v1/shortcode_to_media_id.
type InstagramV1ShortcodeToMediaIDRequest = InstagramV1ConvertShortcodeToMediaIDRequest

// InstagramV1ShortcodeToMediaIDResponse is the response for GET /api/v1/instagram/v1/shortcode_to_media_id.
type InstagramV1ShortcodeToMediaIDResponse = InstagramV1ConvertShortcodeToMediaIDResponse

// ShortcodeToMediaID Shortcode转Media ID/Convert shortcode to media ID
//
// GET /api/v1/instagram/v1/shortcode_to_media_id
func (r InstagramV1Resource) ShortcodeToMediaID(ctx context.Context, request InstagramV1ShortcodeToMediaIDRequest) (*InstagramV1ShortcodeToMediaIDResponse, error) {
	return r.client.InstagramV1ConvertShortcodeToMediaID(ctx, request)
}

// InstagramV1MediaIDToShortcodeRequest is the request for GET /api/v1/instagram/v1/media_id_to_shortcode.
type InstagramV1MediaIDToShortcodeRequest = InstagramV1ConvertMediaIDToShortcodeRequest

// InstagramV1MediaIDToShortcodeResponse is the response for GET /api/v1/instagram/v1/media_id_to_shortcode.
type InstagramV1MediaIDToShortcodeResponse = InstagramV1ConvertMediaIDToShortcodeResponse

// MediaIDToShortcode Media ID转Shortcode/Convert media ID to shortcode
//
// GET /api/v1/instagram/v1/media_id_to_shortcode
func (r InstagramV1Resource) MediaIDToShortcode(ctx context.Context, request InstagramV1MediaIDToShortcodeRequest) (*InstagramV1MediaIDToShortcodeResponse, error) {
	return r.client.InstagramV1ConvertMediaIDToShortcode(ctx, request)
}

// InstagramV1UserIDToUsernameRequest is the request for GET /api/v1/instagram/v1/user_id_to_username.
type InstagramV1UserIDToUsernameRequest = InstagramV1GetUserInfoByUserIDRequest

// InstagramV1UserIDToUsernameResponse is the response for GET /api/v1/instagram/v1/user_id_to_username.
type InstagramV1UserIDToUsernameResponse = InstagramV1GetUserInfoByUserIDResponse

// UserIDToUsername 用户ID转用户信息/Get user info by user ID
//
// GET /api/v1/instagram/v1/user_id_to_username
func (r InstagramV1Resource) UserIDToUsername(ctx context.Context, request InstagramV1UserIDToUsernameRequest) (*InstagramV1UserIDToUsernameResponse, error) {
	return r.client.InstagramV1GetUserInfoByUserID(ctx, request)
}

// InstagramV1FetchUserInfoByUsernameRequest is the request for GET /api/v1/instagram/v1/fetch_user_info_by_username.
type InstagramV1FetchUserInfoByUsernameRequest = InstagramV1GetUserDataByUsernameRequest

// InstagramV1FetchUserInfoByUsernameResponse is the response for GET /api/v1/instagram/v1/fetch_user_info_by_username.
type InstagramV1FetchUserInfoByUsernameResponse = InstagramV1GetUserDataByUsernameResponse

// FetchUserInfoByUsername 根据用户名获取用户数据/Get user data by username
//
// GET /api/v1/instagram/v1/fetch_user_info_by_username
func (r InstagramV1Resource) FetchUserInfoByUsername(ctx context.Context, request InstagramV1FetchUserInfoByUsernameRequest) (*InstagramV1FetchUserInfoByUsernameResponse, error) {
	return r.client.InstagramV1GetUserDataByUsername(ctx, request)
}

// InstagramV1FetchUserInfoByUsernameV2Request is the request for GET /api/v1/instagram/v1/fetch_user_info_by_username_v2.
type InstagramV1FetchUserInfoByUsernameV2Request = InstagramV1GetUserDataByUsernameV2Request

// InstagramV1FetchUserInfoByUsernameV2Response is the response for GET /api/v1/instagram/v1/fetch_user_info_by_username_v2.
type InstagramV1FetchUserInfoByUsernameV2Response = InstagramV1GetUserDataByUsernameV2Response

// FetchUserInfoByUsernameV2 根据用户名获取用户数据V2/Get user data by username V2
//
// GET /api/v1/instagram/v1/fetch_user_info_by_username_v2
func (r InstagramV1Resource) FetchUserInfoByUsernameV2(ctx context.Context, request InstagramV1FetchUserInfoByUsernameV2Request) (*InstagramV1FetchUserInfoByUsernameV2Response, error) {
	return r.client.InstagramV1GetUserDataByUsernameV2(ctx, request)
}

// InstagramV1FetchUserInfoByUsernameV3Request is the request for GET /api/v1/instagram/v1/fetch_user_info_by_username_v3.
type InstagramV1FetchUserInfoByUsernameV3Request = InstagramV1GetUserDataByUsernameV3Request

// InstagramV1FetchUserInfoByUsernameV3Response is the response for GET /api/v1/instagram/v1/fetch_user_info_by_username_v3.
type InstagramV1FetchUserInfoByUsernameV3Response = InstagramV1GetUserDataByUsernameV3Response

// FetchUserInfoByUsernameV3 根据用户名获取用户数据V3/Get user data by username V3
//
// GET /api/v1/instagram/v1/fetch_user_info_by_username_v3
func (r InstagramV1Resource) FetchUserInfoByUsernameV3(ctx context.Context, request InstagramV1FetchUserInfoByUsernameV3Request) (*InstagramV1FetchUserInfoByUsernameV3Response, error) {
	return r.client.InstagramV1GetUserDataByUsernameV3(ctx, request)
}

// InstagramV1FetchUserInfoByIDRequest is the request for GET /api/v1/instagram/v1/fetch_user_info_by_id.
type InstagramV1FetchUserInfoByIDRequest = InstagramV1GetUserDataByUserIDRequest

// InstagramV1FetchUserInfoByIDResponse is the response for GET /api/v1/instagram/v1/fetch_user_info_by_id.
type InstagramV1FetchUserInfoByIDResponse = InstagramV1GetUserDataByUserIDResponse

// FetchUserInfoByID 根据用户ID获取用户数据/Get user data by user ID
//
// GET /api/v1/instagram/v1/fetch_user_info_by_id
func (r InstagramV1Resource) FetchUserInfoByID(ctx context.Context, request InstagramV1FetchUserInfoByIDRequest) (*InstagramV1FetchUserInfoByIDResponse, error) {
	return r.client.InstagramV1GetUserDataByUserID(ctx, request)
}

// InstagramV1FetchUserInfoByIDV2Request is the request for GET /api/v1/instagram/v1/fetch_user_info_by_id_v2.
type InstagramV1FetchUserInfoByIDV2Request = InstagramV1GetUserDataByUserIDV2Request

// InstagramV1FetchUserInfoByIDV2Response is the response for GET /api/v1/instagram/v1/fetch_user_info_by_id_v2.
type InstagramV1FetchUserInfoByIDV2Response = InstagramV1GetUserDataByUserIDV2Response

// FetchUserInfoByIDV2 根据用户ID获取用户数据V2/Get user data by user ID V2
//
// GET /api/v1/instagram/v1/fetch_user_info_by_id_v2
func (r InstagramV1Resource) FetchUserInfoByIDV2(ctx context.Context, request InstagramV1FetchUserInfoByIDV2Request) (*InstagramV1FetchUserInfoByIDV2Response, error) {
	return r.client.InstagramV1GetUserDataByUserIDV2(ctx, request)
}

// InstagramV1FetchUserAboutInfoRequest is the request for GET /api/v1/instagram/v1/fetch_user_about_info.
type InstagramV1FetchUserAboutInfoRequest = InstagramV1GetUserAboutInfoRequest

// InstagramV1FetchUserAboutInfoResponse is the response for GET /api/v1/instagram/v1/fetch_user_about_info.
type InstagramV1FetchUserAboutInfoResponse = InstagramV1GetUserAboutInfoResponse

// FetchUserAboutInfo 获取用户的About信息/Get user about info
//
// GET /api/v1/instagram/v1/fetch_user_about_info
func (r InstagramV1Resource) FetchUserAboutInfo(ctx context.Context, request InstagramV1FetchUserAboutInfoRequest) (*InstagramV1FetchUserAboutInfoResponse, error) {
	return r.client.InstagramV1GetUserAboutInfo(ctx, request)
}

// InstagramV1FetchUserPostsRequest is the request for GET /api/v1/instagram/v1/fetch_user_posts.
type InstagramV1FetchUserPostsRequest = InstagramV1GetUserPostsListRequest

// InstagramV1FetchUserPostsResponse is the response for GET /api/v1/instagram/v1/fetch_user_posts.
type InstagramV1FetchUserPostsResponse = InstagramV1GetUserPostsListResponse

// FetchUserPosts 获取用户帖子列表/Get user posts list
//
// GET /api/v1/instagram/v1/fetch_user_posts
func (r InstagramV1Resource) FetchUserPosts(ctx context.Context, request InstagramV1FetchUserPostsRequest) (*InstagramV1FetchUserPostsResponse, error) {
	return r.client.InstagramV1GetUserPostsList(ctx, request)
}

// InstagramV1FetchUserPostsV2Request is the request for GET /api/v1/instagram/v1/fetch_user_posts_v2.
type InstagramV1FetchUserPostsV2Request = InstagramV1GetUserPostsListV2Request

// InstagramV1FetchUserPostsV2Response is the response for GET /api/v1/instagram/v1/fetch_user_posts_v2.
type InstagramV1FetchUserPostsV2Response = InstagramV1GetUserPostsListV2Response

// FetchUserPostsV2 获取用户帖子列表V2/Get user posts list V2
//
// GET /api/v1/instagram/v1/fetch_user_posts_v2
func (r InstagramV1Resource) FetchUserPostsV2(ctx context.Context, request InstagramV1FetchUserPostsV2Request) (*InstagramV1FetchUserPostsV2Response, error) {
	return r.client.InstagramV1GetUserPostsListV2(ctx, request)
}

// InstagramV1FetchUserReelsRequest is the request for GET /api/v1/instagram/v1/fetch_user_reels.
type InstagramV1FetchUserReelsRequest = InstagramV1GetUserReelsListRequest

// InstagramV1FetchUserReelsResponse is the response for GET /api/v1/instagram/v1/fetch_user_reels.
type InstagramV1FetchUserReelsResponse = InstagramV1GetUserReelsListResponse

// FetchUserReels 获取用户Reels列表/Get user Reels list
//
// GET /api/v1/instagram/v1/fetch_user_reels
func (r InstagramV1Resource) FetchUserReels(ctx context.Context, request InstagramV1FetchUserReelsRequest) (*InstagramV1FetchUserReelsResponse, error) {
	return r.client.InstagramV1GetUserReelsList(ctx, request)
}

// InstagramV1FetchUserRepostsRequest is the request for GET /api/v1/instagram/v1/fetch_user_reposts.
type InstagramV1FetchUserRepostsRequest = InstagramV1GetUserRepostsListRequest

// InstagramV1FetchUserRepostsResponse is the response for GET /api/v1/instagram/v1/fetch_user_reposts.
type InstagramV1FetchUserRepostsResponse = InstagramV1GetUserRepostsListResponse

// FetchUserReposts 获取用户转发列表/Get user reposts list
//
// GET /api/v1/instagram/v1/fetch_user_reposts
func (r InstagramV1Resource) FetchUserReposts(ctx context.Context, request InstagramV1FetchUserRepostsRequest) (*InstagramV1FetchUserRepostsResponse, error) {
	return r.client.InstagramV1GetUserRepostsList(ctx, request)
}

// InstagramV1FetchUserTaggedPostsRequest is the request for GET /api/v1/instagram/v1/fetch_user_tagged_posts.
type InstagramV1FetchUserTaggedPostsRequest = InstagramV1GetUserTaggedPostsRequest

// InstagramV1FetchUserTaggedPostsResponse is the response for GET /api/v1/instagram/v1/fetch_user_tagged_posts.
type InstagramV1FetchUserTaggedPostsResponse = InstagramV1GetUserTaggedPostsResponse

// FetchUserTaggedPosts 获取用户被标记的帖子/Get user tagged posts
//
// GET /api/v1/instagram/v1/fetch_user_tagged_posts
func (r InstagramV1Resource) FetchUserTaggedPosts(ctx context.Context, request InstagramV1FetchUserTaggedPostsRequest) (*InstagramV1FetchUserTaggedPostsResponse, error) {
	return r.client.InstagramV1GetUserTaggedPosts(ctx, request)
}

// InstagramV1FetchRelatedProfilesRequest is the request for GET /api/v1/instagram/v1/fetch_related_profiles.
type InstagramV1FetchRelatedProfilesRequest = InstagramV1GetRelatedProfilesRequest

// InstagramV1FetchRelatedProfilesResponse is the response for GET /api/v1/instagram/v1/fetch_related_profiles.
type InstagramV1FetchRelatedProfilesResponse = InstagramV1GetRelatedProfilesResponse

// FetchRelatedProfiles 获取相关用户推荐/Get related profiles
//
// GET /api/v1/instagram/v1/fetch_related_profiles
func (r InstagramV1Resource) FetchRelatedProfiles(ctx context.Context, request InstagramV1FetchRelatedProfilesRequest) (*InstagramV1FetchRelatedProfilesResponse, error) {
	return r.client.InstagramV1GetRelatedProfiles(ctx, request)
}

// InstagramV1FetchSearchRequest is the request for GET /api/v1/instagram/v1/fetch_search.
type InstagramV1FetchSearchRequest = InstagramV1SearchUsersHashtagsPlacesRequest

// InstagramV1FetchSearchResponse is the response for GET /api/v1/instagram/v1/fetch_search.
type InstagramV1FetchSearchResponse = InstagramV1SearchUsersHashtagsPlacesResponse

// FetchSearch 搜索用户/话题/地点/Search users/hashtags/places
//
// GET /api/v1/instagram/v1/fetch_search
func (r InstagramV1Resource) FetchSearch(ctx context.Context, request InstagramV1FetchSearchRequest) (*InstagramV1FetchSearchResponse, error) {
	return r.client.InstagramV1SearchUsersHashtagsPlaces(ctx, request)
}

// InstagramV1FetchPostByURLRequest is the request for GET /api/v1/instagram/v1/fetch_post_by_url.
type InstagramV1FetchPostByURLRequest = InstagramV1GetPostByURLRequest

// InstagramV1FetchPostByURLResponse is the response for GET /api/v1/instagram/v1/fetch_post_by_url.
type InstagramV1FetchPostByURLResponse = InstagramV1GetPostByURLResponse

// FetchPostByURL 通过URL获取帖子详情/Get post by URL
//
// GET /api/v1/instagram/v1/fetch_post_by_url
func (r InstagramV1Resource) FetchPostByURL(ctx context.Context, request InstagramV1FetchPostByURLRequest) (*InstagramV1FetchPostByURLResponse, error) {
	return r.client.InstagramV1GetPostByURL(ctx, request)
}

// InstagramV1FetchPostByURLV2Request is the request for GET /api/v1/instagram/v1/fetch_post_by_url_v2.
type InstagramV1FetchPostByURLV2Request = InstagramV1GetPostByURLV2Request

// InstagramV1FetchPostByURLV2Response is the response for GET /api/v1/instagram/v1/fetch_post_by_url_v2.
type InstagramV1FetchPostByURLV2Response = InstagramV1GetPostByURLV2Response

// FetchPostByURLV2 通过URL获取帖子详情 V2/Get post by URL V2
//
// GET /api/v1/instagram/v1/fetch_post_by_url_v2
func (r InstagramV1Resource) FetchPostByURLV2(ctx context.Context, request InstagramV1FetchPostByURLV2Request) (*InstagramV1FetchPostByURLV2Response, error) {
	return r.client.InstagramV1GetPostByURLV2(ctx, request)
}

// InstagramV1FetchPostByIDRequest is the request for GET /api/v1/instagram/v1/fetch_post_by_id.
type InstagramV1FetchPostByIDRequest = InstagramV1GetPostByIDRequest

// InstagramV1FetchPostByIDResponse is the response for GET /api/v1/instagram/v1/fetch_post_by_id.
type InstagramV1FetchPostByIDResponse = InstagramV1GetPostByIDResponse

// FetchPostByID 通过ID获取帖子详情/Get post by ID
//
// GET /api/v1/instagram/v1/fetch_post_by_id
func (r InstagramV1Resource) FetchPostByID(ctx context.Context, request InstagramV1FetchPostByIDRequest) (*InstagramV1FetchPostByIDResponse, error) {
	return r.client.InstagramV1GetPostByID(ctx, request)
}

// InstagramV1FetchPostCommentsV2Request is the request for GET /api/v1/instagram/v1/fetch_post_comments_v2.
type InstagramV1FetchPostCommentsV2Request = InstagramV1GetPostCommentsV2Request

// InstagramV1FetchPostCommentsV2Response is the response for GET /api/v1/instagram/v1/fetch_post_comments_v2.
type InstagramV1FetchPostCommentsV2Response = InstagramV1GetPostCommentsV2Response

// FetchPostCommentsV2 获取帖子评论列表V2/Get post comments V2
//
// GET /api/v1/instagram/v1/fetch_post_comments_v2
func (r InstagramV1Resource) FetchPostCommentsV2(ctx context.Context, request InstagramV1FetchPostCommentsV2Request) (*InstagramV1FetchPostCommentsV2Response, error) {
	return r.client.InstagramV1GetPostCommentsV2(ctx, request)
}

// InstagramV1FetchCommentRepliesRequest is the request for GET /api/v1/instagram/v1/fetch_comment_replies.
type InstagramV1FetchCommentRepliesRequest = InstagramV1GetCommentRepliesRequest

// InstagramV1FetchCommentRepliesResponse is the response for GET /api/v1/instagram/v1/fetch_comment_replies.
type InstagramV1FetchCommentRepliesResponse = InstagramV1GetCommentRepliesResponse

// FetchCommentReplies 获取评论的子评论列表/Get comment replies
//
// GET /api/v1/instagram/v1/fetch_comment_replies
func (r InstagramV1Resource) FetchCommentReplies(ctx context.Context, request InstagramV1FetchCommentRepliesRequest) (*InstagramV1FetchCommentRepliesResponse, error) {
	return r.client.InstagramV1GetCommentReplies(ctx, request)
}

// InstagramV1FetchMusicPostsRequest is the request for GET /api/v1/instagram/v1/fetch_music_posts.
type InstagramV1FetchMusicPostsRequest = InstagramV1GetPostsUsingSpecificMusicRequest

// InstagramV1FetchMusicPostsResponse is the response for GET /api/v1/instagram/v1/fetch_music_posts.
type InstagramV1FetchMusicPostsResponse = InstagramV1GetPostsUsingSpecificMusicResponse

// FetchMusicPosts 获取使用特定音乐的帖子/Get posts using specific music
//
// GET /api/v1/instagram/v1/fetch_music_posts
func (r InstagramV1Resource) FetchMusicPosts(ctx context.Context, request InstagramV1FetchMusicPostsRequest) (*InstagramV1FetchMusicPostsResponse, error) {
	return r.client.InstagramV1GetPostsUsingSpecificMusic(ctx, request)
}

// InstagramV1FetchHashtagPostsRequest is the request for GET /api/v1/instagram/v1/fetch_hashtag_posts.
type InstagramV1FetchHashtagPostsRequest = InstagramV1GetPostsByHashtagRequest

// InstagramV1FetchHashtagPostsResponse is the response for GET /api/v1/instagram/v1/fetch_hashtag_posts.
type InstagramV1FetchHashtagPostsResponse = InstagramV1GetPostsByHashtagResponse

// FetchHashtagPosts 获取话题标签下的帖子/Get posts by hashtag
//
// GET /api/v1/instagram/v1/fetch_hashtag_posts
func (r InstagramV1Resource) FetchHashtagPosts(ctx context.Context, request InstagramV1FetchHashtagPostsRequest) (*InstagramV1FetchHashtagPostsResponse, error) {
	return r.client.InstagramV1GetPostsByHashtag(ctx, request)
}

// InstagramV1FetchLocationInfoRequest is the request for GET /api/v1/instagram/v1/fetch_location_info.
type InstagramV1FetchLocationInfoRequest = InstagramV1GetLocationInfoRequest

// InstagramV1FetchLocationInfoResponse is the response for GET /api/v1/instagram/v1/fetch_location_info.
type InstagramV1FetchLocationInfoResponse = InstagramV1GetLocationInfoResponse

// FetchLocationInfo 获取地点信息/Get location info
//
// GET /api/v1/instagram/v1/fetch_location_info
func (r InstagramV1Resource) FetchLocationInfo(ctx context.Context, request InstagramV1FetchLocationInfoRequest) (*InstagramV1FetchLocationInfoResponse, error) {
	return r.client.InstagramV1GetLocationInfo(ctx, request)
}

// InstagramV1FetchLocationPostsRequest is the request for GET /api/v1/instagram/v1/fetch_location_posts.
type InstagramV1FetchLocationPostsRequest = InstagramV1GetPostsByLocationRequest

// InstagramV1FetchLocationPostsResponse is the response for GET /api/v1/instagram/v1/fetch_location_posts.
type InstagramV1FetchLocationPostsResponse = InstagramV1GetPostsByLocationResponse

// FetchLocationPosts 获取地点下的帖子/Get posts by location
//
// GET /api/v1/instagram/v1/fetch_location_posts
func (r InstagramV1Resource) FetchLocationPosts(ctx context.Context, request InstagramV1FetchLocationPostsRequest) (*InstagramV1FetchLocationPostsResponse, error) {
	return r.client.InstagramV1GetPostsByLocation(ctx, request)
}

// InstagramV1FetchCitiesRequest is the request for GET /api/v1/instagram/v1/fetch_cities.
type InstagramV1FetchCitiesRequest = InstagramV1GetCitiesByCountryRequest

// InstagramV1FetchCitiesResponse is the response for GET /api/v1/instagram/v1/fetch_cities.
type InstagramV1FetchCitiesResponse = InstagramV1GetCitiesByCountryResponse

// FetchCities 获取国家城市列表/Get cities by country
//
// GET /api/v1/instagram/v1/fetch_cities
func (r InstagramV1Resource) FetchCities(ctx context.Context, request InstagramV1FetchCitiesRequest) (*InstagramV1FetchCitiesResponse, error) {
	return r.client.InstagramV1GetCitiesByCountry(ctx, request)
}

// InstagramV1FetchLocationsRequest is the request for GET /api/v1/instagram/v1/fetch_locations.
type InstagramV1FetchLocationsRequest = InstagramV1GetLocationsByCityRequest

// InstagramV1FetchLocationsResponse is the response for GET /api/v1/instagram/v1/fetch_locations.
type InstagramV1FetchLocationsResponse = InstagramV1GetLocationsByCityResponse

// FetchLocations 获取城市地点列表/Get locations by city
//
// GET /api/v1/instagram/v1/fetch_locations
func (r InstagramV1Resource) FetchLocations(ctx context.Context, request InstagramV1FetchLocationsRequest) (*InstagramV1FetchLocationsResponse, error) {
	return r.client.InstagramV1GetLocationsByCity(ctx, request)
}

// InstagramV1FetchExploreSectionsResponse is the response for GET /api/v1/instagram/v1/fetch_explore_sections.
type InstagramV1FetchExploreSectionsResponse = InstagramV1GetExplorePageSectionsResponse

// FetchExploreSections 获取探索页面分类/Get explore page sections
//
// GET /api/v1/instagram/v1/fetch_explore_sections
func (r InstagramV1Resource) FetchExploreSections(ctx context.Context) (*InstagramV1FetchExploreSectionsResponse, error) {
	return r.client.InstagramV1GetExplorePageSections(ctx)
}

// InstagramV1FetchSectionPostsRequest is the request for GET /api/v1/instagram/v1/fetch_section_posts.
type InstagramV1FetchSectionPostsRequest = InstagramV1GetPostsBySectionRequest

// InstagramV1FetchSectionPostsResponse is the response for GET /api/v1/instagram/v1/fetch_section_posts.
type InstagramV1FetchSectionPostsResponse = InstagramV1GetPostsBySectionResponse

// FetchSectionPosts 获取分类下的帖子/Get posts by section
//
// GET /api/v1/instagram/v1/fetch_section_posts
func (r InstagramV1Resource) FetchSectionPosts(ctx context.Context, request InstagramV1FetchSectionPostsRequest) (*InstagramV1FetchSectionPostsResponse, error) {
	return r.client.InstagramV1GetPostsBySection(ctx, request)
}

// InstagramV2Resource contains endpoints from the Instagram-V2-API tag.
type InstagramV2Resource struct {
	client *Client
}

// InstagramV2ShortcodeToMediaIDRequest is the request for GET /api/v1/instagram/v2/shortcode_to_media_id.
type InstagramV2ShortcodeToMediaIDRequest = InstagramV2ConvertShortcodeToMediaIDRequest

// InstagramV2ShortcodeToMediaIDResponse is the response for GET /api/v1/instagram/v2/shortcode_to_media_id.
type InstagramV2ShortcodeToMediaIDResponse = InstagramV2ConvertShortcodeToMediaIDResponse

// ShortcodeToMediaID Shortcode转Media ID/Convert shortcode to media ID
//
// GET /api/v1/instagram/v2/shortcode_to_media_id
func (r InstagramV2Resource) ShortcodeToMediaID(ctx context.Context, request InstagramV2ShortcodeToMediaIDRequest) (*InstagramV2ShortcodeToMediaIDResponse, error) {
	return r.client.InstagramV2ConvertShortcodeToMediaID(ctx, request)
}

// InstagramV2MediaIDToShortcodeRequest is the request for GET /api/v1/instagram/v2/media_id_to_shortcode.
type InstagramV2MediaIDToShortcodeRequest = InstagramV2ConvertMediaIDToShortcodeRequest

// InstagramV2MediaIDToShortcodeResponse is the response for GET /api/v1/instagram/v2/media_id_to_shortcode.
type InstagramV2MediaIDToShortcodeResponse = InstagramV2ConvertMediaIDToShortcodeResponse

// MediaIDToShortcode Media ID转Shortcode/Convert media ID to shortcode
//
// GET /api/v1/instagram/v2/media_id_to_shortcode
func (r InstagramV2Resource) MediaIDToShortcode(ctx context.Context, request InstagramV2MediaIDToShortcodeRequest) (*InstagramV2MediaIDToShortcodeResponse, error) {
	return r.client.InstagramV2ConvertMediaIDToShortcode(ctx, request)
}

// InstagramV2UserIDToUsernameRequest is the request for GET /api/v1/instagram/v2/user_id_to_username.
type InstagramV2UserIDToUsernameRequest = InstagramV2GetUserInfoByUserIDRequest

// InstagramV2UserIDToUsernameResponse is the response for GET /api/v1/instagram/v2/user_id_to_username.
type InstagramV2UserIDToUsernameResponse = InstagramV2GetUserInfoByUserIDResponse

// UserIDToUsername 用户ID转用户信息/Get user info by user ID
//
// GET /api/v1/instagram/v2/user_id_to_username
func (r InstagramV2Resource) UserIDToUsername(ctx context.Context, request InstagramV2UserIDToUsernameRequest) (*InstagramV2UserIDToUsernameResponse, error) {
	return r.client.InstagramV2GetUserInfoByUserID(ctx, request)
}

// InstagramV2FetchUserInfoRequest is the request for GET /api/v1/instagram/v2/fetch_user_info.
type InstagramV2FetchUserInfoRequest = InstagramV2GetUserInfoRequest

// InstagramV2FetchUserInfoResponse is the response for GET /api/v1/instagram/v2/fetch_user_info.
type InstagramV2FetchUserInfoResponse = InstagramV2GetUserInfoResponse

// FetchUserInfo 获取用户信息/Get user info
//
// GET /api/v1/instagram/v2/fetch_user_info
func (r InstagramV2Resource) FetchUserInfo(ctx context.Context, request InstagramV2FetchUserInfoRequest) (*InstagramV2FetchUserInfoResponse, error) {
	return r.client.InstagramV2GetUserInfo(ctx, request)
}

// InstagramV2FetchUserPostsRequest is the request for GET /api/v1/instagram/v2/fetch_user_posts.
type InstagramV2FetchUserPostsRequest = InstagramV2GetUserPostsRequest

// InstagramV2FetchUserPostsResponse is the response for GET /api/v1/instagram/v2/fetch_user_posts.
type InstagramV2FetchUserPostsResponse = InstagramV2GetUserPostsResponse

// FetchUserPosts 获取用户帖子/Get user posts
//
// GET /api/v1/instagram/v2/fetch_user_posts
func (r InstagramV2Resource) FetchUserPosts(ctx context.Context, request InstagramV2FetchUserPostsRequest) (*InstagramV2FetchUserPostsResponse, error) {
	return r.client.InstagramV2GetUserPosts(ctx, request)
}

// InstagramV2FetchUserReelsRequest is the request for GET /api/v1/instagram/v2/fetch_user_reels.
type InstagramV2FetchUserReelsRequest = InstagramV2GetUserReelsRequest

// InstagramV2FetchUserReelsResponse is the response for GET /api/v1/instagram/v2/fetch_user_reels.
type InstagramV2FetchUserReelsResponse = InstagramV2GetUserReelsResponse

// FetchUserReels 获取用户Reels/Get user reels
//
// GET /api/v1/instagram/v2/fetch_user_reels
func (r InstagramV2Resource) FetchUserReels(ctx context.Context, request InstagramV2FetchUserReelsRequest) (*InstagramV2FetchUserReelsResponse, error) {
	return r.client.InstagramV2GetUserReels(ctx, request)
}

// InstagramV2FetchUserFollowersRequest is the request for GET /api/v1/instagram/v2/fetch_user_followers.
type InstagramV2FetchUserFollowersRequest = InstagramV2GetUserFollowersRequest

// InstagramV2FetchUserFollowersResponse is the response for GET /api/v1/instagram/v2/fetch_user_followers.
type InstagramV2FetchUserFollowersResponse = InstagramV2GetUserFollowersResponse

// FetchUserFollowers 获取用户粉丝/Get user followers
//
// GET /api/v1/instagram/v2/fetch_user_followers
func (r InstagramV2Resource) FetchUserFollowers(ctx context.Context, request InstagramV2FetchUserFollowersRequest) (*InstagramV2FetchUserFollowersResponse, error) {
	return r.client.InstagramV2GetUserFollowers(ctx, request)
}

// InstagramV2FetchUserFollowingRequest is the request for GET /api/v1/instagram/v2/fetch_user_following.
type InstagramV2FetchUserFollowingRequest = InstagramV2GetUserFollowingRequest

// InstagramV2FetchUserFollowingResponse is the response for GET /api/v1/instagram/v2/fetch_user_following.
type InstagramV2FetchUserFollowingResponse = InstagramV2GetUserFollowingResponse

// FetchUserFollowing 获取用户关注/Get user following
//
// GET /api/v1/instagram/v2/fetch_user_following
func (r InstagramV2Resource) FetchUserFollowing(ctx context.Context, request InstagramV2FetchUserFollowingRequest) (*InstagramV2FetchUserFollowingResponse, error) {
	return r.client.InstagramV2GetUserFollowing(ctx, request)
}

// InstagramV2FetchUserStoriesRequest is the request for GET /api/v1/instagram/v2/fetch_user_stories.
type InstagramV2FetchUserStoriesRequest = InstagramV2GetUserStoriesRequest

// InstagramV2FetchUserStoriesResponse is the response for GET /api/v1/instagram/v2/fetch_user_stories.
type InstagramV2FetchUserStoriesResponse = InstagramV2GetUserStoriesResponse

// FetchUserStories 获取用户故事/Get user stories
//
// GET /api/v1/instagram/v2/fetch_user_stories
func (r InstagramV2Resource) FetchUserStories(ctx context.Context, request InstagramV2FetchUserStoriesRequest) (*InstagramV2FetchUserStoriesResponse, error) {
	return r.client.InstagramV2GetUserStories(ctx, request)
}

// InstagramV2FetchUserHighlightsRequest is the request for GET /api/v1/instagram/v2/fetch_user_highlights.
type InstagramV2FetchUserHighlightsRequest = InstagramV2GetUserHighlightsRequest

// InstagramV2FetchUserHighlightsResponse is the response for GET /api/v1/instagram/v2/fetch_user_highlights.
type InstagramV2FetchUserHighlightsResponse = InstagramV2GetUserHighlightsResponse

// FetchUserHighlights 获取用户精选/Get user highlights
//
// GET /api/v1/instagram/v2/fetch_user_highlights
func (r InstagramV2Resource) FetchUserHighlights(ctx context.Context, request InstagramV2FetchUserHighlightsRequest) (*InstagramV2FetchUserHighlightsResponse, error) {
	return r.client.InstagramV2GetUserHighlights(ctx, request)
}

// InstagramV2FetchHighlightStoriesRequest is the request for GET /api/v1/instagram/v2/fetch_highlight_stories.
type InstagramV2FetchHighlightStoriesRequest = InstagramV2GetHighlightStoriesRequest

// InstagramV2FetchHighlightStoriesResponse is the response for GET /api/v1/instagram/v2/fetch_highlight_stories.
type InstagramV2FetchHighlightStoriesResponse = InstagramV2GetHighlightStoriesResponse

// FetchHighlightStories 获取精选故事详情/Get highlight stories
//
// GET /api/v1/instagram/v2/fetch_highlight_stories
func (r InstagramV2Resource) FetchHighlightStories(ctx context.Context, request InstagramV2FetchHighlightStoriesRequest) (*InstagramV2FetchHighlightStoriesResponse, error) {
	return r.client.InstagramV2GetHighlightStories(ctx, request)
}

// InstagramV2FetchUserTaggedPostsRequest is the request for GET /api/v1/instagram/v2/fetch_user_tagged_posts.
type InstagramV2FetchUserTaggedPostsRequest = InstagramV2GetUserTaggedPostsRequest

// InstagramV2FetchUserTaggedPostsResponse is the response for GET /api/v1/instagram/v2/fetch_user_tagged_posts.
type InstagramV2FetchUserTaggedPostsResponse = InstagramV2GetUserTaggedPostsResponse

// FetchUserTaggedPosts 获取用户被标记的帖子/Get user tagged posts
//
// GET /api/v1/instagram/v2/fetch_user_tagged_posts
func (r InstagramV2Resource) FetchUserTaggedPosts(ctx context.Context, request InstagramV2FetchUserTaggedPostsRequest) (*InstagramV2FetchUserTaggedPostsResponse, error) {
	return r.client.InstagramV2GetUserTaggedPosts(ctx, request)
}

// InstagramV2FetchSimilarUsersRequest is the request for GET /api/v1/instagram/v2/fetch_similar_users.
type InstagramV2FetchSimilarUsersRequest = InstagramV2GetSimilarUsersRequest

// InstagramV2FetchSimilarUsersResponse is the response for GET /api/v1/instagram/v2/fetch_similar_users.
type InstagramV2FetchSimilarUsersResponse = InstagramV2GetSimilarUsersResponse

// FetchSimilarUsers 获取相似用户/Get similar users
//
// GET /api/v1/instagram/v2/fetch_similar_users
func (r InstagramV2Resource) FetchSimilarUsers(ctx context.Context, request InstagramV2FetchSimilarUsersRequest) (*InstagramV2FetchSimilarUsersResponse, error) {
	return r.client.InstagramV2GetSimilarUsers(ctx, request)
}

// SearchUsers 搜索用户/Search users
//
// GET /api/v1/instagram/v2/search_users
func (r InstagramV2Resource) SearchUsers(ctx context.Context, request InstagramV2SearchUsersRequest) (*InstagramV2SearchUsersResponse, error) {
	return r.client.InstagramV2SearchUsers(ctx, request)
}

// GeneralSearch 综合搜索/General search
//
// GET /api/v1/instagram/v2/general_search
func (r InstagramV2Resource) GeneralSearch(ctx context.Context, request InstagramV2GeneralSearchRequest) (*InstagramV2GeneralSearchResponse, error) {
	return r.client.InstagramV2GeneralSearch(ctx, request)
}

// SearchReels 搜索Reels/Search reels
//
// GET /api/v1/instagram/v2/search_reels
func (r InstagramV2Resource) SearchReels(ctx context.Context, request InstagramV2SearchReelsRequest) (*InstagramV2SearchReelsResponse, error) {
	return r.client.InstagramV2SearchReels(ctx, request)
}

// SearchMusic 搜索音乐/Search music
//
// GET /api/v1/instagram/v2/search_music
func (r InstagramV2Resource) SearchMusic(ctx context.Context, request InstagramV2SearchMusicRequest) (*InstagramV2SearchMusicResponse, error) {
	return r.client.InstagramV2SearchMusic(ctx, request)
}

// SearchHashtags 搜索话题标签/Search hashtags
//
// GET /api/v1/instagram/v2/search_hashtags
func (r InstagramV2Resource) SearchHashtags(ctx context.Context, request InstagramV2SearchHashtagsRequest) (*InstagramV2SearchHashtagsResponse, error) {
	return r.client.InstagramV2SearchHashtags(ctx, request)
}

// SearchLocations 搜索地点/Search locations
//
// GET /api/v1/instagram/v2/search_locations
func (r InstagramV2Resource) SearchLocations(ctx context.Context, request InstagramV2SearchLocationsRequest) (*InstagramV2SearchLocationsResponse, error) {
	return r.client.InstagramV2SearchLocations(ctx, request)
}

// InstagramV2SearchByCoordinatesRequest is the request for GET /api/v1/instagram/v2/search_by_coordinates.
type InstagramV2SearchByCoordinatesRequest = InstagramV2SearchLocationsByCoordinatesRequest

// InstagramV2SearchByCoordinatesResponse is the response for GET /api/v1/instagram/v2/search_by_coordinates.
type InstagramV2SearchByCoordinatesResponse = InstagramV2SearchLocationsByCoordinatesResponse

// SearchByCoordinates 根据坐标搜索地点/Search locations by coordinates
//
// GET /api/v1/instagram/v2/search_by_coordinates
func (r InstagramV2Resource) SearchByCoordinates(ctx context.Context, request InstagramV2SearchByCoordinatesRequest) (*InstagramV2SearchByCoordinatesResponse, error) {
	return r.client.InstagramV2SearchLocationsByCoordinates(ctx, request)
}

// InstagramV2FetchPostInfoRequest is the request for GET /api/v1/instagram/v2/fetch_post_info.
type InstagramV2FetchPostInfoRequest = InstagramV2GetPostInfoRequest

// InstagramV2FetchPostInfoResponse is the response for GET /api/v1/instagram/v2/fetch_post_info.
type InstagramV2FetchPostInfoResponse = InstagramV2GetPostInfoResponse

// FetchPostInfo 获取帖子详情/Get post info
//
// GET /api/v1/instagram/v2/fetch_post_info
func (r InstagramV2Resource) FetchPostInfo(ctx context.Context, request InstagramV2FetchPostInfoRequest) (*InstagramV2FetchPostInfoResponse, error) {
	return r.client.InstagramV2GetPostInfo(ctx, request)
}

// InstagramV2FetchPostLikesRequest is the request for GET /api/v1/instagram/v2/fetch_post_likes.
type InstagramV2FetchPostLikesRequest = InstagramV2GetPostLikesRequest

// InstagramV2FetchPostLikesResponse is the response for GET /api/v1/instagram/v2/fetch_post_likes.
type InstagramV2FetchPostLikesResponse = InstagramV2GetPostLikesResponse

// FetchPostLikes 获取帖子点赞列表/Get post likes
//
// GET /api/v1/instagram/v2/fetch_post_likes
func (r InstagramV2Resource) FetchPostLikes(ctx context.Context, request InstagramV2FetchPostLikesRequest) (*InstagramV2FetchPostLikesResponse, error) {
	return r.client.InstagramV2GetPostLikes(ctx, request)
}

// InstagramV2FetchPostCommentsRequest is the request for GET /api/v1/instagram/v2/fetch_post_comments.
type InstagramV2FetchPostCommentsRequest = InstagramV2GetPostCommentsRequest

// InstagramV2FetchPostCommentsResponse is the response for GET /api/v1/instagram/v2/fetch_post_comments.
type InstagramV2FetchPostCommentsResponse = InstagramV2GetPostCommentsResponse

// FetchPostComments 获取帖子评论/Get post comments
//
// GET /api/v1/instagram/v2/fetch_post_comments
func (r InstagramV2Resource) FetchPostComments(ctx context.Context, request InstagramV2FetchPostCommentsRequest) (*InstagramV2FetchPostCommentsResponse, error) {
	return r.client.InstagramV2GetPostComments(ctx, request)
}

// InstagramV2FetchCommentRepliesRequest is the request for GET /api/v1/instagram/v2/fetch_comment_replies.
type InstagramV2FetchCommentRepliesRequest = InstagramV2GetCommentRepliesRequest

// InstagramV2FetchCommentRepliesResponse is the response for GET /api/v1/instagram/v2/fetch_comment_replies.
type InstagramV2FetchCommentRepliesResponse = InstagramV2GetCommentRepliesResponse

// FetchCommentReplies 获取评论回复/Get comment replies
//
// GET /api/v1/instagram/v2/fetch_comment_replies
func (r InstagramV2Resource) FetchCommentReplies(ctx context.Context, request InstagramV2FetchCommentRepliesRequest) (*InstagramV2FetchCommentRepliesResponse, error) {
	return r.client.InstagramV2GetCommentReplies(ctx, request)
}

// InstagramV2FetchMusicPostsRequest is the request for GET /api/v1/instagram/v2/fetch_music_posts.
type InstagramV2FetchMusicPostsRequest = InstagramV2GetMusicPostsRequest

// InstagramV2FetchMusicPostsResponse is the response for GET /api/v1/instagram/v2/fetch_music_posts.
type InstagramV2FetchMusicPostsResponse = InstagramV2GetMusicPostsResponse

// FetchMusicPosts 获取音乐帖子/Get music posts
//
// GET /api/v1/instagram/v2/fetch_music_posts
func (r InstagramV2Resource) FetchMusicPosts(ctx context.Context, request InstagramV2FetchMusicPostsRequest) (*InstagramV2FetchMusicPostsResponse, error) {
	return r.client.InstagramV2GetMusicPosts(ctx, request)
}

// InstagramV2FetchLocationPostsRequest is the request for GET /api/v1/instagram/v2/fetch_location_posts.
type InstagramV2FetchLocationPostsRequest = InstagramV2GetLocationPostsRequest

// InstagramV2FetchLocationPostsResponse is the response for GET /api/v1/instagram/v2/fetch_location_posts.
type InstagramV2FetchLocationPostsResponse = InstagramV2GetLocationPostsResponse

// FetchLocationPosts 获取地点帖子/Get location posts
//
// GET /api/v1/instagram/v2/fetch_location_posts
func (r InstagramV2Resource) FetchLocationPosts(ctx context.Context, request InstagramV2FetchLocationPostsRequest) (*InstagramV2FetchLocationPostsResponse, error) {
	return r.client.InstagramV2GetLocationPosts(ctx, request)
}

// InstagramV2FetchHashtagPostsRequest is the request for GET /api/v1/instagram/v2/fetch_hashtag_posts.
type InstagramV2FetchHashtagPostsRequest = InstagramV2GetHashtagPostsRequest

// InstagramV2FetchHashtagPostsResponse is the response for GET /api/v1/instagram/v2/fetch_hashtag_posts.
type InstagramV2FetchHashtagPostsResponse = InstagramV2GetHashtagPostsResponse

// FetchHashtagPosts 获取话题帖子/Get hashtag posts
//
// GET /api/v1/instagram/v2/fetch_hashtag_posts
func (r InstagramV2Resource) FetchHashtagPosts(ctx context.Context, request InstagramV2FetchHashtagPostsRequest) (*InstagramV2FetchHashtagPostsResponse, error) {
	return r.client.InstagramV2GetHashtagPosts(ctx, request)
}

// InstagramV3Resource contains endpoints from the Instagram-V3-API tag.
type InstagramV3Resource struct {
	client *Client
}

// SearchUsers 搜索用户/Search users
//
// GET /api/v1/instagram/v3/search_users
func (r InstagramV3Resource) SearchUsers(ctx context.Context, request InstagramV3SearchUsersRequest) (*InstagramV3SearchUsersResponse, error) {
	return r.client.InstagramV3SearchUsers(ctx, request)
}

// SearchHashtags 搜索话题标签/Search hashtags
//
// GET /api/v1/instagram/v3/search_hashtags
func (r InstagramV3Resource) SearchHashtags(ctx context.Context, request InstagramV3SearchHashtagsRequest) (*InstagramV3SearchHashtagsResponse, error) {
	return r.client.InstagramV3SearchHashtags(ctx, request)
}

// SearchPlaces 搜索地点/Search places
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/instagram/v3/search_places
func (r InstagramV3Resource) SearchPlaces(ctx context.Context, request InstagramV3SearchPlacesRequest) (*InstagramV3SearchPlacesResponse, error) {
	return r.client.InstagramV3SearchPlaces(ctx, request)
}

// GeneralSearch 综合搜索（支持分页）/General search (with pagination)
//
// GET /api/v1/instagram/v3/general_search
func (r InstagramV3Resource) GeneralSearch(ctx context.Context, request InstagramV3GeneralSearchRequest) (*InstagramV3GeneralSearchResponse, error) {
	return r.client.InstagramV3GeneralSearch(ctx, request)
}

// GetUserIDByUsername 通过用户名获取用户ID/Get user ID by username
//
// GET /api/v1/instagram/v3/get_user_id_by_username
func (r InstagramV3Resource) GetUserIDByUsername(ctx context.Context, request InstagramV3GetUserIDByUsernameRequest) (*InstagramV3GetUserIDByUsernameResponse, error) {
	return r.client.InstagramV3GetUserIDByUsername(ctx, request)
}

// GetUserProfile 获取用户信息/Get user profile
//
// GET /api/v1/instagram/v3/get_user_profile
func (r InstagramV3Resource) GetUserProfile(ctx context.Context, request InstagramV3GetUserProfileRequest) (*InstagramV3GetUserProfileResponse, error) {
	return r.client.InstagramV3GetUserProfile(ctx, request)
}

// InstagramV3GetUserBriefRequest is the request for GET /api/v1/instagram/v3/get_user_brief.
type InstagramV3GetUserBriefRequest = InstagramV3GetUserBriefInfoRequest

// InstagramV3GetUserBriefResponse is the response for GET /api/v1/instagram/v3/get_user_brief.
type InstagramV3GetUserBriefResponse = InstagramV3GetUserBriefInfoResponse

// GetUserBrief 获取用户短详情/Get user brief info
//
// GET /api/v1/instagram/v3/get_user_brief
func (r InstagramV3Resource) GetUserBrief(ctx context.Context, request InstagramV3GetUserBriefRequest) (*InstagramV3GetUserBriefResponse, error) {
	return r.client.InstagramV3GetUserBriefInfo(ctx, request)
}

// GetUserPosts 获取用户帖子列表/Get user posts
//
// GET /api/v1/instagram/v3/get_user_posts
func (r InstagramV3Resource) GetUserPosts(ctx context.Context, request InstagramV3GetUserPostsRequest) (*InstagramV3GetUserPostsResponse, error) {
	return r.client.InstagramV3GetUserPosts(ctx, request)
}

// GetUserTaggedPosts 获取用户被标记的帖子/Get user tagged posts
//
// GET /api/v1/instagram/v3/get_user_tagged_posts
func (r InstagramV3Resource) GetUserTaggedPosts(ctx context.Context, request InstagramV3GetUserTaggedPostsRequest) (*InstagramV3GetUserTaggedPostsResponse, error) {
	return r.client.InstagramV3GetUserTaggedPosts(ctx, request)
}

// GetUserReels 获取用户Reels列表/Get user reels
//
// GET /api/v1/instagram/v3/get_user_reels
func (r InstagramV3Resource) GetUserReels(ctx context.Context, request InstagramV3GetUserReelsRequest) (*InstagramV3GetUserReelsResponse, error) {
	return r.client.InstagramV3GetUserReels(ctx, request)
}

// GetUserHighlights 获取用户精选Highlights列表/Get user highlights
//
// GET /api/v1/instagram/v3/get_user_highlights
func (r InstagramV3Resource) GetUserHighlights(ctx context.Context, request InstagramV3GetUserHighlightsRequest) (*InstagramV3GetUserHighlightsResponse, error) {
	return r.client.InstagramV3GetUserHighlights(ctx, request)
}

// GetHighlightStories 获取Highlight精选详情/Get highlight stories
//
// GET /api/v1/instagram/v3/get_highlight_stories
func (r InstagramV3Resource) GetHighlightStories(ctx context.Context, request InstagramV3GetHighlightStoriesRequest) (*InstagramV3GetHighlightStoriesResponse, error) {
	return r.client.InstagramV3GetHighlightStories(ctx, request)
}

// InstagramV3GetUserAboutRequest is the request for GET /api/v1/instagram/v3/get_user_about.
type InstagramV3GetUserAboutRequest = InstagramV3GetUserAboutInfoRequest

// InstagramV3GetUserAboutResponse is the response for GET /api/v1/instagram/v3/get_user_about.
type InstagramV3GetUserAboutResponse = InstagramV3GetUserAboutInfoResponse

// GetUserAbout 获取用户账户简介/Get user about info
//
// GET /api/v1/instagram/v3/get_user_about
func (r InstagramV3Resource) GetUserAbout(ctx context.Context, request InstagramV3GetUserAboutRequest) (*InstagramV3GetUserAboutResponse, error) {
	return r.client.InstagramV3GetUserAboutInfo(ctx, request)
}

// GetUserFormerUsernames 获取用户曾用用户名/Get user former usernames
//
// GET /api/v1/instagram/v3/get_user_former_usernames
func (r InstagramV3Resource) GetUserFormerUsernames(ctx context.Context, request InstagramV3GetUserFormerUsernamesRequest) (*InstagramV3GetUserFormerUsernamesResponse, error) {
	return r.client.InstagramV3GetUserFormerUsernames(ctx, request)
}

// GetUserStories 获取用户Stories（快拍）/Get user stories
//
// GET /api/v1/instagram/v3/get_user_stories
func (r InstagramV3Resource) GetUserStories(ctx context.Context, request InstagramV3GetUserStoriesRequest) (*InstagramV3GetUserStoriesResponse, error) {
	return r.client.InstagramV3GetUserStories(ctx, request)
}

// InstagramV3GetRecommendedReelsRequest is the request for GET /api/v1/instagram/v3/get_recommended_reels.
type InstagramV3GetRecommendedReelsRequest = InstagramV3GetRecommendedReelsFeedRequest

// InstagramV3GetRecommendedReelsResponse is the response for GET /api/v1/instagram/v3/get_recommended_reels.
type InstagramV3GetRecommendedReelsResponse = InstagramV3GetRecommendedReelsFeedResponse

// GetRecommendedReels 获取Reels推荐列表/Get recommended Reels feed
//
// GET /api/v1/instagram/v3/get_recommended_reels
func (r InstagramV3Resource) GetRecommendedReels(ctx context.Context, request InstagramV3GetRecommendedReelsRequest) (*InstagramV3GetRecommendedReelsResponse, error) {
	return r.client.InstagramV3GetRecommendedReelsFeed(ctx, request)
}

// GetPostInfo 获取帖子详情/Get post info (media_id or URL)
//
// GET /api/v1/instagram/v3/get_post_info
func (r InstagramV3Resource) GetPostInfo(ctx context.Context, request InstagramV3GetPostInfoRequest) (*InstagramV3GetPostInfoResponse, error) {
	return r.client.InstagramV3GetPostInfo(ctx, request)
}

// InstagramV3GetPostInfoByCodeRequest is the request for GET /api/v1/instagram/v3/get_post_info_by_code.
type InstagramV3GetPostInfoByCodeRequest = InstagramV3GetPostInfoByShortcodeRequest

// InstagramV3GetPostInfoByCodeResponse is the response for GET /api/v1/instagram/v3/get_post_info_by_code.
type InstagramV3GetPostInfoByCodeResponse = InstagramV3GetPostInfoByShortcodeResponse

// GetPostInfoByCode 获取帖子详情(code)/Get post info by shortcode
//
// GET /api/v1/instagram/v3/get_post_info_by_code
func (r InstagramV3Resource) GetPostInfoByCode(ctx context.Context, request InstagramV3GetPostInfoByCodeRequest) (*InstagramV3GetPostInfoByCodeResponse, error) {
	return r.client.InstagramV3GetPostInfoByShortcode(ctx, request)
}

// GetPostComments 获取帖子评论/Get post comments
//
// GET /api/v1/instagram/v3/get_post_comments
func (r InstagramV3Resource) GetPostComments(ctx context.Context, request InstagramV3GetPostCommentsRequest) (*InstagramV3GetPostCommentsResponse, error) {
	return r.client.InstagramV3GetPostComments(ctx, request)
}

// GetCommentReplies 获取评论的子评论/回复/Get comment replies
//
// GET /api/v1/instagram/v3/get_comment_replies
func (r InstagramV3Resource) GetCommentReplies(ctx context.Context, request InstagramV3GetCommentRepliesRequest) (*InstagramV3GetCommentRepliesResponse, error) {
	return r.client.InstagramV3GetCommentReplies(ctx, request)
}

// InstagramV3GetPostOembedRequest is the request for GET /api/v1/instagram/v3/get_post_oembed.
type InstagramV3GetPostOembedRequest = InstagramV3GetPostOEmbedInfoRequest

// InstagramV3GetPostOembedResponse is the response for GET /api/v1/instagram/v3/get_post_oembed.
type InstagramV3GetPostOembedResponse = InstagramV3GetPostOEmbedInfoResponse

// GetPostOembed 获取帖子oEmbed内嵌信息/Get post oEmbed info
//
// GET /api/v1/instagram/v3/get_post_oembed
func (r InstagramV3Resource) GetPostOembed(ctx context.Context, request InstagramV3GetPostOembedRequest) (*InstagramV3GetPostOembedResponse, error) {
	return r.client.InstagramV3GetPostOEmbedInfo(ctx, request)
}

// InstagramV3TranslateCommentRequest is the request for GET /api/v1/instagram/v3/translate_comment.
type InstagramV3TranslateCommentRequest = InstagramV3TranslateCommentOrCaptionRequest

// InstagramV3TranslateCommentResponse is the response for GET /api/v1/instagram/v3/translate_comment.
type InstagramV3TranslateCommentResponse = InstagramV3TranslateCommentOrCaptionResponse

// TranslateComment 翻译评论/帖子文本/Translate comment or caption
//
// GET /api/v1/instagram/v3/translate_comment
func (r InstagramV3Resource) TranslateComment(ctx context.Context, request InstagramV3TranslateCommentRequest) (*InstagramV3TranslateCommentResponse, error) {
	return r.client.InstagramV3TranslateCommentOrCaption(ctx, request)
}

// BulkTranslateComments 批量翻译评论/Bulk translate comments
//
// GET /api/v1/instagram/v3/bulk_translate_comments
func (r InstagramV3Resource) BulkTranslateComments(ctx context.Context, request InstagramV3BulkTranslateCommentsRequest) (*InstagramV3BulkTranslateCommentsResponse, error) {
	return r.client.InstagramV3BulkTranslateComments(ctx, request)
}

// InstagramV3GetExploreRequest is the request for GET /api/v1/instagram/v3/get_explore.
type InstagramV3GetExploreRequest = InstagramV3GetExploreFeedRequest

// InstagramV3GetExploreResponse is the response for GET /api/v1/instagram/v3/get_explore.
type InstagramV3GetExploreResponse = InstagramV3GetExploreFeedResponse

// GetExplore 获取探索页推荐帖子/Get explore feed
//
// GET /api/v1/instagram/v3/get_explore
func (r InstagramV3Resource) GetExplore(ctx context.Context, request InstagramV3GetExploreRequest) (*InstagramV3GetExploreResponse, error) {
	return r.client.InstagramV3GetExploreFeed(ctx, request)
}

// InstagramV3GetUserFollowingRequest is the request for GET /api/v1/instagram/v3/get_user_following.
type InstagramV3GetUserFollowingRequest = InstagramV3GetUserFollowingListRequest

// InstagramV3GetUserFollowingResponse is the response for GET /api/v1/instagram/v3/get_user_following.
type InstagramV3GetUserFollowingResponse = InstagramV3GetUserFollowingListResponse

// GetUserFollowing 获取用户关注列表/Get user following list
//
// GET /api/v1/instagram/v3/get_user_following
func (r InstagramV3Resource) GetUserFollowing(ctx context.Context, request InstagramV3GetUserFollowingRequest) (*InstagramV3GetUserFollowingResponse, error) {
	return r.client.InstagramV3GetUserFollowingList(ctx, request)
}

// InstagramV3GetUserFollowersRequest is the request for GET /api/v1/instagram/v3/get_user_followers.
type InstagramV3GetUserFollowersRequest = InstagramV3GetUserFollowersListRequest

// InstagramV3GetUserFollowersResponse is the response for GET /api/v1/instagram/v3/get_user_followers.
type InstagramV3GetUserFollowersResponse = InstagramV3GetUserFollowersListResponse

// GetUserFollowers 获取用户粉丝列表/Get user followers list
//
// GET /api/v1/instagram/v3/get_user_followers
func (r InstagramV3Resource) GetUserFollowers(ctx context.Context, request InstagramV3GetUserFollowersRequest) (*InstagramV3GetUserFollowersResponse, error) {
	return r.client.InstagramV3GetUserFollowersList(ctx, request)
}

// GetLocationInfo 获取地点详情/Get location info
//
// GET /api/v1/instagram/v3/get_location_info
func (r InstagramV3Resource) GetLocationInfo(ctx context.Context, request InstagramV3GetLocationInfoRequest) (*InstagramV3GetLocationInfoResponse, error) {
	return r.client.InstagramV3GetLocationInfo(ctx, request)
}

// GetLocationPosts 获取地点相关帖子/Get location posts
//
// GET /api/v1/instagram/v3/get_location_posts
func (r InstagramV3Resource) GetLocationPosts(ctx context.Context, request InstagramV3GetLocationPostsRequest) (*InstagramV3GetLocationPostsResponse, error) {
	return r.client.InstagramV3GetLocationPosts(ctx, request)
}

// InstagramV3GetLocationNearbyRequest is the request for GET /api/v1/instagram/v3/get_location_nearby.
type InstagramV3GetLocationNearbyRequest = InstagramV3GetNearbyLocationContentRequest

// InstagramV3GetLocationNearbyResponse is the response for GET /api/v1/instagram/v3/get_location_nearby.
type InstagramV3GetLocationNearbyResponse = InstagramV3GetNearbyLocationContentResponse

// GetLocationNearby 获取地点附近内容/Get nearby location content
//
// GET /api/v1/instagram/v3/get_location_nearby
func (r InstagramV3Resource) GetLocationNearby(ctx context.Context, request InstagramV3GetLocationNearbyRequest) (*InstagramV3GetLocationNearbyResponse, error) {
	return r.client.InstagramV3GetNearbyLocationContent(ctx, request)
}

// InstagramV3ShortcodeToMediaIDRequest is the request for GET /api/v1/instagram/v3/shortcode_to_media_id.
type InstagramV3ShortcodeToMediaIDRequest = InstagramV3ConvertShortcodeToMediaIDRequest

// InstagramV3ShortcodeToMediaIDResponse is the response for GET /api/v1/instagram/v3/shortcode_to_media_id.
type InstagramV3ShortcodeToMediaIDResponse = InstagramV3ConvertShortcodeToMediaIDResponse

// ShortcodeToMediaID 短码转媒体ID/Convert shortcode to media ID
//
// GET /api/v1/instagram/v3/shortcode_to_media_id
func (r InstagramV3Resource) ShortcodeToMediaID(ctx context.Context, request InstagramV3ShortcodeToMediaIDRequest) (*InstagramV3ShortcodeToMediaIDResponse, error) {
	return r.client.InstagramV3ConvertShortcodeToMediaID(ctx, request)
}

// InstagramV3MediaIDToShortcodeRequest is the request for GET /api/v1/instagram/v3/media_id_to_shortcode.
type InstagramV3MediaIDToShortcodeRequest = InstagramV3ConvertMediaIDToShortcodeRequest

// InstagramV3MediaIDToShortcodeResponse is the response for GET /api/v1/instagram/v3/media_id_to_shortcode.
type InstagramV3MediaIDToShortcodeResponse = InstagramV3ConvertMediaIDToShortcodeResponse

// MediaIDToShortcode 媒体ID转短码/Convert media ID to shortcode
//
// GET /api/v1/instagram/v3/media_id_to_shortcode
func (r InstagramV3Resource) MediaIDToShortcode(ctx context.Context, request InstagramV3MediaIDToShortcodeRequest) (*InstagramV3MediaIDToShortcodeResponse, error) {
	return r.client.InstagramV3ConvertMediaIDToShortcode(ctx, request)
}

// InstagramV3ExtractShortcodeRequest is the request for GET /api/v1/instagram/v3/extract_shortcode.
type InstagramV3ExtractShortcodeRequest = InstagramV3ExtractShortcodeFromURLRequest

// InstagramV3ExtractShortcodeResponse is the response for GET /api/v1/instagram/v3/extract_shortcode.
type InstagramV3ExtractShortcodeResponse = InstagramV3ExtractShortcodeFromURLResponse

// ExtractShortcode 从URL提取短码/Extract shortcode from URL
//
// GET /api/v1/instagram/v3/extract_shortcode
func (r InstagramV3Resource) ExtractShortcode(ctx context.Context, request InstagramV3ExtractShortcodeRequest) (*InstagramV3ExtractShortcodeResponse, error) {
	return r.client.InstagramV3ExtractShortcodeFromURL(ctx, request)
}

// YouTubeWebResource contains endpoints from the YouTube-Web-API tag.
type YouTubeWebResource struct {
	client *Client
}

// YouTubeWebGetVideoInfoRequest is the request for GET /api/v1/youtube/web/get_video_info.
type YouTubeWebGetVideoInfoRequest = YouTubeWebGetVideoInformationV1Request

// YouTubeWebGetVideoInfoResponse is the response for GET /api/v1/youtube/web/get_video_info.
type YouTubeWebGetVideoInfoResponse = YouTubeWebGetVideoInformationV1Response

// GetVideoInfo 获取视频信息 V1/Get video information V1
//
// GET /api/v1/youtube/web/get_video_info
func (r YouTubeWebResource) GetVideoInfo(ctx context.Context, request YouTubeWebGetVideoInfoRequest) (*YouTubeWebGetVideoInfoResponse, error) {
	return r.client.YouTubeWebGetVideoInformationV1(ctx, request)
}

// YouTubeWebGetVideoInfoV2Request is the request for GET /api/v1/youtube/web/get_video_info_v2.
type YouTubeWebGetVideoInfoV2Request = YouTubeWebGetVideoInformationV2Request

// YouTubeWebGetVideoInfoV2Response is the response for GET /api/v1/youtube/web/get_video_info_v2.
type YouTubeWebGetVideoInfoV2Response = YouTubeWebGetVideoInformationV2Response

// GetVideoInfoV2 获取视频信息 V2/Get video information V2
//
// GET /api/v1/youtube/web/get_video_info_v2
func (r YouTubeWebResource) GetVideoInfoV2(ctx context.Context, request YouTubeWebGetVideoInfoV2Request) (*YouTubeWebGetVideoInfoV2Response, error) {
	return r.client.YouTubeWebGetVideoInformationV2(ctx, request)
}

// YouTubeWebGetVideoInfoV3Request is the request for GET /api/v1/youtube/web/get_video_info_v3.
type YouTubeWebGetVideoInfoV3Request = YouTubeWebGetVideoInformationV3Request

// YouTubeWebGetVideoInfoV3Response is the response for GET /api/v1/youtube/web/get_video_info_v3.
type YouTubeWebGetVideoInfoV3Response = YouTubeWebGetVideoInformationV3Response

// GetVideoInfoV3 获取视频详情 V3/Get video information V3
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/youtube/web/get_video_info_v3
func (r YouTubeWebResource) GetVideoInfoV3(ctx context.Context, request YouTubeWebGetVideoInfoV3Request) (*YouTubeWebGetVideoInfoV3Response, error) {
	return r.client.YouTubeWebGetVideoInformationV3(ctx, request)
}

// GetVideoSubtitles 获取视频字幕/Get video subtitles
//
// GET /api/v1/youtube/web/get_video_subtitles
func (r YouTubeWebResource) GetVideoSubtitles(ctx context.Context, request YouTubeWebGetVideoSubtitlesRequest) (*YouTubeWebGetVideoSubtitlesResponse, error) {
	return r.client.YouTubeWebGetVideoSubtitles(ctx, request)
}

// GetVideoComments 获取视频评论/Get video comments
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/youtube/web/get_video_comments
func (r YouTubeWebResource) GetVideoComments(ctx context.Context, request YouTubeWebGetVideoCommentsRequest) (*YouTubeWebGetVideoCommentsResponse, error) {
	return r.client.YouTubeWebGetVideoComments(ctx, request)
}

// YouTubeWebGetVideoCommentRepliesRequest is the request for GET /api/v1/youtube/web/get_video_comment_replies.
type YouTubeWebGetVideoCommentRepliesRequest = YouTubeWebGetVideoSubCommentsRequest

// YouTubeWebGetVideoCommentRepliesResponse is the response for GET /api/v1/youtube/web/get_video_comment_replies.
type YouTubeWebGetVideoCommentRepliesResponse = YouTubeWebGetVideoSubCommentsResponse

// GetVideoCommentReplies 获取视频二级评论/Get video sub comments
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/youtube/web/get_video_comment_replies
func (r YouTubeWebResource) GetVideoCommentReplies(ctx context.Context, request YouTubeWebGetVideoCommentRepliesRequest) (*YouTubeWebGetVideoCommentRepliesResponse, error) {
	return r.client.YouTubeWebGetVideoSubComments(ctx, request)
}

// GetChannelDescription 获取频道描述信息/Get channel description
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/youtube/web/get_channel_description
func (r YouTubeWebResource) GetChannelDescription(ctx context.Context, request YouTubeWebGetChannelDescriptionRequest) (*YouTubeWebGetChannelDescriptionResponse, error) {
	return r.client.YouTubeWebGetChannelDescription(ctx, request)
}

// YouTubeWebGetRelateVideoRequest is the request for GET /api/v1/youtube/web/get_relate_video.
type YouTubeWebGetRelateVideoRequest = YouTubeWebGetRelatedVideosRequest

// YouTubeWebGetRelateVideoResponse is the response for GET /api/v1/youtube/web/get_relate_video.
type YouTubeWebGetRelateVideoResponse = YouTubeWebGetRelatedVideosResponse

// GetRelateVideo 获取推荐视频/Get related videos
//
// GET /api/v1/youtube/web/get_relate_video
func (r YouTubeWebResource) GetRelateVideo(ctx context.Context, request YouTubeWebGetRelateVideoRequest) (*YouTubeWebGetRelateVideoResponse, error) {
	return r.client.YouTubeWebGetRelatedVideos(ctx, request)
}

// SearchVideo 搜索视频/Search video
//
// GET /api/v1/youtube/web/search_video
func (r YouTubeWebResource) SearchVideo(ctx context.Context, request YouTubeWebSearchVideoRequest) (*YouTubeWebSearchVideoResponse, error) {
	return r.client.YouTubeWebSearchVideo(ctx, request)
}

// YouTubeWebGetGeneralSearchRequest is the request for GET /api/v1/youtube/web/get_general_search.
type YouTubeWebGetGeneralSearchRequest = YouTubeWebGeneralSearchWithFiltersRequest

// YouTubeWebGetGeneralSearchResponse is the response for GET /api/v1/youtube/web/get_general_search.
type YouTubeWebGetGeneralSearchResponse = YouTubeWebGeneralSearchWithFiltersResponse

// GetGeneralSearch 综合搜索（支持过滤条件）/General search with filters
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/youtube/web/get_general_search
func (r YouTubeWebResource) GetGeneralSearch(ctx context.Context, request YouTubeWebGetGeneralSearchRequest) (*YouTubeWebGetGeneralSearchResponse, error) {
	return r.client.YouTubeWebGeneralSearchWithFilters(ctx, request)
}

// YouTubeWebGetShortsSearchRequest is the request for GET /api/v1/youtube/web/get_shorts_search.
type YouTubeWebGetShortsSearchRequest = YouTubeWebYouTubeShortsSearchRequest

// YouTubeWebGetShortsSearchResponse is the response for GET /api/v1/youtube/web/get_shorts_search.
type YouTubeWebGetShortsSearchResponse = YouTubeWebYouTubeShortsSearchResponse

// GetShortsSearch YouTube Shorts短视频搜索/YouTube Shorts search
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/youtube/web/get_shorts_search
func (r YouTubeWebResource) GetShortsSearch(ctx context.Context, request YouTubeWebGetShortsSearchRequest) (*YouTubeWebGetShortsSearchResponse, error) {
	return r.client.YouTubeWebYouTubeShortsSearch(ctx, request)
}

// GetChannelID 获取频道ID/Get channel ID
//
// GET /api/v1/youtube/web/get_channel_id
func (r YouTubeWebResource) GetChannelID(ctx context.Context, request YouTubeWebGetChannelIDRequest) (*YouTubeWebGetChannelIDResponse, error) {
	return r.client.YouTubeWebGetChannelID(ctx, request)
}

// YouTubeWebGetChannelIDV2Request is the request for GET /api/v1/youtube/web/get_channel_id_v2.
type YouTubeWebGetChannelIDV2Request = YouTubeWebGetChannelIDFromURLV2Request

// YouTubeWebGetChannelIDV2Response is the response for GET /api/v1/youtube/web/get_channel_id_v2.
type YouTubeWebGetChannelIDV2Response = YouTubeWebGetChannelIDFromURLV2Response

// GetChannelIDV2 从频道URL获取频道ID V2/Get channel ID from URL V2
//
// GET /api/v1/youtube/web/get_channel_id_v2
func (r YouTubeWebResource) GetChannelIDV2(ctx context.Context, request YouTubeWebGetChannelIDV2Request) (*YouTubeWebGetChannelIDV2Response, error) {
	return r.client.YouTubeWebGetChannelIDFromURLV2(ctx, request)
}

// YouTubeWebGetChannelURLRequest is the request for GET /api/v1/youtube/web/get_channel_url.
type YouTubeWebGetChannelURLRequest = YouTubeWebGetChannelURLFromChannelIDRequest

// YouTubeWebGetChannelURLResponse is the response for GET /api/v1/youtube/web/get_channel_url.
type YouTubeWebGetChannelURLResponse = YouTubeWebGetChannelURLFromChannelIDResponse

// GetChannelURL 从频道ID获取频道URL/Get channel URL from channel ID
//
// GET /api/v1/youtube/web/get_channel_url
func (r YouTubeWebResource) GetChannelURL(ctx context.Context, request YouTubeWebGetChannelURLRequest) (*YouTubeWebGetChannelURLResponse, error) {
	return r.client.YouTubeWebGetChannelURLFromChannelID(ctx, request)
}

// YouTubeWebGetChannelInfoRequest is the request for GET /api/v1/youtube/web/get_channel_info.
type YouTubeWebGetChannelInfoRequest = YouTubeWebGetChannelInformationRequest

// YouTubeWebGetChannelInfoResponse is the response for GET /api/v1/youtube/web/get_channel_info.
type YouTubeWebGetChannelInfoResponse = YouTubeWebGetChannelInformationResponse

// GetChannelInfo 获取频道信息/Get channel information
//
// GET /api/v1/youtube/web/get_channel_info
func (r YouTubeWebResource) GetChannelInfo(ctx context.Context, request YouTubeWebGetChannelInfoRequest) (*YouTubeWebGetChannelInfoResponse, error) {
	return r.client.YouTubeWebGetChannelInformation(ctx, request)
}

// YouTubeWebGetChannelVideosRequest is the request for GET /api/v1/youtube/web/get_channel_videos.
type YouTubeWebGetChannelVideosRequest = YouTubeWebGetChannelVideosV1Request

// YouTubeWebGetChannelVideosResponse is the response for GET /api/v1/youtube/web/get_channel_videos.
type YouTubeWebGetChannelVideosResponse = YouTubeWebGetChannelVideosV1Response

// GetChannelVideos 获取频道视频 V1（即将过时，优先使用 V2）/Get channel videos V1 (deprecated soon, use V2 first)
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/youtube/web/get_channel_videos
func (r YouTubeWebResource) GetChannelVideos(ctx context.Context, request YouTubeWebGetChannelVideosRequest) (*YouTubeWebGetChannelVideosResponse, error) {
	return r.client.YouTubeWebGetChannelVideosV1(ctx, request)
}

// GetChannelVideosV2 获取频道视频 V2/Get channel videos V2
//
// GET /api/v1/youtube/web/get_channel_videos_v2
func (r YouTubeWebResource) GetChannelVideosV2(ctx context.Context, request YouTubeWebGetChannelVideosV2Request) (*YouTubeWebGetChannelVideosV2Response, error) {
	return r.client.YouTubeWebGetChannelVideosV2(ctx, request)
}

// GetChannelVideosV3 获取频道视频 V3/Get channel videos V3
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/youtube/web/get_channel_videos_v3
func (r YouTubeWebResource) GetChannelVideosV3(ctx context.Context, request YouTubeWebGetChannelVideosV3Request) (*YouTubeWebGetChannelVideosV3Response, error) {
	return r.client.YouTubeWebGetChannelVideosV3(ctx, request)
}

// GetChannelShortVideos 获取频道短视频/Get channel short videos
//
// GET /api/v1/youtube/web/get_channel_short_videos
func (r YouTubeWebResource) GetChannelShortVideos(ctx context.Context, request YouTubeWebGetChannelShortVideosRequest) (*YouTubeWebGetChannelShortVideosResponse, error) {
	return r.client.YouTubeWebGetChannelShortVideos(ctx, request)
}

// SearchChannel 搜索频道/Search channel
//
// GET /api/v1/youtube/web/search_channel
func (r YouTubeWebResource) SearchChannel(ctx context.Context, request YouTubeWebSearchChannelRequest) (*YouTubeWebSearchChannelResponse, error) {
	return r.client.YouTubeWebSearchChannel(ctx, request)
}

// GetTrendingVideos 获取趋势视频/Get trending videos
//
// GET /api/v1/youtube/web/get_trending_videos
func (r YouTubeWebResource) GetTrendingVideos(ctx context.Context, request YouTubeWebGetTrendingVideosRequest) (*YouTubeWebGetTrendingVideosResponse, error) {
	return r.client.YouTubeWebGetTrendingVideos(ctx, request)
}

// YouTubeWebV2Resource contains endpoints from the YouTube-Web-V2-API tag.
type YouTubeWebV2Resource struct {
	client *Client
}

// YouTubeWebV2GetVideoInfoRequest is the request for GET /api/v1/youtube/web_v2/get_video_info.
type YouTubeWebV2GetVideoInfoRequest = YouTubeWebV2GetVideoInformationRequest

// YouTubeWebV2GetVideoInfoResponse is the response for GET /api/v1/youtube/web_v2/get_video_info.
type YouTubeWebV2GetVideoInfoResponse = YouTubeWebV2GetVideoInformationResponse

// GetVideoInfo 获取视频详情 /Get video information
//
// GET /api/v1/youtube/web_v2/get_video_info
func (r YouTubeWebV2Resource) GetVideoInfo(ctx context.Context, request YouTubeWebV2GetVideoInfoRequest) (*YouTubeWebV2GetVideoInfoResponse, error) {
	return r.client.YouTubeWebV2GetVideoInformation(ctx, request)
}

// YouTubeWebV2GetVideoInfoV2Request is the request for GET /api/v1/youtube/web_v2/get_video_info_v2.
type YouTubeWebV2GetVideoInfoV2Request = YouTubeWebV2GetVideoInformationV2Request

// YouTubeWebV2GetVideoInfoV2Response is the response for GET /api/v1/youtube/web_v2/get_video_info_v2.
type YouTubeWebV2GetVideoInfoV2Response = YouTubeWebV2GetVideoInformationV2Response

// GetVideoInfoV2 获取视频详情 V2/Get video information V2
//
// GET /api/v1/youtube/web_v2/get_video_info_v2
func (r YouTubeWebV2Resource) GetVideoInfoV2(ctx context.Context, request YouTubeWebV2GetVideoInfoV2Request) (*YouTubeWebV2GetVideoInfoV2Response, error) {
	return r.client.YouTubeWebV2GetVideoInformationV2(ctx, request)
}

// GetVideoComments 获取视频评论/Get video comments
//
// GET /api/v1/youtube/web_v2/get_video_comments
func (r YouTubeWebV2Resource) GetVideoComments(ctx context.Context, request YouTubeWebV2GetVideoCommentsRequest) (*YouTubeWebV2GetVideoCommentsResponse, error) {
	return r.client.YouTubeWebV2GetVideoComments(ctx, request)
}

// YouTubeWebV2GetVideoCommentRepliesRequest is the request for GET /api/v1/youtube/web_v2/get_video_comment_replies.
type YouTubeWebV2GetVideoCommentRepliesRequest = YouTubeWebV2GetVideoSubCommentsRequest

// YouTubeWebV2GetVideoCommentRepliesResponse is the response for GET /api/v1/youtube/web_v2/get_video_comment_replies.
type YouTubeWebV2GetVideoCommentRepliesResponse = YouTubeWebV2GetVideoSubCommentsResponse

// GetVideoCommentReplies 获取视频二级评论/Get video sub comments
//
// GET /api/v1/youtube/web_v2/get_video_comment_replies
func (r YouTubeWebV2Resource) GetVideoCommentReplies(ctx context.Context, request YouTubeWebV2GetVideoCommentRepliesRequest) (*YouTubeWebV2GetVideoCommentRepliesResponse, error) {
	return r.client.YouTubeWebV2GetVideoSubComments(ctx, request)
}

// GetChannelDescription 获取频道描述信息/Get channel description
//
// GET /api/v1/youtube/web_v2/get_channel_description
func (r YouTubeWebV2Resource) GetChannelDescription(ctx context.Context, request YouTubeWebV2GetChannelDescriptionRequest) (*YouTubeWebV2GetChannelDescriptionResponse, error) {
	return r.client.YouTubeWebV2GetChannelDescription(ctx, request)
}

// YouTubeWebV2GetGeneralSearchRequest is the request for GET /api/v1/youtube/web_v2/get_general_search.
type YouTubeWebV2GetGeneralSearchRequest = YouTubeWebV2GeneralSearchRequest

// YouTubeWebV2GetGeneralSearchResponse is the response for GET /api/v1/youtube/web_v2/get_general_search.
type YouTubeWebV2GetGeneralSearchResponse = YouTubeWebV2GeneralSearchResponse

// GetGeneralSearch 综合搜索（原始数据，推荐使用V2）/General search (raw data, recommend V2)
//
// GET /api/v1/youtube/web_v2/get_general_search
func (r YouTubeWebV2Resource) GetGeneralSearch(ctx context.Context, request YouTubeWebV2GetGeneralSearchRequest) (*YouTubeWebV2GetGeneralSearchResponse, error) {
	return r.client.YouTubeWebV2GeneralSearch(ctx, request)
}

// YouTubeWebV2GetGeneralSearchV2Request is the request for GET /api/v1/youtube/web_v2/get_general_search_v2.
type YouTubeWebV2GetGeneralSearchV2Request = YouTubeWebV2GeneralSearchV2Request

// YouTubeWebV2GetGeneralSearchV2Response is the response for GET /api/v1/youtube/web_v2/get_general_search_v2.
type YouTubeWebV2GetGeneralSearchV2Response = YouTubeWebV2GeneralSearchV2Response

// GetGeneralSearchV2 综合搜索V2/General search V2
//
// GET /api/v1/youtube/web_v2/get_general_search_v2
func (r YouTubeWebV2Resource) GetGeneralSearchV2(ctx context.Context, request YouTubeWebV2GetGeneralSearchV2Request) (*YouTubeWebV2GetGeneralSearchV2Response, error) {
	return r.client.YouTubeWebV2GeneralSearchV2(ctx, request)
}

// YouTubeWebV2GetShortsSearchRequest is the request for GET /api/v1/youtube/web_v2/get_shorts_search.
type YouTubeWebV2GetShortsSearchRequest = YouTubeWebV2ShortsSearchRequest

// YouTubeWebV2GetShortsSearchResponse is the response for GET /api/v1/youtube/web_v2/get_shorts_search.
type YouTubeWebV2GetShortsSearchResponse = YouTubeWebV2ShortsSearchResponse

// GetShortsSearch Shorts搜索（原始数据，推荐使用V2）/Shorts search (raw data, recommend V2)
//
// GET /api/v1/youtube/web_v2/get_shorts_search
func (r YouTubeWebV2Resource) GetShortsSearch(ctx context.Context, request YouTubeWebV2GetShortsSearchRequest) (*YouTubeWebV2GetShortsSearchResponse, error) {
	return r.client.YouTubeWebV2ShortsSearch(ctx, request)
}

// YouTubeWebV2GetShortsSearchV2Request is the request for GET /api/v1/youtube/web_v2/get_shorts_search_v2.
type YouTubeWebV2GetShortsSearchV2Request = YouTubeWebV2ShortsSearchV2Request

// YouTubeWebV2GetShortsSearchV2Response is the response for GET /api/v1/youtube/web_v2/get_shorts_search_v2.
type YouTubeWebV2GetShortsSearchV2Response = YouTubeWebV2ShortsSearchV2Response

// GetShortsSearchV2 Shorts搜索V2/Shorts search V2
//
// GET /api/v1/youtube/web_v2/get_shorts_search_v2
func (r YouTubeWebV2Resource) GetShortsSearchV2(ctx context.Context, request YouTubeWebV2GetShortsSearchV2Request) (*YouTubeWebV2GetShortsSearchV2Response, error) {
	return r.client.YouTubeWebV2ShortsSearchV2(ctx, request)
}

// YouTubeWebV2GetChannelIDRequest is the request for GET /api/v1/youtube/web_v2/get_channel_id.
type YouTubeWebV2GetChannelIDRequest = YouTubeWebV2GetChannelIDFromURLRequest

// YouTubeWebV2GetChannelIDResponse is the response for GET /api/v1/youtube/web_v2/get_channel_id.
type YouTubeWebV2GetChannelIDResponse = YouTubeWebV2GetChannelIDFromURLResponse

// GetChannelID 从频道URL获取频道ID /Get channel ID from URL
//
// GET /api/v1/youtube/web_v2/get_channel_id
func (r YouTubeWebV2Resource) GetChannelID(ctx context.Context, request YouTubeWebV2GetChannelIDRequest) (*YouTubeWebV2GetChannelIDResponse, error) {
	return r.client.YouTubeWebV2GetChannelIDFromURL(ctx, request)
}

// YouTubeWebV2GetChannelURLRequest is the request for GET /api/v1/youtube/web_v2/get_channel_url.
type YouTubeWebV2GetChannelURLRequest = YouTubeWebV2GetChannelURLFromChannelIDRequest

// YouTubeWebV2GetChannelURLResponse is the response for GET /api/v1/youtube/web_v2/get_channel_url.
type YouTubeWebV2GetChannelURLResponse = YouTubeWebV2GetChannelURLFromChannelIDResponse

// GetChannelURL 从频道ID获取频道URL/Get channel URL from channel ID
//
// GET /api/v1/youtube/web_v2/get_channel_url
func (r YouTubeWebV2Resource) GetChannelURL(ctx context.Context, request YouTubeWebV2GetChannelURLRequest) (*YouTubeWebV2GetChannelURLResponse, error) {
	return r.client.YouTubeWebV2GetChannelURLFromChannelID(ctx, request)
}

// GetChannelVideos 获取频道视频 /Get channel videos
//
// GET /api/v1/youtube/web_v2/get_channel_videos
func (r YouTubeWebV2Resource) GetChannelVideos(ctx context.Context, request YouTubeWebV2GetChannelVideosRequest) (*YouTubeWebV2GetChannelVideosResponse, error) {
	return r.client.YouTubeWebV2GetChannelVideos(ctx, request)
}

// YouTubeWebV2GetVideoStreamsRequest is the request for GET /api/v1/youtube/web_v2/get_video_streams.
type YouTubeWebV2GetVideoStreamsRequest = YouTubeWebV2GetVideoStreamsInfoRequest

// YouTubeWebV2GetVideoStreamsResponse is the response for GET /api/v1/youtube/web_v2/get_video_streams.
type YouTubeWebV2GetVideoStreamsResponse = YouTubeWebV2GetVideoStreamsInfoResponse

// GetVideoStreams 获取视频流信息/Get video streams info
//
// GET /api/v1/youtube/web_v2/get_video_streams
func (r YouTubeWebV2Resource) GetVideoStreams(ctx context.Context, request YouTubeWebV2GetVideoStreamsRequest) (*YouTubeWebV2GetVideoStreamsResponse, error) {
	return r.client.YouTubeWebV2GetVideoStreamsInfo(ctx, request)
}

// YouTubeWebV2GetVideoStreamsV2Request is the request for GET /api/v1/youtube/web_v2/get_video_streams_v2.
type YouTubeWebV2GetVideoStreamsV2Request = YouTubeWebV2GetVideoStreamsInfoV2Request

// YouTubeWebV2GetVideoStreamsV2Response is the response for GET /api/v1/youtube/web_v2/get_video_streams_v2.
type YouTubeWebV2GetVideoStreamsV2Response = YouTubeWebV2GetVideoStreamsInfoV2Response

// GetVideoStreamsV2 获取视频流信息 V2/Get video streams info V2
//
// GET /api/v1/youtube/web_v2/get_video_streams_v2
func (r YouTubeWebV2Resource) GetVideoStreamsV2(ctx context.Context, request YouTubeWebV2GetVideoStreamsV2Request) (*YouTubeWebV2GetVideoStreamsV2Response, error) {
	return r.client.YouTubeWebV2GetVideoStreamsInfoV2(ctx, request)
}

// YouTubeWebV2GetSignedStreamURLRequest is the request for GET /api/v1/youtube/web_v2/get_signed_stream_url.
type YouTubeWebV2GetSignedStreamURLRequest = YouTubeWebV2GetSignedVideoStreamURLRequest

// YouTubeWebV2GetSignedStreamURLResponse is the response for GET /api/v1/youtube/web_v2/get_signed_stream_url.
type YouTubeWebV2GetSignedStreamURLResponse = YouTubeWebV2GetSignedVideoStreamURLResponse

// GetSignedStreamURL 获取已签名的视频流URL/Get signed video stream URL
//
// GET /api/v1/youtube/web_v2/get_signed_stream_url
func (r YouTubeWebV2Resource) GetSignedStreamURL(ctx context.Context, request YouTubeWebV2GetSignedStreamURLRequest) (*YouTubeWebV2GetSignedStreamURLResponse, error) {
	return r.client.YouTubeWebV2GetSignedVideoStreamURL(ctx, request)
}

// GetVideoCaptions 获取视频字幕/Get video captions
//
// GET /api/v1/youtube/web_v2/get_video_captions
func (r YouTubeWebV2Resource) GetVideoCaptions(ctx context.Context, request YouTubeWebV2GetVideoCaptionsRequest) (*YouTubeWebV2GetVideoCaptionsResponse, error) {
	return r.client.YouTubeWebV2GetVideoCaptions(ctx, request)
}

// GetVideoCaptionsV2 获取视频字幕 V2/Get video captions V2
//
// GET /api/v1/youtube/web_v2/get_video_captions_v2
func (r YouTubeWebV2Resource) GetVideoCaptionsV2(ctx context.Context, request YouTubeWebV2GetVideoCaptionsV2Request) (*YouTubeWebV2GetVideoCaptionsV2Response, error) {
	return r.client.YouTubeWebV2GetVideoCaptionsV2(ctx, request)
}

// GetRelatedVideos 获取视频相似内容/Get related videos
//
// GET /api/v1/youtube/web_v2/get_related_videos
func (r YouTubeWebV2Resource) GetRelatedVideos(ctx context.Context, request YouTubeWebV2GetRelatedVideosRequest) (*YouTubeWebV2GetRelatedVideosResponse, error) {
	return r.client.YouTubeWebV2GetRelatedVideos(ctx, request)
}

// GetChannelShorts 获取频道短视频列表/Get channel shorts
//
// GET /api/v1/youtube/web_v2/get_channel_shorts
func (r YouTubeWebV2Resource) GetChannelShorts(ctx context.Context, request YouTubeWebV2GetChannelShortsRequest) (*YouTubeWebV2GetChannelShortsResponse, error) {
	return r.client.YouTubeWebV2GetChannelShorts(ctx, request)
}

// GetSearchSuggestions 获取搜索推荐词/Get search suggestions
//
// GET /api/v1/youtube/web_v2/get_search_suggestions
func (r YouTubeWebV2Resource) GetSearchSuggestions(ctx context.Context, request YouTubeWebV2GetSearchSuggestionsRequest) (*YouTubeWebV2GetSearchSuggestionsResponse, error) {
	return r.client.YouTubeWebV2GetSearchSuggestions(ctx, request)
}

// SearchChannels 搜索频道/Search channels
//
// GET /api/v1/youtube/web_v2/search_channels
func (r YouTubeWebV2Resource) SearchChannels(ctx context.Context, request YouTubeWebV2SearchChannelsRequest) (*YouTubeWebV2SearchChannelsResponse, error) {
	return r.client.YouTubeWebV2SearchChannels(ctx, request)
}

// GetChannelCommunityPosts 获取频道帖子列表/Get channel community posts
//
// GET /api/v1/youtube/web_v2/get_channel_community_posts
func (r YouTubeWebV2Resource) GetChannelCommunityPosts(ctx context.Context, request YouTubeWebV2GetChannelCommunityPostsRequest) (*YouTubeWebV2GetChannelCommunityPostsResponse, error) {
	return r.client.YouTubeWebV2GetChannelCommunityPosts(ctx, request)
}

// GetPostDetail 获取帖子详情/Get post detail
//
// GET /api/v1/youtube/web_v2/get_post_detail
func (r YouTubeWebV2Resource) GetPostDetail(ctx context.Context, request YouTubeWebV2GetPostDetailRequest) (*YouTubeWebV2GetPostDetailResponse, error) {
	return r.client.YouTubeWebV2GetPostDetail(ctx, request)
}

// GetPostComments 获取帖子评论/Get post comments
//
// GET /api/v1/youtube/web_v2/get_post_comments
func (r YouTubeWebV2Resource) GetPostComments(ctx context.Context, request YouTubeWebV2GetPostCommentsRequest) (*YouTubeWebV2GetPostCommentsResponse, error) {
	return r.client.YouTubeWebV2GetPostComments(ctx, request)
}

// GetPostCommentReplies 获取帖子评论回复/Get post comment replies
//
// GET /api/v1/youtube/web_v2/get_post_comment_replies
func (r YouTubeWebV2Resource) GetPostCommentReplies(ctx context.Context, request YouTubeWebV2GetPostCommentRepliesRequest) (*YouTubeWebV2GetPostCommentRepliesResponse, error) {
	return r.client.YouTubeWebV2GetPostCommentReplies(ctx, request)
}

// LinkedInWebResource contains endpoints from the LinkedIn-Web-API tag.
type LinkedInWebResource struct {
	client *Client
}

// GetUserProfile 获取用户资料/Get user profile
//
// GET /api/v1/linkedin/web/get_user_profile
func (r LinkedInWebResource) GetUserProfile(ctx context.Context, request LinkedInWebGetUserProfileRequest) (*LinkedInWebGetUserProfileResponse, error) {
	return r.client.LinkedInWebGetUserProfile(ctx, request)
}

// GetUserPosts 获取用户帖子/Get user posts
//
// GET /api/v1/linkedin/web/get_user_posts
func (r LinkedInWebResource) GetUserPosts(ctx context.Context, request LinkedInWebGetUserPostsRequest) (*LinkedInWebGetUserPostsResponse, error) {
	return r.client.LinkedInWebGetUserPosts(ctx, request)
}

// GetUserComments 获取用户评论/Get user comments
//
// GET /api/v1/linkedin/web/get_user_comments
func (r LinkedInWebResource) GetUserComments(ctx context.Context, request LinkedInWebGetUserCommentsRequest) (*LinkedInWebGetUserCommentsResponse, error) {
	return r.client.LinkedInWebGetUserComments(ctx, request)
}

// LinkedInWebGetUserContactRequest is the request for GET /api/v1/linkedin/web/get_user_contact.
type LinkedInWebGetUserContactRequest = LinkedInWebGetUserContactInformationRequest

// LinkedInWebGetUserContactResponse is the response for GET /api/v1/linkedin/web/get_user_contact.
type LinkedInWebGetUserContactResponse = LinkedInWebGetUserContactInformationResponse

// GetUserContact 获取用户联系信息/Get user contact information
//
// GET /api/v1/linkedin/web/get_user_contact
func (r LinkedInWebResource) GetUserContact(ctx context.Context, request LinkedInWebGetUserContactRequest) (*LinkedInWebGetUserContactResponse, error) {
	return r.client.LinkedInWebGetUserContactInformation(ctx, request)
}

// GetUserRecommendations 获取用户推荐信/Get user recommendations
//
// GET /api/v1/linkedin/web/get_user_recommendations
func (r LinkedInWebResource) GetUserRecommendations(ctx context.Context, request LinkedInWebGetUserRecommendationsRequest) (*LinkedInWebGetUserRecommendationsResponse, error) {
	return r.client.LinkedInWebGetUserRecommendations(ctx, request)
}

// GetUserVideos 获取用户视频/Get user videos
//
// GET /api/v1/linkedin/web/get_user_videos
func (r LinkedInWebResource) GetUserVideos(ctx context.Context, request LinkedInWebGetUserVideosRequest) (*LinkedInWebGetUserVideosResponse, error) {
	return r.client.LinkedInWebGetUserVideos(ctx, request)
}

// GetUserImages 获取用户图片/Get user images
//
// GET /api/v1/linkedin/web/get_user_images
func (r LinkedInWebResource) GetUserImages(ctx context.Context, request LinkedInWebGetUserImagesRequest) (*LinkedInWebGetUserImagesResponse, error) {
	return r.client.LinkedInWebGetUserImages(ctx, request)
}

// GetCompanyProfile 获取公司资料/Get company profile
//
// GET /api/v1/linkedin/web/get_company_profile
func (r LinkedInWebResource) GetCompanyProfile(ctx context.Context, request LinkedInWebGetCompanyProfileRequest) (*LinkedInWebGetCompanyProfileResponse, error) {
	return r.client.LinkedInWebGetCompanyProfile(ctx, request)
}

// GetCompanyPeople 获取公司员工/Get company people
//
// GET /api/v1/linkedin/web/get_company_people
func (r LinkedInWebResource) GetCompanyPeople(ctx context.Context, request LinkedInWebGetCompanyPeopleRequest) (*LinkedInWebGetCompanyPeopleResponse, error) {
	return r.client.LinkedInWebGetCompanyPeople(ctx, request)
}

// GetCompanyPosts 获取公司帖子/Get company posts
//
// GET /api/v1/linkedin/web/get_company_posts
func (r LinkedInWebResource) GetCompanyPosts(ctx context.Context, request LinkedInWebGetCompanyPostsRequest) (*LinkedInWebGetCompanyPostsResponse, error) {
	return r.client.LinkedInWebGetCompanyPosts(ctx, request)
}

// GetCompanyJobs 获取公司职位/Get company jobs
//
// GET /api/v1/linkedin/web/get_company_jobs
func (r LinkedInWebResource) GetCompanyJobs(ctx context.Context, request LinkedInWebGetCompanyJobsRequest) (*LinkedInWebGetCompanyJobsResponse, error) {
	return r.client.LinkedInWebGetCompanyJobs(ctx, request)
}

// GetCompanyJobCount 获取公司职位数量/Get company job count
//
// GET /api/v1/linkedin/web/get_company_job_count
func (r LinkedInWebResource) GetCompanyJobCount(ctx context.Context, request LinkedInWebGetCompanyJobCountRequest) (*LinkedInWebGetCompanyJobCountResponse, error) {
	return r.client.LinkedInWebGetCompanyJobCount(ctx, request)
}

// GetUserAbout 获取用户简介/Get user about
//
// GET /api/v1/linkedin/web/get_user_about
func (r LinkedInWebResource) GetUserAbout(ctx context.Context, request LinkedInWebGetUserAboutRequest) (*LinkedInWebGetUserAboutResponse, error) {
	return r.client.LinkedInWebGetUserAbout(ctx, request)
}

// GetUserFollowerAndConnection 获取用户粉丝和连接数/Get user follower and connection
//
// GET /api/v1/linkedin/web/get_user_follower_and_connection
func (r LinkedInWebResource) GetUserFollowerAndConnection(ctx context.Context, request LinkedInWebGetUserFollowerAndConnectionRequest) (*LinkedInWebGetUserFollowerAndConnectionResponse, error) {
	return r.client.LinkedInWebGetUserFollowerAndConnection(ctx, request)
}

// GetUserExperience 获取用户工作经历/Get user experience
//
// GET /api/v1/linkedin/web/get_user_experience
func (r LinkedInWebResource) GetUserExperience(ctx context.Context, request LinkedInWebGetUserExperienceRequest) (*LinkedInWebGetUserExperienceResponse, error) {
	return r.client.LinkedInWebGetUserExperience(ctx, request)
}

// GetUserSkills 获取用户技能/Get user skills
//
// GET /api/v1/linkedin/web/get_user_skills
func (r LinkedInWebResource) GetUserSkills(ctx context.Context, request LinkedInWebGetUserSkillsRequest) (*LinkedInWebGetUserSkillsResponse, error) {
	return r.client.LinkedInWebGetUserSkills(ctx, request)
}

// GetUserEducations 获取用户教育背景/Get user educations
//
// GET /api/v1/linkedin/web/get_user_educations
func (r LinkedInWebResource) GetUserEducations(ctx context.Context, request LinkedInWebGetUserEducationsRequest) (*LinkedInWebGetUserEducationsResponse, error) {
	return r.client.LinkedInWebGetUserEducations(ctx, request)
}

// GetUserPublications 获取用户出版物/Get user publications
//
// GET /api/v1/linkedin/web/get_user_publications
func (r LinkedInWebResource) GetUserPublications(ctx context.Context, request LinkedInWebGetUserPublicationsRequest) (*LinkedInWebGetUserPublicationsResponse, error) {
	return r.client.LinkedInWebGetUserPublications(ctx, request)
}

// GetUserCertifications 获取用户认证/Get user certifications
//
// GET /api/v1/linkedin/web/get_user_certifications
func (r LinkedInWebResource) GetUserCertifications(ctx context.Context, request LinkedInWebGetUserCertificationsRequest) (*LinkedInWebGetUserCertificationsResponse, error) {
	return r.client.LinkedInWebGetUserCertifications(ctx, request)
}

// GetUserHonors 获取用户荣誉奖项/Get user honors
//
// GET /api/v1/linkedin/web/get_user_honors
func (r LinkedInWebResource) GetUserHonors(ctx context.Context, request LinkedInWebGetUserHonorsRequest) (*LinkedInWebGetUserHonorsResponse, error) {
	return r.client.LinkedInWebGetUserHonors(ctx, request)
}

// GetUserInterestsGroups 获取用户感兴趣的群组/Get user interests groups
//
// GET /api/v1/linkedin/web/get_user_interests_groups
func (r LinkedInWebResource) GetUserInterestsGroups(ctx context.Context, request LinkedInWebGetUserInterestsGroupsRequest) (*LinkedInWebGetUserInterestsGroupsResponse, error) {
	return r.client.LinkedInWebGetUserInterestsGroups(ctx, request)
}

// GetUserInterestsCompanies 获取用户感兴趣的公司/Get user interests companies
//
// GET /api/v1/linkedin/web/get_user_interests_companies
func (r LinkedInWebResource) GetUserInterestsCompanies(ctx context.Context, request LinkedInWebGetUserInterestsCompaniesRequest) (*LinkedInWebGetUserInterestsCompaniesResponse, error) {
	return r.client.LinkedInWebGetUserInterestsCompanies(ctx, request)
}

// GetJobDetail 获取职位详情/Get job detail
//
// GET /api/v1/linkedin/web/get_job_detail
func (r LinkedInWebResource) GetJobDetail(ctx context.Context, request LinkedInWebGetJobDetailRequest) (*LinkedInWebGetJobDetailResponse, error) {
	return r.client.LinkedInWebGetJobDetail(ctx, request)
}

// SearchJobs 搜索职位/Search jobs
//
// GET /api/v1/linkedin/web/search_jobs
func (r LinkedInWebResource) SearchJobs(ctx context.Context, request LinkedInWebSearchJobsRequest) (*LinkedInWebSearchJobsResponse, error) {
	return r.client.LinkedInWebSearchJobs(ctx, request)
}

// SearchPeople 搜索用户/Search people
//
// GET /api/v1/linkedin/web/search_people
func (r LinkedInWebResource) SearchPeople(ctx context.Context, request LinkedInWebSearchPeopleRequest) (*LinkedInWebSearchPeopleResponse, error) {
	return r.client.LinkedInWebSearchPeople(ctx, request)
}

// GetUserReactions 获取用户点赞反应/Get user reactions
//
// GET /api/v1/linkedin/web/get_user_reactions
func (r LinkedInWebResource) GetUserReactions(ctx context.Context, request LinkedInWebGetUserReactionsRequest) (*LinkedInWebGetUserReactionsResponse, error) {
	return r.client.LinkedInWebGetUserReactions(ctx, request)
}

// GetUserVolunteers 获取用户志愿者经历/Get user volunteers
//
// GET /api/v1/linkedin/web/get_user_volunteers
func (r LinkedInWebResource) GetUserVolunteers(ctx context.Context, request LinkedInWebGetUserVolunteersRequest) (*LinkedInWebGetUserVolunteersResponse, error) {
	return r.client.LinkedInWebGetUserVolunteers(ctx, request)
}

// GetCompanyAffiliatedPages 获取公司关联页面/Get company affiliated pages
//
// GET /api/v1/linkedin/web/get_company_affiliated_pages
func (r LinkedInWebResource) GetCompanyAffiliatedPages(ctx context.Context, request LinkedInWebGetCompanyAffiliatedPagesRequest) (*LinkedInWebGetCompanyAffiliatedPagesResponse, error) {
	return r.client.LinkedInWebGetCompanyAffiliatedPages(ctx, request)
}

// GetCompanyAssociatedMemberInsights 获取公司关联成员洞察/Get company associated member insights
//
// GET /api/v1/linkedin/web/get_company_associated_member_insights
func (r LinkedInWebResource) GetCompanyAssociatedMemberInsights(ctx context.Context, request LinkedInWebGetCompanyAssociatedMemberInsightsRequest) (*LinkedInWebGetCompanyAssociatedMemberInsightsResponse, error) {
	return r.client.LinkedInWebGetCompanyAssociatedMemberInsights(ctx, request)
}

// GetPostDetail 获取帖子详情/Get post detail
//
// GET /api/v1/linkedin/web/get_post_detail
func (r LinkedInWebResource) GetPostDetail(ctx context.Context, request LinkedInWebGetPostDetailRequest) (*LinkedInWebGetPostDetailResponse, error) {
	return r.client.LinkedInWebGetPostDetail(ctx, request)
}

// GetPostComments 获取帖子评论/Get post comments
//
// GET /api/v1/linkedin/web/get_post_comments
func (r LinkedInWebResource) GetPostComments(ctx context.Context, request LinkedInWebGetPostCommentsRequest) (*LinkedInWebGetPostCommentsResponse, error) {
	return r.client.LinkedInWebGetPostComments(ctx, request)
}

// GetPostReactions 获取帖子点赞反应/Get post reactions
//
// GET /api/v1/linkedin/web/get_post_reactions
func (r LinkedInWebResource) GetPostReactions(ctx context.Context, request LinkedInWebGetPostReactionsRequest) (*LinkedInWebGetPostReactionsResponse, error) {
	return r.client.LinkedInWebGetPostReactions(ctx, request)
}

// GetPostReposts 获取帖子转发/Get post reposts
//
// GET /api/v1/linkedin/web/get_post_reposts
func (r LinkedInWebResource) GetPostReposts(ctx context.Context, request LinkedInWebGetPostRepostsRequest) (*LinkedInWebGetPostRepostsResponse, error) {
	return r.client.LinkedInWebGetPostReposts(ctx, request)
}

// LinkedInWebGetCommentsRepliesRequest is the request for GET /api/v1/linkedin/web/get_comments_replies.
type LinkedInWebGetCommentsRepliesRequest = LinkedInWebGetCommentRepliesRequest

// LinkedInWebGetCommentsRepliesResponse is the response for GET /api/v1/linkedin/web/get_comments_replies.
type LinkedInWebGetCommentsRepliesResponse = LinkedInWebGetCommentRepliesResponse

// GetCommentsReplies 获取评论回复/Get comment replies
//
// GET /api/v1/linkedin/web/get_comments_replies
func (r LinkedInWebResource) GetCommentsReplies(ctx context.Context, request LinkedInWebGetCommentsRepliesRequest) (*LinkedInWebGetCommentsRepliesResponse, error) {
	return r.client.LinkedInWebGetCommentReplies(ctx, request)
}

// SearchPosts 搜索帖子/Search posts
//
// GET /api/v1/linkedin/web/search_posts
func (r LinkedInWebResource) SearchPosts(ctx context.Context, request LinkedInWebSearchPostsRequest) (*LinkedInWebSearchPostsResponse, error) {
	return r.client.LinkedInWebSearchPosts(ctx, request)
}

// SearchLocation 搜索地理位置/Search location
//
// GET /api/v1/linkedin/web/search_location
func (r LinkedInWebResource) SearchLocation(ctx context.Context, request LinkedInWebSearchLocationRequest) (*LinkedInWebSearchLocationResponse, error) {
	return r.client.LinkedInWebSearchLocation(ctx, request)
}

// SearchSchools 搜索学校/Search schools
//
// GET /api/v1/linkedin/web/search_schools
func (r LinkedInWebResource) SearchSchools(ctx context.Context, request LinkedInWebSearchSchoolsRequest) (*LinkedInWebSearchSchoolsResponse, error) {
	return r.client.LinkedInWebSearchSchools(ctx, request)
}

// LinkedInWebSearchSuggestionIndustryRequest is the request for GET /api/v1/linkedin/web/search_suggestion_industry.
type LinkedInWebSearchSuggestionIndustryRequest = LinkedInWebSearchIndustrySuggestionsRequest

// LinkedInWebSearchSuggestionIndustryResponse is the response for GET /api/v1/linkedin/web/search_suggestion_industry.
type LinkedInWebSearchSuggestionIndustryResponse = LinkedInWebSearchIndustrySuggestionsResponse

// SearchSuggestionIndustry 搜索行业建议/Search industry suggestions
//
// GET /api/v1/linkedin/web/search_suggestion_industry
func (r LinkedInWebResource) SearchSuggestionIndustry(ctx context.Context, request LinkedInWebSearchSuggestionIndustryRequest) (*LinkedInWebSearchSuggestionIndustryResponse, error) {
	return r.client.LinkedInWebSearchIndustrySuggestions(ctx, request)
}

// GetGroupInfo 获取群组信息/Get group info
//
// GET /api/v1/linkedin/web/get_group_info
func (r LinkedInWebResource) GetGroupInfo(ctx context.Context, request LinkedInWebGetGroupInfoRequest) (*LinkedInWebGetGroupInfoResponse, error) {
	return r.client.LinkedInWebGetGroupInfo(ctx, request)
}

// GetGroupPosts 获取群组帖子/Get group posts
//
// GET /api/v1/linkedin/web/get_group_posts
func (r LinkedInWebResource) GetGroupPosts(ctx context.Context, request LinkedInWebGetGroupPostsRequest) (*LinkedInWebGetGroupPostsResponse, error) {
	return r.client.LinkedInWebGetGroupPosts(ctx, request)
}

// SearchAds 搜索广告/Search ads (Ad Library)
//
// GET /api/v1/linkedin/web/search_ads
func (r LinkedInWebResource) SearchAds(ctx context.Context, request LinkedInWebSearchAdsRequest) (*LinkedInWebSearchAdsResponse, error) {
	return r.client.LinkedInWebSearchAds(ctx, request)
}

// GetAdDetail 获取广告详情/Get ad detail
//
// GET /api/v1/linkedin/web/get_ad_detail
func (r LinkedInWebResource) GetAdDetail(ctx context.Context, request LinkedInWebGetAdDetailRequest) (*LinkedInWebGetAdDetailResponse, error) {
	return r.client.LinkedInWebGetAdDetail(ctx, request)
}

// LinkedInWebV2Resource contains endpoints from the LinkedIn-Web-V2-API tag.
type LinkedInWebV2Resource struct {
	client *Client
}

// GetUserProfile 获取用户主页基础信息（可选附带子节）/Get user profile (optional sub-sections)
//
// GET /api/v1/linkedin/web_v2/get_user_profile
func (r LinkedInWebV2Resource) GetUserProfile(ctx context.Context, request LinkedInWebV2GetUserProfileRequest) (*LinkedInWebV2GetUserProfileResponse, error) {
	return r.client.LinkedInWebV2GetUserProfile(ctx, request)
}

// GetUserPosts 获取用户帖子（动态标签）/Get user posts
//
// GET /api/v1/linkedin/web_v2/get_user_posts
func (r LinkedInWebV2Resource) GetUserPosts(ctx context.Context, request LinkedInWebV2GetUserPostsRequest) (*LinkedInWebV2GetUserPostsResponse, error) {
	return r.client.LinkedInWebV2GetUserPosts(ctx, request)
}

// GetUserComments 获取用户评论（在他人帖子下的评论）/Get user comments
//
// GET /api/v1/linkedin/web_v2/get_user_comments
func (r LinkedInWebV2Resource) GetUserComments(ctx context.Context, request LinkedInWebV2GetUserCommentsRequest) (*LinkedInWebV2GetUserCommentsResponse, error) {
	return r.client.LinkedInWebV2GetUserComments(ctx, request)
}

// LinkedInWebV2GetUserContactInfoRequest is the request for GET /api/v1/linkedin/web_v2/get_user_contact_info.
type LinkedInWebV2GetUserContactInfoRequest = LinkedInWebV2GetContactInfoRequest

// LinkedInWebV2GetUserContactInfoResponse is the response for GET /api/v1/linkedin/web_v2/get_user_contact_info.
type LinkedInWebV2GetUserContactInfoResponse = LinkedInWebV2GetContactInfoResponse

// GetUserContactInfo 获取用户公开联系信息/Get contact info
//
// GET /api/v1/linkedin/web_v2/get_user_contact_info
func (r LinkedInWebV2Resource) GetUserContactInfo(ctx context.Context, request LinkedInWebV2GetUserContactInfoRequest) (*LinkedInWebV2GetUserContactInfoResponse, error) {
	return r.client.LinkedInWebV2GetContactInfo(ctx, request)
}

// LinkedInWebV2GetUserRecommendationsRequest is the request for GET /api/v1/linkedin/web_v2/get_user_recommendations.
type LinkedInWebV2GetUserRecommendationsRequest = LinkedInWebV2GetRecommendationsRequest

// LinkedInWebV2GetUserRecommendationsResponse is the response for GET /api/v1/linkedin/web_v2/get_user_recommendations.
type LinkedInWebV2GetUserRecommendationsResponse = LinkedInWebV2GetRecommendationsResponse

// GetUserRecommendations 获取用户推荐信/Get recommendations
//
// GET /api/v1/linkedin/web_v2/get_user_recommendations
func (r LinkedInWebV2Resource) GetUserRecommendations(ctx context.Context, request LinkedInWebV2GetUserRecommendationsRequest) (*LinkedInWebV2GetUserRecommendationsResponse, error) {
	return r.client.LinkedInWebV2GetRecommendations(ctx, request)
}

// GetUserVideos 获取用户视频帖子/Get user videos
//
// GET /api/v1/linkedin/web_v2/get_user_videos
func (r LinkedInWebV2Resource) GetUserVideos(ctx context.Context, request LinkedInWebV2GetUserVideosRequest) (*LinkedInWebV2GetUserVideosResponse, error) {
	return r.client.LinkedInWebV2GetUserVideos(ctx, request)
}

// GetUserImages 获取用户图片帖子/Get user images
//
// GET /api/v1/linkedin/web_v2/get_user_images
func (r LinkedInWebV2Resource) GetUserImages(ctx context.Context, request LinkedInWebV2GetUserImagesRequest) (*LinkedInWebV2GetUserImagesResponse, error) {
	return r.client.LinkedInWebV2GetUserImages(ctx, request)
}

// GetUserBio 获取用户简介摘要/Get user bio
//
// GET /api/v1/linkedin/web_v2/get_user_bio
func (r LinkedInWebV2Resource) GetUserBio(ctx context.Context, request LinkedInWebV2GetUserBioRequest) (*LinkedInWebV2GetUserBioResponse, error) {
	return r.client.LinkedInWebV2GetUserBio(ctx, request)
}

// LinkedInWebV2GetUserFollowerAndConnectionCountRequest is the request for GET /api/v1/linkedin/web_v2/get_user_follower_and_connection_count.
type LinkedInWebV2GetUserFollowerAndConnectionCountRequest = LinkedInWebV2GetFollowerConnectionCountRequest

// LinkedInWebV2GetUserFollowerAndConnectionCountResponse is the response for GET /api/v1/linkedin/web_v2/get_user_follower_and_connection_count.
type LinkedInWebV2GetUserFollowerAndConnectionCountResponse = LinkedInWebV2GetFollowerConnectionCountResponse

// GetUserFollowerAndConnectionCount 获取用户粉丝/连接数/Get follower & connection count
//
// GET /api/v1/linkedin/web_v2/get_user_follower_and_connection_count
func (r LinkedInWebV2Resource) GetUserFollowerAndConnectionCount(ctx context.Context, request LinkedInWebV2GetUserFollowerAndConnectionCountRequest) (*LinkedInWebV2GetUserFollowerAndConnectionCountResponse, error) {
	return r.client.LinkedInWebV2GetFollowerConnectionCount(ctx, request)
}

// LinkedInWebV2GetUserProfileCardsRequest is the request for GET /api/v1/linkedin/web_v2/get_user_profile_cards.
type LinkedInWebV2GetUserProfileCardsRequest = LinkedInWebV2GetFullProfileCardsRequest

// LinkedInWebV2GetUserProfileCardsResponse is the response for GET /api/v1/linkedin/web_v2/get_user_profile_cards.
type LinkedInWebV2GetUserProfileCardsResponse = LinkedInWebV2GetFullProfileCardsResponse

// GetUserProfileCards 获取用户主页全部卡片原始结构/Get full profile cards
//
// GET /api/v1/linkedin/web_v2/get_user_profile_cards
func (r LinkedInWebV2Resource) GetUserProfileCards(ctx context.Context, request LinkedInWebV2GetUserProfileCardsRequest) (*LinkedInWebV2GetUserProfileCardsResponse, error) {
	return r.client.LinkedInWebV2GetFullProfileCards(ctx, request)
}

// LinkedInWebV2GetUserExperiencesRequest is the request for GET /api/v1/linkedin/web_v2/get_user_experiences.
type LinkedInWebV2GetUserExperiencesRequest = LinkedInWebV2GetExperiencesRequest

// LinkedInWebV2GetUserExperiencesResponse is the response for GET /api/v1/linkedin/web_v2/get_user_experiences.
type LinkedInWebV2GetUserExperiencesResponse = LinkedInWebV2GetExperiencesResponse

// GetUserExperiences 获取用户工作经历/Get experiences
//
// GET /api/v1/linkedin/web_v2/get_user_experiences
func (r LinkedInWebV2Resource) GetUserExperiences(ctx context.Context, request LinkedInWebV2GetUserExperiencesRequest) (*LinkedInWebV2GetUserExperiencesResponse, error) {
	return r.client.LinkedInWebV2GetExperiences(ctx, request)
}

// LinkedInWebV2GetUserSkillsRequest is the request for GET /api/v1/linkedin/web_v2/get_user_skills.
type LinkedInWebV2GetUserSkillsRequest = LinkedInWebV2GetSkillsRequest

// LinkedInWebV2GetUserSkillsResponse is the response for GET /api/v1/linkedin/web_v2/get_user_skills.
type LinkedInWebV2GetUserSkillsResponse = LinkedInWebV2GetSkillsResponse

// GetUserSkills 获取用户技能/Get skills
//
// GET /api/v1/linkedin/web_v2/get_user_skills
func (r LinkedInWebV2Resource) GetUserSkills(ctx context.Context, request LinkedInWebV2GetUserSkillsRequest) (*LinkedInWebV2GetUserSkillsResponse, error) {
	return r.client.LinkedInWebV2GetSkills(ctx, request)
}

// LinkedInWebV2GetUserEducationsRequest is the request for GET /api/v1/linkedin/web_v2/get_user_educations.
type LinkedInWebV2GetUserEducationsRequest = LinkedInWebV2GetEducationsRequest

// LinkedInWebV2GetUserEducationsResponse is the response for GET /api/v1/linkedin/web_v2/get_user_educations.
type LinkedInWebV2GetUserEducationsResponse = LinkedInWebV2GetEducationsResponse

// GetUserEducations 获取用户教育背景/Get educations
//
// GET /api/v1/linkedin/web_v2/get_user_educations
func (r LinkedInWebV2Resource) GetUserEducations(ctx context.Context, request LinkedInWebV2GetUserEducationsRequest) (*LinkedInWebV2GetUserEducationsResponse, error) {
	return r.client.LinkedInWebV2GetEducations(ctx, request)
}

// LinkedInWebV2GetUserPublicationsRequest is the request for GET /api/v1/linkedin/web_v2/get_user_publications.
type LinkedInWebV2GetUserPublicationsRequest = LinkedInWebV2GetPublicationsRequest

// LinkedInWebV2GetUserPublicationsResponse is the response for GET /api/v1/linkedin/web_v2/get_user_publications.
type LinkedInWebV2GetUserPublicationsResponse = LinkedInWebV2GetPublicationsResponse

// GetUserPublications 获取用户出版物/Get publications
//
// GET /api/v1/linkedin/web_v2/get_user_publications
func (r LinkedInWebV2Resource) GetUserPublications(ctx context.Context, request LinkedInWebV2GetUserPublicationsRequest) (*LinkedInWebV2GetUserPublicationsResponse, error) {
	return r.client.LinkedInWebV2GetPublications(ctx, request)
}

// LinkedInWebV2GetUserCertificationsRequest is the request for GET /api/v1/linkedin/web_v2/get_user_certifications.
type LinkedInWebV2GetUserCertificationsRequest = LinkedInWebV2GetCertificationsRequest

// LinkedInWebV2GetUserCertificationsResponse is the response for GET /api/v1/linkedin/web_v2/get_user_certifications.
type LinkedInWebV2GetUserCertificationsResponse = LinkedInWebV2GetCertificationsResponse

// GetUserCertifications 获取用户认证/Get certifications
//
// GET /api/v1/linkedin/web_v2/get_user_certifications
func (r LinkedInWebV2Resource) GetUserCertifications(ctx context.Context, request LinkedInWebV2GetUserCertificationsRequest) (*LinkedInWebV2GetUserCertificationsResponse, error) {
	return r.client.LinkedInWebV2GetCertifications(ctx, request)
}

// LinkedInWebV2GetUserHonorsRequest is the request for GET /api/v1/linkedin/web_v2/get_user_honors.
type LinkedInWebV2GetUserHonorsRequest = LinkedInWebV2GetHonorsRequest

// LinkedInWebV2GetUserHonorsResponse is the response for GET /api/v1/linkedin/web_v2/get_user_honors.
type LinkedInWebV2GetUserHonorsResponse = LinkedInWebV2GetHonorsResponse

// GetUserHonors 获取用户荣誉奖项/Get honors
//
// GET /api/v1/linkedin/web_v2/get_user_honors
func (r LinkedInWebV2Resource) GetUserHonors(ctx context.Context, request LinkedInWebV2GetUserHonorsRequest) (*LinkedInWebV2GetUserHonorsResponse, error) {
	return r.client.LinkedInWebV2GetHonors(ctx, request)
}

// LinkedInWebV2GetUserInterestedGroupsRequest is the request for GET /api/v1/linkedin/web_v2/get_user_interested_groups.
type LinkedInWebV2GetUserInterestedGroupsRequest = LinkedInWebV2GetFollowedGroupsRequest

// LinkedInWebV2GetUserInterestedGroupsResponse is the response for GET /api/v1/linkedin/web_v2/get_user_interested_groups.
type LinkedInWebV2GetUserInterestedGroupsResponse = LinkedInWebV2GetFollowedGroupsResponse

// GetUserInterestedGroups 获取用户关注的群组/Get followed groups
//
// GET /api/v1/linkedin/web_v2/get_user_interested_groups
func (r LinkedInWebV2Resource) GetUserInterestedGroups(ctx context.Context, request LinkedInWebV2GetUserInterestedGroupsRequest) (*LinkedInWebV2GetUserInterestedGroupsResponse, error) {
	return r.client.LinkedInWebV2GetFollowedGroups(ctx, request)
}

// LinkedInWebV2GetUserInterestedCompaniesRequest is the request for GET /api/v1/linkedin/web_v2/get_user_interested_companies.
type LinkedInWebV2GetUserInterestedCompaniesRequest = LinkedInWebV2GetFollowedCompaniesRequest

// LinkedInWebV2GetUserInterestedCompaniesResponse is the response for GET /api/v1/linkedin/web_v2/get_user_interested_companies.
type LinkedInWebV2GetUserInterestedCompaniesResponse = LinkedInWebV2GetFollowedCompaniesResponse

// GetUserInterestedCompanies 获取用户关注的公司/Get followed companies
//
// GET /api/v1/linkedin/web_v2/get_user_interested_companies
func (r LinkedInWebV2Resource) GetUserInterestedCompanies(ctx context.Context, request LinkedInWebV2GetUserInterestedCompaniesRequest) (*LinkedInWebV2GetUserInterestedCompaniesResponse, error) {
	return r.client.LinkedInWebV2GetFollowedCompanies(ctx, request)
}

// LinkedInWebV2GetUserTopCardRequest is the request for GET /api/v1/linkedin/web_v2/get_user_top_card.
type LinkedInWebV2GetUserTopCardRequest = LinkedInWebV2GetProfileTopCardRequest

// LinkedInWebV2GetUserTopCardResponse is the response for GET /api/v1/linkedin/web_v2/get_user_top_card.
type LinkedInWebV2GetUserTopCardResponse = LinkedInWebV2GetProfileTopCardResponse

// GetUserTopCard 获取用户主页顶部卡片/Get profile top card
//
// GET /api/v1/linkedin/web_v2/get_user_top_card
func (r LinkedInWebV2Resource) GetUserTopCard(ctx context.Context, request LinkedInWebV2GetUserTopCardRequest) (*LinkedInWebV2GetUserTopCardResponse, error) {
	return r.client.LinkedInWebV2GetProfileTopCard(ctx, request)
}

// LinkedInWebV2GetUserTopCardSupplementaryRequest is the request for GET /api/v1/linkedin/web_v2/get_user_top_card_supplementary.
type LinkedInWebV2GetUserTopCardSupplementaryRequest = LinkedInWebV2GetTopCardSupplementaryRequest

// LinkedInWebV2GetUserTopCardSupplementaryResponse is the response for GET /api/v1/linkedin/web_v2/get_user_top_card_supplementary.
type LinkedInWebV2GetUserTopCardSupplementaryResponse = LinkedInWebV2GetTopCardSupplementaryResponse

// GetUserTopCardSupplementary 获取用户主页顶部卡片补充信息/Get top card supplementary
//
// GET /api/v1/linkedin/web_v2/get_user_top_card_supplementary
func (r LinkedInWebV2Resource) GetUserTopCardSupplementary(ctx context.Context, request LinkedInWebV2GetUserTopCardSupplementaryRequest) (*LinkedInWebV2GetUserTopCardSupplementaryResponse, error) {
	return r.client.LinkedInWebV2GetTopCardSupplementary(ctx, request)
}

// LinkedInWebV2GetUserRecentActivityRequest is the request for GET /api/v1/linkedin/web_v2/get_user_recent_activity.
type LinkedInWebV2GetUserRecentActivityRequest = LinkedInWebV2GetRecentActivitySummaryRequest

// LinkedInWebV2GetUserRecentActivityResponse is the response for GET /api/v1/linkedin/web_v2/get_user_recent_activity.
type LinkedInWebV2GetUserRecentActivityResponse = LinkedInWebV2GetRecentActivitySummaryResponse

// GetUserRecentActivity 获取用户近期动态聚合/Get recent activity summary
//
// GET /api/v1/linkedin/web_v2/get_user_recent_activity
func (r LinkedInWebV2Resource) GetUserRecentActivity(ctx context.Context, request LinkedInWebV2GetUserRecentActivityRequest) (*LinkedInWebV2GetUserRecentActivityResponse, error) {
	return r.client.LinkedInWebV2GetRecentActivitySummary(ctx, request)
}

// LinkedInWebV2GetDiscoveryRelevantToCompanyRequest is the request for GET /api/v1/linkedin/web_v2/get_discovery_relevant_to_company.
type LinkedInWebV2GetDiscoveryRelevantToCompanyRequest = LinkedInWebV2DiscoveryRelevantToCompanyRequest

// LinkedInWebV2GetDiscoveryRelevantToCompanyResponse is the response for GET /api/v1/linkedin/web_v2/get_discovery_relevant_to_company.
type LinkedInWebV2GetDiscoveryRelevantToCompanyResponse = LinkedInWebV2DiscoveryRelevantToCompanyResponse

// GetDiscoveryRelevantToCompany 发现："基于公司 X"的相关推荐/Discovery relevant to company
//
// GET /api/v1/linkedin/web_v2/get_discovery_relevant_to_company
func (r LinkedInWebV2Resource) GetDiscoveryRelevantToCompany(ctx context.Context, request LinkedInWebV2GetDiscoveryRelevantToCompanyRequest) (*LinkedInWebV2GetDiscoveryRelevantToCompanyResponse, error) {
	return r.client.LinkedInWebV2DiscoveryRelevantToCompany(ctx, request)
}

// LinkedInWebV2GetDiscoveryRelevantToUserRequest is the request for GET /api/v1/linkedin/web_v2/get_discovery_relevant_to_user.
type LinkedInWebV2GetDiscoveryRelevantToUserRequest = LinkedInWebV2DiscoveryRelevantToUserRequest

// LinkedInWebV2GetDiscoveryRelevantToUserResponse is the response for GET /api/v1/linkedin/web_v2/get_discovery_relevant_to_user.
type LinkedInWebV2GetDiscoveryRelevantToUserResponse = LinkedInWebV2DiscoveryRelevantToUserResponse

// GetDiscoveryRelevantToUser 发现："基于用户 X"的相关推荐/Discovery relevant to user
//
// GET /api/v1/linkedin/web_v2/get_discovery_relevant_to_user
func (r LinkedInWebV2Resource) GetDiscoveryRelevantToUser(ctx context.Context, request LinkedInWebV2GetDiscoveryRelevantToUserRequest) (*LinkedInWebV2GetDiscoveryRelevantToUserResponse, error) {
	return r.client.LinkedInWebV2DiscoveryRelevantToUser(ctx, request)
}

// GetCompanyProfile 获取公司主页资料/Get company profile
//
// GET /api/v1/linkedin/web_v2/get_company_profile
func (r LinkedInWebV2Resource) GetCompanyProfile(ctx context.Context, request LinkedInWebV2GetCompanyProfileRequest) (*LinkedInWebV2GetCompanyProfileResponse, error) {
	return r.client.LinkedInWebV2GetCompanyProfile(ctx, request)
}

// LinkedInWebV2GetCompanyEmployeesRequest is the request for GET /api/v1/linkedin/web_v2/get_company_employees.
type LinkedInWebV2GetCompanyEmployeesRequest = LinkedInWebV2GetEmployeesRequest

// LinkedInWebV2GetCompanyEmployeesResponse is the response for GET /api/v1/linkedin/web_v2/get_company_employees.
type LinkedInWebV2GetCompanyEmployeesResponse = LinkedInWebV2GetEmployeesResponse

// GetCompanyEmployees 获取公司员工列表/Get employees
//
// GET /api/v1/linkedin/web_v2/get_company_employees
func (r LinkedInWebV2Resource) GetCompanyEmployees(ctx context.Context, request LinkedInWebV2GetCompanyEmployeesRequest) (*LinkedInWebV2GetCompanyEmployeesResponse, error) {
	return r.client.LinkedInWebV2GetEmployees(ctx, request)
}

// GetCompanyPosts 获取公司主页帖子流/Get company posts
//
// GET /api/v1/linkedin/web_v2/get_company_posts
func (r LinkedInWebV2Resource) GetCompanyPosts(ctx context.Context, request LinkedInWebV2GetCompanyPostsRequest) (*LinkedInWebV2GetCompanyPostsResponse, error) {
	return r.client.LinkedInWebV2GetCompanyPosts(ctx, request)
}

// GetCompanyJobs 获取公司在招职位列表/Get company jobs
//
// GET /api/v1/linkedin/web_v2/get_company_jobs
func (r LinkedInWebV2Resource) GetCompanyJobs(ctx context.Context, request LinkedInWebV2GetCompanyJobsRequest) (*LinkedInWebV2GetCompanyJobsResponse, error) {
	return r.client.LinkedInWebV2GetCompanyJobs(ctx, request)
}

// LinkedInWebV2GetCompanyJobCountRequest is the request for GET /api/v1/linkedin/web_v2/get_company_job_count.
type LinkedInWebV2GetCompanyJobCountRequest = LinkedInWebV2GetJobCountRequest

// LinkedInWebV2GetCompanyJobCountResponse is the response for GET /api/v1/linkedin/web_v2/get_company_job_count.
type LinkedInWebV2GetCompanyJobCountResponse = LinkedInWebV2GetJobCountResponse

// GetCompanyJobCount 获取公司在招职位总数/Get job count
//
// GET /api/v1/linkedin/web_v2/get_company_job_count
func (r LinkedInWebV2Resource) GetCompanyJobCount(ctx context.Context, request LinkedInWebV2GetCompanyJobCountRequest) (*LinkedInWebV2GetCompanyJobCountResponse, error) {
	return r.client.LinkedInWebV2GetJobCount(ctx, request)
}

// LinkedInWebV2GetCompanySimilarCompaniesRequest is the request for GET /api/v1/linkedin/web_v2/get_company_similar_companies.
type LinkedInWebV2GetCompanySimilarCompaniesRequest = LinkedInWebV2GetSimilarCompaniesRequest

// LinkedInWebV2GetCompanySimilarCompaniesResponse is the response for GET /api/v1/linkedin/web_v2/get_company_similar_companies.
type LinkedInWebV2GetCompanySimilarCompaniesResponse = LinkedInWebV2GetSimilarCompaniesResponse

// GetCompanySimilarCompanies 获取相似公司（People also viewed）/Get similar companies
//
// GET /api/v1/linkedin/web_v2/get_company_similar_companies
func (r LinkedInWebV2Resource) GetCompanySimilarCompanies(ctx context.Context, request LinkedInWebV2GetCompanySimilarCompaniesRequest) (*LinkedInWebV2GetCompanySimilarCompaniesResponse, error) {
	return r.client.LinkedInWebV2GetSimilarCompanies(ctx, request)
}

// LinkedInWebV2GetCompanyCompetitorsRequest is the request for GET /api/v1/linkedin/web_v2/get_company_competitors.
type LinkedInWebV2GetCompanyCompetitorsRequest = LinkedInWebV2GetCompetitorsRequest

// LinkedInWebV2GetCompanyCompetitorsResponse is the response for GET /api/v1/linkedin/web_v2/get_company_competitors.
type LinkedInWebV2GetCompanyCompetitorsResponse = LinkedInWebV2GetCompetitorsResponse

// GetCompanyCompetitors 获取公司竞争对手/Get competitors
//
// GET /api/v1/linkedin/web_v2/get_company_competitors
func (r LinkedInWebV2Resource) GetCompanyCompetitors(ctx context.Context, request LinkedInWebV2GetCompanyCompetitorsRequest) (*LinkedInWebV2GetCompanyCompetitorsResponse, error) {
	return r.client.LinkedInWebV2GetCompetitors(ctx, request)
}

// LinkedInWebV2GetCompanyStockQuoteRequest is the request for GET /api/v1/linkedin/web_v2/get_company_stock_quote.
type LinkedInWebV2GetCompanyStockQuoteRequest = LinkedInWebV2GetStockQuoteRequest

// LinkedInWebV2GetCompanyStockQuoteResponse is the response for GET /api/v1/linkedin/web_v2/get_company_stock_quote.
type LinkedInWebV2GetCompanyStockQuoteResponse = LinkedInWebV2GetStockQuoteResponse

// GetCompanyStockQuote 获取上市公司股价/Get stock quote
//
// GET /api/v1/linkedin/web_v2/get_company_stock_quote
func (r LinkedInWebV2Resource) GetCompanyStockQuote(ctx context.Context, request LinkedInWebV2GetCompanyStockQuoteRequest) (*LinkedInWebV2GetCompanyStockQuoteResponse, error) {
	return r.client.LinkedInWebV2GetStockQuote(ctx, request)
}

// LinkedInWebV2GetCompanyCallToActionsRequest is the request for GET /api/v1/linkedin/web_v2/get_company_call_to_actions.
type LinkedInWebV2GetCompanyCallToActionsRequest = LinkedInWebV2GetCtaButtonsRequest

// LinkedInWebV2GetCompanyCallToActionsResponse is the response for GET /api/v1/linkedin/web_v2/get_company_call_to_actions.
type LinkedInWebV2GetCompanyCallToActionsResponse = LinkedInWebV2GetCtaButtonsResponse

// GetCompanyCallToActions 获取公司主页 CTA 按钮配置/Get CTA buttons
//
// GET /api/v1/linkedin/web_v2/get_company_call_to_actions
func (r LinkedInWebV2Resource) GetCompanyCallToActions(ctx context.Context, request LinkedInWebV2GetCompanyCallToActionsRequest) (*LinkedInWebV2GetCompanyCallToActionsResponse, error) {
	return r.client.LinkedInWebV2GetCtaButtons(ctx, request)
}

// LinkedInWebV2GetCompanyEmployeeCountRangesRequest is the request for GET /api/v1/linkedin/web_v2/get_company_employee_count_ranges.
type LinkedInWebV2GetCompanyEmployeeCountRangesRequest = LinkedInWebV2GetEmployeeCountBySegmentRequest

// LinkedInWebV2GetCompanyEmployeeCountRangesResponse is the response for GET /api/v1/linkedin/web_v2/get_company_employee_count_ranges.
type LinkedInWebV2GetCompanyEmployeeCountRangesResponse = LinkedInWebV2GetEmployeeCountBySegmentResponse

// GetCompanyEmployeeCountRanges 获取公司员工数量范围（各 segment）/Get employee count by segment
//
// GET /api/v1/linkedin/web_v2/get_company_employee_count_ranges
func (r LinkedInWebV2Resource) GetCompanyEmployeeCountRanges(ctx context.Context, request LinkedInWebV2GetCompanyEmployeeCountRangesRequest) (*LinkedInWebV2GetCompanyEmployeeCountRangesResponse, error) {
	return r.client.LinkedInWebV2GetEmployeeCountBySegment(ctx, request)
}

// LinkedInWebV2GetCompanyGroupedLocationsRequest is the request for GET /api/v1/linkedin/web_v2/get_company_grouped_locations.
type LinkedInWebV2GetCompanyGroupedLocationsRequest = LinkedInWebV2GetGroupedLocationsRequest

// LinkedInWebV2GetCompanyGroupedLocationsResponse is the response for GET /api/v1/linkedin/web_v2/get_company_grouped_locations.
type LinkedInWebV2GetCompanyGroupedLocationsResponse = LinkedInWebV2GetGroupedLocationsResponse

// GetCompanyGroupedLocations 获取公司全部办公地点（按地理分组）/Get grouped locations
//
// GET /api/v1/linkedin/web_v2/get_company_grouped_locations
func (r LinkedInWebV2Resource) GetCompanyGroupedLocations(ctx context.Context, request LinkedInWebV2GetCompanyGroupedLocationsRequest) (*LinkedInWebV2GetCompanyGroupedLocationsResponse, error) {
	return r.client.LinkedInWebV2GetGroupedLocations(ctx, request)
}

// LinkedInWebV2GetPostDetailRequest is the request for GET /api/v1/linkedin/web_v2/get_post_detail.
type LinkedInWebV2GetPostDetailRequest = LinkedInWebV2GetPostDetailByUrnRequest

// LinkedInWebV2GetPostDetailResponse is the response for GET /api/v1/linkedin/web_v2/get_post_detail.
type LinkedInWebV2GetPostDetailResponse = LinkedInWebV2GetPostDetailByUrnResponse

// GetPostDetail 获取单条帖子详情（按 post URN）/Get post detail by URN
//
// GET /api/v1/linkedin/web_v2/get_post_detail
func (r LinkedInWebV2Resource) GetPostDetail(ctx context.Context, request LinkedInWebV2GetPostDetailRequest) (*LinkedInWebV2GetPostDetailResponse, error) {
	return r.client.LinkedInWebV2GetPostDetailByUrn(ctx, request)
}

// LinkedInWebV2GetPostDetailBySlugRequest is the request for GET /api/v1/linkedin/web_v2/get_post_detail_by_slug.
type LinkedInWebV2GetPostDetailBySlugRequest = LinkedInWebV2GetPostByURLSlugRequest

// LinkedInWebV2GetPostDetailBySlugResponse is the response for GET /api/v1/linkedin/web_v2/get_post_detail_by_slug.
type LinkedInWebV2GetPostDetailBySlugResponse = LinkedInWebV2GetPostByURLSlugResponse

// GetPostDetailBySlug 按 URL slug 获取帖子/Get post by URL slug
//
// GET /api/v1/linkedin/web_v2/get_post_detail_by_slug
func (r LinkedInWebV2Resource) GetPostDetailBySlug(ctx context.Context, request LinkedInWebV2GetPostDetailBySlugRequest) (*LinkedInWebV2GetPostDetailBySlugResponse, error) {
	return r.client.LinkedInWebV2GetPostByURLSlug(ctx, request)
}

// LinkedInWebV2GetPostCommentsRequest is the request for GET /api/v1/linkedin/web_v2/get_post_comments.
type LinkedInWebV2GetPostCommentsRequest = LinkedInWebV2GetPostTopLevelCommentsRequest

// LinkedInWebV2GetPostCommentsResponse is the response for GET /api/v1/linkedin/web_v2/get_post_comments.
type LinkedInWebV2GetPostCommentsResponse = LinkedInWebV2GetPostTopLevelCommentsResponse

// GetPostComments 获取帖子顶层评论/Get post top-level comments
//
// GET /api/v1/linkedin/web_v2/get_post_comments
func (r LinkedInWebV2Resource) GetPostComments(ctx context.Context, request LinkedInWebV2GetPostCommentsRequest) (*LinkedInWebV2GetPostCommentsResponse, error) {
	return r.client.LinkedInWebV2GetPostTopLevelComments(ctx, request)
}

// GetCommentReplies 获取评论的回复/Get comment replies
//
// GET /api/v1/linkedin/web_v2/get_comment_replies
func (r LinkedInWebV2Resource) GetCommentReplies(ctx context.Context, request LinkedInWebV2GetCommentRepliesRequest) (*LinkedInWebV2GetCommentRepliesResponse, error) {
	return r.client.LinkedInWebV2GetCommentReplies(ctx, request)
}

// GetPostReactions 获取帖子点赞/反应人列表/Get post reactions
//
// GET /api/v1/linkedin/web_v2/get_post_reactions
func (r LinkedInWebV2Resource) GetPostReactions(ctx context.Context, request LinkedInWebV2GetPostReactionsRequest) (*LinkedInWebV2GetPostReactionsResponse, error) {
	return r.client.LinkedInWebV2GetPostReactions(ctx, request)
}

// GetHashtagFeed 按 hashtag 获取话题动态流/Get hashtag feed
//
// GET /api/v1/linkedin/web_v2/get_hashtag_feed
func (r LinkedInWebV2Resource) GetHashtagFeed(ctx context.Context, request LinkedInWebV2GetHashtagFeedRequest) (*LinkedInWebV2GetHashtagFeedResponse, error) {
	return r.client.LinkedInWebV2GetHashtagFeed(ctx, request)
}

// GetJobDetail 获取职位详情/Get job detail
//
// GET /api/v1/linkedin/web_v2/get_job_detail
func (r LinkedInWebV2Resource) GetJobDetail(ctx context.Context, request LinkedInWebV2GetJobDetailRequest) (*LinkedInWebV2GetJobDetailResponse, error) {
	return r.client.LinkedInWebV2GetJobDetail(ctx, request)
}

// SearchUsers 搜索用户/Search users
//
// GET /api/v1/linkedin/web_v2/search_users
func (r LinkedInWebV2Resource) SearchUsers(ctx context.Context, request LinkedInWebV2SearchUsersRequest) (*LinkedInWebV2SearchUsersResponse, error) {
	return r.client.LinkedInWebV2SearchUsers(ctx, request)
}

// SearchJobs 搜索职位/Search jobs
//
// GET /api/v1/linkedin/web_v2/search_jobs
func (r LinkedInWebV2Resource) SearchJobs(ctx context.Context, request LinkedInWebV2SearchJobsRequest) (*LinkedInWebV2SearchJobsResponse, error) {
	return r.client.LinkedInWebV2SearchJobs(ctx, request)
}

// BilibiliWebResource contains endpoints from the Bilibili-Web-API tag.
type BilibiliWebResource struct {
	client *Client
}

// BilibiliWebFetchOneVideoRequest is the request for GET /api/v1/bilibili/web/fetch_one_video.
type BilibiliWebFetchOneVideoRequest = BilibiliWebGetSingleVideoDataRequest

// BilibiliWebFetchOneVideoResponse is the response for GET /api/v1/bilibili/web/fetch_one_video.
type BilibiliWebFetchOneVideoResponse = BilibiliWebGetSingleVideoDataResponse

// FetchOneVideo 获取单个视频详情信息/Get single video data
//
// GET /api/v1/bilibili/web/fetch_one_video
func (r BilibiliWebResource) FetchOneVideo(ctx context.Context, request BilibiliWebFetchOneVideoRequest) (*BilibiliWebFetchOneVideoResponse, error) {
	return r.client.BilibiliWebGetSingleVideoData(ctx, request)
}

// BilibiliWebFetchOneVideoV2Request is the request for GET /api/v1/bilibili/web/fetch_one_video_v2.
type BilibiliWebFetchOneVideoV2Request = BilibiliWebGetSingleVideoDataV2Request

// BilibiliWebFetchOneVideoV2Response is the response for GET /api/v1/bilibili/web/fetch_one_video_v2.
type BilibiliWebFetchOneVideoV2Response = BilibiliWebGetSingleVideoDataV2Response

// FetchOneVideoV2 获取单个视频详情信息V2/Get single video data V2
//
// GET /api/v1/bilibili/web/fetch_one_video_v2
func (r BilibiliWebResource) FetchOneVideoV2(ctx context.Context, request BilibiliWebFetchOneVideoV2Request) (*BilibiliWebFetchOneVideoV2Response, error) {
	return r.client.BilibiliWebGetSingleVideoDataV2(ctx, request)
}

// BilibiliWebFetchOneVideoV3Request is the request for GET /api/v1/bilibili/web/fetch_one_video_v3.
type BilibiliWebFetchOneVideoV3Request = BilibiliWebGetSingleVideoDataV3Request

// BilibiliWebFetchOneVideoV3Response is the response for GET /api/v1/bilibili/web/fetch_one_video_v3.
type BilibiliWebFetchOneVideoV3Response = BilibiliWebGetSingleVideoDataV3Response

// FetchOneVideoV3 获取单个视频详情信息V3/Get single video data V3
//
// GET /api/v1/bilibili/web/fetch_one_video_v3
func (r BilibiliWebResource) FetchOneVideoV3(ctx context.Context, request BilibiliWebFetchOneVideoV3Request) (*BilibiliWebFetchOneVideoV3Response, error) {
	return r.client.BilibiliWebGetSingleVideoDataV3(ctx, request)
}

// BilibiliWebFetchVideoDetailRequest is the request for GET /api/v1/bilibili/web/fetch_video_detail.
type BilibiliWebFetchVideoDetailRequest = BilibiliWebGetSingleVideoDetailRequest

// BilibiliWebFetchVideoDetailResponse is the response for GET /api/v1/bilibili/web/fetch_video_detail.
type BilibiliWebFetchVideoDetailResponse = BilibiliWebGetSingleVideoDetailResponse

// FetchVideoDetail 获取单个视频详情/Get single video detail
//
// GET /api/v1/bilibili/web/fetch_video_detail
func (r BilibiliWebResource) FetchVideoDetail(ctx context.Context, request BilibiliWebFetchVideoDetailRequest) (*BilibiliWebFetchVideoDetailResponse, error) {
	return r.client.BilibiliWebGetSingleVideoDetail(ctx, request)
}

// BilibiliWebFetchVideoPlayInfoRequest is the request for GET /api/v1/bilibili/web/fetch_video_play_info.
type BilibiliWebFetchVideoPlayInfoRequest = BilibiliWebGetSingleVideoPlayInfoRequest

// BilibiliWebFetchVideoPlayInfoResponse is the response for GET /api/v1/bilibili/web/fetch_video_play_info.
type BilibiliWebFetchVideoPlayInfoResponse = BilibiliWebGetSingleVideoPlayInfoResponse

// FetchVideoPlayInfo 获取单个视频播放信息/Get single video play info
//
// GET /api/v1/bilibili/web/fetch_video_play_info
func (r BilibiliWebResource) FetchVideoPlayInfo(ctx context.Context, request BilibiliWebFetchVideoPlayInfoRequest) (*BilibiliWebFetchVideoPlayInfoResponse, error) {
	return r.client.BilibiliWebGetSingleVideoPlayInfo(ctx, request)
}

// BilibiliWebFetchVideoSubtitleRequest is the request for GET /api/v1/bilibili/web/fetch_video_subtitle.
type BilibiliWebFetchVideoSubtitleRequest = BilibiliWebGetVideoSubtitleInfoRequest

// BilibiliWebFetchVideoSubtitleResponse is the response for GET /api/v1/bilibili/web/fetch_video_subtitle.
type BilibiliWebFetchVideoSubtitleResponse = BilibiliWebGetVideoSubtitleInfoResponse

// FetchVideoSubtitle 获取视频字幕信息/Get video subtitle info
//
// GET /api/v1/bilibili/web/fetch_video_subtitle
func (r BilibiliWebResource) FetchVideoSubtitle(ctx context.Context, request BilibiliWebFetchVideoSubtitleRequest) (*BilibiliWebFetchVideoSubtitleResponse, error) {
	return r.client.BilibiliWebGetVideoSubtitleInfo(ctx, request)
}

// BilibiliWebFetchHotSearchRequest is the request for GET /api/v1/bilibili/web/fetch_hot_search.
type BilibiliWebFetchHotSearchRequest = BilibiliWebGetHotSearchDataRequest

// BilibiliWebFetchHotSearchResponse is the response for GET /api/v1/bilibili/web/fetch_hot_search.
type BilibiliWebFetchHotSearchResponse = BilibiliWebGetHotSearchDataResponse

// FetchHotSearch 获取热门搜索信息/Get hot search data
//
// GET /api/v1/bilibili/web/fetch_hot_search
func (r BilibiliWebResource) FetchHotSearch(ctx context.Context, request BilibiliWebFetchHotSearchRequest) (*BilibiliWebFetchHotSearchResponse, error) {
	return r.client.BilibiliWebGetHotSearchData(ctx, request)
}

// BilibiliWebFetchGeneralSearchRequest is the request for GET /api/v1/bilibili/web/fetch_general_search.
type BilibiliWebFetchGeneralSearchRequest = BilibiliWebGetGeneralSearchDataRequest

// BilibiliWebFetchGeneralSearchResponse is the response for GET /api/v1/bilibili/web/fetch_general_search.
type BilibiliWebFetchGeneralSearchResponse = BilibiliWebGetGeneralSearchDataResponse

// FetchGeneralSearch 获取综合搜索信息/Get general search data
//
// GET /api/v1/bilibili/web/fetch_general_search
func (r BilibiliWebResource) FetchGeneralSearch(ctx context.Context, request BilibiliWebFetchGeneralSearchRequest) (*BilibiliWebFetchGeneralSearchResponse, error) {
	return r.client.BilibiliWebGetGeneralSearchData(ctx, request)
}

// BilibiliWebFetchVideoPlayurlRequest is the request for GET /api/v1/bilibili/web/fetch_video_playurl.
type BilibiliWebFetchVideoPlayurlRequest = BilibiliWebGetVideoPlayurlRequest

// BilibiliWebFetchVideoPlayurlResponse is the response for GET /api/v1/bilibili/web/fetch_video_playurl.
type BilibiliWebFetchVideoPlayurlResponse = BilibiliWebGetVideoPlayurlResponse

// FetchVideoPlayurl 获取视频流地址/Get video playurl
//
// GET /api/v1/bilibili/web/fetch_video_playurl
func (r BilibiliWebResource) FetchVideoPlayurl(ctx context.Context, request BilibiliWebFetchVideoPlayurlRequest) (*BilibiliWebFetchVideoPlayurlResponse, error) {
	return r.client.BilibiliWebGetVideoPlayurl(ctx, request)
}

// BilibiliWebFetchVipVideoPlayurlRequest is the request for POST /api/v1/bilibili/web/fetch_vip_video_playurl.
type BilibiliWebFetchVipVideoPlayurlRequest = BilibiliWebGetVipVideoPlayurlRequest

// BilibiliWebFetchVipVideoPlayurlResponse is the response for POST /api/v1/bilibili/web/fetch_vip_video_playurl.
type BilibiliWebFetchVipVideoPlayurlResponse = BilibiliWebGetVipVideoPlayurlResponse

// FetchVipVideoPlayurl 获取大会员清晰度视频流地址/Get VIP video playurl
//
// POST /api/v1/bilibili/web/fetch_vip_video_playurl
func (r BilibiliWebResource) FetchVipVideoPlayurl(ctx context.Context, request BilibiliWebFetchVipVideoPlayurlRequest) (*BilibiliWebFetchVipVideoPlayurlResponse, error) {
	return r.client.BilibiliWebGetVipVideoPlayurl(ctx, request)
}

// BilibiliWebFetchUserPostVideosRequest is the request for GET /api/v1/bilibili/web/fetch_user_post_videos.
type BilibiliWebFetchUserPostVideosRequest = BilibiliWebGetUserHomepageVideoDataRequest

// BilibiliWebFetchUserPostVideosResponse is the response for GET /api/v1/bilibili/web/fetch_user_post_videos.
type BilibiliWebFetchUserPostVideosResponse = BilibiliWebGetUserHomepageVideoDataResponse

// FetchUserPostVideos 获取用户主页作品数据/Get user homepage video data
//
// GET /api/v1/bilibili/web/fetch_user_post_videos
func (r BilibiliWebResource) FetchUserPostVideos(ctx context.Context, request BilibiliWebFetchUserPostVideosRequest) (*BilibiliWebFetchUserPostVideosResponse, error) {
	return r.client.BilibiliWebGetUserHomepageVideoData(ctx, request)
}

// BilibiliWebFetchCollectFoldersRequest is the request for GET /api/v1/bilibili/web/fetch_collect_folders.
type BilibiliWebFetchCollectFoldersRequest = BilibiliWebGetUserCollectionFoldersRequest

// BilibiliWebFetchCollectFoldersResponse is the response for GET /api/v1/bilibili/web/fetch_collect_folders.
type BilibiliWebFetchCollectFoldersResponse = BilibiliWebGetUserCollectionFoldersResponse

// FetchCollectFolders 获取用户所有收藏夹信息/Get user collection folders
//
// GET /api/v1/bilibili/web/fetch_collect_folders
func (r BilibiliWebResource) FetchCollectFolders(ctx context.Context, request BilibiliWebFetchCollectFoldersRequest) (*BilibiliWebFetchCollectFoldersResponse, error) {
	return r.client.BilibiliWebGetUserCollectionFolders(ctx, request)
}

// BilibiliWebFetchUserCollectionVideosRequest is the request for GET /api/v1/bilibili/web/fetch_user_collection_videos.
type BilibiliWebFetchUserCollectionVideosRequest = BilibiliWebGetsVideoDataFromACollectionFolderRequest

// BilibiliWebFetchUserCollectionVideosResponse is the response for GET /api/v1/bilibili/web/fetch_user_collection_videos.
type BilibiliWebFetchUserCollectionVideosResponse = BilibiliWebGetsVideoDataFromACollectionFolderResponse

// FetchUserCollectionVideos 获取指定收藏夹内视频数据/Gets video data from a collection folder
//
// GET /api/v1/bilibili/web/fetch_user_collection_videos
func (r BilibiliWebResource) FetchUserCollectionVideos(ctx context.Context, request BilibiliWebFetchUserCollectionVideosRequest) (*BilibiliWebFetchUserCollectionVideosResponse, error) {
	return r.client.BilibiliWebGetsVideoDataFromACollectionFolder(ctx, request)
}

// BilibiliWebFetchUserProfileRequest is the request for GET /api/v1/bilibili/web/fetch_user_profile.
type BilibiliWebFetchUserProfileRequest = BilibiliWebGetInformationOfSpecifiedUserRequest

// BilibiliWebFetchUserProfileResponse is the response for GET /api/v1/bilibili/web/fetch_user_profile.
type BilibiliWebFetchUserProfileResponse = BilibiliWebGetInformationOfSpecifiedUserResponse

// FetchUserProfile 获取指定用户的信息/Get information of specified user
//
// GET /api/v1/bilibili/web/fetch_user_profile
func (r BilibiliWebResource) FetchUserProfile(ctx context.Context, request BilibiliWebFetchUserProfileRequest) (*BilibiliWebFetchUserProfileResponse, error) {
	return r.client.BilibiliWebGetInformationOfSpecifiedUser(ctx, request)
}

// BilibiliWebFetchUserUpStatRequest is the request for GET /api/v1/bilibili/web/fetch_user_up_stat.
type BilibiliWebFetchUserUpStatRequest = BilibiliWebGetUpStatRequest

// BilibiliWebFetchUserUpStatResponse is the response for GET /api/v1/bilibili/web/fetch_user_up_stat.
type BilibiliWebFetchUserUpStatResponse = BilibiliWebGetUpStatResponse

// FetchUserUpStat 获取UP主状态统计/Get UP stat (total likes and views)
//
// GET /api/v1/bilibili/web/fetch_user_up_stat
func (r BilibiliWebResource) FetchUserUpStat(ctx context.Context, request BilibiliWebFetchUserUpStatRequest) (*BilibiliWebFetchUserUpStatResponse, error) {
	return r.client.BilibiliWebGetUpStat(ctx, request)
}

// BilibiliWebFetchUserRelationStatRequest is the request for GET /api/v1/bilibili/web/fetch_user_relation_stat.
type BilibiliWebFetchUserRelationStatRequest = BilibiliWebGetUserRelationStatRequest

// BilibiliWebFetchUserRelationStatResponse is the response for GET /api/v1/bilibili/web/fetch_user_relation_stat.
type BilibiliWebFetchUserRelationStatResponse = BilibiliWebGetUserRelationStatResponse

// FetchUserRelationStat 获取用户关系状态统计/Get user relation stat (following and followers)
//
// GET /api/v1/bilibili/web/fetch_user_relation_stat
func (r BilibiliWebResource) FetchUserRelationStat(ctx context.Context, request BilibiliWebFetchUserRelationStatRequest) (*BilibiliWebFetchUserRelationStatResponse, error) {
	return r.client.BilibiliWebGetUserRelationStat(ctx, request)
}

// BilibiliWebFetchComPopularRequest is the request for GET /api/v1/bilibili/web/fetch_com_popular.
type BilibiliWebFetchComPopularRequest = BilibiliWebGetComprehensivePopularVideoInformationRequest

// BilibiliWebFetchComPopularResponse is the response for GET /api/v1/bilibili/web/fetch_com_popular.
type BilibiliWebFetchComPopularResponse = BilibiliWebGetComprehensivePopularVideoInformationResponse

// FetchComPopular 获取综合热门视频信息/Get comprehensive popular video information
//
// GET /api/v1/bilibili/web/fetch_com_popular
func (r BilibiliWebResource) FetchComPopular(ctx context.Context, request BilibiliWebFetchComPopularRequest) (*BilibiliWebFetchComPopularResponse, error) {
	return r.client.BilibiliWebGetComprehensivePopularVideoInformation(ctx, request)
}

// BilibiliWebFetchVideoCommentsRequest is the request for GET /api/v1/bilibili/web/fetch_video_comments.
type BilibiliWebFetchVideoCommentsRequest = BilibiliWebGetCommentsOnTheSpecifiedVideoRequest

// BilibiliWebFetchVideoCommentsResponse is the response for GET /api/v1/bilibili/web/fetch_video_comments.
type BilibiliWebFetchVideoCommentsResponse = BilibiliWebGetCommentsOnTheSpecifiedVideoResponse

// FetchVideoComments 获取指定视频的评论/Get comments on the specified video
//
// GET /api/v1/bilibili/web/fetch_video_comments
func (r BilibiliWebResource) FetchVideoComments(ctx context.Context, request BilibiliWebFetchVideoCommentsRequest) (*BilibiliWebFetchVideoCommentsResponse, error) {
	return r.client.BilibiliWebGetCommentsOnTheSpecifiedVideo(ctx, request)
}

// BilibiliWebFetchCommentReplyRequest is the request for GET /api/v1/bilibili/web/fetch_comment_reply.
type BilibiliWebFetchCommentReplyRequest = BilibiliWebGetReplyToTheSpecifiedCommentRequest

// BilibiliWebFetchCommentReplyResponse is the response for GET /api/v1/bilibili/web/fetch_comment_reply.
type BilibiliWebFetchCommentReplyResponse = BilibiliWebGetReplyToTheSpecifiedCommentResponse

// FetchCommentReply 获取视频下指定评论的回复/Get reply to the specified comment
//
// GET /api/v1/bilibili/web/fetch_comment_reply
func (r BilibiliWebResource) FetchCommentReply(ctx context.Context, request BilibiliWebFetchCommentReplyRequest) (*BilibiliWebFetchCommentReplyResponse, error) {
	return r.client.BilibiliWebGetReplyToTheSpecifiedComment(ctx, request)
}

// BilibiliWebFetchUserDynamicRequest is the request for GET /api/v1/bilibili/web/fetch_user_dynamic.
type BilibiliWebFetchUserDynamicRequest = BilibiliWebGetDynamicInformationOfSpecifiedUserRequest

// BilibiliWebFetchUserDynamicResponse is the response for GET /api/v1/bilibili/web/fetch_user_dynamic.
type BilibiliWebFetchUserDynamicResponse = BilibiliWebGetDynamicInformationOfSpecifiedUserResponse

// FetchUserDynamic 获取指定用户动态/Get dynamic information of specified user
//
// GET /api/v1/bilibili/web/fetch_user_dynamic
func (r BilibiliWebResource) FetchUserDynamic(ctx context.Context, request BilibiliWebFetchUserDynamicRequest) (*BilibiliWebFetchUserDynamicResponse, error) {
	return r.client.BilibiliWebGetDynamicInformationOfSpecifiedUser(ctx, request)
}

// BilibiliWebFetchDynamicDetailRequest is the request for GET /api/v1/bilibili/web/fetch_dynamic_detail.
type BilibiliWebFetchDynamicDetailRequest = BilibiliWebGetDynamicDetailRequest

// BilibiliWebFetchDynamicDetailResponse is the response for GET /api/v1/bilibili/web/fetch_dynamic_detail.
type BilibiliWebFetchDynamicDetailResponse = BilibiliWebGetDynamicDetailResponse

// FetchDynamicDetail 获取动态详情/Get dynamic detail
//
// GET /api/v1/bilibili/web/fetch_dynamic_detail
func (r BilibiliWebResource) FetchDynamicDetail(ctx context.Context, request BilibiliWebFetchDynamicDetailRequest) (*BilibiliWebFetchDynamicDetailResponse, error) {
	return r.client.BilibiliWebGetDynamicDetail(ctx, request)
}

// BilibiliWebFetchDynamicDetailV2Request is the request for GET /api/v1/bilibili/web/fetch_dynamic_detail_v2.
type BilibiliWebFetchDynamicDetailV2Request = BilibiliWebGetDynamicDetailV2Request

// BilibiliWebFetchDynamicDetailV2Response is the response for GET /api/v1/bilibili/web/fetch_dynamic_detail_v2.
type BilibiliWebFetchDynamicDetailV2Response = BilibiliWebGetDynamicDetailV2Response

// FetchDynamicDetailV2 获取动态详情v2/Get dynamic detail v2
//
// GET /api/v1/bilibili/web/fetch_dynamic_detail_v2
func (r BilibiliWebResource) FetchDynamicDetailV2(ctx context.Context, request BilibiliWebFetchDynamicDetailV2Request) (*BilibiliWebFetchDynamicDetailV2Response, error) {
	return r.client.BilibiliWebGetDynamicDetailV2(ctx, request)
}

// BilibiliWebFetchVideoDanmakuRequest is the request for GET /api/v1/bilibili/web/fetch_video_danmaku.
type BilibiliWebFetchVideoDanmakuRequest = BilibiliWebGetVideoDanmakuRequest

// BilibiliWebFetchVideoDanmakuResponse is the response for GET /api/v1/bilibili/web/fetch_video_danmaku.
type BilibiliWebFetchVideoDanmakuResponse = BilibiliWebGetVideoDanmakuResponse

// FetchVideoDanmaku 获取视频实时弹幕/Get Video Danmaku
//
// GET /api/v1/bilibili/web/fetch_video_danmaku
func (r BilibiliWebResource) FetchVideoDanmaku(ctx context.Context, request BilibiliWebFetchVideoDanmakuRequest) (*BilibiliWebFetchVideoDanmakuResponse, error) {
	return r.client.BilibiliWebGetVideoDanmaku(ctx, request)
}

// BilibiliWebFetchLiveRoomDetailRequest is the request for GET /api/v1/bilibili/web/fetch_live_room_detail.
type BilibiliWebFetchLiveRoomDetailRequest = BilibiliWebGetInformationOfSpecifiedLiveRoomRequest

// BilibiliWebFetchLiveRoomDetailResponse is the response for GET /api/v1/bilibili/web/fetch_live_room_detail.
type BilibiliWebFetchLiveRoomDetailResponse = BilibiliWebGetInformationOfSpecifiedLiveRoomResponse

// FetchLiveRoomDetail 获取指定直播间信息/Get information of specified live room
//
// GET /api/v1/bilibili/web/fetch_live_room_detail
func (r BilibiliWebResource) FetchLiveRoomDetail(ctx context.Context, request BilibiliWebFetchLiveRoomDetailRequest) (*BilibiliWebFetchLiveRoomDetailResponse, error) {
	return r.client.BilibiliWebGetInformationOfSpecifiedLiveRoom(ctx, request)
}

// BilibiliWebFetchLiveVideosRequest is the request for GET /api/v1/bilibili/web/fetch_live_videos.
type BilibiliWebFetchLiveVideosRequest = BilibiliWebGetLiveVideoDataOfSpecifiedRoomRequest

// BilibiliWebFetchLiveVideosResponse is the response for GET /api/v1/bilibili/web/fetch_live_videos.
type BilibiliWebFetchLiveVideosResponse = BilibiliWebGetLiveVideoDataOfSpecifiedRoomResponse

// FetchLiveVideos 获取直播间视频流/Get live video data of specified room
//
// GET /api/v1/bilibili/web/fetch_live_videos
func (r BilibiliWebResource) FetchLiveVideos(ctx context.Context, request BilibiliWebFetchLiveVideosRequest) (*BilibiliWebFetchLiveVideosResponse, error) {
	return r.client.BilibiliWebGetLiveVideoDataOfSpecifiedRoom(ctx, request)
}

// BilibiliWebFetchLiveStreamersRequest is the request for GET /api/v1/bilibili/web/fetch_live_streamers.
type BilibiliWebFetchLiveStreamersRequest = BilibiliWebGetLiveStreamersOfSpecifiedLiveAreaRequest

// BilibiliWebFetchLiveStreamersResponse is the response for GET /api/v1/bilibili/web/fetch_live_streamers.
type BilibiliWebFetchLiveStreamersResponse = BilibiliWebGetLiveStreamersOfSpecifiedLiveAreaResponse

// FetchLiveStreamers 获取指定分区正在直播的主播/Get live streamers of specified live area
//
// GET /api/v1/bilibili/web/fetch_live_streamers
func (r BilibiliWebResource) FetchLiveStreamers(ctx context.Context, request BilibiliWebFetchLiveStreamersRequest) (*BilibiliWebFetchLiveStreamersResponse, error) {
	return r.client.BilibiliWebGetLiveStreamersOfSpecifiedLiveArea(ctx, request)
}

// BilibiliWebFetchAllLiveAreasResponse is the response for GET /api/v1/bilibili/web/fetch_all_live_areas.
type BilibiliWebFetchAllLiveAreasResponse = BilibiliWebGetAListOfAllLiveAreasResponse

// FetchAllLiveAreas 获取所有直播分区列表/Get a list of all live areas
//
// GET /api/v1/bilibili/web/fetch_all_live_areas
func (r BilibiliWebResource) FetchAllLiveAreas(ctx context.Context) (*BilibiliWebFetchAllLiveAreasResponse, error) {
	return r.client.BilibiliWebGetAListOfAllLiveAreas(ctx)
}

// BilibiliWebBvToAidRequest is the request for GET /api/v1/bilibili/web/bv_to_aid.
type BilibiliWebBvToAidRequest = BilibiliWebGenerateAidByBvidRequest

// BilibiliWebBvToAidResponse is the response for GET /api/v1/bilibili/web/bv_to_aid.
type BilibiliWebBvToAidResponse = BilibiliWebGenerateAidByBvidResponse

// BvToAid 通过bv号获得视频aid号/Generate aid by bvid
//
// GET /api/v1/bilibili/web/bv_to_aid
func (r BilibiliWebResource) BvToAid(ctx context.Context, request BilibiliWebBvToAidRequest) (*BilibiliWebBvToAidResponse, error) {
	return r.client.BilibiliWebGenerateAidByBvid(ctx, request)
}

// BilibiliWebFetchVideoPartsRequest is the request for GET /api/v1/bilibili/web/fetch_video_parts.
type BilibiliWebFetchVideoPartsRequest = BilibiliWebGetVideoPartsByBvidRequest

// BilibiliWebFetchVideoPartsResponse is the response for GET /api/v1/bilibili/web/fetch_video_parts.
type BilibiliWebFetchVideoPartsResponse = BilibiliWebGetVideoPartsByBvidResponse

// FetchVideoParts 通过bv号获得视频分p信息/Get Video Parts By bvid
//
// GET /api/v1/bilibili/web/fetch_video_parts
func (r BilibiliWebResource) FetchVideoParts(ctx context.Context, request BilibiliWebFetchVideoPartsRequest) (*BilibiliWebFetchVideoPartsResponse, error) {
	return r.client.BilibiliWebGetVideoPartsByBvid(ctx, request)
}

// BilibiliWebFetchGetUserIDRequest is the request for GET /api/v1/bilibili/web/fetch_get_user_id.
type BilibiliWebFetchGetUserIDRequest = BilibiliWebExtractUserIDRequest

// BilibiliWebFetchGetUserIDResponse is the response for GET /api/v1/bilibili/web/fetch_get_user_id.
type BilibiliWebFetchGetUserIDResponse = BilibiliWebExtractUserIDResponse

// FetchGetUserID 提取用户ID/Extract user ID
//
// GET /api/v1/bilibili/web/fetch_get_user_id
func (r BilibiliWebResource) FetchGetUserID(ctx context.Context, request BilibiliWebFetchGetUserIDRequest) (*BilibiliWebFetchGetUserIDResponse, error) {
	return r.client.BilibiliWebExtractUserID(ctx, request)
}

// BilibiliAppResource contains endpoints from the Bilibili-App-API tag.
type BilibiliAppResource struct {
	client *Client
}

// BilibiliAppFetchOneVideoRequest is the request for GET /api/v1/bilibili/app/fetch_one_video.
type BilibiliAppFetchOneVideoRequest = BilibiliAppGetSingleVideoDataRequest

// BilibiliAppFetchOneVideoResponse is the response for GET /api/v1/bilibili/app/fetch_one_video.
type BilibiliAppFetchOneVideoResponse = BilibiliAppGetSingleVideoDataResponse

// FetchOneVideo 获取单个视频详情信息/Get single video data
//
// GET /api/v1/bilibili/app/fetch_one_video
func (r BilibiliAppResource) FetchOneVideo(ctx context.Context, request BilibiliAppFetchOneVideoRequest) (*BilibiliAppFetchOneVideoResponse, error) {
	return r.client.BilibiliAppGetSingleVideoData(ctx, request)
}

// BilibiliAppFetchVideoCommentsRequest is the request for GET /api/v1/bilibili/app/fetch_video_comments.
type BilibiliAppFetchVideoCommentsRequest = BilibiliAppGetVideoCommentsRequest

// BilibiliAppFetchVideoCommentsResponse is the response for GET /api/v1/bilibili/app/fetch_video_comments.
type BilibiliAppFetchVideoCommentsResponse = BilibiliAppGetVideoCommentsResponse

// FetchVideoComments 获取视频评论列表/Get video comments
//
// GET /api/v1/bilibili/app/fetch_video_comments
func (r BilibiliAppResource) FetchVideoComments(ctx context.Context, request BilibiliAppFetchVideoCommentsRequest) (*BilibiliAppFetchVideoCommentsResponse, error) {
	return r.client.BilibiliAppGetVideoComments(ctx, request)
}

// BilibiliAppFetchReplyDetailRequest is the request for GET /api/v1/bilibili/app/fetch_reply_detail.
type BilibiliAppFetchReplyDetailRequest = BilibiliAppGetReplyDetailRequest

// BilibiliAppFetchReplyDetailResponse is the response for GET /api/v1/bilibili/app/fetch_reply_detail.
type BilibiliAppFetchReplyDetailResponse = BilibiliAppGetReplyDetailResponse

// FetchReplyDetail 获取二级评论回复/Get reply detail
//
// GET /api/v1/bilibili/app/fetch_reply_detail
func (r BilibiliAppResource) FetchReplyDetail(ctx context.Context, request BilibiliAppFetchReplyDetailRequest) (*BilibiliAppFetchReplyDetailResponse, error) {
	return r.client.BilibiliAppGetReplyDetail(ctx, request)
}

// BilibiliAppFetchUserVideosRequest is the request for GET /api/v1/bilibili/app/fetch_user_videos.
type BilibiliAppFetchUserVideosRequest = BilibiliAppGetUserVideosRequest

// BilibiliAppFetchUserVideosResponse is the response for GET /api/v1/bilibili/app/fetch_user_videos.
type BilibiliAppFetchUserVideosResponse = BilibiliAppGetUserVideosResponse

// FetchUserVideos 获取用户投稿视频/Get user videos
//
// GET /api/v1/bilibili/app/fetch_user_videos
func (r BilibiliAppResource) FetchUserVideos(ctx context.Context, request BilibiliAppFetchUserVideosRequest) (*BilibiliAppFetchUserVideosResponse, error) {
	return r.client.BilibiliAppGetUserVideos(ctx, request)
}

// BilibiliAppFetchUserInfoRequest is the request for GET /api/v1/bilibili/app/fetch_user_info.
type BilibiliAppFetchUserInfoRequest = BilibiliAppGetUserInfoRequest

// BilibiliAppFetchUserInfoResponse is the response for GET /api/v1/bilibili/app/fetch_user_info.
type BilibiliAppFetchUserInfoResponse = BilibiliAppGetUserInfoResponse

// FetchUserInfo 获取用户信息/Get user info
//
// GET /api/v1/bilibili/app/fetch_user_info
func (r BilibiliAppResource) FetchUserInfo(ctx context.Context, request BilibiliAppFetchUserInfoRequest) (*BilibiliAppFetchUserInfoResponse, error) {
	return r.client.BilibiliAppGetUserInfo(ctx, request)
}

// BilibiliAppFetchHomeFeedRequest is the request for GET /api/v1/bilibili/app/fetch_home_feed.
type BilibiliAppFetchHomeFeedRequest = BilibiliAppGetHomeFeedRequest

// BilibiliAppFetchHomeFeedResponse is the response for GET /api/v1/bilibili/app/fetch_home_feed.
type BilibiliAppFetchHomeFeedResponse = BilibiliAppGetHomeFeedResponse

// FetchHomeFeed 获取主页推荐视频流/Get home feed
//
// GET /api/v1/bilibili/app/fetch_home_feed
func (r BilibiliAppResource) FetchHomeFeed(ctx context.Context, request BilibiliAppFetchHomeFeedRequest) (*BilibiliAppFetchHomeFeedResponse, error) {
	return r.client.BilibiliAppGetHomeFeed(ctx, request)
}

// BilibiliAppFetchPopularFeedRequest is the request for GET /api/v1/bilibili/app/fetch_popular_feed.
type BilibiliAppFetchPopularFeedRequest = BilibiliAppGetPopularFeedRequest

// BilibiliAppFetchPopularFeedResponse is the response for GET /api/v1/bilibili/app/fetch_popular_feed.
type BilibiliAppFetchPopularFeedResponse = BilibiliAppGetPopularFeedResponse

// FetchPopularFeed 获取热门推荐/Get popular feed
//
// GET /api/v1/bilibili/app/fetch_popular_feed
func (r BilibiliAppResource) FetchPopularFeed(ctx context.Context, request BilibiliAppFetchPopularFeedRequest) (*BilibiliAppFetchPopularFeedResponse, error) {
	return r.client.BilibiliAppGetPopularFeed(ctx, request)
}

// BilibiliAppFetchSearchAllRequest is the request for GET /api/v1/bilibili/app/fetch_search_all.
type BilibiliAppFetchSearchAllRequest = BilibiliAppSearchAllRequest

// BilibiliAppFetchSearchAllResponse is the response for GET /api/v1/bilibili/app/fetch_search_all.
type BilibiliAppFetchSearchAllResponse = BilibiliAppSearchAllResponse

// FetchSearchAll 综合搜索/search all
//
// GET /api/v1/bilibili/app/fetch_search_all
func (r BilibiliAppResource) FetchSearchAll(ctx context.Context, request BilibiliAppFetchSearchAllRequest) (*BilibiliAppFetchSearchAllResponse, error) {
	return r.client.BilibiliAppSearchAll(ctx, request)
}

// BilibiliAppFetchSearchByTypeRequest is the request for GET /api/v1/bilibili/app/fetch_search_by_type.
type BilibiliAppFetchSearchByTypeRequest = BilibiliAppSearchByTypeRequest

// BilibiliAppFetchSearchByTypeResponse is the response for GET /api/v1/bilibili/app/fetch_search_by_type.
type BilibiliAppFetchSearchByTypeResponse = BilibiliAppSearchByTypeResponse

// FetchSearchByType 分类搜索/ search by type
//
// GET /api/v1/bilibili/app/fetch_search_by_type
func (r BilibiliAppResource) FetchSearchByType(ctx context.Context, request BilibiliAppFetchSearchByTypeRequest) (*BilibiliAppFetchSearchByTypeResponse, error) {
	return r.client.BilibiliAppSearchByType(ctx, request)
}

// BilibiliAppFetchCinemaTabResponse is the response for GET /api/v1/bilibili/app/fetch_cinema_tab.
type BilibiliAppFetchCinemaTabResponse = BilibiliAppGetCinemaTabResponse

// FetchCinemaTab 获取影视推荐/Get cinema tab
//
// GET /api/v1/bilibili/app/fetch_cinema_tab
func (r BilibiliAppResource) FetchCinemaTab(ctx context.Context) (*BilibiliAppFetchCinemaTabResponse, error) {
	return r.client.BilibiliAppGetCinemaTab(ctx)
}

// BilibiliAppFetchBangumiTabResponse is the response for GET /api/v1/bilibili/app/fetch_bangumi_tab.
type BilibiliAppFetchBangumiTabResponse = BilibiliAppGetBangumiTabResponse

// FetchBangumiTab 获取番剧推荐/Get bangumi tab
//
// GET /api/v1/bilibili/app/fetch_bangumi_tab
func (r BilibiliAppResource) FetchBangumiTab(ctx context.Context) (*BilibiliAppFetchBangumiTabResponse, error) {
	return r.client.BilibiliAppGetBangumiTab(ctx)
}

// Sora2Resource contains endpoints from the Sora2-API tag.
type Sora2Resource struct {
	client *Client
}

// Sora2GetPostDetailRequest is the request for GET /api/v1/sora2/get_post_detail.
type Sora2GetPostDetailRequest = Sora2FetchSinglePostDetailRequest

// Sora2GetPostDetailResponse is the response for GET /api/v1/sora2/get_post_detail.
type Sora2GetPostDetailResponse = Sora2FetchSinglePostDetailResponse

// GetPostDetail 获取单一作品详情/Fetch single post detail
//
// GET /api/v1/sora2/get_post_detail
func (r Sora2Resource) GetPostDetail(ctx context.Context, request Sora2GetPostDetailRequest) (*Sora2GetPostDetailResponse, error) {
	return r.client.Sora2FetchSinglePostDetail(ctx, request)
}

// Sora2GetPostRemixListRequest is the request for GET /api/v1/sora2/get_post_remix_list.
type Sora2GetPostRemixListRequest = Sora2FetchPostRemixListRequest

// Sora2GetPostRemixListResponse is the response for GET /api/v1/sora2/get_post_remix_list.
type Sora2GetPostRemixListResponse = Sora2FetchPostRemixListResponse

// GetPostRemixList 获取作品的 Remix 列表/Fetch post remix list
//
// GET /api/v1/sora2/get_post_remix_list
func (r Sora2Resource) GetPostRemixList(ctx context.Context, request Sora2GetPostRemixListRequest) (*Sora2GetPostRemixListResponse, error) {
	return r.client.Sora2FetchPostRemixList(ctx, request)
}

// Sora2GetVideoDownloadInfoRequest is the request for GET /api/v1/sora2/get_video_download_info.
type Sora2GetVideoDownloadInfoRequest = Sora2FetchNoneWatermarkVideoDownloadInfoRequest

// Sora2GetVideoDownloadInfoResponse is the response for GET /api/v1/sora2/get_video_download_info.
type Sora2GetVideoDownloadInfoResponse = Sora2FetchNoneWatermarkVideoDownloadInfoResponse

// GetVideoDownloadInfo 获取无水印视频下载信息/Fetch none watermark video download info
//
// GET /api/v1/sora2/get_video_download_info
func (r Sora2Resource) GetVideoDownloadInfo(ctx context.Context, request Sora2GetVideoDownloadInfoRequest) (*Sora2GetVideoDownloadInfoResponse, error) {
	return r.client.Sora2FetchNoneWatermarkVideoDownloadInfo(ctx, request)
}

// Sora2GetPostCommentsRequest is the request for GET /api/v1/sora2/get_post_comments.
type Sora2GetPostCommentsRequest = Sora2FetchPostCommentsRequest

// Sora2GetPostCommentsResponse is the response for GET /api/v1/sora2/get_post_comments.
type Sora2GetPostCommentsResponse = Sora2FetchPostCommentsResponse

// GetPostComments 获取作品一级评论/Fetch post comments
//
// GET /api/v1/sora2/get_post_comments
func (r Sora2Resource) GetPostComments(ctx context.Context, request Sora2GetPostCommentsRequest) (*Sora2GetPostCommentsResponse, error) {
	return r.client.Sora2FetchPostComments(ctx, request)
}

// Sora2GetCommentRepliesRequest is the request for GET /api/v1/sora2/get_comment_replies.
type Sora2GetCommentRepliesRequest = Sora2FetchCommentRepliesRequest

// Sora2GetCommentRepliesResponse is the response for GET /api/v1/sora2/get_comment_replies.
type Sora2GetCommentRepliesResponse = Sora2FetchCommentRepliesResponse

// GetCommentReplies 获取评论的回复/Fetch comment replies
//
// GET /api/v1/sora2/get_comment_replies
func (r Sora2Resource) GetCommentReplies(ctx context.Context, request Sora2GetCommentRepliesRequest) (*Sora2GetCommentRepliesResponse, error) {
	return r.client.Sora2FetchCommentReplies(ctx, request)
}

// Sora2GetUserProfileRequest is the request for GET /api/v1/sora2/get_user_profile.
type Sora2GetUserProfileRequest = Sora2FetchUserProfileRequest

// Sora2GetUserProfileResponse is the response for GET /api/v1/sora2/get_user_profile.
type Sora2GetUserProfileResponse = Sora2FetchUserProfileResponse

// GetUserProfile 获取用户信息档案/Fetch user profile
//
// GET /api/v1/sora2/get_user_profile
func (r Sora2Resource) GetUserProfile(ctx context.Context, request Sora2GetUserProfileRequest) (*Sora2GetUserProfileResponse, error) {
	return r.client.Sora2FetchUserProfile(ctx, request)
}

// Sora2GetUserPostsRequest is the request for GET /api/v1/sora2/get_user_posts.
type Sora2GetUserPostsRequest = Sora2FetchUserPostsRequest

// Sora2GetUserPostsResponse is the response for GET /api/v1/sora2/get_user_posts.
type Sora2GetUserPostsResponse = Sora2FetchUserPostsResponse

// GetUserPosts 获取用户发布的帖子列表/Fetch user posts
//
// GET /api/v1/sora2/get_user_posts
func (r Sora2Resource) GetUserPosts(ctx context.Context, request Sora2GetUserPostsRequest) (*Sora2GetUserPostsResponse, error) {
	return r.client.Sora2FetchUserPosts(ctx, request)
}

// Sora2GetCameoLeaderboardRequest is the request for GET /api/v1/sora2/get_cameo_leaderboard.
type Sora2GetCameoLeaderboardRequest = Sora2FetchCameoLeaderboardRequest

// Sora2GetCameoLeaderboardResponse is the response for GET /api/v1/sora2/get_cameo_leaderboard.
type Sora2GetCameoLeaderboardResponse = Sora2FetchCameoLeaderboardResponse

// GetCameoLeaderboard 获取 Cameo 出镜秀达人排行榜/Fetch Cameo leaderboard
//
// GET /api/v1/sora2/get_cameo_leaderboard
func (r Sora2Resource) GetCameoLeaderboard(ctx context.Context, request Sora2GetCameoLeaderboardRequest) (*Sora2GetCameoLeaderboardResponse, error) {
	return r.client.Sora2FetchCameoLeaderboard(ctx, request)
}

// Sora2GetUserCameoAppearancesRequest is the request for GET /api/v1/sora2/get_user_cameo_appearances.
type Sora2GetUserCameoAppearancesRequest = Sora2FetchUserCameoAppearancesRequest

// Sora2GetUserCameoAppearancesResponse is the response for GET /api/v1/sora2/get_user_cameo_appearances.
type Sora2GetUserCameoAppearancesResponse = Sora2FetchUserCameoAppearancesResponse

// GetUserCameoAppearances 获取用户Cameo出镜秀列表/Fetch user cameo appearances
//
// GET /api/v1/sora2/get_user_cameo_appearances
func (r Sora2Resource) GetUserCameoAppearances(ctx context.Context, request Sora2GetUserCameoAppearancesRequest) (*Sora2GetUserCameoAppearancesResponse, error) {
	return r.client.Sora2FetchUserCameoAppearances(ctx, request)
}

// Sora2GetUserFollowersRequest is the request for GET /api/v1/sora2/get_user_followers.
type Sora2GetUserFollowersRequest = Sora2FetchUserFollowersRequest

// Sora2GetUserFollowersResponse is the response for GET /api/v1/sora2/get_user_followers.
type Sora2GetUserFollowersResponse = Sora2FetchUserFollowersResponse

// GetUserFollowers 获取用户粉丝列表/Fetch user followers
//
// GET /api/v1/sora2/get_user_followers
func (r Sora2Resource) GetUserFollowers(ctx context.Context, request Sora2GetUserFollowersRequest) (*Sora2GetUserFollowersResponse, error) {
	return r.client.Sora2FetchUserFollowers(ctx, request)
}

// Sora2GetUserFollowingRequest is the request for GET /api/v1/sora2/get_user_following.
type Sora2GetUserFollowingRequest = Sora2FetchUserFollowingRequest

// Sora2GetUserFollowingResponse is the response for GET /api/v1/sora2/get_user_following.
type Sora2GetUserFollowingResponse = Sora2FetchUserFollowingResponse

// GetUserFollowing 获取用户关注列表/Fetch user following
//
// GET /api/v1/sora2/get_user_following
func (r Sora2Resource) GetUserFollowing(ctx context.Context, request Sora2GetUserFollowingRequest) (*Sora2GetUserFollowingResponse, error) {
	return r.client.Sora2FetchUserFollowing(ctx, request)
}

// Sora2GetFeedRequest is the request for GET /api/v1/sora2/get_feed.
type Sora2GetFeedRequest = Sora2FetchFeedRequest

// Sora2GetFeedResponse is the response for GET /api/v1/sora2/get_feed.
type Sora2GetFeedResponse = Sora2FetchFeedResponse

// GetFeed 获取Feed流（热门/推荐视频）/Fetch feed
//
// GET /api/v1/sora2/get_feed
func (r Sora2Resource) GetFeed(ctx context.Context, request Sora2GetFeedRequest) (*Sora2GetFeedResponse, error) {
	return r.client.Sora2FetchFeed(ctx, request)
}

// SearchUsers 搜索用户/Search users
//
// GET /api/v1/sora2/search_users
func (r Sora2Resource) SearchUsers(ctx context.Context, request Sora2SearchUsersRequest) (*Sora2SearchUsersResponse, error) {
	return r.client.Sora2SearchUsers(ctx, request)
}

// Sora2UploadImageResponse is the response for POST /api/v1/sora2/upload_image.
type Sora2UploadImageResponse = Sora2UploadImageToGetMediaIDResponse

// UploadImage 上传图片获取media_id/Upload image to get media_id
//
// POST /api/v1/sora2/upload_image
func (r Sora2Resource) UploadImage(ctx context.Context) (*Sora2UploadImageResponse, error) {
	return r.client.Sora2UploadImageToGetMediaID(ctx)
}

// Sora2CreateVideoRequest is the request for POST /api/v1/sora2/create_video.
type Sora2CreateVideoRequest = Sora2DeprecatedCreateVideoFromTextOrImageRequest

// Sora2CreateVideoResponse is the response for POST /api/v1/sora2/create_video.
type Sora2CreateVideoResponse = Sora2DeprecatedCreateVideoFromTextOrImageResponse

// CreateVideo [已弃用/Deprecated] 文本/图片生成视频/Create video from text or image
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// POST /api/v1/sora2/create_video
func (r Sora2Resource) CreateVideo(ctx context.Context, request Sora2CreateVideoRequest) (*Sora2CreateVideoResponse, error) {
	return r.client.Sora2DeprecatedCreateVideoFromTextOrImage(ctx, request)
}

// Sora2GetTaskStatusRequest is the request for GET /api/v1/sora2/get_task_status.
type Sora2GetTaskStatusRequest = Sora2DeprecatedGetTaskStatusRequest

// Sora2GetTaskStatusResponse is the response for GET /api/v1/sora2/get_task_status.
type Sora2GetTaskStatusResponse = Sora2DeprecatedGetTaskStatusResponse

// GetTaskStatus [已弃用/Deprecated] 查询任务状态/Get task status
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/sora2/get_task_status
func (r Sora2Resource) GetTaskStatus(ctx context.Context, request Sora2GetTaskStatusRequest) (*Sora2GetTaskStatusResponse, error) {
	return r.client.Sora2DeprecatedGetTaskStatus(ctx, request)
}

// Sora2GetTaskDetailRequest is the request for GET /api/v1/sora2/get_task_detail.
type Sora2GetTaskDetailRequest = Sora2DeprecatedGetTaskGeneratedPostDetailRequest

// Sora2GetTaskDetailResponse is the response for GET /api/v1/sora2/get_task_detail.
type Sora2GetTaskDetailResponse = Sora2DeprecatedGetTaskGeneratedPostDetailResponse

// GetTaskDetail [已弃用/Deprecated] 获取任务生成的作品详情（无水印版本）/Get task-generated post detail (watermark-free)
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/sora2/get_task_detail
func (r Sora2Resource) GetTaskDetail(ctx context.Context, request Sora2GetTaskDetailRequest) (*Sora2GetTaskDetailResponse, error) {
	return r.client.Sora2DeprecatedGetTaskGeneratedPostDetail(ctx, request)
}

// TempMailResource contains endpoints from the Temp-Mail-API tag.
type TempMailResource struct {
	client *Client
}

// TempMailGetTempEmailAddressResponse is the response for GET /api/v1/temp_mail/v1/get_temp_email_address.
type TempMailGetTempEmailAddressResponse = TempMailGetTempEmailResponse

// GetTempEmailAddress Get Temp Email
//
// GET /api/v1/temp_mail/v1/get_temp_email_address
func (r TempMailResource) GetTempEmailAddress(ctx context.Context) (*TempMailGetTempEmailAddressResponse, error) {
	return r.client.TempMailGetTempEmail(ctx)
}

// TempMailGetEmailsInboxRequest is the request for GET /api/v1/temp_mail/v1/get_emails_inbox.
type TempMailGetEmailsInboxRequest = TempMailGetEmailsRequest

// TempMailGetEmailsInboxResponse is the response for GET /api/v1/temp_mail/v1/get_emails_inbox.
type TempMailGetEmailsInboxResponse = TempMailGetEmailsResponse

// GetEmailsInbox Get Emails
//
// GET /api/v1/temp_mail/v1/get_emails_inbox
func (r TempMailResource) GetEmailsInbox(ctx context.Context, request TempMailGetEmailsInboxRequest) (*TempMailGetEmailsInboxResponse, error) {
	return r.client.TempMailGetEmails(ctx, request)
}

// GetEmailByID Get Email By Id
//
// GET /api/v1/temp_mail/v1/get_email_by_id
func (r TempMailResource) GetEmailByID(ctx context.Context, request TempMailGetEmailByIDRequest) (*TempMailGetEmailByIDResponse, error) {
	return r.client.TempMailGetEmailByID(ctx, request)
}

// TwitterWebResource contains endpoints from the Twitter-Web-API tag.
type TwitterWebResource struct {
	client *Client
}

// TwitterWebFetchTweetDetailRequest is the request for GET /api/v1/twitter/web/fetch_tweet_detail.
type TwitterWebFetchTweetDetailRequest = TwitterWebGetSingleTweetDataRequest

// TwitterWebFetchTweetDetailResponse is the response for GET /api/v1/twitter/web/fetch_tweet_detail.
type TwitterWebFetchTweetDetailResponse = TwitterWebGetSingleTweetDataResponse

// FetchTweetDetail 获取单个推文数据/Get single tweet data
//
// GET /api/v1/twitter/web/fetch_tweet_detail
func (r TwitterWebResource) FetchTweetDetail(ctx context.Context, request TwitterWebFetchTweetDetailRequest) (*TwitterWebFetchTweetDetailResponse, error) {
	return r.client.TwitterWebGetSingleTweetData(ctx, request)
}

// TwitterWebFetchUserProfileRequest is the request for GET /api/v1/twitter/web/fetch_user_profile.
type TwitterWebFetchUserProfileRequest = TwitterWebGetUserProfileRequest

// TwitterWebFetchUserProfileResponse is the response for GET /api/v1/twitter/web/fetch_user_profile.
type TwitterWebFetchUserProfileResponse = TwitterWebGetUserProfileResponse

// FetchUserProfile 获取用户资料/Get user profile
//
// GET /api/v1/twitter/web/fetch_user_profile
func (r TwitterWebResource) FetchUserProfile(ctx context.Context, request TwitterWebFetchUserProfileRequest) (*TwitterWebFetchUserProfileResponse, error) {
	return r.client.TwitterWebGetUserProfile(ctx, request)
}

// TwitterWebFetchUserPostTweetRequest is the request for GET /api/v1/twitter/web/fetch_user_post_tweet.
type TwitterWebFetchUserPostTweetRequest = TwitterWebGetUserPostRequest

// TwitterWebFetchUserPostTweetResponse is the response for GET /api/v1/twitter/web/fetch_user_post_tweet.
type TwitterWebFetchUserPostTweetResponse = TwitterWebGetUserPostResponse

// FetchUserPostTweet 获取用户发帖/Get user post
//
// GET /api/v1/twitter/web/fetch_user_post_tweet
func (r TwitterWebResource) FetchUserPostTweet(ctx context.Context, request TwitterWebFetchUserPostTweetRequest) (*TwitterWebFetchUserPostTweetResponse, error) {
	return r.client.TwitterWebGetUserPost(ctx, request)
}

// TwitterWebFetchSearchTimelineRequest is the request for GET /api/v1/twitter/web/fetch_search_timeline.
type TwitterWebFetchSearchTimelineRequest = TwitterWebSearchRequest

// TwitterWebFetchSearchTimelineResponse is the response for GET /api/v1/twitter/web/fetch_search_timeline.
type TwitterWebFetchSearchTimelineResponse = TwitterWebSearchResponse

// FetchSearchTimeline 搜索/Search
//
// GET /api/v1/twitter/web/fetch_search_timeline
func (r TwitterWebResource) FetchSearchTimeline(ctx context.Context, request TwitterWebFetchSearchTimelineRequest) (*TwitterWebFetchSearchTimelineResponse, error) {
	return r.client.TwitterWebSearch(ctx, request)
}

// TwitterWebFetchPostCommentsRequest is the request for GET /api/v1/twitter/web/fetch_post_comments.
type TwitterWebFetchPostCommentsRequest = TwitterWebGetCommentsRequest

// TwitterWebFetchPostCommentsResponse is the response for GET /api/v1/twitter/web/fetch_post_comments.
type TwitterWebFetchPostCommentsResponse = TwitterWebGetCommentsResponse

// FetchPostComments 获取评论/Get comments
//
// GET /api/v1/twitter/web/fetch_post_comments
func (r TwitterWebResource) FetchPostComments(ctx context.Context, request TwitterWebFetchPostCommentsRequest) (*TwitterWebFetchPostCommentsResponse, error) {
	return r.client.TwitterWebGetComments(ctx, request)
}

// TwitterWebFetchLatestPostCommentsRequest is the request for GET /api/v1/twitter/web/fetch_latest_post_comments.
type TwitterWebFetchLatestPostCommentsRequest = TwitterWebGetTheLatestTweetCommentsRequest

// TwitterWebFetchLatestPostCommentsResponse is the response for GET /api/v1/twitter/web/fetch_latest_post_comments.
type TwitterWebFetchLatestPostCommentsResponse = TwitterWebGetTheLatestTweetCommentsResponse

// FetchLatestPostComments 获取最新的推文评论/Get the latest tweet comments
//
// GET /api/v1/twitter/web/fetch_latest_post_comments
func (r TwitterWebResource) FetchLatestPostComments(ctx context.Context, request TwitterWebFetchLatestPostCommentsRequest) (*TwitterWebFetchLatestPostCommentsResponse, error) {
	return r.client.TwitterWebGetTheLatestTweetComments(ctx, request)
}

// TwitterWebFetchUserTweetRepliesRequest is the request for GET /api/v1/twitter/web/fetch_user_tweet_replies.
type TwitterWebFetchUserTweetRepliesRequest = TwitterWebGetUserTweetRepliesRequest

// TwitterWebFetchUserTweetRepliesResponse is the response for GET /api/v1/twitter/web/fetch_user_tweet_replies.
type TwitterWebFetchUserTweetRepliesResponse = TwitterWebGetUserTweetRepliesResponse

// FetchUserTweetReplies 获取用户推文回复/Get user tweet replies
//
// GET /api/v1/twitter/web/fetch_user_tweet_replies
func (r TwitterWebResource) FetchUserTweetReplies(ctx context.Context, request TwitterWebFetchUserTweetRepliesRequest) (*TwitterWebFetchUserTweetRepliesResponse, error) {
	return r.client.TwitterWebGetUserTweetReplies(ctx, request)
}

// TwitterWebFetchUserHighlightsTweetsRequest is the request for GET /api/v1/twitter/web/fetch_user_highlights_tweets.
type TwitterWebFetchUserHighlightsTweetsRequest = TwitterWebGetUserHighlightsTweetsRequest

// TwitterWebFetchUserHighlightsTweetsResponse is the response for GET /api/v1/twitter/web/fetch_user_highlights_tweets.
type TwitterWebFetchUserHighlightsTweetsResponse = TwitterWebGetUserHighlightsTweetsResponse

// FetchUserHighlightsTweets 获取用户高光推文/Get user highlights tweets
//
// Deprecated: 已弃用 / Deprecated. 此接口已在 TikHub OpenAPI 标记为弃用；请查看在线文档获取替代接口。 This endpoint is marked deprecated in TikHub OpenAPI; check the online docs for a replacement.
//
// GET /api/v1/twitter/web/fetch_user_highlights_tweets
func (r TwitterWebResource) FetchUserHighlightsTweets(ctx context.Context, request TwitterWebFetchUserHighlightsTweetsRequest) (*TwitterWebFetchUserHighlightsTweetsResponse, error) {
	return r.client.TwitterWebGetUserHighlightsTweets(ctx, request)
}

// TwitterWebFetchUserMediaRequest is the request for GET /api/v1/twitter/web/fetch_user_media.
type TwitterWebFetchUserMediaRequest = TwitterWebGetUserMediaRequest

// TwitterWebFetchUserMediaResponse is the response for GET /api/v1/twitter/web/fetch_user_media.
type TwitterWebFetchUserMediaResponse = TwitterWebGetUserMediaResponse

// FetchUserMedia 获取用户媒体/Get user media
//
// GET /api/v1/twitter/web/fetch_user_media
func (r TwitterWebResource) FetchUserMedia(ctx context.Context, request TwitterWebFetchUserMediaRequest) (*TwitterWebFetchUserMediaResponse, error) {
	return r.client.TwitterWebGetUserMedia(ctx, request)
}

// TwitterWebFetchRetweetUserListRequest is the request for GET /api/v1/twitter/web/fetch_retweet_user_list.
type TwitterWebFetchRetweetUserListRequest = TwitterWebReTweetUserListRequest

// TwitterWebFetchRetweetUserListResponse is the response for GET /api/v1/twitter/web/fetch_retweet_user_list.
type TwitterWebFetchRetweetUserListResponse = TwitterWebReTweetUserListResponse

// FetchRetweetUserList 转推用户列表/ReTweet User list
//
// GET /api/v1/twitter/web/fetch_retweet_user_list
func (r TwitterWebResource) FetchRetweetUserList(ctx context.Context, request TwitterWebFetchRetweetUserListRequest) (*TwitterWebFetchRetweetUserListResponse, error) {
	return r.client.TwitterWebReTweetUserList(ctx, request)
}

// TwitterWebFetchTrendingRequest is the request for GET /api/v1/twitter/web/fetch_trending.
type TwitterWebFetchTrendingRequest = TwitterWebTrendingRequest

// TwitterWebFetchTrendingResponse is the response for GET /api/v1/twitter/web/fetch_trending.
type TwitterWebFetchTrendingResponse = TwitterWebTrendingResponse

// FetchTrending 趋势/Trending
//
// GET /api/v1/twitter/web/fetch_trending
func (r TwitterWebResource) FetchTrending(ctx context.Context, request TwitterWebFetchTrendingRequest) (*TwitterWebFetchTrendingResponse, error) {
	return r.client.TwitterWebTrending(ctx, request)
}

// TwitterWebFetchUserFollowingsRequest is the request for GET /api/v1/twitter/web/fetch_user_followings.
type TwitterWebFetchUserFollowingsRequest = TwitterWebUserFollowingsRequest

// TwitterWebFetchUserFollowingsResponse is the response for GET /api/v1/twitter/web/fetch_user_followings.
type TwitterWebFetchUserFollowingsResponse = TwitterWebUserFollowingsResponse

// FetchUserFollowings 用户关注/User Followings
//
// GET /api/v1/twitter/web/fetch_user_followings
func (r TwitterWebResource) FetchUserFollowings(ctx context.Context, request TwitterWebFetchUserFollowingsRequest) (*TwitterWebFetchUserFollowingsResponse, error) {
	return r.client.TwitterWebUserFollowings(ctx, request)
}

// TwitterWebFetchUserFollowersRequest is the request for GET /api/v1/twitter/web/fetch_user_followers.
type TwitterWebFetchUserFollowersRequest = TwitterWebUserFollowersRequest

// TwitterWebFetchUserFollowersResponse is the response for GET /api/v1/twitter/web/fetch_user_followers.
type TwitterWebFetchUserFollowersResponse = TwitterWebUserFollowersResponse

// FetchUserFollowers 用户粉丝/User Followers
//
// GET /api/v1/twitter/web/fetch_user_followers
func (r TwitterWebResource) FetchUserFollowers(ctx context.Context, request TwitterWebFetchUserFollowersRequest) (*TwitterWebFetchUserFollowersResponse, error) {
	return r.client.TwitterWebUserFollowers(ctx, request)
}

// ThreadsWebResource contains endpoints from the Threads-Web-API tag.
type ThreadsWebResource struct {
	client *Client
}

// ThreadsWebFetchUserInfoRequest is the request for GET /api/v1/threads/web/fetch_user_info.
type ThreadsWebFetchUserInfoRequest = ThreadsWebGetUserInfoRequest

// ThreadsWebFetchUserInfoResponse is the response for GET /api/v1/threads/web/fetch_user_info.
type ThreadsWebFetchUserInfoResponse = ThreadsWebGetUserInfoResponse

// FetchUserInfo 获取用户信息/Get user info
//
// GET /api/v1/threads/web/fetch_user_info
func (r ThreadsWebResource) FetchUserInfo(ctx context.Context, request ThreadsWebFetchUserInfoRequest) (*ThreadsWebFetchUserInfoResponse, error) {
	return r.client.ThreadsWebGetUserInfo(ctx, request)
}

// ThreadsWebFetchUserInfoByIDRequest is the request for GET /api/v1/threads/web/fetch_user_info_by_id.
type ThreadsWebFetchUserInfoByIDRequest = ThreadsWebGetUserInfoByIDRequest

// ThreadsWebFetchUserInfoByIDResponse is the response for GET /api/v1/threads/web/fetch_user_info_by_id.
type ThreadsWebFetchUserInfoByIDResponse = ThreadsWebGetUserInfoByIDResponse

// FetchUserInfoByID 根据用户ID获取用户信息/Get user info by ID
//
// GET /api/v1/threads/web/fetch_user_info_by_id
func (r ThreadsWebResource) FetchUserInfoByID(ctx context.Context, request ThreadsWebFetchUserInfoByIDRequest) (*ThreadsWebFetchUserInfoByIDResponse, error) {
	return r.client.ThreadsWebGetUserInfoByID(ctx, request)
}

// ThreadsWebFetchUserPostsRequest is the request for GET /api/v1/threads/web/fetch_user_posts.
type ThreadsWebFetchUserPostsRequest = ThreadsWebGetUserPostsRequest

// ThreadsWebFetchUserPostsResponse is the response for GET /api/v1/threads/web/fetch_user_posts.
type ThreadsWebFetchUserPostsResponse = ThreadsWebGetUserPostsResponse

// FetchUserPosts 获取用户帖子列表/Get user posts
//
// GET /api/v1/threads/web/fetch_user_posts
func (r ThreadsWebResource) FetchUserPosts(ctx context.Context, request ThreadsWebFetchUserPostsRequest) (*ThreadsWebFetchUserPostsResponse, error) {
	return r.client.ThreadsWebGetUserPosts(ctx, request)
}

// ThreadsWebFetchUserRepostsRequest is the request for GET /api/v1/threads/web/fetch_user_reposts.
type ThreadsWebFetchUserRepostsRequest = ThreadsWebGetUserRepostsRequest

// ThreadsWebFetchUserRepostsResponse is the response for GET /api/v1/threads/web/fetch_user_reposts.
type ThreadsWebFetchUserRepostsResponse = ThreadsWebGetUserRepostsResponse

// FetchUserReposts 获取用户转发列表/Get user reposts
//
// GET /api/v1/threads/web/fetch_user_reposts
func (r ThreadsWebResource) FetchUserReposts(ctx context.Context, request ThreadsWebFetchUserRepostsRequest) (*ThreadsWebFetchUserRepostsResponse, error) {
	return r.client.ThreadsWebGetUserReposts(ctx, request)
}

// ThreadsWebFetchUserRepliesRequest is the request for GET /api/v1/threads/web/fetch_user_replies.
type ThreadsWebFetchUserRepliesRequest = ThreadsWebGetUserRepliesRequest

// ThreadsWebFetchUserRepliesResponse is the response for GET /api/v1/threads/web/fetch_user_replies.
type ThreadsWebFetchUserRepliesResponse = ThreadsWebGetUserRepliesResponse

// FetchUserReplies 获取用户回复列表/Get user replies
//
// GET /api/v1/threads/web/fetch_user_replies
func (r ThreadsWebResource) FetchUserReplies(ctx context.Context, request ThreadsWebFetchUserRepliesRequest) (*ThreadsWebFetchUserRepliesResponse, error) {
	return r.client.ThreadsWebGetUserReplies(ctx, request)
}

// ThreadsWebFetchPostDetailRequest is the request for GET /api/v1/threads/web/fetch_post_detail.
type ThreadsWebFetchPostDetailRequest = ThreadsWebGetPostDetailRequest

// ThreadsWebFetchPostDetailResponse is the response for GET /api/v1/threads/web/fetch_post_detail.
type ThreadsWebFetchPostDetailResponse = ThreadsWebGetPostDetailResponse

// FetchPostDetail 获取帖子详情/Get post detail
//
// GET /api/v1/threads/web/fetch_post_detail
func (r ThreadsWebResource) FetchPostDetail(ctx context.Context, request ThreadsWebFetchPostDetailRequest) (*ThreadsWebFetchPostDetailResponse, error) {
	return r.client.ThreadsWebGetPostDetail(ctx, request)
}

// ThreadsWebFetchPostDetailV2Request is the request for GET /api/v1/threads/web/fetch_post_detail_v2.
type ThreadsWebFetchPostDetailV2Request = ThreadsWebGetPostDetailV2Request

// ThreadsWebFetchPostDetailV2Response is the response for GET /api/v1/threads/web/fetch_post_detail_v2.
type ThreadsWebFetchPostDetailV2Response = ThreadsWebGetPostDetailV2Response

// FetchPostDetailV2 获取帖子详情 V2(支持链接)/Get post detail V2(supports URL)
//
// GET /api/v1/threads/web/fetch_post_detail_v2
func (r ThreadsWebResource) FetchPostDetailV2(ctx context.Context, request ThreadsWebFetchPostDetailV2Request) (*ThreadsWebFetchPostDetailV2Response, error) {
	return r.client.ThreadsWebGetPostDetailV2(ctx, request)
}

// ThreadsWebFetchPostCommentsRequest is the request for GET /api/v1/threads/web/fetch_post_comments.
type ThreadsWebFetchPostCommentsRequest = ThreadsWebGetPostCommentsRequest

// ThreadsWebFetchPostCommentsResponse is the response for GET /api/v1/threads/web/fetch_post_comments.
type ThreadsWebFetchPostCommentsResponse = ThreadsWebGetPostCommentsResponse

// FetchPostComments 获取帖子评论/Get post comments
//
// GET /api/v1/threads/web/fetch_post_comments
func (r ThreadsWebResource) FetchPostComments(ctx context.Context, request ThreadsWebFetchPostCommentsRequest) (*ThreadsWebFetchPostCommentsResponse, error) {
	return r.client.ThreadsWebGetPostComments(ctx, request)
}

// ThreadsWebSearchTopRequest is the request for GET /api/v1/threads/web/search_top.
type ThreadsWebSearchTopRequest = ThreadsWebSearchTopContentRequest

// ThreadsWebSearchTopResponse is the response for GET /api/v1/threads/web/search_top.
type ThreadsWebSearchTopResponse = ThreadsWebSearchTopContentResponse

// SearchTop 搜索热门内容/Search top content
//
// GET /api/v1/threads/web/search_top
func (r ThreadsWebResource) SearchTop(ctx context.Context, request ThreadsWebSearchTopRequest) (*ThreadsWebSearchTopResponse, error) {
	return r.client.ThreadsWebSearchTopContent(ctx, request)
}

// ThreadsWebSearchRecentRequest is the request for GET /api/v1/threads/web/search_recent.
type ThreadsWebSearchRecentRequest = ThreadsWebSearchRecentContentRequest

// ThreadsWebSearchRecentResponse is the response for GET /api/v1/threads/web/search_recent.
type ThreadsWebSearchRecentResponse = ThreadsWebSearchRecentContentResponse

// SearchRecent 搜索最新内容/Search recent content
//
// GET /api/v1/threads/web/search_recent
func (r ThreadsWebResource) SearchRecent(ctx context.Context, request ThreadsWebSearchRecentRequest) (*ThreadsWebSearchRecentResponse, error) {
	return r.client.ThreadsWebSearchRecentContent(ctx, request)
}

// SearchProfiles 搜索用户档案/Search profiles
//
// GET /api/v1/threads/web/search_profiles
func (r ThreadsWebResource) SearchProfiles(ctx context.Context, request ThreadsWebSearchProfilesRequest) (*ThreadsWebSearchProfilesResponse, error) {
	return r.client.ThreadsWebSearchProfiles(ctx, request)
}

// RedditAppResource contains endpoints from the Reddit-APP-API tag.
type RedditAppResource struct {
	client *Client
}

// RedditAppFetchHomeFeedRequest is the request for GET /api/v1/reddit/app/fetch_home_feed.
type RedditAppFetchHomeFeedRequest = RedditAppFetchRedditAppHomeFeedRequest

// RedditAppFetchHomeFeedResponse is the response for GET /api/v1/reddit/app/fetch_home_feed.
type RedditAppFetchHomeFeedResponse = RedditAppFetchRedditAppHomeFeedResponse

// FetchHomeFeed 获取Reddit APP首页推荐内容/Fetch Reddit APP Home Feed
//
// GET /api/v1/reddit/app/fetch_home_feed
func (r RedditAppResource) FetchHomeFeed(ctx context.Context, request RedditAppFetchHomeFeedRequest) (*RedditAppFetchHomeFeedResponse, error) {
	return r.client.RedditAppFetchRedditAppHomeFeed(ctx, request)
}

// RedditAppFetchPopularFeedRequest is the request for GET /api/v1/reddit/app/fetch_popular_feed.
type RedditAppFetchPopularFeedRequest = RedditAppFetchRedditAppPopularFeedRequest

// RedditAppFetchPopularFeedResponse is the response for GET /api/v1/reddit/app/fetch_popular_feed.
type RedditAppFetchPopularFeedResponse = RedditAppFetchRedditAppPopularFeedResponse

// FetchPopularFeed 获取Reddit APP流行推荐内容/Fetch Reddit APP Popular Feed
//
// GET /api/v1/reddit/app/fetch_popular_feed
func (r RedditAppResource) FetchPopularFeed(ctx context.Context, request RedditAppFetchPopularFeedRequest) (*RedditAppFetchPopularFeedResponse, error) {
	return r.client.RedditAppFetchRedditAppPopularFeed(ctx, request)
}

// RedditAppFetchGamesFeedRequest is the request for GET /api/v1/reddit/app/fetch_games_feed.
type RedditAppFetchGamesFeedRequest = RedditAppFetchRedditAppGamesFeedRequest

// RedditAppFetchGamesFeedResponse is the response for GET /api/v1/reddit/app/fetch_games_feed.
type RedditAppFetchGamesFeedResponse = RedditAppFetchRedditAppGamesFeedResponse

// FetchGamesFeed 获取Reddit APP游戏推荐内容/Fetch Reddit APP Games Feed
//
// GET /api/v1/reddit/app/fetch_games_feed
func (r RedditAppResource) FetchGamesFeed(ctx context.Context, request RedditAppFetchGamesFeedRequest) (*RedditAppFetchGamesFeedResponse, error) {
	return r.client.RedditAppFetchRedditAppGamesFeed(ctx, request)
}

// RedditAppFetchNewsFeedRequest is the request for GET /api/v1/reddit/app/fetch_news_feed.
type RedditAppFetchNewsFeedRequest = RedditAppFetchRedditAppNewsFeedRequest

// RedditAppFetchNewsFeedResponse is the response for GET /api/v1/reddit/app/fetch_news_feed.
type RedditAppFetchNewsFeedResponse = RedditAppFetchRedditAppNewsFeedResponse

// FetchNewsFeed 获取Reddit APP资讯推荐内容/Fetch Reddit APP News Feed
//
// GET /api/v1/reddit/app/fetch_news_feed
func (r RedditAppResource) FetchNewsFeed(ctx context.Context, request RedditAppFetchNewsFeedRequest) (*RedditAppFetchNewsFeedResponse, error) {
	return r.client.RedditAppFetchRedditAppNewsFeed(ctx, request)
}

// RedditAppFetchExploreFeedRequest is the request for GET /api/v1/reddit/app/fetch_explore_feed.
type RedditAppFetchExploreFeedRequest = RedditAppFetchRedditAppExploreFeedRequest

// RedditAppFetchExploreFeedResponse is the response for GET /api/v1/reddit/app/fetch_explore_feed.
type RedditAppFetchExploreFeedResponse = RedditAppFetchRedditAppExploreFeedResponse

// FetchExploreFeed 获取Reddit APP发现页(社区分类+推荐社区)/Fetch Reddit APP Explore Feed
//
// GET /api/v1/reddit/app/fetch_explore_feed
func (r RedditAppResource) FetchExploreFeed(ctx context.Context, request RedditAppFetchExploreFeedRequest) (*RedditAppFetchExploreFeedResponse, error) {
	return r.client.RedditAppFetchRedditAppExploreFeed(ctx, request)
}

// RedditAppFetchTopicFeedRequest is the request for GET /api/v1/reddit/app/fetch_topic_feed.
type RedditAppFetchTopicFeedRequest = RedditAppFetchRedditAppTopicFeedRequest

// RedditAppFetchTopicFeedResponse is the response for GET /api/v1/reddit/app/fetch_topic_feed.
type RedditAppFetchTopicFeedResponse = RedditAppFetchRedditAppTopicFeedResponse

// FetchTopicFeed 按分类获取Reddit APP feed/Fetch Reddit APP Topic Feed
//
// GET /api/v1/reddit/app/fetch_topic_feed
func (r RedditAppResource) FetchTopicFeed(ctx context.Context, request RedditAppFetchTopicFeedRequest) (*RedditAppFetchTopicFeedResponse, error) {
	return r.client.RedditAppFetchRedditAppTopicFeed(ctx, request)
}

// RedditAppFetchPostDetailsRequest is the request for GET /api/v1/reddit/app/fetch_post_details.
type RedditAppFetchPostDetailsRequest = RedditAppFetchSingleRedditPostDetailsRequest

// RedditAppFetchPostDetailsResponse is the response for GET /api/v1/reddit/app/fetch_post_details.
type RedditAppFetchPostDetailsResponse = RedditAppFetchSingleRedditPostDetailsResponse

// FetchPostDetails 获取单个Reddit帖子详情/Fetch Single Reddit Post Details
//
// GET /api/v1/reddit/app/fetch_post_details
func (r RedditAppResource) FetchPostDetails(ctx context.Context, request RedditAppFetchPostDetailsRequest) (*RedditAppFetchPostDetailsResponse, error) {
	return r.client.RedditAppFetchSingleRedditPostDetails(ctx, request)
}

// RedditAppFetchPostDetailsBatchRequest is the request for GET /api/v1/reddit/app/fetch_post_details_batch.
type RedditAppFetchPostDetailsBatchRequest = RedditAppFetchRedditPostDetailsInBatchRequest

// RedditAppFetchPostDetailsBatchResponse is the response for GET /api/v1/reddit/app/fetch_post_details_batch.
type RedditAppFetchPostDetailsBatchResponse = RedditAppFetchRedditPostDetailsInBatchResponse

// FetchPostDetailsBatch 批量获取Reddit帖子详情(最多5条)/Fetch Reddit Post Details in Batch (Max 5)
//
// GET /api/v1/reddit/app/fetch_post_details_batch
func (r RedditAppResource) FetchPostDetailsBatch(ctx context.Context, request RedditAppFetchPostDetailsBatchRequest) (*RedditAppFetchPostDetailsBatchResponse, error) {
	return r.client.RedditAppFetchRedditPostDetailsInBatch(ctx, request)
}

// RedditAppFetchPostDetailsBatchLargeRequest is the request for GET /api/v1/reddit/app/fetch_post_details_batch_large.
type RedditAppFetchPostDetailsBatchLargeRequest = RedditAppFetchRedditPostDetailsInLargeBatchRequest

// RedditAppFetchPostDetailsBatchLargeResponse is the response for GET /api/v1/reddit/app/fetch_post_details_batch_large.
type RedditAppFetchPostDetailsBatchLargeResponse = RedditAppFetchRedditPostDetailsInLargeBatchResponse

// FetchPostDetailsBatchLarge 大批量获取Reddit帖子详情(最多30条)/Fetch Reddit Post Details in Large Batch (Max 30)
//
// GET /api/v1/reddit/app/fetch_post_details_batch_large
func (r RedditAppResource) FetchPostDetailsBatchLarge(ctx context.Context, request RedditAppFetchPostDetailsBatchLargeRequest) (*RedditAppFetchPostDetailsBatchLargeResponse, error) {
	return r.client.RedditAppFetchRedditPostDetailsInLargeBatch(ctx, request)
}

// RedditAppFetchPostCommentsRequest is the request for GET /api/v1/reddit/app/fetch_post_comments.
type RedditAppFetchPostCommentsRequest = RedditAppFetchRedditAppPostCommentsRequest

// RedditAppFetchPostCommentsResponse is the response for GET /api/v1/reddit/app/fetch_post_comments.
type RedditAppFetchPostCommentsResponse = RedditAppFetchRedditAppPostCommentsResponse

// FetchPostComments 获取Reddit APP帖子评论/Fetch Reddit APP Post Comments
//
// GET /api/v1/reddit/app/fetch_post_comments
func (r RedditAppResource) FetchPostComments(ctx context.Context, request RedditAppFetchPostCommentsRequest) (*RedditAppFetchPostCommentsResponse, error) {
	return r.client.RedditAppFetchRedditAppPostComments(ctx, request)
}

// RedditAppFetchCommentRepliesRequest is the request for GET /api/v1/reddit/app/fetch_comment_replies.
type RedditAppFetchCommentRepliesRequest = RedditAppFetchRedditAppCommentRepliesRequest

// RedditAppFetchCommentRepliesResponse is the response for GET /api/v1/reddit/app/fetch_comment_replies.
type RedditAppFetchCommentRepliesResponse = RedditAppFetchRedditAppCommentRepliesResponse

// FetchCommentReplies 获取Reddit APP评论回复（二级评论）/Fetch Reddit APP Comment Replies (Sub-comments)
//
// GET /api/v1/reddit/app/fetch_comment_replies
func (r RedditAppResource) FetchCommentReplies(ctx context.Context, request RedditAppFetchCommentRepliesRequest) (*RedditAppFetchCommentRepliesResponse, error) {
	return r.client.RedditAppFetchRedditAppCommentReplies(ctx, request)
}

// RedditAppFetchSubredditStyleRequest is the request for GET /api/v1/reddit/app/fetch_subreddit_style.
type RedditAppFetchSubredditStyleRequest = RedditAppFetchRedditAppSubredditRulesAndStyleInfoRequest

// RedditAppFetchSubredditStyleResponse is the response for GET /api/v1/reddit/app/fetch_subreddit_style.
type RedditAppFetchSubredditStyleResponse = RedditAppFetchRedditAppSubredditRulesAndStyleInfoResponse

// FetchSubredditStyle 获取Reddit APP版块规则样式信息/Fetch Reddit APP Subreddit Rules and Style Info
//
// GET /api/v1/reddit/app/fetch_subreddit_style
func (r RedditAppResource) FetchSubredditStyle(ctx context.Context, request RedditAppFetchSubredditStyleRequest) (*RedditAppFetchSubredditStyleResponse, error) {
	return r.client.RedditAppFetchRedditAppSubredditRulesAndStyleInfo(ctx, request)
}

// RedditAppFetchSubredditPostChannelsRequest is the request for GET /api/v1/reddit/app/fetch_subreddit_post_channels.
type RedditAppFetchSubredditPostChannelsRequest = RedditAppFetchRedditAppSubredditPostChannelsRequest

// RedditAppFetchSubredditPostChannelsResponse is the response for GET /api/v1/reddit/app/fetch_subreddit_post_channels.
type RedditAppFetchSubredditPostChannelsResponse = RedditAppFetchRedditAppSubredditPostChannelsResponse

// FetchSubredditPostChannels 获取Reddit APP版块帖子频道信息/Fetch Reddit APP Subreddit Post Channels
//
// GET /api/v1/reddit/app/fetch_subreddit_post_channels
func (r RedditAppResource) FetchSubredditPostChannels(ctx context.Context, request RedditAppFetchSubredditPostChannelsRequest) (*RedditAppFetchSubredditPostChannelsResponse, error) {
	return r.client.RedditAppFetchRedditAppSubredditPostChannels(ctx, request)
}

// RedditAppFetchSubredditInfoRequest is the request for GET /api/v1/reddit/app/fetch_subreddit_info.
type RedditAppFetchSubredditInfoRequest = RedditAppFetchRedditAppSubredditInfoRequest

// RedditAppFetchSubredditInfoResponse is the response for GET /api/v1/reddit/app/fetch_subreddit_info.
type RedditAppFetchSubredditInfoResponse = RedditAppFetchRedditAppSubredditInfoResponse

// FetchSubredditInfo 获取Reddit APP版块信息/Fetch Reddit APP Subreddit Info
//
// GET /api/v1/reddit/app/fetch_subreddit_info
func (r RedditAppResource) FetchSubredditInfo(ctx context.Context, request RedditAppFetchSubredditInfoRequest) (*RedditAppFetchSubredditInfoResponse, error) {
	return r.client.RedditAppFetchRedditAppSubredditInfo(ctx, request)
}

// RedditAppFetchSubredditSettingsRequest is the request for GET /api/v1/reddit/app/fetch_subreddit_settings.
type RedditAppFetchSubredditSettingsRequest = RedditAppFetchRedditAppSubredditSettingsRequest

// RedditAppFetchSubredditSettingsResponse is the response for GET /api/v1/reddit/app/fetch_subreddit_settings.
type RedditAppFetchSubredditSettingsResponse = RedditAppFetchRedditAppSubredditSettingsResponse

// FetchSubredditSettings 获取Reddit APP版块设置/Fetch Reddit APP Subreddit Settings
//
// GET /api/v1/reddit/app/fetch_subreddit_settings
func (r RedditAppResource) FetchSubredditSettings(ctx context.Context, request RedditAppFetchSubredditSettingsRequest) (*RedditAppFetchSubredditSettingsResponse, error) {
	return r.client.RedditAppFetchRedditAppSubredditSettings(ctx, request)
}

// RedditAppFetchSearchTypeaheadRequest is the request for GET /api/v1/reddit/app/fetch_search_typeahead.
type RedditAppFetchSearchTypeaheadRequest = RedditAppFetchRedditAppSearchTypeaheadSuggestionsRequest

// RedditAppFetchSearchTypeaheadResponse is the response for GET /api/v1/reddit/app/fetch_search_typeahead.
type RedditAppFetchSearchTypeaheadResponse = RedditAppFetchRedditAppSearchTypeaheadSuggestionsResponse

// FetchSearchTypeahead 获取Reddit APP搜索自动补全建议/Fetch Reddit APP Search Typeahead Suggestions
//
// GET /api/v1/reddit/app/fetch_search_typeahead
func (r RedditAppResource) FetchSearchTypeahead(ctx context.Context, request RedditAppFetchSearchTypeaheadRequest) (*RedditAppFetchSearchTypeaheadResponse, error) {
	return r.client.RedditAppFetchRedditAppSearchTypeaheadSuggestions(ctx, request)
}

// RedditAppFetchDynamicSearchRequest is the request for GET /api/v1/reddit/app/fetch_dynamic_search.
type RedditAppFetchDynamicSearchRequest = RedditAppFetchRedditAppDynamicSearchResultsRequest

// RedditAppFetchDynamicSearchResponse is the response for GET /api/v1/reddit/app/fetch_dynamic_search.
type RedditAppFetchDynamicSearchResponse = RedditAppFetchRedditAppDynamicSearchResultsResponse

// FetchDynamicSearch 获取Reddit APP动态搜索结果/Fetch Reddit APP Dynamic Search Results
//
// GET /api/v1/reddit/app/fetch_dynamic_search
func (r RedditAppResource) FetchDynamicSearch(ctx context.Context, request RedditAppFetchDynamicSearchRequest) (*RedditAppFetchDynamicSearchResponse, error) {
	return r.client.RedditAppFetchRedditAppDynamicSearchResults(ctx, request)
}

// RedditAppFetchCommunityHighlightsRequest is the request for GET /api/v1/reddit/app/fetch_community_highlights.
type RedditAppFetchCommunityHighlightsRequest = RedditAppFetchRedditAppCommunityHighlightsRequest

// RedditAppFetchCommunityHighlightsResponse is the response for GET /api/v1/reddit/app/fetch_community_highlights.
type RedditAppFetchCommunityHighlightsResponse = RedditAppFetchRedditAppCommunityHighlightsResponse

// FetchCommunityHighlights 获取Reddit APP社区亮点/Fetch Reddit APP Community Highlights
//
// GET /api/v1/reddit/app/fetch_community_highlights
func (r RedditAppResource) FetchCommunityHighlights(ctx context.Context, request RedditAppFetchCommunityHighlightsRequest) (*RedditAppFetchCommunityHighlightsResponse, error) {
	return r.client.RedditAppFetchRedditAppCommunityHighlights(ctx, request)
}

// RedditAppFetchTrendingSearchesRequest is the request for GET /api/v1/reddit/app/fetch_trending_searches.
type RedditAppFetchTrendingSearchesRequest = RedditAppFetchRedditAppTrendingSearchesRequest

// RedditAppFetchTrendingSearchesResponse is the response for GET /api/v1/reddit/app/fetch_trending_searches.
type RedditAppFetchTrendingSearchesResponse = RedditAppFetchRedditAppTrendingSearchesResponse

// FetchTrendingSearches 获取Reddit APP今日热门搜索/Fetch Reddit APP Trending Searches
//
// GET /api/v1/reddit/app/fetch_trending_searches
func (r RedditAppResource) FetchTrendingSearches(ctx context.Context, request RedditAppFetchTrendingSearchesRequest) (*RedditAppFetchTrendingSearchesResponse, error) {
	return r.client.RedditAppFetchRedditAppTrendingSearches(ctx, request)
}

// RedditAppFetchGeneratedPostsRequest is the request for GET /api/v1/reddit/app/fetch_generated_posts.
type RedditAppFetchGeneratedPostsRequest = RedditAppFetchRedditAnswersGeneratedPostsRequest

// RedditAppFetchGeneratedPostsResponse is the response for GET /api/v1/reddit/app/fetch_generated_posts.
type RedditAppFetchGeneratedPostsResponse = RedditAppFetchRedditAnswersGeneratedPostsResponse

// FetchGeneratedPosts 批量获取Reddit Answers卡片精简帖子信息/Fetch Reddit Answers Generated Posts
//
// GET /api/v1/reddit/app/fetch_generated_posts
func (r RedditAppResource) FetchGeneratedPosts(ctx context.Context, request RedditAppFetchGeneratedPostsRequest) (*RedditAppFetchGeneratedPostsResponse, error) {
	return r.client.RedditAppFetchRedditAnswersGeneratedPosts(ctx, request)
}

// RedditAppFetchGeneratedCommentsRequest is the request for GET /api/v1/reddit/app/fetch_generated_comments.
type RedditAppFetchGeneratedCommentsRequest = RedditAppFetchRedditAnswersGeneratedCommentsRequest

// RedditAppFetchGeneratedCommentsResponse is the response for GET /api/v1/reddit/app/fetch_generated_comments.
type RedditAppFetchGeneratedCommentsResponse = RedditAppFetchRedditAnswersGeneratedCommentsResponse

// FetchGeneratedComments 批量获取Reddit Answers卡片精简评论信息/Fetch Reddit Answers Generated Comments
//
// GET /api/v1/reddit/app/fetch_generated_comments
func (r RedditAppResource) FetchGeneratedComments(ctx context.Context, request RedditAppFetchGeneratedCommentsRequest) (*RedditAppFetchGeneratedCommentsResponse, error) {
	return r.client.RedditAppFetchRedditAnswersGeneratedComments(ctx, request)
}

// RedditAppFetchUserProfileRequest is the request for GET /api/v1/reddit/app/fetch_user_profile.
type RedditAppFetchUserProfileRequest = RedditAppFetchRedditAppUserProfileRequest

// RedditAppFetchUserProfileResponse is the response for GET /api/v1/reddit/app/fetch_user_profile.
type RedditAppFetchUserProfileResponse = RedditAppFetchRedditAppUserProfileResponse

// FetchUserProfile 获取Reddit APP用户资料信息/Fetch Reddit APP User Profile
//
// GET /api/v1/reddit/app/fetch_user_profile
func (r RedditAppResource) FetchUserProfile(ctx context.Context, request RedditAppFetchUserProfileRequest) (*RedditAppFetchUserProfileResponse, error) {
	return r.client.RedditAppFetchRedditAppUserProfile(ctx, request)
}

// RedditAppFetchUserActiveSubredditsRequest is the request for GET /api/v1/reddit/app/fetch_user_active_subreddits.
type RedditAppFetchUserActiveSubredditsRequest = RedditAppFetchUserSActiveSubredditsRequest

// RedditAppFetchUserActiveSubredditsResponse is the response for GET /api/v1/reddit/app/fetch_user_active_subreddits.
type RedditAppFetchUserActiveSubredditsResponse = RedditAppFetchUserSActiveSubredditsResponse

// FetchUserActiveSubreddits 获取用户活跃的社区列表/Fetch User's Active Subreddits
//
// GET /api/v1/reddit/app/fetch_user_active_subreddits
func (r RedditAppResource) FetchUserActiveSubreddits(ctx context.Context, request RedditAppFetchUserActiveSubredditsRequest) (*RedditAppFetchUserActiveSubredditsResponse, error) {
	return r.client.RedditAppFetchUserSActiveSubreddits(ctx, request)
}

// FetchUserComments 获取用户评论列表/Fetch User Comments
//
// GET /api/v1/reddit/app/fetch_user_comments
func (r RedditAppResource) FetchUserComments(ctx context.Context, request RedditAppFetchUserCommentsRequest) (*RedditAppFetchUserCommentsResponse, error) {
	return r.client.RedditAppFetchUserComments(ctx, request)
}

// FetchUserPosts 获取用户发布的帖子列表/Fetch User Posts
//
// GET /api/v1/reddit/app/fetch_user_posts
func (r RedditAppResource) FetchUserPosts(ctx context.Context, request RedditAppFetchUserPostsRequest) (*RedditAppFetchUserPostsResponse, error) {
	return r.client.RedditAppFetchUserPosts(ctx, request)
}

// RedditAppFetchSubredditFeedRequest is the request for GET /api/v1/reddit/app/fetch_subreddit_feed.
type RedditAppFetchSubredditFeedRequest = RedditAppFetchRedditAppSubredditFeedRequest

// RedditAppFetchSubredditFeedResponse is the response for GET /api/v1/reddit/app/fetch_subreddit_feed.
type RedditAppFetchSubredditFeedResponse = RedditAppFetchRedditAppSubredditFeedResponse

// FetchSubredditFeed 获取Reddit APP版块Feed内容/Fetch Reddit APP Subreddit Feed
//
// GET /api/v1/reddit/app/fetch_subreddit_feed
func (r RedditAppResource) FetchSubredditFeed(ctx context.Context, request RedditAppFetchSubredditFeedRequest) (*RedditAppFetchSubredditFeedResponse, error) {
	return r.client.RedditAppFetchRedditAppSubredditFeed(ctx, request)
}

// RedditAppCheckSubredditMutedRequest is the request for GET /api/v1/reddit/app/check_subreddit_muted.
type RedditAppCheckSubredditMutedRequest = RedditAppCheckIfSubredditIsMutedRequest

// RedditAppCheckSubredditMutedResponse is the response for GET /api/v1/reddit/app/check_subreddit_muted.
type RedditAppCheckSubredditMutedResponse = RedditAppCheckIfSubredditIsMutedResponse

// CheckSubredditMuted 检查版块是否静音/Check if Subreddit is Muted
//
// GET /api/v1/reddit/app/check_subreddit_muted
func (r RedditAppResource) CheckSubredditMuted(ctx context.Context, request RedditAppCheckSubredditMutedRequest) (*RedditAppCheckSubredditMutedResponse, error) {
	return r.client.RedditAppCheckIfSubredditIsMuted(ctx, request)
}

// RedditAppFetchUserTrophiesRequest is the request for GET /api/v1/reddit/app/fetch_user_trophies.
type RedditAppFetchUserTrophiesRequest = RedditAppFetchUserPublicTrophiesRequest

// RedditAppFetchUserTrophiesResponse is the response for GET /api/v1/reddit/app/fetch_user_trophies.
type RedditAppFetchUserTrophiesResponse = RedditAppFetchUserPublicTrophiesResponse

// FetchUserTrophies 获取用户公开奖杯/Fetch User Public Trophies
//
// GET /api/v1/reddit/app/fetch_user_trophies
func (r RedditAppResource) FetchUserTrophies(ctx context.Context, request RedditAppFetchUserTrophiesRequest) (*RedditAppFetchUserTrophiesResponse, error) {
	return r.client.RedditAppFetchUserPublicTrophies(ctx, request)
}

// HybridParsingResource contains endpoints from the Hybrid-Parsing tag.
type HybridParsingResource struct {
	client *Client
}

// HybridParsingVideoDataRequest is the request for GET /api/v1/hybrid/video_data.
type HybridParsingVideoDataRequest = HybridParsingHybridParsingSingleVideoEndpointRequest

// HybridParsingVideoDataResponse is the response for GET /api/v1/hybrid/video_data.
type HybridParsingVideoDataResponse = HybridParsingHybridParsingSingleVideoEndpointResponse

// VideoData 混合解析单一视频接口/Hybrid parsing single video endpoint
//
// GET /api/v1/hybrid/video_data
func (r HybridParsingResource) VideoData(ctx context.Context, request HybridParsingVideoDataRequest) (*HybridParsingVideoDataResponse, error) {
	return r.client.HybridParsingHybridParsingSingleVideoEndpoint(ctx, request)
}

// IOSShortcutResource contains endpoints from the iOS-Shortcut tag.
type IOSShortcutResource struct {
	client *Client
}

// IOSShortcutShortcutResponse is the response for GET /api/v1/ios_shortcut/shortcut.
type IOSShortcutShortcutResponse = IOsShortcutVersionUpdateInformationForIOsShortcutsResponse

// Shortcut 用于iOS快捷指令的版本更新信息/Version update information for iOS shortcuts
//
// GET /api/v1/ios_shortcut/shortcut
func (r IOSShortcutResource) Shortcut(ctx context.Context) (*IOSShortcutShortcutResponse, error) {
	return r.client.IOsShortcutVersionUpdateInformationForIOsShortcuts(ctx)
}

// DemoResource contains endpoints from the Demo-API tag.
type DemoResource struct {
	client *Client
}

// DemoCacheStatusResponse is the response for GET /api/v1/demo/demo/cache_status.
type DemoCacheStatusResponse = DemoViewDemoCacheStatusResponse

// CacheStatus 查看Demo缓存状态/View Demo Cache Status
//
// GET /api/v1/demo/demo/cache_status
func (r DemoResource) CacheStatus(ctx context.Context) (*DemoCacheStatusResponse, error) {
	return r.client.DemoViewDemoCacheStatus(ctx)
}

// DemoDouyinWebFetchOneVideoResponse is the response for GET /api/v1/demo/douyin/web/fetch_one_video.
type DemoDouyinWebFetchOneVideoResponse = DemoDemoFetchDouyinWebFixedVideoDataWithCacheResponse

// DouyinWebFetchOneVideo 【Demo】抖音Web获取固定作品数据（1小时缓存）/[Demo] Fetch Douyin Web Fixed Video Data with Cache
//
// GET /api/v1/demo/douyin/web/fetch_one_video
func (r DemoResource) DouyinWebFetchOneVideo(ctx context.Context) (*DemoDouyinWebFetchOneVideoResponse, error) {
	return r.client.DemoDemoFetchDouyinWebFixedVideoDataWithCache(ctx)
}

// DemoDouyinAppFetchOneVideoResponse is the response for GET /api/v1/demo/douyin/app/fetch_one_video.
type DemoDouyinAppFetchOneVideoResponse = DemoDemoFetchDouyinAppFixedVideoDataWithCacheResponse

// DouyinAppFetchOneVideo 【Demo】抖音APP获取固定作品数据（1小时缓存）/[Demo] Fetch Douyin APP Fixed Video Data with Cache
//
// GET /api/v1/demo/douyin/app/fetch_one_video
func (r DemoResource) DouyinAppFetchOneVideo(ctx context.Context) (*DemoDouyinAppFetchOneVideoResponse, error) {
	return r.client.DemoDemoFetchDouyinAppFixedVideoDataWithCache(ctx)
}

// DemoGeneralSearchResponse is the response for GET /api/v1/demo/douyin_search/app/general_search.
type DemoGeneralSearchResponse = DemoDemoDouyinGeneralSearchWithCacheResponse

// GeneralSearch 【Demo】抖音搜索综合搜索（1小时缓存）/[Demo] Douyin General Search with Cache
//
// GET /api/v1/demo/douyin_search/app/general_search
func (r DemoResource) GeneralSearch(ctx context.Context) (*DemoGeneralSearchResponse, error) {
	return r.client.DemoDemoDouyinGeneralSearchWithCache(ctx)
}

// DemoKuaishouWebFetchOneVideoResponse is the response for GET /api/v1/demo/kuaishou/web/fetch_one_video.
type DemoKuaishouWebFetchOneVideoResponse = DemoDemoKuaishouFixedVideoWithCacheResponse

// KuaishouWebFetchOneVideo 【Demo】快手获取固定视频信息（1小时缓存）/[Demo] Kuaishou Fixed Video with Cache
//
// GET /api/v1/demo/kuaishou/web/fetch_one_video
func (r DemoResource) KuaishouWebFetchOneVideo(ctx context.Context) (*DemoKuaishouWebFetchOneVideoResponse, error) {
	return r.client.DemoDemoKuaishouFixedVideoWithCache(ctx)
}

// DemoFetchUserProfileResponse is the response for GET /api/v1/demo/tiktok/web/fetch_user_profile.
type DemoFetchUserProfileResponse = DemoDemoTikTokFixedUserProfileWithCacheResponse

// FetchUserProfile 【Demo】TikTok固定用户信息（1小时缓存）/[Demo] TikTok Fixed User Profile with Cache
//
// GET /api/v1/demo/tiktok/web/fetch_user_profile
func (r DemoResource) FetchUserProfile(ctx context.Context) (*DemoFetchUserProfileResponse, error) {
	return r.client.DemoDemoTikTokFixedUserProfileWithCache(ctx)
}

// DemoTikTokAppFetchOneVideoResponse is the response for GET /api/v1/demo/tiktok/app/fetch_one_video.
type DemoTikTokAppFetchOneVideoResponse = DemoDemoTikTokAppFixedVideoDetailWithCacheResponse

// TikTokAppFetchOneVideo 【Demo】TikTok APP获取固定视频详情（1小时缓存）/[Demo] TikTok APP Fixed Video Detail with Cache
//
// GET /api/v1/demo/tiktok/app/fetch_one_video
func (r DemoResource) TikTokAppFetchOneVideo(ctx context.Context) (*DemoTikTokAppFetchOneVideoResponse, error) {
	return r.client.DemoDemoTikTokAppFixedVideoDetailWithCache(ctx)
}

// DemoFetchUserInfoResponse is the response for GET /api/v1/demo/instagram/web/fetch_user_info.
type DemoFetchUserInfoResponse = DemoDemoInstagramFixedUserProfileWithCacheResponse

// FetchUserInfo 【Demo】Instagram获取固定用户信息（1小时缓存）/[Demo] Instagram Fixed User Profile with Cache
//
// GET /api/v1/demo/instagram/web/fetch_user_info
func (r DemoResource) FetchUserInfo(ctx context.Context) (*DemoFetchUserInfoResponse, error) {
	return r.client.DemoDemoInstagramFixedUserProfileWithCache(ctx)
}

// DemoArticleExtractResponse is the response for GET /api/v1/demo/wechat/article_extract.
type DemoArticleExtractResponse = DemoDemoWeChatArticleExtractWithCacheResponse

// ArticleExtract 【Demo】微信公众号文章提取（1小时缓存）/[Demo] WeChat Article Extract with Cache
//
// GET /api/v1/demo/wechat/article_extract
func (r DemoResource) ArticleExtract(ctx context.Context) (*DemoArticleExtractResponse, error) {
	return r.client.DemoDemoWeChatArticleExtractWithCache(ctx)
}
