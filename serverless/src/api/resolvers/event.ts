import { Event } from '../types/event';

export default {
    Event: {
        id: (event: Event) => event.id,
        title: (event: Event) => event.title,
        description: (event: Event) => event.description,
        location: (event: Event) => event.location,
        createdAt: (event: Event) => event.createdAt,
        members: (event: Event) => event.members,
        organizers: (event: Event) => event.organizers
    },
};