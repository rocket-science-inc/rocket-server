import 'reflect-metadata';
import { bootstrap } from './server';
import { GraphQLModule } from '@graphql-modules/core';
import { ApiModule } from './api/api.module';

export const AppModule = new GraphQLModule({
  imports: [
    ApiModule,
  ],
});

bootstrap(AppModule);
