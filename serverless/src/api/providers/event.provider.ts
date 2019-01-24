import { Event } from '../types/event';
import jsonEvents from './events.json';
import { User } from '../types/user';
import { UserProvider } from './user.provider';

const userProvider: UserProvider = new UserProvider();
export class EventProvider {
    public getEvents(): Event[] {
        const events: Event[] = [];
        jsonEvents.forEach(event => {
            events.push(this.extendEvent(event));     
        });
        return events;
    }

    public getEvent(id: number): Event {
        const jsonEvent: any = jsonEvents.find(id => id === id);
        return this.extendEvent(jsonEvent); 
    }

    private getUsersByIds(ids: number[]): User[] {
        const users: User[] = [];
        ids.forEach(userId => {
            const user: User = userProvider.getUser(userId);
            users.push(user);
        })
        return users;
    }

    private extendEvent(jsonEvent: any): Event {
        const members: User[] = this.getUsersByIds(jsonEvent.members);
        const organizers: User[] = this.getUsersByIds(jsonEvent.organizers);
        
        let event = Object.assign({}, jsonEvent);
        event.members = members;
        event.organizers = organizers;

        return event; 
    }
}
