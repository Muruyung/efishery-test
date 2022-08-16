import { Exclude, Expose } from "class-transformer";
import { IsNotEmpty, IsString } from "class-validator";

@Exclude()
export class RegisterUserBody {
  @Expose({ name: "role" })
  @IsString()
  @IsNotEmpty()
  role: string;
  
  @Expose({ name: "phone_number" })
  @IsString()
  @IsNotEmpty()
  phoneNumber: string;
  
  @Expose({ name: "name" })
  @IsString()
  @IsNotEmpty()
  name: string;
}
