import { Exclude, Expose } from "class-transformer";
import { IsNotEmpty, IsString } from "class-validator";

@Exclude()
export class LoginUserBody {
  @Expose({ name: "phone_number" })
  @IsString()
  @IsNotEmpty()
  phoneNumber: string;
  
  @Expose({ name: "password" })
  @IsString()
  @IsNotEmpty()
  password: string;
}
