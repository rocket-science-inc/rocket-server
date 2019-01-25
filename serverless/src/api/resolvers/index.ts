import event from './event';
import user from './user';
import query from './query';
import mutation from './mutation';

export default {
    ...event,
    ...user,
    ...query,
    ...mutation,
};
