import { UserDto } from "../../users/dto/user.dto";
import { Exclude, Expose } from "class-transformer";

@Exclude()
export class RegisterUserResponse {
  @Expose({ name: "user" })
  user: UserDto;

  @Expose({ name: "access_token" })
  accessToken: string;

  @Expose({ name: "refresh_token" })
  refreshToken: string;
}
