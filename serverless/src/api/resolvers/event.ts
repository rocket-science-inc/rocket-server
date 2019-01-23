export default {
    Event: {
        id: event => event.id,
        title: event => event.title,
        description: event => event.description,
        location: event => event.location,
        createdAt: event => event.createdAt,
        members: event => event.members,
        organizers: event => event.organizers
    },
};