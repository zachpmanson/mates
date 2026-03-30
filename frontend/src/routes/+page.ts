// since there's no dynamic data here, we can prerender

import type { MateReport } from '../types/report';
import type { PageLoad } from './$types';

// it so that it gets served as a static asset in production
export const prerender = false;
export const ssr = false;

export const load: PageLoad = async ({ params, url }) => {
	const reports: MateReport[] = [
		{
			id: 1,
			title: 'Mates on the number 5 heading towards the city.',
			timestamp: '2024-08-06T20:00:00.000Z',
			author: 'John Doe',
			coordinates: [-37.81, 144.96],
			type: 'tram'
		},
		{
			id: 2,
			title: 'Mates waiting for the 1 or 6 city bound on Lygon st Brunswick east',
			timestamp: '2024-08-06T07:30:00.000Z',
			author: 'Holly McGrath',
			coordinates: [-37.81, 144.97],
			type: 'tram'
		},
		{
			id: 3,
			title: 'Mates all over the place on Sydney rd tram in Brunswick today',
			timestamp: '2024-08-06T08:00:00.000Z',
			author: 'Jane Doe',
			coordinates: [-37.8, 144.960009],
			type: 'tram'
		},
		{
			id: 4,
			title: 'Mates spotted at brunswick train station',
			timestamp: '2024-08-07T08:00:00.000Z',
			author: 'Emily G',
			coordinates: [-37.79, 144.945009],
			type: 'train'
		}
	];
	await new Promise((r) => setTimeout(r, 200));
	return {
		reports: reports.sort((a, b) => (b.timestamp > a.timestamp ? 1 : -1))
	};
};
