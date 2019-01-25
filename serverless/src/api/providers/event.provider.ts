import { Event } from '../types/event';
import  jsonEvents from '../datasets/events.json';
import { User } from '../types/user';
import { UserProvider } from './user.provider';

const userProvider: UserProvider = new UserProvider();
export class EventProvider {
    public getEvents(): Event[] {
        const extendedEvents: Event[] = [];
        jsonEvents.forEach((jsonEvent: any) => {
            extendedEvents.push(this.extendEvent(jsonEvent));     
        });
        return extendedEvents;
    }

    public getEvent(id: number): Event {
        const jsonEvent: any = jsonEvents.find((jsonEvent: any) => jsonEvent.id === id);
        return this.extendEvent(jsonEvent); 
    }

    public createEvent(title: string): Event {
        var ids: number[] = jsonEvents.map((jsonEvent: any) => jsonEvent.id);
        const nextId: number = Math.max(...ids) + 1;
        jsonEvents.push({
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
        const jsonEvent: any = jsonEvents.find((jsonEvent: any) => jsonEvent.id === id);
        const index: number = jsonEvents.indexOf(jsonEvent, 0);
        if (index === -1) {
            return false;
        }
        jsonEvents.splice(index, 1);
        const deletedEvent: any = jsonEvents.find((jsonEvent: any) => jsonEvent.id === id);
        return deletedEvent === undefined; 
    }

    private getUsersByIds(ids: number[]): User[] {
        const users: User[] = [];
        if (ids === undefined) {
            return users;    
        }
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
