import { User } from '../types/user';

export interface Event {
    id: number,
    title: string,
    description: string,
    location: string,
    createdAt: string,
    members: User[],
    organizers: User[],
}