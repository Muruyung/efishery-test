import { createParamDecorator, ExecutionContext } from "@nestjs/common";
import { Users } from "../../users/entities/user.entity";

export const CurrentUser = createParamDecorator<
  unknown,
  ExecutionContext,
  Users
>((_, ctx) => {
  const request = ctx.switchToHttp().getRequest();
  return request.user;
});
