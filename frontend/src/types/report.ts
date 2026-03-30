import type { Writable } from 'svelte/store';

export type MateReport = {
	id: number;
	title: string;
	timestamp: string;
	author: string;
	coordinates: [number, number];
	type: 'tram' | 'bus' | 'train';
};
