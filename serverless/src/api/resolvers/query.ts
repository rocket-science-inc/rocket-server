import { ModuleContext } from '@graphql-modules/core';
import { EventProvider } from '../providers/event.provider';
import { UserProvider } from '../providers/user.provider';

export default {
  Query: {
    events: (root, args, { injector }: ModuleContext) => 
      injector.get(EventProvider).getEvents(),
    
    event: (root, {id}: any, { injector }: ModuleContext) => 
      injector.get(EventProvider).getEvent(id),
    
    users: (root, args, { injector }: ModuleContext) => 
      injector.get(UserProvider).getUsers(),

    user: (root, {id}: any, { injector }: ModuleContext) => 
      injector.get(UserProvider).getUser(id),
  },
};
