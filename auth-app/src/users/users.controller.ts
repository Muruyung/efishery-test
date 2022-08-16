import { Controller, Get, UseGuards } from '@nestjs/common';
import { CurrentUser } from 'src/auth/decorator/current-user.decorator';
import { JwtAuthGuard } from 'src/auth/guards/jwt-auth.guard';
import { Users } from './entities/user.entity';
import { UserObject } from './dto/user.object';

@Controller('users')
export class UsersController {
    /**
     * getProfile is used to get current user data
     * @param user 
     * @returns current user data
     */
    @UseGuards(JwtAuthGuard)
    @Get("me")
    getProfile(@CurrentUser() user: Users) {
        var userObject = new UserObject(user);
        return userObject;
    }
}
