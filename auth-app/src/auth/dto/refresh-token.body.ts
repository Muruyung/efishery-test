import { Exclude, Expose } from "class-transformer";

@Exclude()
export class RefreshTokenBody {
  @Expose({ name: "refresh_token" })
  refreshToken: string;
}
