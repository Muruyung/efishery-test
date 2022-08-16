import { Exclude, Expose } from "class-transformer";
import { Users } from "../entities/user.entity";

@Exclude()
export class UserObject {
    constructor(
        partial: Pick<Users, "phoneNumber" | "name" | "role" | "createdAt" >,
    ) {
        Object.assign(this, partial);
    }

    @Expose({ name: "phone_number" })
    readonly phoneNumber: string;

    @Expose({ name: "name" })
    readonly name: string;

    @Expose({ name: "role" })
    readonly role: string;

    @Expose({ name: "created_at" })
    readonly createdAt: string;
}
