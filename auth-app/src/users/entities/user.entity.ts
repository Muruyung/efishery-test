import { Entity, Column, PrimaryGeneratedColumn, CreateDateColumn, UpdateDateColumn, DeleteDateColumn, OneToMany, JoinColumn, OneToOne } from 'typeorm';
import { ObjectType, Field } from '@nestjs/graphql';
import { RefreshToken } from 'src/auth/entities/refresh-token.entity';
import { Exclude } from 'class-transformer';

@ObjectType()
@Entity()
export class Users {
    @Field()
    @PrimaryGeneratedColumn('uuid')
    id: string;

    @Field()
    @Column('varchar', { nullable: false })
    @Exclude({ toPlainOnly: true })
    password: string;

    @Field()
    @Column('varchar', { length: 13, name: "phone_number" })
    phoneNumber: string;

    @Field()
    @Column('varchar')
    name: string;

    @Field()
    @Column('varchar')
    role: string;

    @OneToMany(() => RefreshToken, (refreshToken) => refreshToken.user, {
        cascade: true,
    })
    refreshTokens: RefreshToken[];

    @Field()
    @CreateDateColumn({ type: "timestamp with time zone", name: "created_at" })
    createdAt: Date;

    @Field()
    @UpdateDateColumn({ type: "timestamp with time zone", name: "updated_at" })
    updatedAt: Date;

    @Field()
    @DeleteDateColumn({ type: "timestamp with time zone", name: "deleted_at" })
    deletedAt: Date;
}
