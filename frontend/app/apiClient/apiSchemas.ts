/**
 * Generated by @openapi-codegen
 *
 * @version 1.0
 */
/**
 * 座標情報
 */
export type Coordinate = {
  /**
   * 経度
   */
  latitude: number;
  /**
   * 緯度
   */
  longitude: number;
};

/**
 * 配車要求ステータス
 *
 * MATCHING: サービス上でマッチング処理を行なっていて椅子が確定していない
 * DISPATCHING: 椅子が確定し、乗車位置に向かっている
 * DISPATCHED: 椅子が乗車位置に到着して、ユーザーの乗車を待機している
 * CARRYING: ユーザーが乗車し、椅子が目的地に向かっている
 * ARRIVED: 目的地に到着した
 * COMPLETED: ユーザーの決済・椅子評価が完了した
 * CANCELED: 何らかの理由により途中でキャンセルされた(一定時間待ったが椅子を割り当てられなかった場合などを想定)
 */
export type RequestStatus =
  | "MATCHING"
  | "DISPATCHING"
  | "DISPATCHED"
  | "CARRYING"
  | "ARRIVED"
  | "COMPLETED"
  | "CANCELED";

/**
 * 簡易椅子情報
 */
export type Chair = {
  /**
   * 椅子ID
   */
  id: string;
  /**
   * 椅子名
   */
  name: string;
  /**
   * 車種
   */
  chair_model: string;
  /**
   * カーナンバー
   */
  chair_no: string;
};

/**
 * 簡易ユーザー情報
 */
export type User = {
  /**
   * ユーザーID
   */
  id: string;
  /**
   * ユーザー名
   */
  name: string;
};

/**
 * 問い合わせ内容
 */
export type InquiryContent = {
  /**
   * 問い合わせID
   */
  id: string;
  /**
   * 件名
   */
  subject: string;
  /**
   * 問い合わせ内容
   */
  body: string;
  /**
   * 問い合わせ日時
   */
  created_at: number;
};

/**
 * App向け配車要求情報
 */
export type AppRequest = {
  /**
   * 配車要求ID
   */
  request_id: string;
  pickup_coordinate: Coordinate;
  destination_coordinate: Coordinate;
  status: RequestStatus;
  chair?: Chair;
  /**
   * 配車要求日時
   */
  created_at: number;
  /**
   * 配車要求更新日時
   */
  updated_at: number;
};

/**
 * Chair向け配車要求情報
 */
export type ChairRequest = {
  /**
   * 配車要求ID
   */
  request_id: string;
  user: User;
  destination_coordinate: Coordinate;
  status?: RequestStatus;
};

export type Error = {
  message: string;
};
