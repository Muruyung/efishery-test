import { Exclude, Expose } from "class-transformer";
import { Users } from "../entities/user.entity";

@Exclude()
export class UserDto {
  constructor(
    partial: Pick<Users, "id" | "phoneNumber" | "name" | "role" | "password">,
  ) {
    Object.assign(this, partial);
  }

  @Expose({ name: "id" })
  readonly id: string;

  @Expose({ name: "phone_number" })
  readonly phoneNumber: string;

  @Expose({ name: "name" })
  readonly name: string;

  @Expose({ name: "role" })
  readonly role: string;

  @Expose({ name: "password" })
  readonly password: string;
}
