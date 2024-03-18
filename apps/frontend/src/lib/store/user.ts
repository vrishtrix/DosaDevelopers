interface IMessage {
	username?: string;
	message?: string;
	type?: 'sent' | 'received';
}

import { writable } from 'svelte/store';

export const username = writable('');
export const messages = writable<IMessage[]>([]);
