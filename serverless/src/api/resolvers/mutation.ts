import { EventProvider } from '../providers/event.provider';
import { UserProvider } from '../providers/user.provider';
import { User } from '../types/user';
import { Event } from '../types/event';

const eventProvider = new EventProvider();
const userProvider = new UserProvider();

export default {
  Mutation: {
    createEvent: (root, args) => eventProvider.createEvent(args.title),
    deleteEvent: (root, args) => eventProvider.deleteEvent(+args.id),
    createUser: (root, args) => userProvider.createUser(args.email),
    deleteUser: (root, args) => userProvider.deleteUser(+args.id),
  },
};
