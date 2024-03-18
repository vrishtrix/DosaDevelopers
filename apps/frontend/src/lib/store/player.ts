interface ITrack {
	title: string;
	artist: string;
	file: string;
}

import { writable } from 'svelte/store';

export const audioPlayer = writable<HTMLAudioElement>();
export const status = writable<string>('default');
export const isPlaying = writable<boolean>(false);
export const index = writable<number>(0);

export const trackList = writable([
	{
		title: 'Requiem in D minor, K. 626',
		artist: 'Wolfgang Amadeus Mozart',
		file: 'https://sveltejs.github.io/assets/music/mozart.mp3',
	},
	{
		title: 'Symphony no. 5 in Cm, Op. 67',
		artist: 'Ludwig van Beethoven',
		file: 'https://sveltejs.github.io/assets/music/beethoven.mp3',
	},
]);

export const addTrack = (track: ITrack) => {
	trackList.update((v) => [...v, track]);
};
