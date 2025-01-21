// types/api-response.d.ts
export interface ApiResponse<T> {
  code: number;
  data: T;
  message: string;
}
