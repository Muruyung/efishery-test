import { Field, InputType } from "@nestjs/graphql";

@InputType()
export class CreateUserInput {
    @Field()
    phoneNumber: string;

    @Field()
    name: string;

    @Field()
    password: string;
    
    @Field()
    role: string;
}
