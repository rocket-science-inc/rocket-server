import { Injectable } from '@graphql-modules/di';
import events from './events.json';

@Injectable()
export class EventProvider {
    getEvents() {
        return events;
    }
}