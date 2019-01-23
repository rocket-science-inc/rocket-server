import { GraphQLModule } from '@graphql-modules/core';
import { ApolloServer } from 'apollo-server-express';
import express from 'express'
import http from 'http'

export async function bootstrap(AppModule: GraphQLModule) {
  const { schema, context } = AppModule;
  const apollo = new ApolloServer({
    schema,
    context,
    introspection: true,
    playground: true,
  });

  const app = express()
  apollo.applyMiddleware({ app, path: '/graphql' });

  const server = http.createServer(app);
  apollo.installSubscriptionHandlers(server)

  server.listen({ port: 4000 }, () =>
    console.log(
      '🚀  Server ready at',
      `http://localhost:4000${apollo.graphqlPath}`
    )
  )
}
