import { params } from '$lib/util/params';
import type { MateReport } from '../types/report';
import type { PageLoad } from './$types';

export type Feed = { ID: number; Name: string; Desc: { String: string; Valid: boolean } };

export const prerender = false;
export const ssr = false;

export const load: PageLoad = async ({ fetch, url }) => {
	const feedId = url.searchParams.get('feed_id') ?? '';

	const [feedsRes, sightingsRes] = await Promise.all([
		fetch('/api/feeds'),
		feedId
			? fetch('/api/sightings.rss?' + params({ feed_id: feedId }))
			: Promise.resolve(null)
	]);

	const feeds: Feed[] = feedsRes.ok ? await feedsRes.json() : [];

	if (!sightingsRes?.ok) {
		return { feeds, reports: [] };
	}

	const xml = await sightingsRes.text();
	const doc = new DOMParser().parseFromString(xml, 'application/xml');
	const items = Array.from(doc.querySelectorAll('item'));

	const reports: MateReport[] = items.flatMap((item, i) => {
		const pointEl = item.getElementsByTagNameNS('http://www.georss.org/georss', 'point')[0];
		if (!pointEl) return [];

		const [lat, long] = pointEl.textContent!.trim().split(' ').map(Number);
		if (isNaN(lat) || isNaN(long)) return [];

		return [
			{
				id: i + 1,
				title: item.querySelector('title')?.textContent ?? `Sighting #${i + 1}`,
				timestamp: item.querySelector('pubDate')?.textContent ?? new Date().toISOString(),
				coordinates: [lat, long]
			}
		];
	});

	return { feeds, reports };
};
