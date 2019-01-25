import { Event } from '../types/event';
import  events from './events.json';
import { User } from '../types/user';
import { UserProvider } from './user.provider';

const userProvider: UserProvider = new UserProvider();
export class EventProvider {
    public getEvents(): Event[] {
        const extendedEvents: Event[] = [];
        events.forEach(jsonEvent => {
            extendedEvents.push(this.extendEvent(jsonEvent));     
        });
        return extendedEvents;
    }

    public getEvent(id: number): Event {
        const jsonEvent: any = events.find(jsonEvent => jsonEvent.id === id);
        return this.extendEvent(jsonEvent); 
    }

    public createEvent(title: string): Event {
        var ids: number[] = events.map(jsonEvent => jsonEvent.id);
        const nextId: number = Math.max(...ids) + 1;
        events.push({
            id: nextId,
            title: title,
            createdAt: new Date().toDateString(),
            description: "",
            location: "",
            members: [],
            organizers: [],
        });
        return this.getEvent(nextId);
    }

    public deleteEvent(id: number): boolean {
        const jsonEvent: any = events.find(jsonEvent => jsonEvent.id === id);
        const index: number = events.indexOf(jsonEvent, 0);
        if (index === -1) {
            return false;
        }
        events.splice(index, 1);
        const deletedEvent: any = events.find(jsonEvent => jsonEvent.id === id);
        return deletedEvent === null; 
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
