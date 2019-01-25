import { EventProvider } from '../providers/event.provider';
import { UserProvider } from '../providers/user.provider';
import { User } from '../types/user';
import { Event } from '../types/event';

const eventProvider = new EventProvider();
const userProvider = new UserProvider();

export default {
  Mutation: {
    createEvent: (event: Event, title: string) => eventProvider.createEvent(title),
    deleteEvent: (event: Event, id: number) => eventProvider.deleteEvent(id),
    createUser: (user: User, email: string) => userProvider.createUser(email),
    deleteUser: (user: User, id: number) => userProvider.getUser(id),
  },
};
