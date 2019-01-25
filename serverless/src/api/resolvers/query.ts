import { EventProvider } from '../providers/event.provider';
import { UserProvider } from '../providers/user.provider';
import { User } from '../types/user';
import { Event } from '../types/event';

const eventProvider = new EventProvider();
const userProvider = new UserProvider();

export default {
  Query: {
    events: () => eventProvider.getEvents(),
    event: (root, args) => eventProvider.getEvent(+args.id),
    users: () => userProvider.getUsers(),
    user: (root, args) => userProvider.getUser(+args.id),
  },
};
