import { createParamDecorator, ExecutionContext } from "@nestjs/common";
import { GqlExecutionContext } from "@nestjs/graphql";
import { Users } from "../../users/entities/user.entity";

export const GqlCurrentUser = createParamDecorator<
  unknown,
  ExecutionContext,
  Users
>((_, context) => {
  const ctx = GqlExecutionContext.create(context);
  return ctx.getContext().req.user;
});
