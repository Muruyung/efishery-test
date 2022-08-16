import { UserObject } from "src/users/dto/user.object";
import { Exclude, Expose } from "class-transformer";

@Exclude()
export class LoginUserResponse {
  @Expose({ name: "user" })
  user: UserObject;

  @Expose({ name: "access_token" })
  accessToken: string;

  @Expose({ name: "refresh_token" })
  refreshToken: string;
}
