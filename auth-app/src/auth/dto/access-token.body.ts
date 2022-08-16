import { Exclude, Expose } from "class-transformer";

@Exclude()
export class AccessTokenBody {
  @Expose({ name: "access_token" })
  accessToken: string;
}
