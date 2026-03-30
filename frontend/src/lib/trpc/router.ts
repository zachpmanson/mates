import type { Context } from '$lib/trpc/context';
import { initTRPC } from '@trpc/server';
import delay from 'delay';
import { z } from 'zod';

export const t = initTRPC.context<Context>().create();

export const router = t.router({
	greeting: t.procedure.query(async () => {
		await delay(500); // 👈 simulate an expensive operation
		return `Hello tRPC v10 @ ${new Date().toLocaleTimeString()}`;
	}),
	reports: t.procedure
		.input(
			z.object({
				since: z.date(),
				type: z.enum(['tram', 'bus', 'train'])
			})
		)
		.query(async ({ ctx }) => {
			return {
				reports: [
					{
						title: 'Mates on the number 5 heading towards the city.',
						timestamp: '2024-08-06T20:00:00.000Z',
						author: 'John Doe',
						coordinates: [-37.81, 144.96]
					},
					{
						title: 'Mates waiting for the 1 or 6 city bound on Lygon st Brunswick east',
						timestamp: '2024-08-06T07:30:00.000Z',
						author: 'Holly McGrath',
						coordinates: [-37.81, 144.960005]
					},
					{
						title: 'Mates all over the place on Sydney rd tram in Brunswick today',
						timestamp: '2024-08-06T08:00:00.000Z',
						author: 'Jane Doe',
						coordinates: [-37.81, 144.960009]
					}
				]
			};
		})
});

export const createCaller = t.createCallerFactory(router);

export type Router = typeof router;
