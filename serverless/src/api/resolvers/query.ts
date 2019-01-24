import { EventProvider } from '../providers/event.provider';
import { UserProvider } from '../providers/user.provider';
import { User } from '../types/user';
import { Event } from '../types/event';

const eventProvider = new EventProvider();
const userProvider = new UserProvider();

export default {
  Query: {
    events: (events: Event[]) => eventProvider.getEvents(),
    event: (event: Event, id: number) => eventProvider.getEvent(id),
    users: (users: User[]) => userProvider.getUsers(),
    user: (user: User, id: number) => userProvider.getUser(id),
  },
};
