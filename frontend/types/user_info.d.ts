// types/user_info.d.ts

export interface GetUserInfoRequest {}

export interface GetUserInfoResponse {
  user_id: number;
  username: string;
  avatar: string;
  email: string;
}
