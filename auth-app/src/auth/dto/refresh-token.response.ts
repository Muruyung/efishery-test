import { UserDto } from "../../users/dto/user.dto";
import { Exclude, Expose } from "class-transformer";

@Exclude()
export class RefreshTokenResponse {
  @Expose({ name: "user" })
  user: UserDto;

  @Expose({ name: "access_token" })
  accessToken: string;
}
