import { Injectable } from '@graphql-modules/di';
import { ModuleContext } from '@graphql-modules/core';
import { UserProvider } from './user.provider';
import events from './events.json';

@Injectable()
export class EventProvider {
    getEvents() {
        console.log(events);
        return events; 
    }
    getEvent(id: number) { 
        let event = events.find(({ id }) => id === id);
        console.log(event);
        return event; 
    }
}