import { ModuleContext } from '@graphql-modules/core';
import { EventProvider } from '../providers/event.provider';

export default {
  Query: {
    getEvents: (root, args, {injector}: ModuleContext) => injector.get(EventProvider).getEvents(),
  },
};
