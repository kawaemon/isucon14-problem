// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// AppGetNearbyChairs invokes app-get-nearby-chairs operation.
	//
	// ユーザーの近くにいる椅子を取得する.
	//
	// GET /app/nearby-chairs
	AppGetNearbyChairs(ctx context.Context, params AppGetNearbyChairsParams) (*AppGetNearbyChairsOK, error)
	// AppGetNotification invokes app-get-notification operation.
	//
	// 最新の自分のライドの状態を取得・通知する.
	//
	// GET /app/notification
	AppGetNotification(ctx context.Context) (*AppGetNotificationOK, error)
	// AppGetRides invokes app-get-rides operation.
	//
	// ユーザーが完了済みのライド一覧を取得する.
	//
	// GET /app/rides
	AppGetRides(ctx context.Context) (*AppGetRidesOK, error)
	// AppPostPaymentMethods invokes app-post-payment-methods operation.
	//
	// 決済トークンの登録.
	//
	// POST /app/payment-methods
	AppPostPaymentMethods(ctx context.Context, request OptAppPostPaymentMethodsReq) (AppPostPaymentMethodsRes, error)
	// AppPostRideEvaluation invokes app-post-ride-evaluation operation.
	//
	// ユーザーがライドを評価する.
	//
	// POST /app/rides/{ride_id}/evaluation
	AppPostRideEvaluation(ctx context.Context, request OptAppPostRideEvaluationReq, params AppPostRideEvaluationParams) (AppPostRideEvaluationRes, error)
	// AppPostRides invokes app-post-rides operation.
	//
	// ユーザーが配車を要求する.
	//
	// POST /app/rides
	AppPostRides(ctx context.Context, request OptAppPostRidesReq) (AppPostRidesRes, error)
	// AppPostRidesEstimatedFare invokes app-post-rides-estimated-fare operation.
	//
	// ライドの運賃を見積もる.
	//
	// POST /app/rides/estimated-fare
	AppPostRidesEstimatedFare(ctx context.Context, request OptAppPostRidesEstimatedFareReq) (AppPostRidesEstimatedFareRes, error)
	// AppPostUsers invokes app-post-users operation.
	//
	// ユーザーが会員登録を行う.
	//
	// POST /app/users
	AppPostUsers(ctx context.Context, request OptAppPostUsersReq) (AppPostUsersRes, error)
	// ChairGetNotification invokes chair-get-notification operation.
	//
	// 自分に割り当てられた最新のライドの状態を取得・通知する.
	//
	// GET /chair/notification
	ChairGetNotification(ctx context.Context) (*ChairGetNotificationOK, error)
	// ChairPostActivity invokes chair-post-activity operation.
	//
	// 椅子が配車受付を開始・停止する.
	//
	// POST /chair/activity
	ChairPostActivity(ctx context.Context, request OptChairPostActivityReq) error
	// ChairPostChairs invokes chair-post-chairs operation.
	//
	// オーナーが椅子の登録を行う.
	//
	// POST /chair/chairs
	ChairPostChairs(ctx context.Context, request OptChairPostChairsReq) (*ChairPostChairsCreatedHeaders, error)
	// ChairPostCoordinate invokes chair-post-coordinate operation.
	//
	// 椅子が自身の位置情報を送信する.
	//
	// POST /chair/coordinate
	ChairPostCoordinate(ctx context.Context, request OptCoordinate) (*ChairPostCoordinateOK, error)
	// ChairPostRideStatus invokes chair-post-ride-status operation.
	//
	// 椅子がライドのステータスを更新する.
	//
	// POST /chair/rides/{ride_id}/status
	ChairPostRideStatus(ctx context.Context, request OptChairPostRideStatusReq, params ChairPostRideStatusParams) (ChairPostRideStatusRes, error)
	// OwnerGetChairs invokes owner-get-chairs operation.
	//
	// 椅子のオーナーが管理している椅子の一覧を取得する.
	//
	// GET /owner/chairs
	OwnerGetChairs(ctx context.Context) (*OwnerGetChairsOK, error)
	// OwnerGetSales invokes owner-get-sales operation.
	//
	// 椅子のオーナーが指定期間の全体・椅子ごと・モデルごとの売上情報を取得する.
	//
	// GET /owner/sales
	OwnerGetSales(ctx context.Context, params OwnerGetSalesParams) (*OwnerGetSalesOK, error)
	// OwnerPostOwners invokes owner-post-owners operation.
	//
	// 椅子のオーナーが会員登録を行う.
	//
	// POST /owner/owners
	OwnerPostOwners(ctx context.Context, request OptOwnerPostOwnersReq) (OwnerPostOwnersRes, error)
	// PostInitialize invokes post-initialize operation.
	//
	// サービスを初期化する.
	//
	// POST /initialize
	PostInitialize(ctx context.Context, request OptPostInitializeReq) (*PostInitializeOK, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// AppGetNearbyChairs invokes app-get-nearby-chairs operation.
//
// ユーザーの近くにいる椅子を取得する.
//
// GET /app/nearby-chairs
func (c *Client) AppGetNearbyChairs(ctx context.Context, params AppGetNearbyChairsParams) (*AppGetNearbyChairsOK, error) {
	res, err := c.sendAppGetNearbyChairs(ctx, params)
	return res, err
}

func (c *Client) sendAppGetNearbyChairs(ctx context.Context, params AppGetNearbyChairsParams) (res *AppGetNearbyChairsOK, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/app/nearby-chairs"
	uri.AddPathParts(u, pathParts[:]...)

	q := uri.NewQueryEncoder()
	{
		// Encode "latitude" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "latitude",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.IntToString(params.Latitude))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "longitude" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "longitude",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.IntToString(params.Longitude))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "distance" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "distance",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Distance.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeAppGetNearbyChairsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AppGetNotification invokes app-get-notification operation.
//
// 最新の自分のライドの状態を取得・通知する.
//
// GET /app/notification
func (c *Client) AppGetNotification(ctx context.Context) (*AppGetNotificationOK, error) {
	res, err := c.sendAppGetNotification(ctx)
	return res, err
}

func (c *Client) sendAppGetNotification(ctx context.Context) (res *AppGetNotificationOK, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/app/notification"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeAppGetNotificationResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AppGetRides invokes app-get-rides operation.
//
// ユーザーが完了済みのライド一覧を取得する.
//
// GET /app/rides
func (c *Client) AppGetRides(ctx context.Context) (*AppGetRidesOK, error) {
	res, err := c.sendAppGetRides(ctx)
	return res, err
}

func (c *Client) sendAppGetRides(ctx context.Context) (res *AppGetRidesOK, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/app/rides"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeAppGetRidesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AppPostPaymentMethods invokes app-post-payment-methods operation.
//
// 決済トークンの登録.
//
// POST /app/payment-methods
func (c *Client) AppPostPaymentMethods(ctx context.Context, request OptAppPostPaymentMethodsReq) (AppPostPaymentMethodsRes, error) {
	res, err := c.sendAppPostPaymentMethods(ctx, request)
	return res, err
}

func (c *Client) sendAppPostPaymentMethods(ctx context.Context, request OptAppPostPaymentMethodsReq) (res AppPostPaymentMethodsRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/app/payment-methods"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAppPostPaymentMethodsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeAppPostPaymentMethodsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AppPostRideEvaluation invokes app-post-ride-evaluation operation.
//
// ユーザーがライドを評価する.
//
// POST /app/rides/{ride_id}/evaluation
func (c *Client) AppPostRideEvaluation(ctx context.Context, request OptAppPostRideEvaluationReq, params AppPostRideEvaluationParams) (AppPostRideEvaluationRes, error) {
	res, err := c.sendAppPostRideEvaluation(ctx, request, params)
	return res, err
}

func (c *Client) sendAppPostRideEvaluation(ctx context.Context, request OptAppPostRideEvaluationReq, params AppPostRideEvaluationParams) (res AppPostRideEvaluationRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/app/rides/"
	{
		// Encode "ride_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "ride_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.RideID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/evaluation"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAppPostRideEvaluationRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeAppPostRideEvaluationResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AppPostRides invokes app-post-rides operation.
//
// ユーザーが配車を要求する.
//
// POST /app/rides
func (c *Client) AppPostRides(ctx context.Context, request OptAppPostRidesReq) (AppPostRidesRes, error) {
	res, err := c.sendAppPostRides(ctx, request)
	return res, err
}

func (c *Client) sendAppPostRides(ctx context.Context, request OptAppPostRidesReq) (res AppPostRidesRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/app/rides"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAppPostRidesRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeAppPostRidesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AppPostRidesEstimatedFare invokes app-post-rides-estimated-fare operation.
//
// ライドの運賃を見積もる.
//
// POST /app/rides/estimated-fare
func (c *Client) AppPostRidesEstimatedFare(ctx context.Context, request OptAppPostRidesEstimatedFareReq) (AppPostRidesEstimatedFareRes, error) {
	res, err := c.sendAppPostRidesEstimatedFare(ctx, request)
	return res, err
}

func (c *Client) sendAppPostRidesEstimatedFare(ctx context.Context, request OptAppPostRidesEstimatedFareReq) (res AppPostRidesEstimatedFareRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/app/rides/estimated-fare"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAppPostRidesEstimatedFareRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeAppPostRidesEstimatedFareResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AppPostUsers invokes app-post-users operation.
//
// ユーザーが会員登録を行う.
//
// POST /app/users
func (c *Client) AppPostUsers(ctx context.Context, request OptAppPostUsersReq) (AppPostUsersRes, error) {
	res, err := c.sendAppPostUsers(ctx, request)
	return res, err
}

func (c *Client) sendAppPostUsers(ctx context.Context, request OptAppPostUsersReq) (res AppPostUsersRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/app/users"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAppPostUsersRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeAppPostUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ChairGetNotification invokes chair-get-notification operation.
//
// 自分に割り当てられた最新のライドの状態を取得・通知する.
//
// GET /chair/notification
func (c *Client) ChairGetNotification(ctx context.Context) (*ChairGetNotificationOK, error) {
	res, err := c.sendChairGetNotification(ctx)
	return res, err
}

func (c *Client) sendChairGetNotification(ctx context.Context) (res *ChairGetNotificationOK, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/chair/notification"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeChairGetNotificationResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ChairPostActivity invokes chair-post-activity operation.
//
// 椅子が配車受付を開始・停止する.
//
// POST /chair/activity
func (c *Client) ChairPostActivity(ctx context.Context, request OptChairPostActivityReq) error {
	_, err := c.sendChairPostActivity(ctx, request)
	return err
}

func (c *Client) sendChairPostActivity(ctx context.Context, request OptChairPostActivityReq) (res *ChairPostActivityNoContent, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/chair/activity"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeChairPostActivityRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeChairPostActivityResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ChairPostChairs invokes chair-post-chairs operation.
//
// オーナーが椅子の登録を行う.
//
// POST /chair/chairs
func (c *Client) ChairPostChairs(ctx context.Context, request OptChairPostChairsReq) (*ChairPostChairsCreatedHeaders, error) {
	res, err := c.sendChairPostChairs(ctx, request)
	return res, err
}

func (c *Client) sendChairPostChairs(ctx context.Context, request OptChairPostChairsReq) (res *ChairPostChairsCreatedHeaders, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/chair/chairs"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeChairPostChairsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeChairPostChairsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ChairPostCoordinate invokes chair-post-coordinate operation.
//
// 椅子が自身の位置情報を送信する.
//
// POST /chair/coordinate
func (c *Client) ChairPostCoordinate(ctx context.Context, request OptCoordinate) (*ChairPostCoordinateOK, error) {
	res, err := c.sendChairPostCoordinate(ctx, request)
	return res, err
}

func (c *Client) sendChairPostCoordinate(ctx context.Context, request OptCoordinate) (res *ChairPostCoordinateOK, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/chair/coordinate"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeChairPostCoordinateRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeChairPostCoordinateResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ChairPostRideStatus invokes chair-post-ride-status operation.
//
// 椅子がライドのステータスを更新する.
//
// POST /chair/rides/{ride_id}/status
func (c *Client) ChairPostRideStatus(ctx context.Context, request OptChairPostRideStatusReq, params ChairPostRideStatusParams) (ChairPostRideStatusRes, error) {
	res, err := c.sendChairPostRideStatus(ctx, request, params)
	return res, err
}

func (c *Client) sendChairPostRideStatus(ctx context.Context, request OptChairPostRideStatusReq, params ChairPostRideStatusParams) (res ChairPostRideStatusRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/chair/rides/"
	{
		// Encode "ride_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "ride_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.RideID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/status"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeChairPostRideStatusRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeChairPostRideStatusResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OwnerGetChairs invokes owner-get-chairs operation.
//
// 椅子のオーナーが管理している椅子の一覧を取得する.
//
// GET /owner/chairs
func (c *Client) OwnerGetChairs(ctx context.Context) (*OwnerGetChairsOK, error) {
	res, err := c.sendOwnerGetChairs(ctx)
	return res, err
}

func (c *Client) sendOwnerGetChairs(ctx context.Context) (res *OwnerGetChairsOK, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/owner/chairs"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOwnerGetChairsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OwnerGetSales invokes owner-get-sales operation.
//
// 椅子のオーナーが指定期間の全体・椅子ごと・モデルごとの売上情報を取得する.
//
// GET /owner/sales
func (c *Client) OwnerGetSales(ctx context.Context, params OwnerGetSalesParams) (*OwnerGetSalesOK, error) {
	res, err := c.sendOwnerGetSales(ctx, params)
	return res, err
}

func (c *Client) sendOwnerGetSales(ctx context.Context, params OwnerGetSalesParams) (res *OwnerGetSalesOK, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/owner/sales"
	uri.AddPathParts(u, pathParts[:]...)

	q := uri.NewQueryEncoder()
	{
		// Encode "since" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "since",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Since.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "until" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "until",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Until.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOwnerGetSalesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OwnerPostOwners invokes owner-post-owners operation.
//
// 椅子のオーナーが会員登録を行う.
//
// POST /owner/owners
func (c *Client) OwnerPostOwners(ctx context.Context, request OptOwnerPostOwnersReq) (OwnerPostOwnersRes, error) {
	res, err := c.sendOwnerPostOwners(ctx, request)
	return res, err
}

func (c *Client) sendOwnerPostOwners(ctx context.Context, request OptOwnerPostOwnersReq) (res OwnerPostOwnersRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/owner/owners"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeOwnerPostOwnersRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOwnerPostOwnersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PostInitialize invokes post-initialize operation.
//
// サービスを初期化する.
//
// POST /initialize
func (c *Client) PostInitialize(ctx context.Context, request OptPostInitializeReq) (*PostInitializeOK, error) {
	res, err := c.sendPostInitialize(ctx, request)
	return res, err
}

func (c *Client) sendPostInitialize(ctx context.Context, request OptPostInitializeReq) (res *PostInitializeOK, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/initialize"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodePostInitializeRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePostInitializeResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
